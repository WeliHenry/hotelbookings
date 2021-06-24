package main

import (
	"github.com/alexedwards/scs/v2"
	"net/http"
	"time"
)

const portNumber = ":8080"



func main()  {
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	



}

