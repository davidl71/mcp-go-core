package gosdk

import "fmt"

// ValidateRegistration validates common registration parameters
// for tools, prompts, and resources.
func ValidateRegistration(name, description string, handler interface{}) error {
	if name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if description == "" {
		return fmt.Errorf("description cannot be empty")
	}
	if handler == nil {
		return fmt.Errorf("handler cannot be nil")
	}
	return nil
}

// ValidateResourceRegistration validates resource-specific registration parameters
func ValidateResourceRegistration(uri, name, description string, handler interface{}) error {
	if err := ValidateRegistration(name, description, handler); err != nil {
		return err
	}
	if uri == "" {
		return fmt.Errorf("resource URI cannot be empty")
	}
	return nil
}
