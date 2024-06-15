package appWebsocket

import (
	"Systemge/Application"
	"Systemge/Client"
	"Systemge/Message"
	"Systemge/Utilities"
	"Systemge/WebsocketClient"
	"SystemgeSampleChess/topics"
)

type WebsocketApp struct {
	client *Client.Client
}

func New(messageBrokerClient *Client.Client, args []string) (Application.WebsocketApplication, error) {
	return &WebsocketApp{
		client: messageBrokerClient,
	}, nil
}

func (app *WebsocketApp) OnStart() error {
	return nil
}

func (app *WebsocketApp) OnStop() error {
	return nil
}

func (app *WebsocketApp) GetAsyncMessageHandlers() map[string]Application.AsyncMessageHandler {
	return map[string]Application.AsyncMessageHandler{
		topics.MOVE: func(message *Message.Message) error {
			println(app.client.GetName() + " received message from: " + message.GetOrigin())
			return nil
		},
	}
}

func (app *WebsocketApp) GetSyncMessageHandlers() map[string]Application.SyncMessageHandler {
	return map[string]Application.SyncMessageHandler{}
}

func (app *WebsocketApp) GetCustomCommandHandlers() map[string]Application.CustomCommandHandler {
	return map[string]Application.CustomCommandHandler{}
}

func (app *WebsocketApp) GetWebsocketMessageHandlers() map[string]Application.WebsocketMessageHandler {
	return map[string]Application.WebsocketMessageHandler{
		topics.MOVE: func(client *WebsocketClient.Client, message *Message.Message) error {
			/* groups := app.messageBrokerClient.GetWebsocketServer().GetGroups(client.GetId())
			if len(groups) != 1 {
				return Utilities.NewError("Expected exactly one group for client", nil)
			}
			chessRoom := groups[0]
			topic := "move_" + chessRoom
			err := app.messageBrokerClient.AsyncMessage(topic, client.GetId(), message.GetPayload())
			if err != nil {
				return Utilities.NewError("Error sending async message", err)
			} */
			return nil
		},
	}
}

func (app *WebsocketApp) OnConnectHandler(connection *WebsocketClient.Client) {
	_, err := app.client.SyncMessage(topics.NEW, connection.GetId(), "")
	if err != nil {
		panic(Utilities.NewError("Error sending sync message", err))
	}
}

func (app *WebsocketApp) OnDisconnectHandler(connection *WebsocketClient.Client) {
}
