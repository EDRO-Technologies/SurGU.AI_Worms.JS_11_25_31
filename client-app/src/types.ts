export type CHAT = {
  role: string;
  content: string;
  pageNumber: number
};

export type RESPONSE = {
  id: string,
  object: string,
  created: number,
  model: string,
  choices: [
    {
      index: number,
      message: {
        role: string,
        content: string,
        tool_calls: string[]
      },
      logprobs: any,
      finish_reason: string,
      stop_reason: any
    }
  ],
  usage: {
    prompt_tokens: number,
    total_tokens: number,
    completion_tokens: number
  },
  prompt_logprobs: any
}
