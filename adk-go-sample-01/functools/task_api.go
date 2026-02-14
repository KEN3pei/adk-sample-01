package functools

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"time"

	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/functiontool"
)

type TaskAPIClient struct {
	client *http.Client
}

func NewTaskAPIClient() *TaskAPIClient {
	return &TaskAPIClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func NewFunctionTools() []tool.Tool {
	apiClient := NewTaskAPIClient()
	// Add the tool to the agent
	taskTool01, err := functiontool.New(
		functiontool.Config{
			Name:        "get_taskId_by_input_parameter",
			Description: "Retrieves the todo taskId by input parameter",
		},
		apiClient.FindTaskIdByInput,
	)
	if err != nil {
		log.Fatal(err)
	}

	return []tool.Tool{taskTool01}
}

func (tc *TaskAPIClient) FindTaskIdByInput(ctx tool.Context, input FindTaskIdInput) (FindTaskIdResponse, error) {
	req, err := http.NewRequest(http.MethodGet, FindTaskInputToRequestURL(input), nil)
	if err != nil {
		return FindTaskIdResponse{}, err
	}
	res, err := tc.client.Do(req)
	if err != nil {
		return FindTaskIdResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return FindTaskIdResponse{}, err
	}
	slog.Info("status:", res.StatusCode)

	// 非2xxの場合は本文を含めてエラー化
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return FindTaskIdResponse{}, fmt.Errorf("unexpected status %d: %s", res.StatusCode, string(body))
	}

	// 常にトップレベル配列を前提にデコード
	var findTaskIdResponse FindTaskIdResponse
	if err := json.Unmarshal(body, &findTaskIdResponse); err != nil {
		return FindTaskIdResponse{}, err
	}
	return findTaskIdResponse, nil
}

func FindTaskInputToRequestURL(input FindTaskIdInput) string {
	url := fmt.Sprintf("https://47f933ce-a83e-4887-b3d2-498635479533.mock.pstmn.io/agent/v1/tasks/search?user_id=%d", input.UserId)

	if input.Name != nil {
		url += fmt.Sprintf("&name=%s", *input.Name)
	}

	return url
}
