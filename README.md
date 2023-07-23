![Golang gopher image](https://miro.medium.com/v2/resize:fit:256/1*cWqak8OijbTerY420wRgNQ.png)
# **Go API Template**

Quick start basic api template in Go!

## Features

- JWT Auth
- Marketplace template with both user accounts and merchant accounts
- Basic CRUD routes
- PostgreSQL
- Stripe for payments[^1].
- Mailgun for emails and notifications[^2].
- Makefile  

## Setup & Installations

1. Clone the repo
2. Add your credentials for the following vars in a root `.env` file:  
```
DB_USER=                 <dbuser>
DB_PWD=                  <dbpassword>
DB_TCP_HOST=             <host>
DB_PORT=                 <port>
DB_NAME=                 <dbname>
PG_SSL_MODE=             disable
POSTGRESQL_CLOUD_URL=    <gcp cloud url>
POSTGRESQL_LOCAL_URL=    <local url>
STRIPE_KEY=              <stripekey>
STRIPE_SECRET=           <stripe secret>
STRIPE_WEBHOOK_SECRET=   <stripe webhook secret>
DOMAIN=                  <stripe domain>
MAILGUN_API_KEY=         <mailgun key>
MAILGUN_BASE_URL=        <mailgun baseurl>
DOMAIN=                  <mailgun domain>
JWT_SECRET=              <jwt secret>
```
3. Run `go get` to start downloading the necessary modules
4. Run `make migrate-up` to run migrations
5. Run `make run` to build and start the app

[^1]: Work in progress, estimated completion August 2024.
[^2]: Work in progress, estimated completion September 2024.

