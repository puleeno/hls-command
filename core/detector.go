package core

import "github.com/puleeno/tarantula/core/format"

type Detector struct {
	Detectors []format.Format
}

func (t *Detector) DetectFromURL(url string) *format.Format {
}
