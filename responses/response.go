package responses

import (
	"errors"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type RegistrarResponse struct {
	Message string `json:"message"`
}

// generic error response messages
const AlreadyExists = "entity already exists"
const ErrCannotReadRequestBody = "could not read request body"
const ErrInvalidRequestBody = "invalid request body"
const MissingEmailMsg = "email must be provided"
const MissingPermissionLevelMsg = "permission_level must be provided"
const SomethingWentWrong = "something went wrong"
const ErrInvalidParam = "query parameter is invalid"
const NotFoundMsg = "not found"

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrMissingToken = errors.New("missing token")
)

func MissingFields(fields interface{}) string {
	return fmt.Sprintf(`missing required field. Received: %+v`, fields)
}

// generic success messages
const SuccessMsg = "success"

func JSON(w http.ResponseWriter, obj interface{}) {
	err := renderer.JSON(w, http.StatusOK, obj)
	if err != nil {
		log.Error(err)
		UnknownError(w)
	}
}

func Success(w http.ResponseWriter) {
	err := renderer.JSON(w, http.StatusOK, &RegistrarResponse{SuccessMsg})
	if err != nil {
		log.Error(err)
	}
}
