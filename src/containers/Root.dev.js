import React, { Component } from 'react';
import { Provider } from 'react-redux';
import Main from '../index';
import DevTools from './DevTools';
import configureStore from '../redux/store/configureStore';
import { initialState } from '../redux/reducers';

export default class Root extends Component {
  render() {
    const store = configureStore(initialState)

    return (
      <Provider store={store}>
        <div>
            <h1>Hello World...</h1>
          <DevTools />
        </div>
      </Provider>
    );
  }
}