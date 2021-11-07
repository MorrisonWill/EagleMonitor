import React from "react";
import { StyleSheet, Text, View } from "react-native";
import createAuth0Client from "@auth0/auth0-spa-js";
import LoginButton from "./LoginButton";
import LogoutButton from "./LogoutButton";
import { Auth0Provider } from "@auth0/auth0-react";
// import { useAuth0 } from "./react-auth0-spa";
import { LoggedIn } from "./LoggedIn";
import { useAuth0 } from "@auth0/auth0-react";

export default function App() {
  const { isAuthenticated } = useAuth0();

  const { loading } = useAuth0();

  const auth0 = async () =>
    await createAuth0Client({
      domain: "dev-16j507u0.us.auth0.com",
      client_id: "ISdvSESZ7KHOcU3nhmhn5s5uTZWjTj5F",
      redirect_uri: "localhost:19006/callback", //change when live
      cacheLocation: "memory",
    });

  return (
    <View style={styles.container}>
      <div>
        {!isAuthenticated && (
          <View>
            <Text style={styles.font}> Welcome to EagleEye. </Text>
            <Auth0Provider
              domain="dev-16j507u0.us.auth0.com"
              clientId="ISdvSESZ7KHOcU3nhmhn5s5uTZWjTj5F"
              //redirectUri={window.location.origin} // works with web, breaks ios - need to fix
            >
              <LoginButton style={styles.button}></LoginButton>
            </Auth0Provider>
          </View>
        )}
        {isAuthenticated && <LoggedIn /> && (
          <LogoutButton style={styles.button}></LogoutButton>
        )}
      </div>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#eaeaea",
    alignItems: "center",
    justifyContent: "center",
  },
  font: {
    fontSize: 36,
    fontFamily: "Sans",
  },
  button: {
    marginTop: 16,
    paddingVertical: 8,
    borderWidth: 4,
    borderColor: "#20232a",
    borderRadius: 6,
    backgroundColor: "#61dafb",
    color: "#20232a",
    textAlign: "center",
    fontSize: 30,
    fontWeight: "bold",
  },
});
