package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/opentracing/opentracing-go"
)

const USERNAME = "root"      //"your user name"
const PASSWORD = "vitalik88" //"your password"
const HOST = "localhost"     //"your host"
const PORT = "3306"          //"your port"
const DB = "usersdb"         //"your db name"

const DATASOURCENAME = USERNAME + ":" + PASSWORD + "@tcp(" + HOST + ":" + PORT + ")/" + DB

func HandleGet(w http.ResponseWriter, r *http.Request) {
	trace, ctx := opentracing.StartSpanFromContext(r.Context(), "Handle /get")
	time.Sleep(time.Second / 2)
	defer trace.Finish()

	//check login
	if isLogin(ctx) {
		fmt.Fprintf(w, "USER LOGIN\n")
	}

	//get user
	user := getUser(1, ctx) //put here ID to get user
	fmt.Fprintf(w, "User ID that you looking for is: %v", user)
}

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	trace, ctx := opentracing.StartSpanFromContext(r.Context(), "Handle /create")
	time.Sleep(time.Second / 2)
	defer trace.Finish()

	//check login
	if isLogin(ctx) {
		fmt.Fprintf(w, "USER LOGIN\n")
	}

	//create status
	userCreateStatus := createUser(ctx)
	fmt.Fprintln(w, "User: ", userCreateStatus)
}

func isLogin(ctx context.Context) bool {
	var status bool

	trace, ctx := opentracing.StartSpanFromContext(ctx, "func isLogin")
	time.Sleep(time.Second / 2)
	defer trace.Finish()

	//get status
	status = true

	return status
}

func createUser(ctx context.Context) string {
	trace, ctx := opentracing.StartSpanFromContext(ctx, "func createUser")
	time.Sleep(time.Second / 2)
	defer trace.Finish()

	db, err := sql.Open("mysql", DATASOURCENAME)
	if err != nil {
		panic(err.Error())
	}
	db.Query("INSERT INTO `usersdb`.`Users` (`email`, `activation_code`, `name`) VALUES ('test@test.com', '67582', 'Peter')")

	return "created successful"
}

func getUser(Id int, ctx context.Context) int {
	trace, ctx := opentracing.StartSpanFromContext(ctx, "func getUser")
	time.Sleep(time.Second / 2)
	defer trace.Finish()

	db, err := sql.Open("mysql", DATASOURCENAME)
	if err != nil {
		panic(err.Error())
	}
	rows, err := db.Query("select id, name, email, activation_code from usersdb.Users where id = ?", Id)
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.ActivationCode)
		if err != nil {
			panic(err.Error())
		}
		id := user.Id
		return id
	}
	return 0
}

type User struct {
	Id             int
	Name           string
	Email          string
	ActivationCode string
}
