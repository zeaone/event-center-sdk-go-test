package event

import "github.com/cloudevents/sdk-go/v2/event"

const (
	Action_Create = "create"
)
var _ EventWriter = (*IcEvent)(nil)


func (i *IcEvent) SetSource(s string) {
	i.Event.SetSource(s)
}

func (i *IcEvent) SetType(s string) {
	i.Event.SetType(s)
}

func (i *IcEvent) SetSubject(s string) {
	i.Event.SetSubject(s)
}

func (i *IcEvent) SetBody(i2 interface{}, contentType ...string) error {
	if err := i.Event.SetData(event.ApplicationJSON, i2); err != nil {
		return err
	}
	return nil
}

func (i *IcEvent) SetAction(s string) {
	i.Event.SetExtension("action", s)
}

func (i *IcEvent) SetRegion(s string) {
	i.Event.SetExtension("region", s)
}

func (i *IcEvent) SetAccountid(s string) {
	i.Event.SetExtension("accountid", s)
}
func (i *IcEvent) SetResource(s string) {
	i.Event.SetExtension("resource", s)
}
func (i *IcEvent) SetExtension(key string,value interface{}) {
	i.Event.SetExtension(key, value)
}