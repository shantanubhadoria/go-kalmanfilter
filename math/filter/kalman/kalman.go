/*

Introduction

Package Kalman implements Kalman Filter(Linear Quadratic 
Estimation) support for Go language

Source and Bug reports at 

  https://github.com/shantanubhadoria/go-math-filter-kalman

*/
package kalman


/*
Angle: the state sensor value. In a IMU this would be the 
Accelerometer
*/ 
var Angle float64

/*
Bias: the delta sensor calculation. This is the deviation 
from last base state value as calculted from the delta 
sensor. In a IMU this would be the product of time since 
last reading and the delta sensor value
starting value(default): 0

Bias is recalculated(optimised) at each new sensor reading.
*/
var Bias float64

/*
Covariance Matrix a 2d 2x2 matrix (also known as dispersion 
matrix or variance-covariance matrix) is a matrix whose 
element in the i, j position is the covariance between the i
and j elements of a random vector.

  Default: [[0,0][0,0]]
*/
var Covariance [2][2]float64

var QAngle float64 = 0.001
var QBias float64 = 0.003
var RMeasure float64 = 0.003
