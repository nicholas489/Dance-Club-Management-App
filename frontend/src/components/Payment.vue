<template>
    <div class="payment-container">
      <h2 class="payment-title">Pay Invoice</h2>
  
      <div class="payment-details">
        <div class="payment-method">
          <!-- Card icons here -->
        </div>
  
        <div class="payment-amount">
          <label for="amount" class="block text-sm font-medium text-gray-700">Payment amount</label>
          <div class="mt-1 relative rounded-md shadow-sm">
            <input type="text" id="amount" class="form-input block w-full sm:text-sm border-gray-300" placeholder="Canadian Dollars" readonly />
          </div>
        </div>
  
        <div class="input-group">
          <label for="cardName">Name on card</label>
          <input type="text" id="cardName" class="form-input" placeholder="Name on card">
        </div>
  
        <div class="input-group">
          <label for="cardNumber">Card number</label>
          <input type="text" id="cardNumber" class="form-input" placeholder="Card number" maxlength="16">
        </div>
  
        <div class="flex gap-4 mb-4">
          <div class="input-group flex-1">
            <label for="expiry">Expiry date</label>
            <input type="text" id="expiry" class="form-input" placeholder="MM/YY">
          </div>
          <div class="input-group flex-1">
            <label for="securityCode">Security code</label>
            <input type="text" id="securityCode" class="form-input" placeholder="CVV" maxlength="3">
          </div>
        </div>
  
        <div class="input-group">
          <label for="zip">ZIP/Postal code</label>
          <input type="text" id="zip" class="form-input" placeholder="ZIP Code">
        </div>
  
        <!-- PrimeVue Button -->
        <Button label="Pay" @click="submitPayment" class="w-full mt-2" />
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import Button from 'primevue/button';
  import { useRouter } from 'vue-router';
  
  const router = useRouter();
  
  function submitPayment() {
    // Payment submission logic here
    console.log('Payment submitted');
  }
  </script>
  
  <style scoped>
  .payment-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    max-width: 400px;
    margin: 2rem auto;
    padding: 2rem;
    border: 1px solid #ccc;
    border-radius: 0.5rem;
    background: #fff;
    box-shadow: 0 4px 8px rgba(0,0,0,0.1);
  }
  
  .payment-title {
    margin-bottom: 1rem;
  }
  
  .payment-method {
    margin-bottom: 1rem;
  }
  
  .payment-amount,
  .input-group {
    width: 100%;
    margin-bottom: 1rem;
  }
  
  .form-input {
    padding: 0.5rem;
    border: 1px solid #ccc;
    border-radius: 0.375rem;
    width: 100%;
  }
  
  .flex {
    display: flex;
    gap: 1rem;
  }
  
  .flex-1 {
    flex: 1;
  }
  
  .input-group label {
    display: block;
    margin-bottom: 0.5rem;
  }
  
  /* Import PrimeVue button styles if not globally imported */
  </style>
  