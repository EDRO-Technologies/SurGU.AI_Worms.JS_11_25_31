<template>
  <div class="flex gap-1 mb-2" :class="{ 'justify-end': isUser }">
    <div
      class="flex flex-col gap-1 p-2 rounded-xl"
      :class="{ 'bg-gray-200': isUser }"
    >
      <div v-if="!isUser" class="font-bold">Wormy</div>
      <div v-html="markdownContent"></div>
      <el-button type="primary" @click="navigateInPdf(30)">Открыть в документе</el-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue";
import { marked } from "marked";

const props = defineProps<{
  content: string;
  role: string;
}>();

const emit = defineEmits<{
  (e: "navigateToPdf", pageNum: number): void;
}>();

const isUser = computed(() => {
  return props.role === "user";
});

const navigateInPdf = (pageNum: number) => {
  emit("navigateToPdf", pageNum);
};

const markdownContent = computed(() => marked(props.content));
</script>
