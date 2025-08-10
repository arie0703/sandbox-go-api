# Microservices Cobra CLI

GoとCobraフレームワークを使用したマイクロサービスCLIアプリケーションです。7つの番号付きタスクコマンド（task1-task7）を提供し、AWS ECSで個別にデプロイできます。

## 機能

- **7つのタスクコマンド**: task1からtask7まで、それぞれ独立したコマンド
- **Dockerコンテナ化**: マルチステージビルドによる軽量なコンテナイメージ
- **AWS ECSデプロイ**: 各タスクを個別のECSタスク定義としてデプロイ
- **GitHub Actions**: 並列デプロイメントによる効率的なCI/CD

## 使用方法

### ローカル実行

```bash
# プロジェクトをビルド
go build -o microservice-cli .

# ヘルプを表示
./microservice-cli --help

# 特定のタスクを実行
./microservice-cli task1
./microservice-cli task2
# ... task7まで
```

### Dockerでの実行

```bash
# イメージをビルド
docker build -t microservice .

# ヘルプを表示
docker run --rm microservice

# 特定のタスクを実行
docker run --rm microservice task1
docker run --rm microservice task2
# ... task7まで
```

## デプロイメント設定

### 必要なGitHub Secrets

GitHub Actionsワークフローを使用するために、以下のシークレットを設定してください：

```
AWS_ROLE_TO_ASSUME        # AssumeするAWS IAMロールのARN
```

**注意**: AWS Account IDは動的に取得されるため、Secretsに設定する必要はありません。ワークフロー実行時に`aws sts get-caller-identity`コマンドで自動取得されます。

### 必要なGitHub Variables

以下の変数を設定してください（オプション、デフォルト値が設定されています）：

```
ENVIRONMENT               # デプロイ環境 (デフォルト: production)
AWS_REGION               # AWSリージョン (デフォルト: ap-northeast-1)
ECS_CLUSTER              # ECSクラスター名 (デフォルト: microservice-cluster)
```

### 環境変数の設定

`.github/workflows/deploy.yml`で以下の環境変数を適切に設定してください：

```yaml
env:
  AWS_REGION: ap-northeast-1               # AWSリージョン
  ECR_REPOSITORY: microservice             # ECRリポジトリ名
  ECS_CLUSTER: microservice-cluster       # ECSクラスター名
  ENVIRONMENT: production                  # デプロイ環境
```

### 必要なAWSリソース

デプロイメント前に以下のAWSリソースを作成してください：

1. **ECRリポジトリ**: `microservice`
2. **ECSクラスター**: `microservice-cluster`
3. **CloudWatch Logsグループ**: `/ecs/microservice`
4. **IAMロール**:
   - ECSタスクロール（必要な権限を付与）
   - ECS実行ロール（ECRプル権限とCloudWatch Logs書き込み権限）

### IAMロールの権限例

#### ECSタスクロール
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "arn:aws:logs:*:*:log-group:/ecs/microservice:*"
    }
  ]
}
```

#### ECS実行ロール
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ecr:GetAuthorizationToken",
        "ecr:BatchCheckLayerAvailability",
        "ecr:GetDownloadUrlForLayer",
        "ecr:BatchGetImage",
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "*"
    }
  ]
}
```

## プロジェクト構造

```
micro-services/
├── cmd/
│   └── root.go              # Cobraルートコマンドとサブコマンド定義
├── task_definitions/        # ECSタスク定義ファイル
│   ├── task1/
│   │   └── production/
│   │       ├── task-definition.json
│   │       └── container-definitions/
│   │           └── microservice-container.json
│   ├── task2/
│   │   └── production/
│   │       ├── task-definition.json
│   │       └── container-definitions/
│   │           └── microservice-container.json
│   └── ... (task3-task7も同様の構造)
├── Dockerfile              # マルチステージDockerビルド
├── .dockerignore           # Docker除外ファイル
├── go.mod                  # Goモジュール定義
├── go.sum                  # 依存関係チェックサム
├── main.go                 # エントリーポイント
└── README.md               # このファイル
```

### GitHub Actionsワークフロー構造

```
.github/workflows/
├── deploy-micro-services.yml  # メインデプロイワークフロー
└── deploy-common.yml          # 共通デプロイ処理（再利用可能ワークフロー）
```

## デプロイメントフロー

1. **トリガー**: mainブランチへのpushまたは手動実行
2. **ビルド**: Dockerイメージをビルドし、ECRにプッシュ
3. **並列デプロイ**: Matrix戦略を使用して7つのタスク定義を並列でECSに更新
   - `strategy.matrix`で`task1`〜`task7`を並列実行
   - 各タスクは`deploy-common.yml`の再利用可能ワークフローを呼び出し
   - タスク定義とコンテナ定義を分離した構造で管理
   - イメージの動的置換とAWS認証情報の安全な処理
   - `fail-fast: false`により、1つのタスクが失敗しても他のタスクは継続
4. **結果報告**: 各タスクのデプロイメント結果を報告

### Matrix戦略の利点

- **コードの簡潔性**: 7つの個別ジョブを1つのジョブ定義で表現
- **保守性の向上**: タスクの追加・削除が配列の変更だけで可能
- **並列実行**: 全てのタスクが同時に実行され、デプロイ時間を短縮
- **エラー処理**: `fail-fast: false`により部分的な失敗を許容

### プレースホルダー置換機能

タスク定義ファイル内のプレースホルダーは、デプロイ時に自動的に実際の値に置換されます：

- **ACCOUNT**: AWS Account ID（`aws sts get-caller-identity`で動的取得）
- **REGION**: AWSリージョン（ワークフロー設定から取得）
- **イメージURI**: ECRイメージURI（ビルド時に生成）

これにより、環境固有の値をハードコードすることなく、柔軟なデプロイメントが可能です。

### 再利用可能ワークフローの利点

- **コードの重複排除**: 共通のデプロイ処理を一箇所で管理
- **保守性の向上**: デプロイロジックの変更が一箇所で済む
- **一貫性の確保**: 全てのタスクで同じデプロイ処理を使用
- **柔軟性**: 環境ごとの設定を簡単に変更可能

## トラブルシューティング

### よくある問題

1. **ECR認証エラー**: AWS認証情報とECRリポジトリの存在を確認
2. **ECSデプロイエラー**: IAMロールの権限とクラスターの存在を確認
3. **タスク定義エラー**: JSON形式とリソース制限を確認

### ログの確認

- **GitHub Actions**: リポジトリのActionsタブでワークフロー実行ログを確認
- **ECS**: CloudWatch Logsの`/ecs/microservice`ロググループを確認
- **ECR**: ECRコンソールでイメージプッシュ状況を確認