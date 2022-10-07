// Code generated by excel-error-code-generation tool. DO NOT EDIT.
package errcode

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
)

type ErrorCode string

const (
	OK                             ErrorCode = "0"
	Unknown                        ErrorCode = "1000"
	MissingRequiredMetadata        ErrorCode = "1001"
	SessionExpired                 ErrorCode = "1002"
	NotAuthorize                   ErrorCode = "1003"
	PermissionDenied               ErrorCode = "1004"
	MethodNotFound                 ErrorCode = "1005"
	HubNotMatch                    ErrorCode = "1006"
	Invalid                        ErrorCode = "1007"
	DBError                        ErrorCode = "1008"
	InvalidWithError               ErrorCode = "1009"
	InvalidOrderCode               ErrorCode = "1010"
	OrderCodeEmpty                 ErrorCode = "1011"
	InvalidPartnerApiKey           ErrorCode = "1012"
	AppVersionIsNotSupported       ErrorCode = "1013"
	AppVersionNotFound             ErrorCode = "1014"
	AppVersionIsNotSupportedForHub ErrorCode = "1015"
	ApiPartnerNotFound             ErrorCode = "1016"
	InvalidRequest                 ErrorCode = "1017"
	MissingMandatory               ErrorCode = "1018"
	CallOnlineError                ErrorCode = "1019"
	TicketNotFound                 ErrorCode = "20000"
	CreatedTicketExplanationFailed ErrorCode = "20001"
	ExplanationRequireAttachments  ErrorCode = "20002"
	ExplanationTypeIsNotDefined    ErrorCode = "20003"
	IssueTypeIsNotDefined          ErrorCode = "20004"
	IssueIsNotDefined              ErrorCode = "20005"
	WorkflowIsNotDefined           ErrorCode = "20006"
	NotAllowAccessTicket           ErrorCode = "20007"
	TicketStatusClosed             ErrorCode = "20008"
	TicketStatusProcessing         ErrorCode = "20009"
	TicketExplanationNotFound      ErrorCode = "20010"
	NotAllowExplainTicket          ErrorCode = "20011"
	TicketExpiredNotFound          ErrorCode = "20012"
	TicketSyncPayrollIsProcessing  ErrorCode = "20013"
	TicketSyncPayrollIsComplete    ErrorCode = "20014"
	OrderHasBeenCreatedTicket      ErrorCode = "20015"
	EmployeeIsNotYourMember        ErrorCode = "20016"
	RequiredAssignedUserIds        ErrorCode = "20017"
	NotApproveAccessTicket         ErrorCode = "20018"
	TicketExplanationPhaseClosed   ErrorCode = "20019"
	RequireMessage                 ErrorCode = "20020"
	ExplanationCanNotUpdate        ErrorCode = "20021"
	ExplanationCanNotSubmit        ErrorCode = "20022"
	EformPropertyRequired          ErrorCode = "20023"
	IssueTypeIsNotSuitable         ErrorCode = "20024"
	EformFlowNotActive             ErrorCode = "20025"
	StatusNotExplaining            ErrorCode = "20026"
	IssueTriggerTimeInvalid        ErrorCode = "20027"
	InvalidPenaltyType             ErrorCode = "20028"
)

var errorCodeMap *ini.File

func init() {
	var err error
	if errorCodeMap, err = ini.Load("./resources/metadata/error_code.ini"); err != nil {
		log.Fatalln(err)
	}
}

func (c ErrorCode) Error() string {
	return c.String()
}

func (c ErrorCode) String() string {
	key := fmt.Sprintf("%s", string(c))
	return errorCodeMap.Section("").Key(key).String()
}
