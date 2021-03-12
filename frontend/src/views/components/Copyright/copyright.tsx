import {FC} from "react";
import {createStyles, Link, makeStyles, Theme, Typography} from "@material-ui/core";

const useStyles = makeStyles((theme: Theme) => createStyles({
	copyright: {
		position: 'relative',
		bottom: 0,
		margin: '20px 0px',
		left: '50%',
		transform: 'translateX(-50%)'
	}
}))

export const Copyright: FC = () => {
	const classes = useStyles();

	return (
		<Typography className={classes.copyright} variant={'body2'} color={"textSecondary"} align={'center'}>
			{'Copyright © '}
			<Link color="inherit" href="https://newgate.prodreamer.my.id/">
				Infinite Gate
			</Link>{' '}
			{new Date().getFullYear()}
			{'.'}
		</Typography>
	)
}
