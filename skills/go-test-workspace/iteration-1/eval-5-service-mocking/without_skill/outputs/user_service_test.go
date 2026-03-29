package service

import (
	"errors"
	"testing"
)

// MockNotifier is a mock implementation of the Notifier interface
type MockNotifier struct {
	NotifyFunc func(message string) error
	Called     bool
	LastMsg    string
}

func (m *MockNotifier) Notify(message string) error {
	m.Called = true
	m.LastMsg = message
	if m.NotifyFunc != nil {
		return m.NotifyFunc(message)
	}
	return nil
}

func TestUserService_WelcomeUser(t *testing.T) {
	tests := []struct {
		name          string
		userName      string
		mockNotifyErr error
		wantErr       bool
		expectedMsg   string
	}{
		{
			name:        "successful welcome",
			userName:    "Alice",
			wantErr:     false,
			expectedMsg: "Welcome Alice",
		},
		{
			name:          "notifier error",
			userName:      "Bob",
			mockNotifyErr: errors.New("notification failed"),
			wantErr:       true,
			expectedMsg:   "Welcome Bob",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &MockNotifier{
				NotifyFunc: func(message string) error {
					return tt.mockNotifyErr
				},
			}
			s := &UserService{
				notifier: mock,
			}

			err := s.WelcomeUser(tt.userName)

			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.WelcomeUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !mock.Called {
				t.Errorf("UserService.WelcomeUser() expected notifier to be called")
			}

			if mock.LastMsg != tt.expectedMsg {
				t.Errorf("UserService.WelcomeUser() message = %v, want %v", mock.LastMsg, tt.expectedMsg)
			}
		})
	}
}
