package main

import (
	"context"
	"fmt"
	"os"
	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
)

func main() {
	fmt.Println("start------")

	cfg := mysql.Cfg(os.Getenv("DB_CONNECTION_NAME"),os.Getenv("DB_USER"), os.Getenv("DB_PASS"))
	cfg.DBName = "db"
	db, err := mysql.DialCfg(cfg)
	if err != nil {
		fmt.Printf("%s\n",err)
		os.Exit(1)	
	}
	defer db.Close()
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, "SELECT * FROM user")
    if err != nil {
		fmt.Printf("%s\n",err)
        os.Exit(1)
    }
	defer rows.Close()
	for rows.Next() {
        fmt.Printf("%s\n","aaa" )
	}
	
	fmt.Println("end------")
	os.Exit(0)
}
