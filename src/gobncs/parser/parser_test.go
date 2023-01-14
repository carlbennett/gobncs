package gobncs

import (
	"reflect"
	"testing"
)

func Test_WriteMessage(t *testing.T) {
	type args struct {
		messageId   uint8
		messageData []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"message header standard output (SID_NULL)",
			args{0x00, []byte("")},
			[]byte("\xFF\x00\x04\x00"),
		},
		{
			"messageData with a uint32 and a string (SID_JOINCHANNEL)",
			args{0x0C, []byte("\x00\x00\x00\x02The Void\x00")},
			[]byte("\xFF\x0C\x11\x00\x00\x00\x00\x02The Void\x00"),
		},
		{
			"messageData larger than 255 bytes (SID_CHATCOMMAND)",
			args{0x0E, []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc velit libero, dictum id augue a, dictum malesuada magna. Sed in arcu a magna auctor tincidunt. Sed ut magna non metus hendrerit fringilla. Sed id nulla risus. Nunc malesuada, risus ut eleifend congue, ex erat fermentum ante, a aliquet nibh quam non augue. In hac habitasse platea dictumst. Sed congue ex non mauris tincidunt, id egestas sem dapibus. In hac habitasse platea dictumst. Sed ut metus euismod, posuere tellus eget, dictum ex.\x00")},
			[]byte("\xFF\x0E\xF8\x01Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc velit libero, dictum id augue a, dictum malesuada magna. Sed in arcu a magna auctor tincidunt. Sed ut magna non metus hendrerit fringilla. Sed id nulla risus. Nunc malesuada, risus ut eleifend congue, ex erat fermentum ante, a aliquet nibh quam non augue. In hac habitasse platea dictumst. Sed congue ex non mauris tincidunt, id egestas sem dapibus. In hac habitasse platea dictumst. Sed ut metus euismod, posuere tellus eget, dictum ex.\x00"),
		},
		{
			"messageId is not always zero with empty messageData (SID_LEAVECHAT)",
			args{0x10, []byte("")},
			[]byte("\xFF\x10\x04\x00"),
		},
		{
			"messageData with a uint32 (SID_PING)",
			args{0x25, []byte("\xBA\xAD\xF0\x0D")},
			[]byte("\xFF\x25\x08\x00\xBA\xAD\xF0\x0D"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WriteMessage(tt.args.messageId, tt.args.messageData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WriteMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
