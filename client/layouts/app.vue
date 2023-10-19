<script setup lang="ts">
const route = useRoute();
</script>
<template>
  <main
    class="min-h-[100dvh] bg-white dark:bg-black text-black dark:text-stone-200"
  >
    <div class="flex">
      <nav class="flex flex-col gap-y-2">
        <NuxtLink
          to="/"
          class="flex justify-center items-center h-12 w-12 transition duration-200 ease-in-out hover:bg-white/10 rounded-full"
        >
          <div
            class="bg-[url(twitter.svg)] bg-no-repeat w-[35px] h-[28px]"
          ></div>
        </NuxtLink>
        <NuxtLink
          to="/home"
          class="flex items-center h-12 w-[min-content] transition duration-200 ease-in-out hover:bg-white/10 rounded-full"
        >
          <div
            :class="
              route.name?.toString() == 'home'
                ? 'bg-[url(home_active.svg)]'
                : 'bg-[url(home.svg)]'
            "
            class="mx-3 w-6 h-6 dark:invert"
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
                ? 'bg-[url(explore_active.svg)] font-bold'
                : 'bg-[url(explore.svg)]'
            "
            class="mx-3 w-6 h-6 dark:invert"
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
                ? 'bg-[url(notifications_active.svg)] font-bold'
                : 'bg-[url(notifications.svg)]'
            "
            class="mx-3 w-6 h-6 dark:invert"
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
                ? 'bg-[url(messages_active.svg)] font-bold'
                : 'bg-[url(messages.svg)]'
            "
            class="mx-3 w-6 h-6 dark:invert"
          ></div>
          <h1
            :class="route.name?.toString() == 'messages' ? 'font-bold' : ''"
            class="mr-3 text-xl"
          >
            Messages
          </h1>
        </NuxtLink>
        <NuxtLink
          to="/profile"
          class="flex items-center h-12 w-[min-content] transition duration-200 ease-in-out hover:bg-white/10 rounded-full"
        >
          <div
            :class="
              route.name?.toString() == 'profile'
                ? 'bg-[url(profile_active.svg)] font-bold'
                : 'bg-[url(profile.svg)]'
            "
            class="mx-3 w-6 h-6 dark:invert"
          ></div>
          <h1
            :class="route.name?.toString() == 'profile' ? 'font-bold' : ''"
            class="mr-3 text-xl"
          >
            Profile
          </h1>
        </NuxtLink>
        <NuxtLink
          to="/"
          class="flex items-center h-12 w-[min-content] transition duration-200 ease-in-out hover:bg-white/10 rounded-full"
        >
          <div class="mx-3 w-6 h-6 dark:invert bg-[url(more.svg)]"></div>
          <h1 class="mr-3 text-xl">More</h1>
        </NuxtLink>
      </nav>
      <slot />
    </div>
  </main>
</template>
