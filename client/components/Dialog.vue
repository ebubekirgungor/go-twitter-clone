<script>
export default {
  props: ["type", "signup_step", "login_step"],
};
</script>
<template>
  <div
    class="bg-black/40 dark:bg-[#5b708366] inset-x-0 inset-y-0 w-full h-full fixed font-bold text-black dark:text-stone-200 select-none"
  >
    <div
      class="fixed z-2 flex flex-col items-center m-auto inset-x-0 sm:inset-y-0 p-2 w-full h-[100%] sm:min-w-[600px] sm:w-[600px] sm:min-h-[400px] sm:h-[500px] bg-white dark:bg-black sm:rounded-2xl"
    >
      <div v-if="type == 'signup'" class="h-8 self-start flex">
        <NuxtLink
          v-if="signup_step == 1"
          to="/"
          @click="$emit('close')"
          class="transition duration-300 ease-in-out flex justify-center items-center w-10 h-10 rounded-full hover:bg-black/10 dark:hover:bg-white/10"
          ><img
            src="/close.svg"
            class="pointer-events-none w-6 h-6 dark:invert"
        /></NuxtLink>
        <button
          v-else
          @click="$emit('back')"
          class="transition duration-300 ease-in-out flex justify-center items-center w-10 h-10 rounded-full hover:bg-black/10 dark:hover:bg-white/10"
        >
          <img
            src="/back.svg"
            class="pointer-events-none w-6 h-6 dark:invert"
          />
        </button>
        <h1 class="text-xl ml-5 mt-1">Step {{ signup_step }} of 3</h1>
      </div>
      <div v-else class="h-8 self-start">
        <NuxtLink
          v-if="login_step == 1"
          to="/"
          @click="$emit('close')"
          class="transition duration-300 ease-in-out flex justify-center items-center w-10 h-10 rounded-full hover:bg-black/10 dark:hover:bg-white/10"
          ><img
            src="/close.svg"
            class="pointer-events-none w-6 h-6 dark:invert"
        /></NuxtLink>
        <button
          v-else
          @click="$emit('back')"
          class="transition duration-300 ease-in-out flex justify-center items-center w-10 h-10 rounded-full hover:bg-black/10 dark:hover:bg-white/10"
        >
          <img
            src="/back.svg"
            class="pointer-events-none w-6 h-6 dark:invert"
          />
        </button>
      </div>
      <slot />
    </div>
  </div>
</template>
