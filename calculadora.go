package main

import (
	"fmt"
	"strconv"
)

type result struct {
	Value float64 `json:"value" bson:"value,omitempty"`
}

func Sum(op1, op2 string) (*result, error) {
	x, err := strconv.ParseFloat(op1, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid input: %s", op1)
	}
	y, err := strconv.ParseFloat(op2, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid input: %s", op2)
	}
	return &result{Value: x + y}, nil
}
