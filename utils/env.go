package utils

import (
	"os"
	"strconv"
)

func EnvOrDef(key, def string) string {
	v, ok := os.LookupEnv(key)
	if ! ok {
		return  def
	}
	return v
}

func EnvOrDefInt(key string, def int) int {
	v, ok := os.LookupEnv(key)
	if ! ok {
		return  def
	}
	d, err := strconv.Atoi(v)
	if err !=nil {
		///log
		return def
	}
	return d
}
