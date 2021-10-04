package main

/*
typedef enum {
    ASCII_ENCODING = 0,
    EBCDIC_ENCODING = 1
} FileEncoding;
*/
import "C"

import (
	"bytes"
	"encoding/json"
	"unsafe"

	"github.com/moov-io/imagecashletter"
)

type FileEncoding int

const (
	ASCII_ENCODING FileEncoding = iota
	EBCDIC_ENCODING
)

//export image_cash_letter_file_to_json
func image_cash_letter_file_to_json(cFileEncoding C.FileEncoding, iclBytes unsafe.Pointer, n C.int) *C.char {
	fileEncoding := FileEncoding(cFileEncoding)
	buf := bytes.NewBuffer(C.GoBytes(iclBytes, n))

	readerOptions := []imagecashletter.ReaderOption{}
	if fileEncoding == ASCII_ENCODING {
		readerOptions = []imagecashletter.ReaderOption{
			imagecashletter.ReadVariableLineLengthOption(),
		}
	} else if fileEncoding == EBCDIC_ENCODING {
		readerOptions = []imagecashletter.ReaderOption{
			imagecashletter.ReadVariableLineLengthOption(),
			imagecashletter.ReadEbcdicEncodingOption(),
		}
	} else {
		// An invalid file encoding was specified.
		return nil
	}

	r := imagecashletter.NewReader(buf, readerOptions...)
	iclFile, iclFileReadErr := r.Read()
	if iclFileReadErr != nil {
		return nil
	}

	if err := iclFile.Validate(); err != nil {
		return nil
	}

	if err := iclFile.Create(); err != nil {
		return nil
	}

	jsonEncodedIclBuf := new(bytes.Buffer)
	if err := json.NewEncoder(jsonEncodedIclBuf).Encode(iclFile); err != nil {
		return nil
	}

	return C.CString(jsonEncodedIclBuf.String())
}

func main() {}
