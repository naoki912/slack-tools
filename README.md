# slack-tools

openstackコマンド風にslackを操作したかった  
使ってるものしか実装してない

## Install

```
go get github.com/naoki912/slack-tools
```

## 使い方
ドキュメントなんて無かった

```
$ slack-tools --help
```

### 設定ファイル
~/.slack-tools

```yaml
slack:
  token: xoxp-xxxxxxxxxxxxxxxxxxxxxxxx
```

## ToDo
- 冗長すぎなので直す
- cmdとslackAPI叩きに行っている部分を分ける
