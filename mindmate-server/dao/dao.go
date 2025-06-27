/*
 * Data Access Object
 */

package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Import MySQL driver
	"github.com/noahzyl/mindmate/config"
	"github.com/noahzyl/mindmate/pkg/logger"
	"time"
)

var (
	DB  *gorm.DB // Connection to the database
	err error
)

func init() {
	DB, err = gorm.Open("mysql", config.Mysqldb) // Connect to the database
	// If an error happened, write a log
	if err != nil {
		logger.Error(map[string]interface{}{"mysql connect error": err.Error()})
		// If err is nil, then the database connection failed, DB.Error is nil
		panic("Failed to connect database") // Stop the server to avoid the nil error DB.Error
	}
	if DB.Error != nil {
		logger.Error(map[string]interface{}{"database error": DB.Error})
	}
	// Set connection pool
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	DB.DB().SetConnMaxLifetime(30 * time.Minute)
}
