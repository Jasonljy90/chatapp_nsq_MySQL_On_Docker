package main

import (
	"SessionCookies/token"
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Declaration of global variable
var (
	tpl   *template.Template
	mutex sync.Mutex // Concurrency
	db    *sql.DB
	err1  error
)

// Declaration of type patient
type messageTest struct {
	MessageName    string `validate:"required"`
	MessageContent string `validate:"required"`
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func home(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "homePage.gohtml", nil)
}

func userLoginSuccess(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "userLoginSuccess.gohtml", nil)
}

func getEnvVars() {
	err := godotenv.Load("credentials.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

func userChat(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		msgName := req.FormValue("messagename")
		msgContent := req.FormValue("messagecontent")

		m := messageTest{msgName, msgContent}
		validate := validator.New()
		err := validate.Struct(m)
		if err != nil {
			io.WriteString(res, `
			<html>
			<meta http-equiv='refresh' content='5; url=/userchat '/>
			Please fill in all fields!<br>
			You will be redirected shortly in 5 seconds...<br>
			</html>
			`)
			return
		}
		SendMsg(msgName, msgContent) // publish message to producer
	}
	tpl.ExecuteTemplate(res, "chatAgentRider.gohtml", nil)
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3307)/mysql")
	if err != nil {
		panic(err)
	}

	// Generate key and convert to string for password reset feature
	key := generateSecretKey()
	keyStr = string(key)
	maker, err = token.NewJWTMaker(keyStr)
	if err != nil {
		fmt.Println("Error generating token maker!")
	}

	// http multiplexer with gorilla/mux
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/logout", logout)
	router.HandleFunc("/usersignup", userSignup)
	router.HandleFunc("/userlogin", userLogin)
	router.HandleFunc("/userloginsuccess", userLoginSuccess)
	router.HandleFunc("/userchat", userChat)
	router.HandleFunc("/deleteuser", deleteUser)
	router.HandleFunc("/userchangepassword", userChangePassword)
	router.HandleFunc("/userchangelanguage", userChangeLanguage)
	router.HandleFunc("/userresetpassword", userResetPassword)
	router.HandleFunc("/userresetchangepassword", userResetChangePassword)
	router.HandleFunc("/usertoken/{token}", resetUserPasswordLinkClicked)
	fmt.Println("server is up")
	http.ListenAndServe(":5221", router)
}
