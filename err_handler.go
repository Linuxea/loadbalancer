package lb

import (
	"os"
)

func handlerError(err error) {

	log(err.Error())
	os.Exit(1)
}
