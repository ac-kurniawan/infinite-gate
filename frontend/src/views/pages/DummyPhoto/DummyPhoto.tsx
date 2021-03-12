import {FC, PropsWithChildren, useEffect, useState} from "react";
import {Layout} from "../../layouts/default";
import {
	Button, Card, CardActionArea, CardActions,
	CardContent,
	CardMedia, CircularProgress, Container, createStyles,
	makeStyles, Theme,
	Typography
} from "@material-ui/core";
import {PhotoService, TPhoto} from "../../../services/Photo";
import {ApolloClient, InMemoryCache} from "@apollo/client";
import {PhotoGql} from "../../../repositories/Photo";

const useStyles = makeStyles((theme: Theme) => createStyles({
	root: {
		width: 300,
		margin: theme.spacing(2),
	},
	media: {
		height: 150,
	},
	cardContainer: {
		display: 'flex',
		justifyContent: 'space-around',
		flexWrap: 'wrap',
	},
	loading: {
		position: 'relative',
		margin: '20px 0px',
		left: 'calc(50% - 20px)',
	}
}));

// @ts-ignore
export default function MediaCard({title, thumbnailUrl}) {
	const classes = useStyles();

	return (
		<Card className={classes.root}>
			<CardActionArea>
				<CardMedia
					className={classes.media}
					image={thumbnailUrl}
					title={title}
				/>
				<CardContent>
					<Typography gutterBottom variant="h5" component="h2">
						{title}
					</Typography>
				</CardContent>
			</CardActionArea>
			<CardActions>
				<Button size="small" color="primary">
					Share
				</Button>
				<Button size="small" color="primary">
					Learn More
				</Button>
			</CardActions>
		</Card>
	);
}


export const DummyPhoto: FC<PropsWithChildren<any>> = (props) => {
	const classes = useStyles()
	const [photoList, setPhotoList] = useState([] as TPhoto[])

	useEffect(() => {
		const gqlClient = new ApolloClient({
			uri: 'https://graphqlzero.almansi.me/api',
			cache: new InMemoryCache()
		});

		const gqlPhoto = new PhotoGql(gqlClient)
		const photoService = new PhotoService(gqlPhoto)

		photoService.getPhotos().then(res => {
			setPhotoList(res)
		})


	}, [setPhotoList])

	return (
		<Layout>
			<Typography variant={'h4'} align={'center'} style={{margin: '30px 0px'}}>Dummy
				photo test</Typography>
			{photoList.length === 0 ? <CircularProgress className={classes.loading} color="secondary"/> : ''}
			<Container className={classes.cardContainer}>
				{
					photoList.map(photo => <MediaCard key={photo.title} title={photo.title} thumbnailUrl={photo.thumbnailUrl} />)
				}
			</Container>
		</Layout>
	)
}
