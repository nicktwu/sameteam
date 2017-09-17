/**
 * Created by nwu on 9/17/17.
 */
import {LOGIN_URL, SIGNUP_URL} from './urls';

class SessionAPI {
  static login(credentials) {
    const request = new Request(LOGIN_URL, {
      method: 'POST',
      headers: new Headers({
        'Content-Type': 'application/json'
      }),
      body: JSON.stringify({username: credentials.username, password: credentials.password})
    });


    return fetch(request).then(response => {
      return response.json();
    }).catch(error => {
      return error;
    });
  }

  static signup(credentials) {
    const request = new Request(SIGNUP_URL, {
      method: 'POST',
      headers: new Headers({
        'Content-Type': 'application/json'
      }),
      body: JSON.stringify({username: credentials.username, password: credentials.password})
    });


    return fetch(request).then(response => {
      return response.json();
    }).catch(error => {
      return error;
    });
  }
}

export default SessionAPI;