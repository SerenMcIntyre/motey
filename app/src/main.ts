import "./assets/main.css";

import { createApp } from "vue";
import App from "./App.vue";
import ui from "@nuxt/ui/vue-plugin";
import router from "./router";
import config from "./auth-config.json";
import { createAuth0 } from "@auth0/auth0-vue";
import { isTauri } from "@tauri-apps/api/core";
import { platform } from "@tauri-apps/plugin-os";

// Build the URL that Auth0 should redirect back to
const platformName = !isTauri() ? "web" : platform();
let redirect_uri = "";
if (!isTauri()) {
  redirect_uri = `${window.location.origin}/callback`;
} else if (platformName in ["ios", "android"]) {
  redirect_uri = `${config.app_id}://dev-iiluofr2qsjpqnpa.us.auth0.com/capacitor/${config.app_id}/callback`;
} else {
  redirect_uri = `motey://callback`;
}

console.debug("redirect_uri", redirect_uri);

const app = createApp(App);

const auth0Config = {
  domain: config.domain,
  clientId: config.client_id,
  useRefreshTokens: true,
  useRefreshTokenFallback: false,
  authorizationParams: {
    redirect_uri,
  },
};

app.use(createAuth0(auth0Config));
app.use(ui);
app.use(router);

app.mount("#app");
