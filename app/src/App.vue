<script setup lang="ts">
import { useAuth0 } from '@auth0/auth0-vue';

const {  user, isAuthenticated } = useAuth0();
</script>

<template>
  <UApp>
    <RouterLink to="/">Home</RouterLink>
    <RouterLink to="/login">Login</RouterLink>
    <RouterLink to="/settings">Settings</RouterLink>
    <RouterLink to="/tasks/new" class="px-2 underline text-red-500">Add task</RouterLink>
    <main>
      <RouterView />
      <pre class="text-primary-500" v-if="isAuthenticated">
        <code>{{ user }}</code>
      </pre>
    </main>
  </UApp>
</template>


