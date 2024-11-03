<template>
  <el-progress  v-if="isLoading" class="mb-1 mx-4" :duration="2" :show-text="false" :percentage="100" :indeterminate="true" />
  <div v-else class="h-[6px]"/>
  <div class="flex gap-1">
    <div class="flex-1 flex-row relative border border-gray-300 py-2 pl-4 pr-2.5 rounded-[10px] flex">
      <input
        v-model="message"
        type="text"
        class="w-full outline-none"
        placeholder="Спросите wormy..."
        @keydown="onEnter"
      />
      <div disabled="isLoading" class="flex flex-row">
        <el-button type="danger" circle @click="clearChat" :icon="CloseBold" />
        <el-button type="primary" circle @click="sendChats" :icon="ArrowRightBold" />
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

const delay = (delayInms) => {
  return new Promise(resolve => setTimeout(resolve, delayInms));
};

const sendChats = async () => {
  if (!message.value.trim()) return;

  const userMessage = {
    role: "user",
    content: message.value,
  };

  const div = document.querySelector("#chatsContainer");

  message.value = "";

  const pushAndScroll = async (msg) => {
    CHATS.value.push(msg);
    await delay(10);
    div.scrollTop = div.scrollHeight;
  };

  await pushAndScroll(userMessage);

  isLoading.value = true;
  const chatMessage = await sendAllMessages(CHATS.value);
  isLoading.value = false;

  await pushAndScroll(chatMessage);
};

const clearChat = () => {
  clearHistory();
  ElMessage({
    message: "Память очищена",
    type: "error",
  });
};

const onEnter = async (e) => {
  if (e.key !== "Enter") {
    return;
  }

  await sendChats();
};
</script>
