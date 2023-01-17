package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct{}

var singleInstance *single

/*
注目すべき点：

初回は single­Instance であることを確認するための nil の検査が冒頭にあります。
これは、 getinstance メソッドの呼び出しのたびに高価なロック操作を行うことを避けるためです。
この検査が陰性ならばそれは single­Instance フィールドはすでに値を持っているということです。

single­Instance 構造体はロックの範囲内で生成されます。

ロック獲得後にもう一度 nil 検査があります。
これは二つ以上のゴルーチンが最初の検査をすり抜けた場合に、ゴルーチン一つだけがシングルトンのインスタンスを作成できるようにするためです。
そしないと全部のゴルーチンがそれぞれ専用のシングルトン構造体を作成してしまいます。
*/
func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}
	return singleInstance
}
