package config

import (
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {
	tests := map[string]struct {
		envContent string
		expect     Config
	}{
		"full .env": {
			envContent: "HOST=1.2.3.4\nPORT=9000\nDBConn=abc",
			expect: Config{
				ServerHost: "1.2.3.4",
				ServerPort: "9000",
				DBConn:     "abc",
			},
		},

		"empty .env": {
			envContent: "",
			expect: Config{
				ServerHost: "localhost",
				ServerPort: "8080",
				DBConn:     "postgres://postgres:password@localhost:5432/learn_hub_auth?sslmode=disable",
			},
		},

		"partial .env": {
			envContent: "PORT=5050",
			expect: Config{
				ServerHost: "localhost",
				ServerPort: "5050",
				DBConn:     "postgres://postgres:password@localhost:5432/learn_hub_auth?sslmode=disable",
			},
		},

		"missing .env": {
			envContent: "__MISSING__",
			expect: Config{
				ServerHost: "localhost",
				ServerPort: "8080",
				DBConn:     "postgres://postgres:password@localhost:5432/learn_hub_auth?sslmode=disable",
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			keys := []string{"HOST", "PORT", "DBConn"}
			saved := map[string]*string{}
			for _, k := range keys {
				if v, ok := os.LookupEnv(k); ok {
					saved[k] = &v
				}
				_ = os.Unsetenv(k)
			}
			t.Cleanup(func() {
				for _, k := range keys {
					if v, ok := saved[k]; ok && v != nil {
						_ = os.Setenv(k, *v)
					} else {
						_ = os.Unsetenv(k)
					}
				}
			})

			origDir, err := os.Getwd()
			if err != nil {
				t.Fatalf("getwd: %v", err)
			}
			tempDir := t.TempDir()

			if tt.envContent != "__MISSING__" {
				if err := os.WriteFile(tempDir+"/.env", []byte(tt.envContent), 0644); err != nil {
					t.Fatalf("write .env: %v", err)
				}
			}

			if err := os.Chdir(tempDir); err != nil {
				t.Fatalf("chdir: %v", err)
			}
			t.Cleanup(func() {
				_ = os.Chdir(origDir)
			})

			cfg := NewConfig()

			if cfg.ServerHost != tt.expect.ServerHost {
				t.Fatalf("ServerHost = %q, want %q", cfg.ServerHost, tt.expect.ServerHost)
			}
			if cfg.ServerPort != tt.expect.ServerPort {
				t.Fatalf("ServerPort = %q, want %q", cfg.ServerPort, tt.expect.ServerPort)
			}
			if cfg.DBConn != tt.expect.DBConn {
				t.Fatalf("DBConn = %q, want %q", cfg.DBConn, tt.expect.DBConn)
			}
		})
	}
}
