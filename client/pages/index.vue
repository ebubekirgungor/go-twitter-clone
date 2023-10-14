<script setup lang="ts"></script>
<template>
  <main class="h-screen bg-white dark:bg-black">
    <div class="h-screen flex justify-center items-center select-none font-bold text-black dark:text-stone-200">
      <div class="w-[400px] h-[328px] bg-[url(twitter.svg)]"></div>
      <div class="ml-48">
        <div class="h-32">
          <h1 class="text-6xl">Happening now</h1>
        </div>

        <div class="h-16">
          <h1 class="text-3xl">Join now.</h1>
        </div>
        <div class="h-24 flex flex-col justify-between">
          <button class="transition duration-300 ease-in-out w-[300px] h-10 rounded-full bg-cyan-600 hover:bg-cyan-500">
            Create account
          </button>
          <button class="transition duration-300 ease-in-out w-[300px] h-10 rounded-full border border-slate-600 text-cyan-600 hover:bg-cyan-600/10">
            Sign in
          </button>
        </div>
      </div>
    </div>
  </main>
</template>
