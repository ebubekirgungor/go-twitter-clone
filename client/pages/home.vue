<script setup lang="ts">
import { useToken } from "@/store/token";
definePageMeta({
  layout: "app",
});
useHead({
  title: "Home / Twitter",
});
interface Tweet {
  content: string;
}
const { token } = useToken();

const { data: tweets, refresh } = await useAsyncData("tweets", () => {
  $fetch("/api/tweets", {
    method: "get",
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  
});
refresh();
console.log(token);
console.log(tweets);
</script>
<template>
  <div>
    <h1 v-for="tweet in tweets">
      {{ tweet.content }}
    </h1>
  </div>
</template>
