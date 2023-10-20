<script setup lang="ts">
import { useUser } from "@/store/user";
definePageMeta({
  layout: "app",
  middleware: "auth",
});
useHead({
  title: "Profile / Twitter",
});
interface Tweet {
  content: string;
}
interface User {
  tweets: Tweet[];
}
const { user } = useUser();
const { data: userData } = useFetch<User>(
  "http://127.0.0.1/api/users/" + user.id
);
</script>
<template>
  <main
    class="h-screen border-x border-[#2f3336] flex flex-col text-black dark:text-stone-200"
  >
    <div class="flex items-center ml-2 gap-x-6 w-56 h-[53px]">
      <button
        class="transition duration-300 ease-in-out flex justify-center items-center w-9 h-9 rounded-full hover:bg-black/10 dark:hover:bg-white/10"
      >
        <img src="/back.svg" class="pointer-events-none w-5 h-5 dark:invert" />
      </button>
      <div>
        <h1 class="font-bold text-xl">{{ user.name }}</h1>
        <h1 class="text-sm text-gray-500">
          {{ userData!.tweets.length }} posts
        </h1>
      </div>
    </div>
    <div>
      <div class="bg-[#333639] w-[600px] h-[200px]"></div>
      <div class="flex justify-between ml-5 mr-4">
        <div
          class="-mt-16 border-black border-4 rounded-full w-[134px] h-[134px] bg-[url(default_profile.png)] bg-cover"
        ></div>
        <NuxtLink
          to="#"
          class="mt-3 transition duration-300 ease-in-out w-32 h-9 flex justify-center items-center rounded-full border border-[#536471] font-bold text-[15px] hover:bg-cyan-600/10"
        >
          Set up profile
        </NuxtLink>
      </div>
      <div class="mt-3 ml-4">
        <h1 class="font-extrabold text-xl">{{ user.name }}</h1>
        <h1 class="text-[15px] text-gray-500">@{{ user.username }}</h1>
      </div>
    </div>
  </main>
</template>
