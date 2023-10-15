<script setup lang="ts">
const route = useRoute();
const login_dialog = ref(route.query.login === null);
const signup_dialog = ref(route.query.signup === null);
const login_step = ref(1);
const signup_step = ref(1);
const login_identity = ref("");
const login_password = ref("");
const check_error = ref("");
const login_error = ref("");
const user = ref({
  name: "",
  email: "",
  username: "",
  password: "",
  confirm_password: "",
});
const transition = ref(true);
const close_login_dialog = () => {
  login_dialog.value = false;
  login_identity.value = "";
  login_password.value = "";
};
const close_signup_dialog = () => {
  signup_dialog.value = false;
  transition.value = true;
  user.value = {
    name: "",
    email: "",
    username: "",
    password: "",
    confirm_password: "",
  };
};
const signup = async () => {
  const { data: response } = await useFetch("http://localhost/api/users", {
    method: "post",
    body: {
      name: user.value.name,
      email: user.value.email,
      username: user.value.username,
      password: user.value.password,
    },
  });
  console.log(response);
};
const check_user = async () => {
  const { data: response } = await useFetch(
    "http://localhost/api/auth/check_user",
    {
      method: "post",
      body: {
        identity: login_identity.value,
      },
    }
  );
  if (response.value == "Ok") {
    login_step.value++;
    check_error.value = "";
  } else check_error.value = response.value as string;
};
const login = async () => {
  const { data: response } = await useFetch("http://localhost/api/auth/login", {
    method: "post",
    body: {
      identity: login_identity.value,
      password: login_password.value,
    },
  });
  console.log(response)
  if ((response.value as any).status == "success") {
    navigateTo("/main");
  } else login_error.value = (response.value as any).message as string;
};
</script>
<template>
  <main class="h-screen bg-white dark:bg-black">
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
              @click="check_user"
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
            <h1 class="text-center font-normal text-red-600">
              {{ check_error }}
            </h1>
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
        </div>
        <div v-if="login_step == 2">
          <div class="h-40">
            <h1 class="text-3xl mt-5">Sign in to Twitter</h1>
          </div>
          <div class="h-20">
            <input
              v-model="login_password"
              placeholder="Password"
              type="password"
              class="transition duration-300 ease-in-out w-[300px] h-10 rounded-md font-normal dark:bg-black"
            />
          </div>
          <div class="h-32">
            <h1 class="text-center font-normal text-red-600">
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
        <div v-if="signup_step == 1">
          <div class="h-28">
            <h1 class="text-3xl mt-5">Create your account</h1>
          </div>
          <div class="h-16">
            <input
              v-model="user.name"
              placeholder="Name"
              type="text"
              class="transition duration-300 ease-in-out w-[438px] h-12 rounded-md font-normal dark:bg-black"
            />
          </div>
          <div class="h-16">
            <input
              v-model="user.email"
              placeholder="E-mail"
              type="text"
              class="transition duration-300 ease-in-out w-[438px] h-12 rounded-md font-normal dark:bg-black"
            />
          </div>
          <div class="h-32">
            <input
              v-model="user.username"
              placeholder="Username"
              type="text"
              class="transition duration-300 ease-in-out w-[438px] h-12 rounded-md font-normal dark:bg-black"
            />
          </div>
          <div class="h-16">
            <button
              :disabled="user.name == '' || user.email == ''"
              @click="signup_step++"
              class="transition duration-300 ease-in-out w-[438px] h-12 rounded-full text-stone-200 bg-black dark:text-black dark:bg-stone-200 hover:bg-black/80 dark:hover:bg-white/80 disabled:bg-black/40 dark:disabled:bg-white/40 disabled:pointer-events-none"
            >
              Next
            </button>
          </div>
        </div>
        <div v-if="signup_step == 2">
          <div class="h-32">
            <h1 class="text-3xl mt-5">Set password</h1>
          </div>
          <div class="h-16">
            <input
              v-model="user.password"
              placeholder="Password"
              type="password"
              class="transition duration-300 ease-in-out w-[438px] h-12 rounded-md font-normal dark:bg-black"
            />
          </div>
          <div class="h-44">
            <input
              v-model="user.confirm_password"
              placeholder="Confirm Password"
              type="password"
              class="transition duration-300 ease-in-out w-[438px] h-12 rounded-md font-normal dark:bg-black"
            />
          </div>
          <div class="h-16">
            <button
              :disabled="
                user.password == '' || user.confirm_password != user.password
              "
              @click="signup_step++"
              class="transition duration-300 ease-in-out w-[438px] h-12 rounded-full text-stone-200 bg-black dark:text-black dark:bg-stone-200 hover:bg-black/80 dark:hover:bg-white/80 disabled:bg-black/40 dark:disabled:bg-white/40 disabled:pointer-events-none"
            >
              Next
            </button>
          </div>
        </div>
        <div v-if="signup_step == 3">
          <div class="h-32">
            <h1 class="text-3xl mt-5">Sign up</h1>
          </div>
          <div class="h-16">
            <input
              :placeholder="user.name"
              type="text"
              disabled
              class="transition duration-300 ease-in-out w-[438px] h-12 rounded-md font-normal dark:bg-black"
            />
          </div>
          <div class="h-44">
            <input
              :placeholder="user.email"
              type="text"
              disabled
              class="transition duration-300 ease-in-out w-[438px] h-12 rounded-md font-normal dark:bg-black"
            />
          </div>
          <div class="h-16">
            <button
              @click="signup"
              class="transition duration-300 ease-in-out w-[438px] h-12 rounded-full text-stone-200 bg-black dark:text-black dark:bg-stone-200 hover:bg-black/80 dark:hover:bg-white/80 disabled:bg-black/40 dark:disabled:bg-white/40 disabled:pointer-events-none"
            >
              Sign up
            </button>
          </div>
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
  </main>
</template>
