package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/whiteblock/httputils/responses"

	"github.com/coreos/go-oidc"
	"github.com/sirupsen/logrus"
	"github.com/whiteblock/utility/auth"
)

func AuthN(log logrus.Ext1FieldLogger, verifier *oidc.IDTokenVerifier) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "GET" && strings.HasSuffix(r.URL.Path, "/health") {
				next.ServeHTTP(w, r)
				return
			}
			header := r.Header.Get("Authorization")
			if header == "" || !strings.Contains(header, "Bearer") {
				log.Error(responses.ErrMissingToken)
				responses.MissingToken(w)
				return
			}
			token, err := auth.VerifyToken(verifier, r.Context(), header)
			if err != nil {
				log.Error(err)
				responses.InvalidToken(w)
				return
			}

			userContext, err := auth.GetUserContext(token)
			if err != nil {
				log.Error(err)
				responses.InvalidToken(w)
				return
			}

			ctx := context.WithValue(r.Context(), "UserContext", userContext)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
