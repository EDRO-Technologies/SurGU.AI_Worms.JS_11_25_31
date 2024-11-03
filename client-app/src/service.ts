import type { CHAT } from "./types";
import axios from "axios";


interface ApiRequestMessage {
  is_ai: boolean,
  content: string,
}

interface ApiRequestPrompt {
  act_id: number,
  msg: ApiRequestMessage[]
}

interface ApiResponse {
  federalChapter: {
    answer: string,
    article_name: string,
    chapter_name: string,
    page_number: number
  }
}

const URL = `${import.meta.env.VITE_BACKEND_URL}/api/prompt`;

export const sendAllMessages = async (chats: unknown) => {
  const requestBody = {
    act_id: 0,
    msg: chats.map(x => ({ is_ai: x.role === "assistant", content: x.content } as ApiRequestMessage))
  } as ApiRequestPrompt;

  const { data } = await axios.post<ApiResponse>(URL, requestBody);

  let responseMessage: CHAT = {
    role: "assistant",
    content: data.federalChapter.answer,
    pageNumber: data.federalChapter.page_number,
  };

  return responseMessage;
};
