<script setup lang="ts">
import { ref } from "vue";
import { invoke } from "@tauri-apps/api/core";

const greetMsg = ref("");
const name = ref("");

async function greet() {
  // Learn more about Tauri commands at https://tauri.app/develop/calling-rust/
  greetMsg.value = await invoke("greet", { name: name.value });
}
</script>

<template>
  <UApp>
    <RouterLink to="/">Home</RouterLink>
    <RouterLink to="/settings">Settings</RouterLink>
    <RouterLink to="/tasks/new" class="px-2 underline text-red-500">Add task</RouterLink>
    <main>
      <RouterView />
    </main>
  </UApp>
</template>


