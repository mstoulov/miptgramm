import { createSlice } from "@reduxjs/toolkit";
import { schema } from "normalizr";

const userIdSchema = new schema.Entity("userId")
const userNameSchema = new schema.Entity("userName")

const userSchema = new schema.Entity("user", {
  id : userIdSchema,
  name : userNameSchema,
})

const initialState = {
  userNames : {}
}

const userNamesSlice = createSlice({
  name: "userNames",
  initialState,
  reducers: {
    setUserNames: (state, action) => {
      state.userNames = {...state.userNames, ...action.payload}
    },
    addUserName: (state, action) => {
      const copy = {...state.userNames}
      copy[action.payload.id] = action.payload.name
      state.userNames = copy
    },
  },
});

export const userNamesReducer = userNamesSlice.reducer
export const { setUserNames, addUserName } = userNamesSlice.actions
