# React Typescript using Material UI

## Dependencies

1. React
2. React-router-dom
3. Redux
4. React-redux
5. Redux-observable
6. Material-ui/core
7. Material-ui/icons

## Code Architecture
We try to implement hexagonal architecture by splitting Service (model/entity), repository and views.
In other hand we also try to combine with MVVN pattern, so in this adaptation we create rules:

**Services -> Model and ModelView** = In this part we split every model each domain and create business logic for every domain. We use agnostic rules in this part, so we are not consider frontend framework. 

**Views -> View** = We use react library to provide the UI

**Repositories** = This part to communicate with external resources

### Code structure

* **/public**
* **/src**
  + **/repositories**
  + **/services**
  + **/views**
    + **/components** # Store reusable component
    + **/layouts** # Store reusable layout (single page layout, dashboard layout, error page layout, etc)
    + **/pages** # Base on routes
    + **/routes** # Store routes map
    + **/stores** # Global state (action, reducer and middleware)

## Available Scripts

In the project directory, you can run:

### `yarn start`

Runs the app in the development mode.\
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.\
You will also see any lint errors in the console.

### `yarn test`

Launches the test runner in the interactive watch mode.\
See the section about [running tests](https://facebook.github.io/create-react-app/docs/running-tests) for more information.

### `yarn build`

Builds the app for production to the `build` folder.\
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.\
Your app is ready to be deployed!

See the section about [deployment](https://facebook.github.io/create-react-app/docs/deployment) for more information.

### `yarn eject`

**Note: this is a one-way operation. Once you `eject`, you can’t go back!**

If you aren’t satisfied with the build tool and configuration choices, you can `eject` at any time. This command will remove the single build dependency from your project.

Instead, it will copy all the configuration files and the transitive dependencies (webpack, Babel, ESLint, etc) right into your project so you have full control over them. All of the commands except `eject` will still work, but they will point to the copied scripts so you can tweak them. At this point you’re on your own.

You don’t have to ever use `eject`. The curated feature set is suitable for small and middle deployments, and you shouldn’t feel obligated to use this feature. However we understand that this tool wouldn’t be useful if you couldn’t customize it when you are ready for it.

## Learn More

You can learn more in the [Create React App documentation](https://facebook.github.io/create-react-app/docs/getting-started).

To learn React, check out the [React documentation](https://reactjs.org/).

