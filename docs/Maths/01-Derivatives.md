# Mathematics Notation Conventions for Functions

When writing mathematical functions, there are conventions that help standardize their representation and make them easily understandable. These are not strict rules but guidelines commonly adopted in mathematical literature.

## Function Notation

The most common way to denote a function is by using the letter $f$ followed by the independent variable in parentheses, such as $f(x)$. This represents a function named $f$ with $x$ as its input. The output, or dependent variable, is typically written as $y$. Thus, $y = f(x)$ implies that $y$ is the result of applying function $f$ to $x$.

### Example:

If $f(x) = x^2$, then the statement $y = f(x)$ is read as "y equals x squared," indicating that the function $f$ takes $x$ and squares it to produce the output $y$.

# Derivatives: An Introduction

In calculus, a derivative represents the rate at which a function changes as its input changes. Loosely speaking, it's the slope of the function at any given point.

## Concept of Derivatives

The derivative of a function at a point is the limit of the average rate of change of the function in a tiny interval around that point. If the function is graphed, the derivative corresponds to the slope of the tangent line at a point on the curve.

## Notations for Derivatives

There are several notations for derivatives, including:

- Leibniz's notation: $\frac{dy}{dx}$ or $\frac{df}{dx}(x)$
- Lagrange's notation: $f'(x)$
- Newton's notation: $\dot{y}$, often used for time derivatives

# Detailed Explanation and Examples of Derivatives

## Constant Function

Consider $f(x) = c$, where $c$ is a constant.

### Derivative:

For any constant function, the derivative is 0, since the function does not change as $x$ changes.

### Example:

If $f(x) = 5$, then $f'(x) = 0$.

## Linear Function

Consider $f(x) = mx + b$, where $m$ and $b$ are constants.

### Derivative:

The derivative of a linear function is the coefficient of $x$, which is $m$.

### Example:

If $f(x) = 3x + 4$, then $f'(x) = 3$.

## Exponential Function

Consider $f(x) = e^x$, where $e$ is the base of the natural logarithm.

### Derivative:

The derivative of $e^x$ is $e^x$ itself.

### Example:

If $f(x) = e^x$, then $f'(x) = e^x$.

## Sigmoid Function

Consider the sigmoid function $\sigma(x) = \frac{1}{1+e^{-x}}$.

### Derivative:

The derivative of the sigmoid function is $\sigma'(x) = \sigma(x)(1 - \sigma(x))$.

### Example:

If $\sigma(x) = \frac{1}{1+e^{-x}}$, to find $\sigma'(x)$, we first calculate $\sigma(x)$ and then use it to find $\sigma(x)(1 - \sigma(x))$.

# Cheat Sheet of Derivatives

Here is a quick reference for the derivatives of some common functions:

- $\frac{d}{dx}c = 0$ (Constant function)
- $\frac{d}{dx}(cx) = c$ (Constant multiple of $x$)
- $\frac{d}{dx}(x^n) = nx^{n-1}$ (Power rule)
- $\frac{d}{dx}(e^x) = e^x$ (Exponential function with base $e$)
- $\frac{d}{dx}(\ln(x)) = \frac{1}{x}$ (Natural logarithm)
- $\frac{d}{dx}(\sin(x)) = \cos(x)$ (Sine function)
- $\frac{d}{dx}(\cos(x)) = -\sin(x)$ (Cosine function)
- $\frac{d}{dx}(\tan(x)) = \sec^2(x)$ (Tangent function)
- $\frac{d}{dx}(\sigma(x)) = \sigma(x)(1 - \sigma(x))$ (Sigmoid function)

This cheat sheet is a simplified guide and does not cover all the rules of differentiation, such as the product rule,

 quotient rule, or chain rule, which are essential when dealing with more complex functions.