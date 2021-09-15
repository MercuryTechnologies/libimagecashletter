package main

import "C"

import (
	"bytes"
	"encoding/json"
	"unsafe"

	"github.com/moov-io/imagecashletter"
)

//export image_cash_letter_to_json
func image_cash_letter_to_json(iclBytes unsafe.Pointer, n C.int) *C.char {
	buf := bytes.NewBuffer(C.GoBytes(iclBytes, n))

	r := imagecashletter.NewReader(buf, imagecashletter.ReadVariableLineLengthOption())
	iclFile, err := r.Read()
	if err != nil {
		return nil
	}

	if iclFile.Validate(); err != nil {
		return nil
	}

	if iclFile.Create(); err != nil {
		return nil
	}

	jsonEncodedIclBuf := new(bytes.Buffer)
	if err := json.NewEncoder(jsonEncodedIclBuf).Encode(iclFile); err != nil {
		return nil
	}

	return C.CString(jsonEncodedIclBuf.String())
}

func main() {}
