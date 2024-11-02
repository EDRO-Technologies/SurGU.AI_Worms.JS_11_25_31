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

func GetFederalChapter(prompt T.Prompt) (string, *T.ServiceError) {
	req := MakeDimaRequest([]DimaRequestMessage{
		{
			Role:    BoolRoleToString(prompt.Message[0].IsAI),
			Content: prompt.Message[0].Content,
		},
	})

	_ = req

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

	client := http.Client{Timeout: 3 * time.Second}

	res, err := client.Do(promptReq)

	defer res.Body.Close()

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
