````markdown
# Hugging Face Inference API Transfomers Library

This project demonstrates how to create a RESTful API for performing inference using Hugging Face's models. The API allows you to send text input and get inference results based on the loaded model.

## Prerequisites

Before running the project, make sure you have the following installed:

- Go programming language
- Gin framework
- Hugging Face API key

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/hugging-face-inference-api.git
   cd hugging-face-inference-api
   ```
````

2. Create a `.env` file in the root directory and add your Hugging Face API key:

   ```env
   HUGGING_FACE_API_KEY=YOUR_HUGGING_FACE_API_KEY
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

## Usage

1. Start the server:

   ```bash
   go run main.go
   ```

2. Send a POST request to perform inference:

   ```bash
   curl -X POST http://localhost:8080/inference -H "Content-Type: application/json" -d '{"inputs":"Good morning [MASK]"}'
   ```

   Replace `YOUR_HUGGING_FACE_API_KEY` with your actual API key.

## API Endpoints

- POST `/inference`: Perform inference with text input. Example request and response are available in the [Postman documentation](https://documenter.getpostman.com/view/22804089/2sA3XMjigB).

## Contributing

Feel free to contribute to this project by submitting issues or pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

```

Saya telah menambahkan link ke dokumentasi Postman di bagian "API Endpoints" untuk memudahkan pengguna untuk melihat contoh permintaan dan respons menggunakan Postman. Jangan lupa untuk mengganti `YOUR_HUGGING_FACE_API_KEY` dan `yourusername` dengan nilai yang sesuai dengan proyek Anda.
```
