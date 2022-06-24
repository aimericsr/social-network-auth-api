package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/aimericsr/social-network-auth-api/token"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func BasicAuth(h http.HandlerFunc, tokenMaker token.Maker) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authorizationHeader := r.Header.Get(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("authorization header is not provided in middleware"))
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid authorization header format"))
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unsupported authorization type"))
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid token"))
		}

		_ = context.WithValue(r.Context(), authorizationPayloadKey, payload)
		//next(w, r.WithContext(ctx))
	})
}
