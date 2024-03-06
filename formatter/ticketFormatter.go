package formatter

import (
	"core/models"
	"core/response"
	"strings"

	"git.frontiir.net/sa-dev/rtdatacore/pkg/core/domain/entities"
)

func BuildTicketResponse(ticket *models.Ticket, cfValues map[string]string) *response.TicketDetail {
	ticketDetail := &response.TicketDetail{
		ID:           ticket.ID,
		CustomFields: cfValues,
	}
	return ticketDetail
}

func BuildFbytesResponse(ticket *entities.TicketMeta, cfValues map[string]string) *response.FBytesDetail {

	return &response.FBytesDetail{
		ID:           int32(ticket.ID),
		CustomFields: BuildFbytesCustomFiled(cfValues),
	}
}

func BuildFbytesCustomFiled(cfValues map[string]string) *response.FBytesCustomFields {
	return &response.FBytesCustomFields{
		CustomerId:       cfValues["customer_id"],
		LocalServiceId:   cfValues["local_service_id"],
		CpeId:            cfValues["cpe_id"],
		TicketProblem:    cfValues["ticket_problem"],
		ServiceType:      cfValues["service_type"],
		InstallationType: cfValues["installation_type"],
		MainRootCause: &response.MainRootCause{
			ServiceRootCause:      cfValues["service_root_cause"],
			RootCauseCategory:     cfValues["root_cause_category"],
			SuscpectedAreaOfIssue: cfValues["suspected_area_of_issue"],
		},
	}
}

func MapCustomFieldsIdName(cfs *entities.CustomFieldCollection) map[uint]string {
	cfMap := make(map[uint]string)
	for _, cf := range *cfs {
		cfMap[cf.ID] = cf.Name
	}
	return cfMap
}

func MapCustomFieldsIdValue(objCfs *entities.ObjectCustomFieldValueCollection) map[uint]string {
	idValueMap := make(map[uint]string)
	for _, objCf := range *objCfs {
		idValueMap[objCf.CustomFieldID] = objCf.Content
	}
	return idValueMap
}

func MapCustomFieldValueName(cfs *entities.CustomFieldCollection, objCfs *entities.ObjectCustomFieldValueCollection) map[string]string {

	idNameMap := MapCustomFieldsIdName(cfs)
	nameValueMap := make(map[string]string)
	idValueMap := MapCustomFieldsIdValue(objCfs)

	for i, cf := range idNameMap {
		formattedName := FormatCustomFieldName(cf)
		nameValueMap[formattedName] = idValueMap[i]
	}
	return nameValueMap
}

func FormatCustomFieldName(name string) string {
	name = strings.Trim(strings.ToLower(name), " ")
	return strings.ReplaceAll(name, " ", "_")
}
