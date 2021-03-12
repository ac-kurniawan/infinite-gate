// Model
export type TPhoto = {
	id?: number;
	title?: string;
	url?: string;
	thumbnailUrl?: string;
}

// Builder
export class PhotoBuilder {
	private id?: number;
	private title?: string;
	private url?: string;
	private thumbnailUrl?: string;

	public setId(id?: number){
		this.id = id;
		return this;
	}
	public setTitle(title?: string){
		this.title = title;
		return this;
	}
	public setUrl(url?: string){
		this.url = url;
		return this;
	}
	public setThumbnailUrl(thumbnailUrl?: string){
		this.thumbnailUrl = thumbnailUrl;
		return this;
	}

	public build(): TPhoto{
		return {
			id: this.id,
			title: this.title,
			url: this.url,
			thumbnailUrl: this.thumbnailUrl
		}
	}

}

// Service interface
export interface IPhotoService {
	getPhotos(): Promise<TPhoto[]>
}

// Service
export class PhotoService implements IPhotoService {
	private PhotoServiceInf: IPhotoService;

	constructor(photoService: IPhotoService) {
		this.PhotoServiceInf = photoService
	}

	getPhotos(): Promise<TPhoto[]> {
		return this.PhotoServiceInf.getPhotos();
	}
}
