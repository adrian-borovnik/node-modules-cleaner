# Node Modules Cleanner

A simple and efficient command-line tool written in Go to find and delete `node_modules` directories within a specified path.

## Description

This tool recursively searches for all `node_modules` directories starting from a given root directory (or the current directory by default). It then lists all the found directories and prompts the user for confirmation before deleting them. This is a useful utility for freeing up disk space by removing bulky `node_modules` folders from your projects.

## Features

-   **Recursive Search:** Finds all `node_modules` directories in the specified path and its subdirectories.
-   **Interactive Deletion:** Prompts for confirmation before deleting any files, preventing accidental data loss.
-   **Cross-Platform:** As a Go application, it can be built to run on Windows, macOS, and Linux.
-   **Standalone:** The compiled binary is standalone with no external dependencies.

## Installation

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/adrian-borovnik/node-modules-cleaner.git
    cd node-modules-cleaner
    ```

2.  **Build the application:**
    You can use the provided `Makefile` to build the application.
    ```bash
    make build
    ```
    This will create a binary in the `bin/` directory.

## Usage

Run the cleanner by specifying a path, or run it without arguments to scan the current directory.

**Scan the current directory:**

```bash
./bin/nmclean
```

**Scan a specific directory:**

```bash
./bin/nmclean /path/to/your/projects
```

The tool will list all the `node_modules` directories it finds and then ask for your confirmation before proceeding with the deletion.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue if you have any suggestions or find any bugs.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
