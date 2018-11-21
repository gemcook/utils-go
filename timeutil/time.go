package timeutil

import "time"

var dbDatetimeFormat = "2006-01-02 15:04:05"

// NowFunc is current time function.
// Default function is "time.Now".
var NowFunc = time.Now

// Now は `NowFunc` を使って現在時刻を返す
// `timeutil.NowFunc` に別の `func() time.Time` を代入することで、Now()の動作を上書き可能
func Now() time.Time {
	now := NowFunc()
	return now
}

// FormatDBNow はDBフォーマット文字列で現在時刻を返す
func FormatDBNow() string {
	return Now().Format(dbDatetimeFormat)
}

// FormatDB は指定時刻をDBフォーマット文字列に変換する
func FormatDB(t time.Time) string {
	return t.Format(dbDatetimeFormat)
}

// ParseDBDatetime はDBフォーマット文字列を`time.Time`に変換する
func ParseDBDatetime(dbTimestamp string) (time.Time, error) {
	t, err := time.ParseInLocation(dbDatetimeFormat, dbTimestamp, time.Local)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

// MustParseDBDatetime はDBフォーマット文字列を`time.Time`に変換する
func MustParseDBDatetime(dbTimestamp string) time.Time {
	t, err := ParseDBDatetime(dbTimestamp)
	if err != nil {
		panic(err)
	}
	return t
}
