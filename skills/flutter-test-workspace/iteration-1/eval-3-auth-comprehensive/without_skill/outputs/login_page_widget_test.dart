import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:mocktail/mocktail.dart';
import 'package:bloc_test/bloc_test.dart';

// Mock classes
class MockAuthBloc extends MockBloc<AuthEvent, AuthState> implements AuthBloc {}

void main() {
  late MockAuthBloc authBloc;

  setUpAll(() {
    registerFallbackValue(const AuthEvent.loginRequested(
      email: 'test@example.com',
      password: 'password123',
    ));
  });

  setUp(() {
    authBloc = MockAuthBloc();
  });

  Widget buildSubject() {
    return MaterialApp(
      home: BlocProvider<AuthBloc>.value(
        value: authBloc,
        child: const LoginPage(),
      ),
    );
  }

  group('LoginPage', () {
    testWidgets('shows CircularProgressIndicator when state is inProgress', (tester) async {
      when(() => authBloc.state).thenReturn(const AuthState.inProgress());

      await tester.pumpWidget(buildSubject());

      expect(find.byType(CircularProgressIndicator), findsOneWidget);
    });

    testWidgets('shows error message when state is failure', (tester) async {
      const errorMessage = 'Invalid credentials';
      when(() => authBloc.state).thenReturn(const AuthState.failure(errorMessage));

      await tester.pumpWidget(buildSubject());

      expect(find.text(errorMessage), findsOneWidget);
    });

    testWidgets('adds loginRequested event when login button is tapped', (tester) async {
      when(() => authBloc.state).thenReturn(const AuthState.initial());

      await tester.pumpWidget(buildSubject());

      await tester.enterText(find.byKey(const Key('email_input')), 'test@example.com');
      await tester.enterText(find.byKey(const Key('password_input')), 'password123');
      await tester.tap(find.byKey(const Key('login_button')));
      await tester.pump();

      verify(() => authBloc.add(const AuthEvent.loginRequested(
            email: 'test@example.com',
            password: 'password123',
          ))).called(1);
    });

    testWidgets('shows validation errors for empty fields', (tester) async {
      when(() => authBloc.state).thenReturn(const AuthState.initial());

      await tester.pumpWidget(buildSubject());
      await tester.tap(find.byKey(const Key('login_button')));
      await tester.pump();

      expect(find.text('Email is required'), findsOneWidget);
      expect(find.text('Password is required'), findsOneWidget);
      verifyNever(() => authBloc.add(any()));
    });
  });
}
