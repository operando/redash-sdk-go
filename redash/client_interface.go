package redash

type IClient interface {
	// client
	SetConfig(*Config)

	// alerts
	ListAlerts(*ListAlertsInput) *ListAlertsOutput
	GetAlert(*GetAlertInput) *GetAlertOutput
	ListAlertSubscriptions(*ListAlertSubscriptionsInput) *ListAlertSubscriptionsOutput

	// dashboards
	ListDashboards(*ListDashboardsInput) *ListDashboardsOutput
	GetDashboard(*GetDashboardInput) *GetDashboardOutput
	GetPublicDashboard(*GetPublicDashboardInput) *GetPublicDashboardOutput
	GetDashboardTags(*GetDashboardTagsInput) *GetDashboardTagsOutput

	// data_sources
	ListDataSources(*ListDataSourcesInput) *ListDataSourcesOutput
	CreateDataSource(*CreateDataSourceInput) *CreateDataSourceOutput
	ListDataSourcesTypes(*ListDataSourcesTypesInput) *ListDataSourcesTypesOutput
	GetDataSource(*GetDataSourceInput) *GetDataSourceOutput
	UpdateDataSource(*UpdateDataSourceInput) *UpdateDataSourceOutput
	DeleteDataSource(*DeleteDataSourceInput) *DeleteDataSourceOutput
	GetDataSourceSchema(*GetDataSourceSchemaInput) *GetDataSourceSchemaOutput
	PauseDataSource(*PauseDataSourceInput) *PauseDataSourceOutput
	UnpauseDataSource(*UnpauseDataSourceInput) *UnpauseDataSourceOutput
	TestDataSource(*TestDataSourceInput) *TestDataSourceOutput

	// destinations
	ListDestinations(*ListDestinationsInput) *ListDestinationsOutput
	CreateDestination(*CreateDestinationInput) *CreateDestinationOutput
	ListDestinationsTypes(*ListDestinationsTypesInput) *ListDestinationsTypesOutput
	GetDestination(*GetDestinationInput) *GetDestinationOutput
	UpdateDestination(*UpdateDestinationInput) *UpdateDestinationOutput
	DeleteDestination(*DeleteDestinationInput) *DeleteDestinationOutput

	// events
	GetEvents(*GetEventsInput) *GetEventsOutput

	// favorites
	ListQueryFavorites(*ListQueryFavoritesInput) *ListQueryFavoritesOutput
	ListDashboardFavorites(*ListDashboardFavoritesInput) *ListDashboardFavoritesOutput

	// groups
	ListGroups(*ListGroupsInput) *ListGroupsOutput
	GetGroup(*GetGroupInput) *GetGroupOutput
	ListGroupMembers(*ListGroupMembersInput) *ListGroupMembersOutput

	// queries
	ListQueries(*ListQueriesInput) *ListQueriesOutput
	GetQuery(*GetQueryInput) *GetQueryOutput
	GetQuerySearch(*GetQuerySearchInput) *GetQuerySearchOutput
	GetQueryRecent(*GetQueryRecentInput) *GetQueryRecentOutput
	GetMyQueries(*GetMyQueriesInput) *GetMyQueriesOutput
	GetQueryTags(*GetQueryTagsInput) *GetQueryTagsOutput
	CreateQuery(*CreateQueryInput) *CreateQueryOutput
	ModifyQuery(*ModifyQueryInput) *ModifyQueryOutput
	DeleteQuery(*DeleteQueryInput) *DeleteQueryOutput

	// query_results
	GetQueryResult(*GetQueryResultInput) *GetQueryResultOutput
	GetJob(*GetJobInput) *GetJobOutput

	// query_snippets
	ListQuerySnippets(*ListQuerySnippetsInput) *ListQuerySnippetsOutput
	GetQuerySnippet(*GetQuerySnippetInput) *GetQuerySnippetOutput

	// settings
	GetOrganizationSettings(*GetOrganizationSettingsInput) *GetOrganizationSettingsOutput

	// users
	ListUsers(*ListUsersInput) *ListUsersOutput
	CreateUser(*CreateUserInput) *CreateUserOutput
	GetUser(*GetUserInput) *GetUserOutput
	UpdateUser(*UpdateUserInput) *UpdateUserOutput
	// DeleteUser(*DeleteUserInput) *DeleteUserOutput
	// InviteUser(*InviteUserInput) *InviteUserOutput
	// ResetUserPassword(*ResetUserPasswordInput) *ResetUserPasswordOutput
	// RegenerateUserApiKey(*RegenerateUserApiKeyInput) *RegenerateUserApiKeyOutput
	// DiableUser(*DiableUserInput) *DisableUserOutput
	// EnableUser(*EnableUserInput) *EnableUserOutput
}
