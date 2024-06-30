package tempest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"time"
)

// A Network represents a network of hubs and sensors.
// The network will add hubs and sensors as messages
// are received from the network.
//
// Tempest hubs broadcast messages they receive from
// sensors paired with the hub on the network.
//
// Sensors transmit data to the hub via radio.
// The hub broadcasts the messages to the network
// on the broadcast address on port 50222 using UDP.
type Network struct {
	// Name of the network (e.g. "tempest").
	NetworkName string // Name of the network.

	// Hubs on the network.
	// The key is the hub serial number.
	// The network will add hubs as messages are received.
	hubs map[string]*Hub

	// UDP connection to the network.
	udpConn *net.UDPConn

	// Channel to stop the network.
	stop chan struct{}

	// Channel to process raw messages.
	networkMessage chan networkMessage

	// MessageRepo to store messages.
	messageRepo MessageRepo

	// SensorRepo to store sensors.
	sensorRepo SensorRepo

	// HubRepo to store hubs.
	hubRepo HubRepo
}

// networkMessage is a message received on the network.
type networkMessage struct {
	raw   RawMessage
	hubIp net.IP
}

// NewNetwork returns a new network with the given name.
func NewNetwork(name string) *Network {
	return &Network{
		NetworkName: name,
		hubs:        make(map[string]*Hub),
		stop:        make(chan struct{}),
	}
}

// Start the network activates the network to listen for
// messages from the network.
func (n *Network) Start(ip net.IP) error {
	if err := n.connect(ip); err != nil {
		return err
	}

	// Start listening for messages.
	go func() {
		if err := n.listen(); err != nil {
			fmt.Printf("error listening for messages: %v\n", err)
		}
	}()

	messageProcessing := make(chan networkMessage)
	n.networkMessage = messageProcessing

	// Process messages.
	go func() {
		if err := n.processMessage(messageProcessing); err != nil {
			fmt.Printf("error processing messages: %v\n", err)
		}
	}()

	return nil
}

// Stop the network from listening for messages.
func (n *Network) Stop() error {
	close(n.stop)

	return nil
}

// Connect to local network.
// The network will listen for UDP packets on the interface
// with the given IP address.
// Use net.IPv4(0, 0, 0, 0) or net.IPv4zero to listen to all
// interfaces.
func (n *Network) connect(ip net.IP) error {
	conn, err := net.ListenUDP("udp",
		&net.UDPAddr{
			IP:   ip,
			Port: 50222,
		})
	if err != nil {
		return fmt.Errorf("failed to connect to network: %w", err)
	}

	n.udpConn = conn

	return nil
}

// handlePartialMessage handles a partial message received from the hub.
// The message is incomplete and may be missing the closing brace due to
// a network error.
func handlePartialMessage(messageBuffer *bytes.Buffer) {
	var message RawMessage
	err := json.Unmarshal(messageBuffer.Bytes(), &message)
	if err == nil {
		fmt.Printf("recovered partial message: %v\n", message)
	} else {
		fmt.Printf("unable to recover partial message: %s", string(messageBuffer.Bytes()))
	}
}

// processMessage processes a message received from the network.
func (n *Network) processMessage(networkMessage <-chan networkMessage) error {
	for {
		select {
		case <-n.stop:
			fmt.Println("stopping message processing")

			return nil

			// Process the message.
		case msg := <-networkMessage:
			msgType, err := msg.raw.Type()
			if err != nil {
				fmt.Printf("error processing message: %v\n", err)
				continue
			}

			switch msgType {
			case MessageTypeRapidWind:
				hubSerialNumber, err := msg.raw.HubSerial()
				if err != nil {
					fmt.Printf("error getting hub serial: %v\n", err)
					continue
				}

				sensorSerialNumber, err := msg.raw.SensorSerial()
				if err != nil {
					fmt.Printf("error getting sensor serial: %v\n", err)
					continue
				}

				hub := n.updateHub(hubSerialNumber, msg.hubIp)
				sensor := n.updateSensor(hubSerialNumber, sensorSerialNumber, msg.hubIp)

				rapidWindEvent, err := NewRapidWindEvent(msg.raw, sensor, hub)
				if err != nil {
					fmt.Printf("error creating rapid wind event: %v\n", err)
					continue
				}

				fmt.Printf("Wind speed (m/s) - %0f", rapidWindEvent.WindSpeed)

			default:
				fmt.Printf("unhandled message type: %s\n", msgType)
			}
		}
	}
}

