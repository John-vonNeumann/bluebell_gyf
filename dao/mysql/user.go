package mysql

import (
	"bluebell_gyf/models"
	"crypto/md5"
	"encoding/hex"
)

const name = "guoyifan"
const secret = "guoyifan.com"

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int64
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	user.Password = encryptPassword(user.Password)
	sqlStr := `insert into user (user_id, username, password) values (?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}

func Login(user *models.User) (err error) {
	oldPassword := user.Password
	sqlStr := `select user_id, username, password from user where username=? `
	err = db.Get(user, sqlStr, user.Username)
	if err != nil {
		return err
	}
	if encryptPassword(oldPassword) != user.Password {
		return ErrorInvalidPassword
	}
	return
}
