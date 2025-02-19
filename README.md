
# 📌Basic Banking Transaction API

## 📖 Overview
basic banking transaction to fullfill the PT Ihsan Solusi Informatika for backend engineer test



## 🚀 Features
- ✅ Create Account
- ✅ Top-Up Balance
- ✅ Withdraw Balance
- ✅ Balance Checking


## 🛠️ Tech Stack  
This project is built using the following technologies:  
- **Go (Golang)**  
- **Fiber**
- **GORM**
- **PostgreSQL**
- **Cobra**
## 📌 API Endpoints  

| Method | Endpoint         | Description            |
|--------|----------------|------------------------|
| **POST**    | `/daftar`    | Create user account     |
| **POST**   | `/tabung`     | Add funds to balance   |
| **POST**   | `/tarik`   | Withdraw balance      |
| **GET**    | `/saldo/:no_rekening` | Get user balance |

## 📦 Installation

1. Clone the repository:  
   ```sh
   git clone https://github.com/EputraP/Test_IhsanSolusi.git

   go mod tidy
   ```
2. Copy and configure the environment file:
- This project includes a .env.example file. You need to copy and rename it to .env:
   ```sh
   cp .env.example .env
   ```
- Then, open .env and fill in the required environment variables:
   ```sh
   DB_HOST="localhost"
   DB_PORT="5435"
   DB_USER="postgres"
   DB_PASS="postgres"
   DB_NAME="postgres"
   TIMEZONE="Asia/Jakarta"
   ```
3. Start dependencies using Docker Compose:  
   ```sh
   git clone https://github.com/EputraP/Test_IhsanSolusi.git

   go mod tidy
   ```
4. Run the application: 
   ```sh
   go run main.go -H localhost -P 8080
   ```
This application supports command-line arguments for customizing the server settings.  
### 🔹 Available Options  

| Flag  | Description          | Default Value |
|-------|----------------------|--------------|
| `-H`  | Set the server host  | `localhost`  |
| `-P`  | Set the server port  | `8080`       |

    
## 📜 License  
This project is licensed under the **[MIT](https://choosealicense.com/licenses/mit/)**, which allows commercial and personal use, modification, and distribution.  