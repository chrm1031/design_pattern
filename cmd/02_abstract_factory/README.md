# Abstract Factory

関連したオブジェクトの集りを、 具象クラスを指定することなく生成することを可能とする

![abstract factory](https://refactoring.guru/images/patterns/content/abstract-factory/abstract-factory-ja.png)


## 構造
![UML](https://refactoring.guru/images/patterns/diagrams/abstract-factory/structure-indexed.png)

1. 抽象製品 （Abstract Product） は、 製品ファミリーを構成する個別の関連製品に対するインターフェースを宣言します。

2. 具象製品 （Concrete Product） は、 抽象製品の種々の変種ごとにグループ化されています。 個々の抽象製品 （椅子とかソファー） は、 全部の変種 （ビクトリア調、 現代風） において実装されている必要があります。

3. 抽象ファクトリー （Abstract Factory） インターフェースは、 抽象製品のそれぞれを生成するメソッドの集合です。

4. 具象ファクトリー （Concrete Factory） は、 抽象ファクトリーの生成メソッドを実装します。 個々の具象ファクトリーは、 特定の異種 （様式） に対応しており、 そのような異種の製品のみを作成します。

5. 具象ファクトリーは、 具象製品をインスタンス化しますが、 生成メソッドのシグネチャーは、 抽象 製品である必要があります。 そうすることで、 クライアント側コードは、 あるファクトリーから返される特定の変種の製品に縛られることがなくなります。 抽象インターフェースによって通信する限り、 クライアント （Client） は、 いかなる具象ファクトリーや製品変種も扱うことができます。

## 実装方法
1. 明確に区別できる製品の型と、 製品型の変種の表を作ります。

2. 全部の製品型について、 抽象製品インターフェースを宣言します。 そして、 全インターフェースを実装した製品の具象クラスを作ります。

3. 全抽象製品型に対する生成メソッドを含んだ抽象ファクトリー・インターフェースを宣言します。

4. 各変種ごとに、 具象ファクトリー・クラスを実装します。

5. アプリのどこかに、 ファクトリー初期化コードを追加します。 そこでアプリケーションの構成か環境設定に従って具象ファクトリー・クラスのインスタンスを作成します。 製品を構築するすべてのクラスに、 このファクトリー・オブジェクトを渡します。

6. コードをスキャンして、 製品クラスのコンストラクターへの直接の呼び出しを探します。 これらの呼び出しを、 ファクトリー・オブジェクトに対する適切な作成メソッドの呼び出しと置き換えます。

## 他のパターンとの関係
- 多くの設計は、 まず比較的単純でサブクラスによりカスタマイズ可能な、 Factory Method から始まり、 次第に、 もっと柔軟だが複雑な Abstract Factory や Prototype や Builder へと発展していきます。

- Builder は、 複雑なオブジェクトを段階的に構築することに重点を置いています。 Abstract Factory は、 関連するオブジェクトの集団を作成することに特化しています。 Abstract Factory がすぐにプロダクトを返すのに対して、 Builder ではプロダクトの取得前に、 いくつかの追加の構築のステップを踏まなければなりません。

- Abstract Factory クラスは、 多くの場合 Factory Methods の集まりですが、 Prototype を使ってメソッドを書くこともできます。

- サブシステムがオブジェクトを作成する方法をクライアントから隠蔽することだけが目的なら、 Abstract Factory を Facade の代わりに使えます。

- Abstract Factory は、 Bridge と一緒に使用できます。 この組み合わせは、 Bridge によって定義された抽象化層のいくつかが特定の実装としか動作しない場合に便利です。 この場合、 Abstract Factory はこれらの関係をカプセル化し、 クライアント・コードから複雑さを隠すことができます。

- Abstract Factories、 Builders、 Prototypes はどれも Singletons で実装可能です。

## 引用元

> https://refactoring.guru/ja/design-patterns/abstract-factory
> https://refactoring.guru/images/patterns/content/abstract-factory/abstract-factory-ja.png
> https://refactoring.guru/images/patterns/diagrams/abstract-factory/structure-indexed.png