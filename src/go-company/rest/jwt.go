package rest

import (
	"fmt"
	"net/http"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

var signString []byte = []byte("signing_string")
var Key string = "1111"

func  CreateJWT()(string,error){
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	signedToken, err  := token.SignedString(signString)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GetJWT(w http.ResponseWriter, req *http.Request) {

	if req.Header["Access"] != nil {
		if req.Header["Access"][0] != Key {
			return
		} else {
			token, err := CreateJWT()
			if err != nil {
				return
			}
			fmt.Fprint(w, token)
		}
	}

}



func validateJWT(w http.ResponseWriter, r *http.Request) (err error) {
	if r.Header["Token"] == nil {
		fmt.Fprintf(w, "can not find token in header")
		return errors.New("Token error")
	}

	token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return signString, nil
	})

	if token == nil {
		fmt.Fprintf(w, "invalid token")
		return errors.New("Token error")
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Fprintf(w, "couldn't parse claims")
		return errors.New("Token error")
	}


	return nil
}

