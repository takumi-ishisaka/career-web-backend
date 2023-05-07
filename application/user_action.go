package application

import (
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/pkg/errors"

	"github.com/nokin-all-of-career/career-web-backend/domain"
	"github.com/nokin-all-of-career/career-web-backend/domain/repository"
	"github.com/nokin-all-of-career/career-web-backend/infra"
)

// UserActionApplication : userAction application
type UserActionApplication interface {
	TodoUserActionsGet(c echo.Context) error
	DoneUserActionsGet(c echo.Context) error
	ApprovalUserActionsGet(c echo.Context) error
	AgainUserActionsGet(c echo.Context) error
	DoneApprovalUserActionsGet(c echo.Context) error
	UserActionStatusChange(c echo.Context) error
	UserActionReflectionGet(c echo.Context) error
	UserActionReflectionInsert(c echo.Context) error
	DoneUserActionList(c echo.Context) error
}

type userActionApplication struct {
	userActionRepository repository.UserActionRepository
	actionRepository     repository.ActionRepository
	userRepository       repository.UserRepository
}

// NewUserActionApplication : create userAction application
func NewUserActionApplication(uar repository.UserActionRepository, ar repository.ActionRepository, ur repository.UserRepository) UserActionApplication {
	return &userActionApplication{
		userActionRepository: uar,
		actionRepository:     ar,
		userRepository:       ur,
	}
}

