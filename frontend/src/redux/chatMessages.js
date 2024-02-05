import {createSlice} from "@reduxjs/toolkit";
import {schema} from "normalizr";

const isMessageReadSchema = new schema.Entity("isMessageRead")
const senderIdSchema = new schema.Entity("senderId")
const receiverIdSchema = new schema.Entity("receiverId")
const dateTimeSchema = new schema.Entity("dateTime")
const messageIdSchema = new schema.Entity("messageId")
const messageTextSchema = new schema.Entity("messageText")

export const messageSchema = new schema.Entity("message", {
  id : messageIdSchema,
  senderId : senderIdSchema,
  receiverId : receiverIdSchema,
  text : messageTextSchema,
  dateTime : dateTimeSchema,
  isRead : isMessageReadSchema,
})

const chatMessagesSchema = new schema.Entity("chatMessages", {
  messages : [messageSchema]
})

const initialState = {
  messages : []
}

const userNamesSlice = createSlice({
  name: "chatMessages",
  initialState,
  reducers: {
    setMessages: (state, action) => {
      state.messages = action.payload;
    },
  },
});

export const chatMessagesReducer = userNamesSlice.reducer
export const { setMessages } = userNamesSlice.actions
