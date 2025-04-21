import { createMemoryHistory, createRouter } from "vue-router";

import HomeView from "./views/Home.vue";
import SettingsView from "./views/Settings.vue";
import EditTaskView from "./views/tasks/Edit.vue";


const routes = [
  { path: "/", component: HomeView },
  { path: "/settings", component: SettingsView },
  { path: "/tasks/new", component: EditTaskView },
];

const router = createRouter({
  history: createMemoryHistory(),
  routes,
});

export default router;

