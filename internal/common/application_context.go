package common

type ApplicationContext struct {
	InstanceId string
	AppBaseUrl string
	SecretKey  string
}

func NewApplicationContext(instanceId, appBaseUrl, secretKey string) *ApplicationContext {
	return &ApplicationContext{
		InstanceId: instanceId,
		AppBaseUrl: appBaseUrl,
		SecretKey:  secretKey,
	}
}
