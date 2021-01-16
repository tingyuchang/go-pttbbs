package api

import (
	"reflect"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

func TestGetFavorites(t *testing.T) {
	setupTest()
	defer teardownTest()

	content0 := []byte{
		0x23, 0x0d, 0x03, 0x00, 0x02, 0x01, 0x01, 0x01,
		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x03, 0x01, 0x01, 0x02,
		0x01, 0x01, 0xb7, 0x73, 0xaa, 0xba, 0xa5, 0xd8,
		0xbf, 0xfd, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x03, 0x01, 0x02, 0x01, 0x01,
		0x09, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x08, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x01,
		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}

	params0 := &GetFavoritesParams{}
	path0 := &GetFavoritesPath{UserID: "CodingMan"}
	result0 := &GetFavoritesResult{Content: content0}

	type args struct {
		remoteAddr string
		uuserID    bbs.UUserID
		params     interface{}
		path       interface{}
	}
	tests := []struct {
		name           string
		args           args
		expectedResult *GetFavoritesResult
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			args:           args{remoteAddr: testIP, uuserID: "CodingMan", params: params0, path: path0},
			expectedResult: result0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := GetFavorites(tt.args.remoteAddr, tt.args.uuserID, tt.args.params, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFavorites() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := gotResult.(*GetFavoritesResult)
			got.MTime = 0
			if !reflect.DeepEqual(gotResult, tt.expectedResult) {
				t.Errorf("GetFavorites() = %v, want %v", gotResult, tt.expectedResult)
			}
		})
	}
}
