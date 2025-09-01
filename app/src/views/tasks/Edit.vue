<script setup lang="ts">
  import { FormSubmitEvent } from "@nuxt/ui";
  import { useToast } from "@nuxt/ui/runtime/composables/useToast.js";
  import * as v from "valibot";
  import { reactive } from "vue";

  const MeasurementSchema = v.object({
    unit: v.string(),
    sticker_value: v.number(),
  });

  const AddTaskRequestSchema = v.object({
    title: v.pipe(v.string(), v.minLength(1, "Title is required")),
    background: v.pipe(v.string(), v.minLength(7), v.maxLength(7), v.startsWith("#")), // Expecting a hex color code like #ffffff
    sticker: v.file("A valid file must pe provided for the sticker"),
    isMeasured: v.boolean(),
    measurement: v.union([MeasurementSchema, v.null()]),
  });
  
/*   const NotificationSchema = v.object({
    frequency_type: v.picklist(['daily', 'weekly', 'monthly']),
    frequency_value: v.number(),
    time: v.string(), // Consider adding a regex check for HH:mm format
  }); */
  
  const state = reactive({
    title: "",
    background: "#ffffff",
    sticker: null as File | null,
    isMeasured: false,
    measurement: null as { unit: string; sticker_value: number } | null,
  })
  
  const toast = useToast();
  
  async function onSubmit(event: FormSubmitEvent<v.InferOutput<typeof AddTaskRequestSchema>>) {
    event.preventDefault();
    toast.add({
      title: "Task Saved",
      color: "secondary",
    });
    console.debug("Form submitted with data:", event.data);
  }

</script>

<template>
  <UForm :schema="AddTaskRequestSchema" :state="state" 
          class="space-y-4" @submit="onSubmit">

    <div class="grid grid-cols-2 space-x-2">
      <UFormField name="sticker" label="Sticker">
        <UFileUpload v-model="state.sticker" class="" />
      </UFormField>
      <UFormField name="title" label="Title">
        <UInput v-model="state.title" placeholder="Enter task title" />
      </UFormField>
    </div>

    <UFormField name="background" label="Background">
      <UColorPicker v-model="state.background" />
    </UFormField>



    <UFormField name="isMeasured" label="Is Measured">
      <USwitch v-model="state.isMeasured" />
    </UFormField>

    <UFormField v-if="state.isMeasured" name="measurement" label="Measurement">
      <UForm :schema="MeasurementSchema">
        <UFormField name="unit" label="Unit">
          <UInput placeholder="Enter measurement unit (e.g., kg, cm)" />
        </UFormField>
        <UFormField name="sticker_value" label="Sticker Value">
          <UNumberInput placeholder="Enter sticker value" />
        </UFormField>
      </UForm>
    </UFormField>

    <UButton type="submit" color="primary">Save Task</UButton>
  </UForm>
</template>
