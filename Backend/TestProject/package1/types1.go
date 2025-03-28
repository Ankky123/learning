package package1

import (
	"fmt"
	"log"
	"os"
)

type Bottle struct {
	Name     string
	Material string
	Volume   int
}

func GetBottle(bottle Bottle) {
	fmt.Printf("Name : %s, Materia : %s, Volume : %v\n", bottle.Name, bottle.Material, bottle.Volume)
}

func CreateFile() {
	path := "C:\\Users\\Ankit Mangal\\Desktop\\Learning\\Backend\\TestProject\\assets\\file.txt"
	goPath := os.Getenv("GOPATH")
	fmt.Println("GOPATH is", goPath)
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File is created and written")
	defer file.Close()

	file.WriteString("Hi this is some file")

}
