<script setup lang="ts">
import { useUser } from "@/store/user";
const route = useRoute();
const { user } = useUser();
</script>
<template>
  <main
    class="min-h-[100dvh] bg-white dark:bg-black text-black dark:text-stone-200"
  >
    <div class="flex justify-center">
      <nav class="flex flex-col gap-y-2 mr-2 mb-3">
        <NuxtLink
          to="/home"
          class="flex justify-center items-center h-12 w-12 transition duration-200 ease-in-out hover:bg-white/10 rounded-full"
        >
          <div
            class="bg-[url(/twitter.svg)] bg-no-repeat w-[27px] h-[22px]"
          ></div>
        </NuxtLink>
        <NuxtLink
          to="/home"
          class="flex items-center h-12 w-[min-content] transition duration-200 ease-in-out hover:bg-white/10 rounded-full"
        >
          <div
            :class="
              route.name?.toString() == 'home'
                ? 'bg-[url(/home_active.svg)]'
                : 'bg-[url(/home.svg)]'
            "
            class="ml-3 mr-5 w-6 h-6 dark:invert"
          ></div>
          <h1
            :class="route.name?.toString() == 'home' ? 'font-bold' : ''"
            class="mr-3 text-xl"
          >
            Home
          </h1>
        </NuxtLink>
        <NuxtLink
          to="/"
          class="flex items-center h-12 w-[min-content] transition duration-200 ease-in-out hover:bg-white/10 rounded-full"
        >
          <div
            :class="
              route.name?.toString() == 'explore'
                ? 'bg-[url(/explore_active.svg)] font-bold'
                : 'bg-[url(/explore.svg)]'
            "
            class="ml-3 mr-5 w-6 h-6 dark:invert"
          ></div>
          <h1
            :class="route.name?.toString() == 'explore' ? 'font-bold' : ''"
            class="mr-3 text-xl"
          >
            Explore
          </h1>
        </NuxtLink>
        <NuxtLink
          to="/"
          class="flex items-center h-12 w-[min-content] transition duration-200 ease-in-out hover:bg-white/10 rounded-full"
        >
          <div
            :class="
              route.name?.toString() == 'notifications'
                ? 'bg-[url(/notifications_active.svg)] font-bold'
                : 'bg-[url(/notifications.svg)]'
            "
            class="ml-3 mr-5 w-6 h-6 dark:invert"
          ></div>
          <h1
            :class="
              route.name?.toString() == 'notifications' ? 'font-bold' : ''
            "
            class="mr-3 text-xl"
          >
            Notifications
          </h1>
        </NuxtLink>
        <NuxtLink
          to="/"
          class="flex items-center h-12 w-[min-content] transition duration-200 ease-in-out hover:bg-white/10 rounded-full"
        >
          <div
            :class="
              route.name?.toString() == 'messages'
                ? 'bg-[url(/messages_active.svg)] font-bold'
                : 'bg-[url(/messages.svg)]'
            "
            class="ml-3 mr-5 w-6 h-6 dark:invert"
          ></div>
          <h1
            :class="route.name?.toString() == 'messages' ? 'font-bold' : ''"
            class="mr-3 text-xl"
          >
            Messages
          </h1>
        </NuxtLink>
        <NuxtLink
          :to="(user.username as any)"
          class="flex items-center h-12 w-[min-content] transition duration-200 ease-in-out hover:bg-white/10 rounded-full"
        >
          <div
            :class="
              route.name?.toString().startsWith('username')
                ? 'bg-[url(/profile_active.svg)] font-bold'
                : 'bg-[url(/profile.svg)]'
            "
            class="ml-3 mr-5 w-6 h-6 dark:invert"
          ></div>
          <h1
            :class="
              route.name?.toString().startsWith('username') ? 'font-bold' : ''
            "
            class="mr-3 text-xl"
          >
            Profile
          </h1>
        </NuxtLink>
        <NuxtLink
          to="/"
          class="flex items-center h-12 w-[min-content] transition duration-200 ease-in-out hover:bg-white/10 rounded-full"
        >
          <div class="ml-3 mr-5 w-6 h-6 dark:invert bg-[url(/more.svg)]"></div>
          <h1 class="mr-3 text-xl">More</h1>
        </NuxtLink>
        <div class="grow">
          <NuxtLink
            to="#"
            class="mt-5 transition duration-300 ease-in-out w-56 h-[52px] flex justify-center items-center rounded-full bg-[#1e9bf0] hover:bg-[#1a8cd8] text-white text-[17px] font-bold"
          >
            Post
          </NuxtLink>
        </div>
        <NuxtLink
          to="#"
          class="transition duration-300 ease-in-out w-64 h-16 flex justify-center items-center rounded-full hover:bg-white/10 text-white text-[17px] font-bold"
        >
          User
        </NuxtLink>
      </nav>

      <slot />
    </div>
  </main>
</template>
