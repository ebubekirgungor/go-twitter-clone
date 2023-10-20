<script setup lang="ts">
const route = useRoute();
import { useUser } from "@/store/user";
const login_dialog = ref(route.query.login === null);
const signup_dialog = ref(route.query.signup === null);
const login_step = ref(1);
const signup_step = ref(1);
const login_identity = ref("");
const login_password = ref("");
const check_error = ref("");
const login_error = ref("");
const signup_error = ref("");
const user = ref({
  name: "",
  email: "",
  password: "",
  confirm_password: "",
});
const transition = ref(true);
const close_login_dialog = () => {
  transition.value = true;
  login_dialog.value = false;
  login_identity.value = "";
  login_password.value = "";
  check_error.value = "";
};
const open_signup_dialog = () => {
  transition.value = false;
  login_dialog.value = false;
  login_identity.value = "";
  check_error.value = "";
  signup_dialog.value = true;
};
const close_signup_dialog = () => {
  signup_dialog.value = false;
  transition.value = true;
  user.value = {
    name: "",
    email: "",
    password: "",
    confirm_password: "",
  };
  check_error.value = "";
};
const signup = async () => {
  const { data: response } = await useFetch("/api/users", {
    method: "post",
    body: {
      name: user.value.name,
      email: user.value.email,
      username: user.value.name.replace(" ", "_"),
      password: user.value.password,
    },
  });
  if ((response.value as any).email) {
    transition.value = false;
    signup_dialog.value = false;
    signup_step.value = 1;
    login_dialog.value = true;
    user.value = {
      name: "",
      email: "",
      password: "",
      confirm_password: "",
    };
    check_error.value = "";
  } else signup_error.value = (response.value as any).message as string;
};
const check_user = async (step: string) => {
  const { data: response } = await useFetch("/api/auth/check_user", {
    method: "post",
    body: {
      identity: step == "login" ? login_identity.value : user.value.email,
    },
  });

  if (
    (step == "login" && response.value == "Ok") ||
    (step != "login" && response.value != "Ok")
  ) {
    if (step == "login") login_step.value++;
    else signup_step.value++;
    check_error.value = "";
  } else
    check_error.value =
      step == "login"
        ? (response.value as string)
        : "This email is already in use";
};
const login = async () => {
  const { data: response } = await useFetch("/api/auth/login", {
    method: "post",
    body: {
      identity: login_identity.value,
      password: login_password.value,
    },
  });
  if ((response.value as any).status == "success") {
    const user = useUser();
    user.user = (response.value as any).data;
    navigateTo("/home");
  } else login_error.value = (response.value as any).message as string;
};
</script>
<template>
  <main class="min-h-[100dvh] bg-white dark:bg-black">
    <transition :css="transition" name="modal-fade" mode="out-in">
      <Dialog
        v-if="login_dialog"
        @close="close_login_dialog"
        @back="login_step--"
        type="login"
        :login_step="login_step"
      >
        <div v-if="login_step == 1">
          <div class="h-32">
            <h1 class="text-3xl mt-5">Sign in to Twitter</h1>
          </div>
          <div class="h-16">
            <input
              v-model="login_identity"
              placeholder="E-mail or username"
              type="text"
              class="transition duration-300 ease-in-out w-[300px] h-10 rounded-md font-normal dark:bg-black"
            />
          </div>
          <div class="h-16">
            <button
              :disabled="login_identity == ''"
              @click="check_user('login')"
              class="transition duration-300 ease-in-out w-[300px] h-10 rounded-full text-stone-200 bg-black dark:text-black dark:bg-stone-200 hover:bg-black/80 dark:hover:bg-white/80 disabled:bg-black/40 dark:disabled:bg-white/40 disabled:pointer-events-none"
            >
              Next
            </button>
          </div>
          <div class="h-20">
            <button
              class="transition duration-300 ease-in-out w-[300px] h-10 rounded-full border border-slate-600 hover:bg-white/10"
            >
              Forgot password?
            </button>
          </div>
          <div class="h-12">
            <h1 class="text-center font-normal text-[#ff0000]">
              {{ check_error }}
            </h1>
          </div>
          <h1 class="font-normal text-gray-500">
            Don't have an account?
            <NuxtLink
              @click="open_signup_dialog"
              to="?signup"
              class="text-cyan-500"
              >Sign up</NuxtLink
            >
          </h1>
        </div>
        <div class="h-full flex flex-col" v-if="login_step == 2">
          <div class="h-40">
            <h1 class="text-3xl mt-5">Enter your password</h1>
          </div>
          <div class="h-16">
            <input
              :placeholder="login_identity"
              type="text"
              disabled
              class="w-[300px] h-10 rounded-md font-normal dark:bg-black brightness-50 disabled:bg-neutral-800 border-none"
            />
          </div>
          <div class="h-16">
            <input
              v-model="login_password"
              placeholder="Password"
              type="password"
              class="transition duration-300 ease-in-out w-[300px] h-10 rounded-md font-normal dark:bg-black"
            />
          </div>
          <div class="grow">
            <h1 class="text-center font-normal text-[#ff0000]">
              {{ login_error }}
            </h1>
          </div>
          <div class="h-16">
            <button
              :disabled="login_password == ''"
              @click="login"
              class="transition duration-300 ease-in-out w-[300px] h-10 rounded-full text-stone-200 bg-black dark:text-black dark:bg-stone-200 hover:bg-black/80 dark:hover:bg-white/80 disabled:bg-black/40 dark:disabled:bg-white/40 disabled:pointer-events-none"
            >
              Log in
            </button>
          </div>
        </div>
      </Dialog>
    </transition>
    <transition :css="transition" name="modal-fade" mode="out-in">
      <Dialog
        v-if="signup_dialog"
        @close="close_signup_dialog"
        @back="signup_step--"
        type="signup"
        :signup_step="signup_step"
      >
        <div class="h-full flex flex-col" v-if="signup_step == 1">
          <div class="h-40 sm:h-32">
            <h1 class="text-3xl mt-5">Create your account</h1>
          </div>
          <div class="h-16">
            <input
              v-model="user.name"
              placeholder="Name"
              type="text"
              class="transition duration-300 ease-in-out w-[300px] sm:w-[438px] h-12 rounded-md font-normal dark:bg-black"
            />
          </div>
          <div class="h-56 sm:h-24">
            <input
              v-model="user.email"
              placeholder="E-mail"
              type="text"
              class="transition duration-300 ease-in-out w-[300px] sm:w-[438px] h-12 rounded-md font-normal dark:bg-black"
            />
          </div>
          <div class="grow">
            <h1 class="text-center font-normal text-[#ff0000]">
              {{ check_error }}
            </h1>
          </div>
          <div class="h-16">
            <button
              :disabled="user.name == '' || user.email == ''"
              @click="check_user('signup')"
              class="transition duration-300 ease-in-out w-[300px] sm:w-[438px] h-12 rounded-full text-stone-200 bg-black dark:text-black dark:bg-stone-200 hover:bg-black/80 dark:hover:bg-white/80 disabled:bg-black/40 dark:disabled:bg-white/40 disabled:pointer-events-none"
            >
              Next
            </button>
          </div>
        </div>
        <div class="h-full flex flex-col" v-if="signup_step == 2">
          <div class="h-48 sm:h-32">
            <h1 class="text-3xl mt-5">Set password</h1>
          </div>
          <div class="h-16">
            <input
              v-model="user.password"
              placeholder="Password"
              type="password"
              class="transition duration-300 ease-in-out w-[300px] sm:w-[438px] h-12 rounded-md font-normal dark:bg-black"
            />
          </div>
          <div class="grow">
            <input
              v-model="user.confirm_password"
              placeholder="Confirm Password"
              type="password"
              class="transition duration-300 ease-in-out w-[300px] sm:w-[438px] h-12 rounded-md font-normal dark:bg-black"
            />
          </div>
          <div class="h-16">
            <button
              :disabled="
                user.password == '' || user.confirm_password != user.password
              "
              @click="signup_step++"
              class="transition duration-300 ease-in-out w-[300px] sm:w-[438px] h-12 rounded-full text-stone-200 bg-black dark:text-black dark:bg-stone-200 hover:bg-black/80 dark:hover:bg-white/80 disabled:bg-black/40 dark:disabled:bg-white/40 disabled:pointer-events-none"
            >
              Next
            </button>
          </div>
        </div>
        <div class="h-full flex flex-col" v-if="signup_step == 3">
          <div class="h-48 sm:h-32">
            <h1 class="text-3xl mt-5">Sign up</h1>
          </div>
          <div class="h-16">
            <input
              :placeholder="user.name"
              type="text"
              disabled
              class="w-[300px] sm:w-[438px] h-12 rounded-md font-normal dark:bg-black brightness-50 disabled:bg-neutral-800 border-none"
            />
          </div>
          <div class="h-16">
            <input
              :placeholder="user.email"
              type="text"
              disabled
              class="w-[300px] sm:w-[438px] h-12 rounded-md font-normal dark:bg-black brightness-50 disabled:bg-neutral-800 border-none"
            />
          </div>
          <div class="grow">
            <h1 class="text-center font-normal text-[#ff0000]">
              {{ signup_error }}
            </h1>
          </div>
          <div class="h-16">
            <button
              @click="signup"
              class="transition duration-300 ease-in-out w-[300px] sm:w-[438px] h-12 rounded-full text-stone-200 bg-black dark:text-black dark:bg-stone-200 hover:bg-black/80 dark:hover:bg-white/80 disabled:bg-black/40 dark:disabled:bg-white/40 disabled:pointer-events-none"
            >
              Sign up
            </button>
          </div>
        </div>
      </Dialog>
    </transition>
    <div
      class="sm:h-screen flex flex-col sm:flex-row justify-center items-center select-none font-bold text-black dark:text-stone-200"
    >
      <div class="flex mx-12 flex-col sm:flex-row gap-x-80 gap-y-12">
        <div
          class="mt-10 sm:mt-0 w-[45px] h-[36px] sm:w-[400px] sm:h-[328px] bg-[url(twitter.svg)] bg-no-repeat"
        ></div>
        <div class="flex flex-col">
          <div class="h-32">
            <h1 class="text-5xl sm:text-6xl">Happening now</h1>
          </div>
          <div class="h-16">
            <h1 class="text-2xl sm:text-3xl">Join today.</h1>
          </div>
          <div class="h-24 flex flex-col justify-between">
            <NuxtLink
              to="?signup"
              @click="signup_dialog = true"
              class="transition duration-300 ease-in-out w-[300px] h-10 flex justify-center items-center rounded-full bg-cyan-600 hover:bg-cyan-500 text-stone-200"
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
    </div>
  </main>
</template>
~/store/user