export default defineNuxtConfig({
  devtools: { enabled: true },
  app: {
    head: {
      title: "Twitter. It's what's happening",
      link: [{ rel: "icon", type: "image/ico", href: "twitter.ico" }],
    },
  },
  css: ["~/assets/css/main.css"],
  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
});
