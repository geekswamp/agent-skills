import 'package:bloc_test/bloc_test.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class CounterCubit extends Cubit<int> {
  CounterCubit() : super(0);
  void increment() => emit(state + 1);
  void decrement() => emit(state - 1);
}

void main() {
  group('CounterCubit', () {
    late CounterCubit counterCubit;

    setUp(() {
      counterCubit = CounterCubit();
    });

    tearDown(() {
      counterCubit.close();
    });

    test('initial state is 0', () {
      expect(counterCubit.state, 0);
    });

    blocTest<CounterCubit, int>(
      'emits [1] when increment is called',
      build: () => counterCubit,
      act: (cubit) => cubit.increment(),
      expect: () => [1],
    );

    blocTest<CounterCubit, int>(
      'emits [-1] when decrement is called',
      build: () => counterCubit,
      act: (cubit) => cubit.decrement(),
      expect: () => [-1],
    );

    blocTest<CounterCubit, int>(
      'emits [1, 2] when increment is called twice',
      build: () => counterCubit,
      act: (cubit) => cubit..increment()..increment(),
      expect: () => [1, 2],
    );

    blocTest<CounterCubit, int>(
      'emits [1, 0] when increment and then decrement is called',
      build: () => counterCubit,
      act: (cubit) => cubit..increment()..decrement(),
      expect: () => [1, 0],
    );
  });
}
