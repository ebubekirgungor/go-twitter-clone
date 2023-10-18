import { defineStore } from "pinia";
export const useToken = defineStore("main", {
  state: () => ({
    token: null,
  }),
  persist: true,
});
