# Bridge
巨大なクラスや密接に関連したクラスの集まりを、抽象部分と実装部分という二つの階層に分離し、それぞれが独立して開発できるようにする

![bridge](https://refactoring.guru/images/patterns/content/bridge/bridge.png)

## 構造
![UML](https://refactoring.guru/images/patterns/diagrams/bridge/structure-ja-indexed.png)

1. 抽象化層 （Abstraction） は、 高レベルの制御ロジックを提供します。 実際の低レベルの作業は実装オブジェクトに任されています。

2. 実装 （Implementation） は、 すべての具象実装に共通のインターフェースを宣言します。 抽象化層は、 ここで宣言されたメソッドを介してのみ実装オブジェクトと通信することができます。抽象化は、 実装と同じメソッドを並べただけでもかまいません。 しかし通常、 抽象化は複雑な振る舞いを宣言します。 それらの振る舞いは、 実装によって宣言された多種多様な基本操作を使用します。

3. 具象的実装 （Concrete Implementation） には、 プラットフォーム固有のコードが含まれています。

4. 整った抽象化 （Refined Abstraction） は、 制御ロジックの変種を提供します。 親と同様に、 一般的な実装インターフェースを介して各種の異なる実装を使います。

5. 通常、 クライアント （Client） は、 抽象化層とだけやりとりをします。 しかし、 抽象化層のオブジェクトと実装オブジェクトを結びつけるのは、 クライアントの仕事です。

## 実装方法
1. クラス内の直交する次元を特定します。 これらの独立した概念は、 たとえば抽象化とプラットフォーム （OS） だったり、 ドメインとインフラストラクチャーだったり、 フロントエンドとバックエンドだったり、 インターフェースと実装だったりします。

2. クライアントが必要とする操作は何であるかを考え、 抽象化層の基底クラスで定義します。

3. すべてのプラットフォームで利用可能な操作は何かを決定します。 抽象化層が必要とするものを一般的な実装インターフェースで宣言します。

4. サポートするすべてのプラットフォームに対して具体的な実装クラスを作成します。 すべて実装クラスはプラットフォームの実装インターフェースに従うようにしてください。

5. 抽象化層のクラス内に、 実装型の参照フィールドを追加します。 抽象化層は、 このフィールドから参照される実装オブジェクトにほとんどの仕事を移譲します。

6. 高レベルのロジックにいくつかの変種がある場合は、 各変種ごとに抽象化層の基底クラスを拡張して、 より洗練した抽象化層クラスを作成します。

7. クライアント・コードは、 抽象化層のクラスのコンストラクターに実装オブジェクトを渡して、 互いを関連づけます。 その後、 クライアントは実装のことは忘れ、 抽象化層のオブジェクトとだけやりとりをします。

## 他のパターンとの関係
- Bridge は通常、 アプリケーションの部分部分を独立して開発できるように、 設計当初から使われます。 一方、 Adapter は、 既存のアプリケーションに対して利用され、 本来は互換性のないクラスとうまく動作させるために使われます。

- Bridge、 State、 Strategy （と限られた意味合いでは、 Adapter も） は、 非常に似た構造をしています。 実際のところ、 これらの全てのパターンは、 合成に基づいており、 仕事を他のオブジェクトに委任します。 しかしながら、 違う問題を解決します。 パターンは、 単にコードを特定の方法で構造化するためのレシピではありません。 パターンが解決する問題に関して、 開発者同士がするコミュニケーションの道具でもあります。

- Abstract Factory は、 Bridge と一緒に使用できます。 この組み合わせは、 Bridge によって定義された抽象化層のいくつかが特定の実装としか動作しない場合に便利です。 この場合、 Abstract Factory はこれらの関係をカプセル化し、 クライアント・コードから複雑さを隠すことができます。

- Builder と Bridge を組み合わせることができます： ディレクター・クラスは抽象化層の役割を果たし、 ビルダーは実装です。

## 引用元

> https://refactoring.guru/ja/design-patterns/bridge
> https://refactoring.guru/images/patterns/content/bridge/bridge.png
> https://refactoring.guru/images/patterns/diagrams/bridge/structure-ja-indexed.png