import React, { Component } from 'react';
import { Provider } from 'react-redux';
import Main from './Main';
import configureStore from '../redux/store/configureStore';
import { initialState } from '../redux/reducers';
import { Router, Route } from 'react-router-dom';
import DevTools from './DevTools';

import {history} from '../redux/store/configureStore';


export default class Root extends Component {
  render() {
    const store = configureStore(initialState);

    return (
      <Provider store={store}>
        <Router history={history}>
          <Route path="/" component={Main} />
        </Router>
      </Provider>
    );
  }
}