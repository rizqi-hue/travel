import { configureStore, ThunkAction, Action } from "@reduxjs/toolkit";
import TravelSlice from "../modules/Travel/Services/TravelSlice";

export const store = configureStore({
  reducer: {
    travel: TravelSlice,
  },
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;
