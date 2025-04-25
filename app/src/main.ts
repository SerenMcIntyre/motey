import './assets/main.css';
import vue3GoogleLogin from "vue3-google-login";

import { createApp } from "vue";
import App from "./App.vue";
import ui from "@nuxt/ui/vue-plugin";
import router from "./router";

const app = createApp(App);

app.use(vue3GoogleLogin, {
  clientId: "clientid",
});

app.use(ui);
app.use(router);

app.mount("#app");
