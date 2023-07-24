package carbonclient

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"

	"github.com/MacIt/pickle"
)

const PLAINTEXT_PORT int = 2003 // Not currently supported!
const PICKLE_PORT int = 2004

type CarbonClient struct {
	Host string
	Port int
}

type TimedMetricValue struct {
	Timestamp time.Time
	Value     interface{}
}

type TimedMetric struct {
	Path  string
	Value TimedMetricValue
}

func NewCarbonClient(host string, port int) *CarbonClient {
	return &CarbonClient{
		Host: host,
		Port: port,
	}
}

func (c *CarbonClient) prepareMetrics(metrics []TimedMetric) []pickle.Tuple {
	var stats []pickle.Tuple
	for _, metric := range metrics {
		stats = append(stats, pickle.Tuple{
			metric.Path,
			pickle.Tuple{metric.Value.Timestamp.Unix(), metric.Value.Value},
		})
	}
	return stats
}

func (c *CarbonClient) makeMessage(payload *bytes.Buffer) []byte {
	message := bytes.Buffer{}
	header := make([]byte, 4)
	binary.BigEndian.PutUint32(header, uint32(payload.Len()))

	message.Write(header)
	message.Write(payload.Bytes())

	return message.Bytes()
}

func (c *CarbonClient) SendMetrics(metrics []TimedMetric) error {
	stats := c.prepareMetrics(metrics)

	w := bytes.NewBuffer([]byte{})
	e := pickle.NewEncoder(w)

	err := e.Encode(stats)
	if err != nil {
		return err
	}

	message := c.makeMessage(w)

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port))
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(message)
	if err != nil {
		return err
	}

	return nil
}
