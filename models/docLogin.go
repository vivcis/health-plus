package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"unicode"

	//"golang.org/x/crypto/bcrypt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var tpl *template.Template
var db *sql.DB

func main() {
	tpl, _ = template.ParseGlob("pages/*.html")
	var err error
	db, err = sql.Open("mysql", "appliCATION123@#@tcp(localhost:3306)/healthPlus_DB")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/loginauth", loginAuthHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/registerauth", registerAuthHandler)
	http.ListenAndServe(":8080", nil)
}

// loginHandler handles the login page, serves form for doctors to login with
func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*********loginHandler running*********")
	tpl.ExecuteTemplate(w, "doctorLogin.html", nil)
}

//loginAuthHandler handles the login authentication, checks if the doctor is in the database
func loginAuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*********loginAuthHandler running*********")
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")
		fmt.Println("username: ", username, "password: ", password)
		// retrive password from db to compare (hash) with user supplied password hash 
		var hash string
		statement := "SELECT Hash FROM bcrypt WHERE Username = ?"
		row := db.QueryRow(statement, username)
		err := row.Scan(&hash)
		fmt.Println("hash from db: ", hash)
		if err != nil {
			fmt.Println("error selecting Hash in db by Username: ", err)
			tpl.ExecuteTemplate(w, "doctorLogin.html", "check username and passsword")
			return
		} 
		//func compareHashAndPassword(hashedPassword, password string) error 
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))	
	//returns nill on success
	if err != nil {
		fmt.Fprint(w, "You have successfully logged in : )")
		return
}
    fmt.Println("incorrect password")
	tpl.ExecuteTemplate(w, "doctorLogin.html", "confirm username and  password")
}

//registerHandler serves form for registering new user
func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*********registerHandler running*********")
	tpl.ExecuteTemplate(w, "doctorRegister.html", nil)
}

//registerAuthHandler creates newuser in the database
func registerAuthHandler(w http.ResponseWriter, r *http.Request){
	/*
	   1. check username criteria
	   2. check password criteria
	   3. check if username already exists
	   4. create bcrypt hash from password
	   5. insert username and password hash into db
	*/
	fmt.Println("*********registerAuthHandler running*********")
	r.ParseForm()
	username := r.FormValue("username")

	//check username for only alphanumeric characters
	//var nameAlphanumeric  = true
	for _, char := range username {
		//func isLetter(r rune) bool, func isNumber(r rune) bool
		//if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
		if unicode.IsLetter(char) == false && unicode.IsNumber(char) == false {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode(map[string]string{"error": "username must be alphanumeric"})
			return
			//nameAlphanumeric = false
		}
	}
	//check username passwordlength
	//var nameLength bool
	if 5 <= len(username) && len(username) <= 50 {
		//nameLength = true
	}

	age := r.FormValue("age")
	//check age for only numbers
	//var ageNumeric bool
	for _, char := range age {
		if unicode.IsNumber(char) == false {
			//ageNumeric = false
		}
	}
	//check age length
	//var ageLength bool
	if 1 <= len(age) && len(age) <= 3 {
		//ageLength = true
	}

	email := r.FormValue("email")
	//check email for only alphanumeric characters
	//var emailAlphanumeric bool
	for _, char := range email {
		if unicode.IsLetter(char) == false && unicode.IsNumber(char) == false {
			//emailAlphanumeric = false
		}
	}
	

	//check password criteria
	password := r.FormValue("password")
	fmt.Println("password: ", username, "\npswdLength: ", len(password))
	//variables that must pass for password creation criteria
	var pswdLoweracase, pswdUppercase, pswdNumber, pswdSpecial, pswdLength, pswdNoSpaces bool
	pswdNoSpaces = true
	for _, char := range password {
		switch {
			//func isLower (r rune) bool
		case unicode.IsLower(char):
			pswdLoweracase = true
			//func isUpper (r rune) bool
		case unicode.IsUpper(char):
			pswdUppercase = true
			//func isNumber (r rune) bool
		case unicode.IsNumber(char):
			pswdNumber = true
			//func isLetter (r rune) bool, func isSymbol (r rune) bool
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			pswdSpecial = true
			//func isSpace (r rune) bool, type rune = int32
		case unicode.IsSpace(int32(char)):
			pswdNoSpaces = false
		}
	}
	if 11 < len(password) && len(password) < 60{
		pswdLength = true
	}
	fmt.Println("pswdLowercase: ", pswdLoweracase, "\npswdUppercase: ", pswdUppercase, "\npswdNumber: ", pswdNumber, "\npswdSpecial: ", pswdSpecial, "\npswdLength: ", pswdLength, "\npswdNoSpaces: ", pswdNoSpaces)
	if !pswdLoweracase || !pswdUppercase || !pswdNumber || !pswdSpecial || !pswdLength || !pswdNoSpaces {
		tpl.ExecuteTemplate(w, "doctorRegister.html", "please check username and password criteria")
		return
	}
	//check if username already exists for availability
	stmt := "SELECT Username FROM bcrypt WHERE Username = ?"
	row := db.QueryRow(stmt, username)
	var uID string
	err := row.Scan(&uID)
	if err == sql.ErrNoRows {
		fmt.Println("username already exists, err:", err)
	    tpl.ExecuteTemplate(w, "doctorRegister.html", "username already taken")
	    return
	}

	//create hash from password
	var hash []byte
	//func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		fmt.Println("bcrypt error: ", err)
		tpl.ExecuteTemplate(w, "doctorRegister.html", "there was a problem registering account")
		return
	}
	fmt.Println("hash:", hash)
	fmt.Println("string(hash):", string(hash))
	//func (db *DB) Prepare(query string) (*stmt, error)
	var insertStmt *sql.Stmt
	insertStmt, err = db.Prepare("INSERT INTO bcrypt (username, Hash) VALUES(? ?);")
	if err != nil{
		fmt.Println("error preparing statement:", err)
		tpl.ExecuteTemplate(w, "doctorRegister.html", "there was a problem registering account")
		return
	}
	defer insertStmt.Close()

	var result sql.Result
	//func (s *Stmt) Exec (args ...interfaces{}) (Result, error)
	result, err = insertStmt.Exec(username, hash)
	rowsAff, _ := result.RowsAffected()
	lastIns, _ := result.LastInsertId()
	fmt.Println("rowsAff:", rowsAff)
	fmt.Println("lastIns:", lastIns)
	fmt.Println("err:", err)
	if err != nil{
		fmt.Println("error inserting new user")
		tpl.ExecuteTemplate(w, "doctorRegister.html", "there was a problem registering account")
		return
	}
	fmt.Fprint(w, "congrats, your account has been successfully created")
    }
