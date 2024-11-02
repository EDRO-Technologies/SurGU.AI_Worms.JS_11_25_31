<template>
  <div>
    <iframe :src="pdfUrl" width="100%" height="680"></iframe>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watch } from "vue";

const props = defineProps<{
  page: number;
}>();

const pageNum = ref(0);

const pdfUrl = computed<string>(() => {
  const baseUrl = new URL("@/assets/source.pdf", import.meta.url).href;
  return `${baseUrl}#page=${pageNum.value}`;
});

watch(
  () => props.page,
  () => {
    pageNum.value = props.page;
  },
);
</script>

<style scoped></style>
