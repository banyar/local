package utils

import (
	"encoding/json"
	"log"
	"strconv"
)

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func (req RestErr) IsRequest() bool {
	return true
}

type ArrayRestErr struct {
	Message []string `json:"message"`
	Status  int      `json:"status"`
	Error   string   `json:"error"`
}

func (req ArrayRestErr) IsRequest() bool {
	return true
}

func BadRequest(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  400,
		Error:   "bad request",
	}
}

func NotFound(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  404,
		Error:   "not found",
	}
}

func InputErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  422,
		Error:   "unprocessable entity",
	}
}

func ValidationErr(message []string) *ArrayRestErr {
	return &ArrayRestErr{
		Message: message,
		Status:  422,
		Error:   "unprocessable entity",
	}
}

func SuccessMessage(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  200,
	}
}

func DisplayJsonFormat(message string, class interface{}) {
	if message == "" {
		message = "CLASS"
	}
	p, err := json.Marshal(class)
	LogError(err)
	log.Println(message+" ===>", string(p))
}

func StringToInt32(x string) (int32, error) {
	i, err := strconv.ParseInt(x, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(i), nil
}

func LogError(err error) {
	if err != nil {
		log.Fatal("ERROR : ", err)
	}
}

func InternalErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  500,
		Error:   "internal server error",
	}
}

func ErrResponse(result error) *RestErr {
	if result != nil {
		if result.Error() == "record not found" {
			return NotFound(result.Error())
		} else {
			return InternalErr(result.Error())
		}
	}
	return &RestErr{}
}
