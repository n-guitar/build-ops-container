package database

import "gorm.io/gorm"

// データベース接続へのポインターとなるグローバル変数
var (
	DBConn *gorm.DB
)
