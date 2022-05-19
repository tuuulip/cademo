import axios from "axios";
import Vue from "vue";
import { Notification } from "element-ui";

// create an axios instance
const instance = axios.create({
  baseURL: process.env.VUE_APP_BASE_API, // url = base url + request url
  withCredentials: true, // send cookies when cross-domain requests
  timeout: 1200000 // request timeout
});

// catch response err
function handleResponse(response) {
  const res = response.data;
  if (res.code === 0) {
    return res;
  } else {
    return handleError(new Error(res.msg || "error"));
  }
}

// cat response err
function handleError(error) {
  Notification({
    title: "Server return error",
    message: error ? error.toString() : "error",
    type: "error"
  });
  return Promise.reject(error || "error");
}
instance.interceptors.response.use(handleResponse, handleError);

// set vue plugin
const requestPlugin = {};
requestPlugin.install = function(Vue) {
  Vue.prototype.$request = instance;
};

Vue.use(requestPlugin);

export default instance;
