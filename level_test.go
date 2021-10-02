package pq_logger

import (
	"fmt"
	"strings"
	"testing"
)

func TestLevel_String(t *testing.T) {
	lvls := []Level{DebugLevel, InfoLevel, WarnLevel, ErrorLevel}
	lvlStrings := []LevelString{DebugLevelString, InfoLevelString, WarnLevelString, ErrorLevelString}

	for i, lvl := range lvls {
		if strings.EqualFold(fmt.Sprintf("%s", lvl.String()), string(lvlStrings[i])) {
			t.Log("valid string level")
		} else {
			t.Error("invalid string level")
		}
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("invalid level string must throw panic")
		} else {
			t.Log("valid check invalid level")
		}
	}()

	iLvl := Level(uint8(255))
	if iLvl.String() == "" {
		t.Errorf("invalid level string must not have custom level")
	} else {
		t.Errorf("this should be never executed")
	}
}
