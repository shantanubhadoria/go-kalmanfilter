# kalmanfilter
--
    import "github.com/shantanubhadoria/go-kalmanfilter/kalmanfilter"

Package kalmanfilter implements Kalman Filter(Linear Quadratic Estimation)
support for Go language

[![Travis CI](https://img.shields.io/travis/shantanubhadoria/go-kalmanfilter.svg?style=flat-square)](https://travis-ci.org/shantanubhadoria/go-kalmanfilter) [![GoDoc](https://godoc.org/github.com/shantanubhadoria/go-kalmanfilter/kalmanfilter?status.svg)](https://godoc.org/github.com/shantanubhadoria/go-kalmanfilter/kalmanfilter)

### Introduction

Source and Bug reports at https://github.com/shantanubhadoria/go-kalmanfilter

### Synopsis

    package main

    import (
      "fmt"
      "time"
      "github.com/shantanubhadoria/go-kalmanfilter/"
    )

    myFilterData = new(kalmanfilter.FilterData)

    var oldTime time.Time = time.Now()
    for {
      stateReading := float64(getStateSensorReading()) // in units X
      deltaReading := float64(getDeltaSensorReading()) // in unit X per nanosecond

      var newTime time.Time = time.Now()
      var duration Duration = newTime.Sub(oldTime)
      oldTime = newTime
      newState := myFilterData.Update(stateReading, deltaReading, int64(duration/time.Nanosecond))
      fmt.Println(newState)
    }


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
axis(not z axis as Z axis faces the direction opposite the direction of
gravitional force) by measuring the direction in which the force of gravity is
felt.

Gyroscope measures the rate of rotation about one or all the axes of a body.
While it gives fairly accurate estimation of the angular velocity, if we use it
to calculate the current inclination based on the starting inclination and the
angular velocity, there is a lot of drift, which means the gyroscope error will
accumulate over time as we calculate newer angles based on previous angle and
angular velocity and the error in angular velocity piles on leading to
increasingly inaccurate estimations as time passes.

A real life example of how Kalman filter works is noticed while driving on a
highway in a car. If you take the time passed since when your started driving
and your estimated average speed since then and use it to calculate the distance
you have traveled your calculation will become more inaccurate as time passes.

This is drift in value. However if you correct based on each milestone marker
that you pass through and re-calculate your distance travelled using milestone
data and your average speed since you pass the last milestone your result will
be much more accurate irrespective of how much time has passed. That is
approximately close to how Kalman filter and sensor fusion work.

State Sensor: ![Milestone](/corpus/milestone.jpg)

Delta Sensor: ![Speedometer](/corpus/speedometer.png)


### Author

Shantanu Bhadoria <shantanu att cpan dot org> https://www.shantanubhadoria.com

## Usage

#### type FilterData

```go
type FilterData struct {

	/*
	   State the state sensor value. In a IMU this would be the
	   Accelerometer
	*/
	State float64

	/*
	   Bias: the delta sensor error. This is the deviation
	   from sensor reading and actual value. Bias can be caused by
	   electromagnetic interference and represents a permanent error
	   in delta sensor reading. Bias is detected by averaging the
	   delta sensor reading at stationary state of delta sensor
	*/
	Bias float64

	/*
	   Covariance Matrix a 2d 2x2 matrix (also known as dispersion
	   matrix or variance-covariance matrix) is a matrix whose
	   element in the i, j position is the covariance between the i
	   and j elements of a random vector. Leave this at default
	   value of [[0,0],[0,0]]
	*/
	Covariance [2][2]float64

	QAngle   float64
	QBias    float64
	RMeasure float64
}
```

FilterData struct, initialize this struct before commencing any operations, as
sensors are read, this struct must be updated alongside

#### func (*FilterData) Update

```go
func (filterData *FilterData) Update(stateReading, deltaReading, deltaTime float64) float64
```
Update Method Call this method to update the state value based on sensor fusion
of state and delta sensor and the previously calculate reading to get
progressively more accurate state value
