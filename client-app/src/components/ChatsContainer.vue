<template>
  <div ref="chatsContainer">
    <Message
      v-for="(chat, index) of CHATS"
      :key="index"
      :content="chat.content"
      :role="chat.role"
      @navigate-to-pdf="navigateToPdf"
    />
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue";
import Message from "./Message.vue";
import { CHATS } from "@/stores/chat";

const emit = defineEmits<{
  (e: "navigateToPdf", pageNum: number): void;
}>();

const chatsContainer = ref<HTMLElement | null>(null);

const navigateToPdf = (pageNum: number) => {
  emit("navigateToPdf", pageNum);
};

onMounted(() => {
  chatsContainer.value?.scrollTo(0, chatsContainer.value.scrollHeight);
});
</script>
