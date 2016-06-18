package goptv

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var UTCDATA = "2006-01-02 15:04:05.999999999 +0000 UTC"

func IntV(pointer interface{}) (int64, error) {
	var buf bytes.Buffer
	buf = BufV(pointer)
	return strconv.ParseInt(buf.String(), 10, 64)

}
func TimeV(pointer interface{}) (time.Time, error) {
	ts := strings.Trim(StringV(pointer), "time.Time{")
	return time.Parse(UTCDATA, strings.Trim(ts, "}"))

}
func StringV(pointer interface{}) string {
	var buf bytes.Buffer
	buf = BufV(pointer)
	return strings.Trim(buf.String(), "\"")

}
func BufV(pointer interface{}) (buf bytes.Buffer) {
	v := reflect.ValueOf(pointer)
	pointerValue(&buf, v)
	return

}
func pointerValue(w io.Writer, val reflect.Value) {
	if val.Kind() == reflect.Ptr && val.IsNil() {
		w.Write([]byte("<nil>"))
		return
	}

	v := reflect.Indirect(val)

	switch v.Kind() {
	case reflect.String:
		fmt.Fprintf(w, `"%s"`, v)
	case reflect.Slice:
		w.Write([]byte{'['})
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				w.Write([]byte{' '})
			}

			pointerValue(w, v.Index(i))
		}

		w.Write([]byte{']'})
		return
	case reflect.Struct:
		if v.Type().Name() != "" {
			w.Write([]byte(v.Type().String()))
		}

		// special handling of Timestamp values
		if v.Type() == reflect.TypeOf(time.Time{}) {
			fmt.Fprintf(w, "{%s}", v.Interface())
			return
		}

		w.Write([]byte{'{'})

		var sep bool
		for i := 0; i < v.NumField(); i++ {
			fv := v.Field(i)
			if fv.Kind() == reflect.Ptr && fv.IsNil() {
				continue
			}
			if fv.Kind() == reflect.Slice && fv.IsNil() {
				continue
			}

			if sep {
				w.Write([]byte(", "))
			} else {
				sep = true
			}

			w.Write([]byte(v.Type().Field(i).Name))
			w.Write([]byte{':'})
			pointerValue(w, fv)
		}

		w.Write([]byte{'}'})
	default:
		if v.CanInterface() {
			fmt.Fprint(w, v.Interface())
		}
	}
}
