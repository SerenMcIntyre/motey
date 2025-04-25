<script setup lang="ts">
    const callback = async (response: any) => {
        console.log(response);
        // Handle the response from Google login
        // You can use the response to authenticate the user in your application
        // For example, you might want to send the token to your backend for verification
    };
</script>

<template>
    <GoogleLogin :callback="callback" prompt />
</template>