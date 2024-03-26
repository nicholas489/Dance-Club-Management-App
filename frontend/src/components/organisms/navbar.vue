<script setup lang="ts">
import Menubar from 'primevue/menubar';
import {ref} from "vue";
import {useRouter} from 'vue-router';
const router = useRouter();
const items = ref([
    {
        label: 'Home',
        icon: 'pi pi-home',
        command: () => {
            router.push('/');
        }
    },
    {
        label: 'Dashboard',
        icon: 'pi pi-user',
        command: () => {
            router.push('/login');
        }
    }
])

</script>

<template>
<Menubar :model="items" />
</template>

<style scoped>

</style>