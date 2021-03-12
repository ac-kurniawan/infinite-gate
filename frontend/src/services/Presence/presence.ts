// Model / entity
export type TPresence = {
	id?: number;
	date?: Date,
	checkInDate?: Date,
	checkOutDate?: Date,
}

// Builder
export class PresenceBuilder {
	private id?: number;
	private date?: Date;
	private checkInDate?: Date;
	private checkOutDate?: Date;

	public setId(id: number) {
		this.id = id;
		return this
	}

	public setDate(date: Date) {
		this.date = date
		return this
	}

	public setCheckInDate(date: Date) {
		this.checkInDate = date;
		return this;
	}

	public setCheckOutDate(date: Date) {
		this.checkOutDate = date;
		return this;
	}

	public build(): TPresence {
		return {
			id: this.id,
			date: this.date,
			checkInDate: this.checkInDate,
			checkOutDate: this.checkOutDate,
		}
	}
}

// Service interface
export interface IPresenceService {
	getPresence(): Promise<TPresence[]>
}

// Service
export class PresenceService implements IPresenceService {
	private PresenceRepo: IPresenceService;

	constructor(PresenceRepo: IPresenceService) {
		this.PresenceRepo = PresenceRepo
	}

	getPresence(): Promise<TPresence[]> {
		return this.PresenceRepo.getPresence()
	};

}
