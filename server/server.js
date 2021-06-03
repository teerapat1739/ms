// server/server/js
const express = require("express");
const bodyParser = require("body-parser");
const cors = require("cors");
const app = express();
const port = 8000;
const jwt = require("express-jwt");
const jwksRsa = require("jwks-rsa");
const { default: axios } = require("axios");

app.use(bodyParser.json());
app.use(cors());
app.use(express.urlencoded({ extended: true }));
require('dotenv').config()


// Set up Auth0 configuration
const authConfig = {
  domain: "dev-bzic0sop.auth0.com",
  audience: "https://vue-express-api.com",
  secret: "vBFpjeHQiy6pKSMKx_CVboYlb_8V6lw9Y3OSGRhM9ugeSPYpXhfmL6X3yDxJAQqt"
};

// Create middleware to validate the JWT using express-jwt
const checkJwt = jwt({
  // Provide a signing key based on the key identifier in the header and the signing keys provided by your Auth0 JWKS endpoint.
  secret: jwksRsa.expressJwtSecret({
    cache: true,
    rateLimit: true,
    jwksRequestsPerMinute: 5,
    jwksUri: `https://${authConfig.domain}/.well-known/jwks.json`
  }),

  // Validate the audience (Identifier) and the issuer (Domain).
  audience: authConfig.audience,
  issuer: `https://${authConfig.domain}/`,
  algorithms: ["RS256"]
});

// mock data to send to our frontend
let events = [
  {
    id: 1,
    name: "Charity Ball",
    category: "Fundraising",
    description:
      "Spend an elegant night of dinner and dancing with us as we raise money for our new rescue farm.",
    featuredImage: "https://placekitten.com/500/500",
    images: [
      "https://placekitten.com/500/500",
      "https://placekitten.com/500/500",
      "https://placekitten.com/500/500"
    ],
    location: "1234 Fancy Ave",
    date: "12-25-2019",
    time: "11:30"
  },
  {
    id: 2,
    name: "Rescue Center Goods Drive",
    category: "Adoptions",
    description:
      "Come to our donation drive to help us replenish our stock of pet food, toys, bedding, etc. We will have live bands, games, food trucks, and much more.",
    featuredImage: "https://placekitten.com/500/500",
    images: ["https://placekitten.com/500/500"],
    location: "1234 Dog Alley",
    date: "11-21-2019",
    time: "12:00"
  },
  {
    id: 3,
    name: "Charity Ball",
    category: "Fundraising",
    description:
      "Spend an elegant night of dinner and dancing with us as we raise money for our new rescue farm.",
    featuredImage: "https://placekitten.com/500/500",
    images: [
      "https://placekitten.com/500/500",
      "https://placekitten.com/500/500",
      "https://placekitten.com/500/500"
    ],
    location: "1234 Fancy Ave",
    date: "12-25-2019",
    time: "11:30"
  },
  {
    id: 5,
    name: "Rescue Center Goods Drive",
    category: "Adoptions",
    description:
      "Come to our donation drive to help us replenish our stock of pet food, toys, bedding, etc. We will have live bands, games, food trucks, and much more.",
    featuredImage: "https://placekitten.com/500/500",
    images: ["https://placekitten.com/500/500"],
    location: "1234 Dog Alley",
    date: "11-21-2019",
    time: "12:00"
  },
  {
    id: 6,
    name: "Rescue Center Goods Drive",
    category: "Adoptions",
    description:
      "Come to our donation drive to help us replenish our stock of pet food, toys, bedding, etc. We will have live bands, games, food trucks, and much more.",
    featuredImage: "https://placekitten.com/500/500",
    images: ["https://placekitten.com/500/500"],
    location: "1234 Dog Alley",
    date: "11-21-2019",
    time: "12:00"
  }
];
const base_url = {
  party_service: process.env.API_PARTY_SERVICE_URL || "http://localhost:7777",
  // party_service: "https://dbe34a7a6452.ngrok.io",
  payment_service: process.env.API_PAYMENT_SERVICE_URL || "http://localhost:9999"
}
// get all events
app.get("/events", (req, res) => {
  res.send(events);
});

app.get("/events/:id", checkJwt, (req, res) => {
  const id = Number(req.params.id);
  const event = events.find(event => event.id === id);
  res.send(event);
});

app.get("/", (req, res) => {
  res.send(`Hi! Server is listening on port ${port}`);
});

app.get("/public/api-party/*", async (req, res) => {
  try {
    let {data} = await axios.get(`${base_url.party_service}${req.url}`)
    res.send(data);
  } catch (error) {
    res.send(error);
  }
});


app.post("/public/api-party/*", async (req, res) => {
  try {
    let res = await axios({
      method: 'post',
      url: `${base_url.party_service}${req.url}`,
      data: req.body
    });
    res.send(res);
  } catch (error) {
    res.send(error);
  }
});


app.get("/private/api-party/*", checkJwt, async (req, res) => {
  try {
    let {data} = await axios.get(`${base_url.party_service}${req.url}`)
    res.send(data);
  } catch (error) {
    res.send(error);
  }
});

app.post("/private/api-party/*", checkJwt, async (req, res) => {
  try {
    let res = await axios({
      method: 'post',
      url: `${base_url.party_service}${req.url}`,
      data: req.body
    });
    res.send(res);
  } catch (error) {
    res.send(error);
  }
});

app.post("/private/api-payment/*", checkJwt, async (req, res) => {
  console.log(`${base_url.payment_service}${req.url}`)
  try {
    let res = await axios({
      method: 'post',
      url: `${base_url.payment_service}${req.url}`,
      data: req.body
    });
    res.send(res);
  } catch (error) {
    res.send(error);
  }

});


// listen on the port
app.listen(port);

