package handler

import (
	"BelajarAPI/features/todo"
	"BelajarAPI/helper"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type controller struct {
	s todo.TodoService
}

func NewHandler(service todo.TodoService) todo.TodoController {
	return &controller{
		s: service,
	}
}

func (ct *controller) AddActivity() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input TodoRequest
		err := c.Bind(&input)
		if err != nil {
			log.Println("error bind data:", err.Error())
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.UserInputFormatError, nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}
		var inputProcess todo.Todo
		inputProcess.Kegiatan = input.Kegiatan
		result, err := ct.s.AddActivity(token, inputProcess)
		if err != nil {
			log.Println("error insert db:", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
		}

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "berhasil menambahkan kegiatan", result))
	}
}

func (ct *controller) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		activityID, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		var input TodoRequest
		err = c.Bind(&input)
		if err != nil {
			log.Println("error bind data:", err.Error())
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.UserInputFormatError, nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		var updateProcess todo.Todo
		updateProcess.Kegiatan = input.Kegiatan
		err = ct.s.UpdateActivity(updateProcess, uint(activityID), token)
		if err != nil {
			log.Println("error insert db:", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
		}

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "berhasil merubah kegiatan", nil))
	}
}

func (ct *controller) GetData() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		listActivity, err := ct.s.GetAllDataById(token)
		if err != nil {
			// Jika terjadi kesalahan lain selain "record not found",
			// kembalikan respons kesalahan internal server
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return c.JSON(http.StatusInternalServerError,
					helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
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
