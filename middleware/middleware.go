package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Dparty/common/utils"
	coreModel "github.com/Dparty/model/core"
)

var UserCtxKey = &ContextKey{"user"}

type ContextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorization := r.Header.Get("Authorization")
			splited := strings.Split(authorization, " ")
			if authorization == "" || len(splited) != 2 {
				next.ServeHTTP(w, r)
				return
			}
			claims, err := utils.VerifyJwt(splited[1])
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			account, _ := coreModel.FindAccount(utils.StringToUint(claims["id"].(string)))
			ctx := context.WithValue(r.Context(), UserCtxKey, account)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
