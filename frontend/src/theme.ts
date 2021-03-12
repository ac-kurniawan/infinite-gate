import {createMuiTheme} from "@material-ui/core";
import createPalette from "@material-ui/core/styles/createPalette";

export const Theme = createMuiTheme({
	palette: createPalette({
		common: {
			black: "#000",
			white: "#fff"
		},
		background: {
			paper: "rgba(255, 255, 255, 1)",
			default: "#fafafa"
		},
		primary: {
			light: "#77d4f6",
			main: "#56CAF4",
			dark: "#3c8daa",
		},
		secondary: {
			light: "#f6b35a",
			main: "#F4A031",
			dark: "#aa7022",
		},
		error: {
			light: "#e57373",
			main: "#f44336",
			dark: "#d32f2f",
			contrastText: "#fff"
		},
		text: {
			primary: "rgba(0, 0, 0, 0.87)",
			secondary: "rgba(0, 0, 0, 0.54)",
			disabled: "rgba(0, 0, 0, 0.38)",
			hint: "rgba(0, 0, 0, 0.38)"
		}
	})
})
