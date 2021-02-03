import React, {Component} from 'react';
import "./App.css";
import { ApolloProvider, useMutation } from "@apollo/react-hooks";
import { client } from "./apollo-client";
import { Route } from "react-router";
import { Link } from "react-router-dom";
import {
    GetAllRestaurants,
    GetRestaurantsByID
} from "./gqlComponents/restaurant/RestaurantContainer";
import {
    GetAllCategories,
    GetCategoryByID
} from "./gqlComponents/category/CategoryContainer";
import { useState, useEffect } from "react";

const CategoryPage = ({match}) => {
    const {
        params: { categoryId },
    } = match;
    const [isLoading, setIsLoading] = useState(true);
    const [data, setData] = useState();

    useEffect(() => {
    },[categoryId]);
    return (
        <div>
            <ApolloProvider client={client}>
            <div class="splash-container">
              <div class="splash">
                  <p class="splash-in">
                    <a><b>
                     <GetCategoryByID filter={categoryId} />
                    </b></a>
                  </p>
              </div>
          </div>
            </ApolloProvider>
        </div>
    )
}

const RestaurantPage = ({match}) => {
  const {
    params: { restaurantId },
  } = match;
  const [isLoading, setIsLoading] = useState(true);
  const [data, setData] = useState();

  useEffect(() => {
  },[restaurantId]);
  return (
    <div>
      <ApolloProvider client={client}>
        <div class="splash-container">
              <div class="splash">
                  <p class="splash-in">
                    <b>
                      <GetRestaurantsByID filter={restaurantId} />
                   </b>
                </p>
            </div>
        </div>
      </ApolloProvider>
    </div>
  )
}

function Categories() {
    return (
        <div>
            <ApolloProvider client={client}>
            <div class="splash-container">
              <div class="splash">
                <h1 class="splash-head">RESTAURANTES</h1>
                  <p class="splash-subhead">
                    <a><b>
                     <GetAllCategories />
                    </b></a>
                   </p>
                  </div>
               </div>
            </ApolloProvider>
        </div>
    );
}

function Restaurants() {
  return (
    <div>
      <ApolloProvider client={client}>
        <div class="splash-container">
              <div class="splash">
                <h1 class="splash-head">RESTAURANTES</h1>
                  <p class="splash-subhead">
                    <b>
                    <GetAllRestaurants />
                  </b>
                </p>
              </div>
        </div>
      </ApolloProvider>
    </div>
  );
}

function HomePage() {
  return (
    <div>
      <ApolloProvider client={client}>
        <div class="splash-container">
            <div class="splash">
                <h1 class="splash-head">easyFood</h1>
                  <p class="splash-subhead">
                    <a><b>
                      <Link to="/restaurant">RESTAURANTES<br></br></Link>
                      <Link to="/category">CATEGORIAS<br></br></Link>
                      <Link to="/dish">PRATOS<br></br></Link>
                    </b></a>
                </p>
            </div>
        </div>
      </ApolloProvider>
    </div>
  );
}

class App extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    let button;
    button = (
      <div>
        <Route exact path="/restaurant" component={Restaurants}>
          <Restaurants />
        </Route>
        <Route path="/restaurant/:restaurantId" component={RestaurantPage} />

        <Route exact path="/category" component={Categories}>
          <Categories />
        </Route>
        <Route path="/category/:categoryId" component={CategoryPage} />

        <Route exact path="/" component={HomePage}>
          <HomePage />
        </Route>

      </div>
    );

    return (
      <div>
        {button}
      </div>
    );
  }
}

export default App;