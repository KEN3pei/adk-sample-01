package functools

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/functiontool"
)

type HTTPClient struct {
	client http.Client
}

type FindTaskIdInput struct {
	Name   *string `json:"name", jsonschema:"todo task name. for used search taskID"`
	UserId int     `json:"user_id", jsonschema:"Allow task operation user"`
}

type FindTaskIdOutputRecord struct {
	TaskId      int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type FindTaskIdResponse struct {
	Data []FindTaskIdOutputRecord `json:"result"`
}

func NewFunctionTools() []tool.Tool {
	// Add the tool to the agent
	taskTool01, err := functiontool.New(
		functiontool.Config{
			Name:        "get_taskId_by_input_parameter",
			Description: "Retrieves the todo taskId by input parameter",
		},
		FindTaskIdByInput,
	)
	if err != nil {
		log.Fatal(err)
	}

	return []tool.Tool{taskTool01}
}

func FindTaskIdByInput(ctx tool.Context, input FindTaskIdInput) (FindTaskIdResponse, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, FindTaskInputToRequestURL(input), nil)
	if err != nil {
		return FindTaskIdResponse{}, err
	}
	res, err := client.Do(req)
	if err != nil {
		return FindTaskIdResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return FindTaskIdResponse{}, err
	}
	fmt.Println("[FindTaskIdByInput] response body:", string(body))

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
