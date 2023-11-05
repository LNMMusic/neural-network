package neuron

import "math"

// Edge represents an edge that holds info about the neuron and the weight
type Edge struct {
	// Weight is the weight of the edge
	Weight float64
	// Neuron is the neuron of the edge
	Neuron Neuron
}

// NeuronNode is a neuron node
type NeuronNode struct {
	// ActivationValue is the activation value of the neuron
	ActivationValue float64
	// Bias is the bias value of the neuron
	Bias float64
	// Edges are the edges of the neuron
	Edges []Edge
	// LearningRate is the learning rate of the neuron
	LearningRate float64
}

// Activate returns the activation value of the neuron
func (n *NeuronNode) Activate() (a float64) {
	// calculate the activation value
	// - calculate the sum of the weights * activation values of the edges
	var weightedSum float64
	for _, edge := range n.Edges {
		weightedSum += edge.Weight * edge.Neuron.Activate()
	}

	// - add the bias
	weightedSum += n.Bias

	// - calculate the activation value
	a = 1.0 / (1.0 + math.Exp(-weightedSum))

	// set the activation value
	n.ActivationValue = a

	return
}

// Optimize optimizes the neuron
func (n *NeuronNode) Optimize(o OptimizeInfo) (err error) {
	// calculations
	// - diff is the difference between the target and the activation value
	diff := n.ActivationValue - o.Target
	// - delta is the delta value
	delta := diff * n.ActivationValue * (1.0 - n.ActivationValue)
	// - alpha is the learning rate
	alpha := n.LearningRate

	// updates
	// - update the bias
	n.Bias -= alpha * delta

	// - update the weights
	for _, edge := range n.Edges {
		edge.Weight -= alpha * delta * edge.Neuron.Activate()
	}

	return
}