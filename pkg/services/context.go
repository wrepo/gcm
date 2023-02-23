package services

type (
	Context interface {
		//Data() *model.Data
		//SetData(d *model.Data) error
		SessionId() int
		SetSessionId(id int) error
		MetaData(key string) (interface{}, error)
		SetMetaData(key string, value interface{}) error
	}
)
