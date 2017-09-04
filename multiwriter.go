// Copyright 2017 John Scherff
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"path/filepath"
	"bufio"
	"fmt"
	"log"
	"os"
	"io"

	"github.com/RackSec/srslog"
)

// MultiWriter is an io.Writer that sends output to multiple destinations.
type MultiWriter struct {
	writers	[]io.Writer
	buffers	[]*bufio.Writer
	files	[]*os.File
}

// NewMultiWriter returns an initialized MultiWriter object.
func NewMultiWriter() (this *MultiWriter) {
	return new(MultiWriter)
}

// Add adds one or more writers to MultiWriter.
func (this *MultiWriter) Add(writers ...io.Writer) {
	for _, w := range writers {
		this.writers = append(this.writers, w)
	}
}

// AddFiles adds one or more file writers to MultiWriter.
func (this *MultiWriter) AddFiles(files ...string) {

	for _, f := range files {

		var err error
		var h *os.File

		if err = os.MkdirAll(filepath.Dir(f), LogDirMode); err == nil {
			if h, err = os.OpenFile(f, LogFileFlags, LogFileMode); err == nil {
				b := bufio.NewWriter(h)
				this.Add(b)
				this.buffers = append(this.buffers, b)
				this.files = append(this.files, h)
			}
		}

		if err != nil {
			log.Printf("%v", err)
		}
	}
}

// AddSyslog appends a syslog writer to MultiWriter.
func (this *MultiWriter) AddSyslog(proto, raddr, tag string, pri srslog.Priority) {

	if s, err := srslog.Dial(proto, raddr, pri, tag); err == nil {
		this.Add(s)
	} else {
		log.Printf("%v", err)
	}
}

// Write writes output to each writer in MultiWriter.
func (this *MultiWriter) Write(p []byte) (n int, err error) {

	var errs int

	for _, w := range this.writers {
		if n, err = w.Write(p); err != nil {
			errs++
		}
	}

	if errs > 0 {
		err = fmt.Errorf("%d write errors", errs)
	}

	return n, err
}

//WriteString converts string input to []byte and then calls Write.
func (this *MultiWriter) WriteString(s string) (n int, err error) {
	return this.Write([]byte(s))
}

// Count returns the number of writers in MultiWriter.
func (this *MultiWriter) Count() (n int) {
	return len(this.writers)
}

// Flush flushes underlying bufio writers in MultiWriter.
func (this *MultiWriter) Flush() (err error) {

	var errs int

	for _, b := range this.buffers {
		if err = b.Flush(); err != nil {
			errs++
		}
	}

	if errs > 0 {
		err = fmt.Errorf("%d flush errors", errs)
	}

	return err
}

// Sync syncs underlying file writers in MultiWriter.
func (this *MultiWriter) Sync() {

	for _, f := range this.files {
		f.Sync()
	}
}

// Close syncs and closes underlying file writers in MultiWriter.
func (this *MultiWriter) Close() (err error) {

	err = this.Flush()
	this.Sync()

	for _, f := range this.files {
		f.Close()
	}

	return err
}
