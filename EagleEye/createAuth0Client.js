// import createAuth0Client from "@auth0/auth0-spa-js";

// export default function LoginPlox() {
//   const auth0 = async () =>
//     await createAuth0Client({
//       domain: "dev-16j507u0.us.auth0.com",
//       client_id: "ISdvSESZ7KHOcU3nhmhn5s5uTZWjTj5F",
//       redirect_uri: "localhost:3000/callback",
//       cacheLocation: "memory",
//     });

//   //with async/await

//   // Login
//   //redirect to the Universal Login Page
//   document.getElementById("login").addEventListener("click", async () => {
//     await auth0.loginWithRedirect();
//   });

//   // //in your callback route (<MY_CALLBACK_URL>)
//   window.addEventListener("load", async () => {
//     const redirectResult = await auth0.handleRedirectCallback();
//     //logged in. you can get the user profile like this:
//     const user = await auth0.getUser();
//     console.log(user);
//   });
// }

// // // Call an API
// // //with async/await
// // document.getElementById("call-api").addEventListener("click", async () => {
// //   const accessToken = await auth0.getTokenSilently();
// //   const result = await fetch("https://myapi.com", {
// //     method: "GET",
// //     headers: {
// //       Authorization: `Bearer ${accessToken}`,
// //     },
// //   });
// //   const data = await result.json();
// //   console.log(data);
// // });
