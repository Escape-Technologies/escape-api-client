# Escape API CLI Client

Escape API CLI Client is a command-line interface client built using Cobra for interacting with the Escape API.

## Usage

### Get Scan Details
Fetch details of a specific scan using its scan ID.
```sh
./escape-api-client get-scan [scanId] --api-key your_api_key_here
```

### List Applications in an Organization
Retrieve a list of applications associated with an organization using its organization ID.
```sh
./escape-api-client organization-applications [organizationId] --api-key your_api_key_here
```

### Start a Scan
Initiate a scan for a given application using its application ID.
```sh
./escape-api-client start-scan [applicationId] --api-key your_api_key_here --config config_override --commit commit_hash --introspection json_data_here
```

### Upload Introspection Data
Upload introspection data for a specific application using its application ID.
```sh
./escape-api-client upload-introspection [applicationId] --api-key your_api_key_here --introspection-response introspection_json_here
```

## Flags

- `--api-key`: API key for authorization (required for all commands)
- `--config`: Configuration override (optional, `start-scan` command)
- `--commit`: Commit hash (optional, `start-scan` command)
- `--introspection`: Introspection JSON (optional, `start-scan` command)
- `--introspection-response`: Introspection response (required for `upload-introspection` command)


## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/Escape-Technologies/escape-api-client.git
   ```

2. Navigate to the project directory:
   ```sh
   cd escape-api-client
   ```

3. Build the CLI client:
   ```sh
   make build 
   ```

4. Verify the installation by running:
   ```sh
   ./dist/{arch}/escape-api-client --help
   ```

## License

This project is licensed under the Apache 2.0 License.
