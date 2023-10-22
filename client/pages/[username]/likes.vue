<script setup lang="ts">
import { useUser } from "@/store/user";
const { user } = useUser();
definePageMeta({
  layout: "profile",
  middleware: "auth",
});
useHead({
  title: `Posts liked by ${user.name} (${user.username}) / Twitter`,
});
</script>
<template>
  <div>
    <h1>test</h1>
  </div>
</template>
