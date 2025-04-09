package osenv

import (
	"os"
)

// # MQTT broker
// MQTT_BROKER=
// MQTT_USERNAME=
// MQTT_PASSWORD=
// MQTT_TOPIC=
func GetMQTTBroker() string {
	return os.Getenv("MQTT_BROKER")
}
func GetMQTTUsername() string {
	return os.Getenv("MQTT_USERNAME")
}
func GetMQTTPassword() string {
	return os.Getenv("MQTT_PASSWORD")
}
func GetMQTTTopic() string {
	return os.Getenv("MQTT_TOPIC")
}
