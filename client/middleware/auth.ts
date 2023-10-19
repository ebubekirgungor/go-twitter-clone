import { useToken } from "@/store/token";
export default defineNuxtRouteMiddleware((to, from) => {
  const { token } = useToken();
  if (!token) {
    return navigateTo("/?login");
  }
});
