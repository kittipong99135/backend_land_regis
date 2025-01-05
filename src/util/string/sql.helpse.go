package util_string

import (
	"fmt"
	"strconv"
)

type UtilsString struct{}

func UseUtilsString() UtilsString {
	return UtilsString{}
}

func (u UtilsString) IsStringPointer(str *string) string {
	if str != nil {
		return *str
	}
	return ""
}

func (u UtilsString) AutoIncarmentAccountId(maxId, maxIndex int) string {
	index := strconv.Itoa(maxIndex)
	nextId := fmt.Sprintf("%0"+index+"d", maxId+1)
	return nextId
}

func (u UtilsString) WhereCluaseConcatnate(data []map[string]map[string]string) string {
	// for index, val := range data {
	// 	log.Println(val[index])
	// }
	return ""
}
