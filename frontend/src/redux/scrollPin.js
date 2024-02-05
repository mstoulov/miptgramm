import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  pinned : true,
}

const scrollPinSlice = createSlice({
  name: "scrollPin",
  initialState,
  reducers: {
    setPinned: (state, action) => {
      state.pinned = action.payload;
    },
  },
});

export const scrollPinReducer = scrollPinSlice.reducer
export const { setPinned} = scrollPinSlice.actions
