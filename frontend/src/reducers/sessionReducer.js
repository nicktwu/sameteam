/**
 * Created by nwu on 9/17/17.
 */
import * as types from '../session/types';
import initialState from './initialState';

export default function sessionReducer(state = initialState.session, action) {
  switch(action.type) {
    case types.LOG_IN_SUCCESS:
      return !!sessionStorage.token;
    default:
      return state;
  }
}