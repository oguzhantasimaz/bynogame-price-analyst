# Price Analyst & Comparison Tool

Welcome to the Price Analyst & Comparison Tool! This application allows you to analyze and compare prices from various sources, including Bynogame.com and the Steam market.

## Table of Contents

- [Overview](#overview)
- [Technologies Used](#technologies-used)
- [Deployment](#deployment)
- [Architecture](#architecture)
- [Frontend](#frontend)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Overview

This tool provides price analysis and comparison features for products available on Bynogame.com and the Steam market. It includes a backend system for data processing and storage, as well as a frontend interface for users to interact with the analyzed data.

The frontend GitHub repository can be found [here](https://github.com/oguzhantasimaz/bynogame-price-analyst-fe).

## Technologies Used

- **Golang**: Used for building the main backend API, as well as the Producer and Consumer components.
- **Kafka**: Messaging system for data communication.
- **MongoDB**: Used to store and retrieve data.
- **React.js**: Frontend framework for the user interface.

## Deployment

- **Backend**:
    - Core API running on  [Fly.io](https://fly.io/).
    - Producer and Consumer components are also hosted on [Fly.io](https://fly.io/).
    - Kafka is hosted serverlessly on [Upstash](https://upstash.com/).
    - MongoDB is hosted on [MongoDB Atlas](https://www.mongodb.com/cloud/atlas).
  

- **Frontend**: The frontend is deployed on Vercel and can be accessed at [https://bynogamepriceanalyst.vercel.app/](https://bynogamepriceanalyst.vercel.app/).

## Architecture

The architecture of the system is as follows:
- Producer gathers data from Bynogame.com.
- Consumer consumes data from Kafka and processes it, also gathers data from the Steam market.
- Kafka is used for asynchronous data transfer between components.
- A core API handles data writing to MongoDB and reading from it.

## Frontend

The frontend provides an intuitive interface for users to:
- View price analysis results.
- Compare prices from different sources.
- See the latest price
- Percentage between the two sources.

## Usage

1. Access the [frontend](https://bynogamepriceanalyst.vercel.app/) to start using the tool.
2. Explore products and compare prices from Bynogame.com and the Steam market.
3. Utilize the search and filter options to refine your results.
4. Stay up-to-date with the latest price trends and historical data.

## TODO

- [ ] Add price history graph.
- [ ] Add search and filter options.
- [ ] Add pagination.
- [ ] And more...

## Contributing

Feel free to open issues or submit pull requests to this GitHub repository.

## License

This project is licensed under the [MIT License](LICENSE).
