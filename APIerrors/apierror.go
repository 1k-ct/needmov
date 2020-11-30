package apierrors

//ApplicationError エラーtype
type ApplicationError struct {
	Code  int    `json:"code"`
	Level string `json:"level"`
	Msg1  string `json:"msg1"` //英語で書く
	Msg2  string `json:"msg2"` //日本語で書く
}

//ErrInvalidURL 無効なURLのとき。UCで始まり２４文字ではない
var ErrInvalidURL = &ApplicationError{
	Code:  1,
	Level: "Error",
	Msg1:  "This CHANNEL URL is not valid.The length of the string is not 24 characters.Or it is a string that does not begin with 'UC'.",
	Msg2:  "このCHANNEL URLは無効です。文字列の長さが２４文字ではありません。または、”UC”　で始まらない文字列です。",
}

//ErrDuplicateURL URLが同じ場合
var ErrDuplicateURL = &ApplicationError{
	Code:  2,
	Level: "Error",
	Msg1:  "That URL is a duplicate.It's already in the DB.",
	Msg2:  "そのURLは重複する。すでにDBにあります。",
}

//ErrDB データベース全般のエラーについて
var ErrDB = &ApplicationError{
	Code:  3,
	Level: "Error",
	Msg1:  "I couldn't manipulate the database.",
	Msg2:  "データベースの操作が出来ませんでした。",
}
var ErrJson = &ApplicationError{
	Code:  4,
	Level: "Error",
	Msg1:  "I couldn't read the json.",
	Msg2:  "jsonの読み取りか出来ませんでした。",
}
var ErrDecoding = &ApplicationError{
	Code:  5,
	Level: "Error",
	Msg1:  "That character could not be decoded(Base64)",
	Msg2:  "その文字は、変換(Base64)出来ませんでした。",
}
var ErrAPIKey = &ApplicationError{
	Code:  6,
	Level: "Error",
	Msg1:  "Its API key is not registered.",
	Msg2:  "そのAPIキーは、登録されていません。",
}
