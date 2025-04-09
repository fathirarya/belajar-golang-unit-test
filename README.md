# Belajar Golang Unit Test

## Repository Summary

This repository contains learning materials and examples focused on unit testing in Go programming language. Created by [fathirarya](https://github.com/fathirarya), it serves as a practical guide to understanding and implementing unit tests in Go applications.

## Key Components

### Helper Package
- **hello_world_test.go**: Contains basic unit tests for a hello world function
- **hello_world.go**: Implementation of simple hello world functions
- **table_test.go**: Demonstrates table-driven tests in Go

### Service and Repository Pattern Testing
- Demonstrates mocking techniques using the service-repository pattern
- Includes tests for both success and failure scenarios
- Shows dependency injection for testable code

### Repository Mock Implementation
- **category_repository_mock.go**: Mock implementation of repository interfaces
- Uses the `testify/mock` library to simulate repository behavior

### Entity Models
- **category.go**: Basic entity structures used in the application

### Service Layer Testing
- **category_service_test.go**: Tests for the category service
- Shows how to test different scenarios using mock repositories
- Validates service behavior independent of actual data storage

## Testing Techniques Demonstrated

1. Basic unit tests with `testing` package
2. Table-driven tests for testing multiple scenarios
3. Mocking external dependencies
4. Assertion techniques using the `testify/assert` package
5. Testing success and error cases
6. Setup and teardown patterns

## Key Takeaways

- Proper organization of test files alongside implementation
- Effective use of Go's testing conventions
- Implementation of mock objects for isolated testing
- Practical examples of service-repository pattern testing
- Clear demonstration of dependency injection for testability

This repository serves as an excellent resource for Go developers looking to improve their unit testing skills and understand best practices for writing testable code in Go applications.