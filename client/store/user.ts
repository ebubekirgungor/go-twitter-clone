import { defineStore } from "pinia";
export const useUser = defineStore("user", {
  state: () => ({
    user: {
      id: null,
      username: null,
      email: null,
      name: null,
      joined: new Date(),
      token: null,
    },
  }),
  persist: true,
});
