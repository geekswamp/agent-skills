# Unit Testing BLoC / Cubit (Freezed-based)
This reference describes best practices for unit testing BLoC or Cubit that use Freezed for events and states.

## Core principles
- Test **state transitions**, not implementation details.
- Each test must clearly describe:
  - Given (initial state)
  - When (event / method call)
  - Then (expected emitted states)
- Prefer table-like grouping via multiple `blocTest` blocks.

## Freezed considerations
- Freezed provides:
  - Value equality
  - Exhaustive unions
- Always assert **concrete state instances**, not abstract conditions.

✅ Good:

```dart
expect: () => [
  const MyState.loading(),
  const MyState.success(data),
]
```

❌ Avoid:

```dart
expect: () => [isA<MyState>()]
```

## Error Handling
- Always test error paths if the BLoC can fail.
- Prefer domain errors (Failure objects) over raw exceptions.
