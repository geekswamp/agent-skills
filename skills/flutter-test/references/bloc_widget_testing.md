# Widget Testing with BLoC (Freezed-based)
This reference covers widget testing when UI depends on BLoC states defined using Freezed unions.

## Injecting BLoC
Use `BlocProvider.value` with a mocked BLoC:

```dart
BlocProvider<MyBloc>.value(
  value: mockBloc,
  child: MyWidget(),
);
```

## Mocking BLoC States
Success test:

```dart
final initialState = const MyState();

whenListen(
  mockBloc,
  Stream.fromIterable([
    initialState.copyWith(status: MyStatus.loading),
    initialState.copyWith(
      status: MyStatus.success,
      items: data,
    ),
  ]),
  initialState: initialState,
);
```

Failure test:

```dart
final initialState = const MyState();

whenListen(
  mockBloc,
  Stream.fromIterable([
    initialState.copyWith(status: MyStatus.loading),
    initialState.copyWith(
      status: MyStatus.failure,
      errorMessage: 'Failed to load data',
    ),
  ]),
  initialState: initialState,
);
```

## Asserting Freezed States via UI
Instead of checking state directly, assert visible UI:

- loading → `CircularProgressIndicator`
- success → `Text` / `ListView`
- error → Error message

## Example Assertions
```dart
expect(find.byType(CircularProgressIndicator), findsOneWidget);
expect(find.text('User loaded'), findsOneWidget);
```

## Avoid
- Testing `.when()` logic inside widget tests  
- Asserting exact widget tree structure

### Status-based Freezed State with whenListen
If the BLoC state is a single Freezed data class with a `status` field:

- Do NOT use union-style factories in `whenListen`.
- Emit concrete state objects with expected field values.
- Prefer using `copyWith` from the initial state.

Example:

```dart
final initialState = const MyState();

whenListen(
  mockBloc,
  Stream.fromIterable([
    initialState.copyWith(status: MyStatus.loading),
    initialState.copyWith(status: MyStatus.success, data: result),
  ]),
  initialState: initialState,
);
```