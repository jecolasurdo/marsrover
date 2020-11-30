# Mars Rover
[![PkgGoDev](https://pkg.go.dev/badge/github.com/jecolasurdo/marsrover)](https://pkg.go.dev/github.com/jecolasurdo/marsrover)
[![Build Status](https://travis-ci.org/jecolasurdo/marsrover.svg?branch=master)](https://travis-ci.org/jecolasurdo/marsrover)
[![Go Report Card](https://goreportcard.com/badge/github.com/jecolasurdo/marsrover)](https://goreportcard.com/report/github.com/jecolasurdo/marsrover)
[![Maintainability](https://api.codeclimate.com/v1/badges/42f4814205a407bf7ca1/maintainability)](https://codeclimate.com/github/jecolasurdo/marsrover/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/42f4814205a407bf7ca1/test_coverage)](https://codeclimate.com/github/jecolasurdo/marsrover/test_coverage)

## Specifications

A squad of robotic rovers are to be landed by NASA on a plateau on Mars.
This plateau, which is curiously rectangular, must be navigated by the
rovers so that their on-board cameras can get a complete view of the
surrounding terrain to send back to Earth.

A rover's position and location is represented by a combination of x and y
co-ordinates and a letter representing one of the four cardinal compass
points. The plateau is divided up into a grid to simplify navigation. An
example position might be 0, 0, N, which means the rover is in the bottom
left corner and facing North.

In order to control a rover, NASA sends a simple string of letters. The
possible letters are 'L', 'R' and 'M'. 'L' and 'R' makes the rover spin 90
degrees left or right respectively, without moving from its current spot.
'M' means move forward one grid point, and maintain the same heading.

Assume that the square directly North from (x, y) is (x, y+1).

### Input 
The first line of input is the upper-right coordinates of the plateau, the
lower-left coordinates are assumed to be 0,0.

The rest of the input is information pertaining to the rovers that have
been deployed. Each rover has two lines of input. The first line gives the
rover's position, and the second line is a series of instructions telling
the rover how to explore the plateau.

The position is made up of two integers and a letter separated by spaces,
corresponding to the x and y co-ordinates and the rover's orientation.

Each rover will be finished sequentially, which means that the second rover
won't start to move until the first one has finished moving, and each rover
stays on the plateau once finished.

### Output
The output for each rover should be its final co-ordinates and heading.

### Example 1
**Input**
```
5 5
1 2 N
LMLMLMLMM
3 3 E
MMRMMRMRRM
```
**Output**
1 3 N
5 1 E

### Example 2
**Input**
```
3 3
0 0 S
LMMLM
1 2 W
LMLMRM
```

**Output**
```
2 1 N
1 0 S
```

## Assumptions
In some cases, the specification was not clear about certain behavior.
Assumptions regarding the behavior of the system are documented here.

### Handling placing or moving a rover out of bounds 
- The specification isn't clear about what to do if a rover is initialized outside
the bounds of the environment. It is also not clear about how a rover should
behave if it moves out of the bounds of the environment.
- It seemed reasonable that a rover should not be allowed to be placed outside
the bounds of the environment, so this is prohibited, and will return an error.
- It also seemed reasonable that a rover not allow itself to wander outside the
boundaries of its environment. Thus, if a rover is commanded to move beyond the
boundaries of the environment, the rover will not move, and will instead return
an error.

### Handling moving a rover into a space already occupied by another rover
- The main specifications do not address this scenario, but the specification's 
second example alludes to a presumption that if a rover encounters another rover
, it should just fail to move into the space occupied by the second rover.
- This seems like a reasonable presumption, so rovers are currently designed
such that they will return an error if commanded to move into a position occupied
by another rover.

### Negative plateau dimensions
- The specification does not state whether the dimensions of the plateau must be
expressed as positive values (though this seems reasonable).
- The current system does not deny the creation of plateau's with
negative dimensions, nor does it deny the ability to place a rover in a negative
position (if the environment is also in negative space). However, the system
isn't explicitly tested for these scenarios, so the behavior in that space remains
be undefined.

## System Architecture
The marsrover system is composed of two primary components:
1. The marsrover API: a library which handles the majority of system behavior
1. The marsrover CLI: a command line interface which allows commands to be executed via stdin.

### API
The API is composed of the following top level components:
1. Environments: An environment (such as a Plateau) has dimensions and contains objects (such as Rovers)
1. Objects: Objects (such as Rovers) are any discrete thing that can interact with an environment (or other objects)
1. Mission Control: The high level API responsible solely for establishing environments and objects via a series of text commands. Mission Control is primarily responsible for parsing command input, and marshalling results between the internal API and some other interface (such as a CLI or rest API)

### CLI
The CLI is a thin command line interface that allows commands from a systems stdin be passed into
mission control, and, conversly, allow mission control to report results back via stdin (or stderr).

Example 1:
```
$ cat ./sampleinput/testinput1.txt | ./marsrover
1 3 N
5 1 E
$
```

Example 2:
```
$ cat ./sampleinput/testinput2.txt | ./marsrover
2 1 N
1 0 S
$
```