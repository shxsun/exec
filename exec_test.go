package exec

import (
	"testing"
	"time"
)

func TestNormal(t *testing.T) {
	cmd := Command("echo", "-n", "hello")
	output, err := cmd.Output()
	if err != nil {
		t.Errorf("expect normal, but got error: %s\n", err.Error())
	}
	if string(output) != "hello" {
		t.Errorf("expect 'hello', but got: %s\n", string(output))
	}
}

func TestNormalTimeout(t *testing.T) {
    cmd := Command("sleep", "2")
    err := cmd.Start()
    if err != nil {
        t.Error(err)
    }
    err = cmd.WaitTimeout(time.Second * 1)
    if err == nil {
        t.Errorf("expect ErrTimeout, but err is nil")
    }
    if err != ErrTimeout {
        t.Errorf("expect ErrTimeout, but receive %s", err.Error())
    }
}

func TestTimeout(t *testing.T) {
	cmd := Command("sleep", "2")
	cmd.Timeout = time.Second * 1
	_, err := cmd.Output()
	if err != ErrTimeout {
		t.Errorf("expect ErrTimeout, but got: %s\n", err.Error())
	}
}

// FIXME
func TestCleanNormal(t *testing.T) {
}

// FIXME
func TestCleanTimeout(t *testing.T) {
}
