export default defineNuxtRouteMiddleware((to, from) => {
  if (!localStorage.getItem("token")) {
    return navigateTo("/");
  }
});
