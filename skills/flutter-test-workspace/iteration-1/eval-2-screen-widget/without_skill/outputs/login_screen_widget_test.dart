import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:mocktail/mocktail.dart';

// Mocking dependencies
class MockAuthBloc extends Mock implements AuthBloc {}
class MockAuthState extends Mock implements AuthState {}

// Define the required states and bloc for the test to be self-contained or referenceable
abstract class AuthState {}
class AuthInitial extends AuthState {}
class AuthLoading extends AuthState {}

class AuthBloc extends Mock implements Bloc<AuthEvent, AuthState> {}
abstract class AuthEvent {}

class LoginScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: BlocBuilder<AuthBloc, AuthState>(
        builder: (context, state) {
          if (state is AuthLoading) return CircularProgressIndicator();
          return ElevatedButton(onPressed: () {}, child: Text('Login'));
        },
      ),
    );
  }
}

void main() {
  late AuthBloc authBloc;

  setUp(() {
    authBloc = MockAuthBloc();
  });

  testWidgets('shows loading indicator when state is AuthLoading', (WidgetTester tester) async {
    // Arrange
    when(() => authBloc.state).thenReturn(AuthLoading());
    when(() => authBloc.stream).thenAnswer((_) => Stream.value(AuthLoading()));

    // Act
    await tester.pumpWidget(
      MaterialApp(
        home: BlocProvider<AuthBloc>.value(
          value: authBloc,
          child: LoginScreen(),
        ),
      ),
    );

    // Assert
    expect(find.byType(CircularProgressIndicator), findsOneWidget);
    expect(find.byType(ElevatedButton), findsNothing);
  });

  testWidgets('shows login button when state is AuthInitial', (WidgetTester tester) async {
    // Arrange
    when(() => authBloc.state).thenReturn(AuthInitial());
    when(() => authBloc.stream).thenAnswer((_) => Stream.value(AuthInitial()));

    // Act
    await tester.pumpWidget(
      MaterialApp(
        home: BlocProvider<AuthBloc>.value(
          value: authBloc,
          child: LoginScreen(),
        ),
      ),
    );

    // Assert
    expect(find.byType(ElevatedButton), findsOneWidget);
    expect(find.byType(CircularProgressIndicator), findsNothing);
  });
}
