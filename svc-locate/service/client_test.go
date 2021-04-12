package service

import (
	"net/rpc"
	"testing"
)

func TestClient_Close(t *testing.T) {
	type fields struct {
		conn *rpc.Client
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:   "default",
			fields: fields(*NewClient()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				conn: tt.fields.conn,
			}
			if err := c.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Client.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_Locate(t *testing.T) {
	type fields struct {
		conn *rpc.Client
	}
	type args struct {
		scanList ScanList
		location *LocationInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"default",
			fields(*NewClient()),
			args{
				ScanList{
					{"a", ""},
					{"b", ""},
				},
				&LocationInfo{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				conn: tt.fields.conn,
			}
			if err := c.Locate(tt.args.scanList, tt.args.location); (err != nil) != tt.wantErr {
				t.Errorf("Client.Locate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"default",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(); got == nil {
				t.Errorf("NewClient() = %v", got)
			}
		})
	}
}

func setup() {
	go RunDefaultServer(Config{LogPath: "../"})
}

func TestMain(t *testing.M) {
	setup()
	t.Run()
}
