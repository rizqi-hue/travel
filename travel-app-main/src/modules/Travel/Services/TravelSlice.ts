/* eslint-disable @typescript-eslint/no-explicit-any */
import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import api from "../../../app/api";

export interface TravelInterface {
    field?: string;
    label?: string;
    total?: number;
}

export interface TravelParamInterface {
    search?: string;
}

export const getsTravel = createAsyncThunk(
    "getsTravel",
    async (params: { filter?: TravelParamInterface, pagination?: { page: number, perPage: number }, sort?: { field: string, order: string } }, thunkAPI) => {

        const config = {
            method: "get",
            url: "travelled_country",
            params
        };

        return api(config).then(
            (res: any) => {
                return res.data;
            },
            (err: any) => {
                return thunkAPI.rejectWithValue(err.response.data.message);
            }
        );
    }
);

const TravelSlice = createSlice({
    name: "TravelSlice",
    initialState: {
        data: {} as TravelInterface,
        list: {
            "GETS_TRAVEL": {
                "next": false,
                "back": true,
                "limit": 0,
                "totalPage": 0,
                "currentPage": 0,
                "sort": "",
                "order": "",
                "recordsTotal": 0,
                data: [{} as TravelInterface],
            }
        },
        isFetching: false,
        isSuccess: false,
        isError: false,
        errorMessage: "" as string,
    },
    reducers: {
        clearState: (state) => {
            state.isFetching = false;
            state.isSuccess = false;
            state.isError = false;
            state.errorMessage = "";
        },

    },
    extraReducers: (builder) => {
        // getsTravel
        builder.addCase(getsTravel.fulfilled, (state, { payload }) => {
            state.isFetching = false;
            state.isSuccess = true;
            state.list = payload;
            return state;
        });
        builder.addCase(getsTravel.rejected, (state, payload: any) => {
            state.isFetching = false;
            state.isError = true;
            state.errorMessage = payload.payload;
        });
        builder.addCase(getsTravel.pending, (state) => {
            state.isFetching = true;
        });
    },
});

export const { clearState } = TravelSlice.actions;
export default TravelSlice.reducer;
