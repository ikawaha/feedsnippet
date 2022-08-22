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
    - cron:  '0 0 * * *'  

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
      run: feedsnippet -config feedsnippet.yml -diff -file README.md

    - name: git commit
      run: |
        git config --local user.email "ikawaha@users.noreply.github.com"
        git config --local user.name "ikawaha"
        git add README.md
        git diff --cached --quiet || (git commit -m "Update feed snippet" && git push origin main)
```

Outputs:


<!--[START github.com/ikawaha/feedsnippet]--><!--[2022-08-22T00:14:06Z]-->
* ![](./icon/zenn.png)[Goa 更新情報 v3.8.3](https://zenn.dev/ikawaha/articles/20220821-b5356517f44727)
* ![](./icon/zenn.png)[Goa 更新情報 v3.8.2](https://zenn.dev/ikawaha/articles/20220811-66619772b2b8a4)
* ![](./icon/zenn.png)[ドドスコするオートマトン考](https://zenn.dev/ikawaha/articles/20220806-55c9db03732a09)
* ![](./icon/zenn.png)[Goa 更新情報 v3.7.14](https://zenn.dev/ikawaha/articles/20220801-ad319824c789e8)
* ![](./icon/zenn.png)[さよなら WebDriver Client Agouti](https://zenn.dev/ikawaha/articles/20220722-4719183593a3cb)
* ![](./icon/zenn.png)[Goa 更新情報 v3.7.12](https://zenn.dev/ikawaha/articles/20220717-c8f77d794dd18e)
* ![](./icon/zenn.png)[clue 更新情報 v0.9.0](https://zenn.dev/ikawaha/articles/20220715-024418951cf58d)
* ![](./icon/zenn.png)[Goa 更新情報 v3.7.10](https://zenn.dev/ikawaha/articles/20220702-2fb029d5b376fd)
* ![](./icon/zenn.png)[Go の workspace を使って go. mod を汚さずに replace を書く](https://zenn.dev/ikawaha/articles/20220701-a053ec54b77435)
* ![](./icon/zenn.png)[単純な Slack Bot を単純に書くための Go のライブラリをメンテした](https://zenn.dev/ikawaha/articles/20220628-0bd46584a3063d)
<!--[END github.com/ikawaha/feedsnippet]-->
---
MIT
