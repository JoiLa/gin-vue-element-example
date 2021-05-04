package models

import (
    "database/sql"
    "fmt"
    "log"

    "Api/pkg/logging"
    "Api/pkg/setting"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

type Model struct {
    ID         int           `gorm:"primary_key;column:id" json:"id"`
    CreatedOn  sql.NullInt64 `gorm:"column:created_on" json:"-"`
    ModifiedOn sql.NullInt64 `gorm:"column:modified_on" json:"-"`
    DeletedOn  sql.NullInt64 `gorm:"column:deleted_on" json:"-"`
}

func init1() {
    var (
        err                                               error
        dbType, dbName, user, password, host, tablePrefix string
    )

    sec, err := setting.Cfg.GetSection("database")
    if err != nil {
        log.Fatal(2, "Fail to get section 'database': %v", err)
    }

    dbType = sec.Key("TYPE").String()
    dbName = sec.Key("NAME").String()
    user = sec.Key("USER").String()
    password = sec.Key("PASSWORD").String()
    host = sec.Key("HOST").String()
    tablePrefix = sec.Key("TABLE_PREFIX").String()

    Db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
        user,
        password,
        host,
        dbName))

    if err != nil {
        logging.Info("数据库连接失败", err)
    }

    gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
        return tablePrefix + defaultTableName
    }

    Db.SingularTable(true)
    Db.LogMode(true)
    Db.DB().SetMaxIdleConns(10)
    Db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
    defer Db.Close()
}
