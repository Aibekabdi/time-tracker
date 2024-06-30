package user

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/Aibekabdi/time-tracker/internal/handlers/utils"
	"github.com/gin-gonic/gin"
)

var (
	errPassportIn     = errors.New("invalid user passport")
	errPassportSerie  = errors.New("invalid user passport serie")
	errPassportNumber = errors.New("invalid user passport number")
)

type userPassportIn struct {
	PassportInfo string `json:"passportNumber"`
}

func (u *userPassportIn) validate() error {
	arr := strings.Split(u.PassportInfo, " ")
	if len(arr) != 2 {
		return errPassportIn
	}
	_, err := strconv.Atoi(arr[0])
	if err != nil {
		return errPassportSerie
	}
	_, err = strconv.Atoi(arr[1])
	if err != nil {
		return errPassportNumber
	}
	return nil
}

func (h *Handler) Create(c *gin.Context) {
	const op = "handler.User.Create"
	var input userPassportIn
	log := h.log.With(slog.String("op", op))
	if err := c.BindJSON(&input); err != nil {
		utils.ErrorHandler(c, http.StatusBadRequest, fmt.Sprintf("can not parse json: %s", err.Error()))
	}
	if err := input.validate(); err != nil {
		utils.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	log.Debug("got user passport", slog.String("passport", input.PassportInfo))
}
