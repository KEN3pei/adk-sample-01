### ADK

https://pkg.go.dev/google.golang.org/adk

## packages

### agent
- llmagent: LLM ベースのエージェントを提供します。
  - パッケージllmagentはLLMベースのエージェントを提供します。
  - LLMエージェントは、大規模言語モデルを用いて、指示やユーザー入力に基づいてタスクを実行し、取るべき行動を決定し、利用可能なツールを使用して行動を実行するか、サブエージェントに委任します。

- remoteagent: リモート ADK エージェントを使用できます。
- workflowagents/loopagent: 指定された回数だけ、または終了条件が満たされるまでサブエージェントを繰り返し実行するエージェントを提供します。
- workflowagents/parallelagent: サブエージェントを並列に実行するエージェントを提供します。
- workflowagents/sequentialagent: サブエージェントを順番に実行するエージェントを提供します。

リモートADKエージェントとA2Aとは(=Agent to Agent)
- NewA2AはリモートA2Aエージェントを作成します。A2A（エージェント間通信）プロトコルは、別プロセスまたは別ホスト上で実行可能なエージェントとの通信に使用されます。

サブエージェントとは？
- メインのエージェント（親エージェント）が大きなタスクを受け持ち、その一部のサブタスクを別のエージェントに渡して処理させるような構造になっている。
- Root Agent → Sub Agents
  - Root Agent は：
    - タスク分解
    - どのサブエージェントに投げるか判断
    - 結果を統合
    を担当します。

### artifact: (アーティファクトを管理するためのサービスを提供します)
- gcsartifact: Google Cloud Storage (GCS) のアーティファクト.Service を提供します。
  - このパッケージは、GCS バケットへのアーティファクトの保存と取得を可能にします。アーティファクトはアプリケーション名、ユーザー ID、セッション ID、ファイル名で整理され、バージョン管理をサポートします。

### cmd
- adkgo: adkgo は、ADKアプリケーションのデプロイとテストを支援する CLIツール
- launcher: Package launcherは、エージェントと対話する方法を提供します。
- launcher/console: コンソール アプリケーションからエージェントと対話する簡単な方法を提供します。
- launcher/full: パッケージの完全版は、利用可能なすべてのオプションを備えたADKで簡単にプレイする方法を提供します
- launcher/prod:
- launcher/universal:
- launcher/web: 
- launcher/weba2a: 
- launcher/web/api
- launcher/web/webui


### example(とりあえず見ておくと便利そうなexample)
- workflowagents/loop: ループ エージェントを実行するワークフロー エージェント
- workflowagents/parallel: サブエージェントを並列に実行するワークフロー エージェント
- workflowagents/sequential: サブエージェントを順番に実行するワークフロー エージェント

ループエージェントとは？
- サブエージェントを**反復実行（ループ）**するためのもの
- 指定した 回数だけ繰り返すか、終了条件が満たされるまで同じ一連のサブエージェントを何度も動かす。
- あくまでワークフロー制御のためのオーケストレーションのみ担当する。

### memory
- エージェント メモリ (長期的な知識) と対話するエンティティを定義します。
  - ここでいうメモリとは？
    - memory packageは、In-MemoryDB,RDB,VectorDB,Cloud RAG のどれでも使えるようにintererfaceやヘルパメソッドなどを用意している。
    - 過去の会話やイベントを保持 → 次回以降の対話でコンテキストとして活用するための仕組み
    - ユーザーの好みを覚えて将来の応答に反映することができる。


### model
- geminii: Gemini モデルの model.LLM インターフェースを実装します。

### runner
- ADK エージェントのランタイムを提供します。

### server
- adka2a: A2A 経由で ADK エージェントを公開できます。
- adkrest: パッケージ コントローラーには、ADK REST API のコントローラーが含まれています。
- adkrest/controllers

※a2a=Agent to Agent
- NewA2AはリモートA2Aエージェントを作成します。A2A（エージェント間通信）プロトコルは、別プロセスまたは別ホスト上で実行可能なエージェントとの通信に使用されます。

