package nats

import (
	"context"
	"fmt"

   "github.com/nats-io/nats.go"
	//"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
)

var triggerMd = trigger.NewMetadata(&Settings{}, &HandlerSettings{}, &Output{})

func init() {
	_ = trigger.Register(&Trigger{}, &Factory{})
}

type Trigger struct {
	settings *Settings
	nconn    *nats.Conn
	natsHandlers []*Handler
}

type Factory struct {
}

func (*Factory) New(config *trigger.Config) (trigger.Trigger, error) {
	s := &Settings{}
	err := metadata.MapToStruct(config.Settings, s, true)
	if err != nil {
		return nil, err
	}

	return &Trigger{settings: s}, nil
}

func (f *Factory) Metadata() *trigger.Metadata {
	return triggerMd
}

// Metadata implements trigger.Trigger.Metadata
func (t *Trigger) Metadata() *trigger.Metadata {
	return triggerMd
}

func (t *Trigger) Initialize(ctx trigger.InitContext) error {

	logger := ctx.Logger()
	var err error

	natsURL := fmt.Sprintf("nats://%s:%d",t.settings.Host,t.settings.Port)

	t.nconn, err = nats.Connect(natsURL)
	if err != nil {
		return err
	}
	logger.Infof("Connected to nats: %s",natsURL)

	// Init handlers
	for _, handler := range ctx.GetHandlers() {
		natsHandler, err := NewNatsHandler(ctx.Logger(), handler, t.nconn)
		if err != nil {
			return err
		}
		t.natsHandlers = append(t.natsHandlers, natsHandler)
	}

	return nil
}

// Start implements util.Managed.Start
func (t *Trigger) Start() error {
	for _, handler := range t.natsHandlers {
		_ = handler.Start()
	}

	return nil
}

// Stop implements util.Managed.Stop
func (t *Trigger) Stop() error {
	for _, handler := range t.natsHandlers {
		_ = handler.Stop()
	}

	t.nconn.Close()

	return nil

}

//NewNatsHandler creates a new nats handler to handle a topic
func NewNatsHandler(logger log.Logger, handler trigger.Handler,  nc *nats.Conn) (*Handler, error) {
	natsHandler := &Handler{logger: logger, shutdown: make(chan struct{}), handler: handler}

	handlerSetting := &HandlerSettings{}
	err := metadata.MapToStruct(handler.Settings(), handlerSetting, true)
	if err != nil {
		return nil, err
	}

	if handlerSetting.Topic == "" {
		return nil, fmt.Errorf("topic string was not provided for handler: [%s]", handler)
	}

	logger.Debugf("Subscribing to topic [%s]", handlerSetting.Topic)

	//TODO - need to think how to create/contain nats subscribers..
	//and return inside handler

	//6/1/2020 - let's have a go with a chan subscriber
	//returns *Subscription
	subchan := make(chan *nats.Msg, 64)
	sub, err := nc.ChanSubscribe(handlerSetting.Topic, subchan)

	if err != nil {
		return nil, err
	}

	natsHandler.subscribers = append(natsHandler.subscribers, Subscriber{subscriber: sub, ch: subchan})

	return natsHandler, nil
}

type Subscriber struct {
	subscriber *nats.Subscription
	ch chan *nats.Msg
}

type Handler struct {
	shutdown chan struct{}
	logger log.Logger
	handler trigger.Handler
	subscribers []Subscriber
}

//TBD.....
func (h *Handler) consumeTopic(s Subscriber) {
	for {
		select {
		case <-h.shutdown:
			return
		case msg := <-s.ch:
			out := &Output{}
			out.Message = string(msg.Data)
			h.logger.Errorf("Received message data: [%v]", out.Message)

			_, err := h.handler.Handle(context.Background(), out)
			if err != nil {
				h.logger.Errorf("Run action for handler [%s] failed for reason [%s] message lost", h.handler.Name(), err)
			}
		}
	}
}

//Start starts the handler
func (h *Handler) Start() error {
	//TBD iterate over subscribers....?
	for _, s := range h.subscribers {
		go h.consumeTopic(s)
	}

	return nil
}

//Stop stops the handler
func (h *Handler) Stop() error {
	//TBD iterate over subscribers....?

	close(h.shutdown)

	for _, s := range h.subscribers {
		//catch err and return???
		s.subscriber.Unsubscribe()
	}

	return nil
}


