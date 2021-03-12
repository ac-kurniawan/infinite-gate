import {FC, PropsWithChildren, useCallback, useEffect} from "react";
import {Layout} from "../../layouts/default";
import {
	CircularProgress, Container, createStyles, makeStyles,
	Paper,
	Table, TableBody, TableCell,
	TableContainer,
	TableHead, TableRow, Theme, Typography,
} from "@material-ui/core";
import {useDispatch, useSelector} from "react-redux";
import {LastPresenceAction, PresenceViewModel, TRootStore} from "../../stores";

const useStyles = makeStyles((theme: Theme) => createStyles({
	loading: {
		position: 'relative',
		margin: '20px 0px',
		left: 'calc(50% - 20px)',
	}
}))

const PresenceTableData: FC<PropsWithChildren<any>> = ({children, data}) => {
	// @ts-ignore
	return data.map((data: PresenceViewModel, i) => {
		return (
			<TableRow key={`${data.date?.toDateString()}-${i}`}>
				<TableCell>
					{data.date?.toDateString()}
				</TableCell>
				<TableCell>
					{data.checkInDate?.toTimeString()}
				</TableCell>
				<TableCell>
					{data.checkOutDate?.toTimeString()}
				</TableCell>
			</TableRow>
		)
	})
}

const PresenceTable: FC<PropsWithChildren<any>> = ({children}) => {
	return (
		<TableContainer component={Paper}>
			<Table aria-label={'Presence Table'}>
				<TableHead>
					<TableRow>
						<TableCell align={'center'}><strong>Date</strong></TableCell>
						<TableCell align={'center'}><strong>Check In</strong></TableCell>
						<TableCell align={'center'}><strong>Check Out</strong></TableCell>
					</TableRow>
				</TableHead>
				<TableBody>
					{children}
				</TableBody>
			</Table>
		</TableContainer>
	)
}

export const DummyPage: FC<PropsWithChildren<any>> = (props) => {
	const classes = useStyles()
	const dispatch = useDispatch();
	// create getLastPresence dispatcher
	const dispatchLastPresence = useCallback(
		() => dispatch(LastPresenceAction.getLastPresence()), [dispatch]);

	useEffect(() => {
		dispatchLastPresence()
	}, [dispatchLastPresence]);

	// get lastPresences state
	const lastPresences = useSelector((state: TRootStore) => state.lastPresences);

	return (
		<Layout>
			<Typography variant={'h4'} align={'center'} style={{margin: '30px 0px'}}>Dummy
				test</Typography>
			<Container>
				<PresenceTable>
					<PresenceTableData data={lastPresences}/>
				</PresenceTable>
			</Container>
			{
				lastPresences.length === 0 ? <CircularProgress className={classes.loading} color="secondary"/> : ''
			}
		</Layout>
	);
}
