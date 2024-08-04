/* eslint-disable @typescript-eslint/no-explicit-any */
import axios from "axios";

interface ApiInterface {
  method: string;
  url: string;
  data?: any;
  params?: any;
  withCredentials?: boolean;
}

const _api = axios.create({
  baseURL: `${import.meta.env.VITE_APP_BASE_URL}/v1/`,
});

_api.interceptors.request.use(
  function (config: any) {
    config.headers["Content-Type"] = `Application/json`;
    return config;
  },
  function (error: any) {
    return Promise.reject(error);
  }
);

const api = async ({
  method,
  url,
  data = null,
  params = null,
  withCredentials = false,
}: ApiInterface) => {
  try {
    const response = await _api({
      method,
      url,
      data,
      params,
      withCredentials,
    });

    return response;
  } catch (err: any) {
    if (err.response.status === 401) {
      // 
    }

    throw err;
  }
};

export default api;
