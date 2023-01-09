# Template Pattern
スーパークラス内でアルゴリズムの骨格を定義しておき、サブクラスは構造を変えることなくアルゴリズムの特定のステップを上書きする

![template_pattarn](https://refactoring.guru/images/patterns/content/template-method/template-method.png)

## 構造
![UML](https://refactoring.guru/images/patterns/diagrams/template-method/structure-indexed.png)

1. 抽象クラス （Abstract Class） は、 アルゴリズムの各ステップに対応するメソッドと、 これらのメソッドを特定の順序で呼び出す実際のテンプレート・メソッドを宣言します。 ステップは、 abstract と宣言されるか、 何らかのデフォルトの実装を持つかのどちらかです。

2. 具象クラス （Concrete Class） は、 テンプレート・メソッドを除くすべてのステップを上書きできます。

## 実装方法
1. 対象となるアルゴリズムを分析し、 それをステップに分割できるかどうかを検討します。 どのステップがサブクラス間で共通で、 どのステップが常にサブクラスに固有かを考察します。

2. 抽象基底クラスを作成し、 テンプレート・メソッドと、 アルゴリズムの各ステップに対応する抽象メソッドの組を宣言します。 アルゴリズムの輪郭は、 ステップを実行するテンプレート・メソッドに描かれています。 テンプレート・メソッドがサブクラスで上書きされるのを防ぐために、 テンプレート メソッドを final と宣言することを検討してください。

3. 全部のステップが抽象メソッドとなってしまったとしても、 大丈夫です。 いくつかのメソッドにはデフォルトの実装があった方が有益かもしれません。 サブクラスがそれらのメソッドを実装する必要がなくなります。

4. アルゴリズムの重要なステップの間にフックを追加することを検討してください。

5. アルゴリズムの異種ごとに、 新しい具象サブクラスを作成します。 サブクラスは、 抽象ステップのすべてを実装しなければならないですが、 任意のステップは上書きしてもよいです。


## 他のパターンとの関係
- Factory Method は、 Template Method の特別な場合です。 同時に、 Factory Method は、 大きな Template Method の一つのステップとして使うこともできます。

- Template Method は、 継承に基づくもので、 サブクラスでその一部を拡張することによって、 アルゴリズムを部分的に変更できます。 Strategy は、 合成に基き、 オブジェクトの振る舞いを、 新しい振る舞いに関連した異なるストラテジーを与えることにより変更します。 Template Method は、 クラスのレベルで機能するので、 静的です。 Strategy はオブジェクトのレベルで機能するため、 実行時に動作を切り替えることができます。

## 引用元

> https://refactoring.guru/ja/design-patterns/template-method
> https://refactoring.guru/images/patterns/content/template-method/template-method.png
> https://refactoring.guru/images/patterns/diagrams/template-method/structure-indexed.png