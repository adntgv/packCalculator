# Pack Sizes Calculator

The Pack Sizes Calculator is a simple web application built in Go (Golang) that allows users to calculate the required number of packs for a given number of items. It also provides the ability to set and view available pack sizes through a user-friendly web interface.

## Getting Started

### Prerequisites

- Go (Golang) installed on your machine
- Basic understanding of Go programming

### Installation and Usage

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/yourusername/pack-sizes-calculator.git
   cd pack-sizes-calculator
   ```

2. Run the application:

   ```bash
   go run main.go
   ```

3. Open your web browser and navigate to http://localhost:8080.

## Features

- Calculate the required number of packs for a given number of items.
- View and set available pack sizes.
- Interactive web interface for user interaction.

## API Endpoints

- `GET /calculate`: Calculate required packs for a given number of items.
- `GET /sizes`: View available pack sizes.
- `POST /sizes/set`: Set new pack sizes.

## How to Use

1. Access the application through your web browser at http://localhost:8080.
2. Use the provided UI to calculate required packs, view available pack sizes, and set new pack sizes.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please feel free to submit a pull request.
