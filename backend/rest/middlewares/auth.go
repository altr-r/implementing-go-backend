package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"net/http"
	"strings"

	"github.com/altr-r/implementing-go-backend/config"
	"github.com/altr-r/implementing-go-backend/utils"
)

func AuthenticateJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// parse jwt
		// parse header and payload or claims
		// hmac-sha-256 algorithm -> hash(header, payload, secret key)
		// parse signature part from the jwt
		// if the signature and hash is same => forward to create products
		// otherwise 401 status code with Unauthorized
		cnf := config.GetConfig()

		header := r.Header.Get("Authorization")

		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		headerArr := strings.Split(header, " ")

		if len(headerArr) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		accessToken := headerArr[1]

		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		signature := tokenParts[2]

		message := jwtHeader + "." + jwtPayload
		messageBytArr := []byte(message)

		secretBytArr := []byte(cnf.JwtSecretKey)

		hashed := hmac.New(sha256.New, secretBytArr)
		hashed.Write(messageBytArr)

		hash := hashed.Sum(nil)
		newSignature := utils.Base64UrlEncode(hash)

		if newSignature != signature {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
