<template>
  <div>
    <iframe :src="pdfUrl" width="100%" height="680"></iframe>
  </div>
</template>

<script setup lang="ts">
const pdfUrl = new URL('@/assets/source.pdf', import.meta.url).href;
</script>

<style scoped></style>
