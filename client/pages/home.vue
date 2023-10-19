<script setup lang="ts">
import { useToken } from "@/store/token";
definePageMeta({
  layout: "app",
  middleware: "auth",
});
useHead({
  title: "Home / Twitter",
});
interface Tweet {
  content: string;
}
const { token } = useToken();
const { data: tweets } = useFetch<Array<Tweet>>("http://127.0.0.1/api/tweets", {
  headers: {
    Authorization: `Bearer ${token}`,
  },
});
</script>
<template>
  <div>
    <h1 v-for="tweet in tweets">
      {{ tweet.content }}
    </h1>
  </div>
</template>
