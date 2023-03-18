# ultrachat

## コマンド
```bash
go run github.com/99designs/gqlgen generate
```

### テーブル作成
```
aws dynamodb --region ap-northeast-1 \
--endpoint-url http://localhost:8000 \
  create-table \
--table-name ChatMessages \
--attribute-definitions \
  AttributeName=id,AttributeType=S \
  AttributeName=message,AttributeType=S \
--key-schema \
  AttributeName=id,KeyType=HASH \
  AttributeName=message,KeyType=RANGE \
--billing-mode PAY_PER_REQUEST
```

### カラム追加
```
aws dynamodb update-table \
  --region ap-northeast-1 \
  --endpoint-url http://localhost:8000 \
  --table-name ChatMessages \
  --attribute-definitions AttributeName=createdAt,AttributeType=S \
  --billing-mode PAY_PER_REQUEST
```

### テーブルの中身確認
aws dynamodb scan --table-name ChatMessages --endpoint-url http://localhost:8000 --region ap-northeast-1


## 準備
### s3のバケット作成
```
aws s3 mb s3://parent-yoyoyo
aws s3 cp ./cfn/parent.yaml s3://parent-yoyoyo/
aws s3 cp ./cfn/child.yaml s3://parent-yoyoyo/
```

### 親スタックの作成
```
// スタック作成
aws cloudformation create-stack --stack-name parent-stack --template-body file://cfn/parent.yaml --capabilities CAPABILITY_AUTO_EXPAND

// スタック削除
aws cloudformation delete-stack --stack-name parent-stack
```
