import 'package:bloc_test/bloc_test.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';

// Mock classes
class MockAuthRepository extends Mock implements AuthRepository {}

void main() {
  late MockAuthRepository repository;
  late AuthBloc authBloc;

  setUp(() {
    repository = MockAuthRepository();
    authBloc = AuthBloc(authRepository: repository);
  });

  tearDown(() {
    authBloc.close();
  });

  group('AuthBloc', () {
    test('initial state is AuthState.initial()', () {
      expect(authBloc.state, const AuthState.initial());
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
        return authBloc;
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
        ).thenThrow(Exception('Invalid credentials'));
        return authBloc;
      },
      act: (bloc) => bloc.add(
        const AuthEvent.loginRequested(
          email: 'alice@fake.test',
          password: 'wrong-password',
        ),
      ),
      expect: () => const [
        AuthState.inProgress(),
        AuthState.failure('Exception: Invalid credentials'),
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
  });
}
