package repositories

import (
	"../entities"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var database *sql.DB

func GetAllUsers(args ...interface{}) ([]*entities.User, error) {
	db, err := sql.Open("mysql", "docker:docker@tcp(db:3306)/test_db")


	rows, err := db.Query("SELECT * FROM test_db.user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := make([]*entities.User, 0, 10)

	for rows.Next() {
		us := new(entities.User)
		err := rows.Scan(&us.Id, &us.Username, &us.Nick, &us.CreatedAt, &us.ConfirmToken, &us.ConfirmTokenExpire, &us.Status, &us.Password, &us.Email)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, us)
	}

	return users, nil
}

func GetUserById(Id int) (*entities.User, error) {
	rows, err := database.Query("SELECT * FROM test_db.user WHERE id = ?", Id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	us := new(entities.User)

	errorScan := rows.Scan(&us.Id, &us.Username, &us.Nick, &us.Email, &us.CreatedAt, &us.Status)
	if errorScan != nil {
		log.Fatal(errorScan)
	}


	return us, nil
}


func UpdateUserById(Id int, Username, Nick string) (*entities.User, error) {
	_, err := database.Exec("UPDATE test_db.user set username = ?, nick = ? where id = ?", Username, Nick, Id)
	if err != nil {
		log.Fatal(err)
	}

	finalUser, err := GetUserById(Id)
	if err != nil {
		log.Fatal(err)
	}

	return finalUser, nil
}


func DeleteUserById(Id int) {
	_, err := database.Exec("delete from test_db.use where id = ?", Id)
	if err != nil{
		panic(err)
	}
}

func CreateUser(CreatedAt int, ConfirmToken string, ConfirmTokenExpire int, Status int, Password, Email string) {
	db, err := sql.Open("mysql", "docker:docker@tcp(db:3306)/test_db")

	_, err = db.Exec(
		"insert into test_db.user " +
			"(created_at, confirm_token, confirm_token_expire, status, password, email) " +
			"values (?, ?, ?, ?, ?, ?)",
		CreatedAt, ConfirmToken, ConfirmTokenExpire, Status, Password, Email)

	if err != nil {
		log.Println(err)
	}
}