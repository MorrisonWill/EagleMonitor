import React, { Fragment } from "react";
import { useAuth0 } from "@auth0/auth0-react";

const LoginButton = () => {
  const { loginWithRedirect } = useAuth0();

  return (
    <Fragment>
      <div className="container">
        <div className="jumbotron text-center mt-5">
          <button
            className="btn btn-primary btn-lg btn-login btn-block"
            onClick={() => loginWithRedirect()}
          >
            Log In
          </button>
        </div>
      </div>
    </Fragment>
  );
};

export default LoginButton;

// const styles = StyleSheet.create({
//   container: {
//     flex: 1,
//     backgroundColor: "#eaeaea",
//     alignItems: "center",
//     justifyContent: "center",
//   },
//   button: {
//     marginTop: 16,
//     paddingVertical: 8,
//     borderWidth: 4,
//     borderColor: "#20232a",
//     borderRadius: 6,
//     backgroundColor: "#61dafb",
//     color: "#20232a",
//     textAlign: "center",
//     fontSize: 30,
//     fontWeight: "bold",
//   },
// });
