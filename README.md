
# Psoriasis Image Submission and Clinical Feedback Web Application

## Introduction
Psoriasis Image Submission and Clinical Feedback Web Application is a platform designed to facilitate the submission of images of psoriasis-affected skin by patients and enable authorized clinicians to provide feedback. This README provides an overview of the project and its features.

## Features

1.  User Registration:
	-   Users can register using their email address and password to access the website's features.

2.  Image Submission:
    -   Authenticated users can upload pictures of their psoriasis-affected skin, which are securely stored on the server-side.

3.  Image Management:
    - Authenticated users can view, update, and delete their uploaded images.
    - Users can filter images based on date and body positions.

4.  Clinical Feedback:
    - Users can authorize clinicians to view their uploaded images.
    - Authorized clinicians can access the images and provide feedback to the patients.
    - Multiple clinicians can be authorized for the same patient's images.
    - Clinicians can leave textual feedback for the patient.

5.  Image Security:
    - All images are encrypted using a secure mechanism to ensure privacy.
    - Only the patient and authorized clinicians can view the images by logging into the web application.
    -   The web application employs appropriate security measures, such as SSL/TLS encryption, secure password storage, and two-factor authentication.

6.  Future Enhancements:  
	- Add  support  for  video  submissions  to  allow  users  to  capture  the  
progression of their psoriasis over time.  
	- Implement a feature to allow users to share their images with others for  
research or educational purposes, with their consent  
	- Be able to relate a new image to a previous one in order to follow the  
clinical evolution of a given patient (timeline)

Every feature from 1. to 5. are done, with the exception of two-factor authentication.

## Technologies Used
- React.js
- Golang with Gin
- MariaDB
- Redis
- Docker

## Other Relevant Repositories
REST API: https://github.com/JFSL98/pawb_backend_2023/

## Requirements
[Go](https://go.dev/doc/install)

[Node.js](https://nodejs.org/en/download)

[Docker](https://www.docker.com/get-started/)

## Installation

### Backend Setup
To set up the backend of the Psoriasis Image Submission and Clinical Feedback Web Application, follow these steps:

1.  Clone the repository to your local machine:
```bash
git clone https://github.com/JFSL98/pawb_backend_2023
cd pawb_backend_2023
``` 
    
2.  Install the required dependencies:
```bash
go mod download
``` 
    
3. Create a `.env` file in the root of the `pawb_backend_2023` directory and provide the necessary environment variables

4. Install Air to get Live Reload:
```
go install github.com/cosmtrek/air@latest
```
5. Start the development server:
```bash
air
``` 
    
7. Access the API through a client (frontend, Postman, etc) at `http://localhost:3000`.

### Docker Compose Setup
To set up the Psoriasis Image Submission and Clinical Feedback Web Application using Docker Compose, the following docker-compose.yml file is required in the same folder as the backend and the frontend:

```dockerfile
version: "3.9"
services:
	app:
		container_name: app_backend
		build: ./pawb_backend_2023
		ports:
			- 3000:3000
		restart: on-failure
		volumes:
			- ./pawb_backend_2023:/app
		depends_on:
			- database
			- redis
	database:
		container_name: mariadb
		image: mariadb:10.11.2
		environment:
		MYSQL_ROOT_PASSWORD: ${DB_PASSWORD}
		MYSQL_DATABASE: ${DB_DATABASE}
		MYSQL_USER: ${DB_USERNAME}
		MYSQL_PASSWORD: ${DB_PASSWORD}
		restart: on-failure
		volumes:
			- ./mariadb:/var/lib/mysql
		ports:
			- 3306:3306
	react:
		container_name: app_frontend
		build: ./pawb_frontend_2023
		restart: on-failure
		volumes:
			- ./pawb_frontend_2023/src:/app/src
		depends_on:
			- app
		ports:
			- 3001:3000
	redis:
		container_name: redis
		image: redis:7.0.11
		restart: on-failure
		ports:
			- 6379:6379
```

Then follow these steps:
1. Clone the repositories:
```bash
git clone https://github.com/JFSL98/pawb_frontend_2023
git clone https://github.com/JFSL98/pawb_backend_2023
``` 

2. Setup the `.env`
3. Build and run the application using Docker Compose:
```bash
docker-compose up --build
```
4. Access the web application in your browser at `http://localhost:3001`.

Make sure you have Docker and Docker Compose installed and properly configured on your machine before running the above steps.

## Deployment
The services for the frontend, backend and Redis are being hosted on [Render](https://render.com/), while the database is on [PlanetScale](https://planetscale.com/).

Here is the URL for the web app in production:
https://pawb-react-app-2023.onrender.com/
