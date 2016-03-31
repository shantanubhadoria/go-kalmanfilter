# kalman
--
    import "."

### Introduction

Package Kalman implements Kalman Filter(Linear Quadratic Estimation) support for
Go language

Source and Bug reports at

    https://github.com/shantanubhadoria/go-math-filter-kalman


### Description

The Kalman filter(https://en.wikipedia.org/wiki/Kalman_filter), also known as
linear quadratic estimation (LQE), is an algorithm that uses a series of
measurements observed over time, containing noise (random variations) and other
inaccuracies, and produces estimates of unknown variables that tend to be more
precise than those based on a single measurement alone.

Algorithm is recursive, which means it takes the output of its previous
calculations as a factor in calculating the next step which improves its
accuracy over time. The key to Kalman filters are two sensors with different
kind of accuracy issues in each. Sensor A or the state sensor might give
in-accurate value for a measurement on the whole but it doesn't drift. Sensor B
or delta sensor gives gives much more accurate rate of change in value(or delta)
but it drifts over time due to its small inaccuracies as it only measures rate
of change in value and not the actual value. Kalman filter uses this knowledge
to fuse results from both sensors to give a state value which is more accurate
than state value received from any of these filters alone.

An example of application for this is calculating orientation of objects using
Gyroscopes and Accelerometers.

While Accelerometer is usually used to measure gravity it can be used to measure
the inclination of a body with respect to the surface of earth along the x and y
axis(not z axis as Z axis is usually facing the opposite direction as the force
of gravity) by measuring the direction in which the force of gravity is applied.

Gyroscope measures the rate of rotation about one or all the axis of a body.
while it gives fairly accurate estimation of the angular velocity, if we use it
to calculate the current inclination based on the starting inclination and the
angular velocity, there is a lot of drift, which means the gyroscope error will
accumulate over time as we calculate newer angles based on previous angle and
angular velocity and the error in angular velocity piles on.

A real life example of how Kalman filter works is while driving on a highway in
a car. If you take the time passed since when your started driving and your
estimated average speed every hour and use it to calculate the distance you have
traveled your calculation will become more inaccurate as you drive on.

This is drift in value. However if you watch each milestone and calculate your
current position using milestone data and your speed since the last milestone
your result will be much more accurate. That is approximately close to how
Kalman filter works.

## Usage

```go
var Angle float64
```
Angle the state sensor value. In a IMU this would be the Accelerometer

```go
var Bias float64
```
Bias: the delta sensor calculation. This is the deviation from last base state
value as calculted from the delta sensor. In a IMU this would be the product of
time since last reading and the delta sensor value starting value(default): 0

Bias is recalculated(optimised) at each new sensor reading.

```go
var Covariance [2][2]float64
```
Covariance Matrix a 2d 2x2 matrix (also known as dispersion matrix or
variance-covariance matrix) is a matrix whose element in the i, j position is
the covariance between the i and j elements of a random vector.

    Default: [[0,0][0,0]]

```go
var QAngle float64 = 0.001
```

```go
var QBias float64 = 0.003
```

```go
var RMeasure float64 = 0.003
```
