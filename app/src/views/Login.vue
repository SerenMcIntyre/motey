<script setup lang="ts">
import { useAuth0 } from "@auth0/auth0-vue";
import { openUrl } from "@tauri-apps/plugin-opener";
import { isTauri } from "@tauri-apps/api/core";
import { onOpenUrl } from "@tauri-apps/plugin-deep-link";

const { loginWithRedirect } = useAuth0();

const { handleRedirectCallback } = useAuth0();

if (isTauri()) {
  onOpenUrl(async (urls) => {
    console.debug("[DEEPLINK]", urls);
    for (const url of urls) {
      if (
        url.includes("state") &&
        (url.includes("code") || url.includes("error"))
      ) {
        await handleRedirectCallback(url);
      }
    }
  });
}



const login = async () => {
    console.debug("[LOGIN]");
    console.debug(isTauri());
  if (!isTauri()) {
    await loginWithRedirect();
    return;
  }
  await loginWithRedirect({
    openUrl: (url: string) =>
      openUrl(url).then(() => {
        console.log("URL opened:", url);
      }),
  });
};
</script>

<template>
  <button @click="login" class="bg-blue-500 text-white px-4 py-2 rounded">
    login with auth0
  </button>
</template>
