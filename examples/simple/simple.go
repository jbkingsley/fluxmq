package simple

import (
	"fmt"
	"strings"

	"github.com/mainflux/fluxmq/pkg/auth"
	"github.com/mainflux/fluxmq/pkg/client"
	"go.uber.org/zap"
)

var _ auth.Handler = (*Handler)(nil)

// Handler implements mqtt.Handler interface
type Handler struct {
	logger *zap.Logger
}

// New creates new Event entity
func New(logger *zap.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

// AuthConnect is called on device connection,
// prior forwarding to the MQTT broker
func (h *Handler) AuthConnect(c *client.ClientInfo) error {
	h.logger.Info(fmt.Sprintf("AuthRegister() - clientID: %s, username: %s, password: %s", c.ID, c.Username, string(c.Password)))
	return nil
}

// AuthPublish is called on device publish,
// prior forwarding to the MQTT broker
func (h *Handler) AuthPublish(c *client.ClientInfo, topic *string, payload *[]byte) error {
	h.logger.Info(fmt.Sprintf("AuthPublish() - clientID: %s, topic: %s, payload: %s", c.ID, *topic, string(*payload)))
	return nil
}

// AuthSubscribe is called on device publish,
// prior forwarding to the MQTT broker
func (h *Handler) AuthSubscribe(c *client.ClientInfo, topics *[]string) error {
	h.logger.Info(fmt.Sprintf("AuthSubscribe() - clientID: %s, topics: %s", c.ID, strings.Join(*topics, ",")))
	return nil
}

// Connect - after client successfully connected
func (h *Handler) Connect(c *client.ClientInfo) {
	h.logger.Info(fmt.Sprintf("Register() - username: %s, clientID: %s", c.Username, c.ID))
}

// Publish - after client successfully published
func (h *Handler) Publish(c *client.ClientInfo, topic *string, payload *[]byte) {
	h.logger.Info(fmt.Sprintf("Publish() - username: %s, clientID: %s, topic: %s, payload: %s", c.Username, c.ID, *topic, string(*payload)))
}

// Subscribe - after client successfully subscribed
func (h *Handler) Subscribe(c *client.ClientInfo, topics *[]string) {
	h.logger.Info(fmt.Sprintf("Subscribe() - username: %s, clientID: %s, topics: %s", c.Username, c.ID, strings.Join(*topics, ",")))
}

// Unsubscribe - after client unsubscribed
func (h *Handler) Unsubscribe(c *client.ClientInfo, topics *[]string) {
	h.logger.Info(fmt.Sprintf("Unsubscribe() - username: %s, clientID: %s, topics: %s", c.Username, c.ID, strings.Join(*topics, ",")))
}

// Disconnect on conection lost
func (h *Handler) Disconnect(c *client.ClientInfo) {
	h.logger.Info(fmt.Sprintf("Disconnect() - client with username: %s and ID: %s disconenected", c.Username, c.ID))
}
