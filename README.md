## Description
This is a simple tool written in Go that performs a mass reverse IP lookup using the Webscan API. It retrieves all domains associated with a given IP address and saves the results in a file.

## Features
- Supports mass reverse IP lookups
- Multi-threaded processing for faster results
- Saves results in real-time to an output file
- Colorful CLI output for better readability
- ASCII banner for aesthetics

## Requirements
- Go 1.16 or later
- Internet connection

## Installation
1. Clone this repository:
   ```sh
   git clone https://github.com/maw3six/Maw_rev.git
   ```
2. Navigate to the project directory:
   ```sh
   cd Maw_rev
   ```
3. Build the project:
   ```sh
   go build -o Maw_rev
   ```

## Usage
Run the tool with the following command:
```sh
./Maw_rev
```

### Input Parameters
- **Masukan list**: The filename containing a list of IP addresses (e.g., `ips.txt`)
- **Jumlah thread**: Number of concurrent threads for processing
- **Result Save as**: Output filename where the reversed domains will be stored

### Example
```sh
./Maw_rev
List : ips.txt
Thread : 10
Save as : results.txt
```

## Output Format
The tool will save only the reversed domains into the specified output file. Example:
```
example.com
anotherdomain.com
somedomain.net
```

## License
This project is open-source and available under the MIT License.

## Author
Coded by: [@maw3six]

