package iterator

import "fmt"

/*
クライアント・コード
*/
func main() {
	user1 := &User{
		name: "a",
		age:  30,
	}

	user2 := &User{
		name: "b",
		age:  20,
	}

	userCollection := &UserCollection{
		users: []*User{user1, user2},
	}

	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}

/*
実行結果
User is &{name:a age:30}
User is &{name:b age:20}
*/
