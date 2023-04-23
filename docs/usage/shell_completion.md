# シェル補完を有効にする

Qiita CLI はシェル補完の有効化をサポートしています。  
`qiita completion <SHELL>` を実行すると補完用のシェルスクリプトが出力されます。

```sh
# <SHELL> には使用しているシェルを指定する
$ qiita completion <SHELL>
```

サポートしているシェルは次の通りです。

```sh
# Bash
$ qiita completion bash

# Fish
$ qiita completion fish

# PowerShell
$ qiita completion powershell

# Zsh
$ qiita completion zsh
```

実際にシェル補完を有効にする手順についてはそれぞれのシェルのヘルプをご参照ください。

```sh
$ qiita completion <SHELL> --help
```

例えば Zsh を使用している場合、次のように実行すると現在のセッションでシェル補完が有効になります。

```sh
$ source <(qiita completion zsh)
```
