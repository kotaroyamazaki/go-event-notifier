package entites

type serverity string

const (
	ServerityError serverity = "ERROR"
)

type resourceType string

const (
	ResourceTypeCloudRun      resourceType = "cloud_run_revision"
	ResourceTypeCloudFunction resourceType = "cloud_function"
	ResourceTypeAppEngine     resourceType = "gae_app"
)

type JsonPayload struct {
	Message     string `json:"error"`
	FirebaseUID string `json:"firebase_uid"`
	AppVersion  string `json:"app_version"`
	Caller      string `json:"caller"`
}

type Labels struct {
	ProjectID    string `json:"project_id"`
	RevisionName string `json:"revision_name"`
	ServiceName  string `json:"service_name"`
	FunctionName string `json:"function_name"`
	ExecutionID  string `json:"execution_id"`
	VersionID    string `json:"version_id"`
	ModuleID     string `json:"module_id"`
}
type Resource struct {
	Type   resourceType `json:"type"`
	Labels Labels       `json:"labels"`
}
type AppLog struct {
	InsertID    string      `json:"insertId"`
	JsonPayload JsonPayload `json:"jsonPayload"`
	Resource    Resource    `json:"resource"`
	TextPayload string      `json:"textPayload"`
	Timestamp   string      `json:"timestamp"`
	Severity    string      `json:"severity"`
	Trace       string      `json:"trace"`
}

func (a *AppLog) MakeLoggingURL() string {
	return "https://console.cloud.google.com/logs/query;query=timestamp%3D%22" + a.Timestamp + "%22%0AinsertId%3D%22" + a.InsertID + "%22?project=" + a.Resource.Labels.ProjectID
}
