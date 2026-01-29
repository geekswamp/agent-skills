# Unit Testing BLoC / Cubit (Freezed-based)
This reference describes best practices for unit testing BLoC or Cubit that use Freezed for events and states.

Freezed supports two valid state modeling patterns. Tests MUST match the pattern used by the production code.

## Core principles
- Test **state transitions**, not implementation details.
- Each test must clearly describe:
  - Given (initial state)
  - When (event / method call)
  - Then (expected emitted states)
- Prefer multiple focused `blocTest` blocks over complex logic.
- Tests must be deterministic and isolated.

## State Pattern 1: Union-Based State (Sealed Classes)
Example state:

```dart
@freezed
class MyState with _$MyState {
  const factory MyState.initial() = _Initial;
  const factory MyState.inProgress() = _InProgress;
  const factory MyState.success(List<Item> items) = _Success;
  const factory MyState.failure(@Default('') String message) = _Failure;
}
```

Example test:

```dart
blocTest<MyBloc, MyState>(
  'emits [Loading, Success] when fetch succeeds',
  build: () => myBloc,
  act: (bloc) => bloc.add(const MyEvent.fetch()),
  expect: () => [
    const MyState.inProgress(),
    MyState.success(data),
  ],
);
```

### Assertion Rules (Union-Based)
- ✅ Assert concrete union factories
- ❌ **Do NOT** use abstract matchers (`isA<MyState>()`)
- ❌ **Do NOT** partially match union states

## State Pattern 2: Status-Based State (Single Data Class)
Example State:

```dart
enum MyStatus { initial, inProgress, success, failure }

@freezed
class MyState with _$MyState {
  const factory MyState({
    @Default(MyStatus.initial) MyStatus status,
    @Default(<Item>[]) List<Item> items,
    @Default('') String message,
  }) = _MyState;
}
```

Example test:

```dart
blocTest<MyBloc, MyState>(
  'emits loading then success when fetch succeeds',
  build: () => myBloc,
  act: (bloc) => bloc.add(const MyEvent.fetch()),
  expect: () => [
    const MyState(status: MyStatus.inProgress),
    MyState(
      status: MyStatus.success,
      items: data,
    ),
  ],
);
```

### Error Handling (Both Patterns)
- Always test failure paths if applicable.
- Error flow must `emit`:
  - loading → failure
- Assert error data only if it affects behavior.

Union-Based Error Example:

```dart
expect: () => [
  const MyState.inProgress(),
  const MyState.failure('Something went wrong'),
];
```

Status-Based Error Example:

```dart
expect: () => [
  initial.copyWith(status: MyStatus.inProgress),
  initial.copyWith(
    status: MyStatus.failure,
    message: 'Something went wrong',
  ),
];
```

## Anti-Patterns (ALL)
- Mixing union-based expectations with status-based states
- Using `isA<T>()` to assert state
- Testing private implementation details
- Adding fake assertions for coverage

## Final Rule (IMPORTANT)
A correct test for the wrong state pattern is still a wrong test.