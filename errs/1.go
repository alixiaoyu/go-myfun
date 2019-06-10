package main

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

func main() {
	err := make(0)
	if err == nil {
		fmt.Println("OK")
	} else if err == gorm.ErrRecordNotFound {
		fmt.Println("99999")
	} else {
		fmt.Println(err)
	}

}

func make(i int) error {
	if i > 0 {
		return gorm.ErrRecordNotFound
	} else if i == 0 {
		return nil
	} else {
		return errors.New("KKKKKK")
	}

}
