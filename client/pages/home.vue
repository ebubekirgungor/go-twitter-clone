<script setup lang="ts">
import { useUser } from "@/store/user";
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
const { user } = useUser();
const { data: tweets } = useFetch<Array<Tweet>>("http://127.0.0.1/api/tweets", {
  headers: {
    Authorization: `Bearer ${user.token}`,
  },
});
</script>
<template>
  <div class="border-x border-[#2f3336]">
    <h1 v-for="tweet in tweets">
      {{ tweet.content }}
    </h1>
  </div>
</template>
