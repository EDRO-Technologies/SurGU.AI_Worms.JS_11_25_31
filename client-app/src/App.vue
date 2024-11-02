<template>
  <div class="h-screen w-full max-w-4xl mx-auto p-4 pt-12 flex flex-col justify-between">
    <div class="absolute left-10 top-1 font-bold">WormsAI</div>
    <ChatsContainer class="flex-1 p-2 overflow-auto" />
    <ChatBox />
  </div>
</template>

<script setup lang="ts">
import ChatBox from "./components/ChatBox.vue";
import ChatsContainer from "./components/ChatsContainer.vue";
</script>
