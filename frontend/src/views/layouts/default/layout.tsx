import {FC, PropsWithChildren} from "react";
import {Header} from "../../components/Header";
import {Copyright} from "../../components/Copyright";

export const Layout: FC<PropsWithChildren<any>> = ({children}) => {
	return (
		<>
			<Header />
			{children}
			<Copyright />
		</>
	)
}
