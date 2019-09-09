package redfish

import (
	"testing"

	_ "github.com/bmc-toolbox/bmclib/logging"
	"github.com/stmcginnis/gofish"
)

func TestNew1(t *testing.T) {
	type args struct {
		ip       string
		username string
		password string
	}

	var expected = args{ip: "127.0.0.1", username: "foo", password: "bar"}

	tests := []struct {
		name    string
		args    args
		wantR   *Redfish
		wantErr bool
	}{
		{name: "TestNew", args: expected, wantR: &Redfish{ip: expected.ip, clientConfig: &gofish.ClientConfig{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotR, err := New(tt.args.ip, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %#v, wantErr %#v", err, tt.wantErr)
				return
			}

			if gotR == nil {
				t.Errorf("expected a *Redfish, got nil")
			}

			if gotR.clientConfig.HTTPClient == nil {
				t.Errorf("expected a http.Client in clientConfig instance, got nil")
			}

			if gotR.clientConfig == nil {
				t.Errorf("expected a *gofish.ClientConfig, got nil")
			}

			if gotR.ip != expected.ip {
				t.Errorf("expected IP in Redfish instance - %s got %s", expected.ip, gotR.ip)
			}

			if gotR.clientConfig.Username != expected.username {
				t.Errorf("expected username in clientConfig instance - %s got %s", expected.username, gotR.clientConfig.Username)
			}

			if gotR.clientConfig.Password != expected.password {
				t.Errorf("expected password in clientConfig instance - %s got %s", expected.password, gotR.clientConfig.Password)
			}

		})
	}
}
