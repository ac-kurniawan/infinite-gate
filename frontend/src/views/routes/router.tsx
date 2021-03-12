import {FC, PropsWithChildren} from "react";
import {BrowserRouter, Switch, Route} from "react-router-dom";
import {DummyPage} from "../pages/Dummy";
import {Typography} from "@material-ui/core";
import {DummyPhoto} from "../pages/DummyPhoto";

export const Router: FC<PropsWithChildren<any>> = () => {
	return (
		<BrowserRouter>
			<Switch>
				<Route path={'/dummy'} exact={true}>
					<DummyPage />
				</Route>
				<Route path={'/dummy-photo'} exact={true}>
					<DummyPhoto />
				</Route>
				<Route path={'/*'}>
					<Typography variant={'h3'}>
						404 (Not Found)
					</Typography>
				</Route>
			</Switch>
		</BrowserRouter>
	)
}
