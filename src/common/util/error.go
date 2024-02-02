package util

import (
	"errors"
	"strings"
)

var (
	NotFoundErr     = errors.New("not found")
	AlreadyExistErr = errors.New("already exist")
	DataTypeErr     = errors.New("date type")
)

func IsNotFoundErr(err error) bool {

	return strings.Contains(err.Error(), NotFoundErr.Error())
}

func IsAlreadyExistErr(err error) bool {

	return strings.Contains(err.Error(), AlreadyExistErr.Error())
}

func IsDataTypeErr(err error) bool {

	return strings.Contains(err.Error(), DataTypeErr.Error())
}
