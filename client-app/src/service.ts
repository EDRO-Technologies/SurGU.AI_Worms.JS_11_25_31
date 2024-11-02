import type { CHAT } from "./types";
import axios from "axios";


export const sendAllMessages = async (chats: unknown) => {
  console.log(chats);
  console.log(import.meta.env.URL_ADDRESS);
  const headers = new axios.AxiosHeaders(
    {
      dangerouslyAllowBrowser: true,
    },
  );
  const requestBody = {
    model: import.meta.env.VITE_MODEL_TYPE,
    messages: chats,
  };
  const { data } = await axios.post(import.meta.env.VITE_SERVER_ADDRESS, requestBody);

  let responseMessage: CHAT = {
    role: data.choices[0].message.role,
    content: data.choices[0].message.content!,
  };

  return responseMessage;
};
