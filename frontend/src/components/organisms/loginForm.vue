<script setup lang="ts">
import InputText from 'primevue/inputtext';
import FloatLabel from 'primevue/floatlabel';
import Button from 'primevue/button';
import Toast from 'primevue/toast';


import {ref} from 'vue';
import {useToast} from "primevue/usetoast";
import {useAuthStore} from "@/stores/authStore";
import router from "@/router";

const toast = useToast();

const username = ref('');
const password = ref('');
const submitForm = async () => {
    if (!username.value || !password.value) {
        toast.add({severity: 'error', summary: 'Login', detail: 'Please fill in all fields', life: 3000});
        return;
    }
    const authStore = useAuthStore();
    const loginResult = await authStore.login({username: username.value, password: password.value});
    if (loginResult[0]) {
        toast.add({severity: 'success', summary: 'Login', detail: 'Successfully Logged in', life: 3000});
        await router.push({name: 'dashboard'});
    } else {
        toast.add({severity: 'error', summary: 'Login', detail: loginResult[1], life: 3000});
    }
};
</script>

<template>
    <div class="flex flex-col">
        <h1 class="text-3xl text-center">Login</h1>
        <form class="flex flex-col space-y-8 mt-6 w-full" @submit.prevent="submitForm">
            <FloatLabel>
                <InputText id="username" v-model="username"/>
                <label for="username">Username</label>
            </FloatLabel>
            <FloatLabel>
                <InputText id="password" v-model="password" type="password"/>
                <label for="password">Password</label>
            </FloatLabel>
            <Button label="Submit" type="submit"/>
        </form>
    </div>
    <Toast/>


</template>

<style scoped>

</style>