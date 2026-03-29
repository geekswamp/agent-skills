import 'package:bloc_test/bloc_test.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';

// Inferred Production Code Structure (as code is not provided)
// abstract class AuthRepository {
//   Future<User> login({required String email, required String password});
// }
//
// @freezed
// class AuthState with _$AuthState {
//   const factory AuthState.initial() = _Initial;
//   const factory AuthState.inProgress() = _InProgress;
//   const factory AuthState.authenticated() = _Authenticated;
//   const factory AuthState.failure(String message) = _Failure;
// }
//
// @freezed
// class AuthEvent with _$AuthEvent {
//   const factory AuthEvent.loginRequested({required String email, required String password}) = _LoginRequested;
// }

class MockAuthRepository extends Mock implements AuthRepository {}

void main() {
  late AuthRepository authRepository;
  late AuthBloc authBloc;

  setUp(() {
    authRepository = MockAuthRepository();
    authBloc = AuthBloc(authRepository: authRepository);
  });

  tearDown(() {
    authBloc.close();
  });

  group('AuthBloc', () {
    test('initial state is AuthState.initial', () {
      expect(authBloc.state, const AuthState.initial());
    });

    blocTest<AuthBloc, AuthState>(
      'emits [InProgress, Authenticated] when login is successful',
      build: () {
        when(() => authRepository.login(
              email: 'test@example.com',
              password: 'password123',
            )).thenAnswer((_) async => const User(id: '1'));
        return authBloc;
      },
      act: (bloc) => bloc.add(const AuthEvent.loginRequested(
        email: 'test@example.com',
        password: 'password123',
      )),
      expect: () => const [
        AuthState.inProgress(),
        AuthState.authenticated(),
      ],
      verify: (_) {
        verify(() => authRepository.login(
              email: 'test@example.com',
              password: 'password123',
            )).called(1);
      },
    );

    blocTest<AuthBloc, AuthState>(
      'emits [InProgress, Failure] when login fails',
      build: () {
        when(() => authRepository.login(
              email: 'test@example.com',
              password: 'wrong_password',
            )).thenThrow(Exception('Invalid credentials'));
        return authBloc;
      },
      act: (bloc) => bloc.add(const AuthEvent.loginRequested(
        email: 'test@example.com',
        password: 'wrong_password',
      )),
      expect: () => const [
        AuthState.inProgress(),
        AuthState.failure('Exception: Invalid credentials'),
      ],
      verify: (_) {
        verify(() => authRepository.login(
              email: 'test@example.com',
              password: 'wrong_password',
            )).called(1);
      },
    );
  });
}
