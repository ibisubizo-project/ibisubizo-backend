import React, { Component } from 'react';
import { Provider } from 'react-redux';
import Main from '../index';
import configureStore from '../redux/store/configureStore';
import { initialState } from '../redux/reducers';
import { Router, Route } from 'react-router-dom';

import createHistory from 'history/createBrowserHistory';

export default class Root extends Component {
  render() {
    const store = configureStore(initialState);



    return (
      <Provider store={store}>
        <div>Hello</div>
        {/* <Router>

        </Router> */}
      </Provider>
    );
  }
}