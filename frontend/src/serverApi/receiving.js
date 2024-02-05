import {setMessages} from "../redux/chatMessages";
import {setChatPreviews} from "../redux/chatPreviews";
import {addUserName, setUserNames} from "../redux/userNames";
import {store} from "../store";
import {setMyId} from "../redux/myId";
import {setChatId} from "../redux/chatId";
import {fetchData} from "./sending";
import {setInfoModal} from "../redux/modal";



export function checkLogin(msg) {
  console.log("login checking")
  const msgType = "login success:"
  if (msg.length >= msgType.length && msg.substring(0, msgType.length) === msgType) {
    const id = msg.substring(msgType.length)
    store.dispatch(setMyId(id))
    store.dispatch(setChatId(id))
    store.dispatch(setInfoModal("login success to @" + id))
    console.log("logged in")
    fetchData(id, id)
    return true
  }
  return false
}

export function checkRegister(msg) {
  const msgType = "register success:"
  if (msg.length >= msgType.length && msg.substring(0, msgType.length) === msgType) {
    const id = msg.substring(msgType.length)
    store.dispatch(setMyId(id))
    store.dispatch(setChatId(id))
    store.dispatch(setInfoModal("register success for @" + id))
    fetchData(id, id)
    return true
  }
  return false
}

export function checkGotUserName(msg) {
  const msgType = "user:"
  if (msg.length >= msgType.length && msg.substring(0, msgType.length) === msgType) {
    const {ID, Name} = JSON.parse(msg.substring(msgType.length))
    store.dispatch(addUserName({id: ID, name: Name}))
    store.dispatch(setChatId(ID))
    fetchData(ID, ID)
    return true
  }
  return false
}

export function checkGotError(msg) {
  const msgType = "error:"
  if (msg.length >= msgType.length && msg.substring(0, msgType.length) === msgType) {
    store.dispatch(setInfoModal(msg))
    return true
  }
  return false
}

export function checkGotData(msg) {
  const msgType = "data:"
  if (msg.length >= msgType.length && msg.substring(0, msgType.length) === msgType) {
    const {ChatMessages, LastMessages, UserNames} = JSON.parse(msg.substring(msgType.length))
    const chatMessages = []
    const lastMessages = []
    const userNames = {}
    for (const ind in ChatMessages) {
      chatMessages.push(convertBackendMessage(ChatMessages[ind]))
    }
    for (const ind in LastMessages) {
      lastMessages.push(convertBackendMessage(LastMessages[ind]))
    }
    for (const ind in UserNames) {
      const user = UserNames[ind]
      userNames[user.ID] = user.Name
    }
    store.dispatch(setUserNames(userNames))
    store.dispatch(setChatPreviews(lastMessages))
    store.dispatch(setMessages(chatMessages))
    return true
  }
  return false
}

function convertBackendMessage(msg) {
  return {
    receiverId: msg.ReceiverID,
    senderId: msg.SenderID,
    dateTime: new Date(msg.Time),
    text: msg.Text,
    isRead: msg.IsRead,
    id: msg.ID
  }
}
