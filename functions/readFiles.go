package functions

import (
	"fmt"
	"os"
)

func ReadFile(fileName string) []byte {
	datas, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des données: ", err)
		os.Exit(1)
	}

	return datas
}
