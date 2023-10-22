<script setup lang="ts">
import { useUser } from "@/store/user";
const route = useRoute();
interface Tweet {
  content: string;
}
interface User {
  tweets: Tweet[];
  following: User[];
  followers: User[];
}
const { user } = useUser();
const { data: userData } = useFetch<User>(
  "http://127.0.0.1/api/users/" + user.id
);
</script>
<template>
  <NuxtLayout name="app">
    <main
      class="h-screen border-x border-[#2f3336] flex flex-col text-black dark:text-stone-200"
    >
      <div class="flex items-center ml-2 gap-x-6 w-56 h-[53px]">
        <button
          class="transition duration-300 ease-in-out flex justify-center items-center w-9 h-9 rounded-full hover:bg-black/10 dark:hover:bg-white/10"
        >
          <img
            src="/back.svg"
            class="pointer-events-none w-5 h-5 dark:invert"
          />
        </button>
        <div>
          <h1 class="font-bold text-xl">{{ user.name }}</h1>
          <h1 class="text-sm text-gray-500">
            {{ userData!.tweets.length }}
            {{ userData!.tweets.length == 1 ? "post" : "posts" }}
          </h1>
        </div>
      </div>
      <div>
        <div class="bg-[#333639] w-[600px] h-[200px]"></div>
        <div class="flex justify-between ml-4 mr-4">
          <NuxtLink
            to="#"
            class="-mt-[69px] transition duration-200 ease-in-out border-black border-4 rounded-full w-[138px] h-[138px] bg-[url(/default_profile.png)] bg-cover hover:brightness-[85%]"
          ></NuxtLink>
          <NuxtLink
            to="#"
            class="mt-3 transition duration-200 ease-in-out w-32 h-9 flex justify-center items-center rounded-full border border-[#536471] font-bold text-[15px] hover:bg-white/10"
          >
            Set up profile
          </NuxtLink>
        </div>
        <div class="mt-3 ml-4">
          <h1 class="font-extrabold text-xl">{{ user.name }}</h1>
          <h1 class="text-[15px] text-gray-500">@{{ user.username }}</h1>
          <div class="mt-3 flex gap-x-1">
            <div
              class="mt-[1px] w-[19px] h-[19px] bg-[url(/calendar.svg)] bg-cover"
            ></div>
            <h1 class="text-[15px] text-gray-500">
              Joined
              {{
                new Intl.DateTimeFormat("en", {
                  month: "long",
                  year: "numeric",
                }).format(new Date(user.joined))
              }}
            </h1>
          </div>
          <div class="mt-2 flex gap-x-5 text-sm decoration-slate-200">
            <NuxtLink to="#" class="text-gray-500 hover:underline"
              ><span class="text-slate-200">
                {{ userData!.following.length }}
              </span>
              Following</NuxtLink
            >
            <NuxtLink to="#" class="text-gray-500 hover:underline"
              ><span class="text-slate-200">
                {{ userData!.followers.length }}
              </span>
              Followers</NuxtLink
            >
          </div>
        </div>
        <nav class="mt-4 flex w-full h-[53px] text-gray-500 border-b-[1px] border-[#2f3336]">
          <NuxtLink
            :class="
              route.name?.toString() == 'username'
                ? 'font-bold text-slate-200'
                : ''
            "
            class="transition duration-200 ease-in-out flex justify-center items-center w-full hover:bg-white/10"
            :to="'/' + user.username"
          >
            <div class="w-14 h-full flex flex-col justify-between items-center">
              <div></div>
              <h1>Posts</h1>
              <div
                v-if="route.name?.toString() == 'username'"
                class="w-full h-1 rounded-full bg-[#1d9bf0]"
              ></div>
              <div v-else></div>
            </div>
          </NuxtLink>
          <NuxtLink
            :class="
              route.name?.toString().split('-')[1] == 'replies'
                ? 'font-bold text-slate-200'
                : ''
            "
            class="transition duration-200 ease-in-out flex justify-center items-center w-full hover:bg-white/10"
            :to="'/' + user.username + '/replies'"
            ><div
              class="w-14 h-full flex flex-col justify-between items-center"
            >
              <div></div>
              <h1>Replies</h1>
              <div
                v-if="route.name?.toString().split('-')[1] == 'replies'"
                class="w-full h-1 rounded-full bg-[#1d9bf0]"
              ></div>
              <div v-else></div></div
          ></NuxtLink>
          <NuxtLink
            :class="
              route.name?.toString().split('-')[1] == 'media'
                ? 'font-bold text-slate-200'
                : ''
            "
            class="transition duration-200 ease-in-out flex justify-center items-center w-full hover:bg-white/10"
            :to="'/' + user.username + '/media'"
            ><div
              class="w-14 h-full flex flex-col justify-between items-center"
            >
              <div></div>
              <h1>Media</h1>
              <div
                v-if="route.name?.toString().split('-')[1] == 'media'"
                class="w-full h-1 rounded-full bg-[#1d9bf0]"
              ></div>
              <div v-else></div></div
          ></NuxtLink>
          <NuxtLink
            :class="
              route.name?.toString().split('-')[1] == 'likes'
                ? 'font-bold text-slate-200'
                : ''
            "
            class="transition duration-200 ease-in-out flex justify-center items-center w-full hover:bg-white/10"
            :to="'/' + user.username + '/likes'"
            ><div
              class="w-14 h-full flex flex-col justify-between items-center"
            >
              <div></div>
              <h1>Likes</h1>
              <div
                v-if="route.name?.toString().split('-')[1] == 'likes'"
                class="w-full h-1 rounded-full bg-[#1d9bf0]"
              ></div>
              <div v-else></div></div
          ></NuxtLink>
        </nav>
        <section>
          <slot />
        </section>
      </div>
    </main>
  </NuxtLayout>
</template>
