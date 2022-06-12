# ディレクトリ構成

```text
.
├── aws.yml    // github公式から引用して編集
└── base64.sh  // task-definition.jsonをbase64にエンコードするコマンド
```

# 概要
1. Actions secretsにAWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEYを登録
2. task-definition.jsonをエンコードしたものをTASK_DEFINITIONに登録
