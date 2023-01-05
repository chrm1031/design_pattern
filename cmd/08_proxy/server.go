package proxy

/*
サーバー
*/
type server interface {
	handleRequest(string, string) (int, string)
}
