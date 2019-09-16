package twiml

import (
	"bytes"
	"testing"
)

func TestSay(t *testing.T) {
	r := NewResponse()
	r.Say("Hello, world!", &SayAttr{})

	var b bytes.Buffer
	err := r.ToXML(&b)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	expected := "<Response>\n  <Say>Hello, world!</Say>\n</Response>"
	actual := b.String()
	if expected != actual {
		t.Errorf("Expected \"%s\" but got \"%s\"", expected, actual)
	}
}

func TestDialSimple(t *testing.T) {
	r := NewResponse()
	r.DialSimple("+12223334455", &DialAttr{})

	var b bytes.Buffer
	err := r.ToXML(&b)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	expected := "<Response>\n  <Dial>+12223334455</Dial>\n</Response>"
	actual := b.String()
	if expected != actual {
		t.Errorf("Expected \"%s\" but got \"%s\"", expected, actual)
	}
}

func TestDialNumber(t *testing.T) {
	r := NewResponse()
	r.Dial(&DialAttr{}).Number("+12223334455", &NumberAttr{})

	var b bytes.Buffer
	err := r.ToXML(&b)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	expected := "<Response>\n  <Dial>\n    <Number>+12223334455</Number>\n  </Dial>\n</Response>"
	actual := b.String()
	if expected != actual {
		t.Errorf("Expected \"%s\" but got \"%s\"", expected, actual)
	}
}

func TestDialConference(t *testing.T) {
	r := NewResponse()
	r.Dial(&DialAttr{}).Conference("Conference", &ConferenceAttr{Record: "do-not-record"})

	var b bytes.Buffer
	err := r.ToXML(&b)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	expected := "<Response>\n  <Dial>\n    <Conference record=\"do-not-record\">Conference</Conference>\n  </Dial>\n</Response>"
	actual := b.String()
	if expected != actual {
		t.Errorf("Expected \"%s\" but got \"%s\"", expected, actual)
	}
}

func TestDialClient(t *testing.T) {
	r := NewResponse()
	r.Dial(&DialAttr{}).Client("client", &ClientAttr{})

	var b bytes.Buffer
	err := r.ToXML(&b)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	expected := "<Response>\n  <Dial>\n    <Client>client</Client>\n  </Dial>\n</Response>"
	actual := b.String()
	if expected != actual {
		t.Errorf("Expected \"%s\" but got \"%s\"", expected, actual)
	}
}

func TestDialSIP(t *testing.T) {
	r := NewResponse()
	r.Dial(&DialAttr{}).Sip("sip", &SipAttr{})

	var b bytes.Buffer
	err := r.ToXML(&b)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	expected := "<Response>\n  <Dial>\n    <Sip>sip</Sip>\n  </Dial>\n</Response>"
	actual := b.String()
	if expected != actual {
		t.Errorf("Expected \"%s\" but got \"%s\"", expected, actual)
	}
}

func TestRecord(t *testing.T) {
	r := NewResponse()
	r.Record(&RecordAttr{})

	var b bytes.Buffer
	err := r.ToXML(&b)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	expected := "<Response>\n  <Record></Record>\n</Response>"
	actual := b.String()
	if expected != actual {
		t.Errorf("Expected \"%s\" but got \"%s\"", expected, actual)
	}
}

func TestHangup(t *testing.T) {
	r := NewResponse()
	r.Hangup()

	var b bytes.Buffer
	err := r.ToXML(&b)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	expected := "<Response>\n  <Hangup></Hangup>\n</Response>"
	actual := b.String()
	if expected != actual {
		t.Errorf("Expected \"%s\" but got \"%s\"", expected, actual)
	}
}

func TestReject(t *testing.T) {
	r := NewResponse()
	r.Reject(&RejectAttr{})

	var b bytes.Buffer
	err := r.ToXML(&b)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	expected := "<Response>\n  <Reject></Reject>\n</Response>"
	actual := b.String()

	if expected != actual {
		t.Errorf("Expected \"%s\" but got \"%s\"", expected, actual)
	}
}

func TestPause(t *testing.T) {
	r := NewResponse()
	r.Pause(&PauseAttr{})

	var b bytes.Buffer
	err := r.ToXML(&b)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	expected := "<Response>\n  <Pause></Pause>\n</Response>"
	actual := b.String()
	if expected != actual {
		t.Errorf("Expected \"%s\" but got \"%s\"", expected, actual)
	}
}

func TestPlay(t *testing.T) {
	r := NewResponse()
	r.Play("value", &PlayAttr{})

	var b bytes.Buffer
	err := r.ToXML(&b)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	expected := "<Response>\n  <Play>value</Play>\n</Response>"
	actual := b.String()
	if expected != actual {
		t.Errorf("Expected \"%s\" but got \"%s\"", expected, actual)
	}
}

func TestGatherSay(t *testing.T) {
	r := NewResponse()
	r.Gather(&GatherAttr{}).Say("Hello, world!", &SayAttr{})

	var b bytes.Buffer
	err := r.ToXML(&b)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	expected := "<Response>\n  <Gather>\n    <Say>Hello, world!</Say>\n  </Gather>\n</Response>"
	actual := b.String()
	if expected != actual {
		t.Errorf("Expected \"%s\" but got \"%s\"", expected, actual)
	}
}

func TestGatherPlay(t *testing.T) {
	r := NewResponse()
	r.Gather(&GatherAttr{}).Play("value", &PlayAttr{})

	var b bytes.Buffer
	err := r.ToXML(&b)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	expected := "<Response>\n  <Gather>\n    <Play>value</Play>\n  </Gather>\n</Response>"
	actual := b.String()
	if expected != actual {
		t.Errorf("Expected \"%s\" but got \"%s\"", expected, actual)
	}
}

func TestGatherPause(t *testing.T) {
	r := NewResponse()
	r.Gather(&GatherAttr{}).Pause(&PauseAttr{})

	var b bytes.Buffer
	err := r.ToXML(&b)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	expected := "<Response>\n  <Gather>\n    <Pause></Pause>\n  </Gather>\n</Response>"
	actual := b.String()
	if expected != actual {
		t.Errorf("Expected \"%s\" but got \"%s\"", expected, actual)
	}
}
