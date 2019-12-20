package nats

import (
	"context"

   "github.com/nats-io/nats.go"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/trigger"
)

var triggerMd = trigger.NewMetadata(&Settings{}, &HandlerSettings{}, &Output{}, &Reply{})

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
func NewNatsHandler(logger log.Logger, handler trigger.Handler,  nc *nats.Connection) (*Handler, error) {
	natsHandler := &Handler{logger: logger, shutdown: make(chan struct{}), handler: handler}

	handlerSetting := &HandlerSetting{}
	err := metadata.MapToStruct(handler,Settings(), handlerSetting, true)
	if err != nil {
		return nil, err
	}

	if handlerSetting.Topic == "" {
		return nil, fmt.Errorf("topic string was not provided for handler: [%s]", handler)
	}

	logger.Debugf("Subscribing to topic [%s]", handlerSetting.Topic)

	//TODO - need to think how to create/contain nats subscribers..
	//and return inside handler

	return natsHandler, nil
}

type Handler struct {
	shutdown chan struct{}
	logger log.Logger
	handler trigger.Handler
	//TBD nats subscriber....
}

//TBD.....
func (h *Handler) consumeTopic() {
}

//Start starts the handler
func (h *Handler) Start() error {
	//TBD iterate over subscribers....?

	return nil
}

//Stop stops the handler
func (h *Handler) Stop() error {
	//TBD iterate over subscribers....?

	return nil
}


