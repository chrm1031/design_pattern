package prototype

/*
プロトタイプ・インターフェース
*/
type Inode interface {
	print(string)
	clone() Inode
}
