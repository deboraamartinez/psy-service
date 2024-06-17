# Psychologist Service 

## Overview

We provide cutting-edge technology solutions for psychologists aiming to create smarter and more modern clinics or practices. Our system enhances the efficiency and effectiveness of psychological services, ensuring a seamless experience for both practitioners and clients.

## Technologies Used

- [Gin Gonic](https://github.com/gin-gonic/gin): A high-performance HTTP web framework written in Go.
- [MongoDB Go Driver](https://go.mongodb.org/mongo-driver/mongo): The official MongoDB driver for Go, enabling efficient database operations.
- [JWT (JSON Web Tokens)](https://github.com/golang-jwt/jwt/v4): A robust package for generating and verifying JWTs, used for securing API endpoints.
- [Godotenv](https://github.com/joho/godotenv): A convenient package for loading environment variables from a `.env` file into the application.

### Prerequisites

- Go (version 1.16+)
- MongoDB (running instance or MongoDB Atlas)

### Installation

1. **Clone the repository:**
   ```sh
  git clone https://github.com/deboraamartinez/psy-service.git
   cd psy-service


2. **Install dependencies:**
   ```sh
   go mod download


3. **Set environment variables:**
   ```sh
   MONGODB_URI=mongodb://localhost:27017
   JWT_SECRET=your_secret_key

4. **Run the application:**
   ```sh
   go run cmd/main.go  


### Nexts Steps

- Add Google Calendar
- Add login with Google SSO
- Add login with code sent by email