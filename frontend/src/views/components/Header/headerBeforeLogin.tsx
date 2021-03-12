import {FC, PropsWithChildren, useEffect, useState} from "react";
import {
	AppBar,
	Button,
	createStyles, LinearProgress,
	makeStyles, Theme,
	Toolbar,
	Typography
} from "@material-ui/core";
import { Link } from 'react-router-dom'
import Logo from './updated.svg'

const useStyles = makeStyles((theme: Theme) =>
	createStyles({
		root: {
			flexGrow: 1,
		},
		title: {
			flexGrow: 1,
		},
		logo: {
			maxWidth: 160,
			marginRight: theme.spacing(2),
		}
	}),
);

export const HeaderBeforeLogin: FC<PropsWithChildren<any>> = ({children}) => {
	const classes = useStyles();
	const [progress, setProgress] = useState(0);

	useEffect(() => {
		const timer = setInterval(() => {
			setProgress((oldProgress) => {
				if (oldProgress === 100) {
					return 0;
				}
				const diff = Math.random() * 10;
				return Math.min(oldProgress + diff, 100);
			});
		}, 500);

		return () => {
			clearInterval(timer);
		};
	}, [])

	return (
		<AppBar position={'static'}>
			<Toolbar>
				<img src={Logo} alt="logo"
				     className={classes.logo}/>
				<Typography variant="h6" className={classes.title}>
					Infinite Gate
				</Typography>
				<Button><Link style={{textDecoration: 'none', color: 'whitesmoke'}} to={'/dummy'}>Dummy</Link></Button>
				<Button><Link style={{textDecoration: 'none', color: 'whitesmoke'}} to={'/dummy-photo'}>Dummy Photo</Link></Button>

				<Button color="secondary" variant="contained">Sign up</Button>
			</Toolbar>
			<LinearProgress variant={'determinate'} color={'secondary'} value={progress} />
		</AppBar>
	)
}
