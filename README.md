# Feed Snippet

The feedsnippet is a tool for displaying the latest feed snippet in README.md, and so on.
It creates a markdown feed snippet from an RSS, ATOM, or JSON feeds and replace the old snippet with the new one.
The format of the generated feed can be formatted using the text template.

# Install

```
go install github.com/ikawaha/feedsnippet@latest
```

# Config

## Minimal configuration

```yaml
- urls:
    - https://zenn.dev/ikawaha/feed
  limit: 5
- urls:
    - https://qiita.com/ikawaha/feed
  limit: 5
```

Outputs:

* [その文字が JIS X 0208 に含まれるか？ あるいは unicode.RangeTable の使い方](https://zenn.dev/ikawaha/articles/20210116-ab1ac4a692ae8bb4d9cf)
* [Go製全文検索エンジンBlugeで日本語形態素解析をおこなう](https://zenn.dev/ikawaha/articles/20201230-84b042603ccbbce645d5)
* [実践：形態素解析 kagome v2](https://zenn.dev/ikawaha/books/kagome-v2-japanese-tokenizer)
* [形態素解析器 kagome を Google App Engine で動かす](https://zenn.dev/ikawaha/articles/hatena-20161004-221708)
* [形態素解析器 kagome を brew tap でインストールできるようにした](https://zenn.dev/ikawaha/articles/20201029-391c049a13fb8506361d)
* [Qiitaの記事をZenn形式のMarkdownで保存して管理する](https://qiita.com/ikawaha/items/ab9906581e34f26993a9)
* [Goa v3 のテストをシュッとする](https://qiita.com/ikawaha/items/e0c2b3ed0fedb12f4847)
* [人生で何度目かのダブル配列TRIEを書いた](https://qiita.com/ikawaha/items/edb4e18960ae6e4babc3)
* [形態素解析器 kagome のユーザー辞書の使い方](https://qiita.com/ikawaha/items/9ebe3e1104fb80706d99)
* [goa でデザイン・ファーストをシュッとする](https://qiita.com/ikawaha/items/6638ee8b6978aef50d65)

## Listing multiple feeds

```yaml
- urls:
    - https://zenn.dev/ikawaha/feed
  template: |-
    **Zenn**
    {{range . -}}
      * ![](./icon/zenn.png) [{{ .Title }}]({{ .Link }})
    {{ end }}
  limit: 5
- urls:
    - https://qiita.com/ikawaha/feed
  template: |-
    **Qiita**
    {{range . -}}
      * ![](./icon/qiita.png) [{{ .Title }}]({{ .Link }})
    {{ end }}
  limit: 5
```

Outputs:

#### Zenn
* ![](./icon/zenn.png) [その文字が JIS X 0208 に含まれるか？ あるいは unicode.RangeTable の使い方](https://zenn.dev/ikawaha/articles/20210116-ab1ac4a692ae8bb4d9cf)
* ![](./icon/zenn.png) [Go製全文検索エンジンBlugeで日本語形態素解析をおこなう](https://zenn.dev/ikawaha/articles/20201230-84b042603ccbbce645d5)
* ![](./icon/zenn.png) [実践：形態素解析 kagome v2](https://zenn.dev/ikawaha/books/kagome-v2-japanese-tokenizer)
* ![](./icon/zenn.png) [形態素解析器 kagome を Google App Engine で動かす](https://zenn.dev/ikawaha/articles/hatena-20161004-221708)
* ![](./icon/zenn.png) [形態素解析器 kagome を brew tap でインストールできるようにした](https://zenn.dev/ikawaha/articles/20201029-391c049a13fb8506361d)
#### Qiita
* ![](./icon/qiita.png) [Qiitaの記事をZenn形式のMarkdownで保存して管理する](https://qiita.com/ikawaha/items/ab9906581e34f26993a9)
* ![](./icon/qiita.png) [Goa v3 のテストをシュッとする](https://qiita.com/ikawaha/items/e0c2b3ed0fedb12f4847)
* ![](./icon/qiita.png) [人生で何度目かのダブル配列TRIEを書いた](https://qiita.com/ikawaha/items/edb4e18960ae6e4babc3)
* ![](./icon/qiita.png) [形態素解析器 kagome のユーザー辞書の使い方](https://qiita.com/ikawaha/items/9ebe3e1104fb80706d99)
* ![](./icon/qiita.png) [goa でデザイン・ファーストをシュッとする](https://qiita.com/ikawaha/items/6638ee8b6978aef50d65)


## Mixing multiple feeds

```yaml
- urls:
    - https://zenn.dev/ikawaha/feed
    - https://qiita.com/ikawaha/feed
  template: |-
    {{range . -}}
      * {{ if eq .Header.FeedLink "https://zenn.dev/ikawaha/feed" -}}
            ![](./icon/zenn.png)
        {{- else }}{{ if eq .Header.FeedLink "https://qiita.com/ikawaha/feed" -}}
            ![](./icon/qiita.png)
        {{- end }}{{ end -}}
      [{{ .Title }}]({{ .Link }})
    {{ end }}
  limit: 10
  sort_by_published: true
```

Outputs:

* ![](./icon/zenn.png)[その文字が JIS X 0208 に含まれるか？ あるいは unicode.RangeTable の使い方](https://zenn.dev/ikawaha/articles/20210116-ab1ac4a692ae8bb4d9cf)
* ![](./icon/zenn.png)[Go製全文検索エンジンBlugeで日本語形態素解析をおこなう](https://zenn.dev/ikawaha/articles/20201230-84b042603ccbbce645d5)
* ![](./icon/zenn.png)[実践：形態素解析 kagome v2](https://zenn.dev/ikawaha/books/kagome-v2-japanese-tokenizer)
* ![](./icon/zenn.png)[形態素解析器 kagome を Google App Engine で動かす](https://zenn.dev/ikawaha/articles/hatena-20161004-221708)
* ![](./icon/zenn.png)[形態素解析器 kagome を brew tap でインストールできるようにした](https://zenn.dev/ikawaha/articles/20201029-391c049a13fb8506361d)
* ![](./icon/zenn.png)[Goa v3 入門](https://zenn.dev/ikawaha/books/goa-design-v3)
* ![](./icon/zenn.png)[Qiita/はてブの記事をZennでも管理する](https://zenn.dev/ikawaha/articles/20201012-e56b19cd33c396ae0465)
* ![](./icon/zenn.png)[はてなの記事をZenn形式のMarkdownで保存して管理する](https://zenn.dev/ikawaha/articles/hatena-20201012-205602)
* ![](./icon/qiita.png)[Qiitaの記事をZenn形式のMarkdownで保存して管理する](https://qiita.com/ikawaha/items/ab9906581e34f26993a9)
* ![](./icon/zenn.png)[Qiitaの記事をZenn形式のMarkdownで保存して管理する](https://zenn.dev/ikawaha/articles/qiita-ab9906581e34f26993a9)

# Automatic update using github workflow

```yaml
name: Update feed snippet

on:
  workflow_dispatch:
  schedule:
    - cron:  '0 * * * *'  

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2

    - name: Set up Go
      uses: actions/setup-go@v2
      with:
        go-version: 1.16

    - name: Install feedsnippet
      run: go install github.com/ikawaha/feedsnippet@latest

    - name: Update README.md
      run: feedsnippet -config feedsnippet.yml -file README.md

    - name: git commit
      run: |
        git config --local user.email "ikawaha@users.noreply.github.com"
        git config --local user.name "ikawaha"
        git add README.md
        git commit -m "Update feed snippet"
        git push origin main
```

Outputs:


<!--[START github.com/ikawaha/feedsnippet]--><!--[2021-02-21T03:43:30Z]-->
* ![](./icon/qiita.png)[travis setup releases で Validation Failed になるときの解消法](https://qiita.com/ikawaha/items/c4859421d67a27e9ffc6)
* ![](./icon/qiita.png)[Travis CI で goxc を使うとなぜかgo1.3のときだけビルドがこける](https://qiita.com/ikawaha/items/f0a046428f907bf9ba98)
* ![](./icon/qiita.png)[pythonクライアントで始める「はじめてのElasticsearch」](https://qiita.com/ikawaha/items/c654f746cfe76b888a27)
* ![](./icon/qiita.png)[デフォルトのElasticsearchは常にクラスタを組みたがっているので気をつける，というかクラスタ設定とめとく](https://qiita.com/ikawaha/items/f43a78f171836f1ec86b)
* ![](./icon/qiita.png)[goをソースからコンパイルしてgodocが見つからない](https://qiita.com/ikawaha/items/ea1ca5ee90296307c4bc)
* ![](./icon/qiita.png)[Elasticsearch で 文字の正規化を icu_normalizer でおこなう](https://qiita.com/ikawaha/items/79fdd69c524310db065e)
* ![](./icon/qiita.png)[golang で string を []byte にキャストするとメモリコピーが走ります](https://qiita.com/ikawaha/items/3c3994559dfeffb9f8c9)
* ![](./icon/qiita.png)[golang のリファクタリングには gofmt が使える](https://qiita.com/ikawaha/items/73c0a1d975680276b2c7)
* ![](./icon/qiita.png)[【kagome】形態素解析のラティスをグラフで表示できるようにしてみた](https://qiita.com/ikawaha/items/6fd4071187640f6bf66d)
* ![](./icon/qiita.png)[json を pretty print するのに echo '{"apple": "red", "lemon": "yellow"}' | python -m json.tool は冗長じゃないですか？なので go でコマンド用意しました](https://qiita.com/ikawaha/items/8a01c5739401e26e8794)
<!--[END github.com/ikawaha/feedsnippet]-->







---
MIT
