<template>
  <div class="flex gap-1">
    <div class="flex-1 relative border-2 border-gray-300 p-3 rounded-3xl flex">
      <input
        v-model="message"
        type="text"
        class="w-full outline-none"
        placeholder="Message ChatGPT..."
      />
      <div class="flex gap-2">
        <button @click="sendChats">➤</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { sendAllMessages } from "@/service";
import { CHATS } from "@/stores/chat";

const message = ref("");

const sendChats = async () => {
  const userMessage = {
    role: "user",
    content: message.value,
  };
  CHATS.value.push(userMessage);

  let chatGPTMessage = await sendAllMessages(CHATS.value);
  CHATS.value.push(chatGPTMessage);
}
</script>
