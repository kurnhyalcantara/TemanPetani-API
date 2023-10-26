package controller

import (
	"github.com/kurnhyalcantara/TemanPetani-API/apis/users/model"
	"github.com/kurnhyalcantara/TemanPetani-API/apis/users/service"
	"github.com/kurnhyalcantara/TemanPetani-API/app/utils"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService  service.UserServiceInterface
	responseJson utils.ResponseJSON
}

// PostUserController implements UserControllerInterface
func (controller *UserController) PostUserController(c echo.Context) error {
	payload := model.CreateUser{}
	if errBind := c.Bind(&payload); errBind != nil {
		return controller.responseJson.StatusBadRequestResponse(c, errBind.Error())
	}
	return nil
}

func New(userService service.UserServiceInterface) UserControllerInterface {
	return &UserController{
		userService:  userService,
		responseJson: *utils.NewResponseJSON(),
	}
}
