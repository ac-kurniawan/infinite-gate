import {
	IPresenceService,
	PresenceBuilder,
	TPresence
} from "../../services/Presence";
import {from} from "rxjs";
import {map, mergeMap} from "rxjs/operators";

export class PresenceAPIImlp implements IPresenceService {
	getPresence(): Promise<TPresence[]> {
		// Create fetch endpoint
		const fetchToEndpoint = fetch(
			'https://618ec5a150e24d0017ce144b.mockapi.io/presence');

		// Fetch data
		const getData = from(fetchToEndpoint).pipe(
			map(response => response.json())
		);

		// Transform to Presence model
		const presence = getData.pipe(
			mergeMap(datas => from(datas).pipe(
				map(data => {
						// @ts-ignore
					return data.map(({checkInTime, checkOutTime, date, id}) => new PresenceBuilder()
							.setId(id)
							.setDate(new Date(date))
							.setCheckInDate(new Date(checkInTime))
							.setCheckOutDate(new Date(checkOutTime))
							.build())
					}
				)
			))
		)
		return presence.toPromise()
	};
}
