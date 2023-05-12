# 概要

アプリケーション上で発生したイベントを通知するアプリケーション。
CloudRun にデプロイして PubSub から受け取った特定のイベントを通知先に通知する。
現在 CloudLogging で発生したエラーレベル以上のイベントを Slack への通知を想定した実装になっている。

# CloudRun へのデプロイ手順

1. 環境の設定

``` bash
gcloud auth login
ENV={dev|prd}
PROJECT_ID=sample-$ENV
gcloud config set project $PROJECT_ID
```

2. Container Registory へ image の push

``` bash
IMAGE_NAME=sample-logging-notifer-$ENV
gcloud builds submit --tag gcr.io/$PROJECT_ID/$IMAGE_NAME
```

3. Cloud Run へのデプロイ

```
gcloud run deploy --image gcr.io/$PROJECT_ID/$IMAGE_NAME --platform managed
```

# PubSub の設定

## 権限の付与

Cloud Run サービスデプロイ後、メッセージを push するように Pub/Sub を構成する。
PROJECT_NUMBER は gcp console から確認する
※はじめの一回のみ

``` bash
gcloud projects add-iam-policy-binding $PROJECT_ID \
     --member=serviceAccount:service-$PROJECT_NUMBER@gcp-sa-pubsub.iam.gserviceaccount.com \
     --role=roles/iam.serviceAccountTokenCreator
```

``` bash
gcloud iam service-accounts create cloud-run-pubsub-invoker \
     --display-name "Cloud Run Pub/Sub Invoker"
```

``` bash
gcloud run services add-iam-policy-binding $IMAGE_NAME \
   --member=serviceAccount:cloud-run-pubsub-invoker@$PROJECT_ID.iam.gserviceaccount.com \
   --role=roles/run.invoker
```

## PubSub Subscription の設定

``` bash
gcloud pubsub topics create cloudlogging
$ENDOPOINT={https://...}
gcloud pubsub subscriptions create error-sub --topic cloudlogging \
 --push-endpoint=$ENDOPOINT\
 --push-auth-service-account=cloud-run-pubsub-invoker@$PROJECT_ID.iam.gserviceaccount.com
```

# Cloud Logging の設定

## ログルーターの設定

GCP コンソールの`logging` から`ログルーター` を選択し、以下の条件で PUbSub のトピックにレベルがエラー以上のものを push するように設定する

``` bash
resource.type = cloud_run_revision AND severity >= ERROR
```

# テスト

.env の`SLACK_WEBHOOK_TEST_URL`にテスト送信先の WEBHOOK URL を設定

``` bash
go test -run ./logging/
```
