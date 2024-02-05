package handler

import (
	"course_fullstack/backend/internal/entity"
	"course_fullstack/backend/internal/service"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type Client struct {
	service *service.Service
	conn    *websocket.Conn

	pingChan chan struct{}

	// channels to write request responses
	getUserNameChan chan *entity.FrontendUser // one username
	dataChan        chan *AllDataResponse     // all frontend required data - opened chat messages, last messages from all started chats, all usernames for started chats
	errChan         chan error                // user exists, user not found etc
	otherChan       chan string               // login success, register success messages

	stopChan chan struct{}

	chatIDMutex  sync.RWMutex
	openedChatID string
	myID         string
}

func (c *Client) closeConnection() {
	log.Println("closing connection")
	c.service.Logout(c.myID)
	c.conn.Close()
	close(c.stopChan)
}

func NewClient(service *service.Service, conn *websocket.Conn) *Client {
	return &Client{
		service: service,
		conn:    conn,

		pingChan: make(chan struct{}, 1), // add len to not ping if another ping not yet handled

		dataChan:        make(chan *AllDataResponse),
		errChan:         make(chan error),
		otherChan:       make(chan string),
		getUserNameChan: make(chan *entity.FrontendUser),

		stopChan: make(chan struct{}),
	}
}

func (c *Client) Ping() {
	if len(c.pingChan) >= 1 {
		return
	}
	select {
	case c.pingChan <- struct{}{}:
	case <-c.stopChan:
	}
}

func (c *Client) ListenSocket() {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		c.handleRequest(message)
	}
	c.closeConnection()
}

func (c *Client) WriteSocket() {
	for {
		select {
		case data := <-c.dataChan:
			resp := []byte("data:")
			jsonData, err := json.Marshal(data)
			resp = append(resp, jsonData...)
			if err != nil {
				c.errChan <- err
			}
			c.conn.WriteMessage(websocket.TextMessage, resp)
		case err := <-c.errChan:
			resp := []byte("error:")
			resp = append(resp, err.Error()...)
			c.conn.WriteMessage(websocket.TextMessage, resp)
		case user := <-c.getUserNameChan:
			resp := []byte("user:")
			jsonData, err := json.Marshal(user)
			resp = append(resp, jsonData...)
			if err != nil {
				c.errChan <- err
			}
			c.conn.WriteMessage(websocket.TextMessage, resp)
		case msg := <-c.otherChan:
			c.conn.WriteMessage(websocket.TextMessage, []byte(msg))
		case <-c.stopChan:
			return
		}
	}
}

func (c *Client) ListenPing() {
	for {
		select {
		case <-c.pingChan:
			data, err := c.fetchData()
			if err != nil {
				select {
				case <-c.stopChan:
					return
				case c.errChan <- err:
				}
			} else {
				select {
				case <-c.stopChan:
					return
				case c.dataChan <- data:
				}
			}
		case <-c.stopChan:
			return
		}
	}
}

type AllDataResponse struct {
	ChatMessages []entity.FrontendMessage
	LastMessages []entity.FrontendMessage
	UserNames    []entity.FrontendUser
}

func (c *Client) fetchData() (*AllDataResponse, error) {
	var chatMessages []entity.Message
	var lastMessages []entity.Message
	var userNames []entity.User
	var err error
	c.chatIDMutex.RLock()
	chatID := c.openedChatID
	c.chatIDMutex.RUnlock()
	chatMessages, err = c.service.GetChatMessages(c.myID, chatID)
	if err != nil {
		return nil, err
	}
	lastMessages, err = c.service.GetLastMessages(c.myID)
	if err != nil {
		return nil, err
	}

	var userIDs []string
	for _, message := range lastMessages {
		var userID string
		if message.SenderID == c.myID {
			userID = message.ReceiverID
		} else {
			userID = message.SenderID
		}
		userIDs = append(userIDs, userID)
	}
	userIDs = append(userIDs, c.myID)
	userNames, err = c.service.GetUserList(userIDs)
	if err != nil {
		return nil, err
	}

	var responseData AllDataResponse
	for _, message := range chatMessages {
		responseData.ChatMessages = append(responseData.ChatMessages, message.FrontendMessage)
	}
	for _, message := range lastMessages {
		responseData.LastMessages = append(responseData.LastMessages, message.FrontendMessage)
	}
	for _, user := range userNames {
		responseData.UserNames = append(responseData.UserNames, user.FrontendUser)
	}
	return &responseData, nil
}
