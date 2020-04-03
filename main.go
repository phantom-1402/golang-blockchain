package main

import (
	"fmt"
	"strconv"

	"github.com/jairajsahgal/golang-blockchain/blockchain"
	//"database/sql"
	//_ "github.com/go-sql-driver/mysql"
)

func main() {
	chain := blockchain.InitBlockChain()
	db,err :=sql.Open("mysql","username:password@tcp(localhost:3306)/testdb")
	if err!=nil {
		panic(err.Error())
	}
	_,err = db.Exec("USE testdb")
	if err != nil {
	fmt.Println(err.Error())
	} else {
	fmt.Println("DB selected successfully..")
	}
	stmt, err := db.Prepare("CREATE Table blockchain(id int NOT NULL AUTO_INCREMENT, data varchar(50), PRIMARY KEY (id));")
	if err != nil {
	fmt.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
	fmt.Println(err.Error())
	} else {
	fmt.Println("Table created successfully..")
	}
	
	
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")


	for _, block := range chain.Blocks {
		stmtIns, err := db.Prepare("INSERT INTO blockchain VALUES(?, ? )")
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		_, err = stmtIns.Exec(1,block.Data) // Insert tuples (i, i^2)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))

	}
	defer db.Close()
}
