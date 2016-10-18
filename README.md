# slack-tools

## なにこれ？

コマンドの実行結果とかをstdinに流すとSlackにsnippet形式で投げてくれる  
ICTSCとかISUCON用に作ったやつ  

## Install

```
go get github.com/naoki912/slack-tools
```

もしくは

```
git clone https://github.com/naoki912/slack-tools
cd slack-tools
go build
```

## Usage

```
command | slack-tools -token="xxxx-xxxxxxxxx-xxxx" -channels="random"
```

### リモートのコマンド実行結果をSlackに投げる

ローカル側から実行する
```
ssh user@10.0.0.1 command | ./slack-tools -token="xxxx-xxxxxxxxx-xxxx" -channels="random"
```

サーバ側で実行する(外部のライブラリ等使用していないので静的リンクとかしなくても大丈夫なはず)
```
scp slack-tools 10.0.0.1:
command | ./slack-tools -token="xxxx-xxxxxxxxx-xxxx" -channels="random"
```

### オプション

|Argument|Example|Required|Description|
|---|---|---|---|
|`-token`|`xxxx-xxxxxxxxx-xxxx`|Required|Authentication token|
|`-channels`|`random` or `random,general`|Optional|投稿するチャンネル名 カンマ区切りで複数指定|
|`-filename`|`filename.txt`|Optional|snippetのファイル名|

### いちいちオプションでtokenとかchannelsとかを指定するのが面倒くさい

tokenとかchannelsをバイナリに埋め込める  
`main.go` 内、15行目ぐらいのvar()で定義されている変数にデフォルトで使用したい値を突っ込めばok(オプションで上書き可能)  
サーバ上に実行バイナリを配置して使う場合にオプションを指定無くて済むので楽  

## ToDo
* filetypeを渡せるようにする
* バイナリをアップロード出来るようにする
* 環境変数でtokenやchannelsを設定できるようにする
* confファイルでtokenやchannelsを設定できるようにする
    - プロファイルを切り替えられるようにする
