<template>
  <div
    class="bg-[#5b708366] inset-x-0 inset-y-0 w-full h-full fixed font-bold text-black dark:text-stone-200 select-none"
  >
    <div
      class="fixed z-2 flex flex-col items-center m-auto inset-x-0 inset-y-0 p-2 min-w-[600px] w-[600px] min-h-[400px] h-[500px] bg-white dark:bg-black rounded-2xl"
    >
      <div class="h-8 self-start">
        <NuxtLink
          to="/"
          @click="$emit('close')"
          class="transition duration-300 ease-in-out flex justify-center items-center w-10 h-10 rounded-full hover:bg-white/10"
          ><img src="/close.svg" class="pointer-events-none w-6 h-6"
        /></NuxtLink>
      </div>
      <slot />
    </div>
  </div>
</template>
