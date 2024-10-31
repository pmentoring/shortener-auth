package common

type ApplicationContext struct {
	InstanceId string
	AppBaseUrl string
	SecretKey  string
}

func NewApplicationContext(instanceId, appBaseUrl string) *ApplicationContext {
	return &ApplicationContext{
		InstanceId: instanceId,
		AppBaseUrl: appBaseUrl,
	}
}
