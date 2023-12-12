# Products Test Summary

## Objective
The objective of this project was to develop a RESTful API capable of managing product data from a provided JSON file. Tasks included implementing various functionalities such as data loading, searching, aggregation, modification, and validation.

## Completion Status
I have successfully completed the tasks outlined by implementing a Go-based RESTful API using the Gin framework. The API efficiently manages product data, offering functionalities for data manipulation and retrieval based on specific criteria such as category and price range.

## Implementation Overview
This RESTful API, built with Go and utilizing the Gin framework, efficiently processes and manages product data sourced from a provided JSON file. The API leverages a map-based data storage system, where product details are stored with unique IDs as keys and corresponding attributes as values. Additionally, an auxiliary map structure has been implemented, indexing category names as keys and arrays of associated IDs as values. These design choices optimize data retrieval, enabling swift search operations and organized access based on categories.

The API seamlessly integrates Swagger for comprehensive API documentation and views, ensuring clarity and accessibility in understanding available endpoints and their functionalities. Its structured architecture adheres to clean coding practices and industry-standard tools, fostering scalability and maintainability for future enhancements.

## Design Decisions
- **Backend Technology Choice:** Utilized Go language with the Gin framework for the backend, ensuring robustness and leveraging efficient web routing and middleware functionalities.

- **Data Storage Strategy:** Employed a map-based storage system where product details are stored with unique IDs as keys and corresponding attributes as values. Additionally, incorporated a secondary map structure to index category names for organized retrieval.

- **API Documentation Framework:** Integrated Swagger for comprehensive API documentation, offering clear views of endpoints and their functionalities, enhancing usability and understanding.

- **Code Structure and Cleanliness:** Emphasized clean coding practices to ensure well-structured and maintainable code, allowing for scalability and ease of future enhancements.

## Running the Code Locally
- **Clone the Repository:** Clone the repository using the following command:
    ```
    git clone https://github.com/FaresAbuIram products_test.git
    ```
- **Install Go (version 1.21.4):** If you don't have Go installed or need version 1.21.4, you can download it from the official Go downloads page: [Go Downloads](https://golang.org/dl/)

- **Download Go Libraries:** Once Go is installed, navigate to the project directory and run the following command to download the required Go libraries:
    ```
    go get ./...
    ```

- **Run the Application:** After downloading the libraries, execute the following command:
    ```
    go run main.go
    ```

- **Access Swagger Documentation:** Once the application is running, open your browser and go to http://localhost:8080/swagger/index.html to access the Swagger API documentation.