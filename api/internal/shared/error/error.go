package error

import "fmt"

const (
	NotFound = iota
	InvalidArgument
	Custom
	Unknown
	Database
)

type errorOption struct {
	type_    int
	property string
}

func (o *errorOption) Property(property string) *errorOption {
	o.property = property
	return o
}

func (o *errorOption) Build() *Error {
	return &Error{option: *o}
}

type Error struct {
	option errorOption
}

func ErrorBuilder(type_ int) *errorOption {
	return &errorOption{type_: type_}
}

func (e Error) Error() string {
	switch e.option.type_ {
	case NotFound:
		return "データが見つかりません"
	case InvalidArgument:
		if e.option.property != "" {
			return fmt.Sprintf("%sが不正です", e.option.property)
		}
		return "パラメーターが不正です"
	case Unknown:
		return "予期せぬエラーが発生しました"
	case Database:
		return "データベースの処理に失敗しました"
	default:
		return "予期せぬエラーが発生しました"
	}
}
