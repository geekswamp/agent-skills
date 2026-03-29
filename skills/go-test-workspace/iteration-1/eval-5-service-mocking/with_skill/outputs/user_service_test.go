package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"go-test-workspace/iteration-1/eval-5-service-mocking/with_skill/testutil"
)

// MockNotifier is a mock implementation of the Notifier interface.
type MockNotifier struct {
	mock.Mock
}

func (m *MockNotifier) Notify(message string) error {
	args := m.Called(message)
	return args.Error(0)
}

func TestUserService_WelcomeUser(t *testing.T) {
	tests := []struct {
		name          string
		userName      string
		mockNotifyMsg string
		mockError     error
		expectError   bool
	}{
		{
			name:          "successful welcome",
			userName:      testutil.NameAlice,
			mockNotifyMsg: "Welcome " + testutil.NameAlice,
			mockError:     nil,
			expectError:   false,
		},
		{
			name:          "notifier error",
			userName:      testutil.NameBob,
			mockNotifyMsg: "Welcome " + testutil.NameBob,
			mockError:     errors.New("notification failed"),
			expectError:   true,
		},
		{
			name:          "empty name welcome",
			userName:      "",
			mockNotifyMsg: "Welcome ",
			mockError:     nil,
			expectError:   false,
		},
	}

	table := make([]testutil.TableTestCase, 0, len(tests))
	for _, tt := range tests {
		tc := tt
		table = append(table, testutil.TableTestCase{
			Name: tc.name,
			Run: func(t *testing.T) {
				// Arrange
				mockNotifier := new(MockNotifier)
				mockNotifier.On("Notify", tc.mockNotifyMsg).Return(tc.mockError)

				svc := &UserService{
					notifier: mockNotifier,
				}

				// Act
				err := svc.WelcomeUser(tc.userName)

				// Assert
				if tc.expectError {
					testutil.RequireError(t, err)
					testutil.AssertEqual(t, tc.mockError, err)
				} else {
					testutil.RequireNoError(t, err)
				}

				mockNotifier.AssertExpectations(t)
			},
		})
	}

	testutil.RunTableTests(t, table)
}
