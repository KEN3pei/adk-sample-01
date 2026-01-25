package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"adk-go-sample/functools"
	"adk-go-sample/session_manages"

	"github.com/joho/godotenv"
	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/cmd/launcher"
	"google.golang.org/adk/cmd/launcher/web"
	"google.golang.org/adk/cmd/launcher/web/api"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/adk/session"
	"google.golang.org/genai"
)

func main() {
	ctx := context.Background()
	err := godotenv.Load(".env")
	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	model, err := gemini.NewModel(ctx, "gemini-3-pro-preview", &genai.ClientConfig{
		APIKey: os.Getenv("GOOGLE_API_KEY"),
	})
	if err != nil {
		log.Fatalf("Failed to create model: %v", err)
	}

	timeAgent, err := llmagent.New(llmagent.Config{
		Name:        "todo_app_operation_agent",
		Model:       model,
		Description: "An Agent that calls the REST API for Todo app operations and implements the given Todo app operations.",
		Instruction: `You are a workflow agent that calls the Todo list RestAPI.
Please operate the ToDo app as requested by calling the RestAPI as needed.
Each tool has the following specifications:
get_taskId_by_input_parameter: This API is used to get the task ID of the target task when an operation request is made with a specific task name.`,
		Tools: functools.NewFunctionTools(),
	})
	if err != nil {
		log.Fatalf("Failed to create agent: %v", err)
	}

	// /sessions 用と /api/run 用で同じ SessionService インスタンスを共有する
	sessionSvc := session.InMemoryService()

	config := &launcher.Config{
		AgentLoader:    agent.NewSingleLoader(timeAgent),
		SessionService: sessionSvc,
	}

	l := web.NewLauncher(
		api.NewLauncher(),
		&session_manages.SessionLauncher{
			SessionService: sessionSvc,
		},
	)
	// コマンドライン引数から有効なサブランチャー(api, session)を決定する
	remaining, err := l.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("Failed to parse args: %v\n\n%s", err, l.CommandLineSyntax())
	}
	if len(remaining) != 0 {
		log.Fatalf("Unparsed arguments: %v\n\n%s", remaining, l.CommandLineSyntax())
	}

	if err = l.Run(ctx, config); err != nil {
		log.Fatalf("Run failed: %v\n\n%s", err, l.CommandLineSyntax())
	}
}
