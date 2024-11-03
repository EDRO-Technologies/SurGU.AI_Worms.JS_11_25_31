package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"
	"time"
	C "worm-pack/config"
	T "worm-pack/types"
)

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

type DimaRequestMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type DimaRequest struct {
	Model    string               `json:"model"`
	Messages []DimaRequestMessage `json:"messages"`
}

func MakeDimaRequest(msgs []DimaRequestMessage) *DimaRequest {
	return &DimaRequest{
		Model:    "Vikhrmodels/Vikhr-Llama-3.2-1B-Instruct",
		Messages: msgs,
	}
}

func BoolRoleToString(isAI bool) string {
	if isAI {
		return "assistant"
	} else {
		return "user"
	}
}

func GetAnswered(prompt T.Prompt) (string, *T.ServiceError) {
	if len(prompt.Message) == 0 {
		return "", nil
	}

	msgs := Map(prompt.Message, func(item T.Message) DimaRequestMessage {
		return DimaRequestMessage{
			Role:    BoolRoleToString(item.IsAI),
			Content: item.Content,
		}
	})

	req := MakeDimaRequest(msgs)

	reqJSON, err := json.Marshal(req)
	if err != nil {
		return "", &T.ServiceError{
			Message: "Косяк перепаковки реквеста",
			Error:   err,
			Code:    fiber.StatusInternalServerError,
		}
	}

	reqJSONBytes := bytes.NewBuffer(reqJSON)

	apiURL := fmt.Sprintf("%s", C.Conf.ModelHost)

	promptReq, err := http.NewRequest("POST", apiURL, reqJSONBytes)
	if err != nil {
		return "", &T.ServiceError{
			Message: "Косяк в реквесте",
			Error:   err,
			Code:    fiber.StatusInternalServerError,
		}
	}

	promptReq.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: 80 * time.Second}

	res, err := client.Do(promptReq)

	if err != nil {
		return "", &T.ServiceError{
			Message: "Ploho",
			Error:   err,
			Code:    fiber.StatusInternalServerError,
		}
	}

	if res.Body == nil {
		return "", &T.ServiceError{
			Message: "Flawed response from LLM",
		}
	}

	body, err := io.ReadAll(res.Body)

	responseBody := string(body)

	if err != nil {
		return "", &T.ServiceError{
			Message: "Unable to get federal act",
			Error:   err,
			Code:    fiber.StatusInternalServerError,
		}
	}

	return responseBody, nil
}

//func GetFederalArticle(ctx context.Context, prompt *T.Prompt) (*T.Prompt, *T.ServiceError) {
//	federalChapter, err := //TODO: second API call
//
//	if err != nil {
//		return nil, &T.ServiceError{
//			Message: "Unable to get products",
//			Error:   err,
//			Code:    fiber.StatusInternalServerError,
//		}
//	}
//
//	return products, nil
//}
//
//func GetFinalAnswer(ctx context.Context) (*T.Prompt, *T.ServiceError) {
//	//federalAct, err := http.//TODO: first API call
//
//	if err != nil {
//		return nil, &T.ServiceError{
//			Message: "Unable to get federal act",
//			Error:   err,
//			Code:    fiber.StatusInternalServerError,
//		}
//	}
//
//	return federalAct, nil
//}
