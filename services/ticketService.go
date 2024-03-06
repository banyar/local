package services

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"core/formatter"
	"core/response"
	"core/utils"

	driver "git.frontiir.net/sa-dev/rtdatacore/pkg/driver"
)

// get Ticket by id ...
func GetTicketByID(id uint) (*response.FBytesResponse, *utils.RestErr) {
	dbConnection := ConnectDatabase()
	ticketAdapter := driver.NewMySQLTicketAdapter(dbConnection)
	fmt.Println(reflect.TypeOf(ticketAdapter))
	ticket, tktErr := ticketAdapter.TicketMetaRepo.GetByID(id)

	if tktErr != nil {
		return nil, utils.ErrResponse(tktErr)
	}

	customFieldNames := strings.Split(os.Getenv("CUSTOM_FIELD_NAMES"), ",")
	customFieldAdapter := driver.NewCustomFieldAdapter(dbConnection)
	customFields, cfErr := customFieldAdapter.CustomFieldRepo.GetByNameList(customFieldNames, 0)

	if cfErr != nil {
		return nil, utils.ErrResponse(cfErr)
	}

	objCustomFieldValues, ocfvErr := customFieldAdapter.ObjectCustomFieldValueRepo.GetByTicketID(id, customFields.ExtractIDList(), 2)

	if ocfvErr != nil {
		return nil, utils.ErrResponse(ocfvErr)
	}

	customFieldsResponse := formatter.MapCustomFieldValueName(customFields, objCustomFieldValues)
	fBytesDetail := formatter.BuildFbytesResponse(ticket, customFieldsResponse)
	resResult := &response.FBytesResponse{Status: 200, Data: fBytesDetail}
	return resResult, nil
}