// addHub adds a hub to the network.
func (n *Network) addHub(serialNumber string, ip net.IP) {
	n.hubs[serialNumber] = &Hub{
		HubSerialNumber: serialNumber,
		WeatherSensors:  make(map[string]*WeatherSensor),
		LastReported:    time.Now(),
		IPAddress:       ip,
	}
}

// updateHub updates the hub on the network.
func (n *Network) updateHub(serialNumber string, ip net.IP) *Hub {
	hub, found := n.hubs[serialNumber]
	if !found {
		fmt.Printf("hub not found, adding hub: %s\n", serialNumber)
		n.addHub(serialNumber, ip)
		return n.hubs[serialNumber]
	}

	hub.IPAddress = ip
	hub.LastReported = time.Now()

	return hub
}

// addSensor adds a sensor to the network.
func (n *Network) addSensor(hubSerial string, sensorSerial string, ip net.IP) {
	hub, hubFound := n.hubs[hubSerial]
	if !hubFound {
		// No hub found, add the hub and sensor.
		n.addHub(hubSerial, ip)
		hub = n.hubs[hubSerial]

		hub.WeatherSensors[sensorSerial] = &WeatherSensor{
			SensorSerial: sensorSerial,
			LastMessage:  time.Now(),
		}

		return
	}

	// Hub found, add the sensor.
	hub.WeatherSensors[sensorSerial] = &WeatherSensor{
		SensorSerial: sensorSerial,
		LastMessage:  time.Now(),
	}
}

// updateSensor updates the sensor on the network.
func (n *Network) updateSensor(hubSerial string, sensorSerial string, ip net.IP) *WeatherSensor {
	if hub, hubFound := n.hubs[hubSerial]; hubFound {
		sensor, sensorFound := hub.WeatherSensors[sensorSerial]
		if !sensorFound {
			// No sensor found, add the sensor.
			n.addSensor(hubSerial, sensorSerial, ip)
			return n.hubs[hubSerial].WeatherSensors[sensorSerial]
		}

		sensor.LastMessage = time.Now()
		return sensor
	}

	// No hub found, add the hub and sensor.
	n.addSensor(hubSerial, sensorSerial, ip)

	return n.hubs[hubSerial].WeatherSensors[sensorSerial]
}

// listen for messages on the network.
func (n *Network) listen() error {
	// Create a buffer to read messages into.
	buffer := make([]byte, 1024)
	var messageBuffer bytes.Buffer
	var data []byte
	openBraces := 0

	// Listen for messages.
	for {
		select {
		case <-n.stop:
			// Stop listening for messages and close the connection.
			errorClosing := n.udpConn.Close()
			return errorClosing
		default:
			// Read from the connection.
			err := n.udpConn.SetReadDeadline(time.Now().Add(30 * time.Second))
			if err != nil {
				return fmt.Errorf("error setting read deadline: %w", err)
			}

			bytesRead, hubIp, err := n.udpConn.ReadFromUDP(buffer)
			if err != nil {
				var e net.Error
				if errors.As(err, &e) && e.Timeout() {
					if messageBuffer.Len() > 0 {
						fmt.Println("connection lost mid-message, attempting to recover")
						handlePartialMessage(&messageBuffer)
						messageBuffer.Reset()
						break
					}
				}

				return fmt.Errorf("error reading from UDP connection: %w", err)
			}

			// Write the message to the buffer.
			messageBuffer.Write(buffer[:bytesRead])
			data = messageBuffer.Bytes()

			// Process the message.
			for len(data) > 0 {
				startIndex := bytes.IndexByte(data, '{')
				if startIndex == -1 {
					messageBuffer.Reset()
					break
				}

				// Scan through the buffer to find the end of the message.
				for i := startIndex; i < len(data); i++ {
					if data[i] == '{' {
						openBraces++
					} else if data[i] == '}' {
						openBraces--
						if openBraces == 0 {
							message := data[startIndex : i+1]
							data = data[i+1:]

							// Try to get the message
							var msg RawMessage
							err := json.Unmarshal(message, &msg)
							if err == nil {
								fmt.Printf("received message: %v\n", msg)
							} else {
								fmt.Printf("unable to unmarshal message: %v\n", err)
							}

							if msgType, err := msg.Type(); err == nil && msgType.Valid() {
								fmt.Printf("received message type: %s\n", msgType)
								n.networkMessage <- networkMessage{
									raw:   msg,
									hubIp: hubIp.IP,
								}
							} else {
								fmt.Printf("invalid message type: %v\n", err)
							}

							messageBuffer.Reset()
							messageBuffer.Write(data)
							break
						}
					}
				}

				if openBraces > 0 {
					break
				}
			}
		}
	}
}
