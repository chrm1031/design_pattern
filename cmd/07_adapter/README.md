# Adapter

非互換なインターフェースのオブジェクト同士の協働を可能とする

![adapter](https://refactoring.guru/images/patterns/content/adapter/adapter-ja.png)

## 構造
![UML](https://refactoring.guru/images/patterns/diagrams/adapter/structure-object-adapter-indexed.png)

1. クライアント （Client） は、 既存のビジネス・ロジックを含んだクラスです。

2. クライアント・インターフェース （Client Interface） は、 クライアント側コードと協働するためには従わなければいけないプロトコル＝決まり事を記述しています。

3. サービス （Service） は、 外部から提供されるか、 昔から使われてきた、 ある役に立つクラスです。 このクラスは、 非互換のインターフェースのため、 クライアントから直接使うことができません。

4. アダプター （Adapter） は、 クライアントとサービスの両方と機能できるコードです： クライアント・インターフェースを実装すると同時にサービス・オブジェクトをラップします。 アダプターは、 アダプターのインターフェースを介してクライアントから呼び出され、 それをラップされたサービス・オブジェクトが理解できる形式に変換して呼び出します。

5. クライアント側コードは、 クライアント・インタフェースを介してアダプターとやりとりする限り、 具象アダプター・クラスと結合されることはありません。 このおかげで、 新しい種類のアダプターをプログラムに導入しても既存のクライアント側コードは問題なく動作します。 これは、 サービス・クラスのインターフェースを変更したり置き換えたりする際に便利です。 クライアント側コードを変更することなく、 新しいアダプター・クラスを作成することができます。

## 実装方法
1. 最低二つの非互換なクラスがあることを確認してください：

  - 有益だが変更不能なサービスクラス。 ​ （多くの場合、 外部提供だから、 過去から引き継がれたコードだから、 あるは、 依存する要素の数が多すぎるためといった理由で変更できない。）
  - サービス・クラスを利用したい 1 個以上の クライアント・クラス。

2. クライアント・インタフェースを宣言し、 クライアントとサービスとの情報伝達の方法を記述する。

3. クライアント・インターフェースに従うアダプター・クラスを作成。 メソッドはとりあえず空のままにする。

4. アダプター・クラスに、 サービス・オブジェクトへの参照を格納するためのフィールドを追加します。 このフィールドは、 コンストラクタで初期化するのが一般的な方法ですが、 アダプターのメソッドの呼び出し時に渡す方が便利な場合もあります。

5. クライアント・インタフェースの全メソッドをアダプター・クラスで一つずつ実装していってください。 アダプターは、 実際の仕事のほとんどはサービス・オブジェクトに委ね、 インターフェースやデータ形式の変換だけを行うようにします。

6. クライアントは、 クライアント・インタフェースを介してアダプターを使用する必要があります。 そうすることで、 クライアントのコードに影響を与えることなくアダプターの変更や拡張が可能となります。

## 他のパターンとの関係
- Bridge は通常、 アプリケーションの部分部分を独立して開発できるように、 設計当初から使われます。 一方、 Adapter は、 既存のアプリケーションに対して利用され、 本来は互換性のないクラスとうまく動作させるために使われます。

- Adapter は既存のオブジェクトのインターフェースを変更するのに対し、 Decorator はインターフェースの変更なしにオブジェクトを強力にします。 さらに、 Decorator は、 再帰的な合成をサポートします。 これは、 Adapter を使用する時には不可能です。

- Adapter はラップされたオブジェクトに対しては異なるインターフェースを提供し、 Proxy は同じインターフェースを提供し、 Decorator は強化したインターフェースを提供します。

- Facade が既存のオブジェクトに対して新しいインターフェースを定義するのに対し、 Adapter は既存のインターフェースを使えるようにしようとするものです。 Adapter は通常一つのオブジェクトだけを包み込みますが、 Facade は複数のオブジェクトからなるサブシステム全体を相手にします。

- Bridge、 State、 Strategy （と限られた意味合いでは、 Adapter も） は、 非常に似た構造をしています。 実際のところ、 これらの全てのパターンは、 合成に基づいており、 仕事を他のオブジェクトに委任します。 しかしながら、 違う問題を解決します。 パターンは、 単にコードを特定の方法で構造化するためのレシピではありません。 パターンが解決する問題に関して、 開発者同士がするコミュニケーションの道具でもあります。

## 引用元

> https://refactoring.guru/ja/design-patterns/adapter
> https://refactoring.guru/images/patterns/content/adapter/adapter-ja.png
> https://refactoring.guru/images/patterns/diagrams/adapter/structure-object-adapter-indexed.png