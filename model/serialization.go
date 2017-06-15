package model

// Serializable is an interface representing an object that can be serialized to JSON
type Serializable interface {
	Serialize() ([]byte, error)
}

// Deserializable is an interface representing JSON data that can be turned into an object
type Deserializable interface {
	Deserialize(interface{}) error
}
