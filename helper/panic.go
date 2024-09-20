package helper

import "log"

func PanicIfError(err error, message string) {
	if err != nil {
		log.Printf("============= \n %s \n===================\n", message)
		panic(err)
	}
}
