package main

import (
	"encoding/json"
	"log"
	"strconv"
)

type appError struct {
	//Error   error
	Message string
	Code    int
}

func (source appError) Error() string {
	return source.Message
}

// error handling
func PanicIf(err error) {
	if err != nil {
		logError(err.Error())
		panic(err)
	}
}

// logging
func logError(message string) {
	log.Println("[Error]" + message)
}

func logInfo(message string) {
	log.Println("[Info]" + message)
}

func logDebug(message string) {
	log.Println("[Debug]" + message)
}

//json
func toJSON(target interface{}) (string, error) {
	byteArry, err := json.Marshal(target)
	if err != nil {
		return "", err
	}
	result := string(byteArry[:])
	//logDebug("toJson" + result)
	return result, nil
}

func fromJSON(target interface{}, jsonString string) error {
	byteArray := []byte(jsonString)
	err := json.Unmarshal(byteArray, target)
	if err != nil {
		return err
	}
	return nil
}

func PadRight(str, pad string, lenght int) string {
	for {
		str += pad
		if len(str) > lenght {
			return str[0:lenght]
		}
	}
}

func PadLeft(str, pad string, lenght int) string {
	for {
		str = pad + str
		if len(str) > lenght {
			return str[0:lenght]
		}
	}
}

//parse
func ToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

func ToFloat64(str string) (float64, error) {
	return strconv.ParseFloat(str, 10)
}
