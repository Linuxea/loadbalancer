package lb

import (
	"fmt"
	"time"
)

func log(message string) {
	fmt.Println(fmt.Sprintf("%s - %s", time.Now().String(), message))
}
