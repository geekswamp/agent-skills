# Widget Testing Fundamentals
Widget tests verify UI behavior, not business logic.

## What to Test
- Visible widgets
- User interactions
- State-driven UI changes

## What NOT to Test
- Private widget methods
- Flutter framework internals
- Business logic (belongs in BLoC tests)

## Pumping Strategy
- Use `pump()` for predictable state changes
- Avoid excessive `pumpAndSettle()`

Example:

```dart
await tester.pump();
await tester.pump(const Duration(milliseconds: 300));
```

## Widget Testing - Navigation
Use `mocktail` to verify navigation behavior in widget tests. Do not rely on real Navigator observers.

### Required Test Setup
Always define the following test utilities:

```dart
import 'package:mocktail/mocktail.dart';

class MockNavigatorObserver extends Mock implements NavigatorObserver {}

class FakeRoute extends Fake implements Route {}
```

Register fallback values before tests run:

```dart
setUpAll(() {
  registerFallbackValue(FakeRoute());
});
```

### Injecting `NavigatorObserver`
Provide the mocked observer via `MaterialApp`:

```dart
MaterialApp(
  navigatorObservers: [mockNavigatorObserver],
  home: TestedWidget(),
);
```

### Verifying Navigation
Verify navigation behavior using `mocktail`:

```dart
verify(() => mockNavigatorObserver.didPush(any(), any())).called(1);
```

Use `any()` for Route arguments unless route identity matters.

### Navigation Testing Rules
- Assert navigation via `NavigatorObserver`, not widget tree.
- Do not assert internal Navigator stack state.
- Avoid integration-style navigation flows in widget tests.
- Each navigation action should have exactly one verification.