package auth

import (
	"net/http"

	"github.com/BeryJu/gopyazo/pkg/drivers"
)

type AuthDriver interface {
	drivers.HTTPDriver
	AuthenticateRequest(w http.ResponseWriter, r *http.Request, next http.Handler)
}