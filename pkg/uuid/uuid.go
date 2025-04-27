package uuid

import "github.com/google/uuid"

func GetUUID() string {
	number, err := uuid.NewRandom()
	for err != nil {
		number, err = uuid.NewRandom()
	}
	return number.String()
}
