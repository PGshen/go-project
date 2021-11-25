package cmd

import (
	"github.com/PGshen/go-project/tool/internal/sql2struct"
	"github.com/spf13/cobra"
	"log"
)

var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "", "", "数据库账号")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "", "数据库密码")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "数据库Host")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "数据库编码")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "数据库实例类型")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "数据库名")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "数据表名")

}

var sqlCmd = &cobra.Command{
	Use: "sql",
	Short: "sql工具",
	Long: "sql工具",
	Run: func(cmd *cobra.Command, args []string) {

	},

}

var sql2structCmd = &cobra.Command{
	Use: "2struct",
	Short: "sql转结构体",
	Long: "sql转结构体",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo :=&sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}
		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect err: %v", err)
		}
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v", err)
		}
		template := sql2struct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %v", err)
		}
	},
}
