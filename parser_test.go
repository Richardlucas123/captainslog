package captainslog

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestUnmarshal(t *testing.T) {
	b := []byte("<191>2006-01-02T15:04:05.999999-07:00 host.example.org test: hello world\n")

	var msg SyslogMsg
	err := Unmarshal(b, &msg)
	if err != nil {
		t.Error(err)
	}

	if want, got := Local7, msg.Pri.Facility; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := Debug, msg.Pri.Severity; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	ts := msg.Time

	if want, got := 2006, ts.Year(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := time.Month(1), ts.Month(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 2, ts.Day(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 15, ts.Hour(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 4, ts.Minute(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 5, ts.Second(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 999999, ts.Nanosecond()/1000; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	_, zoneOffsetSecs := ts.Zone()
	if want, got := -25200, zoneOffsetSecs; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := "host.example.org", msg.Host; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := "test:", msg.Tag; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := false, msg.IsCee; want != got {
		t.Errorf("want '%v', got '%v'", want, got)
	}

	if want, got := " hello world", msg.Content; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := 0, bytes.Compare(b, msg.Bytes()); want != got {
		t.Errorf("want '%v', got '%v'", want, got)
	}

	if want, got := string(b), msg.String(); want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}

func TestUnmarshalDateNoMicros(t *testing.T) {
	b := []byte("<191>2006-01-02T15:04:05.999-07:00 host.example.org test: hello world\n")

	var msg SyslogMsg
	err := Unmarshal(b, &msg)
	if err != nil {
		t.Error(err)
	}

	if want, got := Local7, msg.Pri.Facility; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := Debug, msg.Pri.Severity; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	ts := msg.Time

	if want, got := 2006, ts.Year(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := time.Month(1), ts.Month(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 2, ts.Day(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 15, ts.Hour(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 4, ts.Minute(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 5, ts.Second(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 999000, ts.Nanosecond()/1000; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	_, zoneOffsetSecs := ts.Zone()
	if want, got := -25200, zoneOffsetSecs; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := "host.example.org", msg.Host; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := "test:", msg.Tag; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := false, msg.IsCee; want != got {
		t.Errorf("want '%v', got '%v'", want, got)
	}

	if want, got := " hello world", msg.Content; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := 0, bytes.Compare(b, msg.Bytes()); want != got {
		t.Errorf("want '%v', got '%v'", want, got)
	}

	if want, got := string(b), msg.String(); want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}

func TestUnmarshalDateNoMillis(t *testing.T) {
	b := []byte("<171>2015-12-18T18:08:17+00:00 host.example.org test: hello world\n")

	var msg SyslogMsg
	err := Unmarshal(b, &msg)
	if err != nil {
		t.Error(err)
	}

	if want, got := Local5, msg.Pri.Facility; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := Err, msg.Pri.Severity; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	ts := msg.Time

	if want, got := 2015, ts.Year(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := time.Month(12), ts.Month(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 18, ts.Day(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 18, ts.Hour(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 8, ts.Minute(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 17, ts.Second(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 0, ts.Nanosecond()/1000; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	_, zoneOffsetSecs := ts.Zone()
	if want, got := 0, zoneOffsetSecs; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := "host.example.org", msg.Host; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := "test:", msg.Tag; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := false, msg.IsCee; want != got {
		t.Errorf("want '%v', got '%v'", want, got)
	}

	if want, got := " hello world", msg.Content; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := 0, bytes.Compare(b, msg.Bytes()); want != got {
		t.Errorf("want '%v', got '%v'", want, got)
	}

	if want, got := string(b), msg.String(); want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}

func TestUnmarshalCeeSpace(t *testing.T) {
	b := []byte("<191>2006-01-02T15:04:05.999999-07:00 host.example.org test: @cee:{\"a\":\"b\"}\n")

	var msg SyslogMsg
	err := Unmarshal(b, &msg)
	if err != nil {
		t.Error(err)
	}

	if want, got := true, msg.IsCee; want != got {
		t.Errorf("want '%v', got '%v'", want, got)
	}

	if want, got := " @cee:", msg.Cee; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := "{\"a\":\"b\"}", msg.Content; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := string(b), msg.String(); want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}

func TestUnmarshalCeeNoSpace(t *testing.T) {
	b := []byte("<191>2006-01-02T15:04:05.999999-07:00 host.example.org test:@cee:{\"a\":\"b\"}\n")
	var msg SyslogMsg
	err := Unmarshal(b, &msg)
	if err != nil {
		t.Error(err)
	}

	if want, got := true, msg.IsCee; want != got {
		t.Errorf("want '%v', got '%v'", want, got)
	}

	if want, got := "@cee:", msg.Cee; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := string(b), msg.String(); want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}

func TestUnmarshalCeeEarlyBufferBeforeColon(t *testing.T) {
	b := []byte("<191>2006-01-02T15:04:05.999999-07:00 host.example.org test:@cee\n")
	var msg SyslogMsg
	err := Unmarshal(b, &msg)
	if err != nil {
		t.Error(err)
	}

	if want, got := false, msg.IsCee; want != got {
		t.Errorf("want '%v', got '%v'", want, got)
	}

	if want, got := "", msg.Cee; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := "@cee", msg.Content; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}

func TestUnmarshalCeeEarlyBufferAfterColon(t *testing.T) {
	b := []byte("<191>2006-01-02T15:04:05.999999-07:00 host.example.org test:@cee:\n")
	var msg SyslogMsg
	err := Unmarshal(b, &msg)

	if want, got := ErrBadContent, err; want != got {
		t.Errorf("want '%v', got '%v'", want, got)
	}

	if want, got := true, msg.IsCee; want != got {
		t.Errorf("want '%v', got '%v'", want, got)
	}
}

func unmarshalCeeButNotCee(t *testing.T, b []byte) {
	var msg SyslogMsg
	err := Unmarshal(b, &msg)
	if err != nil {
		t.Error(err)
	}

	if want, got := false, msg.IsCee; want != got {
		t.Errorf("want '%v', got '%v'", want, got)
	}
}

func TestUnmarshalCeeButNotCee(t *testing.T) {
	b := []byte("<191>2006-01-02T15:04:05.999999-07:00 host.example.org test:@cee{\"a\":\"b\"}\n")
	unmarshalCeeButNotCee(t, b)

	b = []byte("<191>2006-01-02T15:04:05.999999-07:00 host.example.org test:@ce{\"a\":\"b\"}\n")
	unmarshalCeeButNotCee(t, b)

	b = []byte("<191>2006-01-02T15:04:05.999999-07:00 host.example.org test:@c{\"a\":\"b\"}\n")
	unmarshalCeeButNotCee(t, b)

	b = []byte("<191>2006-01-02T15:04:05.999999-07:00 host.example.org test:@{\"a\":\"b\"}\n")
	unmarshalCeeButNotCee(t, b)
}

func TestUnmarshalNoContent(t *testing.T) {
	b := []byte("<191>2006-01-02T15:04:05.999999-07:00 host.example.org test:\n")

	var msg SyslogMsg
	err := Unmarshal(b, &msg)

	if want, got := ErrBadContent, err; want != got {
		t.Errorf("want '%v', got '%v'", want, got)
	}
}

func TestUnmarshalTagEndHandling(t *testing.T) {
	b := []byte("<191>2006-01-02T15:04:05.999999-07:00 host.example.org test: hello world\n")

	var msg SyslogMsg
	err := Unmarshal(b, &msg)
	if err != nil {
		t.Error(err)
	}

	if want, got := "test:", msg.Tag; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	b = []byte("<191>2006-01-02T15:04:05.999999-07:00 host.example.org test hello world\n")
	err = Unmarshal(b, &msg)
	if err != nil {
		t.Error(err)
	}

	if want, got := "test", msg.Tag; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := string(b), msg.String(); want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}

func TestUnmarshalUnixTime(t *testing.T) {
	b := []byte("<38>Mon Jan  2 15:04:05 host.example.org test: hello world\n")

	var msg SyslogMsg
	err := Unmarshal(b, &msg)
	if err != nil {
		t.Error(err)
	}

	ts := msg.Time

	if want, got := 0, ts.Year(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := time.Month(1), ts.Month(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 2, ts.Day(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 15, ts.Hour(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 4, ts.Minute(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 5, ts.Second(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}
}

func TestUnmarshalTimeANSIC(t *testing.T) {
	b := []byte("<38>Mon Jan  2 15:04:05 2006 host.example.org test: hello world\n")

	var msg SyslogMsg
	err := Unmarshal(b, &msg)
	if err != nil {
		t.Error(err)
	}

	ts := msg.Time

	if want, got := time.Month(1), ts.Month(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 2, ts.Day(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 15, ts.Hour(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 4, ts.Minute(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 5, ts.Second(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := string(b), msg.String(); want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}

func TestUnmarshalTimeUnixDate(t *testing.T) {
	b := []byte("<38>Mon Jan  2 15:04:05 MST 2006 host.example.org test: hello world\n")

	var msg SyslogMsg
	err := Unmarshal(b, &msg)
	if err != nil {
		t.Error(err)
	}

	ts := msg.Time

	if want, got := time.Month(1), ts.Month(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 2, ts.Day(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 15, ts.Hour(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 4, ts.Minute(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 5, ts.Second(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	zone, _ := ts.Zone()
	if want, got := "MST", zone; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}

func TestUnmarshalTimeNoYear(t *testing.T) {
	b := []byte("<38>Mon Jan  2 15:04:05 host.example.org test: hello world\n")

	var msg SyslogMsg
	err := Unmarshal(b, &msg)
	if err != nil {
		t.Error(err)
	}

	ts := msg.Time

	if want, got := time.Month(1), ts.Month(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 2, ts.Day(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 15, ts.Hour(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 4, ts.Minute(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := 5, ts.Second(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}
}

func TestUnmarshalNoPriority(t *testing.T) {
	b := []byte("2006-01-02T15:04:05.999999-07:00 host.example.org test: hello world\n")

	var msg SyslogMsg
	err := Unmarshal(b, &msg)

	if want, got := ErrBadPriority, err; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}

func TestUnmarshalNoPriorityEnd(t *testing.T) {
	b := []byte("<1912006-01-02T15:04:05.999999-07:00 host.example.org test: hello world\n")

	var msg SyslogMsg
	err := Unmarshal(b, &msg)

	if want, got := ErrBadPriority, err; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}

func TestUnmarshalPriorityTooLong(t *testing.T) {
	b := []byte("<9999>2006-01-02T15:04:05.999999-07:00 host.example.org test: hello world\n")

	var msg SyslogMsg
	err := Unmarshal(b, &msg)

	if want, got := ErrBadPriority, err; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}

func TestUnmarshalPriorityTruncated(t *testing.T) {
	b := []byte("<99\n")

	var msg SyslogMsg
	err := Unmarshal(b, &msg)

	if want, got := ErrBadPriority, err; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}

func TestUnmarshalDateTruncated(t *testing.T) {
	b := []byte("<191>2006-01-02T15:0")

	var msg SyslogMsg
	err := Unmarshal(b, &msg)

	if want, got := ErrBadTime, err; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}

func TestUnmarshalHostTruncated(t *testing.T) {
	b := []byte("<191>2006-01-02T15:04:05.999999-07:00 host.examp")

	var msg SyslogMsg
	err := Unmarshal(b, &msg)

	if want, got := ErrBadHost, err; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}

func TestUnmarshalNoHost(t *testing.T) {
	b := []byte("<191>2006-01-02T15:04:05.999999-07:00 ")

	var msg SyslogMsg
	err := Unmarshal(b, &msg)

	if want, got := ErrBadHost, err; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}

func TestUnmarshalTagTruncated(t *testing.T) {
	b := []byte("<191>2006-01-02T15:04:05.999999-07:00 host.example.org tes")

	var msg SyslogMsg
	err := Unmarshal(b, &msg)

	if want, got := ErrBadTag, err; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}

func TestUnmarshalNoTag(t *testing.T) {
	b := []byte("<191>2006-01-02T15:04:05.999999-07:00 host.example.org ")

	var msg SyslogMsg
	err := Unmarshal(b, &msg)

	if want, got := ErrBadTag, err; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}

func TestUnmarshalContentNotTerminated(t *testing.T) {
	b := []byte("<191>2006-01-02T15:04:05.999999-07:00 host.example.org test: hello wo")

	var msg SyslogMsg
	err := Unmarshal(b, &msg)

	if want, got := ErrBadContent, err; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}

func TestUnmarshalPriNotNumber(t *testing.T) {
	b := []byte("<1a1>2006-01-02T15:04:05.999999-07:00 host.example.org test: hello world\n")
	var msg SyslogMsg
	err := Unmarshal(b, &msg)

	if want, got := ErrBadPriority, err; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}

func testFuzzFindings(fuzzData string, t *testing.T) {
	b := []byte(fuzzData)

	var msg SyslogMsg
	err := Unmarshal(b, &msg)

	if want, got := false, err == nil; want != got {
		t.Errorf("want '%v', got '%v'", want, got)
	}
}

func TestFuzzFindings(t *testing.T) {
	inputs := []string{
		"<0>Mon Jan 00 00:00:000 0 ",
		"<0>Mon Jan 00 00:00:000 :",
	}

	for _, fuzzData := range inputs {
		testFuzzFindings(fuzzData, t)
	}
}

func ExampleUnmarshal() {
	b := []byte("<191>2006-01-02T15:04:05.999999-07:00 host.example.org test: hello world\n")

	var msg SyslogMsg
	err := Unmarshal(b, &msg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Syslog message was from host '%s'", msg.Host)
	// Output: Syslog message was from host 'host.example.org'

}

func BenchmarkParserParse(b *testing.B) {
	m := []byte("<191>2006-01-02T15:04:05.999999-07:00 host.example.org test: hello world\n")

	var msg SyslogMsg

	for i := 0; i < b.N; i++ {
		b.SetBytes(int64(len(m)))
		err := Unmarshal(m, &msg)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkParserParseCEE(b *testing.B) {
	m := []byte("<191>2006-01-02T15:04:05.999999-07:00 host.example.org test: @cee:{\"a\":\"b\"}\n")

	var msg SyslogMsg

	for i := 0; i < b.N; i++ {
		err := Unmarshal(m, &msg)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkParserParseLeastLikelyTime(b *testing.B) {
	m := []byte("<38>Mon Jan  2 15:04:05 host.example.org test: hello world\n")

	var msg SyslogMsg

	for i := 0; i < b.N; i++ {
		err := Unmarshal(m, &msg)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkParserParseAndString(b *testing.B) {
	m := []byte("<191>2006-01-02T15:04:05.999999-07:00 host.example.org test: hello world\n")

	var msg SyslogMsg

	for i := 0; i < b.N; i++ {
		err := Unmarshal(m, &msg)
		if err != nil {
			panic(err)
		}
		_ = msg.String()
	}
}

func BenchmarkParserParseAndBytes(b *testing.B) {
	m := []byte("<191>2006-01-02T15:04:05.999999-07:00 host.example.org test: hello world\n")

	var msg SyslogMsg

	for i := 0; i < b.N; i++ {
		err := Unmarshal(m, &msg)
		if err != nil {
			panic(err)
		}
		_ = msg.Bytes()
	}
}
