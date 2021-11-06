import React, { useState } from "react";
import { useAuth0 } from "../react-auth0-spa";

const LoggedIn = () => {
  const { getTokenSilently, loading, user, logout, isAuthenticated } = useAuth0();

  if (loading || !user) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      {isAuthenticated && <button onClick={() => logout()}>Sign out</button>}
      <h1>Placeholder</h1>
      <p>
        Hi, {user.name}.
      </p>
    </div>
  );
}
export default LoggedIn;
