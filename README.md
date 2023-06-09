# ultrachat

## コマンド

### graphql
```bash
go run github.com/99designs/gqlgen generate
```

### dbコンテナ

```bash
docker run --rm -d --name mysql -p 3307:3306 -e MYSQL_DATABASE=ultrachat -e MYSQL_USER=wuser -e MYSQL_ROOT_PASSWORD=password -e MYSQL_PASSWORD=password -e MYSQL_HOST=localhost mysql:8
```

### xo
```
xo schema 'mysql://wuser:password@localhost:3307/ultrachat?parseTime=true'
```

### migration

```bash
sql-migrate up -config config/dbconfig.yml
```

### dynamodb
```
# テーブル作成
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

```
# カラム追加
aws dynamodb update-table \
  --region ap-northeast-1 \
  --endpoint-url http://localhost:8000 \
  --table-name ChatMessages \
  --attribute-definitions AttributeName=createdAt,AttributeType=S \
  --billing-mode PAY_PER_REQUEST
```

```
# テストデータ
aws dynamodb batch-write-item \
    --request-items file://testdata.json \
    --endpoint-url http://localhost:8000
```

```
# テーブルの中身確認
aws dynamodb scan --table-name ChatMessages --endpoint-url http://localhost:8000 --region ap-northeast-1
```


### cfn
```
# s3のバケット作成
aws s3 mb s3://parent-yoyoyo
aws s3 cp ./cfn/parent.yaml s3://parent-yoyoyo/
aws s3 cp ./cfn/child.yaml s3://parent-yoyoyo/
```

```
// スタック作成
aws cloudformation create-stack --stack-name parent-stack --template-body file://cfn/parent.yaml --capabilities CAPABILITY_AUTO_EXPAND

// スタック削除
aws cloudformation delete-stack --stack-name parent-stack
```

### appsync

```
# appsyncのテストエンドポイントの叩き方
aws appsync \
  --endpoint-url https://5buzdr2hpnbcdowjjoihjp4izu.appsync-api.ap-northeast-1.amazonaws.com/graphql \
  --region ap-northeast-1 \
  graphql \
  --query 'query listChatMessages { listChatMessages { id message createdAt } }' \
  --output json \
  --auth-type API_KEY \
  --api-key da2-zlr6zecdc5ctreox4z4h7mrvju
```

### s3
```
aws s3 sync . s3://sample-sample-sample/ \
    --include "script/sample/*" \
    --acl public-read \
    --cache-control "max-age=3600"
```
