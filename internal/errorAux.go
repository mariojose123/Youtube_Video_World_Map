package Erroraux

import "fmt"

type ErrorPrint struct {
}

func ReturnError(err error) error {

	fmt.Println(err)
	return err
}
