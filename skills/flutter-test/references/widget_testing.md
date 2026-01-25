# Widget testing fundamentals
Widget tests verify UI behavior, not business logic.

## What to test
- Visible widgets
- User interactions
- State-driven UI changes

## What NOT to test
- Private widget methods
- Flutter framework internals
- Business logic (belongs in BLoC tests)

## Pumping strategy
- Use `pump()` for predictable state changes
- Avoid excessive `pumpAndSettle()`

Example:

```dart
await tester.pump();
await tester.pump(const Duration(milliseconds: 300));
```