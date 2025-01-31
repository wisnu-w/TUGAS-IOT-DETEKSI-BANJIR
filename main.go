package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

const (
	MQTT_BROKER   = "tcp://103.89.165.110:1883"
	MQTT_TOPIC    = "sensor/air_level"
	MQTT_USER     = "admingsi"
	MQTT_PASSWORD = "Broker9s!"
	INFLUX_URL    = "https://us-east-1-1.aws.cloud2.influxdata.com"
	INFLUX_TOKEN  = "emVB6BMlS8iDAPttd3RN1xq57D7LARAeGxUKmByThk4q5x7AWVoxjPjixgB2Yud2HEvR4-tHv65FemKKzPyugg=="
	INFLUX_ORG    = "dev"
	INFLUX_BUCKET = "monitoring-river"
)

func main() {
	// Inisialisasi InfluxDB Client
	client := influxdb2.NewClient(INFLUX_URL, INFLUX_TOKEN)
	defer client.Close()
	writeAPI := client.WriteAPIBlocking(INFLUX_ORG, INFLUX_BUCKET)

	// Cek koneksi ke InfluxDB
	pong, err := client.Ping(context.Background())
	if err != nil {
		log.Fatal("Failed to connect to InfluxDB:", err)
	} else {
		log.Println("Connected to InfluxDB:", pong)
	}

	// MQTT Options
	opts := mqtt.NewClientOptions()
	opts.AddBroker(MQTT_BROKER)
	opts.SetUsername(MQTT_USER)
	opts.SetPassword(MQTT_PASSWORD)
	mqttClient := mqtt.NewClient(opts)

	// Koneksi ke MQTT Broker
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	// Callback untuk menerima pesan
	mqttClient.Subscribe(MQTT_TOPIC, 0, func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("Received message: %s from topic: %s", msg.Payload(), msg.Topic())

		// Definisikan struct untuk mem-parsing JSON
		var payload struct {
			WaterLevel int `json:"water_level"`
		}

		// Parse JSON ke dalam struct
		err := json.Unmarshal(msg.Payload(), &payload)
		if err != nil {
			log.Printf("Error parsing JSON: %v", err)
			return
		}

		// Simpan data ke InfluxDB
		point := influxdb2.NewPoint(
			"water_level",
			map[string]string{"location": "sungaiA"},
			map[string]interface{}{"level": payload.WaterLevel}, // Gunakan nilai dari struct
			time.Now(),
		)
		err = writeAPI.WritePoint(context.Background(), point)
		if err != nil {
			log.Printf("Error writing to InfluxDB: %v", err)
		} else {
			log.Println("Data written to InfluxDB successfully")
		}
	})

	// Loop agar program tetap berjalan
	select {}
}
