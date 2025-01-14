package api

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
)

func TestLoadBoardSummary(t *testing.T) {
	setupTest(t.Name())
	defer teardownTest(t.Name())

	params := &LoadBoardSummaryParams{}

	path := &LoadBoardSummaryPath{
		BBoardID: "6_ALLPOST",
	}

	expected := testBoardSummary6

	type args struct {
		remoteAddr string
		uuserID    bbs.UUserID
		params     interface{}
		path       interface{}
	}
	tests := []struct {
		name            string
		args            args
		expectedResults LoadBoardSummaryResult
		wantErr         bool
	}{
		// TODO: Add test cases.
		{
			args:            args{remoteAddr: testIP, uuserID: "SYSOP", params: params, path: path},
			expectedResults: expected,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotResults, err := LoadBoardSummary(tt.args.remoteAddr, tt.args.uuserID, tt.args.params, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadBoardSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			testutil.TDeepEqual(t, "got", gotResults, tt.expectedResults)
		})
		wg.Wait()
	}
}
