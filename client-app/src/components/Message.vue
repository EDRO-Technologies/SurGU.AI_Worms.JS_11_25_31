<template>
  <div class="flex gap-1" :class="{ 'justify-end': isUser }">
    <div class="flex flex-col gap-1">
      <div v-if="isUser" class="font-bold self-end">You</div>
      <div v-else class="font-bold">Вихрь</div>
      <div>{{ content }}</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue";

const props = defineProps<{
  content: string;
  role: string;
}>();

const isUser = computed(() => {
  return props.role === "user";
});
</script>
