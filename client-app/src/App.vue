<template>
  <div class="h-screen w-full max-w-4xl mx-auto p-4 pt-6 flex flex-col justify-between">
    <div class="absolute left-10 top-1 font-bold">WormsAI</div>
    <el-tabs v-model="activeName" stretch>
      <el-tab-pane label="Чат" name="first">
        <ChatsContainer
          class="flex-1 p-2 h-[80vh] overflow-scroll"
          @navigate-to-pdf="navigateToPdf"
        />
        <ChatBox />
      </el-tab-pane>
      <el-tab-pane label="Документ" name="second">
        <PdfView :page="computedPdfPage" />
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script setup lang="ts">
import ChatBox from "./components/ChatBox.vue";
import ChatsContainer from "./components/ChatsContainer.vue";
import { computed, ref } from "vue";
import PdfView from "@/components/PdfView.vue";

const activeName = ref("first");
const pdfPage = ref(1);

const computedPdfPage = computed({
  get: () => pdfPage.value,
  set: (newValue) => pdfPage.value = newValue,
})

const navigateToPdf = (pageNum: number) => {
  activeName.value = "second";
  computedPdfPage.value = pageNum;
  console.log(computedPdfPage.value);
};
</script>
