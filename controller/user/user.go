package user

import (
	"BelajarAPI/helper"
	"BelajarAPI/middleware"
	"BelajarAPI/model/user"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	Model user.UserModel
}

func (us *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterRequest
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, "format data tidak didukung", nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirmkan tidak sesuai", nil))
		}

		validate := validator.New(validator.WithRequiredStructEnabled())
		err = validate.Struct(input)

		if err != nil {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirim kurang sesuai", nil))
		}

		// // mencari UserID terakhir dari database
		// lastUserID, err := us.Model.GetLastUserID()
		// if err != nil {
		// 	return c.JSON(http.StatusInternalServerError,
		// 		helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem", nil))
		// }

		// // menentukan UserID untuk pengguna baru
		// newUserID := lastUserID + 1

		processInput := user.User{
			// UserID:   newUserID,
			Hp:       input.Hp,
			Nama:     input.Nama,
			Password: input.Password,
		}

		err = us.Model.Register(processInput)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem", nil))
		}
		return c.JSON(http.StatusCreated,
			helper.ResponseFormat(http.StatusCreated, "selamat data sudah terdaftar", nil))
	}
}

func (us *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginRequest
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, "format data tidak didukung", nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirmkan tidak sesuai", nil))
		}

		validate := validator.New(validator.WithRequiredStructEnabled())
		err = validate.Struct(input)

		if err != nil {
			for _, val := range err.(validator.ValidationErrors) {
				fmt.Println(val.Error())
			}
		}

		result, err := us.Model.Login(input.Hp, input.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem", nil))
		}
		token, err := middleware.GenerateJWT(result.Hp)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem, gagal memproses data", nil))
		}

		var responseData LoginResponse
		// responseData.UserID = result.UserID
		responseData.Hp = result.Hp
		responseData.Nama = result.Nama
		responseData.Token = token

		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "selamat anda berhasil login", responseData))

	}
}
