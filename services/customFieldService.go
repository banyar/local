package services

import (
	"fmt"
	"os"
	"strings"

	driver "git.frontiir.net/sa-dev/rtdatacore/pkg/driver"
)

func GetFormattedRootCauseValues() *map[string][]string {
	dbConnection := ConnectDatabase()
	customFieldAdapter := driver.NewCustomFieldAdapter(dbConnection)
	customFieldNames := strings.Split(os.Getenv("CUSTOM_FIELD_NAME_LIST"), ",")
	customFields, cfErr := customFieldAdapter.CustomFieldRepo.GetByNameList(customFieldNames, 0)

	if cfErr != nil {
		fmt.Println(cfErr)
	}

	customFieldValueAdapter := driver.NewCustomFieldValueAdapter(dbConnection)
	rootCauseValues, cfvErr := customFieldValueAdapter.CustomFieldValueRepo.GetCustomFieldValuesByCustomFieldIds(customFields.ExtractIDList())
	if cfvErr != nil {
		fmt.Println(cfvErr)
	}
	return rootCauseValues.MapNameToCategory()
}
