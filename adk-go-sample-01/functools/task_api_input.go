package functools

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
