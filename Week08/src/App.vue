<template>
  <div style="padding: 20px; font-family: sans-serif; max-width: 800px; margin: 0 auto;">
    <h2>📦 ระบบจัดการพัสดุหอพัก (Full-stack Week 8)</h2>

    <!-- ช่องค้นหาพัสดุ -->
    <div style="margin-bottom: 20px; background-color: #eef2f5; padding: 15px; border-radius: 8px;">
      <label style="font-weight: bold; margin-right: 10px;">🔍 ค้นหาพัสดุ:</label>
      <input
        v-model="searchQuery"
        placeholder="พิมพ์เลข Tracking หรือชื่อผู้รับ..."
        style="padding: 8px; width: 250px; border: 1px solid #ccc; border-radius: 4px;"
      />
    </div>

    <!-- ฟอร์มเพิ่มพัสดุใหม่ -->
    <div style="margin-bottom: 20px; background-color: #f9f9f9; padding: 15px; border-radius: 8px; border: 1px solid #ddd;">
      <h3 style="margin-top: 0; font-size: 16px;">➕ เพิ่มข้อมูลพัสดุใหม่</h3>
      <input
        v-model="newTracking"
        placeholder="เลข Tracking..."
        style="padding: 8px; margin-right: 10px; border: 1px solid #ccc; border-radius: 4px;"
      />
      <input
        v-model="newName"
        placeholder="ชื่อผู้รับ..."
        style="padding: 8px; margin-right: 10px; border: 1px solid #ccc; border-radius: 4px;"
      />
      <button
        @click="addParcel"
        style="padding: 8px 15px; background-color: #4CAF50; color: white; border: none; border-radius: 4px; cursor: pointer;"
      >
        เพิ่มพัสดุลงฐานข้อมูล
      </button>
    </div>

    <!-- ตารางแสดงรายการพัสดุ -->
    <table border="1" style="width: 100%; text-align: left; border-collapse: collapse; border: 1px solid #ddd;">
      <thead>
        <tr style="background-color: #f2f2f2;">
          <th style="padding: 10px;">ID</th>
          <th style="padding: 10px;">เลข Tracking</th>
          <th style="padding: 10px;">ชื่อผู้รับ</th>
          <th style="padding: 10px;">สถานะ</th>
          <th style="padding: 10px;">จัดการ</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="parcel in filteredParcels" :key="parcel.id">
          <td style="padding: 10px;">{{ parcel.id }}</td>
          <td style="padding: 10px;">{{ parcel.tracking }}</td>
          <td style="padding: 10px;">{{ parcel.name }}</td>
          <td
            style="padding: 10px; font-weight: bold;"
            :style="{ color: parcel.status === 'รับแล้ว' ? 'green' : 'orange' }"
          >
            {{ parcel.status }}
          </td>
          <td style="padding: 10px;">
            <button
              v-if="parcel.status === 'รอรับ'"
              @click="updateStatus(parcel)"
              style="padding: 5px 10px; background-color: #2196F3; color: white; border: none; border-radius: 4px; cursor: pointer;"
            >
              รับของแล้ว
            </button>

            <button
              @click="deleteParcel(parcel.id)"
              style="padding: 5px 10px; background-color: #f44336; color: white; border: none; border-radius: 4px; cursor: pointer;"
            >
              ลบ
            </button>
            
          </td>
        </tr>
        <tr v-if="filteredParcels.length === 0">
          <td colspan="5" style="padding: 15px; text-align: center; color: #888;">
            ไม่พบข้อมูลพัสดุ
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue' // เพิ่ม onMounted สำหรับโหลดข้อมูลตอนเปิดเว็บ

const newTracking = ref('')
const newName = ref('')
const searchQuery = ref('')
const parcels = ref([]) // เปลี่ยนเป็น Array ว่างๆ รอรับข้อมูลจาก API

// 1. ฟังก์ชันดึงข้อมูลจาก Go (GET)
const fetchParcels = async () => {
  try {
    const response = await fetch('http://localhost:3000/api/parcels')
    const data = await response.json()
    parcels.value = data || []
  } catch (error) {
    console.error('โหลดข้อมูลไม่สำเร็จ ให้เช็คว่ารัน Go อยู่หรือไม่:', error)
  }
}

// โหลดข้อมูลทันทีเมื่อเปิดหน้าเว็บ
onMounted(() => {
  fetchParcels()
})

// 2. ฟังก์ชันเพิ่มพัสดุ (POST)
const addParcel = async () => {
  if (newTracking.value !== '' && newName.value !== '') {
    try {
      const response = await fetch('http://localhost:3000/api/parcels', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ tracking: newTracking.value, name: newName.value })
      })

      if (response.ok) {
        newTracking.value = ''
        newName.value = ''
        fetchParcels() // โหลดตารางใหม่หลังเพิ่มเสร็จ
      }
    } catch (error) {
      alert('เพิ่มพัสดุไม่สำเร็จ')
    }
  } else {
    alert('กรุณากรอกข้อมูลให้ครบถ้วน!')
  }
}

// 3. ฟังก์ชันเปลี่ยนสถานะ (PUT)
const updateStatus = async (targetParcel) => {
  try {
    const response = await fetch(`http://localhost:3000/api/parcels/${targetParcel.id}`, {
      method: 'PUT'
    })
    
    if (response.ok) {
      fetchParcels() // โหลดตารางใหม่ให้สถานะเปลี่ยนเป็นสีเขียว
    }
  } catch (error) {
    alert('อัปเดตสถานะไม่สำเร็จ')
  }
}

// 4. ฟังก์ชันลบพัสดุ (DELETE) พร้อมระบบยืนยัน
const deleteParcel = async (id) => {
  // แจ้งเตือนให้กดยืนยันก่อนลบ ป้องกันการกดผิด
  if (confirm('คุณแน่ใจหรือไม่ว่าต้องการลบรายการพัสดุนี้?')) {
    try {
      const response = await fetch(`http://localhost:3000/api/parcels/${id}`, {
        method: 'DELETE'
      })
      
      if (response.ok) {
        fetchParcels() // โหลดตารางใหม่ให้รายการหายไป
      }
    } catch (error) {
      alert('ลบพัสดุไม่สำเร็จ')
    }
  }
}

// ระบบค้นหาอัตโนมัติ
const filteredParcels = computed(() => {
  return parcels.value.filter(parcel => {
    return parcel.tracking.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
           parcel.name.toLowerCase().includes(searchQuery.value.toLowerCase())
  })
})
</script>