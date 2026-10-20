<template>
  <div style="padding: 20px; font-family: sans-serif; max-width: 800px; margin: 0 auto;">
    <h2>📦 ระบบจัดการพัสดุหอพัก (Week 5 Mockup)</h2>

    <!-- 1. ช่องค้นหาพัสดุ -->
    <div style="margin-bottom: 20px; background-color: #eef2f5; padding: 15px; border-radius: 8px;">
      <label style="font-weight: bold; margin-right: 10px;">🔍 ค้นหาพัสดุ:</label>
      <input 
        v-model="searchQuery"
        placeholder="พิมพ์เลข Tracking หรือชื่อผู้รับ..."
        style="padding: 8px; width: 250px; border: 1px solid #ccc; border-radius: 4px;"
      />
    </div>

    <!-- 2. ฟอร์มเพิ่มพัสดุใหม่ -->
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
        เพิ่มพัสดุ
      </button>
    </div>

    <!-- 3. ตารางแสดงรายการพัสดุ -->
    <table border="1" style="width: 100%; text-align: left; border-collapse: collapse; border: 1px solid #ddd;">
      <thead>
        <tr style="background-color: #f2f2f2;">
          <th style="padding: 10px;">เลข Tracking</th>
          <th style="padding: 10px;">ชื่อผู้รับ</th>
          <th style="padding: 10px;">สถานะ</th>
          <th style="padding: 10px;">จัดการ</th>
        </tr>
      </thead>
      <tbody>
        <!-- วนลูปเฉพาะรายการที่ตรงกับคำค้นหา -->
        <tr v-for="(parcel, index) in filteredParcels" :key="index">
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
          </td>
        </tr>
        <!-- กรณีค้นหาแล้วไม่เจอข้อมูล -->
        <tr v-if="filteredParcels.length === 0">
          <td colspan="4" style="padding: 15px; text-align: center; color: #888;">
            ไม่พบข้อมูลพัสดุที่ค้นหา
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

// ช่องรับค่าจาก Input
const newTracking = ref('')
const newName = ref('')
const searchQuery = ref('') // คำค้นหา

// ข้อมูลพัสดุจำลอง
const parcels = ref([
  { tracking: 'TH12345', name: 'ฟาติมะห์', status: 'รอรับ' },
  { tracking: 'SP98765', name: 'พี่วัน', status: 'รับแล้ว' },
  { tracking: 'TH99999', name: 'มะห์ฮะจัง', status: 'รอรับ' }
])

// ฟังก์ชันเพิ่มพัสดุ
const addParcel = () => {
  if (newTracking.value !== '' && newName.value !== '') {
    parcels.value.push({
      tracking: newTracking.value,
      name: newName.value,
      status: 'รอรับ'
    })
    newTracking.value = ''
    newName.value = ''
  } else {
    alert('กรุณากรอกข้อมูลให้ครบถ้วน!')
  }
}

// ฟังก์ชันเปลี่ยนสถานะเป็นรับแล้ว
const updateStatus = (targetParcel) => {
  targetParcel.status = 'รับแล้ว'
}

// ระบบค้นหาอัตโนมัติ (พิมพ์ปุ๊บ กรองตารางปั๊บ)
const filteredParcels = computed(() => {
  return parcels.value.filter(parcel => {
    return parcel.tracking.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
           parcel.name.toLowerCase().includes(searchQuery.value.toLowerCase())
  })
})
</script>