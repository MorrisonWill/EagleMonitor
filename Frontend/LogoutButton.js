import React, { Fragment } from "react";
import { useAuth0 } from "@auth0/auth0-react";

const LogoutButton = () => {
  const { isAuthenticated, logout } = useAuth0();

  return (
    <Fragment>
      <div className="container">
        <div className="jumbotron text-center mt-5">
          <button
            className="btn btn-primary btn-lg btn-login btn-block"
            onClick={() => logout()}
          >
            Log Out
          </button>
        </div>
      </div>
    </Fragment>
  );
};

export default LogoutButton;
