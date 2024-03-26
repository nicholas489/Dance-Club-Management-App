<script setup lang="ts">
import InputText from 'primevue/inputtext';
import FloatLabel from 'primevue/floatlabel';
import {ref} from 'vue';

const username = ref('');
const password = ref('');
</script>

<template>
    <form>
        <FloatLabel>
            <InputText id="username" v-model="username"/>
            <label for="username">Username</label>
        </FloatLabel>
        <FloatLabel>
            <InputText id="password" v-model="password"/>
            <label for="password">Password</label>
        </FloatLabel>
    </form>
</template>

<style scoped>

</style>