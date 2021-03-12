import {LastPresences, PresenceViewModel} from "./presence";
import {combineReducers} from "redux";

export type TRootStore = {
	lastPresences: PresenceViewModel[]
}

export const rootReducer = combineReducers({
	lastPresences: LastPresences
})

