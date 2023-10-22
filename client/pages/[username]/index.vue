<script setup lang="ts">
import { useUser } from "@/store/user";
const { user } = useUser();
definePageMeta({
  layout: "profile",
  middleware: "auth",
});
useHead({
  title: `${user.name} (${user.username}) / Twitter`,
});
interface Tweet {
  content: string;
  UpdatedAt: Date;
}
interface User {
  tweets: Tweet[];
}
const { data: userData } = useFetch<User>(
  "http://127.0.0.1/api/users/" + user.id
);
</script>
<template>
  <div>
    <div
      class="flex gap-x-3 px-4 py-2 border-b-[1px] border-[#2f3336] text-[15px]"
      v-for="tweet in userData!.tweets"
    >
      <div
        class="w-10 h-10 bg-[url(/default_profile_small.png)] bg-cover rounded-full"
      ></div>
      <div class="flex flex-col">
        <div class="flex gap-x-1">
          <h1 class="font-bold">{{ user.name }}</h1>
          <h1 class="text-gray-500">@{{ user.username }}</h1>
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
