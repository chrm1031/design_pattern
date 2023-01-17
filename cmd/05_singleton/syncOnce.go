package singleton

/*
init 関数の内部で単一のインスタンスを作成することが可能です。
これはインスタンスの早期初期化が許されている場合のみ適用できます。
init 関数は、パッケージ中のファイルあたり 1 回だけ呼ばれます。
ですのでただ一つのインスタンスが作成されることを保証できます。
*/
// var once sync.Once

// type single struct {
// }

// var singleInstance *single

// func getInstance() *single {
// 	if singleInstance == nil {
// 		once.Do(
// 			func() {
// 				fmt.Println("Creating single instance now.")
// 				singleInstance = &single{}
// 			})
// 	} else {
// 		fmt.Println("Single instance already created.")
// 	}

// 	return singleInstance
// }
