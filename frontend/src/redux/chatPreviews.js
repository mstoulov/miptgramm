import { createSlice } from "@reduxjs/toolkit";
import {schema} from "normalizr";
import {messageSchema} from "./chatMessages";


const previewListSchema = new schema.Entity("previewList", {
  lastMessages : [messageSchema],
})

const initialState = {
  lastMessages : []
}

const chatPreviewsSlice = createSlice({
  name: "chatPreviews",
  initialState,
  reducers: {
    setChatPreviews: (state, action) => {
      state.lastMessages = action.payload;
    },
  },
});

export const chatPreviewsReducer = chatPreviewsSlice.reducer
export const { setChatPreviews} = chatPreviewsSlice.actions
