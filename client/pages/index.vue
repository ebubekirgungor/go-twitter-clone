<script setup lang="ts">
const route = useRoute();
const login_dialog = ref(route.query.login === null);
const signup_dialog = ref(route.query.signup === null);
const login_identity = ref("");
const user = ref({
  name: "",
  email: "",
});
const transition = ref(true);
const close_login_dialog = () => {
  transition.value = false;
  login_dialog.value = false;
  signup_dialog.value = true;
};
</script>
<template>
  <main class="h-screen bg-white dark:bg-black">
    <transition :css="transition" name="modal-fade" mode="out-in">
      <Dialog v-if="login_dialog" @close="login_dialog = false">
        <div class="h-32">
          <h1 class="text-3xl mt-5">Sign in to Twitter</h1>
        </div>
        <div class="h-16">
          <input
            v-model="login_identity"
            placeholder="E-mail or username"
            type="text"
            class="transition duration-300 ease-in-out w-[300px] h-10 rounded-md font-normal bg-black"
          />
        </div>
        <div class="h-16">
          <button
            :disabled="login_identity == ''"
            class="transition duration-300 ease-in-out w-[300px] h-10 rounded-full text-black bg-stone-200 hover:bg-white/80 disabled:bg-white/40 disabled:pointer-events-none"
          >
            Next
          </button>
        </div>
        <div class="h-32">
          <button
            class="transition duration-300 ease-in-out w-[300px] h-10 rounded-full border border-slate-600 hover:bg-white/10"
          >
            Forgot password?
          </button>
        </div>
        <h1 class="font-normal text-gray-500">
          Don't have an account?
          <NuxtLink
            @click="close_login_dialog"
            to="?signup"
            class="text-cyan-500"
            >Sign up</NuxtLink
          >
        </h1>
      </Dialog>
    </transition>
    <transition :css="transition" name="modal-fade" mode="out-in">
      <Dialog v-if="signup_dialog" @close="signup_dialog = false; transition = true;">
        <div class="h-32">
          <h1 class="text-3xl mt-5">Create your account</h1>
        </div>
        <div class="h-16">
          <input
            v-model="user.name"
            placeholder="Name"
            type="text"
            class="transition duration-300 ease-in-out w-[438px] h-12 rounded-md font-normal bg-black"
          />
        </div>
        <div class="grow">
          <input
            v-model="user.email"
            placeholder="E-mail"
            type="text"
            class="transition duration-300 ease-in-out w-[438px] h-12 rounded-md font-normal bg-black"
          />
        </div>
        <div class="h-16">
          <button
            :disabled="user.name == '' || user.email == ''"
            class="transition duration-300 ease-in-out w-[438px] h-12 rounded-full text-black bg-stone-200 hover:bg-white/80 disabled:bg-white/40 disabled:pointer-events-none"
          >
            Next
          </button>
        </div>
      </Dialog>
    </transition>
    <div
      class="h-screen flex justify-center items-center select-none font-bold text-black dark:text-stone-200"
    >
      <div class="w-[400px] h-[328px] bg-[url(twitter.svg)]"></div>
      <div class="ml-80">
        <div class="h-32">
          <h1 class="text-6xl">Happening now</h1>
        </div>

        <div class="h-16">
          <h1 class="text-3xl">Join today.</h1>
        </div>
        <div class="h-24 flex flex-col justify-between">
          <NuxtLink
            to="?signup"
            @click="signup_dialog = true"
            class="transition duration-300 ease-in-out w-[300px] h-10 flex justify-center items-center rounded-full bg-cyan-600 hover:bg-cyan-500"
          >
            Create account
          </NuxtLink>
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
