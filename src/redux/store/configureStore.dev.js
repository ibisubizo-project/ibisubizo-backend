import { createStore, applyMiddleware, compose, combineReducers } from 'redux';
import { persistState } from 'redux-devtools';
import thunk from 'redux-thunk';
import rootReducer from '../reducers';
import DevTools from '../containers/DevTools';
import problemsReducer from '../reducers/problems/index';
import commentsReducer from '../reducers/comments/index';


const enhancers = compose(
    applyMiddleware(logger, thunk),
    DevTools.instrument(),
    persistState(getDebugSessionKey())
)

function getDebugSessionKey() {
    const matches = window.location.href.match(/[?&]debug_session=([^&]+)\b/);
    return matches && matches.length > 0 ? matches[1] : null;
}


// export default  configureStore = preloadedState =>  {
//     // const rootReducer = combineReducers({problemsReducer, commentsReducer})
//     // const store = createStore(rootReducer, preloadedState, enhancers);

//     // // Hot reload reducers (requires Webpack or Browserify HMR to be enabled)
//     // if (module.hot) {
//     //     module.hot.accept('../reducers', () =>
//     //         store.replaceReducer(
//     //             require('../reducers') /*.default if you use Babel 6+ */
//     //         )
//     //     );
//     // }
//     // return store;
// }

/* eslint-disable */
const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;

export function configureStore(initialState) {
  const store = createStore(
    rootReducer,
    initialState,
    composeEnhancers(
      applyMiddleware(logger, thunk),
    ),
  );

  return store;
}