func (uaa userActionApplication) TodoUserActionsGet(c echo.Context) error {
	// get userID from JWT
	var userID string
	if token := c.Get("user"); token != nil {
		user := token.(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userID = claims["sub"].(string)
	} else {
		err := errors.New("not *jwt.Token")
		return errors.Wrap(err, "profileApplication.TodoUserActionsGet()")
	}

	userActions, err := uaa.userActionRepository.SelectByUserIDAndStatus(userID, domain.STATETODO)
	if err != nil {
		return errors.Wrap(err, "userActionApplication.TodoUserActionsGet()")
	}

	var todoActionIDs []interface{}
	for _, v := range userActions {
		todoActionIDs = append(todoActionIDs, v.ActionID)
	}
	// get user todo actions detail
	actions, err := uaa.actionRepository.SelectByPrimaryKeys(todoActionIDs)
	if err != nil {
		return errors.Wrap(err, "actionApplication.TodoUserActionsGet()")
	}

	// get recommendations of user todo actions
	recommendations, err := uaa.userActionRepository.GetRecommendationsMapByActionID(todoActionIDs)
	if err != nil {
		return errors.Wrap(err, "actionApplication.TodoUserActionsGet()")
	}

	useractionDetailShowsResponse := []UserActionDetailShowResponse{}
	for _, v := range userActions {
		useractionDetailShowResponse := UserActionDetailShowResponse{
			UserActionID:        v.UserActionID,
			Action:              actions[v.ActionID],
			CategoryName:        actions[v.ActionID].CategoryID,
			Status:              v.Status,
			RecommendationValue: recommendations[v.ActionID].RecommendationValue,
			DoneUserNum:         recommendations[v.ActionID].DoneUserNum,
		}

		useractionDetailShowsResponse = append(useractionDetailShowsResponse, useractionDetailShowResponse)

	}

	return c.JSON(http.StatusOK, UserActionDetailShowsResponse{UserActionDetailShowsResponse: useractionDetailShowsResponse})
}

func (uaa userActionApplication) DoneUserActionsGet(c echo.Context) error {
	// get userID from JWT
	var userID string
	if token := c.Get("user"); token != nil {
		user := token.(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userID = claims["sub"].(string)
	} else {
		err := errors.New("not *jwt.Token")
		return errors.Wrap(err, "profileApplication.DoneUserActionsGet()")
	}

	userActions, err := uaa.userActionRepository.SelectByUserIDAndStatus(userID, domain.STATEDONE)
	if err != nil {
		return errors.Wrap(err, "userActionApplication.DoneUserActionsGet()")
	}

	var doneActionIDs []interface{}
	for _, v := range userActions {
		doneActionIDs = append(doneActionIDs, v.ActionID)
	}

	// get user done actions
	actions, err := uaa.actionRepository.SelectByPrimaryKeys(doneActionIDs)
	if err != nil {
		return errors.Wrap(err, "actionApplication.DoneUserActionsGet()")
	}

	// get recommendations of user done acitons
	recommendations, err := uaa.userActionRepository.GetRecommendationsMapByActionID(doneActionIDs)
	if err != nil {
		return errors.Wrap(err, "actionApplication.DoneUserActionsGet()")
	}

	useractionDetailShowsResponse := []UserActionDetailShowResponse{}
	for _, v := range userActions {
		useractionDetailShowResponse := UserActionDetailShowResponse{
			UserActionID:        v.UserActionID,
			Action:              actions[v.ActionID],
			CategoryName:        actions[v.ActionID].CategoryID,
			Status:              v.Status,
			RecommendationValue: recommendations[v.ActionID].RecommendationValue,
			DoneUserNum:         recommendations[v.ActionID].DoneUserNum,
		}

		useractionDetailShowsResponse = append(useractionDetailShowsResponse, useractionDetailShowResponse)

	}

	return c.JSON(http.StatusOK, UserActionDetailShowsResponse{UserActionDetailShowsResponse: useractionDetailShowsResponse})
}

func (uaa userActionApplication) ApprovalUserActionsGet(c echo.Context) error {
	// get userID from JWT
	var userID string
	if token := c.Get("user"); token != nil {
		user := token.(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userID = claims["sub"].(string)
	} else {
		err := errors.New("not *jwt.Token")
		return errors.Wrap(err, "profileApplication.ApprovalUserActionsGet()")
	}

	userActions, err := uaa.userActionRepository.SelectByUserIDAndStatus(userID, domain.STATEAPPROVAL)
	if err != nil {
		return errors.Wrap(err, "userActionApplication.ApprovalUserActionsGet()")
	}

	var approvalActionIDs []interface{}
	for _, v := range userActions {
		approvalActionIDs = append(approvalActionIDs, v.ActionID)
	}

	// get approval user actions
	actions, err := uaa.actionRepository.SelectByPrimaryKeys(approvalActionIDs)
	if err != nil {
		return errors.Wrap(err, "actionApplication.ApprovalUserActionsGet()")
	}

	// get recommendations of approval user actions
	recommendations, err := uaa.userActionRepository.GetRecommendationsMapByActionID(approvalActionIDs)
	if err != nil {
		return errors.Wrap(err, "actionApplication.ApprovalUserActionsGet()")
	}

	useractionDetailShowsResponse := []UserActionDetailShowResponse{}
	for _, v := range userActions {
		useractionDetailShowResponse := UserActionDetailShowResponse{
			UserActionID:        v.UserActionID,
			Action:              actions[v.ActionID],
			CategoryName:        actions[v.ActionID].CategoryID,
			Status:              v.Status,
			RecommendationValue: recommendations[v.ActionID].RecommendationValue,
			DoneUserNum:         recommendations[v.ActionID].DoneUserNum,
		}
		useractionDetailShowsResponse = append(useractionDetailShowsResponse, useractionDetailShowResponse)

	}

	return c.JSON(http.StatusOK, UserActionDetailShowsResponse{UserActionDetailShowsResponse: useractionDetailShowsResponse})
}

// AgainUserActionsGet : get action the status is "again"
func (uaa userActionApplication) AgainUserActionsGet(c echo.Context) error {
	// get userID from JWT
	var userID string
	if token := c.Get("user"); token != nil {
		user := token.(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userID = claims["sub"].(string)
	} else {
		err := errors.New("not *jwt.Token")
		return errors.Wrap(err, "userActionApplication.AgainUserActionsGet()")
	}

	userActions, err := uaa.userActionRepository.SelectByUserIDAndStatus(userID, domain.STATEAGAIN)
	if err != nil {
		return errors.Wrap(err, "userActionApplication.AgainUserActionsGet()")
	}

	var againActionIDs []interface{}
	for _, v := range userActions {
		againActionIDs = append(againActionIDs, v.ActionID)
	}

	// get again user actions
	actions, err := uaa.actionRepository.SelectByPrimaryKeys(againActionIDs)
	if err != nil {
		return errors.Wrap(err, "userActionApplication.AgainUserActionsGet()")
	}

	// get recommendations of again user actions
	recommendations, err := uaa.userActionRepository.GetRecommendationsMapByActionID(againActionIDs)
	if err != nil {
		return errors.Wrap(err, "userActionApplication.AgainUserActionsGet()")
	}

	useractionDetailShowsResponse := []UserActionDetailShowResponse{}
	for _, v := range userActions {
		useractionDetailShowResponse := UserActionDetailShowResponse{
			UserActionID:        v.UserActionID,
			Action:              actions[v.ActionID],
			CategoryName:        actions[v.ActionID].CategoryID,
			Status:              v.Status,
			RecommendationValue: recommendations[v.ActionID].RecommendationValue,
			DoneUserNum:         recommendations[v.ActionID].DoneUserNum,
		}
		useractionDetailShowsResponse = append(useractionDetailShowsResponse, useractionDetailShowResponse)

	}

	return c.JSON(http.StatusOK, UserActionDetailShowsResponse{UserActionDetailShowsResponse: useractionDetailShowsResponse})
}

func (uaa userActionApplication) DoneApprovalUserActionsGet(c echo.Context) error {
	// get userID from JWT
	var userID string
	if token := c.Get("user"); token != nil {
		user := token.(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userID = claims["sub"].(string)
	} else {
		err := errors.New("not *jwt.Token")
		return errors.Wrap(err, "userActionApplication.DoneApprovalUserActionsGet()")
	}

	doneUserActions, err := uaa.userActionRepository.SelectByUserIDAndStatus(userID, domain.STATEDONE)
	if err != nil {
		return errors.Wrap(err, "userActionApplication.DoneApprovalUserActionsGet()")
	}
	// fmt.Println(doneUserActions[0].UserActionID)
	approvalUserActions, err := uaa.userActionRepository.SelectByUserIDAndStatus(userID, domain.STATEAPPROVAL)
	if err != nil {
		return errors.Wrap(err, "userActionApplication.DoneApprovalUserActionsGet()")
	}

	var doneApprovalUserActions []*domain.UserAction
	var doneApprovalActionIDs []interface{}
	for _, v := range doneUserActions {
		doneApprovalActionIDs = append(doneApprovalActionIDs, v.ActionID)
		doneApprovalUserActions = append(doneApprovalUserActions, v)
	}

	for _, v := range approvalUserActions {
		doneApprovalActionIDs = append(doneApprovalActionIDs, v.ActionID)
		doneApprovalUserActions = append(doneApprovalUserActions, v)
	}

	// get done and approval user actions detail
	actions, err := uaa.actionRepository.SelectByPrimaryKeys(doneApprovalActionIDs)
	if err != nil {
		return errors.Wrap(err, "actionApplication.DoneApprovalUserActionsGet()")
	}

	// get recommendations of done and approval user actions
	recommendations, err := uaa.userActionRepository.GetRecommendationsMapByActionID(doneApprovalActionIDs)
	if err != nil {
		return errors.Wrap(err, "actionApplication.DoneApprovalUserActionsGet()")
	}

	useractionDetailShowsResponse := []UserActionDetailShowResponse{}
	for _, v := range doneApprovalUserActions {
		useractionDetailShowResponse := UserActionDetailShowResponse{
			UserActionID:        v.UserActionID,
			Action:              actions[v.ActionID],
			CategoryName:        actions[v.ActionID].CategoryID,
			Status:              v.Status,
			RecommendationValue: recommendations[v.ActionID].RecommendationValue,
			DoneUserNum:         recommendations[v.ActionID].DoneUserNum,
		}
		useractionDetailShowsResponse = append(useractionDetailShowsResponse, useractionDetailShowResponse)
	}

	return c.JSON(http.StatusOK, UserActionDetailShowsResponse{UserActionDetailShowsResponse: useractionDetailShowsResponse})
}

func (uaa userActionApplication) UserActionStatusChange(c echo.Context) error {
	// get userID from JWT
	var userID string
	if token := c.Get("user"); token != nil {
		user := token.(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userID = claims["sub"].(string)
	} else {
		err := errors.New("not *jwt.Token")
		return errors.Wrap(err, "profileApplication.UserActionStatusChange()")
	}

	userActionRequest := new(UserActionRequest)
	if err := c.Bind(userActionRequest); err != nil {
		return errors.Wrap(err, "userActionApplication.UserActionStatusChange()")
	}

	if userActionRequest.Status == domain.NOTSELECT {
		err := uaa.userActionRepository.DeleteUserAction(userID, userActionRequest.ActionID)
		if err != nil {
			return errors.Wrap(err, "userActionApplication.UserActionStatusChange()")
		}
	} else if userActionRequest.Status == domain.TODO {
		userActionID, err := domain.GenerateUserActionID()
		if err != nil {
			return errors.Wrap(err, "userActionApplication.UserActionStatusChange()")
		}
		err = uaa.userActionRepository.InsertUserAction(userActionID, userID, userActionRequest.ActionID, domain.STATETODO)
		if err != nil {
			return errors.Wrap(err, "userActionApplication.UserActionStatusChange()")
		}
	} else if userActionRequest.Status == domain.DONE {
		err := uaa.userActionRepository.UpdateUserAction(userID, userActionRequest.ActionID, domain.STATEDONE)
		if err != nil {
			return errors.Wrap(err, "userActionApplication.UserActionStatusChange()")
		}
	}

	log.SetOutput(os.Stdout)
	log.Println(userActionRequest.Status)
	return nil
}

// UserActionDetailShowResponse : user action detail response
type UserActionDetailShowResponse struct {
	UserActionID        string        `json:"user_action_id"`
	Action              domain.Action `json:"action"`
	CategoryName        string        `json:"category_name"`
	Status              int           `json:"status"`
	RecommendationValue float32       `json:"recommendation_value"`
	DoneUserNum         int           `json:"done_user_num"`
}

// UserActionDetailShowsResponse : user action detail slice response
type UserActionDetailShowsResponse struct {
	UserActionDetailShowsResponse []UserActionDetailShowResponse `json:"user_action_detail_shows_response"`
}

// UserActionRequest : userAction request struct
type UserActionRequest struct {
	ActionID string `json:"action_id"`
	Status   string `json:"status"`
}

// UserActionResponse : userAction response struct
type UserActionResponse struct {
	UserActions []*domain.UserAction `json:"user_actions"`
}

func (uaa userActionApplication) UserActionReflectionGet(c echo.Context) error {
	// get userID from JWT
	var userID string
	if token := c.Get("user"); token != nil {
		user := token.(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userID = claims["sub"].(string)
	} else {
		err := errors.New("not *jwt.Token")
		return errors.Wrap(err, "profileApplication.UserActionReflectionGet()")
	}

	reflectionGetRequest := new(ReflectionGetRequest)
	if err := c.Bind(reflectionGetRequest); err != nil {
		return errors.Wrap(err, "userActionApplication.UserActionReflectionGet()")
	}
	reflection, err := uaa.userActionRepository.SelectReflectionByPrimaryKey(userID, reflectionGetRequest.ActionID)
	if err != nil {
		return errors.Wrap(err, "userActionApplication.UserActionReflectionGet()")
	}

	return c.JSON(http.StatusOK, reflection)
}

func (uaa userActionApplication) UserActionReflectionInsert(c echo.Context) error {
	// get userID from JWT
	var userID string
	if token := c.Get("user"); token != nil {
		user := token.(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userID = claims["sub"].(string)
	} else {
		err := errors.New("not *jwt.Token")
		return errors.Wrap(err, "profileApplication.UserActionReflectionInsert()")
	}

	// get form data
	reflectionRequest := new(ReflectionRequest)
	if err := c.Bind(reflectionRequest); err != nil {
		return errors.Wrap(err, "userActionApplication.UserActionReflectionInsert()")
	}
	actionID := c.Param("action_id")
	if actionID == "" {
		return errors.New("userActionApplication.UserActionReflectionInsert() : could not get action_id")
	}

	err := domain.ReflectionValidation(reflectionRequest.Do, reflectionRequest.Reflection, reflectionRequest.NextAction)
	if err != nil {
		return errors.Wrap(err, "userActionApplication.UserActionReflectionInsert()")
	}

	// insert form data to user_action table
	err = uaa.userActionRepository.InsertReflection(userID, actionID, reflectionRequest.Do, reflectionRequest.Reflection, reflectionRequest.NextAction, reflectionRequest.EvaluateValue)
	if err != nil {
		return errors.Wrap(err, "userActionApplication.UserActionReflectionInsert()")
	}
	action, err := uaa.actionRepository.SelectByPrimaryKey(actionID)
	if err != nil {
		return errors.Wrap(err, "userActionApplication.UserActionReflectionInsert()")
	}
	user, err := uaa.userRepository.SelectByPrimaryKey(userID)
	if err != nil {
		return errors.Wrap(err, "userActionApplication.UserActionReflectionInsert()")
	}
	err = uaa.userActionRepository.SendEmailToOperation(infra.OperationEmail, infra.OperationEmail, action.Title, user.Email, reflectionRequest.Do, reflectionRequest.Reflection, reflectionRequest.NextAction)
	if err != nil {
		return errors.Wrap(err, "userActionApplication.UserActionReflectionInsert()")
	}
	return nil
}

func (uaa userActionApplication) DoneUserActionList(c echo.Context) error {
	userActions, err := uaa.userActionRepository.SelectUserActionsOfAllUsers(domain.STATEDONE)
	if err != nil {
		return errors.Wrap(err, "userActionApplication.AllUserDoneUserActionList()")
	}
	var actionIDs []interface{}
	for _, v := range userActions {
		actionIDs = append(actionIDs, v.ActionID)
	}
	actions, err := uaa.actionRepository.SelectByPrimaryKeys(actionIDs)

	var allUserDoneUserAction DoneUserAction
	var allUserDoneUserActions []DoneUserAction
	for _, v := range userActions {
		allUserDoneUserAction = DoneUserAction{
			UserAction: *v,
			Action:     actions[v.ActionID],
		}
		allUserDoneUserActions = append(allUserDoneUserActions, allUserDoneUserAction)
	}

	return c.JSON(http.StatusOK, DoneUserActionsRequest{DoneUserActions: allUserDoneUserActions})
}

// ReflectionGetRequest : reflection get request struct
type ReflectionGetRequest struct {
	ActionID string `query:"action_id"`
}

// DoneUserAction : done user action
type DoneUserAction struct {
	UserAction domain.UserAction `json:"user_action" form:"user_action"`
	Action     domain.Action     `json:"action" form:"action"`
}

// DoneUserActionsRequest : done user action slice request
type DoneUserActionsRequest struct {
	DoneUserActions []DoneUserAction `json:"done_user_actions"`
}

// ReflectionRequest : reflection request struct
type ReflectionRequest struct {
	Do            string `form:"do"`
	Reflection    string `form:"reflection"`
	NextAction    string `form:"next_action"`
	EvaluateValue int    `form:"evaluate_value"`
}
