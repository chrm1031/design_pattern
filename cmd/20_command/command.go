package command

/*
コマンド・インターフェース
*/
type Command interface {
	execute()
}
