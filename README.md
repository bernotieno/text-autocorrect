## Text Modification Tool in Go

### Introduction
This project aims to develop a command-line tool written in Go for modifying text files according to specified rules. The tool reads an input text file, applies various modifications, and writes the modified text to an output file. It adheres to good coding practices and includes unit tests for ensuring correctness.

### Features
- **Hexadecimal and Binary Conversion**: Replace hexadecimal and binary numbers with their decimal equivalents.
- **Case Modification**: Convert words to uppercase, lowercase, or capitalized as specified, including the ability to modify a specific number of words.
- **Punctuation Formatting**: Adjust spacing around punctuation marks and handle special cases like ellipses and exclamation/question marks.
- **Apostrophe Handling**: Ensure proper placement of apostrophes around words.
- **Article Correction**: Replace "a" with "an" based on the following word.

### Implementation
- **File Handling**: Utilizes Go's standard packages for reading from and writing to text files.
- **Text Parsing**: Tokenizes the input text into words and punctuation marks for easier manipulation.
- **Rule-based Modification**: Implements separate functions for each modification rule to ensure modularity and maintainability.
- **Unit Testing**: Includes test files for validating the correctness of the modification functions.

### Usage
The tool is executed from the command line using the following syntax:
<pre>$ go run . "input_file" "output_file"</pre>


### How to Run
1. Clone the repository.
<pre>$ git clone https://learn.zone01kisumu.ke/git/danyonyi/go-reloaded.git</pre>
2. Navigate to the project directory.
<pre>$ cd go-reloaded</pre>
3. Run the tool with appropriate input and output file paths.
<pre>$ go run . sample.txt result.txt</pre>

### Examples
<pre>
$ cat sample.txt
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.
$ go run . sample.txt result.txt
$ cat result.txt
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
$ cat sample.txt
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
$ go run . sample.txt result.txt
$ cat result.txt
Simply add 66 and 2 and you will see the result is 68.
$ cat sample.txt
There is no greater agony than bearing a untold story inside you.
$ go run . sample.txt result.txt
$ cat result.txt
There is no greater agony than bearing an untold story inside you.
$ cat sample.txt
Punctuation tests are ... kinda boring ,don't you think !?
$ go run . sample.txt result.txt
$ cat result.txt
Punctuation tests are... kinda boring, don't you think!?
</pre>

### Dependencies
- Standard Go packages are used for file handling and text manipulation.

### Future Improvements
- Support for additional modification rules.
- Enhanced error handling and logging.
- Improved performance optimizations for large text files.

### Contributing
Contributions are welcome! Feel free to fork the repository, make improvements, and submit pull requests.

### License
This project is licensed under the MIT License. See the LICENSE file for details.
