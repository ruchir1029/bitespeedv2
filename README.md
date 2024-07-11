# Bitespeed Backend Task: Identity Reconciliation

This project implements a web service for Bitespeed to identify and keep track of a customer's identity across multiple purchases based on email and phone numbers.

## Project Structure


## Prerequisites

- Go 1.16+
- PostgreSQL 13+

## Setup

### PostgreSQL

1. **Install PostgreSQL:**
   Follow the instructions for your operating system from the [PostgreSQL website](https://www.postgresql.org/download/).

2. **Create a Database:**
   Open the PostgreSQL shell or use a GUI tool like pgAdmin and run:
   ```sql
   CREATE DATABASE bitespeed;


CREATE TABLE Contact (
    id SERIAL PRIMARY KEY,
    phoneNumber VARCHAR(20),
    email VARCHAR(100),
    linkedId INT,
    linkPrecedence VARCHAR(10),
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deletedAt TIMESTAMP
);


git clone https://github.com/yourusername/bitespeed.git
cd bitespeed

go mod init bitespeed

go get -u github.com/gin-gonic/gin
go get -u github.com/jinzhu/gorm
go get -u github.com/jinzhu/gorm/dialects/postgres


DB_HOST=localhost
DB_PORT=5432
DB_USER=yourusername
DB_NAME=bitespeed
DB_PASSWORD=yourpassword


go run main.go


curl -X POST http://localhost:8080/identify 
