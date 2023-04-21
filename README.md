<h1 align="center">
Qiita CLI
</h1>

<p align="center">
<a href="https://github.com/koki-develop/qiita-cli/releases/latest"><img src="https://img.shields.io/github/v/release/koki-develop/qiita-cli?style=flat-square" alt="GitHub release (latest by date)"></a>
<a href="https://github.com/koki-develop/qiita-cli/actions/workflows/ci.yml"><img src="https://img.shields.io/github/actions/workflow/status/koki-develop/qiita-cli/ci.yml?logo=github&amp;style=flat-square" alt="GitHub Workflow Status"></a>
<a href="https://codeclimate.com/github/koki-develop/qiita-cli/maintainability"><img src="https://img.shields.io/codeclimate/maintainability/koki-develop/qiita-cli?style=flat-square&amp;logo=codeclimate" alt="Maintainability"></a>
<a href="https://goreportcard.com/report/github.com/koki-develop/qiita-cli"><img src="https://goreportcard.com/badge/github.com/koki-develop/qiita-cli?style=flat-square" alt="Go Report Card"></a>
<a href="./LICENSE"><img src="https://img.shields.io/github/license/koki-develop/qiita-cli?style=flat-square" alt="LICENSE"></a>
</p>

## 目次

- [インストール](TODO)
- [クイックスタート](TODO)
- [ライセンス](TODO)

## :zap: インストール

### Homebrew

```console
$ brew install koki-develop/tap/qiita
```

### `go install`

```console
$ go install github.com/koki-develop/qiita-cli/cmd/qiita@latest
```

### リリース

[リリースページ](https://github.com/koki-develop/qiita-cli/releases/latest)からバイナリをダウンロードしてください。

## :beginner: クイックスタート

### 1. Qiita アクセストークンを発行

[Qiita](https://qiita.com) にログイン後、[アクセストークンの発行ページ](https://qiita.com/settings/tokens/new)にアクセスしてください。  
それぞれの項目を次のように入力します。

| 項目 | 説明 |
| --- | --- |
| `アクセストークンの説明` | 任意のテキスト。 |
| `スコープ` | `read_qiita` と `write_qiita` を選択。 |

入力後に `発行する` をクリックするとアクセストークンが発行されるため、ひかえておきます。

### 2. Qiita CLI の設定

まず Qiita CLI の設定を行います。  
`qiita configure` を実行してアクセストークンとデフォルトの出力フォーマットを対話的に設定します。

```sh
$ qiita configure
```

これで Qiita CLI の準備は完了です！ :tada:

### 3. 自分の記事一覧を表示する

試しに Qiita CLI を使って自分の記事一覧を表示してみましょう。  
`qiita items list` を実行します。

```sh
$ qiita items list
```

Qiita CLI の設定を正しく行えていればコンソール上にあなたの記事一覧が出力されるはずです！  

他にも Qiita CLI は様々な操作を行うことができます。  
詳しくは[使い方](./docs/usage.md)をご参照ください。

## :memo: ライセンス

[MIT](./LICENSE)
