**﻿# linkshortening**

**URL Shortener Service**

**Overview**

This project implements a URL shortener service, allowing users to shorten long URLs into concise, memorable links. The service provides features like link customization, analytics, and security.

**Features**

- Shorten lengthy URLs into shorter links
- Customizable links with user-defined keywords
- Analytics dashboard for tracking click-through rates and referrers
- Security screening for suspicious URLs
- User authentication and authorization

**Technologies Used**

- Programming language: Go
- Framework: Echo
- Dependencies: [List dependencies, e.g., go-get, go-mod]

**Setup and Installation**

1. Clone the repository: git clone [(https://github.com/Eyuvasri27/linkshortening)]
2. Install dependencies: go get -u [go get github.com/labstack/echo/v4
  25 go get github.com/labstack/echo/v4/middleware]
4. Run the service: go run main.go shortner.go store.go

**Usage**

**API Endpoints**

- POST /shorten: Shorten a URL "http://localhost:8080/shorten"
- GET /:shortCode: Redirect to original URL "http://localhost:8080/:shortenurloutput"
- GET /metrics: Get top 3 domain names with most shortened UR "http://localhost:8080/metrics"

  
**API request and Response **
- POST /shorten: Shorten a URL "http://localhost:8080/shorten"

- input as json
- {
  "url": "url"
  }
Response

{
  "shortURL": "d18a7939"
}

- GET /:shortCode: Redirect to original URL "http://localhost:8080/:shortenurloutput"
- GET /metrics: Get top 3 domain names with most shortened UR "http://localhost:8080/metrics"

Response

{
  "docs.digitalocean.com": 1,
  "docs.google.com": 3,
  "mail.google.com": 1
}



