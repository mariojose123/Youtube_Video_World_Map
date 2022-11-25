package Auxfunctions

import "log"

func ReturnError(err error) error {

	log.Print(err)
	return err
}
