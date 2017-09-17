/**
 * Created by nwu on 9/17/17.
 */
import * as types from '../session/types';
import initialState from './initialState';

export default function sessionReducer(state = initialState, action) {
  switch(action.type) {
    case types.LOG_IN_SUCCESS:
      return {token: sessionStorage.getItem("token"), error: "", user: action.user};
    case types.LOG_IN_FAILURE:
      return {token: "", error: action.message, user: null};
    case types.SIGN_UP_SUCCESS:
      return {token: sessionStorage.getItem("token"), error: "", user: action.user};
    case types.SIGN_UP_FAILURE:
      return {token: "", error: action.message, user: null};
    default:
      return state;
  }
}