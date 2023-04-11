package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

// type Task struct {
// 	ID              string
// 	State           string
// 	OriginalArticle string
// 	ErrorMessage    string
// }

func main() {
	tokenParam := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJhZTA5ZTIxYS1lOWM2LTQzYjktYjZiNy1mN2U0N2UzNzg4MGEiLCJleHAiOjE2ODEyMTEyOTMsInZpZXdlcl9pZCI6IjE1MDUxMzg1IiwiaXNfYXBwX2luc3RhbGxlZCI6ZmFsc2UsImlzX21haW5fcHJvZHVjdF9hY3RpdmUiOmZhbHNlLCJpc19hcHBfdGFrZW5fZm9yX2ZyZWUiOnRydWUsImFjdGl2ZV9wcm9kdWN0cyI6W10sImVtYWlsX3N1YnNjcmlwdGlvbiI6eyJzdGF0ZSI6IlNUQVRFX1VOU1BFQ0lGSUVEIn0sInVybCI6IiIsImxhbmciOiJlbiJ9._R1FGBtLRtoIejRB4sx-OMwJYmQr4kNwftTAqoSsSsU"

	tokenParts := strings.Split(tokenParam, ".")
	tokenHeaderB64, tokenPayloadB64, tokenSignature := tokenParts[0], tokenParts[1], tokenParts[2]
	tokenHeader, err := base64.StdEncoding.DecodeString(tokenHeaderB64)
	log.Println("tokenHeader err --->", err)
	tokenPayload, err := base64.StdEncoding.DecodeString(tokenPayloadB64)
	log.Println("tokenPayload err --->", err)
	// tokenSignature, err := base64.StdEncoding.DecodeString(tokenSignatureB64)
	// log.Println("tokenSignature err --->", err)
	//sum := sha256.Sum256([]byte(fmt.Sprintf("%s.%s.", tokenHeaderB64, tokenPayloadB64)))

	h := hmac.New(sha256.New, []byte("fd583962-3439-4226-9729-2b00c41b343c"))
	// Write Data to it
	h.Write([]byte(fmt.Sprintf("%s.%s", tokenHeaderB64, tokenPayloadB64)))
	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))
	fmt.Println("Result: ", sha, tokenSignature)

	sEnc := base64.StdEncoding.EncodeToString([]byte(sha))
	fmt.Println("Result: ", sEnc, tokenSignature)

	log.Println("tokenHeader --->", string(tokenHeader))
	log.Println("tokenPayload --->", string(tokenPayload))
	// log.Println("tokenSignature --->", string(tokenSignature))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	// token, error := jwt.Parse(tokenParam, func(token *jwt.Token) (interface{}, error) {
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, fmt.Errorf("There was an error")
	// 	}
	// 	return []byte("fd583962-3439-4226-9729-2b00c41b343c"), nil
	// })
	// if error != nil {
	// 	fmt.Println("error.Error()", error.Error())
	// 	return
	// }
	// if token.Valid {
	// 	fmt.Println("token valid")
	// 	// 	// var user User
	// 	// 	// mapstructure.Decode(token.Claims, &user)

	// 	// 	// vars := mux.Vars(req)
	// 	// 	name := vars["userId"]
	// 	// 	if name != user.Username {
	// 	// 		json.NewEncoder(w).Encode(ErrorMsg{Message: "Invalid authorization token - Does not match UserID"})
	// 	// 		return
	// 	// 	}

	// 	// 	context.Set(req, "decoded", token.Claims)
	// 	// 	next(w, req)
	// 	// } else {
	// 	// 	json.NewEncoder(w).Encode(ErrorMsg{Message: "Invalid authorization token"})
	// }
}
