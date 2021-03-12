import {FC, PropsWithChildren} from "react";
import {HeaderBeforeLogin} from "./headerBeforeLogin";

export const Header:FC<PropsWithChildren<any>> = ({children}) => {
	const isLogin = false;
	return (
		<>
			{isLogin ? (<h1>Signed</h1>) : (<HeaderBeforeLogin />)}
		</>
	)
}
