<template>
  <div class="box">
    <p>This is Finance</p>
    <!-- Using PrimeVue Button component -->
    <Button label="Go to Payment" @click="goToPayment" />
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import Button from 'primevue/button';

const router = useRouter();

function goToPayment() {
  // The name 'payment' should match the name given to the route in your router configuration
  router.push({ name: 'payment' });
}
</script>

<style scoped>
.box {
  display: flex;
  flex-direction: column; /* Align children vertically */
  align-items: center;
  justify-content: center;
  border: 2px solid #ccc;
  border-radius: 8px;
  height: 300px;
  padding: 20px;
  gap: 20px; /* Spacing between elements */
}
</style>
