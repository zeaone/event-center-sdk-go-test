package event

type EventWriter interface {
	SetSource(string)
	SetType(string)
	SetSubject(string)
	SetBody(interface{}, ...string) error
	//以下为自定义字段
	SetAction(string)
	SetRegion(string)
	SetAccountid(string)
	SetResource(string)
	SetExtension(string, interface{})
}

type EventReader interface {
	Source()string
	Type()string
	Subject()string
	Body() []byte
	//以下为自定义字段
	Action() string
	Region() string
	Accountid() string
	Resource() string
	Extensions() map[string]interface{}
}