package bbs

import (
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/types"
)

type BoardDetail struct {
	Brdname            string            `json:"brdname"`
	RealTitle          []byte            `json:"title"` // Require to separate RealTitle, BoardClass, BoardType, because it's hard to parse in utf8
	BoardClass         []byte            `json:"class"`
	BoardType          []byte            `json:"type"` //□, ◎, Σ
	BM                 []UUserID         `json:"moderators"`
	BrdAttr            ptttype.BrdAttr   `json:"attr"`
	Gid                ptttype.Bid       `json:"pttgid"`
	Bid                ptttype.Bid       `json:"pttbid"`
	BBoardID           BBoardID          `json:"bid"`
	ChessCountry       ptttype.ChessCode `json:"chesscountry"`
	VoteLimitLogins    uint8             `json:"votelimitlogins"`
	BUpdate            types.Time4       `json:"bupdate"`
	PostLimitLogins    uint8             `json:"postlimitlogins"`
	BVote              uint8             `json:"bvote"`
	VTime              types.Time4       `json:"vtime"`
	Level              ptttype.PERM      `json:"level"`
	PermReload         types.Time4       `json:"permreload"`
	NUser              int32             `json:"nuser"`
	PostExpire         int32             `json:"postexpire"`
	EndGamble          types.Time4       `json:"endgamble"`
	PostType           []byte            `json:"posttype"`
	FastRecommendPause uint8             `json:"fastrecommendpause"`
	VoteLimitBadpost   uint8             `json:"votelimitbadpost"`

	LastPostTime types.Time4 `json:"last_post_time"`
	Total        int32       `json:"total"`

	IdxByName  string `json:"idx_name"`
	IdxByClass string `json:"idx_class"`

	Reason ptttype.RestrictReason `json:"reason"`
}

func NewBoardDetailFromRaw(boardDetailRaw *ptttype.BoardDetailRaw, bid ptttype.Bid) *BoardDetail {
	bmsRaw := boardDetailRaw.BM.ToBMs()
	bms := make([]UUserID, len(bmsRaw))
	for idx, each := range bmsRaw {
		bms[idx] = UUserID(types.CstrToString(each[:]))
	}

	boardDetail := &BoardDetail{
		Brdname:            types.CstrToString(boardDetailRaw.Brdname[:]),
		BoardClass:         types.CstrToBytes(boardDetailRaw.Title[:4]),
		BoardType:          types.CstrToBytes(boardDetailRaw.Title[5:7]),
		RealTitle:          types.CstrToBytes(boardDetailRaw.Title[7:]),
		BM:                 bms,
		BrdAttr:            boardDetailRaw.BrdAttr,
		Gid:                boardDetailRaw.Gid,
		Bid:                bid,
		BBoardID:           ToBBoardID(bid, &boardDetailRaw.Brdname),
		ChessCountry:       boardDetailRaw.ChessCountry,
		VoteLimitLogins:    boardDetailRaw.VoteLimitLogins,
		BUpdate:            boardDetailRaw.BUpdate,
		PostLimitLogins:    boardDetailRaw.VoteLimitLogins,
		BVote:              boardDetailRaw.BVote,
		VTime:              boardDetailRaw.VTime,
		Level:              boardDetailRaw.Level,
		PermReload:         boardDetailRaw.PermReload,
		NUser:              boardDetailRaw.ChildCount,
		PostExpire:         boardDetailRaw.PostExpire,
		EndGamble:          boardDetailRaw.EndGamble,
		PostType:           types.CstrToBytes(boardDetailRaw.PostType[:]),
		FastRecommendPause: boardDetailRaw.FastRecommendPause,
		VoteLimitBadpost:   boardDetailRaw.VoteLimitBadpost,

		LastPostTime: boardDetailRaw.LastPostTime,
		Total:        boardDetailRaw.Total,
	}
	boardDetail.IdxByName = serializeBoardIdxByNameStr(boardDetail.Brdname)
	boardDetail.IdxByClass = serializeBoardIdxByClassStr(boardDetail.BoardClass, boardDetail.Brdname)

	return boardDetail
}
