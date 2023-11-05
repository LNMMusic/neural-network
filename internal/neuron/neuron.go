package neuron

// OptimizeInfo is the info for optimizing a neuron
type OptimizeInfo struct {
	// Target is the target value
	Target float64
}

// Neuron is an interface representing a neuron
type Neuron interface {
	// Activate returns the activation value of the neuron
	Activate() (a float64)

	// Optimize optimizes the neuron
	Optimize(o OptimizeInfo) (err error)
}