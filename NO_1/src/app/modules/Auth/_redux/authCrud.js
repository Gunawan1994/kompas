import axios from "axios";
import {customAxios} from "../../../services/customAxios";
import qs from 'qs';

export const LOGIN_URL = "login";
export const REGISTER_URL = "api/auth/register";
export const REQUEST_PASSWORD_URL = "api/auth/forgot-password";

export const ME_URL = "api/me";

export function login(email, password) {
  const bodyForm = {
    email:email,
    password:password
  }
  return customAxios({
    url:LOGIN_URL,
    method:"POST",
    data:qs.stringify(bodyForm)
  })
}

export function register(email, fullname, username, password) {
  return axios.post(REGISTER_URL, { email, fullname, username, password });
}

export function requestPassword(emailData) {
    return customAxios({
        url:'/password_reset',
        method:'POST',
        data:qs.stringify({email:emailData}),
      })
}

export function getUserByToken() {
  // Authorization head should be fulfilled in interceptor.
  return axios.get(ME_URL);
}
