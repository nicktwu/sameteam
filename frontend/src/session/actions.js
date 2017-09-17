/**
 * Created by nwu on 9/17/17.
 */
import * as types from './types';
import sessionApi from '../api/sessionApi';

export function loginSuccess() {
  return {type: types.LOG_IN_SUCCESS}
}

export function loginFailure() {
  return {type: types.LOG_IN_FAILURE}
}

export function signupSuccess() {
  return {type: types.SIGN_UP_SUCCESS}
}

export function signupFailure() {
  return {type: types.SIGN_UP_FAILURE}
}

export function loginUser(dispatch) {
  return function(credentials) {
    return sessionApi.login(credentials).then(res => {
      sessionStorage.setItem('token', res.token);
      dispatch(loginSuccess());
    }).catch(error => {
      console.log(error);
      sessionStorage.removeItem('token');
      dispatch(loginFailure());
      throw(error);
    })
  }
}

export function signupUser(dispatch) {
  return function(credentials) {
    return sessionApi.signup(credentials).then(res => {
      sessionStorage.setItem('token', res.token);
      dispatch(signupSuccess());
    }).catch(error => {
      console.log(error);
      sessionStorage.removeItem('token');
      dispatch(signupFailure());
      throw(error);
    })
  }
}

export function logoutUser(dispatch) {
  return () => {sessionStorage.removeItem("token");}
}

export function isAuth() {
  let token = sessionStorage.getItem('token');
  console.log(token);
  if (token) {
    return token.length > 0
  } else {
    return false;
  }
}