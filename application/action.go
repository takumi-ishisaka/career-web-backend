package application

import (
	"net/http"
	"sort"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/pkg/errors"

	"github.com/nokin-all-of-career/career-web-backend/domain"
	"github.com/nokin-all-of-career/career-web-backend/domain/repository"
)

// ActionApplication : action application
type ActionApplication interface {
	ActionsShow(c echo.Context) error
	ActionDetailShow(c echo.Context) error
}

type actionApplication struct {
	actionRepository     repository.ActionRepository
	useractionRepository repository.UserActionRepository
	categoryRepository   repository.CategoryRepository
}

// NewActionApplication : create application about Action
func NewActionApplication(ar repository.ActionRepository, ur repository.UserActionRepository, cr repository.CategoryRepository) ActionApplication {
	return &actionApplication{
		actionRepository:     ar,
		useractionRepository: ur,
		categoryRepository:   cr,
	}
}

func (aa *actionApplication) ActionsShow(c echo.Context) error {
	// get userID from JWT
	var userID string
	if token := c.Get("user"); token != nil {
		user := token.(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userID = claims["sub"].(string)
	} else {
		err := errors.New("not *jwt.Token")
		return errors.Wrap(err, "actionApplication.ActionsShow()")
	}

	categoryRequest := new(CategoryRequest)
	if err := c.Bind(categoryRequest); err != nil {
		return errors.Wrap(err, "actionApplication.ActionsShow()")
	}
	// get disignated category actions
	actions, err := aa.actionRepository.SelectByCategoryID(categoryRequest.CategoryID)
	if err != nil {
		return errors.Wrap(err, "actionApplication.ActionsShow()")
	}

	doneAction, err := aa.useractionRepository.SelectByUserIDAndStatus(userID, domain.STATEDONE)
	todoAction, err := aa.useractionRepository.SelectByUserIDAndStatus(userID, domain.STATETODO)
	var actionIDs []interface{}
	responseActions := map[string]domain.Action{}

	// recognize not todo and not done action
	for i := 0; i < len(actions); i++ {
		userActionExistFlag := 0
		for j := 0; j < len(doneAction); j++ {
			if doneAction[j].ActionID == actions[i].ActionID {
				userActionExistFlag = 1
				// fmt.Println("doneAction")
				break
			}
		}
		if userActionExistFlag != 1 {
			for j := 0; j < len(todoAction); j++ {
				if todoAction[j].ActionID == actions[i].ActionID {
					userActionExistFlag = 1
					// fmt.Println("todoAction")
					break
				}
			}
		}
		if userActionExistFlag != 1 {
			actionIDs = append(actionIDs, actions[i].ActionID)
			responseActions[actions[i].ActionID] = actions[i]
		}
	}
	sort.Slice(actionIDs, func(i, j int) bool { return actionIDs[i].(string) < actionIDs[j].(string) })
	// fmt.Println(actionIDs)
	recommendations, err := aa.useractionRepository.GetRecommendationsByActionID(actionIDs)
	if err != nil {
		return errors.Wrap(err, "actionApplication.ActionsShow()")
	}

	// create response
	var actionsShowResponse ActionsShowResponse
	for _, actionID := range actionIDs {
		// データがrecommendationテーブルにあるかどうかを示す変数
		existOnRecommendationFlag := 0
		if len(recommendations) != 0 {
			for i := 0; i < len(recommendations); i++ {
				//もし、recommendations変数に該当actionIDがあったら、該当データをいれる
				if actionID == recommendations[i].ActionID {
					actionsShowResponse.Actions = append(actionsShowResponse.Actions, ActionShowResponse{
						Action:              responseActions[actionID.(string)],
						RecommendationValue: recommendations[i].RecommendationValue,
						DoneUserNum:         recommendations[i].DoneUserNum,
					})
					existOnRecommendationFlag = 1
					break
				}
			}
		}
		// 該当recommendationRowがなかったら、初期値を入れる
		if existOnRecommendationFlag != 1 {
			actionsShowResponse.Actions = append(actionsShowResponse.Actions, ActionShowResponse{
				Action:              responseActions[actionID.(string)],
				RecommendationValue: 3,
				DoneUserNum:         0,
			})
		}
	}

	return c.JSON(http.StatusOK, actionsShowResponse)
}

func (aa *actionApplication) ActionDetailShow(c echo.Context) error {
	actionRequest := new(ActionRequest)
	if err := c.Bind(actionRequest); err != nil {
		return errors.Wrap(err, "actionApplication.ActionDetailShow()")
	}

	action, err := aa.actionRepository.SelectByPrimaryKey(actionRequest.ActionID)
	if err != nil {
		return errors.Wrap(err, "actionApplication.ActionDetailShow()")
	}

	category, err := aa.categoryRepository.SelectByPrimaryKey(action.CategoryID)
	if err != nil {
		return errors.Wrap(err, "actionApplication.ActionDetailShow()")
	}

	recommendation, err := aa.useractionRepository.GetRecommendationByActionID(actionRequest.ActionID)
	if err != nil {
		return errors.Wrap(err, "actionApplication.ActionDetailShow()")
	}

	var actionDetailShowResponse = ActionDetailShowResponse{
		Action:              action,
		CategoryName:        category.Name,
		RecommendationValue: recommendation.RecommendationValue,
		DoneUserNum:         recommendation.DoneUserNum,
	}

	return c.JSON(http.StatusOK, actionDetailShowResponse)
}

// CategoryRequest : category request
type CategoryRequest struct {
	CategoryID string `query:"category_id"`
}

// ActionShowResponse : action show response
type ActionShowResponse struct {
	Action              domain.Action `json:"action"`
	RecommendationValue float32       `json:"recommendation_value"`
	DoneUserNum         int           `json:"done_user_num"`
}

// ActionsShowResponse : actions show response
type ActionsShowResponse struct {
	Actions []ActionShowResponse `json:"actions"`
}

// ActionRequest : action request
type ActionRequest struct {
	ActionID string `query:"action_id"`
}

// ActionDetailShowResponse : action detail show response
type ActionDetailShowResponse struct {
	Action              domain.Action `json:"action"`
	CategoryName        string        `json:"category_name"`
	Status              int           `json:"status"`
	RecommendationValue float32       `json:"recommendation_value"`
	DoneUserNum         int           `json:"done_user_num"`
}

// ActionDetailShowsResponse : export to useraction
type ActionDetailShowsResponse struct {
	ActionDetailShowResponse []ActionDetailShowResponse `json:"actionDetailShowResponse"`
}
