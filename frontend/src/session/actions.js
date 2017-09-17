/**
 * Created by nwu on 9/17/17.
 */
import * as types from './types';
import sessionApi from '../api/sessionApi';

export function loginSuccess() {
  return {type: types.LOG_IN_SUCCESS}
}

export function loginFailure(message) {
  return {type: types.LOG_IN_FAILURE, message: message}
}

export function signupSuccess() {
  return {type: types.SIGN_UP_SUCCESS}
}

export function signupFailure(message) {
  return {type: types.SIGN_UP_FAILURE, message: message}
}

export function loginUser(dispatch) {
  return function(credentials) {
    return sessionApi.login(credentials).then(res => {
      if (res.token) {
        sessionStorage.setItem('token', res.token);
        dispatch(loginSuccess());
      } else {
        sessionStorage.removeItem('token');
        dispatch(loginFailure("Login failed"));
      }
    }).catch(error => {
      console.log(error);
      sessionStorage.removeItem('token');
      dispatch(loginFailure(error.toString()));
      throw(error);
    })
  }
}

export function signupUser(dispatch) {
  return function(credentials) {
    return sessionApi.signup(credentials).then(res => {
      if (res.token) {
        sessionStorage.setItem('token', res.token);
        dispatch(signupSuccess());
      } else {
        sessionStorage.removeItem('token');
        dispatch(signupFailure("Sign up failed"));
      }
    }).catch(error => {
      console.log(error);
      sessionStorage.removeItem('token');
      dispatch(signupFailure(error.toString()));
      throw(error);
    })
  }
}

export function logoutUser(dispatch) {
  return () => {sessionStorage.removeItem("token");}
}

export function isAuth() {
  let token = sessionStorage.getItem('token');
  if (token) {
    return token.length > 0
  } else {
    return false;
  }
}