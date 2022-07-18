# DynamoDBのuserテーブルにデータを作成できるlambdaFunctionサンプル

# 実行
1. dynamoDBにuserテーブル作成（パーティションキーはuser_id）
2. `$ sh build.sh`を実行。（権限エラーが出たら `chmod 777 build.sh`を実行）
3. lambdaに関数を作成してfunction.zipをアップロード。ハンドラを `bin/main`に変更。
4. APIGatewayか lambdaの関数URLを設定してPOST。

- POST時のサンプルjson

```json
{
  "user_id": 1,
  "last_name": "テスト",
  "first_name": "太郎"
}
```