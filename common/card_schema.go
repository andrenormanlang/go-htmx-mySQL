package common

// CardSchema represents a schema for a card in the system.
type CardSchema struct {
	// Title is the title of the schema.
	Title string `json:"title"`
	// Schema holds the JSON schema definition.
	Schema string `json:"schema"`
	// Cards holds a list of card UUIDs associated with the schema.
	Cards []string `json:"cards"`
}
