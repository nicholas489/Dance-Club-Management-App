<!-- Dashboard.vue -->
<script setup lang="ts">
import Schedule from './dash_comp/Schedule.vue'
import Booking from './dash_comp/Bookings.vue'
import Finance from './dash_comp/Finances.vue'
import BoxFour from './dash_comp/BoxFour.vue'//rename it later or remove completely just to fill up space 
</script>

<template>
  <div class="dashboard">
    <h2>Welcome this is your Dashboard</h2>
    <div class="grid grid-cols-2 gap-4">
      <Schedule />
      <Booking />
      <Finance />
      <BoxFour /> <!-- Remove or rename later just for testing reasons -->
    </div>
  </div>
</template>

<style scoped>
.dashboard {
  padding: 20px;
}
.grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 20px;
}

.h2 {
    padding-bottom: 10px;
    margin-bottom: 10px;
}
/* Additional styling as needed */
</style>
