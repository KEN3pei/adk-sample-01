### sample agents

<https://github.com/google/adk-samples/tree/main/python/agents>

### Function tools

<https://google.github.io/adk-docs/tools-custom/function-tools/>

API呼び出しを、OpenAPIToolsetでできそう

- <https://future-architect.github.io/articles/20250723a/#API%E3%83%84%E3%83%BC%E3%83%AB%E3%81%AE%E5%AE%9F%E8%A3%85%E6%89%8B%E9%A0%86>
- <https://google.github.io/adk-docs/tools-custom/openapi-tools/>

### Goで始めるには？

https://google.github.io/adk-docs/get-started/go/

CLI
- `go run agent.go`

web interface
- `go run agent.go web api webui`

### Function tools

GoでFunction toolを定義する場合、パラメータを明確にするため構造体タグを使ってスキーマを定義する。

```
```
// GetWeatherParams defines the arguments for the getWeather tool.
type GetWeatherParams struct {
    // This field is REQUIRED (no "omitempty").
    // The jsonschema tag provides the description.
    Location string `json:"location" jsonschema:"The city and state, e.g., San Francisco, CA"`

    // This field is also REQUIRED.
    Unit     string `json:"unit" jsonschema:"The temperature unit, either 'celsius' or 'fahrenheit'"`
}
```
```


optionalなパラメータは`omitempty`などを使う。

> LLMから期待するすべてのデータについては、明示的に定義されたパラメータに依存するのが最善です。

> 戻り値は可能な限り説明的なものにするよう努めてください。例えば、数値のエラーコードを返す代わりに、人間が理解できる説明を含む「error_message」キーを持つ辞書を返します。結果を理解する必要があるのはコードではなくLLMであることを忘れないでください。ベストプラクティスとして、返却辞書に「status」キーを含め、全体的な結果（例：「success」「error」「pending」）を示すことで、LLMに操作の状態を明確に伝達してください。

-> つまりFunction ToolでRestAPIを呼び出して返すなどする場合、その返り値はLLM向けのコンテキストが多い設計とすべきということ。




