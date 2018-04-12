// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package domain

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

func easyjsonB6915918DecodeGithubComKamilskClickDomain(in *jlexer.Lexer, out *Metadata) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "cookie":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Cookie = make(map[string]string)
				} else {
					out.Cookie = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 string
					v1 = string(in.String())
					(out.Cookie)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "header":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Header = make(map[string][]string)
				} else {
					out.Header = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 []string
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						in.Delim('[')
						if v2 == nil {
							if !in.IsDelim(']') {
								v2 = make([]string, 0, 4)
							} else {
								v2 = []string{}
							}
						} else {
							v2 = (v2)[:0]
						}
						for !in.IsDelim(']') {
							var v3 string
							v3 = string(in.String())
							v2 = append(v2, v3)
							in.WantComma()
						}
						in.Delim(']')
					}
					(out.Header)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
			}
		case "query":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Query = make(map[string][]string)
				} else {
					out.Query = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v4 []string
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						in.Delim('[')
						if v4 == nil {
							if !in.IsDelim(']') {
								v4 = make([]string, 0, 4)
							} else {
								v4 = []string{}
							}
						} else {
							v4 = (v4)[:0]
						}
						for !in.IsDelim(']') {
							var v5 string
							v5 = string(in.String())
							v4 = append(v4, v5)
							in.WantComma()
						}
						in.Delim(']')
					}
					(out.Query)[key] = v4
					in.WantComma()
				}
				in.Delim('}')
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
func easyjsonB6915918EncodeGithubComKamilskClickDomain(out *jwriter.Writer, in Metadata) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Cookie) != 0 {
		const prefix string = ",\"cookie\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v6First := true
			for v6Name, v6Value := range in.Cookie {
				if v6First {
					v6First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v6Name))
				out.RawByte(':')
				out.String(string(v6Value))
			}
			out.RawByte('}')
		}
	}
	if len(in.Header) != 0 {
		const prefix string = ",\"header\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v7First := true
			for v7Name, v7Value := range in.Header {
				if v7First {
					v7First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v7Name))
				out.RawByte(':')
				if v7Value == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
					out.RawString("null")
				} else {
					out.RawByte('[')
					for v8, v9 := range v7Value {
						if v8 > 0 {
							out.RawByte(',')
						}
						out.String(string(v9))
					}
					out.RawByte(']')
				}
			}
			out.RawByte('}')
		}
	}
	if len(in.Query) != 0 {
		const prefix string = ",\"query\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v10First := true
			for v10Name, v10Value := range in.Query {
				if v10First {
					v10First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v10Name))
				out.RawByte(':')
				if v10Value == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
					out.RawString("null")
				} else {
					out.RawByte('[')
					for v11, v12 := range v10Value {
						if v11 > 0 {
							out.RawByte(',')
						}
						out.String(string(v12))
					}
					out.RawByte(']')
				}
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Metadata) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB6915918EncodeGithubComKamilskClickDomain(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Metadata) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB6915918EncodeGithubComKamilskClickDomain(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Metadata) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB6915918DecodeGithubComKamilskClickDomain(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Metadata) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB6915918DecodeGithubComKamilskClickDomain(l, v)
}
func easyjsonB6915918DecodeGithubComKamilskClickDomain1(in *jlexer.Lexer, out *Log) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint64(in.Uint64())
		case "link_id":
			out.LinkID = string(in.String())
		case "alias_id":
			out.AliasID = uint64(in.Uint64())
		case "target_id":
			out.TargetID = uint64(in.Uint64())
		case "uri":
			out.URI = string(in.String())
		case "code":
			out.Code = int(in.Int())
		case "context":
			(out.Context).UnmarshalEasyJSON(in)
		case "created_at":
			out.CreatedAt = string(in.String())
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
func easyjsonB6915918EncodeGithubComKamilskClickDomain1(out *jwriter.Writer, in Log) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"link_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LinkID))
	}
	{
		const prefix string = ",\"alias_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.AliasID))
	}
	{
		const prefix string = ",\"target_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.TargetID))
	}
	{
		const prefix string = ",\"uri\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.URI))
	}
	{
		const prefix string = ",\"code\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Code))
	}
	{
		const prefix string = ",\"context\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Context).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"created_at\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CreatedAt))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Log) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB6915918EncodeGithubComKamilskClickDomain1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Log) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB6915918EncodeGithubComKamilskClickDomain1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Log) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB6915918DecodeGithubComKamilskClickDomain1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Log) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB6915918DecodeGithubComKamilskClickDomain1(l, v)
}
