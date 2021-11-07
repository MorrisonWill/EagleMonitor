// import React from "react";
// import * as AuthSession from "expo-auth-session";

// const auth0 = async () =>
//   await createAuth0Client({
//     domain: "dev-16j507u0.us.auth0.com",
//     client_id: "ISdvSESZ7KHOcU3nhmhn5s5uTZWjTj5F",
//     redirect_uri: "localhost:3000/callback",
//     cacheLocation: "memory",
//   });

// const auth0Domain = "dev-16j507u0.us.auth0.com";
// const auth0ClientId = "ISdvSESZ7KHOcU3nhmhn5s5uTZWjTj5F";
// const auth0RedirectUrl = "localhost:3000";

// export default class Auth0LoginContainer extends React.Component {
//   _loginWithAuth0 = async () => {
//     const redirectUrl = auth0RedirectUrl;
//     let authUrl =
//       `${auth0Domain}/authorize` +
//       toQueryString({
//         client_id: auth0ClientId,
//         response_type: "token",
//         scope: "openid profile email",
//         redirect_uri: redirectUrl,
//       });
//     console.log(`Redirect URL (add this to Auth0): ${redirectUrl}`);
//     console.log(`AuthURL is:  ${authUrl}`);
//     const result = await AuthSession.startAsync({
//       authUrl: authUrl,
//     });

//     if (result.type === "success") {
//       console.log(result);
//       cachedToken = result;
//     }
//   };

//   render() {
//     return <Login onLogin={() => this._loginWithAuth0} />;
//   }
// }

import * as AuthSessionNew from "expo-auth-session";
import jwtDecode from "jwt-decode";
import _ from "lodash";

const {
  auth0: { auth0Domain = "dev-16j507u0.us.auth0.com"},
} 

const toQueryString = () =>
  "?" +
  Object.entries(params)
    .map(
      ([key, value]) =>
        `${encodeURIComponent(key)}=${encodeURIComponent(value)}`
    )
    .join("&");

const logout = async () => {
  return AuthSession.dismiss();
};

const login = async (settings) => {
  // Retrieve the redirect URL, add this to the callback URL list
  // of your Auth0 application.
  const redirectUrl = "http://localhost:19006";

  // Structure the auth parameters and URL
  const params = {
    client_id: "ISdvSESZ7KHOcU3nhmhn5s5uTZWjTj5F",
    redirect_uri: redirectUrl,
    // response_type:
    // id_token will return a JWT token with the profile
    // token will return access_token to use with api calls
    response_type: "token id_token",
    nonce: "nonce", 
  };

  const queryParams = toQueryString(params);
  const authUrl = `https://${auth0Domain}/authorize${queryParams}`;

  // const response = await WebBrowser.openBrowserAsync(authUrl, {showInRecents: true});
  // const response = await WebBrowser.openAuthSessionAsync(authUrl, {showInRecents: true});
  const response = await AuthSessionNew.startAsync({
    authUrl,
    showInRecents: true,
  });

  // const response = await startAuth(authUrl);
  return handleLoginResponse(response);
};

const handleLoginResponse = (response) => {
  if (response.error || response.type !== "success") {
    return;
  }

  const decodedJwtIdToken = jwtDecode(response.params.id_token);
  const fullName = decodedJwtIdToken["placeholder"];
  return {
    ...decodedJwtIdToken,
    name: fullName,
    firstName: _.words(fullName)[0],
    meta: decodedJwtIdToken["placeholder"],
    primaryUserId: decodedJwtIdToken,
  };
};

export default { login, logout };
