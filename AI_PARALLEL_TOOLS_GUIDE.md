# AI並列ツール使用ガイド

## なぜこのツールが必要なのか？

### 毎日こんな時間を無駄にしていませんか？
- **コード理解に30分** - 「この関数何してる？」の繰り返し
- **思考の中断** - ファイル間を行ったり来たり
- **実装の手戻り** - 理解不足による設計ミス
- **AIへの質問を1つずつ** - 待ち時間の累積

### 5分後のあなた
- ✅ **50個の疑問を13秒で解決** - 並列処理の威力
- ✅ **プロジェクト全体を瞬時に理解** - 横断的分析
- ✅ **思考が途切れない開発** - 理解→実装がスムーズ
- ✅ **期待通りのコード生成** - コンテキスト機能

## ツール概要

**see_parallel** と **code_parallel** は、AI-First開発を実現する2つの核となるツールです。

### 基本コンセプト
- **see_parallel**: 理解・分析に特化（複数ファイルを並列で分析）
- **code_parallel**: 生成・実装に特化（複数ファイルを並列で生成）

### 実測パフォーマンス
- **see_parallel**: 2つの深い分析を11秒で完了
- **code_parallel**: 3ファイルを6.7秒で生成
- **スケーラビリティ**: 1つでも50個でも実行時間はほぼ一定

## 共通仕様

### コマンド構造
```bash
# どちらも同じ構造
tool_name queue '["内容", "ファイル1", "ファイル2", ...]'
tool_name queue '["内容", "ファイル", "deep"]'
tool_name queue run --parallel N
tool_name queue list
tool_name queue clear
```

### 入力フォーマット
```
配列形式：["内容", "ファイル1", "ファイル2", ..., "オプション"]
- [0]: 質問/タスク（必須）
- [1..n]: ファイルパス（1個以上必須）
- [last]: "deep"で上位モデル使用
```

## see_parallel - 理解・分析ツール

### 基本的な使い方

#### 単一ファイル分析
```bash
see_parallel queue '["このファイルの主要な関数は？", "lib/auth.ts"]'
```

#### 複数ファイル横断分析
```bash
see_parallel queue '["認証システム全体の仕組みは？", "lib/auth.ts", "lib/jwt.ts", "middleware.ts"]'
```

#### 深い分析（上位モデル）
```bash
see_parallel queue '["セキュリティリスクを詳細分析", "lib/auth.ts", "deep"]'
```

#### ワイルドカード使用
```bash
see_parallel queue '["プロジェクト全体の構造は？", "**/*.ts", "**/*.tsx"]'
```

### 実用的な質問例

#### コード理解
```bash
see_parallel queue '["このコンポーネントの責任は？", "components/UserCard.tsx"]'
see_parallel queue '["関数の依存関係は？", "utils/helper.ts"]'
see_parallel queue '["データフローは？", "store/userStore.ts"]'
```

#### 品質分析
```bash
see_parallel queue '["パフォーマンスの問題は？", "pages/dashboard.tsx", "deep"]'
see_parallel queue '["コードの改善点は？", "lib/database.ts", "deep"]'
see_parallel queue '["テストカバレッジが必要な箇所は？", "services/api.ts"]'
```

#### セキュリティ分析
```bash
see_parallel queue '["セキュリティホールはあるか？", "lib/auth.ts", "deep"]'
see_parallel queue '["入力検証は適切か？", "api/routes.ts", "deep"]'
```

#### アーキテクチャ分析
```bash
see_parallel queue '["モジュール間の結合度は？", "src/**/*.ts"]'
see_parallel queue '["設計パターンの使用状況は？", "lib/*.ts", "deep"]'
```

## code_parallel - 生成・実装ツール

### 基本的な使い方

#### 単一ファイル生成
```bash
code_parallel queue '["認証機能を実装", "lib/auth-new.ts"]'
```

#### 複数ファイル生成
```bash
code_parallel queue '["CRUD APIを実装", "api/users.ts", "api/posts.ts", "api/comments.ts"]'
```

#### 複雑な実装（上位モデル）
```bash
code_parallel queue '["高性能なアルゴリズムを実装", "lib/optimization.ts", "deep"]'
```

### 実用的なタスク例

#### 基本的なコンポーネント
```bash
code_parallel queue '["再利用可能なボタンコンポーネント", "components/Button.tsx"]'
code_parallel queue '["データテーブルコンポーネント", "components/DataTable.tsx"]'
code_parallel queue '["モーダルダイアログ", "components/Modal.tsx"]'
```

#### ビジネスロジック
```bash
code_parallel queue '["ユーザー管理サービス", "services/userService.ts"]'
code_parallel queue '["決済処理ロジック", "lib/payment.ts", "deep"]'
code_parallel queue '["データ変換ユーティリティ", "utils/transform.ts"]'
```

#### API実装
```bash
code_parallel queue '["RESTful ユーザーAPI", "api/users.ts"]'
code_parallel queue '["GraphQL リゾルバ", "graphql/resolvers.ts"]'
code_parallel queue '["WebSocket ハンドラ", "ws/handlers.ts"]'
```

#### テストコード
```bash
code_parallel queue '["単体テスト", "tests/auth.test.ts"]'
code_parallel queue '["統合テスト", "tests/integration.test.ts"]'
code_parallel queue '["E2Eテスト", "e2e/user-flow.test.ts"]'
```

## 効率的なワークフロー

