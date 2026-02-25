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

## Long Placeholder Text
If long body text is required (e.g. for UI rendering, scroll tests, layout tests), use the following standardized placeholder:

```
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris vitae aliquet nibh. Aliquam neque mi, accumsan volutpat ex eget, faucibus lobortis nisi. Sed elementum eu nisl vitae dapibus. Ut at dictum erat. Phasellus ac massa gravida, accumsan lectus ac, feugiat nunc. Nam mi erat, volutpat nec libero eu, dictum tempor ipsum. Morbi condimentum neque sit amet nibh condimentum, at aliquam est porta. Pellentesque ac metus sed erat varius tincidunt. In hac habitasse platea dictumst.
```

Do not generate random paragraphs or varying lorem ipsum text.

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

## Auth Feature Output Contract
When the feature is authentication, produce separate files:
- `auth_bloc_test.dart` for BLoC/Cubit unit tests.
- `login_page_widget_test.dart` (or equivalent) for widget tests.

Minimum expectations:
- Unit tests:
  - login success emits loading/inProgress -> authenticated
  - login failure emits loading/inProgress -> failure or unauthenticated
- Widget tests:
  - loading indicator is visible in loading/inProgress state
  - error message is visible in failure state
  - tapping login triggers expected BLoC event/action

## Minimal Auth Example (Reference Shape)
Use this section when the user asks for a concrete authentication test example and does not provide production code.

Important:
- Keep unit tests and widget tests in separate files.
- Match the exact state pattern used by production code (union-based or status-based).
- Do not mix assertion styles across patterns.

### `auth_bloc_test.dart` (Union-Based Freezed State)
```dart
import 'package:bloc_test/bloc_test.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';

class MockAuthRepository extends Mock implements AuthRepository {}

void main() {
  late MockAuthRepository repository;

  setUp(() {
    repository = MockAuthRepository();
  });

  blocTest<AuthBloc, AuthState>(
    'emits [inProgress, authenticated] when login succeeds',
    build: () {
      when(
        () => repository.login(
          email: 'alice@fake.test',
          password: 'password123',
        ),
      ).thenAnswer((_) async => const User(id: '550e8400-e29b-41d4-a716-446655440000'));

      return AuthBloc(authRepository: repository);
    },
    act: (bloc) => bloc.add(
      const AuthEvent.loginRequested(
        email: 'alice@fake.test',
        password: 'password123',
      ),
    ),
    expect: () => const [
      AuthState.inProgress(),
      AuthState.authenticated(),
    ],
    verify: (_) {
      verify(
        () => repository.login(
          email: 'alice@fake.test',
          password: 'password123',
        ),
      ).called(1);
    },
  );

  blocTest<AuthBloc, AuthState>(
    'emits [inProgress, failure] when login fails',
    build: () {
      when(
        () => repository.login(
          email: 'alice@fake.test',
          password: 'wrong-password',
        ),
      ).thenThrow(AuthException('Invalid credentials'));

      return AuthBloc(authRepository: repository);
    },
    act: (bloc) => bloc.add(
      const AuthEvent.loginRequested(
        email: 'alice@fake.test',
        password: 'wrong-password',
      ),
    ),
    expect: () => const [
      AuthState.inProgress(),
      AuthState.failure('Invalid credentials'),
    ],
    verify: (_) {
      verify(
        () => repository.login(
          email: 'alice@fake.test',
          password: 'wrong-password',
        ),
      ).called(1);
    },
  );
}
```

