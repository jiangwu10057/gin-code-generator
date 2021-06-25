package model


import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cengsin/oracle"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"gin-code-generator/internal/pkg/config"
)

var DB *gorm.DB

func InitOrm(fullConfig config.FullConfig) (error) {
	dbType := fullConfig.SystemConfig.DbType
	switch dbType {
	case "mysql":
		return InitMysql(fullConfig)
	case "oracle":
		return InitOracle(fullConfig)
	default:
		return fmt.Errorf("db-type只支持配置:oracle,mysql")
	}
}

func NewOrmConfig() (*gorm.Config)  {
	return &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold: 1 * time.Millisecond,
			LogLevel:      logger.Warn,
			Colorful:      true,
		}),
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true, //表名后面不加s
		},
	}
}

func InitMysql(fullConfig config.FullConfig) (error)  {
	config := fullConfig.MySqlConfig
	dsn := config.Username + ":" + config.Password + "@tcp(" + config.Path + ")/" + config.Dbname + "?" + config.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	orm, err := gorm.Open(mysql.New(mysqlConfig), NewOrmConfig())
	DB = orm
	return err
}

func InitOracle(fullConfig config.FullConfig) (error) {
	config := fullConfig.OracleConfig
	
	dsn := config.Username+"/"+ config.Password+"@"+ config.Path+"/"+ config.Dbname
	orm, err := gorm.Open(oracle.Open(dsn), NewOrmConfig())
	DB = orm
	return err
}
