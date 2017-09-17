import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import {Provider} from 'react-redux';
import registerServiceWorker from './registerServiceWorker';
import root from './reducers/root';
import {createStore} from 'redux';

let store = createStore(root);

ReactDOM.render(<Provider store={store}><App /></Provider>, document.getElementById('root'));
registerServiceWorker();
