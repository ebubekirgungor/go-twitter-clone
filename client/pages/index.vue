<script setup lang="ts">
const route = useRoute();
const login_dialog = ref(route.query.login === null);
</script>
<template>
  <main class="h-screen bg-white dark:bg-black">
    <transition name="modal-fade" mode="out-in">
      <div
        v-if="login_dialog"
        class="bg-[#5b708366] inset-x-0 inset-y-0 w-full h-full fixed font-bold text-black dark:text-stone-200"
      >
        <div
          class="fixed z-2 flex flex-col m-auto inset-x-0 inset-y-0 p-2 min-w-[600px] w-[600px] min-h-[400px] h-[500px] bg-white dark:bg-black rounded-2xl"
        >
          <NuxtLink
            to="/"
            @click="login_dialog = false"
            class="transition duration-300 ease-in-out flex justify-center items-center w-10 h-10 rounded-full hover:bg-white/10"
            >x</NuxtLink
          >
          <h1 class="text-3xl text-center grow select-none pb-6">
            Log in to Twitter
          </h1>
        </div>
      </div>
    </transition>
    <div
      class="h-screen flex justify-center items-center select-none font-bold text-black dark:text-stone-200"
    >
      <div class="w-[400px] h-[328px] bg-[url(twitter.svg)]"></div>
      <div class="ml-48">
        <div class="h-32">
          <h1 class="text-6xl">Happening now</h1>
        </div>

        <div class="h-16">
          <h1 class="text-3xl">Join now.</h1>
        </div>
        <div class="h-24 flex flex-col justify-between">
          <button
            class="transition duration-300 ease-in-out w-[300px] h-10 rounded-full bg-cyan-600 hover:bg-cyan-500"
          >
            Create account
          </button>
          <NuxtLink
            to="?login"
            @click="login_dialog = true"
            class="transition duration-300 ease-in-out w-[300px] h-10 flex justify-center items-center rounded-full border border-slate-600 text-cyan-600 hover:bg-cyan-600/10"
          >
            Sign in
          </NuxtLink>
        </div>
      </div>
    </div>
  </main>
</template>