### session
- database: ユーザー セッションとその状態を管理するための型を提供します。

### telemetry
- ADK イベントが発行されるカスタム テレメトリ プロセッサを設定できます。

### tool: エージェントによって呼び出すことができるツールのインターフェイスを定義します。
- agenttool: エージェントが別のエージェントを呼び出すことを可能にするツールを提供します。
- exitlooptool: エージェントがループを終了できるようにするツールを提供します。
- functiontool: Go 関数をラップするツールを提供します。
- geminitool: Gemini ネイティブ ツールへのアクセスを提供します。
- loadartifactstool: アーティファクトをロードするためのツールを定義します。
- mcptoolset: MCP ツール セットを提供します。

### util
- instructionutil: エージェントの指示を操作するためのユーティリティを提供します。


## LLMエージェント

https://google.github.io/adk-docs/agents/llm-agents/

```go
// Example: Defining the basic identity
agent, err := llmagent.New(llmagent.Config{
    Name:        "capital_agent",
    Model:       model,
    Description: "Answers user questions about the capital city of a given country.",
    // instruction and tools will be added next
})
```
```
```

- name(required): 各エージェントには一意の文字列識別子が必要。

- description(option): 

> エージェントの能力について簡潔な概要を記入してください。この説明は主に、 他のLLMエージェントがこのエージェントにタスクをルーティングするかどうかを判断する際に使用され
他のエージェントと区別できるよう、具体的に記入してください（例：「請求エージェント」ではなく、「最新の請求明細に関する問い合わせに対応」など）。

- model(required): このエージェントの推論を支えるための基盤、LLMを指定する。

### Guiding the Agent: Instructions 

https://google.github.io/adk-docs/agents/llm-agents/#guiding-the-agent-instructions-instruction

Instructionパラメータ
-> Agentの動きを決めるための制約や、役割、性格、などを規定するパラメータであり以下ような部分を決定する。

- その中核となるタスクまたは目標。
- その性格またはペルソナ (例:「あなたは役に立つアシスタントです」、「あなたは機知に富んだ海賊です」)。
- 動作に対する制約 (例: 「X に関する質問にのみ回答する」、「Y を決して公開しない」)。
- どのように、いつ使用するかtools。各ツールの目的と、それを呼び出すべき状況について、ツール自体の説明を補足しながら説明する必要があります。
- 出力の希望する形式 (例: 「JSON で応答する」、「箇条書きのリストを提供する」)。

あとInstructionには{var}構文で動的な値を設定できる。

```go
    // Example: Adding instructions
    agent, err := llmagent.New(llmagent.Config{
        Name:        "capital_agent",
        Model:       model,
        Description: "Answers user questions about the capital city of a given country.",
        Instruction: `You are an agent that provides the capital city of a country.
When a user asks for the capital of a country:
1. Identify the country name from the user's query.
2. Use the 'get_capital_city' tool to find the capital.
3. Respond clearly to the user, stating the capital city.
Example Query: "What's the capital of {country}?"
Example Response: "The capital of France is Paris."`,
        // tools will be added next
    })
```
```
```


### tools

LLMは、会話と指示に基づいて、関数/ツールの名前、説明 (docstringsまたはdescriptionフィールドから取得)、およびパラメータスキーマを使用して、どのツールを呼び出すかを決定します。

```go
// 何かしらの関数を用意する
getCapitalCity := func(){...}

// Add the tool to the agent
capitalTool, err := functiontool.New(
    functiontool.Config{
        Name:        "get_capital_city",
        Description: "Retrieves the capital city for a given country.",
    },
    getCapitalCity,
)
if err != nil {
    log.Fatal(err)
}
agent, err := llmagent.New(llmagent.Config{
    Name:        "capital_agent",
    Model:       model,
    Description: "Answers user questions about the capital city of a given country.",
    Instruction: "You are an agent that provides the capital city of a country... (previous instruction text)",
    Tools:       []tool.Tool{capitalTool},
})
```
```

```

## Sessions
