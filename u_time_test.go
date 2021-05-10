package kerbalwzygo

import (
	"testing"
	"time"
)

func TestNowDatetimeStr(t *testing.T) {
	t.Log(Timestamp2DatetimeStr(NowTimestamp(), BJS))
	t.Log(BJSNowDatetimeStr())
	t.Log(BJSTodayDateStr())
	t.Log(UTCNowDatetimeStr())
	t.Log(UTCTodayDateStr())
	t.Log(Time2BJS(time.Now().In(time.UTC)))
}
