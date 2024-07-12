package main

import (
	"errors"
	"strconv"
)

// ParseUint is a utility function to parse a string to uint.
func ParseUint(param string) (uint, error) {
	id, err := strconv.ParseUint(param, 10, 32)
	if err != nil {
		return 0, errors.New("invalid id format")
	}
	return uint(id), nil
}

// SuccessResponse formats success responses.
func SuccessResponse(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
		"data":   data,
	}
}

// ErrorResponse formats error responses.
func ErrorResponse(err error) map[string]interface{} {
	return map[string]interface{}{
		"status":  "error",
		"message": err.Error(),
	}
}
