// Package objectiface provides the basic contract that any object must satisfy to exist within an environment.
package objectiface

// An Objecter is anything that knows how to be an object.
type Objecter interface {
	// ID returns a value that identifies this object uniquely from other
	// objects.
	ID() string
}
