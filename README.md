# CRUD Operations in Go

This is a project that demonstrates CRUD (Create, Read, Update, Delete) operations using the Go programming language.

## Installation

To use this project, you need to follow these steps:

1. Clone the repository: `git clone https://github.com/TiagoNora/GoCRUD.git`
2. Install the dependencies: `go mod download`
3. Build the application: `go build`
4. Run the application: `./main`

> :warning: After installing all the dependencies you can just use the command `make` to generate the API documentation using Swag and to run the application

## Usage
Once the project is running, you can interact with the CRUD operations using the following endpoints:

##### Products Endpoints

- GET /api/v1/products: Retrieve all products 
- GET /api/v1/product?id={id}: Retrieve an product by ID 
- POST /api/v1/secured/product: Create a new product 
- PUT /api/v1/secured/product?id={id}: Update an existing product by ID 
- DELETE /api/v1/secured/product?id={id}: Delete an product by ID

##### Users Endpoints
- GET /api/v1/secured/users: Retrieve all users
- GET /api/v1/secured/user?id={id}: Retrieve an user by ID 
- POST /api/v1/user: Create a new user
- POST /api/v1/user/token: Generate a token
- PUT /api/v1/secured/user?id={id}: Update an existing user by ID 
- DELETE /api/v1/secured/user?id={id}: Delete an user by ID

Replace {id} with the actual ID of the product/user you want to retrieve, update, or delete.

> :warning: To interact with some funcionalities you have to authenticate.

## Contributing
Contributions to this project are welcome. To contribute, follow these steps:

- Fork the repository.
- Create a new branch.
- Make your changes and commit them.
- Push your changes to your forked repository.
- Submit a pull request.

## License
This project is licensed under the MIT License.