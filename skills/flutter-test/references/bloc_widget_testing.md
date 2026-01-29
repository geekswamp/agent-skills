# Widget Testing with BLoC (Freezed-based)
This reference covers widget testing when UI depends on BLoC states defined using Freezed.

Freezed supports two valid state modeling patterns. Widget tests MUST match the pattern used by the production code.

## Injecting BLoC
Always inject a mocked BLoC using `BlocProvider.value`:

```dart
BlocProvider<MyBloc>.value(
  value: mockBloc,
  child: MyWidget(),
);
```
### State Pattern 1: Union-Based Freezed State
Example State:

```dart
@freezed
class MyState with _$MyState {
  const factory MyState.initial() = _Initial;
  const factory MyState.inProgress() = _InProgress;
  const factory MyState.success(List<Item> items) = _Success;
  const factory MyState.failure(@Default('') String message) = _Failure;
}
```

### Mocking BLoC States (Union-Based)
Success flow:

```dart
whenListen(
  mockBloc,
  Stream.fromIterable([
    const MyState.inProgress(),
    MyState.success(data),
  ]),
  initialState: const MyState.initial(),
);
```

Failure flow:

```dart
whenListen(
  mockBloc,
  Stream.fromIterable([
    const MyState.inProgress(),
    const MyState.failure('Failed to load data'),
  ]),
  initialState: const MyState.initial(),
);
```

### Assertion Rules (Union-Based)
- Emit concrete union factories only
- **Do NOT** emit partial or abstract states
- **Do NOT** mix with status-based patterns

## State Pattern 2: Status-Based Freezed State
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

### Mocking BLoC States (Status-Based)
Success flow:

```dart
final initialState = const MyState();

whenListen(
  mockBloc,
  Stream.fromIterable([
    initialState.copyWith(status: MyStatus.inProgress),
    initialState.copyWith(
      status: MyStatus.success,
      items: data,
    ),
  ]),
  initialState: initialState,
);
```

Failure flow:

```dart
final initialState = const MyState();

whenListen(
  mockBloc,
  Stream.fromIterable([
    initialState.copyWith(status: MyStatus.inProgress),
    initialState.copyWith(
      status: MyStatus.failure,
      message: 'Failed to load data',
    ),
  ]),
  initialState: initialState,
);
```

## Assertion Rules (Status-Based)
- Emit full state objects
- Prefer `copyWith` from initial state
- **Do NOT** use union-style factories
- **Do NOT** emit abstract matchers

Example Assertions:

```dart
expect(find.byType(CircularProgressIndicator), findsOneWidget);
expect(find.text('User loaded'), findsOneWidget);
```