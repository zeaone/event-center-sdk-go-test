package client

import (
	"context"
	"fmt"

	 srchttp "net/http"

	"github.com/zeaone/event-center-sdk-go-test/event"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/binding"
	"github.com/cloudevents/sdk-go/v2/protocol/http"
)

var EventSink string

var _ Client = (*icClient)(nil)

type Client interface {
	Send(ctx context.Context, event *event.IcEvent) error
	Receive(ctx context.Context, req *srchttp.Request) (*event.IcEvent, error)
}

type icClient struct {
	client cloudevents.Client
	url string
}

func Init(opt... Option) (*icClient, error) {
	c := &icClient{}
	httpClient, err := cloudevents.NewClientHTTP()
	if err != nil {
		return nil, err
	}
	c.client = httpClient
	if err = c.applyOption(opt...); err != nil {
		return nil, err
	}
	return  c, nil
}

func (i *icClient)applyOption(option... Option) error {
	for _, fn := range option{
		if err := fn(i) ; err != nil{
			return err
		}
	}
	return nil
}


func (i *icClient) Send(ctx1 context.Context, event *event.IcEvent) error {
	err := event.CheckActionBeforeSend()
	if err != nil {
		return err
	}
	ctx := cloudevents.ContextWithTarget(ctx1, i.url)
	result := i.client.Send(ctx, *event.Event)
	if cloudevents.IsUndelivered(result) {
		return  fmt.Errorf("failed to send, %v", result)
	}
	return nil
}

func (i *icClient) Receive(ctx context.Context, req *srchttp.Request) (*event.IcEvent, error) {
	m := http.NewMessageFromHttpRequest(req)
	e, eventErr := binding.ToEvent(context.TODO(), m)
	if eventErr != nil {
		return nil, eventErr
	}
	icEvent := &event.IcEvent{Event: e}
	return icEvent, nil
}


func NewClientHttp() (*icClient, error) {
	c := &icClient{}
	httpClient, err := cloudevents.NewClientHTTP()
	if err != nil {
		return nil, err
	}
	c.client = httpClient
	return c, nil
}

//func Send(ctx1 context.Context, event event.IcEvent) error {
//	c := &icClient{}
//	httpClient, err := cloudevents.NewClientHTTP()
//	if err != nil {
//		return err
//	}
//	c.client = httpClient
//
//	err = event.CheckActionBeforeSend()
//	if err != nil {
//		return err
//	}
//	ctx := cloudevents.ContextWithTarget(ctx1, EventSink)
//	result := c.client.Send(ctx, event.Event)
//	if cloudevents.IsUndelivered(result) {
//		return  fmt.Errorf("failed to send, %v", result)
//	}
//	return nil
//}



