package application

import (
	"database/sql"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/nokin-all-of-career/career-web-backend/domain"
	"github.com/nokin-all-of-career/career-web-backend/domain/repository"
	"github.com/nokin-all-of-career/career-web-backend/infra"
	"github.com/pkg/errors"
)

// FeedBackApplication : feedback application
type FeedBackApplication interface {
	FeedBackInsert(c echo.Context) error
	FeedBacksGet(c echo.Context) error
	FeedBackList(c echo.Context) error
}

type feedbackApplication struct {
	feedbackRepository   repository.FeedBackRepository
	userActionRepository repository.UserActionRepository
	actionRepository     repository.ActionRepository
	userRepository       repository.UserRepository
}

// NewFeedBackApplication : create feedback application
func NewFeedBackApplication(fr repository.FeedBackRepository, uar repository.UserActionRepository, ar repository.ActionRepository, ur repository.UserRepository) FeedBackApplication {
	return &feedbackApplication{
		feedbackRepository:   fr,
		userActionRepository: uar,
		actionRepository:     ar,
		userRepository:       ur,
	}
}

func (fa feedbackApplication) FeedBackInsert(c echo.Context) error {

	var userID string
	if token := c.Get("user"); token != nil {
		user := token.(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userID = claims["sub"].(string)
	} else {
		err := errors.New("not *jwt.Token")
		return errors.Wrap(err, "feedbackApplication.FeedBackInsert()")
	}

	feedbackInsertRequest := new(FeedBackInsertRequest)
	if err := c.Bind(feedbackInsertRequest); err != nil {
		return errors.Wrap(err, "feedbackApplication.FeedBackInsert()")
	}

	userAction, err := fa.userActionRepository.SelectByPrimaryKey(feedbackInsertRequest.UserActionID)
	if err != nil {
		return errors.Wrap(err, "feedbackApplication.FeedBackInsert()")
	}
	action, err := fa.actionRepository.SelectByPrimaryKey(userAction.ActionID)
	user, err := fa.userRepository.SelectByPrimaryKey(userID)

	feedbackID, err := domain.GenerateFeedBackID()

	// TODO:トランザクション
	err = Transact(infra.DB, func(tx *sql.Tx) error {
		//Insert
		err = fa.feedbackRepository.InsertFeedBack(tx, feedbackID, feedbackInsertRequest.UserActionID, feedbackInsertRequest.Comment)
		if err != nil {
			return errors.Wrap(err, "feedbackApplication.FeedBackInsert()")
		}
		// fmt.Println(feedbackInsertRequest.Approval)
		// change user_action status by approval
		if feedbackInsertRequest.Approval == 1 {
			err := fa.userActionRepository.UpdateUserActionByUserActionID(tx, feedbackInsertRequest.UserActionID, domain.STATEAPPROVAL)
			if err != nil {
				return errors.Wrap(err, "feedbackApplication.FeedBackInsert()")
			}
		}
		if feedbackInsertRequest.Approval == 0 {
			err := fa.userActionRepository.UpdateUserActionByUserActionID(tx, feedbackInsertRequest.UserActionID, domain.STATEAGAIN)
			if err != nil {
				return errors.Wrap(err, "feedbackApplication.FeedBackInsert()")
			}
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "feedbackApplication.FeedBackInsert()")
	}
	// 通知メールの送信
	err = fa.feedbackRepository.SendEmailToUser(infra.OperationEmail, user.Email, action.Title, feedbackInsertRequest.Comment)
	if err != nil {
		return errors.Wrap(err, "feedbackApplication.FeedBackInsert()")
	}
	return nil
}

// func (fa feedbackApplication) FeedbackTest(c echo.Context) error {
// 	err := fa.feedbackRepository.SendEmailToUser("allofcareer@gmail.com", "yuukiyuuki327@gmail.com", "あああ")
// 	if err != nil {
// 		return errors.Wrap(err, "feedbackApplication.FeedbackTest")
// 	}
// 	return nil
// }

func (fa feedbackApplication) FeedBacksGet(c echo.Context) error {
	feedBackGetRequest := new(FeedBackGetRequest)
	if err := c.Bind(feedBackGetRequest); err != nil {
		return errors.Wrap(err, "feedbackApplication.FeedBacksGet()")
	}
	feedBacks, err := fa.feedbackRepository.SelectByPrimaryKey(feedBackGetRequest.UserActionID)
	if err != nil {
		return errors.Wrap(err, "feedbackApplication.FeedBacksGet()")
	}
	return c.JSON(http.StatusOK, GetFeedBacksResponse{GetFeedBacksResponse: feedBacks})
}

func (fa feedbackApplication) FeedBackList(c echo.Context) error {
	var userID string
	if token := c.Get("user"); token != nil {
		user := token.(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userID = claims["sub"].(string)
	} else {
		err := errors.New("not *jwt.Token")
		return errors.Wrap(err, "feedbackApplication.FeedBackList()")
	}

	TodoUserAction, err := fa.userActionRepository.SelectByUserIDAndStatus(userID, domain.STATETODO)
	ApprovalUserAction, err := fa.userActionRepository.SelectByUserIDAndStatus(userID, domain.STATEAPPROVAL)
	var userActionID []string
	for _, userAction := range TodoUserAction {
		userActionID = append(userActionID, userAction.UserActionID)
	}
	for _, userAction := range ApprovalUserAction {
		userActionID = append(userActionID, userAction.UserActionID)
	}
	feedBacks, err := fa.feedbackRepository.SelectByPrimaryKeys(userActionID)
	if err != nil {
		return errors.Wrap(err, "feedbackApplication.FeedBackList()")
	}
	return c.JSON(http.StatusOK, GetFeedBacksResponse{GetFeedBacksResponse: feedBacks})
}

// FeedBackInsertRequest : feedback/insert request
type FeedBackInsertRequest struct {
	UserActionID string `json:"user_action_id" form:"user_action_id"`
	Comment      string `json:"comment" form:"comment"`
	Approval     int    `json:"approval" form:"approval"`
}

// FeedBackGetRequest : feedback get request
type FeedBackGetRequest struct {
	UserActionID string `query:"user_action_id" json:"user_action_id"`
}

// GetFeedBacksResponse : get feedback response
type GetFeedBacksResponse struct {
	GetFeedBacksResponse *[]domain.FeedBack `json:"get_feedbacks_response"`
}

// Transact : transaction
func Transact(db *sql.DB, txFunc func(*sql.Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = txFunc(tx)
	return err
}
