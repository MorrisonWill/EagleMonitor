import React from "react";
import { useAuth0 } from "../react-auth0-spa";

const Home = () => {
  const { isAuthenticated, loginWithRedirect } = useAuth0();
  return (
    <div>
      <h1>Placeholder</h1>
      {!isAuthenticated && (
        <button onClick={() => loginWithRedirect()}>Sign in</button>
      )}
    </div>
  );
};

export default Home;
