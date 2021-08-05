package utils

func StrPtr2Str(in *string) string {
	if in == nil {
		return ""
	}
	return *in
}
