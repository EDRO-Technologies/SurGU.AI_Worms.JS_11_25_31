import type { CHAT } from "@/types";
import { useLocalStorage } from "@vueuse/core";

export const CHATS = useLocalStorage<CHAT[]>("CHATS", [
  { role: "assistant", content: "Привет, меня зовут wormy, я готов вам помочь с документом!" },
]);

export const clearHistory = () => {
  CHATS.value = [
    { role: "assistant", content: "Привет, меня зовут wormy, я готов вам помочь с документом!" }
  ];
};
