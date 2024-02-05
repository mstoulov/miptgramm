import {store} from "../store";
import {conn} from "./socket";

export function sendMessage(senderId, receiverId, text) {
  conn.send("send_message" + `{ "senderID": "${senderId}", "receiverID": "${receiverId}", "text": "${text}" }`)
}

export function markRead(lastMessage, chatId) {
  if (lastMessage === undefined || lastMessage === null) {
    return
  }
  const myId = store.getState().myId.myId
  conn.send("mark_read" + `{ "senderID": "${chatId}", "receiverID": "${myId}", "time": ${lastMessage.dateTime.getTime()} }`)
}

export function fetchUserName(userId) {
  conn.send("get_user_name" + `{ "userID": "${userId}" }`)
  console.log("fetching user name")
}

export function fetchData(myId, chatId) {
  console.log("fetching data")
  conn.send("get_data" + `{ "myID": "${myId}", "chatID": "${chatId}" }`)
}

export function loginToServer(myId, password) {
  conn.send("login" + `{ "userID": "${myId}", "password": "${password}" }`)
  console.log("login sending")
}

export function registerToServer(id, userName, password) {
  conn.send("register" + `{ "userID": "${id}", "password": "${password}", "userName": "${userName}" }`)
}
