package Db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// 定义db全局变量

func GetDB() (*gorm.DB, error) {
	slowLogger := logger.New(
		//将标准输出作为Writer
		log.New(os.Stdout, "\r\n", log.LstdFlags),

		logger.Config{
			//设定慢查询时间阈值为1000ms
			SlowThreshold: 5 * time.Second,
			//设置日志级别，只有Warn和Info级别会输出慢查询日志
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	//dsn := env.Mysql.Username + ":" + env.Mysql.Password + "@(" + env.Mysql.Url + ":" + strconv.Itoa(env.Mysql.Port) + ")/genshinstudio?charset=utf8mb4&parseTime=True&loc=Local&timeout=3s"
	dsn := "xmsama:wodemima@(101.42.249.68:3306)/face?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 slowLogger,
		SkipDefaultTransaction: true, //禁用事务
		PrepareStmt:            true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用表名加s
		},

		//Logger:                                   logger.Default.LogMode(logger.Info), // 打印sql语句
		DisableAutomaticPing:                     true, //禁用连接时自动ping
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用创建外键约束
	})
	if err != nil {
		fmt.Println("Mysql Connect Failed")
		return nil, err
	}

	sqlDB, err := db.DB()
	//sqlDB.SetMaxIdleConns(20)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(50)

	//// SetConnMaxLifetime 设置了连接可复用的最大时间。
	//sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		fmt.Println("Mysql Connect Failed")
		return nil, err
	}
	fmt.Println("数据库连接成功")
	return db, err
}

//
//func init() {
//	var err error
//	url := "root:genshin@(127.0.0.1:6665)/genshinstudio?charset=utf8&parseTime=True&loc=Local"
//	DB, err = gorm.Open(mysql.Open(url), &gorm.Config{
//		//SkipDefaultTransaction: true, //启用事务
//		NamingStrategy: schema.NamingStrategy{
//			SingularTable: true, // 禁用表名加s
//		},
//		//Logger:                                   logger.Default.LogMode(logger.Info), // 打印sql语句
//		DisableAutomaticPing:                     true, //禁用连接时自动ping
//		DisableForeignKeyConstraintWhenMigrating: true, // 禁用创建外键约束
//	})
//	if err != nil {
//		panic("Connecting database failed: " + err.Error())
//	}
//	sqlDB, err := DB.DB()
//	sqlDB.SetMaxIdleConns(50) //30个连接池
//}

//
//// GetDB对象
//func GetDB() *gorm.DB {
//	return DB
//}
