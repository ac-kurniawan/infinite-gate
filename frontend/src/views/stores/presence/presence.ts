import {Epic, ofType} from "redux-observable";
import {map, mergeMap, switchMapTo} from "rxjs/operators";
import {PresenceAPIImlp} from "../../../repositories/Presence";
import {from} from "rxjs";
import {PresenceService} from "../../../services/Presence";

// VIEW MODEL
export type PresenceViewModel = {
	date?: Date;
	checkInDate?: Date;
	checkOutDate?: Date;
}

export class PresenceViewModelBuilder {
	private date?: Date;
	private checkInDate?: Date;
	private checkOutDate?: Date;
	public setDate(date?: Date) {
		this.date = date;
		return this;
	}
	public setCheckInDate(date?: Date) {
		this.checkInDate = date
		return this;
	}
	public setCheckOutDate(date?: Date) {
		this.checkOutDate = date
		return this;
	}
	public build(): PresenceViewModel {
		return {
			date: this.date,
			checkInDate: this.checkInDate,
			checkOutDate: this.checkOutDate
		};
	}
}

// ACTION MAP
export const LastPresenceAction = {
	getLastPresence: () => ({
		type: 'GET_LAST_PRESENCE'
	}),
	setLastPresences: (payload: PresenceViewModel[]) => ({
		type: `SET_LAST_PRESENCES`,
		payload
	})
}

const defaultState: PresenceViewModel[] = []

// REDUCER
export const LastPresences = (state = defaultState,
	action: { type: string, payload: PresenceViewModel[] }) => {
	switch (action.type) {
		case 'GET_LAST_PRESENCE':
			return state
		case 'SET_LAST_PRESENCES':
			return action.payload
		default:
			return state
	}
}

// MIDDLEWARE (aka EPIC in REDUX-OBSERVABLE)
export const lastPresenceEpic: Epic = (action$) => {
	const presenceFetchApi = new PresenceAPIImlp()

	const presenceRepo = new PresenceService(presenceFetchApi)

	return action$.pipe(
		ofType('GET_LAST_PRESENCE'),
		switchMapTo(from(presenceRepo.getPresence())),
		map(presences => LastPresenceAction.setLastPresences(presences.map(presence =>
			new PresenceViewModelBuilder()
				.setDate(presence.date)
				.setCheckInDate(presence.checkInDate)
				.setCheckOutDate(presence.checkOutDate)
				.build()
		)))
	)
}
