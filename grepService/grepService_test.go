package grep

import (
	"testing"
)

func TestGrepCount(t *testing.T) {
	grep_service := new(GrepService)
	args := &GrepFormat{Flags: "c", Pattern: "Windows", File: "../test_log/test.log"}

	var reply ReplyFormat
	grep_service.Grep(args, &reply)

	want_match := FileMatchLines{"../test_log/test.log", 7}

	if len(reply.MatchStats) != 1 || reply.MatchStats[0] != want_match {
		t.Errorf("want %v, but got %v", want_match, reply.MatchStats[0])
	}
}

func TestGrepRegrexp(t *testing.T) {
	grep_service := new(GrepService)
	args := &GrepFormat{Flags: "Ec", Pattern: "html?", File: "../test_log/test.log"}

	var reply ReplyFormat
	grep_service.Grep(args, &reply)

	want_match := FileMatchLines{"../test_log/test.log", 6}

	if len(reply.MatchStats) != 1 || reply.MatchStats[0] != want_match {
		t.Errorf("want %v, but got %v", want_match, reply.MatchStats[0])
	}
}
