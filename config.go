package main

type Config struct {
	Env                             string                                `hcl:"env"`
	Region                          string                                `hcl:"region"`
	Profile                         string                                `hcl:"profile"`
	Service                         map[string]ServiceDetail              `hcl:"service"`
	Team                            map[string][]string                   `hcl:"team"`
	AllTeamsToTeamGrant             []string                              `hcl:"all_teams_to_team_grant"`
	TeamToTeamGrant                 map[string][]string                   `hcl:"team_to_team_grant"`
	AllServicesToTeamGrant          []string                              `hcl:"all_services_to_team_grant"`
	ServiceToTeamGrant              map[string][]string                   `hcl:"service_to_team_grant"`
	TeamToServiceGrant              map[string][]string                   `hcl:"team_to_service_grant"`
	PciReaderToTeamGrant            []string                              `hcl:"pci_reader_to_team_grant"`
	PciReaderToServiceGrant         []string                              `hcl:"pci_reader_to_service_grant"`
	Warehouse                       map[string]WarehouseDetail            `hcl:"warehouse"`
	DataSource                      map[string]DataSourceDetail           `hcl:"data_source"`
	DataMart                        map[string]DataMartDetail             `hcl:"data_mart"`
	ServiceNetworkPolicy            map[string]ServiceNetworkPolicyDetail `hcl:"service_network_policy"`
	AccountNetworkPolicyAllowedCidr []string                              `hcl:"account_network_policy_allowed_cidr"`
	AccountNetworkPolicyBlockedCidr []string                              `hcl:"account_network_policy_blocked_cidr"`
	AllAccessTeams                  AllAccessTeamsDetail                  `hcl:"all_access_teams"`
	AllAccessServices               AllAccessServicesDetail               `hcl:"all_access_services"`
	Buckets                         map[string]BucketsDetail              `hcl:"buckets"`
}

type ServiceDetail struct {
	DefaultWarehouse string `cty:"default_warehouse"`
}

type WarehouseDetail struct {
	Size                      string   `cty:"size"`
	RelationshipService       []string `cty:"relationship_service"`
	RelationshipTeam          []string `cty:"relationship_team"`
	RelationshipTeamMonitor   []string `cty:"relationship_team_monitor"`
	MinClusterCount           int      `cty:"min_cluster_count"`
	MaxClusterCount           int      `cty:"max_cluster_count"`
	StatementTimeoutInSeconds int      `cty:"statement_timeout_in_seconds"`
}

type DataSourceDetail struct {
	DataSourceDomain          string   `cty:"data_source_domain"`
	DataSourceName            string   `cty:"data_source_name"`
	StorageIntegration        string   `cty:"storage_integration"`
	ServiceOperator           []string `cty:"relationship_service_operator"`
	ServiceOperatorCreator    []string `cty:"relationship_service_operator_creator"`
	ServiceConsumer           []string `cty:"relationship_service_consumer"`
	TeamDeveloperCreator      []string `cty:"relationship_team_developer_creator"`
	TeamDeveloper             []string `cty:"relationship_team_developer"`
	TeamConsumer              []string `cty:"relationship_team_consumer"`
	TeamAnalyst               []string `cty:"relationship_team_analyst"`
	DatabaseTimeRetentionDays int      `cty:"database_time_retention_days"`
}

type DataMartDetail struct {
	DataMartDomain     string   `cty:"data_mart_domain"`
	DataMartName       string   `cty:"data_mart_name"`
	StorageIntegration string   `cty:"storage_integration"`
	ServiceOperator    []string `cty:"relationship_service_operator"`
	ServiceOperatorSub []string `cty:"relationship_service_operator_sub"`
	ServiceConsumer    []string `cty:"relationship_service_consumer"`
	ServiceConsumerSub []string `cty:"relationship_service_consumer_sub"`
	TeamDeveloper      []string `cty:"relationship_team_developer"`
	TeamConsumer       []string `cty:"relationship_team_consumer"`
	TeamAnalyst        []string `cty:"relationship_team_analyst"`
}

type ServiceNetworkPolicyDetail struct {
	NetworkPolicyName       string   `cty:"network_policy_name"`
	NetworkPolicyComment    string   `cty:"network_policy_comment"`
	AllowedIpList           []string `cty:"allowed_ip_list"`
	BlockedIpList           []string `cty:"blocked_ip_list"`
	RelationshipServiceUser []string `cty:"relationship_service_user"`
}

type AllAccessTeamsDetail struct {
	RelationshipTeamDeveloper     []string `cty:"relationship_team_developer"`
	RelationshipTeamAnalyst       []string `cty:"relationship_team_analyst"`
	RelationshipTeamConsumer      []string `cty:"relationship_team_consumer"`
	RelationshipTeamPowerConsumer []string `cty:"relationship_team_power_consumer"`
}

type AllAccessServicesDetail struct {
	RelationshipServiceOperator        []string `cty:"relationship_service_operator"`
	RelationshipServiceConsumer        []string `cty:"relationship_service_consumer"`
	RelationshipServiceOperatorAllPrep []string `cty:"relationship_service_operator_all_prep"`
	RelationshipServiceOperatorAllMart []string `cty:"relationship_service_operator_all_mart"`
	RelationshipServiceConsumerAllPrep []string `cty:"relationship_service_consumer_all_prep"`
	RelationshipServiceConsumerAllMart []string `cty:"relationship_service_consumer_all_mart"`
}

type BucketsDetail struct {
	Name       string `cty:"name"`
	PolicyName string `cty:"policy_name"`
}
