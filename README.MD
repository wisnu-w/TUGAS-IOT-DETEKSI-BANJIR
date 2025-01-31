# Tugas IoT Prediksi Banjir

Kelompok:
- **Annisa Nuraini** - 2170231111
- **Nabila Aulia Putri** - 2170231109
- **Wisnu Wardanu** - 2170231125

Informatika Unkris

**Link Wokwi**: [https://wokwi.com/projects/320964045035274834](https://wokwi.com/projects/320964045035274834)

## Deskripsi Proyek

Proyek ini bertujuan untuk memprediksi banjir menggunakan data yang dikirimkan oleh sensor melalui protokol MQTT. Data sensor kemudian diproses dan disimpan ke dalam database **InfluxDB** menggunakan bahasa pemrograman **Go (Golang)**. Aplikasi ini mendengarkan data dari broker MQTT dan menyimpan data tersebut sebagai time-series di InfluxDB.

### Arsitektur Sistem
- **MQTT**: Protokol komunikasi yang digunakan oleh perangkat IoT untuk mengirimkan data.
- **InfluxDB**: Database time-series yang digunakan untuk menyimpan data sensor secara efisien.
- **Golang**: Bahasa pemrograman yang digunakan untuk memproses data dan menyimpannya ke InfluxDB.

## Teknologi yang Digunakan

- **Golang**: Bahasa pemrograman untuk mengimplementasikan aplikasi yang menerima data dari MQTT dan menyimpannya ke InfluxDB.
- **MQTT Broker**: Digunakan untuk menerima dan mengirimkan pesan dari sensor.
- **InfluxDB**: Sistem database untuk menyimpan data time-series.
- **Wokwi**: Platform simulasi untuk perangkat IoT dan pengiriman data MQTT.

## Langkah-langkah Implementasi

1. **Koneksi ke MQTT Broker**: Aplikasi Go akan terhubung ke broker MQTT untuk menerima data sensor.
2. **Parsing Data**: Data yang diterima dalam format JSON, misalnya `{"water_level": 57}`, akan diparse menggunakan `json.Unmarshal`.
3. **Koneksi ke InfluxDB**: Data yang telah diparse akan disimpan ke dalam InfluxDB menggunakan client library `influxdb-client-go`.
4. **Menyimpan Data ke InfluxDB**: Data disimpan dalam format time-series di InfluxDB untuk analisis lebih lanjut.

## Struktur Kode

- **main.go**: Program utama yang mengatur koneksi ke MQTT, menerima data sensor, memparsing data, dan menyimpannya ke InfluxDB.
- **go.mod**: File modul Go yang berisi dependensi yang digunakan dalam proyek.

## Persyaratan

- **Go (Golang)**: Pastikan Go sudah terinstal di sistem Anda. Unduh di [https://golang.org/dl/](https://golang.org/dl/).
- **InfluxDB**: Anda perlu memiliki instance InfluxDB yang berjalan, baik lokal atau menggunakan layanan cloud.
- **MQTT Broker**: Gunakan broker MQTT yang sudah ada atau buat sendiri.

## Cara Menjalankan

1. **Jalankan Simulasi Wokwi**:
   - Kunjungi [link Wokwi](https://wokwi.com/projects/320964045035274834) untuk menjalankan simulasi perangkat IoT.
   - Perangkat ini akan mengirimkan data sensor ke broker MQTT. Pastikan broker MQTT sudah disiapkan dan dapat menerima data.
   
2. **Clone Repositori**:
   Clone repositori ini ke mesin lokal Anda:
   ```bash
   git clone https://github.com/username/repository.git
