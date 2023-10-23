<script setup lang="ts">
const route = useRoute();
definePageMeta({
  layout: "profile",
  middleware: "auth",
});
const { data: user } = useFetch<User>(
  "http://127.0.0.1/api/users/" + route.params.username
);
useHead({
  title: `${user!.value?.name} (${route.params.username}) / Twitter`,
});
interface Tweet {
  content: string;
  UpdatedAt: Date;
}
interface User {
  tweets: Tweet[];
  name: string;
}
</script>
<template>
  <div>
    <div
      class="flex gap-x-3 px-4 py-2 border-b-[1px] border-[#2f3336] text-[15px]"
      v-for="tweet in user!.tweets"
    >
      <div
        class="w-10 h-10 bg-[url(/default_profile_small.png)] bg-cover rounded-full"
      ></div>
      <div class="flex flex-col">
        <div class="flex gap-x-1">
          <h1 class="font-bold">{{ user!.name }}</h1>
          <h1 class="text-gray-500">@{{ route.params.username.toString() }}</h1>
          <h1 class="text-gray-500">
            ·
            {{
              new Intl.DateTimeFormat("en", {
                month: "short",
                day: "numeric",
              }).format(new Date(tweet.UpdatedAt))
            }}
          </h1>
        </div>
        <h1>{{ tweet.content }}</h1>
      </div>
    </div>
  </div>
</template>
