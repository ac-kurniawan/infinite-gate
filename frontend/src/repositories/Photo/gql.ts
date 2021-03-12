import {
	ApolloClient, gql,
	NormalizedCacheObject
} from "@apollo/client";
import {from} from "rxjs";
import {map} from "rxjs/operators";
import {IPhotoService, PhotoBuilder, TPhoto} from "../../services/Photo";

export class PhotoGql implements IPhotoService {
	private gqlClient: ApolloClient<NormalizedCacheObject>;

	constructor(gql: ApolloClient<NormalizedCacheObject>) {
		this.gqlClient = gql;
	}

	getPhotos(): Promise<TPhoto[]> {
		const res = this.gqlClient.query({
			query: gql`query {
						  photos(options: { paginate: { page: 1, limit: 10 } }) {
						    data {
						      id
						      title
						      url
						      thumbnailUrl
						    }
						  }
						}`
		})

		const transformResponse = from(res).pipe(
			map(responses =>
				responses.data.photos.data.map((photo: { id: number | undefined; title: string | undefined; url: string | undefined; thumbnailUrl: string | undefined; }) =>
					new PhotoBuilder()
						.setId(photo.id)
						.setTitle(photo.title)
						.setUrl(photo.url)
						.setThumbnailUrl(photo.thumbnailUrl)
						.build()
				)
			)
		)
		return transformResponse.toPromise();
	}

}
