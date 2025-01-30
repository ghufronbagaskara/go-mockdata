package data

import (
	"fmt"
	"math/rand"
)

func Generate(dataType string) any {
	switch dataType {
	case TYPE_NAME:
		return generateName()
	case TYPE_DATE:
		return generateDate()
	case TYPE_ADDRESS:
		return generateAddress()
	case TYPE_PHONE:
		return generatePhone()
	}
	
	return ""
}


func generateName() string {
	nameLen := len(name)

	index := rand.Intn(nameLen)
	return name[index]
}

func generateDate() string {
	return ""	
}

func generateAddress() string {
		
	streetLen := len(address[SUBTYPE_STREET])
	cityLen := len(address[SUBTYPE_CITY])

	streetIndex := rand.Intn(streetLen)
	cityIndex := rand.Intn(cityLen)
	number := rand.Intn(100)

	return fmt.Sprintf("Jl. %s No. %d, %s", address[SUBTYPE_STREET][streetIndex], number, address[SUBTYPE_CITY][cityIndex])
}

func generatePhone() string {
	return ""	
}
