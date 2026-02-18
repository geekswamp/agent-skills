# Optional Test Data Conventions
This section defines standardized placeholder data that may be used when generating example tests, mock data, or UI assertions.

These values are optional and should be used only when realistic sample data is required.

## Person Names
If example person names are required, use only the following:

- Alice
- Bob
- John Doe
- Jean Doe
- Janet

Do not invent additional names unless explicitly required.

## Email Addresses
If example email addresses are required, use only the following:

- alice@fake.test
- bob@fake.test
- john.doe@fake.test
- jean.doe@test.fake
- janet@fake.test

Do not use real or production-like domains.

## Optional Fake UUID Values
Use these predefined UUID values when generating test data that requires `UUIDv4`.

### Valid UUIDv4
Use only the following valid UUIDv4 values:

- `550e8400-e29b-41d4-a716-446655440000`
- `3fa85f64-5717-4562-b3fc-2c963f66afa6`
- `9b2c1f4e-8d5a-4c6b-9f1e-123456789abc`

These values:

- Follow UUIDv4 format
- Use correct version (4)
- Are safe for deterministic testing

### Invalid UUID Examples
Use the following values when testing validation or error scenarios:

- `550e8400-e29b-11d4-a716-446655440000` (wrong version)
- `not-a-uuid`
- `123456`
- `550e8400e29b41d4a716446655440000` (missing hyphens)
- `550e8400-e29b-41d4-a716-44665544` (too short)

Do not generate random invalid UUID strings.

### UUID Usage Guidelines
- Use valid UUIDs for success scenarios.
- Use invalid UUIDs only when explicitly testing validation logic.
- Avoid mixing multiple UUID values in one test unless required.

## Optional Fake URLs
Use standardized URLs for network, routing, or validation tests.

### Valid URLs
Use only the following:

- `https://example.test`
- `https://api.example.test/v1/users`
- `https://cdn.example.test/assets/image.png`
- `https://app.example.test/login`

These domains are:

- Non-production
- Safe
- Deterministic

### Invalid URLs
Use the following for validation failure tests:

- `htp://invalid-url`
- `www.example.test` (missing scheme)
- `https:/broken-url.com`
- `://missing-scheme.com`
- `not-a-url`

Do not use real company domains or public services.

### URL Usage Guidelines
- Always include scheme (https://) in valid URLs.
- Prefer HTTPS over HTTP.
- Do not use localhost unless specifically testing local environment logic.
- Do not generate random URLs.

## Usage Guidelines
- Use consistent placeholder values across a single test file.
- Do not mix multiple placeholder identities unless necessary.
- Avoid excessive mock data unless it contributes to behavior validation.
- These values are for testing clarity only and must not represent real users.