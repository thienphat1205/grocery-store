package repositories

import (
	"my-store/internal/database"
)

func NewFactory(mongoClient database.Instance, mainDbName string) Factory {
	return Factory{
		mongoClient: mongoClient,
		mainDbName:  mainDbName,
	}
}

type Factory struct {
	mongoClient database.Instance
	mainDbName  string
}

func (f Factory) Client() database.Instance {
	return f.mongoClient
}

func (f Factory) SessionManager() database.SessionManager {
	return f.mongoClient
}

// ========================== Ticket ==========================

// func (f Factory) TicketRepo() TicketRepo {
// 	return NewTicketRepo(f.mongoClient, f.mainDbName, "ticket")
// }

// func (f Factory) TicketExplanationRepo() TicketExplanationRepo {
// 	return NewTicketExplanationRepo(f.mongoClient, f.mainDbName, "ticket_explanation")
// }

// func (f Factory) TicketActionHistoryRepo() TicketActionHistoryRepo {
// 	return NewTicketActionHistoryRepo(f.mongoClient, f.mainDbName, "ticket_action_history")
// }

// ========================== Issue ==========================

// func (f Factory) IssueRepo() IssueRepo {
// 	return NewIssueRepo(f.mongoClient, f.mainDbName, "issue")
// }

// func (f Factory) HubIssueRepo() HubIssueRepo {
// 	return NewHubIssue(f.mongoClient, f.mainDbName, "hub_issue")
// }

// func (f Factory) SortIssueRepo() SortIssueRepo {
// 	return NewSortIssueRepo(f.mongoClient, f.mainDbName, "sorting_issue")
// }

// func (f Factory) OrderEventRepo() OrderEventRepo {
// 	return NewOrderEventRepo(f.mongoClient, f.mainDbName, "order_event")
// }

// ========================== Metadata ==========================

// func (f Factory) IssueTypeRepo() IssueTypeRepo {
// 	return NewIssueTypeRepo(f.mongoClient, f.mainDbName, "metadata_issue_type")
// }

// ========================== Config ==========================

// func (f Factory) MetadataHubIssueRepo() MetadataHubIssueRepo {
// 	return NewMetadataHubIssueRepo(f.mongoClient, f.mainDbName, "config_hub_issue")
// }

// func (f Factory) ConfigSortingIssueRepo() ConfigSortingIssueRepo {
// 	return NewConfigSortingIssueRepo(f.mongoClient, f.mainDbName, "config_sorting_issue")
// }

// func (f Factory) MetadataVersionRepo() MetadataVersionRepo {
// 	return NewMetadataVersionRepo(f.mongoClient, f.mainDbName, "metadata_version")
// }

// func (f Factory) ConfigMetadataVersionRepo() ConfigMetadataVersionRepo {
// 	return NewConfigMetadataVersionRepo(f.mongoClient, f.mainDbName, "config_metadata_version")
// }

// func (f Factory) ConfigWorkflowRepo() ConfigWorkflowRepo {
// 	return NewConfigWorkflowRepo(f.mongoClient, f.mainDbName, "config_workflow")
// }

// ========================== customer complaint ==========================

// func (f Factory) CustomerComplaintRepo() CustomerComplaintRepo {
// 	return NewCustomerComplaintRepo(f.mongoClient, f.mainDbName, "customer_complaint")
// }
