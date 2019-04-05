import { applyMiddleware, createStore, combineReducers } from 'redux';
import logger from 'redux-logger';
import createHistory from 'history/createBrowserHistory';
import { routerReducer, routerMiddleware } from 'react-router-redux';

import rootReducer from '../reducers/index';

const enhancers = applyMiddleware(logger, routerMiddleware(history));

export const history = createHistory();

export default function configureStore(initialState) {
    return createStore(combineReducers({reducer: rootReducer, routing: routerReducer}), initialState, enhancers);
}