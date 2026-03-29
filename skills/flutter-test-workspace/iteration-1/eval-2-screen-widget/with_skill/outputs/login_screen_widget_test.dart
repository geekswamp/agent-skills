import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:bloc_test/bloc_test.dart';
import 'package:mocktail/mocktail.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

// Assuming these are the types used in the LoginScreen
abstract class AuthEvent {}
abstract class AuthState {}
class AuthInitial extends AuthState {}
class AuthLoading extends AuthState {}

class MockAuthBloc extends MockBloc<AuthEvent, AuthState> implements AuthBloc {}

class AuthBloc extends Bloc<AuthEvent, AuthState> {
  AuthBloc() : super(AuthInitial());
}

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

  Widget createWidgetUnderTest() {
    return MaterialApp(
      home: BlocProvider<AuthBloc>.value(
        value: authBloc,
        child: LoginScreen(),
      ),
    );
  }

  testWidgets(
    'shows CircularProgressIndicator when state is AuthLoading',
    (WidgetTester tester) async {
      // Arrange
      when(() => authBloc.state).thenReturn(AuthLoading());

      // Act
      await tester.pumpWidget(createWidgetUnderTest());

      // Assert
      expect(find.byType(CircularProgressIndicator), findsOneWidget);
      expect(find.byType(ElevatedButton), findsNothing);
    },
  );

  testWidgets(
    'shows ElevatedButton when state is AuthInitial',
    (WidgetTester tester) async {
      // Arrange
      when(() => authBloc.state).thenReturn(AuthInitial());

      // Act
      await tester.pumpWidget(createWidgetUnderTest());

      // Assert
      expect(find.byType(ElevatedButton), findsOneWidget);
      expect(find.text('Login'), findsOneWidget);
      expect(find.byType(CircularProgressIndicator), findsNothing);
    },
  );
}
