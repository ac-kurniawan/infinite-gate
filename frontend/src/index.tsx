import React from 'react';
import ReactDOM from 'react-dom';
import {Theme} from "./theme";
import {CssBaseline, MuiThemeProvider} from "@material-ui/core";
import {Router} from "./views/routes";
import {Provider} from "react-redux";
import {lastPresenceEpic, rootReducer} from "./views/stores";
import {combineEpics, createEpicMiddleware} from "redux-observable";
import {applyMiddleware, createStore} from "redux";
import {composeWithDevTools} from "redux-devtools-extension";

const _epicMiddleware = createEpicMiddleware();

const _rootEpic = combineEpics(lastPresenceEpic)

const _store = createStore(
	rootReducer,
	composeWithDevTools(applyMiddleware(_epicMiddleware))
)

_epicMiddleware.run(_rootEpic)

ReactDOM.render(
	<React.StrictMode>
		<Provider store={_store}>
			<MuiThemeProvider theme={Theme}>
				<CssBaseline />
				<Router />
			</MuiThemeProvider>
		</Provider>
	</React.StrictMode>,
	document.getElementById('root')
);

