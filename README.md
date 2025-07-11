# GoLarsa 🚀

A powerful CLI tool for generating standardized Go microservice folder structures with complete boilerplate code.

## Features

✨ **Complete Service Architecture Generation**

- Service layer with comprehensive CRUD operations
- Repository layer with MongoDB integration
- Handler layer with REST API endpoints
- Models with validation and DTOs
- Dependency injection setup

🔧 **Template-Based Code Generation**

- Uses Go templates for flexible code generation
- Automatically adapts to your project's module name
- Follows enterprise-grade patterns and best practices

🏗️ **Generated Service Structure**

```
pkg/services/<service-name>/
├── <service-name>.go              # Service interface + implementation
├── repo/<service-name>.repo.go    # MongoDB repository
├── models/<service-name>.models.go # DTOs, entities, validation
├── handler/<service-name>.handler.go # REST API handlers
└── di/di.go                       # Dependency injection setup
```

## Installation

### Install Globally

```bash
go install github.com/fadi-alhabib/golarsa@latest
```

### Verify Installation

```bash
golarsa --help
```

## Usage

### Prerequisites

- Go 1.21 or higher
- Must be run from a Go module root directory (where `go.mod` exists)

### Generate a New Service

```bash
# Basic usage
golarsa service user

# This creates a complete user service with all layers
golarsa service product
golarsa service order
```

### Command Structure

```
golarsa service [service-name]
```

## Generated Code Features

### Service Layer (`<service-name>.go`)

- **Get**: Paginated retrieval with MongoDB aggregation pipelines
- **GetAll**: Non-paginated retrieval with filtering
- **GetById**: Single entity retrieval with validation
- **Save**: Create operations (template for implementation)
- **Update**: Update operations (template for implementation)
- **Delete**: Soft delete with user tracking

### Repository Layer (`repo/<service-name>.repo.go`)

- Generic repository interface
- MongoDB implementation with `samber/do` dependency injection
- Aggregation pipeline support
- Filter-based operations

### Handler Layer (`handler/<service-name>.handler.go`)

- Complete REST API endpoints:
  - `GET /` - Paginated list
  - `GET /all` - Non-paginated list
  - `GET /{id}` - Single entity
  - `POST /` - Create entity
  - `PUT /{id}` - Update entity
  - `PATCH /{id}/toggle` - Toggle operations
  - `DELETE /{id}` - Soft delete
  - `POST /v2` - Alternative create endpoint
- Chi router integration
- Authentication middleware support
- JSON request/response handling

### Models Layer (`models/<service-name>.models.go`)

- DTO structures for API requests
- Main entity with audit fields (CreatedAt, CreatedBy, UpdatedAt, UpdatedBy)
- Pagination wrapper
- Validation methods
- Helper methods

### Dependency Injection (`di/di.go`)

- Service registration with `samber/do`
- Repository initialization
- Handler setup and route registration

## Example Output

When you run `golarsa service user`, it generates:

```
✓ Created: pkg/services/user/di
✓ Created: pkg/services/user/handler
✓ Created: pkg/services/user/models
✓ Created: pkg/services/user/repo
✓ Created: pkg/services/user/user.go
✓ Created: pkg/services/user/repo/user.repo.go
✓ Created: pkg/services/user/models/user.models.go
✓ Created: pkg/services/user/handler/user.handler.go
✓ Created: pkg/services/user/di/di.go

🎉 Service 'user' created successfully!
📁 Service structure created at: pkg/services/user
```

## Key Technologies Integrated

- **MongoDB**: Document database with aggregation pipelines
- **Chi Router**: Fast HTTP router for Go
- **samber/do**: Dependency injection container
- **Go Templates**: Template-based code generation
- **Microservices Architecture**: Enterprise-grade service patterns

## Development

### Build from Source

```bash
git clone https://github.com/fadi-alhabib/golarsa.git
cd golarsa
go build -o golarsa .
```

### Install Locally

```bash
go install .
```

## Requirements

- **Go 1.21+**: Modern Go version with generics support
- **Go Module**: Must run from a directory with `go.mod` file
- **File Permissions**: Write access to create service directories

## Architecture

GoLarsa follows a clean architecture pattern:

1. **Handler Layer**: HTTP request handling, validation, response formatting
2. **Service Layer**: Business logic, orchestration, data transformation
3. **Repository Layer**: Data access, MongoDB operations, query building
4. **Models Layer**: Data structures, validation, serialization
5. **DI Layer**: Dependency injection, service registration

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

**Fadi Bassam Al Habib** - [GitHub](https://github.com/fadi-alhabib)

---

⭐ If this tool helps you, please consider giving it a star!
