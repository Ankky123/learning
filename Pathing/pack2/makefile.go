package pack2

import (
	"fmt"
	"log"
	"os"
)

func MakeAFile() {
	//path is found to be in relation to the main.go file
	//the ./ and direct are found to be same
	/*
		file0, err := os.Create("outfile0.txt")

		if err != nil {
			log.Fatal(err)
		}

		defer file0.Close()

		file0.WriteString("Something in the file0")

		file1, err := os.Create("./pack1/outfile1.txt")

		if err != nil {
			log.Fatal(err)
		}

		defer file1.Close()

		file1.WriteString("Something in the file1")

		file2, err := os.Create("pack1/outfile2.txt")

		if err != nil {
			log.Fatal(err)
		}

		defer file2.Close()

		file2.WriteString("Something in the file2")

		file3, err := os.Create("./pack2/outfile3.txt")

		if err != nil {
			log.Fatal(err)
		}

		defer file3.Close()

		file3.WriteString("Something in the file3")

		file4, err := os.Create("./pack2/pack1/outfile4.txt")

		if err != nil {
			log.Fatal(err)
		}

		defer file4.Close()

		file4.WriteString("Something in the file4")
	*/
	goPath := os.Getenv("GOPATH")
	goPath += "/src/New folder/file.txt"
	fmt.Println(goPath)

	file5, err := os.Create(goPath)

	if err != nil {
		log.Fatal(err)
	}

	defer file5.Close()

	file5.WriteString("Something in the file5")

	path2 := "C:\\Users\\Ankit Mangal\\Desktop\\Learning\\Pathing\\outfile6.txt"
	file6, err := os.Create(path2)

	if err != nil {
		log.Fatal(err)
	}

	defer file6.Close()

	file6.WriteString("Something in the file6")

	//osPath := os.Getenv("PATH")
	//we need to set another variable inside the PATH variable
	//during installation of the setup, this variable will be set
	//we need to make a condition of whether this variable is already present or not
	//fmt.Println(osPath)

	// file6, err := os.Create(path2)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer file6.Close()

	// file6.WriteString("Something in the file6")

	//setting an environment variable
	value, ok := os.LookupEnv("POCPATH")
	if !ok {
		fmt.Println("POCPATH IS NOT EVEN DEFINED YET")
	}
	fmt.Println("POCPATH :", value)

	//The setup will put the fiiles in the following directory in the users pc
	// pathVarName := "POCPATH"
	// completePath := "C:\\Program Files\\Company\\POC"
	// os.Setenv(pathVarName, completePath)
	// fmt.Println(pathVarName, os.Getenv("POCPATH"))

	// value, ok = os.LookupEnv(pathVarName)
	// if !ok {
	// 	fmt.Println(pathVarName, " IS NOT EVEN DEFINED YET")
	// }
	// fmt.Println(pathVarName, value)

	// for _, vari := range os.Environ() {
	// 	fmt.Println(vari)
	// }

}
