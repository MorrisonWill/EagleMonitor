import React, { useEffect, useState } from "react";
import { useAuth0 } from "@auth0/auth0-react";
import LogoutButton from "./LogoutButton";

export const LoggedIn = () => {
  const { getTokenSilently, loading, user, logout, isAuthenticated } =
    useAuth0();

  useEffect(() => {
    const getCatalog = async () => {
      try {
        // Get JWT Token from
        const token = await getTokenSilently();
        const response = await fetch("http://localhost:8081/catalong?", {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

        const responseData = await response.json();

        setCatalog(responseData);
      } catch (error) {
        console.error(error);
      }
    };
    getCatalog();
  }, []);

  if (loading || !user) {
    return <div>Loading...</div>;
  }

  return (
    <div className="container">
      <div className="jumbotron text-center mt-5">
        <h1>EagleEye</h1>
        {isAuthenticated && <LogoutButton></LogoutButton>}
        <p>
          Hi, {user.name}! Below you'll find the course catalog for BC Spring
          Semester 2022.
        </p>
        <div className="row">
          {products.map(function (product) {
            return (
              <div className="col-sm-4">
                <div className="card">
                  <div className="card-header">
                    {product.Name}
                    <span className="float-left">{voted}</span>
                  </div>
                  <div className="card-body">{product.Description}</div>
                  <div className="card-footer">
                    <a
                      onClick={() => vote("Upvoted")}
                      className="btn btn-default float-left"
                    ></a>
                    <a
                      onClick={() => vote("Downvoted")}
                      className="btn btn-default float-right"
                    ></a>
                  </div>
                </div>
              </div>
            );
          })}
        </div>
      </div>
    </div>
  );
};
