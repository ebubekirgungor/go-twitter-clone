<script setup lang="ts">
definePageMeta({
  layout: "app",
  middleware: "auth"
});
useHead({
  title: "Profile / Twitter",
});
</script>
<template>
  <div>
    <h1>
      profile
    </h1>
  </div>
</template>