package handler

import (
	"course_fullstack/backend/internal/entity"
	"encoding/json"
)

type sendMessageData struct {
	SenderID   string
	ReceiverID string
	Text       string
}

func (c *Client) handleSendMessage(request []byte) error {
	var message sendMessageData
	err := json.Unmarshal(request[len(sendMessageRequest):], &message)
	if err != nil {
		return entity.ErrInvalidRequest
	}
	if c.myID == "" || c.myID != message.SenderID {
		return entity.ErrPermissionDenied
	}
	return c.service.SendMessage(message.SenderID, message.ReceiverID, message.Text)
}

const sendMessageRequest = "send_message"

func isSendMessage(request []byte) bool {
	const reqType = sendMessageRequest
	prefixLen := len(reqType)
	if len(request) >= prefixLen && string(request[0:prefixLen]) == reqType {
		return true
	}
	return false
}

type markReadData struct {
	SenderID   string
	ReceiverID string
	Time       int64
}

func (c *Client) handleMarkRead(request []byte) error {
	var markRead markReadData
	err := json.Unmarshal(request[len(markReadRequest):], &markRead)
	if err != nil {
		return entity.ErrInvalidRequest
	}
	if c.myID == "" || c.myID != markRead.ReceiverID {
		return entity.ErrPermissionDenied
	}
	return c.service.MarkReadTill(markRead.SenderID, markRead.ReceiverID, markRead.Time)
}

const markReadRequest = "mark_read"

func isMarkRead(request []byte) bool {
	const reqType = markReadRequest
	prefixLen := len(reqType)
	if len(request) >= prefixLen && string(request[0:prefixLen]) == reqType {
		return true
	}
	return false
}

type getUserNameData struct {
	UserID string
}

func (c *Client) handleGetUserName(request []byte) error {
	var getUserName getUserNameData
	err := json.Unmarshal(request[len(getUserNameRequest):], &getUserName)
	if err != nil {
		return entity.ErrInvalidRequest
	}
	var user *entity.User
	user, err = c.service.GetUser(getUserName.UserID)
	if err != nil {
		return err
	}
	c.getUserNameChan <- &user.FrontendUser
	return nil
}

const getUserNameRequest = "get_user_name"

func isGetUserName(request []byte) bool {
	const reqType = getUserNameRequest
	prefixLen := len(reqType)
	if len(request) >= prefixLen && string(request[0:prefixLen]) == reqType {
		return true
	}
	return false
}

type LoginData struct {
	UserID   string
	Password string
}

func (c *Client) handleLogin(request []byte) error {
	var login LoginData
	err := json.Unmarshal(request[len(loginRequest):], &login)
	if err != nil {
		return entity.ErrInvalidRequest
	}
	c.myID = login.UserID
	err = c.service.Login(login.UserID, login.Password, c)
	if err != nil {
		c.myID = ""
		return err
	}
	c.otherChan <- "login success:" + c.myID
	return nil
}

const loginRequest = "login"

func isLogin(request []byte) bool {
	const reqType = loginRequest
	const prefixLen = len(reqType)
	if len(request) >= prefixLen && string(request[0:prefixLen]) == reqType {
		return true
	}
	return false
}

type registerData struct {
	LoginData
	UserName string
}

func (c *Client) handleRegister(request []byte) error {
	var register registerData
	err := json.Unmarshal(request[len(registerRequest):], &register)
	if err != nil {
		return entity.ErrInvalidRequest
	}
	c.myID = register.UserID
	err = c.service.Register(register.UserID, register.UserName, register.Password, c)
	if err != nil {
		c.myID = ""
		return err
	}
	c.otherChan <- "register success:" + c.myID
	return nil
}

const registerRequest = "register"

func isRegister(request []byte) bool {
	const reqType = registerRequest
	const prefixLen = len(reqType)
	if len(request) >= prefixLen && string(request[0:prefixLen]) == reqType {
		return true
	}
	return false
}

type getDataData struct {
	MyID   string
	ChatID string
}

func (c *Client) handleGetData(request []byte) error {
	var getData getDataData
	err := json.Unmarshal(request[len(getDataRequest):], &getData)
	if err != nil {
		return entity.ErrInvalidRequest
	}
	c.chatIDMutex.Lock()
	c.openedChatID = getData.ChatID
	c.chatIDMutex.Unlock()
	c.Ping()
	return nil
}

const getDataRequest = "get_data"

func isGetData(request []byte) bool {
	const reqType = getDataRequest
	const prefixLen = len(reqType)
	if len(request) >= prefixLen && string(request[0:prefixLen]) == reqType {
		return true
	}
	return false
}

func (c *Client) handleRequest(request []byte) {
	var err error
	if isLogin(request) {
		err = c.handleLogin(request)
	} else if isRegister(request) {
		err = c.handleRegister(request)
	} else if isSendMessage(request) {
		err = c.handleSendMessage(request)
	} else if isMarkRead(request) {
		err = c.handleMarkRead(request)
	} else if isGetUserName(request) {
		err = c.handleGetUserName(request)
	} else if isGetData(request) {
		err = c.handleGetData(request)
	} else {
		err = entity.ErrInvalidRequestType
	}
	if err != nil {
		c.errChan <- err
	}
}
