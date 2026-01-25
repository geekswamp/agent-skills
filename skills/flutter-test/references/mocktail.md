# Mocking with mocktail (Flutter & Freezed)
mocktail is the preferred mocking library due to null-safety support.

## Creating Mocks
```dart
class MockUserRepository extends Mock implements UserRepository {}
```

## registerFallbackValue
Required when mocking methods with non-primitive arguments:

```dart
setUpAll(() {
  registerFallbackValue(const UserEvent.fetchUser(id: 0));
});
```

## Stubbing
```dart
when(() => repository.fetchUser(any()))
  .thenAnswer((_) async => user);
```

## Verification
```dart
verify(() => repository.fetchUser(1)).called(1);
verifyNever(() => repository.deleteUser(any()));
```