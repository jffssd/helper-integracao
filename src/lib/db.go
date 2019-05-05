package lib
/*
import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ctx context.Context
	db  *sql.DB
)

func main() {

	fmt.Println("Conectando no banco de dados")
	db, err := sql.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	cli := "clientname"

	query := fmt.Sprintf("SELECT id, nome FROM db.table WHERE key = '%s'", cli)
	fmt.Println(query)
	results, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	count := 0

	

	for results.Next() {

		var cliente Cliente
		err = results.Scan(&cliente.ID, &cliente.Nome) //, &cliente.Chave, &cliente.ERP, &cliente.WsURL, &cliente.WsUser, &cliente.WsPass)
		if err != nil {
			panic(err.Error())
		}

		if err := rows.Scan(&count); err != nil {
			log.Fatal(err)
		}

		fmt.Println(cliente.ID, cliente.Nome) //, cliente.Chave, cliente.ERP, cliente.WsURL, cliente.WsUser, cliente.WsPass)
	}
}*/