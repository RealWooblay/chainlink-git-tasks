package env

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetString(key string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}

	value, exists := os.LookupEnv(key)
	if !exists {
		return "", errors.New("getEnvString: Missing environment variable '" + key + "'")
	}

	return value, nil
}

func GetInt(key string) (int, error) {
	strValue, err := GetString(key)
	if err != nil {
		return 0, err
	}

	intValue, err := strconv.Atoi(strValue)
	if err != nil {
		return 0, err
	}

	return intValue, nil
}

func GetInt64(key string) (int64, error) {
	strValue, err := GetString(key)
	if err != nil {
		return 0, err
	}

	int64Value, err := strconv.ParseInt(strValue, 10, 64)
	if err != nil {
		return 0, err
	}

	return int64Value, nil
}

func GetUInt64(key string) (uint64, error) {
	strValue, err := GetString(key)
	if err != nil {
		return 0, err
	}

	uint64Value, err := strconv.ParseUint(strValue, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint64Value, nil
}