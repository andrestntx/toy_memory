package memory

type Location struct {
	Lat float32
	Lng float32
}

type ToyMemoryInterface interface {
	Put(orderId string, location Location)
	Get(orderId string, max int) []Location
	Delete(orderId string)
}

var ToyMemory ToyMemoryInterface

func init() {
	ToyMemory = &toyMemory{memory: make(map[string][]Location)}
}

type toyMemory struct {
	memory map[string][]Location
}

func (t *toyMemory) Put(orderId string, location Location) {
	t.memory[orderId] = append(t.memory[orderId], location)
}

func (t *toyMemory) Get(orderId string, max int) []Location {
	locations := t.memory[orderId]

	if locations != nil && len(locations) > max && max > 0 {
		return locations[0:max]
	}

	return locations
}

func (t *toyMemory) Delete(orderId string) {
	delete(t.memory, orderId)
}
