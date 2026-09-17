<template>
  <div style="padding: 20px; font-family: sans-serif;">
    <h2>📝 My To-Do List (รวบยอด Week 3-4)</h2>
    
    <!-- Week 2: v-model รับค่าจากผู้ใช้ -->
    <input v-model="newTask" placeholder="พิมพ์สิ่งที่ต้องทำ..." />
    
    <!-- Week 4: @click ดักจับเหตุการณ์เมื่อกดปุ่ม -->
    <button @click="addTask" style="margin-left: 5px;">เพิ่มรายการ</button>

    <ul style="margin-top: 20px;">
      <!-- Week 3: v-for วนลูปแสดงรายการข้อมูล -->
      <li v-for="(task, index) in tasks" :key="index" style="margin-bottom: 10px;">
        {{ task }}
        <!-- Week 4: ส่งค่า index ไปในฟังก์ชันเพื่อลบให้ถูกตัว -->
        <button @click="removeTask(index)" style="margin-left: 10px; color: red;">ลบ</button>
      </li>
    </ul>
  </div>
</template>

<script setup>
import { ref } from 'vue'

// ตัวแปรเก็บข้อความใหม่ (Week 2)
const newTask = ref('')

// ตัวแปรเก็บ Array ของรายการที่ต้องทำ (Week 3)
const tasks = ref(['อ่านหนังสือ Vue.js', 'เขียนโค้ด', 'พักผ่อน'])

// ฟังก์ชันเพิ่มรายการ (Week 4)
const addTask = () => {
  if (newTask.value !== '') {
    tasks.value.push(newTask.value) // เพิ่มข้อมูลลงใน Array
    newTask.value = '' // ล้างช่องกรอกข้อความ
  }
}

// ฟังก์ชันลบรายการ (Week 4)
const removeTask = (index) => {
  tasks.value.splice(index, 1) // ลบข้อมูลออกจาก Array ตามตำแหน่ง
}
</script>