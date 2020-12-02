package testutil

type testcommon interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Helper()
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Name() string
	Skip(args ...interface{})
	SkipNow()
	Skipf(format string, args ...interface{})
	Skipped() bool
}

func AssertEqual(t testcommon, a, b interface{}) {
	t.Helper()
	if a != b {
		t.Errorf("Not Equal. %d %d", a, b)
	}
}

func AssertNotEqual(t testcommon, a, b interface{}) {
	t.Helper()
	if a == b {
		t.Errorf("Equal. %d %d", a, b)
	}
}

func AssertTrue(t testcommon, a bool) {
	t.Helper()
	if !a {
		t.Errorf("Not True %t", a)
	}
}

func AssertFalse(t testcommon, a bool) {
	t.Helper()
	if a {
		t.Errorf("Not True %t", a)
	}
}

func AssertNil(t testcommon, a interface{}) {
	t.Helper()
	if a != nil {
		t.Error("Not Nil")
	}
}

func AssertNotNil(t testcommon, a interface{}) {
	t.Helper()
	if a == nil {
		t.Error("Is Nil")
	}
}
