<template>
  <div style="padding: 20px;">
    <h2>ทดลอง v-model </h2>
    
    <!-- ช่อง Input ที่ผูกกับตัวแปร message -->
    <input v-model="message" placeholder="ลองพิมพ์ข้อความตรงนี้..." />

    <!-- แสดงผลค่าของตัวแปร message ทันที -->
    <p>ข้อความของคุณคือ: <strong>{{ message }}</strong></p>
  </div>
</template>

<script setup>
import { ref } from 'vue'

// สร้างตัวแปรมารับค่าและตั้งค่าเริ่มต้นเป็นค่าว่าง
const message = ref('')
</script>