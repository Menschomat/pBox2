# pBox2
pBox2 is a software application written in Golang for controlling 3D printer enclosures. It uses the go-chi framework and an Angular dashboard.

## Features
* Connects to Pi Picos to control hardware in the enclosures, such as lights, fans, and sensors
* Pi Picos are connected via MQTT
* Dashboard is connected via WebSockets
## Installation
1. Clone the repository: `git clone https://github.com/Menschomat/pBox2.git`
2. Navigate to the project directory: `cd pBox2`
3. Install dependencies: `go mod download`
## Usage
1. Start the application: go run main.go
2. Open your browser and navigate to http://localhost:8080 to access the dashboard
## Configuration

- `broker` (string): The address of the MQTT broker to connect to.
- `port` (int): The port number to use for the MQTT broker connection.
- `client_id` (string): The client ID to use for the MQTT connection.
- `username` (string): The username to use for the MQTT connection.
- `password` (string): The password to use for the MQTT connection.
- `topic` (string): The MQTT topic to publish data to.

### Enclosure Configuration

- `id` (string): The unique identifier of the enclosure.
- `name` (string): The name of the enclosure.
- `location` (string): The physical location of the enclosure.
- `boxes` (array of objects): An array of boxes contained within the enclosure.

#### Box Configuration

- `id` (string): The unique identifier of the box.
- `name` (string): The name of the box.
- `location` (string): The physical location of the box.
- `lights` (array of objects): An array of lights contained within the box.
- `fans` (array of objects): An array of fans contained within the box.
- `sensors` (array of objects): An array of sensors contained within the box.

##### Light Configuration

- `id` (string): The unique identifier of the light.
- `name` (string): The name of the light.
- `type` (string): The type of light (e.g., RGB, mono).

##### Fan Configuration

- `id` (string): The unique identifier of the fan.
- `name` (string): The name of the fan.

##### Sensor Configuration

- `id` (string): The unique identifier of the sensor.
- `name` (string): The name of the sensor.
- `type` (string): The type of sensor (e.g., temperature, humidity).
- `unit` (string): The unit of measurement for the sensor data.

```
Config-Structure
├── mqtt
│   ├── broker
│   ├── port
│   ├── client_id
│   ├── username
│   ├── password
│   └── topic
└── enclosure
    ├── id
    ├── name
    ├── location
    └── boxes
        ├── id
        ├── name
        ├── location
        ├── lights
        │   ├── id
        │   ├── name
        │   └── type
        ├── fans
        │   ├── id
        │   └── name
        └── sensors
            ├── id
            ├── name
            ├── type
            └── unit
```