package event

import (
	"fmt"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)


type IcEvent struct {
	Event *cloudevents.Event
}

func NewEvent() *IcEvent  {
	e := &IcEvent{}
	tmpEvent := cloudevents.NewEvent()
	e.Event = &tmpEvent
	return  e
}

func (i IcEvent)CheckActionBeforeSend() error {
	actionType := [4]string {"create", "list", "modify", "delete"}
	subject := i.Event.Subject()
	action, ok := i.Event.Extensions()["action"].(string)
	if !ok {
		return fmt.Errorf("filed <action> set error, expect type <string>" )
	}
	inActionFlag := false
	if subject == "operate" {
		for _, v := range actionType {
			if action == v {
				inActionFlag = true
				break
			}
		}
		if !inActionFlag {
			return fmt.Errorf("when subject is <operate>, the action must in %v", actionType )
		}
	}
	return nil
}
