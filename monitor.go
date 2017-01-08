package main

import (
	"fmt"
	"io"
)

type MonitorWriter struct {
	w     io.Writer
	total int64
	count int64
}

func (m *MonitorWriter) Write(p []byte) (int, error) {
	l := int64(len(p))
	m.count += l
	progress := float64(m.count) / float64(m.total) * float64(100)
	str := fmt.Sprintf("\r%0.1f %%", progress)
	m.w.Write([]byte(str))
	return int(l), nil
}

func NewMonitorWriter(w io.Writer, total int64) *MonitorWriter {
	return &MonitorWriter{w, total, 0}
}