### `login_page_widget_test.dart` (Union-Based Freezed State)
```dart
import 'package:bloc_test/bloc_test.dart';
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:mocktail/mocktail.dart';

class MockAuthBloc extends MockBloc<AuthEvent, AuthState> implements AuthBloc {}
class MockNavigatorObserver extends Mock implements NavigatorObserver {}
class FakeRoute extends Fake implements Route<dynamic> {}

void main() {
  late MockAuthBloc authBloc;
  late MockNavigatorObserver navigatorObserver;

  setUpAll(() {
    registerFallbackValue(
      const AuthEvent.loginRequested(
        email: 'alice@fake.test',
        password: 'password123',
      ),
    );
    registerFallbackValue(FakeRoute());
  });

  setUp(() {
    authBloc = MockAuthBloc();
    navigatorObserver = MockNavigatorObserver();
  });

  Widget buildSubject() {
    return MaterialApp(
      navigatorObservers: [navigatorObserver],
      home: BlocProvider<AuthBloc>.value(
        value: authBloc,
        child: const LoginPage(),
      ),
    );
  }

  testWidgets('shows loading indicator when bloc emits inProgress', (tester) async {
    when(() => authBloc.state).thenReturn(const AuthState.initial());
    whenListen(
      authBloc,
      Stream.fromIterable(const [AuthState.inProgress()]),
      initialState: const AuthState.initial(),
    );

    await tester.pumpWidget(buildSubject());
    await tester.pump();

    expect(find.byType(CircularProgressIndicator), findsOneWidget);
    verifyNever(() => navigatorObserver.didPush(any(), any()));
  });

  testWidgets('shows error text when bloc emits failure', (tester) async {
    when(() => authBloc.state).thenReturn(const AuthState.initial());
    whenListen(
      authBloc,
      Stream.fromIterable(const [
        AuthState.inProgress(),
        AuthState.failure('Invalid credentials'),
      ]),
      initialState: const AuthState.initial(),
    );

    await tester.pumpWidget(buildSubject());
    await tester.pump();
    await tester.pump(const Duration(milliseconds: 300));

    expect(find.text('Invalid credentials'), findsOneWidget);
  });

  testWidgets('adds loginRequested event when login button is tapped', (tester) async {
    when(() => authBloc.state).thenReturn(const AuthState.initial());
    whenListen(
      authBloc,
      const Stream<AuthState>.empty(),
      initialState: const AuthState.initial(),
    );

    await tester.pumpWidget(buildSubject());

    await tester.enterText(find.byKey(const Key('email_input')), 'alice@fake.test');
    await tester.enterText(find.byKey(const Key('password_input')), 'password123');
    await tester.tap(find.byKey(const Key('login_button')));
    await tester.pump();

    verify(
      () => authBloc.add(
        const AuthEvent.loginRequested(
          email: 'alice@fake.test',
          password: 'password123',
        ),
      ),
    ).called(1);
  });

  testWidgets('pushes next page when bloc emits authenticated', (tester) async {
    when(() => authBloc.state).thenReturn(const AuthState.initial());
    whenListen(
      authBloc,
      Stream.fromIterable(const [AuthState.authenticated()]),
      initialState: const AuthState.initial(),
    );

    await tester.pumpWidget(buildSubject());
    await tester.pump();
    await tester.pump(const Duration(milliseconds: 300));

    verify(() => navigatorObserver.didPush(any(), any())).called(1);
  });
}
```

### Status-Based Variant (When Production State Uses `status`)
Use this variant only if production `AuthState` is a single data class with a `status` field:

```dart
final initialState = const AuthState();

blocTest<AuthBloc, AuthState>(
  'emits loading then failure when login fails',
  build: () => authBloc,
  act: (bloc) => bloc.add(
    const AuthEvent.loginRequested(
      email: 'alice@fake.test',
      password: 'wrong-password',
    ),
  ),
  expect: () => [
    initialState.copyWith(status: AuthStatus.inProgress),
    initialState.copyWith(
      status: AuthStatus.failure,
      message: 'Invalid credentials',
    ),
  ],
);

whenListen(
  mockAuthBloc,
  Stream.fromIterable([
    initialState.copyWith(status: AuthStatus.inProgress),
    initialState.copyWith(status: AuthStatus.failure, message: 'Invalid credentials'),
  ]),
  initialState: initialState,
);
```
