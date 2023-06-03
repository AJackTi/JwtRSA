package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/AJackTi/JwtRSA/token"
)

func main() {
	prvKey, err := ioutil.ReadFile("cert/id_rsa")
	if err != nil {
		log.Fatalln(err)
	}
	pubKey, err := ioutil.ReadFile("cert/id_rsa.pub")
	if err != nil {
		log.Fatalln(err)
	}

	jwtToken := token.NewJWT(prvKey, pubKey)

	// 1. Create a new JWT token.
	tok, err := jwtToken.Create(time.Hour, "Can be anything")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("TOKEN:", tok)

	// 2. Validate an existing JWT token.
	content, err := jwtToken.Validate(tok)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("CONTENT:", content)
}
