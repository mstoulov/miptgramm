import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  myId : ""
}

const myIdSlice = createSlice({
  name: "myId",
  initialState,
  reducers: {
    setMyId: (state, action) => {
      state.myId = action.payload;
    },
  },
});

export const myIdReducer = myIdSlice.reducer
export const { setMyId} = myIdSlice.actions
