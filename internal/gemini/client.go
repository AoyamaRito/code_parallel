package gemini

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"code_parallel/internal/config"
)

type Client struct {
	client *genai.Client
}

func NewClient() (*Client, error) {
	apiKey, err := config.GetAPIKey()
	if err != nil {
		return nil, fmt.Errorf("failed to get API key: %w", err)
	}

	if apiKey == "" {
		return nil, fmt.Errorf("API key not set. Use 'code_parallel api set <key>' to set it")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create Gemini client: %w", err)
	}

	return &Client{client: client}, nil
}

func (c *Client) Close() {
	if c.client != nil {
		c.client.Close()
	}
}

func (c *Client) GenerateCode(ctx context.Context, prompt string, useDeep bool, existingContents map[string]string) (string, error) {
	modelName := "gemini-2.0-flash-exp"
	if useDeep {
		modelName = "gemini-exp-1206"
	}

	model := c.client.GenerativeModel(modelName)
	
	// プロジェクトコンテキストを取得
	projectContext, err := config.GetContext()
	if err != nil {
		return "", fmt.Errorf("failed to get context: %w", err)
	}

	// 既存ファイルの内容を含めるかチェック
	var existingCodeSection string
	if len(existingContents) > 0 {
		existingCodeSection = "\n\n既存のコード:\n"
		for filePath, content := range existingContents {
			existingCodeSection += fmt.Sprintf("\n### %s\n```\n%s\n```\n", filePath, content)
		}
		existingCodeSection += "\n上記の既存コードを基に、タスクに従って編集・改善してください。\n"
	}

	var fullPrompt string
	if projectContext != "" {
		fullPrompt = fmt.Sprintf(`プロジェクト背景: %s
%s
タスク: %s

以下の要件に従ってコードを生成してください：
- プロジェクト背景を考慮した実装をしてください
- 実装可能で動作するコードを生成
- 適切なエラーハンドリングを含める
- コメントは日本語で記述
- ベストプラクティスに従う

コードのみを出力してください（説明文は不要）:`, projectContext, existingCodeSection, prompt)
	} else {
		fullPrompt = fmt.Sprintf(`%s
タスク: %s

以下の要件に従ってコードを生成してください：
- 実装可能で動作するコードを生成
- 適切なエラーハンドリングを含める
- コメントは日本語で記述
- ベストプラクティスに従う

コードのみを出力してください（説明文は不要）:`, existingCodeSection, prompt)
	}

	resp, err := model.GenerateContent(ctx, genai.Text(fullPrompt))
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %w", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no content generated")
	}

	content := ""
	for _, part := range resp.Candidates[0].Content.Parts {
		if text, ok := part.(genai.Text); ok {
			content += string(text)
		}
	}

	// コードフェンスを除去
	content = cleanCodeFences(content)

	return content, nil
}

func cleanCodeFences(content string) string {
	// 最初と最後のコードフェンスを削除
	if len(content) > 0 {
		lines := strings.Split(content, "\n")
		if len(lines) > 2 {
			// 最初の行がコードフェンスかチェック
			if strings.HasPrefix(lines[0], "```") {
				lines = lines[1:]
			}
			// 最後の行がコードフェンスかチェック
			if len(lines) > 0 && strings.TrimSpace(lines[len(lines)-1]) == "```" {
				lines = lines[:len(lines)-1]
			}
		}
		content = strings.Join(lines, "\n")
	}
	return content
}