package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jecolasurdo/marsrover/pkg/environment"
	"github.com/jecolasurdo/marsrover/pkg/environment/environmentiface"
	"github.com/jecolasurdo/marsrover/pkg/missioncontrol"
	"github.com/jecolasurdo/marsrover/pkg/objects"
	"github.com/jecolasurdo/marsrover/pkg/spatial"
	"github.com/spf13/cobra"
)

type roverBuilder struct{}

func (*roverBuilder) LaunchRover(h spatial.Heading, p spatial.Point, env environmentiface.Environmenter) (RoverAPI, error) {
	return objects.Rover{}.LaunchRover(h, p, env)
}

type envBuilder struct{}

func (*envBuilder) NewEnvironment(p spatial.Point) Environmenter {
	return environment.Plateau{}.NewPlateau(p)
}

var rootCmd = &cobra.Command{
	Use:   "marsrover",
	Short: "A system that simulates exploring mars.",
	RunE: func(cmd *cobra.Command, args []string) error {
		mission := missioncontrol.NewMission(new(roverBuilder), nil)

		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return error
		}

		commands := strings.Split(string(data), "\n")
		stats, err := mission.ExecuteMission(commands)
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
