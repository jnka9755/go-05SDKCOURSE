package main

import (
	"errors"
	"fmt"
	"os"

	courseSDK "github.com/jnka9755/go-05SDKCOURSE/course"
)

func main() {

	courseTrans := courseSDK.NewHttpClient("http://localhost:5002", "")

	course, err := courseTrans.Get("11857977-8602-41c6-825c-e0b2e5e17476")

	if err != nil {
		if errors.As(err, &courseSDK.ErrNotFound{}) {
			fmt.Println("Not found: ", err.Error())
			os.Exit(1)
		}

		fmt.Println("Internal server error: ", err.Error())
		os.Exit(1)
	}

	fmt.Println(course)
}
