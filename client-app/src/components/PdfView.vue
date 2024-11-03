<template>
  <iframe :key="pdfUrl" :src="pdfUrl" width="100%" height="100%"></iframe>
</template>

<script setup lang="ts">
import { computed } from "vue";

const props = defineProps<{
  page: number;
}>();


const pdfUrl = computed<string>(() => {
  const baseUrl = new URL("@/assets/source.pdf", import.meta.url).href;
  const timestamp = Date.now();

  return `${baseUrl}?t=${timestamp}#page=${props.page}`;
});

</script>

<style scoped></style>
