package main

import (
	"fmt"
	"github.com/binqibang/mini-douyin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"log"
	"strings"
)

// Use `gen` pkg to generate `gorm` struct from mysql.
func main() {
	conf, err := config.LoadConfig("E:/字节青训营/mini-douyin/config/settings_dev.yml")
	if err != nil {
		log.Fatalf("cannot read conf file: %s", err)
	}
	// connect database
	mysqlConf := conf.Database.Mysql
	dsn := mysqlConf.UserName + ":" + mysqlConf.Password
	dsn += "@(" + mysqlConf.Address + ")/" + mysqlConf.Database + "?charset=utf8"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println(dsn)
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}

	// define generator rules
	g := gen.NewGenerator(gen.Config{
		// in the output path, you can find a `table.gen.go` file, just copy
		// the struct definition.
		OutPath:           "tmp",
		Mode:              gen.WithDefaultQuery,
		FieldNullable:     false, // generate pointer when field is nullable
		FieldCoverable:    false, // generate pointer when field has default value
		FieldSignable:     false, // detect integer field's unsigned type, adjust generated data type
		FieldWithIndexTag: false, // generate with gorm index tag
		FieldWithTypeTag:  true,  // generate with gorm column type tag
	})
	g.UseDB(db)

	// add datatype maps, like mysql:int -> go:int64
	dataMap := map[string]func(detailType string) (dataType string){
		"int": func(detailType string) (dataType string) { return "int64" },
	}
	g.WithDataTypeMap(dataMap)

	// add json tags
	jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
		toStringField := `balance, `
		if strings.Contains(toStringField, columnName) {
			return columnName + ",string"
		}
		return columnName
	})

	// generate model
	comment := g.GenerateModel("video_comment", jsonField)
	g.ApplyBasic(comment)
	g.Execute()
}
