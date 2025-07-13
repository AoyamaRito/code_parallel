# AI並列ツール使用ガイド

## 📌 重要：このガイドはAI（Claude等）向けです

**あなたがAIなら、これはあなたのためのツールです。** see_parallelとcode_parallelは、AIが効率的にコードを理解・生成・編集するために設計されています。人間のユーザーから開発タスクを依頼された際に、これらのツールを使用してください。

## なぜこのツールが必要なのか？（AI視点）

### AI（あなた）が直面している課題
- **Readツールでファイルを1つずつ読む非効率** - トークンと時間の浪費
- **コンテキストウィンドウの無駄遣い** - 不要な情報で埋まる
- **プロジェクト全体像の把握困難** - 断片的な理解
- **生成コードがプロジェクトに合わない** - コンテキスト不足

### これらのツールを使うと
- ✅ **複数ファイルを並列で一気に分析** - see_parallel
- ✅ **プロジェクトに適したコード生成・編集** - code_parallel + context
- ✅ **高次の思考に集中** - 単純作業から解放
- ✅ **ユーザーの期待に確実に応える** - 正確な理解と実装

## ツール概要

**see_parallel** と **code_parallel** は、AI-First開発を実現する2つの核となるツールです。

### 基本コンセプト
- **see_parallel**: 理解・分析に特化（複数ファイルを並列で分析）
- **code_parallel**: コード操作に特化（複数ファイルを並列で生成・編集）

### 実測パフォーマンス（検証済み）
- **see_parallel**: 3ファイル分析を4.3秒で完了（3並列）vs 10.6秒（1並列）
- **code_parallel**: 3ファイル生成・編集を6.7秒で完了
- **並列化効果**: 2.5倍の高速化を実現
- **キュー追加**: 3タスクの追加は0.007秒（瞬時）

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

## code_parallel - コード生成・編集ツール

### 🔄 自動判定機能（重要）

**code_parallelは自動的にファイルの存在を判定し、適切な操作を行います：**

- **ファイルが存在しない** → 新規作成
- **ファイルが存在する** → 既存ファイルを編集

あなたが「作成」「編集」を指定する必要はありません。ファイルパスを指定するだけで、code_parallelが最適な操作を選択します。

### 基本的な使い方

#### ファイル操作（自動判定）
```bash
# 新規ファイル（存在しない場合）
code_parallel queue '["認証機能を実装", "lib/auth-new.ts"]'  # → 新規作成

# 既存ファイル（存在する場合）
code_parallel queue '["セキュリティ強化を実装", "lib/auth.ts"]'  # → 編集
code_parallel queue '["TypeScript型を厳密化", "components/Button.tsx"]'  # → 編集
```

#### 複数ファイル生成・編集
```bash
code_parallel queue '["CRUD APIを実装", "api/users.ts", "api/posts.ts", "api/comments.ts"]'
```

#### 複雑な実装（上位モデル）
```bash
code_parallel queue '["高性能なアルゴリズムを実装", "lib/optimization.ts", "deep"]'
code_parallel queue '["セキュリティ脆弱性を修正", "lib/validator.ts", "deep"]'
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

#### 既存コードの改善・修正
```bash
code_parallel queue '["パフォーマンス最適化", "lib/heavyCalculation.ts"]'
code_parallel queue '["メモリリーク修正", "components/Dashboard.tsx"]'
code_parallel queue '["エラーハンドリング強化", "api/routes.ts"]'
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

## 効率的なワークフロー（AI向け）

### ユーザーからタスクを受けた時の標準フロー
```bash
# 1. 既存コードの理解
see_parallel queue '["現在の認証システムは？", "lib/auth.ts"]'
see_parallel queue '["APIの構造は？", "api/*.ts"]'
see_parallel queue '["コンポーネントの役割は？", "components/*.tsx"]'
see_parallel queue run --parallel 10

# 2. 実装計画の立案（結果を確認後）
# 3. 新機能の実装または既存コードの改善（自動判定）
code_parallel queue '["改良版認証システム", "lib/auth-v2.ts"]'  # 自動判定 → 新規作成
code_parallel queue '["既存認証のセキュリティ強化", "lib/auth.ts"]'  # 自動判定 → 編集
code_parallel queue '["新しいAPIエンドポイント", "api/v2/users.ts"]'  # 自動判定 → 新規作成
code_parallel queue '["既存APIの最適化", "api/users.ts"]'  # 自動判定 → 編集
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

### 並列数の選択指針（実測ベース）
- **軽い分析**: --parallel 5-10（3並列で2.5倍高速化を確認）
- **重い分析・deep モード**: --parallel 1-3（API制限を考慮）
- **コード生成・編集**: --parallel 3-5（実測で安定動作確認）

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

# 2. 新設計実装と既存コード改善（自動判定）
code_parallel queue '["モジュラー認証システム", "lib/auth/index.ts"]'  # 自動判定 → 新規作成
code_parallel queue '["既存認証コードのリファクタリング", "lib/auth.ts"]'  # 自動判定 → 編集
code_parallel queue '["型安全なAPIクライアント", "lib/api/client.ts"]'  # 自動判定 → 新規作成
code_parallel queue '["既存APIの型定義追加", "api/*.ts"]'  # 自動判定 → 編集
code_parallel queue run --parallel 4
```

