/*
 * Copyright (c) 2014 Mark Samman <https://github.com/marksamman/bencode>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */

package bencode

import (
	"bytes"
	"strconv"
)

type encoder struct {
	bytes.Buffer
}

func (encoder *encoder) writeString(str string) {
	encoder.WriteString(strconv.Itoa(len(str)))
	encoder.WriteByte(':')
	encoder.WriteString(str)
}

func (encoder *encoder) writeInt(v int64) {
	encoder.WriteByte('i')
	encoder.WriteString(strconv.FormatInt(v, 10))
	encoder.WriteByte('e')
}

func (encoder *encoder) writeUint(v uint64) {
	encoder.WriteByte('i')
	encoder.WriteString(strconv.FormatUint(v, 10))
	encoder.WriteByte('e')
}

func (encoder *encoder) writeInterfaceType(v interface{}) {
	switch v.(type) {
	case string:
		encoder.writeString(v.(string))
	case []interface{}:
		encoder.writeList(v.([]interface{}))
	case map[string]interface{}:
		encoder.writeDictionary(v.(map[string]interface{}))
	case int:
		encoder.writeInt(int64(v.(int)))
	case int8:
		encoder.writeInt(int64(v.(int8)))
	case int16:
		encoder.writeInt(int64(v.(int16)))
	case int32:
		encoder.writeInt(int64(v.(int32)))
	case int64:
		encoder.writeInt(v.(int64))
	case uint8:
		encoder.writeUint(uint64(v.(uint8)))
	case uint16:
		encoder.writeUint(uint64(v.(uint16)))
	case uint32:
		encoder.writeUint(uint64(v.(uint32)))
	case uint64:
		encoder.writeUint(v.(uint64))
	}
}

func (encoder *encoder) writeList(list []interface{}) {
	encoder.WriteByte('l')
	for _, v := range list {
		encoder.writeInterfaceType(v)
	}
	encoder.WriteByte('e')
}

func (encoder *encoder) writeDictionary(dict map[string]interface{}) {
	encoder.WriteByte('d')
	for k, v := range dict {
		// Key
		encoder.writeString(k)

		// Value
		encoder.writeInterfaceType(v)
	}
	encoder.WriteByte('e')
}

// Encode takes a bencode dictionary and returns a
// bencode byte array representation of the dictionary
func Encode(dict interface{}) []byte {
	encoder := encoder{}
	encoder.writeDictionary(dict.(map[string]interface{}))
	return encoder.Bytes()
}
