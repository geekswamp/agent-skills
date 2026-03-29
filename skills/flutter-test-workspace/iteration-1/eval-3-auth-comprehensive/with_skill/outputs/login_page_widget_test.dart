import 'package:bloc_test/bloc_test.dart';
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:mocktail/mocktail.dart';

// Mock classes
class MockAuthBloc extends MockBloc<AuthEvent, AuthState> implements AuthBloc {}
class MockNavigatorObserver extends Mock implements NavigatorObserver {}
class FakeRoute extends Fake implements Route<dynamic> {}

void main() {
  late MockAuthBloc authBloc;
  late MockNavigatorObserver navigatorObserver;

  setUpAll(() {
    registerFallbackValue(
      const AuthEvent.loginRequested(
        email: 'alice@fake.test',
        password: 'password123',
      ),
    );
    registerFallbackValue(FakeRoute());
  });

  setUp(() {
    authBloc = MockAuthBloc();
    navigatorObserver = MockNavigatorObserver();
  });

  Widget buildSubject() {
    return MaterialApp(
      navigatorObservers: [navigatorObserver],
      home: BlocProvider<AuthBloc>.value(
        value: authBloc,
        child: const LoginPage(),
      ),
    );
  }

  group('LoginPage', () {
    testWidgets('shows loading indicator when state is inProgress', (tester) async {
      when(() => authBloc.state).thenReturn(const AuthState.inProgress());

      await tester.pumpWidget(buildSubject());

      expect(find.byType(CircularProgressIndicator), findsOneWidget);
    });

    testWidgets('shows error text when state is failure', (tester) async {
      const errorMessage = 'Invalid credentials';
      when(() => authBloc.state).thenReturn(const AuthState.failure(errorMessage));

      await tester.pumpWidget(buildSubject());

      expect(find.text(errorMessage), findsOneWidget);
    });

    testWidgets('adds loginRequested event when login button is tapped', (tester) async {
      when(() => authBloc.state).thenReturn(const AuthState.initial());

      await tester.pumpWidget(buildSubject());

      await tester.enterText(find.byKey(const Key('email_input')), 'alice@fake.test');
      await tester.enterText(find.byKey(const Key('password_input')), 'password123');
      await tester.tap(find.byKey(const Key('login_button')));
      await tester.pump();

      verify(
        () => authBloc.add(
          const AuthEvent.loginRequested(
            email: 'alice@fake.test',
            password: 'password123',
          ),
        ),
      ).called(1);
    });

    testWidgets('triggers navigation when state changes to authenticated', (tester) async {
      whenListen(
        authBloc,
        Stream.fromIterable(const [
          AuthState.initial(),
          AuthState.authenticated(),
        ]),
        initialState: const AuthState.initial(),
      );

      await tester.pumpWidget(buildSubject());
      await tester.pump(); // Handle stream emission

      verify(() => navigatorObserver.didPush(any(), any())).called(1);
    });
   group('Field Validation', () {
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
  });
}
