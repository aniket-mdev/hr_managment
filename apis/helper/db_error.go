package helper

import (
	"errors"
	"fmt"
	"strings"

	"github.com/lib/pq"
)

func HandleDBErr(err error) (err_ error) {
	fmt.Println("DB-Handler get called...")
	switch e := err.(type) {
	case *pq.Error:
		switch e.Code {
		case "23502":
			fmt.Println("some required data was left out: \n\n", e.Message)
			return
		case "23505":
			fmt.Println("some required data was left out 2 : \n\n", e.Detail)
			err_ = errors.New(e.Detail)
			return
		case "23514":
			// check constarint validation
			// for different fields
			fmt.Println("some required data was left out 3 : \n\n", e.Message)

			if strings.Contains(e.Message, "contact") {
				err_ = errors.New("contact should not be empty")
				return
			} else if strings.Contains(e.Message, "email") {
				err_ = errors.New("email should not be empty")
				return
			} else {
				fmt.Println("Un-Define Error code : ", e.Code, e.Message)
			}
		case "23503":
			fmt.Println("some required data was left out 4 : \n\n", e.Message)
			err_ = errors.New("invalid id has been provided,please try with valid id's")
			return
		default:
			msg := e.Message

			if d := e.Detail; d != "" {
				msg += "\n\n" + d
			}

			if h := e.Hint; h != "" {
				msg += "\n\n" + h
			}

			fmt.Println("Message from default : ", e.Code)
			err_ = errors.New(msg)
			return
		}
	default:
		fmt.Println("DB Error from parent def : ")
		err_ = nil
		return
	}
	fmt.Println("function returing")
	return
}
