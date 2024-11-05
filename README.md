# Contract Driven Development (CDD)

## Abstract

Contract Driven Development (CDD) is an approach to software development that emphasizes defining and adhering to clear contracts between different components or services in a system. This method helps teams work independently on various parts of a project while ensuring compatibility and reducing integration issues. CDD combines ideas from Design by Contract, Test-Driven Development, and API-first design principles to create a more efficient and robust development process.

## Tools

This repository demonstrates the use of Contract Driven Development with the following tools:

### OpenAPI TypeScript

[openapi-typescript](https://github.com/drwpow/openapi-typescript) is a TypeScript code generator for OpenAPI 3.0 and Swagger 2.0 schemas. It allows you to generate TypeScript types from your OpenAPI specifications, ensuring type safety and consistency between your API contract and client-side code.

### oapi-codegen (Go)

[oapi-codegen](https://github.com/deepmap/oapi-codegen) is a Go code generator for OpenAPI 3.0 schemas. It generates Go server and client boilerplate code from OpenAPI 3.0 specifications, making it easier to implement and consume APIs that adhere to the defined contract.

## Getting Started

1. Clone this repository
2. Install the required dependencies
3. Generate code from the OpenAPI specification using the provided tools
4. Implement your services following the generated contracts

## Benefits of CDD

- Improved collaboration between frontend and backend teams
- Early detection of compatibility issues
- Easier maintenance and documentation of APIs
- Facilitates parallel development of different system components

## Contributing

We welcome contributions to this project. Please feel free to submit issues and pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
