package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	S "worm-pack/api/services"
	_ "worm-pack/docs"
	H "worm-pack/handler"
	T "worm-pack/types"
)

type message struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

type choice struct {
	Message message `json:"message"`
}

type federalChapter struct {
	Choices []choice `json:"choices"`
	ID      string   `json:"id"`
}

type ApiResponse struct {
	Answer      string `json:"answer"`
	ArticleName string `json:"article_name"`
	ChapterName string `json:"chapter_name"`
	PageNumber  int    `json:"page_number"`
}

func MakeApiRequest(answer federalChapter) *ApiResponse {
	return &ApiResponse{
		Answer: answer.Choices[0].Message.Content,
	}
}

// PostPrompt
// @Summary        PostPrompt
// @Description    Post prompt
// @Tags           Prompts
// @Accept         json
// @Produce        json
// @Param          prompt  body    types.Prompt   true    "prompt"  example({"model": "text-davinci-003", "messages": [{"role": "user", "content": "Hello, how are you?"}]})
// @Success        200     {object}  map[string]interface{}
// @Router         /api/prompt [post]
func PostPrompt(ctx *fiber.Ctx) error {
	var prompt T.Prompt

	if err := ctx.BodyParser(&prompt); err != nil {
		return H.BuildError(ctx, "Invalid JSON", fiber.StatusBadRequest, err)
	}

	response, serviceErr := S.GetAnswered(prompt)
	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	var content federalChapter

	err := json.Unmarshal([]byte(response), &content)
	if err != nil {
		return err
	}

	var apiResponse = MakeApiRequest(content)

	return H.Success(ctx, fiber.Map{
		"ok":             1,
		"federalChapter": apiResponse,
	})
}

//
//// PingPong
//// @Summary        PingPong
//// @Description    Post prompt
//// @Tags           Prompts
//// @Accept         json
//// @Produce        json
//// @Param          prompt  body    types.Prompt   true    "prompt"  example({"model": "text-davinci-003", "messages": [{"role": "user", "content": "Hello, how are you?"}]})
//// @Success        200     {object}  map[string]interface{}
//// @Router         /api/legendaryPrompt [post]
//func PingPong(ctx *fiber.Ctx) error {
//	var prompt T.Prompt
//
//	if err := ctx.BodyParser(&prompt); err != nil {
//		return H.BuildError(ctx, "Invalid JSON", fiber.StatusBadRequest, err)
//	}
//
//	response, serviceErr := S.GetAnswered(prompt)
//	if serviceErr != nil {
//		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
//	}
//
//	var content federalChapter
//
//	err := json.Unmarshal([]byte(response), &content)
//	if err != nil {
//		return err
//	}
//
//	var chapter = MakeApiRequest(content).
//
//	//federalArticle, serviceErr := S.GetFederalArticle(ctx.UserContext())
//	//if serviceErr != nil {
//	//	return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
//	//}
//
//	//answer, serviceErr := S.GetFinalAnswer(ctx.UserContext())
//	//if serviceErr != nil {
//	//	return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
//	//}
//
//	return H.Success(ctx, fiber.Map{
//		"ok":             1,
//		"federalChapter": apiResponse,
//	})
//}
