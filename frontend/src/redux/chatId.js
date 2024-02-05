import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  chatId : ""
}

const chatIdSlice = createSlice({
  name: "chatId",
  initialState,
  reducers: {
    setChatId: (state, action) => {
      state.chatId = action.payload;
    },
  },
});

export const chatIdReducer = chatIdSlice.reducer
export const { setChatId} = chatIdSlice.actions
