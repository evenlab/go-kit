// Copyright © 2020-2022 The EVEN Solutions Developers Team

package timestamp_test

import (
	"reflect"
	"testing"
	"time"

	json "github.com/json-iterator/go"
	"google.golang.org/protobuf/proto"

	"github.com/evenlab/go-kit/timestamp"
	"github.com/evenlab/go-kit/timestamp/proto/pb"
)

const (
	testTimestampLayout = "1970-12-31T23:23:59.123456789Z+0000UTC"
)

func Benchmark_Now(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = timestamp.Now()
	}
}

func Benchmark_DecodeTimestamp(b *testing.B) {
	ts := timestamp.Now()
	pbuf, _ := ts.Encode()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := timestamp.DecodeTimestamp(pbuf); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Timestamp_Decode(b *testing.B) {
	ts := timestamp.Now()
	pbuf, _ := ts.Encode()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ts := timestamp.Timestamp{}
		if err := ts.Decode(pbuf); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Timestamp_Encode(b *testing.B) {
	ts := timestamp.Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := ts.Encode(); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Timestamp_Marshal(b *testing.B) {
	ts := timestamp.Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := ts.Marshal(); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Timestamp_MarshalJSON(b *testing.B) {
	ts := timestamp.Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := ts.MarshalJSON(); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Timestamp_Parse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ts := timestamp.Timestamp{}
		if err := ts.Parse(testTimestampLayout); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Timestamp_Pretty(b *testing.B) {
	ts := timestamp.Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ts.Pretty()
	}
}

func Benchmark_Timestamp_String(b *testing.B) {
	ts := timestamp.Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ts.String()
	}
}

func Benchmark_Timestamp_UnixNanoStr(b *testing.B) {
	ts := timestamp.Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ts.UnixNanoStr()
	}
}

func Benchmark_Timestamp_Unmarshal(b *testing.B) {
	ts := timestamp.Now()
	blob, _ := ts.Marshal()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ts := timestamp.Timestamp{}
		if err := ts.Unmarshal(blob); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Timestamp_UnmarshalJSON(b *testing.B) {
	ts := timestamp.Now()
	blob, _ := ts.MarshalJSON()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ts := timestamp.Timestamp{}
		if err := ts.UnmarshalJSON(blob); err != nil {
			b.Fatal(err)
		}
	}
}

func Test_Now(t *testing.T) {
	t.Parallel()

	zone, offset := time.Now().UTC().Zone()

	tests := [1]struct {
		name       string
		wantZone   string
		wantOffset int
	}{
		{
			name:       "UTC",
			wantZone:   zone,
			wantOffset: offset,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			zone, offset := timestamp.Now().Zone()
			if zone != test.wantZone {
				t.Errorf("Now() zone: %#v | want: %#v", zone, test.wantZone)
			}
			if offset != test.wantOffset {
				t.Errorf("Now() offset: %#v | want: %#v", offset, test.wantOffset)
			}
		})
	}
}

func Test_DecodeTimestamp(t *testing.T) {
	t.Parallel()

	ts := timestamp.Now()
	blob, _ := ts.MarshalBinary()

	tests := [4]struct {
		name    string
		pbuf    *pb.Timestamp
		want    timestamp.Timestamp
		wantErr bool
	}{
		{
			name: "OK",
			pbuf: &pb.Timestamp{Blob: blob},
			want: ts,
		},
		{
			name: "version_Unsupported_ERR",
			pbuf: func() *pb.Timestamp {
				blob := make([]byte, 16)
				// byte on index=0 position encodes version
				blob[0] = 15 // 15 is unsupported version
				return &pb.Timestamp{Blob: blob}
			}(),
			wantErr: true,
		},
		{
			name:    "empty_BLOB_ERR",
			pbuf:    &pb.Timestamp{Blob: make([]byte, 0)},
			wantErr: true,
		},
		{
			name:    "invalid_Len_ERR",
			pbuf:    &pb.Timestamp{Blob: make([]byte, 3)},
			wantErr: true,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := timestamp.DecodeTimestamp(test.pbuf)
			if (err != nil) != test.wantErr {
				t.Errorf("DecodeTimestamp() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("DecodeTimestamp() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_Timestamp_Decode(t *testing.T) {
	t.Parallel()

	ts := timestamp.Now()
	blob, _ := ts.MarshalBinary()

	tests := [4]struct {
		name    string
		pbuf    *pb.Timestamp
		want    timestamp.Timestamp
		wantErr bool
	}{
		{
			name: "OK",
			pbuf: &pb.Timestamp{Blob: blob},
			want: ts,
		},
		{
			name: "version_Unsupported_ERR",
			pbuf: func() *pb.Timestamp {
				blob := make([]byte, 16)
				// byte on index=0 position encodes version
				blob[0] = 15 // 15 is unsupported version
				return &pb.Timestamp{Blob: blob}
			}(),
			wantErr: true,
		},
		{
			name:    "empty_BLOB_ERR",
			pbuf:    &pb.Timestamp{Blob: make([]byte, 0)},
			wantErr: true,
		},
		{
			name:    "invalid_Len_ERR",
			pbuf:    &pb.Timestamp{Blob: make([]byte, 3)},
			wantErr: true,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := timestamp.Timestamp{}
			err := got.Decode(test.pbuf)
			if (err != nil) != test.wantErr {
				t.Errorf("Decode() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Decode() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_Timestamp_Encode(t *testing.T) {
	t.Parallel()

	ts := timestamp.Now()
	blob, _ := ts.Time.MarshalBinary()

	tests := [2]struct {
		name    string
		time    timestamp.Timestamp
		want    *pb.Timestamp
		wantErr bool
	}{
		{
			name: "OK",
			time: ts,
			want: &pb.Timestamp{Blob: blob},
		},
		{
			name:    "unexpected_Zone_Offset_ERR",
			time:    timestamp.Timestamp{Time: ts.In(time.FixedZone("unexpected zone offset", -60))},
			wantErr: true,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := test.time.Encode()
			if (err != nil) != test.wantErr {
				t.Errorf("Encode() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Encode() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_Timestamp_Marshal(t *testing.T) {
	t.Parallel()

	ts := timestamp.Now()
	pbuf, _ := ts.Encode()
	want, _ := proto.Marshal(pbuf)

	tests := [2]struct {
		name    string
		time    timestamp.Timestamp
		want    []byte
		wantErr bool
	}{
		{
			name: "OK",
			time: ts,
			want: want,
		},
		{
			name:    "unexpected_Zone_Offset_ERR",
			time:    timestamp.Timestamp{Time: ts.In(time.FixedZone("unexpected zone offset", -60))},
			wantErr: true,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := test.time.Marshal()
			if (err != nil) != test.wantErr {
				t.Errorf("Marshal() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Marshal() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_Timestamp_MarshalJSON(t *testing.T) {
	t.Parallel()

	ts := timestamp.Now()
	pbuf, _ := ts.Encode()
	want, _ := json.Marshal(pbuf)

	tests := [2]struct {
		name    string
		time    timestamp.Timestamp
		want    []byte
		wantErr bool
	}{
		{
			name: "OK",
			time: ts,
			want: want,
		},
		{
			name:    "unexpected_Zone_Offset_ERR",
			time:    timestamp.Timestamp{Time: ts.In(time.FixedZone("unexpected zone offset", -60))},
			wantErr: true,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := test.time.MarshalJSON()
			if (err != nil) != test.wantErr {
				t.Errorf("MarshalJSON() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("MarshalJSON() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_Timestamp_Parse(t *testing.T) {
	t.Parallel()

	tests := [2]struct {
		name    string
		text    string
		want    string
		wantErr bool
	}{
		{
			name: "OK",
			text: testTimestampLayout,
			want: testTimestampLayout,
		},
		{
			name:    "ERR",
			text:    "0000-00-00T00:00:00.000000000Z-0000UTC",
			wantErr: true,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ts := timestamp.Timestamp{}
			if err := ts.Parse(test.text); (err != nil) != test.wantErr {
				t.Errorf("Parse() error: %v | want: %v", err, test.wantErr)
				return
			}
			if test.wantErr {
				return
			}
			if got := ts.Format(timestamp.LayoutTimestamp); got != test.want {
				t.Errorf("Parse() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_Timestamp_Pretty(t *testing.T) {
	t.Parallel()

	ts := timestamp.Timestamp{}
	if err := ts.Parse(testTimestampLayout); err != nil {
		t.Errorf("Parse() error: %v", err)
		return
	}

	tests := [1]struct {
		name string
		time timestamp.Timestamp
		want string
	}{
		{
			name: "OK",
			time: ts,
			want: "Thu Dec 31 23:23:59 +0000 1970",
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.time.Pretty(); got != test.want {
				t.Errorf("Pretty() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_Timestamp_String(t *testing.T) {
	t.Parallel()

	ts := timestamp.Timestamp{}
	if err := ts.Parse(testTimestampLayout); err != nil {
		t.Errorf("Parse() error: %v", err)
		return
	}

	tests := [1]struct {
		name string
		time timestamp.Timestamp
		want string
	}{
		{
			name: "OK",
			time: ts,
			want: testTimestampLayout,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.time.String(); got != test.want {
				t.Errorf("String() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_Timestamp_UnixNanoStr(t *testing.T) {
	t.Parallel()

	ts := timestamp.Timestamp{}
	if err := ts.Parse(testTimestampLayout); err != nil {
		t.Errorf("Parse() error: %v", err)
		return
	}

	tests := [1]struct {
		name string
		time timestamp.Timestamp
		want string
	}{
		{
			name: "OK",
			time: ts,
			want: "31533839123456789",
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.time.UnixNanoStr(); got != test.want {
				t.Errorf("UnixNanoStr() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_Timestamp_Unmarshal(t *testing.T) {
	t.Parallel()

	ts := timestamp.Now()
	blob, _ := ts.Marshal()

	tests := [2]struct {
		name    string
		blob    []byte
		want    timestamp.Timestamp
		wantErr bool
	}{
		{
			name: "OK",
			blob: blob,
			want: ts,
		},
		{
			name:    "invalid_BLOB_ERR",
			blob:    []byte(":"), // invalid data
			wantErr: true,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := timestamp.Timestamp{}
			if err := got.Unmarshal(test.blob); (err != nil) != test.wantErr {
				t.Errorf("Unmarshal() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Unmarshal() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_Timestamp_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	ts := timestamp.Now()
	blob, _ := ts.MarshalJSON()

	tests := [2]struct {
		name    string
		blob    []byte
		want    timestamp.Timestamp
		wantErr bool
	}{
		{
			name: "OK",
			blob: blob,
			want: ts,
		},
		{
			name:    "invalid_JSON_ERR",
			blob:    []byte(":"), // invalid json
			wantErr: true,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := timestamp.Timestamp{}
			if err := got.UnmarshalJSON(test.blob); (err != nil) != test.wantErr {
				t.Errorf("UnmarshalJSON() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("UnmarshalJSON() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}
