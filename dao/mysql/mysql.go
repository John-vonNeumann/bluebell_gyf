package mysql

import (
	"fmt"
	"gin-cli/setting"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// Init 初始化MySQL连接
func Init(cfg *setting.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	return nil
}

// Close 关闭MySQL连接
func Close() {
	_ = db.Close()
}
