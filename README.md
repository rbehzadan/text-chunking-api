# Text Chunking REST API

This project provides a REST API for chunking a given text into pieces of a specified size using Go. It uses the `prose` library for natural language processing and `gorilla/mux` for routing.

## Features

- Split text into chunks based on the specified maximum number of tokens.
- Does not split text mid-sentence.
- Log each request with method, URI, remote address, and duration.
- Error handling for invalid input and processing errors.

## Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/rbehzadan/text-chunking-api.git
   cd text-chunking-api
   ```

2. **Install dependencies:**
   Ensure you have Go installed. If not, download and install it from [golang.org](https://golang.org/dl/).

   ```sh
   go mod tidy
   ```

## Usage

1. **Run the server:**
   ```sh
   go run main.go
   ```

2. **Make requests to the API:**
   The server runs on port 8080. You can use tools like `curl` or Postman to interact with the API.

   **Example Request:**
   ```sh
   curl -X POST http://localhost:8080/chunk -H "Content-Type: application/json" -d '{"text": "Your text here.", "max_tokens": 50}'
   ```

   **Example Response:**
   ```json
   {
     "chunks": [
       "Your chunked text here."
     ]
   }
   ```

## API Endpoints

### `POST /chunk`

- **Request Body:**
  ```json
  {
    "text": "Your text here.",
    "max_tokens": 50
  }
  ```

- **Response:**
  ```json
  {
    "chunks": ["Chunk 1", "Chunk 2", ...]
  }
  ```

- **Error Responses:**
  - `400 Bad Request` for invalid request payload or non-positive `max_tokens`.
  - `500 Internal Server Error` for errors during text processing.

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [prose](https://github.com/jdkato/prose) for NLP functionalities.
- [gorilla/mux](https://github.com/gorilla/mux) for routing.

