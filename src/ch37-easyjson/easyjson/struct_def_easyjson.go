// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package jsontest

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson7c82d03DecodeCh37EasyjsonEasyjson(in *jlexer.Lexer, out *JobInfo) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "skills":
			if in.IsNull() {
				in.Skip()
				out.Skills = nil
			} else {
				in.Delim('[')
				if out.Skills == nil {
					if !in.IsDelim(']') {
						out.Skills = make([]string, 0, 4)
					} else {
						out.Skills = []string{}
					}
				} else {
					out.Skills = (out.Skills)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Skills = append(out.Skills, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7c82d03EncodeCh37EasyjsonEasyjson(out *jwriter.Writer, in JobInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"skills\":"
		out.RawString(prefix[1:])
		if in.Skills == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Skills {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v JobInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7c82d03EncodeCh37EasyjsonEasyjson(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v JobInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7c82d03EncodeCh37EasyjsonEasyjson(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *JobInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7c82d03DecodeCh37EasyjsonEasyjson(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *JobInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7c82d03DecodeCh37EasyjsonEasyjson(l, v)
}
func easyjson7c82d03DecodeCh37EasyjsonEasyjson1(in *jlexer.Lexer, out *Employee) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "basic_info":
			(out.BasicInfo).UnmarshalEasyJSON(in)
		case "job_info":
			(out.JobInfo).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7c82d03EncodeCh37EasyjsonEasyjson1(out *jwriter.Writer, in Employee) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"basic_info\":"
		out.RawString(prefix[1:])
		(in.BasicInfo).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"job_info\":"
		out.RawString(prefix)
		(in.JobInfo).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Employee) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7c82d03EncodeCh37EasyjsonEasyjson1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Employee) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7c82d03EncodeCh37EasyjsonEasyjson1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Employee) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7c82d03DecodeCh37EasyjsonEasyjson1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Employee) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7c82d03DecodeCh37EasyjsonEasyjson1(l, v)
}
func easyjson7c82d03DecodeCh37EasyjsonEasyjson2(in *jlexer.Lexer, out *BasicInfo) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "age":
			out.Age = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7c82d03EncodeCh37EasyjsonEasyjson2(out *jwriter.Writer, in BasicInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"age\":"
		out.RawString(prefix)
		out.Int(int(in.Age))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BasicInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7c82d03EncodeCh37EasyjsonEasyjson2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BasicInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7c82d03EncodeCh37EasyjsonEasyjson2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BasicInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7c82d03DecodeCh37EasyjsonEasyjson2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BasicInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7c82d03DecodeCh37EasyjsonEasyjson2(l, v)
}