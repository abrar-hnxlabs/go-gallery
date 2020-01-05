import React from "react";
import {
  BrowserRouter as Router,
  Switch,
  Route
} from "react-router-dom";

import MainPage from './page/main'

export default function App() {
  return (
    <Router>
        <Switch>
          <Route path="/ui" component={MainPage} />
        </Switch>
    </Router>
  );
}
