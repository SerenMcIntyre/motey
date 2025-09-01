<script setup lang="ts">
  import { useAuth0 } from '@auth0/auth0-vue';

  const {  user, isAuthenticated } = useAuth0();
</script>

<template>
  <UApp>
  <div class="min-h-screen">
    <main class="h-screen p-4 bg-primary-950/25">
        <pre class="text-primary-500" v-if="isAuthenticated">
          <code>{{ user }}</code>
        </pre>
        <RouterView  />
    </main>
    <nav class="fixed bottom-0 p-4 border-t-2 border-primary-900 w-screen">
      <RouterLink class="p-4" to="/">Home</RouterLink>
      <RouterLink class="p-4" to="/login">Login</RouterLink>
      <RouterLink class="p-4" to="/settings" active-class="text-primary">Settings</RouterLink>
      <RouterLink class="p-4" to="/tasks/new" active-class="text-secondary-500">Add task</RouterLink>
    </nav>
  </div>
  </UApp>
</template>


