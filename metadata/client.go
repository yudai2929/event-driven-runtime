package metadata

// Client is the interface that wraps the basic Metadata methods.
type Client interface {
	ListCronSettings() []*CronSetting
}

type metadataClient struct {
	cronMetadata map[string]CronSetting
}

// CronSetting represents a cron setting.
type CronSetting struct {
	Schedule     string
	FunctionName string
}

// NewClient creates a new Metadata client.
func NewClient() Client {
	return &metadataClient{
		cronMetadata: map[string]CronSetting{
			"logging": {
				Schedule:     "@every 10s",
				FunctionName: "logging",
			},
		},
	}
}

// ListCronSettings lists cron settings.
func (m *metadataClient) ListCronSettings() []*CronSetting {
	var cronSettings []*CronSetting
	for _, v := range m.cronMetadata {
		cronSettings = append(cronSettings, &v)
	}
	return cronSettings
}
