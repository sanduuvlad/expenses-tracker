package config

import (
	"testing"
)

func TestGetEnv(t *testing.T) {
	t.Setenv("DB_HOST", "localhost")

	value, err := getEnv("DB_HOST")
	if err != nil {
		t.Fatal(err)
	}

	if value != "localhost" {
		t.Fatal(value)
	}
}

func TestGetEnvMissing(t *testing.T) {
	t.Setenv("DB_HOST", "")

	_, err := getEnv("DB_HOST")
	if err == nil {
		t.Fatal(err)
	}
}

func TestLoad(t *testing.T) {
	t.Setenv("SERVER_PORT", "8080")
	t.Setenv("DB_HOST", "localhost")
	t.Setenv("DB_PORT", "5433")
	t.Setenv("DB_USER", "expense_user")
	t.Setenv("DB_PASSWORD", "test_password")
	t.Setenv("DB_NAME", "expense_tracker")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	intTests := []struct {
		name string
		got  int
		want int
	}{
		{
			name: "server port",
			got:  cfg.Server.Port,
			want: 8080,
		},

		{
			name: "database port",
			got:  cfg.Database.Port,
			want: 5433,
		},
	}

	stringTests := []struct {
		name string
		got  string
		want string
	}{
		{
			name: "database host",
			got:  cfg.Database.Host,
			want: "localhost",
		},

		{
			name: "database user",
			got:  cfg.Database.User,
			want: "expense_user",
		},

		{
			name: "database password",
			got:  cfg.Database.Password,
			want: "test_password",
		},

		{
			name: "database name",
			got:  cfg.Database.Name,
			want: "expense_tracker",
		},
	}

	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("got %v, want %v", tt.got, tt.want)
			}
		})
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("got %v, want %v", tt.got, tt.want)
			}
		})
	}
}
