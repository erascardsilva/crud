package conect

import (
	"fmt"
	"log"
)

func ReadDados() {
	db := Conect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM usuarios")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Dados no banco de dados ")
	fmt.Println("------------------------------------------------------------------------------")
	for rows.Next() {
		var id int
		var nome, email, whats string

		err := rows.Scan(&id, &nome, &email, &whats)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID =   %v   |   Nome =   %v   |   Email =   %v   |    Whatsapp =   %v   \n", id, nome, email, whats)
	}
	fmt.Println("------------------------------------------------------------------------------")

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
