# bloc_test usage guidelines (with Freezed)
Use `bloc_test` for all Bloc/Cubit unit tests.

## Basic structure
Success test:

```dart
blocTest<MyBloc, MyState>(
  'emits loading then success when fetch succeeds',
  build: () => myBloc,
  act: (bloc) => bloc.add(const MyEvent.fetch()),
  expect: () => [
    myBloc.state.copyWith(status: MyStatus.loading),
    myBloc.state.copyWith(
      status: MyStatus.success,
      items: data,
    ),
  ],
);
```

Failure test:

```dart
blocTest<MyBloc, MyState>(
  'emits loading then failure when fetch fails',
  build: () => myBloc,
  act: (bloc) => bloc.add(const MyEvent.fetch()),
  expect: () => [
    myBloc.state.copyWith(status: MyStatus.loading),
    myBloc.state.copyWith(
      status: MyStatus.failure,
      errorMessage: 'Something went wrong',
    ),
  ],
);
```

## Freezed events
- Always use `const` constructors for Freezed events.
- Prefer explicit event names in test descriptions.

Example:

```dart
act: (bloc) => bloc.add(const UserEvent.fetchUser(id: 1));
```

## Mock interaction verification
Use `verify()` when behavior matters:

```dart
verify(() => repository.fetchUser(1)).called(1);
```

Do not verify interactions unless they are part of the behavior contract.

### Status-based Freezed State
If the state is a single Freezed data class with a `status` field (not a union per factory):

- Do NOT expect per-factory states.
- Always assert concrete state objects with expected field values.
- Prefer `copyWith` from the initial state for resilience.

Example:

```dart
expect: () => [
  initialState.copyWith(status: MyStatus.loading),
  initialState.copyWith(status: MyStatus.success, data: result),
];
```