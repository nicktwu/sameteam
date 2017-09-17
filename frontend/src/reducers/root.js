/**
 * Created by nwu on 9/17/17.
 */
import {combineReducers} from 'redux';
import session from './sessionReducer';


const rootReducer = combineReducers({
  session,
});

export default rootReducer;