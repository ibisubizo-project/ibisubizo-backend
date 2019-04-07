import React, { Component } from 'react';
import { Provider } from 'react-redux';
import {Router, Route, IndexRoute} from 'react-router-dom';
import DevTools from './DevTools';
import configureStore from '../redux/store/configureStore';
import { initialState } from '../redux/reducers';
import Main from './Main';

export default class Root extends Component {
  render() {
    const store = configureStore(initialState)

    return (
      <Provider store={store}>
        <Router>
          <Route path="/" component={Main} />
        </Router>
        <DevTools />
      </Provider>
    );
  }
}