### 理想的な開発フロー
```bash
# 1. 既存コードの理解
see_parallel queue '["現在の認証システムは？", "lib/auth.ts"]'
see_parallel queue '["APIの構造は？", "api/*.ts"]'
see_parallel queue '["コンポーネントの役割は？", "components/*.tsx"]'
see_parallel queue run --parallel 10

# 2. 実装計画の立案（結果を確認後）
# 3. 新機能の実装
code_parallel queue '["改良版認証システム", "lib/auth-v2.ts"]'
code_parallel queue '["新しいAPIエンドポイント", "api/v2/users.ts"]'
code_parallel queue '["対応するフロントエンド", "components/AuthV2.tsx"]'
code_parallel queue run --parallel 5

# 4. 品質確認
see_parallel queue '["生成されたコードの品質は？", "lib/auth-v2.ts", "deep"]'
see_parallel queue run --parallel 1
```

### バッチ処理のコツ
```bash
# 関連する質問をまとめて投入
see_parallel queue '["認証の仕組み", "lib/auth.ts"]'
see_parallel queue '["セッション管理", "lib/session.ts"]'
see_parallel queue '["権限制御", "lib/permissions.ts"]'
see_parallel queue '["セキュリティ対策", "lib/security.ts", "deep"]'
see_parallel queue run --parallel 4
```

## 実行とモニタリング

### キューの管理
```bash
# 現在のキューを確認
see_parallel queue list
code_parallel queue list

# 実行
see_parallel queue run --parallel 10
code_parallel queue run --parallel 5

# キューをクリア
see_parallel queue clear
code_parallel queue clear
```

### 並列数の選択指針
- **軽い質問・タスク**: --parallel 10-20
- **重い分析・実装**: --parallel 3-5
- **deep モード**: --parallel 1-3

## コンテキスト機能（重要）

### プロジェクトコンテキストの設定
両ツールで最も重要な機能です。一度設定すれば、すべての操作に自動適用されます。

```bash
# プロジェクト情報を設定
code_parallel context set "Next.js 15 TypeScriptプロジェクト、Tailwind CSS使用"
see_parallel context set "Next.js 15 TypeScriptプロジェクト、Tailwind CSS使用"

# より詳細な設定例
code_parallel context set "Next.js TypeScript、AI-First原則でコード重複推奨、Turso DB使用"
see_parallel context set "工場-顧客コミュニケーションシステム、Magic Link認証"
```

### コンテキストの確認・管理
```bash
# 現在の設定を確認
code_parallel context show
see_parallel context get

# 設定をクリア
code_parallel context clear
see_parallel context clear
```

### コンテキストの効果
```bash
# Before（コンテキストなし）
code_parallel queue '["REST API実装", "api/users.ts"]'
→ PythonのFlaskコードが生成される...？

# After（コンテキストあり）
code_parallel queue '["REST API実装", "api/users.ts"]'
→ Next.js API Routes、TypeScript、Tailwind CSSで生成！
```

## APIキー設定

### 初回設定
```bash
see_parallel api set "your-gemini-api-key"
code_parallel api set "your-gemini-api-key"
```

### 設定確認
```bash
see_parallel api status
code_parallel api status
```

## ベストプラクティス

### ✅ 推奨事項
1. **コンテキストを最初に設定**: プロジェクト開始時に必ず実行
2. **具体的な質問・タスク**: 「認証機能を分析」より「JWT トークンの有効期限設定は適切か？」
3. **関連ファイルをまとめて指定**: 機能単位でファイルをグループ化
4. **deep モードの適切な使用**: 複雑な分析・実装のみ
5. **並列数の調整**: APIレート制限を考慮

### ❌ 避けるべき事項
1. **曖昧な指示**: 「何か作って」「適当に分析して」
2. **無関係なファイルの混在**: 認証とUIコンポーネントを同時分析
3. **過度な並列実行**: APIレート制限違反
4. **deep モードの乱用**: 簡単なタスクでのコスト増加

## トラブルシューティング

### よくある問題
```bash
# キューが空の場合
ERROR: キューは空です
→ まず queue でタスクを追加

# APIキーエラー
ERROR: API key expired
→ see_parallel api set "new-key"

# ファイルが見つからない
ERROR: File not found
→ ファイルパスを確認（相対パス推奨）
```

## 高度な使用例

### プロジェクト全体の分析
```bash
see_parallel queue '["アーキテクチャ概要", "src/**/*.ts", "components/**/*.tsx"]'
see_parallel queue '["技術的負債", "**/*.ts", "**/*.tsx", "deep"]'
see_parallel queue '["パフォーマンス分析", "pages/**/*.tsx", "deep"]'
see_parallel queue '["セキュリティ監査", "lib/**/*.ts", "api/**/*.ts", "deep"]'
see_parallel queue run --parallel 4
```

### 大規模リファクタリング準備
```bash
# 1. 現状分析
see_parallel queue '["依存関係マップ", "src/**/*.ts"]'
see_parallel queue '["重複コード検出", "**/*.ts", "deep"]'
see_parallel queue run --parallel 2

# 2. 新設計実装
code_parallel queue '["モジュラー認証システム", "lib/auth/index.ts"]'
code_parallel queue '["型安全なAPIクライアント", "lib/api/client.ts"]'
code_parallel queue '["共通ユーティリティ", "lib/utils/index.ts"]'
code_parallel queue run --parallel 3
```

---

これらのツールを活用することで、**理解から実装まで**のAI-First開発サイクルを効率的に回すことができます。