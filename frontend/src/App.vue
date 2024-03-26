<script setup lang="ts">
import {RouterView} from 'vue-router'
import Navbar from "@/components/organisms/navbar.vue";
</script>

<template>
    <div class="flex flex-col h-full">
        <Navbar />
        <RouterView/>
    </div>
</template>

<style scoped>
</style>
