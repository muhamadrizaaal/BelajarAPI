package todo

import (
	"BelajarAPI/helper"
	"BelajarAPI/middleware"
	"BelajarAPI/model/todo"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TodoController struct {
	Model todo.TodoModel
}

// func (tc *TodoController) AddActivity() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var input ActivityRequest
// 		err := c.Bind(&input)
// 		if err != nil {
// 			log.Println("error bind data:", err.Error())
// 			if strings.Contains(err.Error(), "unsupport") {
// 				return c.JSON(http.StatusUnsupportedMediaType,
// 					helper.ResponseFormat(http.StatusUnsupportedMediaType, "format data tidak didukung", nil))
// 			}
// 			return c.JSON(http.StatusBadRequest,
// 				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirmkan tidak sesuai", nil))
// 		}

// 		user_id := middleware.DecodeToken(c.Get("user").(*jwt.Token))

// 		if user_id == "" {
// 			log.Println("error decode token:", "hp tidak ditemukan")
// 			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "tidak dapat mengakses fitur ini", nil))
// 		}

// 		// userid, _ := strconv.ParseUint(c.Param("user_id"), 10, 32)
// 		processInput := todo.Todo{
// 			Kegiatan: input.Activity,
// 			UserID:   input.UserID,
// 		}
// 		err = tc.Model.AddActivity(processInput)

// 		var response ActivityResponse
// 		response.UserID = processInput.UserID
// 		response.Activity = processInput.Kegiatan

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError,
// 				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem", nil))
// 		}
// 		return c.JSON(http.StatusCreated,
// 			helper.ResponseFormat(http.StatusCreated, "selamat data sudah terdaftar", response))
// 	}
// }

func (tc *TodoController) AddActivity() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input ActivityRequest
		err := c.Bind(&input)
		if err != nil {
			log.Println("error bind data:", err.Error())
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, "format data tidak didukung", nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirmkan tidak sesuai", nil))
		}

		hp := middleware.DecodeToken(c.Get("user").(*jwt.Token))

		if hp == "" {
			log.Println("error decode token:", "hp tidak ditemukan")
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "tidak dapat mengakses fitur ini", nil))
		}

		log.Println(hp)

		processInput := todo.Todo{
			Kegiatan: input.Activity,
			Hp:       hp,
		}

		err = tc.Model.AddActivity(processInput)

		var response ActivityResponse
		response.Hp = processInput.Hp
		response.Activity = processInput.Kegiatan

		if err != nil {
			log.Println("error insert db:", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada proses server", nil))
		}

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "berhasil menambahkan kegiatan", response))

	}
}

func (tc *TodoController) UpdateActivity() echo.HandlerFunc {
	return func(c echo.Context) error {
		// // Ambil user_id dari parameter URL
		// userID, _ := strconv.ParseUint(c.Param("user_id"), 10, 32)
		// Ambil id kegiatan dari parameter URL
		activityID, err := strconv.ParseUint(c.Param("id"), 10, 32)

		if err != nil {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirmkan tidak sesuai", nil))
		}

		var input ActivityRequest
		err = c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirim tidak sesuai", nil))
		}

		hp := middleware.DecodeToken(c.Get("user").(*jwt.Token))

		if hp == "" {
			log.Println("error decode token:", "hp tidak ditemukan")
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "tidak dapat mengakses fitur ini", nil))
		}

		newData := todo.Todo{
			Kegiatan: input.Activity,
		}
		err = tc.Model.UpdateActivity(newData, uint(activityID), hp)

		response := ActivityResponse{
			Hp:       hp,
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

func (tc *TodoController) GetAllDataById() echo.HandlerFunc {
	return func(c echo.Context) error {
		hp := middleware.DecodeToken(c.Get("user").(*jwt.Token))

		if hp == "" {
			log.Println("error decode token:", "hp tidak ditemukan")
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "tidak dapat mengakses fitur ini", nil))
		}

		listActivity, err := tc.Model.GetAllDataById(hp)
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
