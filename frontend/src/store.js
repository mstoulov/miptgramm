import { configureStore } from "@reduxjs/toolkit";
import { chatPreviewsReducer } from "./redux/chatPreviews";
import { chatMessagesReducer } from "./redux/chatMessages";
import { userNamesReducer } from "./redux/userNames";
import {myIdReducer} from "./redux/myId";
import {chatIdReducer} from "./redux/chatId";
import {modalReducer} from "./redux/modal";
import {scrollPinReducer} from "./redux/scrollPin";

export const store = configureStore({
  reducer: {
    chatPreviews: chatPreviewsReducer,
    chatMessages: chatMessagesReducer,
    userNames: userNamesReducer,
    myId: myIdReducer,
    chatId: chatIdReducer,
    modal: modalReducer,
    scrollPin: scrollPinReducer,
  },
  devTools: true,
});
