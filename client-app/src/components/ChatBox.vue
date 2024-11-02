<template>
  <div class="flex gap-1">
    <div class="flex-1 relative border border-gray-300 py-2 pl-4 pr-2.5 rounded-[40px] flex">
      <input
        v-model="message"
        type="text"
        class="w-full outline-none"
        placeholder="Спросите wormy..."
      />
      <div class="flex gap-2">
        <el-button type="info" circle @click="sendChats" :icon="ArrowRightBold">
        </el-button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { ArrowRightBold } from "@element-plus/icons-vue";
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
};
</script>
