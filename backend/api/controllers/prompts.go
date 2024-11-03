package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	S "worm-pack/api/services"
	_ "worm-pack/docs"
	H "worm-pack/handler"
	T "worm-pack/types"
)

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

	federalChapter, serviceErr := S.GetFederalChapter(prompt)
	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	var chapterJSON map[string]interface{}

	err := json.Unmarshal([]byte(federalChapter), &chapterJSON)
	if err != nil {
		return err
	}

	//federalArticle, serviceErr := S.GetFederalArticle(ctx.UserContext())
	//if serviceErr != nil {
	//	return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	//}

	//answer, serviceErr := S.GetFinalAnswer(ctx.UserContext())
	//if serviceErr != nil {
	//	return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	//}

	return H.Success(ctx, fiber.Map{
		"ok":             1,
		"federalChapter": chapterJSON,
	})
}
