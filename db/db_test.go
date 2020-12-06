package db

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInsertCreateKEY(t *testing.T) {
	type args struct {
		myKEY string
		uID   string
	}
	tests := []struct {
		args    args
		want    string
		wantErr bool // err nil check
		wantOk  bool // myKEY == want
	}{
		{
			args:    args{myKEY: "test-key-no1", uID: "test-user"},
			want:    "test-key-no1",
			wantErr: false,
			wantOk:  true,
		},
		{
			args:    args{myKEY: "test-key-false", uID: "test-user-false"},
			want:    "test-key-no1",
			wantErr: false,
			wantOk:  false,
		},
	}
	for _, tt := range tests {
		k, err := InsertCreateKEY(tt.args.myKEY, tt.args.uID)
		if (err != nil) != tt.wantErr {
			t.Errorf("InsertCreateKEY() error = %v, wantErr %v", err, tt.wantErr)
		}
		if (k.SelfKey != tt.want) == tt.wantOk {
			t.Errorf("InsertCreateKEY() = %v, want %v", k.SelfKey, tt.want)
		}
	}
}
