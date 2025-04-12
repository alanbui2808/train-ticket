package initialize

import (
	"fmt"
	"time"

	"github.com/alanbui/train-ticket/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}
func InitMySql() {
	m := global.Config.Mysql

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	s := fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.DBname)

	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		/*
			SkipDefaultTransaction: false
				1. Every CREATE, UPDATE, DELETE will be wrapped in a database transaction.
				2. Adds safety ensuring:
					- If fails during operation, the whole transaction is rolled back.
					- Atomic behavior out of the box

				Downsides: Little overhead => reduce latency.

		*/
		SkipDefaultTransaction: false,
	})
	checkErrorPanic(err, "InitMySql initialization error!")

	global.Logger.Info("Initializing MySQL successfully!")
	global.Mdb = db

	// Set Pool
	SetPool()
}

// Connect to Pooling settings
// Used when do low-level DB operations taht GORM doesnt expose directly.
func SetPool() {
	m := global.Config.Mysql

	sqlDB, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("MySQL error: %s::\n", err)
	}

	sqlDB.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))

}

// Create a table in MySQL
func migrateTables() {
	// private function
}
