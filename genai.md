### Google Gen AI SDK

https://blog.g-gen.co.jp/entry/migrate-from-vertex-ai-sdk-to-google-gen-ai-sdk

https://pkg.go.dev/google.golang.org/genai

Google の生成モデルを Go アプリケーションに統合するためのインターフェースを提供します。
Gemini Developer APIとVertex AI APIをサポートしている。

### 実装

Gemini API Client:

```golang
```
client, err := genai.NewClient(ctx, &genai.ClientConfig{
	APIKey:   apiKey,
	Backend:  genai.BackendGeminiAPI,
})
```
```

Vertex AI Client:

```golang
```
client, err := genai.NewClient(ctx, &genai.ClientConfig{
	Project:  project,
	Location: location,
	Backend:  genai.BackendVertexAI,
})
```

### API Key

```shell
```
Gemini Developer API: Set GOOGLE_API_KEY as shown below:

export GOOGLE_API_KEY='your-api-key'
Gemini API on Vertex AI: Set GOOGLE_GENAI_USE_VERTEXAI, GOOGLE_CLOUD_PROJECT and GOOGLE_CLOUD_LOCATION, as shown below:

export GOOGLE_GENAI_USE_VERTEXAI=true
export GOOGLE_CLOUD_PROJECT='your-project-id'
export GOOGLE_CLOUD_LOCATION='us-central1'
```