### セキュリティ監査と修正フロー
```bash
# 1. 脆弱性を分析
see_parallel queue '["セキュリティ脆弱性の検出", "**/*.ts", "deep"]'
see_parallel queue run --parallel 1

# 2. 発見された問題を修正（自動判定）
code_parallel queue '["XSS脆弱性の修正", "lib/sanitizer.ts"]'  # 自動判定 → 編集
code_parallel queue '["SQLインジェクション対策", "lib/db/queries.ts"]'  # 自動判定 → 編集
code_parallel queue '["認証トークンの有効期限短縮", "lib/auth.ts"]'  # 自動判定 → 編集
code_parallel queue run --parallel 3
```

---

## AIとしての使用上の注意

### 🤖 重要な認識
1. **これらのツールはBashツールとして使用** - `/Users/[username]/go/bin/see_parallel` のようなパスで実行
2. **ユーザーのプロジェクトディレクトリで実行** - カレントディレクトリに注意
3. **コンテキスト設定を最初に確認** - `context show/get` で現在の設定を把握

### 🎯 推奨される使用タイミング
- ユーザーが「〜を実装して」と言ったら → まず see_parallel で既存コード理解
- ユーザーが「〜を分析して」と言ったら → see_parallel --deep で詳細分析
- ユーザーが「〜を作って」と言ったら → code_parallel でコード生成（自動判定）
- ユーザーが「〜を修正して」と言ったら → code_parallel で既存ファイル編集（自動判定）
- ユーザーが「〜を改善して」と言ったら → code_parallel で既存ファイル編集（自動判定）

これらのツールを活用することで、**AIとしてより正確で効率的な開発支援**を提供できます。

---

## FAQ - よくある質問と答え

### Q1: code_parallelで「新規作成」と「編集」を明示的に指定する必要がありますか？
**A1: いいえ、必要ありません。** code_parallelは自動的にファイルの存在を判定し、適切な操作を選択します。
- ファイルが存在しない → 新規作成
- ファイルが存在する → 編集

### Q2: プロジェクトのコンテキストが設定されていないとどうなりますか？
**A2: 間違った技術スタックでコードが生成される可能性があります。** 例：
- Next.js プロジェクトなのに Python Flask コードが生成される
- TypeScript プロジェクトなのに JavaScript で生成される

必ず最初に `context set` でプロジェクト情報を設定してください。

### Q3: 並列実行数はどのように決めればよいですか？
**A3: 実測に基づく推奨値：**
- **軽い分析・質問**: `--parallel 5-10`（高速化効果大）
- **重い分析・deep モード**: `--parallel 1-3`（API制限を考慮）
- **コード生成・編集**: `--parallel 3-5`（安定動作確認済み）

### Q4: deep モードはいつ使うべきですか？
**A4: 複雑な分析・実装のみに使用してください：**
- ✅ セキュリティ脆弱性の詳細分析
- ✅ パフォーマンス最適化の実装
- ✅ 複雑なアルゴリズムの実装
- ❌ 簡単なコンポーネント生成
- ❌ 基本的なファイル内容確認

### Q5: エラーが発生した時の対処法は？
**A5: 主なエラーと対処法：**
```bash
# キューが空
ERROR: キューは空です
→ まず queue でタスクを追加

# APIキー期限切れ
ERROR: API key expired
→ ツール api set "new-api-key"

# ファイルが見つからない
ERROR: File not found
→ ファイルパスを確認（相対パス推奨）

# API制限エラー
ERROR: Rate limit exceeded
→ 並列数を減らして再実行
```

### Q6: 従来のReadツールとの使い分けは？
**A6: 明確な使い分け基準：**
- **単一ファイルの確認** → Readツール
- **複数ファイルの分析** → see_parallel
- **プロジェクト全体の理解** → see_parallel
- **関連ファイル群の調査** → see_parallel

### Q7: なぜ配列形式の入力なのですか？
**A7: AIの認識精度向上のため：**
- 質問・タスクとファイル名が明確に分離される
- ファイル数に制限がない（可変引数対応）
- オプション（deep等）の位置が明確
- JSON構文でパースしやすい

### Q8: コンテキスト設定は毎回必要ですか？
**A8: 一度設定すれば永続化されます：**
```bash
# 一度設定すれば自動保存
code_parallel context set "Next.js TypeScript プロジェクト"

# 確認
code_parallel context show

# 必要に応じてクリア
code_parallel context clear
```