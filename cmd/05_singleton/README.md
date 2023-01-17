# Singleton
Singleton （シングルトン）は、クラスが一つのインスタンスのみを持つことを保証するとともに、 このインスタンスへの大域アクセス・ポイントを提供する

![singleton](https://refactoring.guru/images/patterns/content/singleton/singleton.png)

## 構造
![UML](https://refactoring.guru/images/patterns/diagrams/singleton/structure-ja-indexed.png)

1. シングルトン （Singleton） クラスは、 そのクラスの同じインスタンスを返す静的メソッド get­Instance を宣言します。シングルトンのコンストラクターは、 クライアント・コードから隠れている必要があります。 get­Instance メソッドの呼び出しが、 シングルトンのオブジェクトを取得する唯一の方法でなければなりません。

## 実装方法
1. クラスにシングルトンのインスタンスを格納する非公開 （private） の静的 （static） フィールドを追加します。

2. シングルトンのインスタンスを取得するための公開 （public） 静的 （static） 作成用メソッドを宣言します。

3. 静的メソッド内に、 遅延初期化コードを実装します。 初回の呼び出し時だけ新規オブジェクトを作成し、 それを静的フィールドに格納します。 その後の呼び出しでは、 メソッドは常にそのインスタンスを返すようにします。

4. コンストラクターを非公開 （private） とします。 クラス中の静的メソッドからはこのコンストラクターを呼べますが、 他のオブジェクトからは呼べません。

5. クライアントのコード中のコンストラクター呼び出しのすべてを静的生成用メソッドの呼び出しと置き換えます。

## 他のパターンとの関係
- 多くの場合、 ファサード・オブジェクトは一つだけあれば十分なので、 Facade はしばしば Singleton に変換可能です。

- Flyweight で、 共有状態の全部を一つのフライウェイト・オブジェクトに何らかの方法で削減できた場合、 それは Singleton に似たものになります。 しかし、 この二つのパターンには、 根本的な違いが二箇所あります。

  1. Singleton のインスタンスは一つだけですが、 Flyweight クラスは、 異なる内因的状態を持つ複数のインスタンスがある可能性があります。
  2. 2Singleton オブジェクトは変更可能かもしれませんが、 Flyweight のオブジェクトは不変です。

- Abstract Factories、 Builders、 Prototypes はどれも Singletons で実装可能です。

## 引用元

> https://refactoring.guru/ja/design-patterns/singleton
> https://refactoring.guru/images/patterns/content/singleton/singleton.png
> https://refactoring.guru/images/patterns/diagrams/singleton/structure-ja-indexed.png