package responses

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
	"github.com/whiteblock/utility/common"
)

var renderer = render.New()

func HandleError(w http.ResponseWriter, err error) {
	switch err {
	case common.ValidationError:
		_ = renderer.JSON(w, http.StatusBadRequest, &RegistrarResponse{common.ValidationError.Error()})
	case common.AccessDenied:
		_ = renderer.JSON(w, http.StatusForbidden, &RegistrarResponse{common.AccessDenied.Error()})
	case common.NotAMember:
		_ = renderer.JSON(w, http.StatusNotFound, &RegistrarResponse{common.NotAMember.Error()})
	case common.EmptyDBResult:
		_ = renderer.JSON(w, http.StatusNotFound, &RegistrarResponse{common.EmptyDBResult.Error()})
	case common.AlreadyExists:
		_ = renderer.JSON(w, http.StatusConflict, &RegistrarResponse{common.AlreadyExists.Error()})
	default:
		_ = renderer.JSON(w, http.StatusInternalServerError, &RegistrarResponse{SomethingWentWrong})
	}
}
func resp(w http.ResponseWriter, stat int, msg string) {
	err := renderer.JSON(w, stat, &RegistrarResponse{msg})
	if err != nil {
		log.Error(err)
	}
}

func InvalidParam(w http.ResponseWriter) {
	resp(w, http.StatusBadRequest, ErrInvalidParam)
}

func MissingUserInfo(w http.ResponseWriter) {
	resp(w, http.StatusBadRequest, "missing user info")
}

func InvalidRequestBody(w http.ResponseWriter) {
	resp(w, http.StatusBadRequest, ErrInvalidRequestBody)
}

func CannotReadRequestBody(w http.ResponseWriter) {
	resp(w, http.StatusInternalServerError, ErrCannotReadRequestBody)
}

func MissingReqFields(w http.ResponseWriter, vars map[string]string) {
	resp(w, http.StatusBadRequest, MissingFields(vars))
}

func UnknownError(w http.ResponseWriter) {
	resp(w, http.StatusInternalServerError, SomethingWentWrong)
}

func MissingEmail(w http.ResponseWriter) {
	resp(w, http.StatusBadRequest, MissingEmailMsg)
}

func NotFound(w http.ResponseWriter) {
	resp(w, http.StatusNotFound, NotFoundMsg)
}

func MissingPermissionLevel(w http.ResponseWriter) {
	resp(w, http.StatusBadRequest, MissingPermissionLevelMsg)
}

func MissingToken(w http.ResponseWriter) {
	resp(w, http.StatusUnauthorized, ErrMissingToken.Error())
}

func InvalidToken(w http.ResponseWriter) {
	resp(w, http.StatusUnauthorized, ErrInvalidToken.Error())
}
