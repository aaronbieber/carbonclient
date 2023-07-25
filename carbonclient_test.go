package carbonclient

import (
	"bytes"
	"reflect"
	"testing"
	"time"

	"github.com/MacIt/pickle"
)

func TestPrepareMetrics(t *testing.T) {
	client, _ := NewCarbonClient("", 0)

	var metrics []TimedMetric
	metrics = append(metrics, TimedMetric{
		Path: "stats.path",
		Value: TimedMetricValue{
			Timestamp: time.Unix(1234567890, 0), // Friday, February 13, 2009 11:31:30 PM
			Value:     42,
		}})

	var want []pickle.Tuple
	want = append(want, pickle.Tuple{
		"stats.path",
		pickle.Tuple{1234567890, 42},
	})

	got := client.prepareMetrics(metrics)

	if reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestMakeMessage(t *testing.T) {
	client, _ := NewCarbonClient("", 0)

	// "hello" in bytes, preceded by little-endian size, which is 5.
	want := []byte{0, 0, 0, 5, 104, 101, 108, 108, 111}

	b := bytes.Buffer{}
	b.WriteString("hello")
	got := client.makeMessage(&b)

	if !bytes.Equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
