# Text Transformation Program

This Go program processes text input, applies various transformations based on specified rules, and writes the modified text to an output file. It includes functions for converting numbers to decimal, capitalizing words, adjusting punctuation placement, and more.

## Usage

### Prerequisites

- Go programming language installed on your system

### Running the Program

1. Clone this repository to your local machine:

   ```bash
   git clone https://learn.zone01kisumu.ke/git/bernaotieno/go-reloaded.git

2. Navigate to the project directory:

    ```bash
    $ cd go-reloaded

3. Compile and run the program:

    ```bash

    $ go run main.go sample.txt result.txt

    Replace sample.txt with your input file containing the text to be processed and result.txt with the desired output file name.

## Sample Input File

The sample.txt file should contain the text you want to process without modification. Ensure the file exists and is accessible to the program.

## Running Tests

The program includes test files to ensure the correctness of various functions. To run tests, use the following command:

    ```bash

    $ go test -v

### Test Files

The program includes test files for unit testing various functions. These test files cover functions such as converting numbers, capitalizing words, handling punctuation, and more. To run tests for specific functions, modify the test files accordingly.
Test Coverage

The provided tests cover all functions except for wordsModifier and main. Ensure to test these functions separately or integrate them into the existing test suite as needed.
