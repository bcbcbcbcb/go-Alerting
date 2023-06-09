package services

type Hook struct {
	EvalMatches []EvalMatches `json:"evalMatches"`
	ImageURL    string        `json:"imageUrl"`
	Message     string        `json:"message"`
	RuleID      int           `json:"ruleId"`
	RuleName    string        `json:"ruleName"`
	RuleURL     string        `json:"ruleUrl"`
	State       string        `json:"state"`
	Tags        Tags          `json:"tags"`
	Title       string        `json:"title"`
	ID          int
}
type Tags struct {
}

type EvalMatches struct {
	Value  int         `json:"value"`
	Metric string      `json:"metric"`
	Tags   interface{} `json:"tags"`
}
