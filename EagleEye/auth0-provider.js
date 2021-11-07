import React from "react";
import ReactDOM from "react-dom";
import App from "./App";
import { Auth0Provider } from "@auth0/auth0-react";

export default ReactDOM.render(
  <Auth0Provider
    domain="dev-16j507u0.us.auth0.com"
    clientId="ISdvSESZ7KHOcU3nhmhn5s5uTZWjTj5F"
    redirectUri={window.location.origin}
  >
    <App />
  </Auth0Provider>,
  document.getElementById("root")
);
