import {createSlice} from "@reduxjs/toolkit";

const initialState = {
  data : "",
  type : "",
}

const modalSlice = createSlice({
  name : "modal",
  initialState,
  reducers : {
    setLoginModal(state) {
      state.type = "login"
    },
    setInfoModal(state, action) {
      state.type = "info"
      state.data = action.payload
    },
    setModalClose(state) {
      state.type = ""
      state.data = ""
    },
  }
})

export const modalReducer = modalSlice.reducer
export const { setLoginModal, setInfoModal, setModalClose } = modalSlice.actions
