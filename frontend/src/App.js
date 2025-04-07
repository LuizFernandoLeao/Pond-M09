import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import UserForm from './components/UserForm';
import UserList from './components/UserList';

function App() {
  return (
    <Router>
      <div>
        <h1>Aplicação WEB</h1>
        <Switch>
          <Route path="/users" component={UserList} />
          <Route path="/user-form" component={UserForm} />
          <Route path="/" exact>
            <h2>Bem vindo a minha Aplicação WEB</h2>
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

export default App;