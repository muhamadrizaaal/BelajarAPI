package user

import (
	"BelajarAPI/helper"
	"BelajarAPI/middleware"
	"BelajarAPI/model"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController struct {
	Model model.UserModel
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

		// mencari UserID terakhir dari database
		lastUserID, err := us.Model.GetLastUserID()
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem", nil))
		}

		// menentukan UserID untuk pengguna baru
		newUserID := lastUserID + 1

		processInput := model.User{
			UserID:   newUserID,
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
		responseData.UserID = result.UserID
		responseData.Hp = result.Hp
		responseData.Nama = result.Nama
		responseData.Token = token

		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "selamat anda berhasil login", responseData))

	}
}

func (us *UserController) AddActivity() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input ActivityRequest
		err := c.Bind(&input)

		userid, _ := strconv.ParseUint(c.Param("user_id"), 10, 32)
		processInput := model.DaftarKegiatan{
			Kegiatan: input.Activity,
			UserID:   uint(userid),
		}
		err = us.Model.AddActivity(processInput, userid)

		var response ActivityResponse
		response.UserID = uint(userid)
		response.Activity = processInput.Kegiatan
		// response.ID = processInput.ID

		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem", nil))
		}
		return c.JSON(http.StatusCreated,
			helper.ResponseFormat(http.StatusCreated, "selamat data sudah terdaftar", response))
	}
}

func (us *UserController) UpdateActivity() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input ActivityRequest
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirim tidak sesuai", nil))
		}

		// Ambil user_id dari parameter URL
		userID, _ := strconv.ParseUint(c.Param("user_id"), 10, 32)
		// Ambil id kegiatan dari parameter URL
		activityID, _ := strconv.ParseUint(c.Param("id"), 10, 32)

		newData := model.DaftarKegiatan{
			Kegiatan: input.Activity,
			UserID:   uint(userID),
		}
		err = us.Model.UpdateActivity(newData, uint(userID), uint(activityID))

		response := ActivityResponse{
			UserID:   uint(userID),
			Activity: newData.Kegiatan,
		}

		if err != nil {
			log.Println("masalah database :", err.Error())
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan saat update data", nil))
		}

		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "data berhasil di update", response))
	}
}

func (us *UserController) GetAllDataById() echo.HandlerFunc {
    return func(c echo.Context) error {
        userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
        if err != nil {
            return c.JSON(http.StatusBadRequest,
                helper.ResponseFormat(http.StatusBadRequest, "user_id tidak valid", nil))
        }

        listActivity, err := us.Model.GetAllDataById(userID)
        if err != nil {
            // Jika terjadi kesalahan lain selain "record not found",
            // kembalikan respons kesalahan internal server
            if !errors.Is(err, gorm.ErrRecordNotFound) {
                return c.JSON(http.StatusInternalServerError,
                    helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem", nil))
            }
            // Jika record tidak ditemukan, kembalikan respons "not found"
            return c.JSON(http.StatusNotFound,
                helper.ResponseFormat(http.StatusNotFound, "data tidak ditemukan", nil))
        }

        // Jika hasil query kosong (tidak ada kegiatan yang ditemukan untuk user_id tersebut),
        // kembalikan respons "not found"
        if len(listActivity) == 0 {
            return c.JSON(http.StatusNotFound,
                helper.ResponseFormat(http.StatusNotFound, "data tidak ditemukan", nil))
        }

        // Jika data ditemukan, kembalikan respons OK dengan data yang ditemukan
        return c.JSON(http.StatusOK,
            helper.ResponseFormat(http.StatusOK, "berhasil mendapatkan data", listActivity))
    }
}


