package package2

import (
	"fmt"
	"test-project/package1"
)

func IncreaseCapacity(bottle *package1.Bottle) {
	bottle.Volume = bottle.Volume + 1
	fmt.Println("Volume is increased by 1")
	// SomeInt = 2

}
