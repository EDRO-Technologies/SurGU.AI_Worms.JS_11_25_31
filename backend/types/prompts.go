package types

type Message struct {
	IsAI    bool   `json:"is_ai"`
	Content string `json:"content"`
}

type Prompt struct {
	ActId   int       `json:"act_id"`
	Message []Message `json:"msg"`
}

type Article struct {
	Title            string `json:"title"`
	ShortDescription string `json:"short_description"`
	OriginalText     string `json:"original_text"`
	StartPageNumber  int    `json:"start_page_number"`
}

type Chapter struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Page        int       `json:"page"`
	Articles    []Article `json:"articles"`
	Keywords    []string  `json:"keywords"`
}
