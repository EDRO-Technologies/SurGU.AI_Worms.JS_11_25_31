<template>
  <div class="flex gap-1">
    <div class="flex-1 relative border border-gray-300 py-2 pl-4 pr-2.5 rounded-[40px] flex">
      <input
        v-model="message"
        type="text"
        class="w-full outline-none"
        placeholder="Спросите wormy..."
      />
      <div class="flex" v-loading="isLoading">
        <el-button type="primary" circle @click="sendChats" :icon="ArrowRightBold" />
        <el-button type="danger" circle @click="clearChat" :icon="CloseBold" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { ArrowRightBold, CloseBold } from "@element-plus/icons-vue";
import { sendAllMessages } from "@/service";
import { CHATS, clearHistory } from "@/stores/chat";
import { ElMessage } from "element-plus";


const message = ref("");
const isLoading = ref(false);

const sendChats = async () => {
  if (!message.value.trim()) return;

  const userMessage = {
    role: "user",
    content: message.value,
  };

  CHATS.value.push(userMessage);
  isLoading.value = true;
  let chatMessage = await sendAllMessages(CHATS.value);
  CHATS.value.push(chatMessage);
  isLoading.value = false;

  message.value = "";
};

const clearChat = () => {
  clearHistory();
  ElMessage({
    message: "Память очищена",
    type: "error",
  });
};
</script>
