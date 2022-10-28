package main

import (
	"fmt"
	"strings"

	goTour02 "go_explorer/02_flow_control"
	goTour03 "go_explorer/03_more_types"
	goTour04 "go_explorer/04_methods_and_interfaces"
	goTour05 "go_explorer/05_generics"
	goTour06 "go_explorer/06_concurrency"
	albumRestApi "go_explorer/album_rest_api"
)

func announceAndRun(name string, run func()) {
	announce := "---===== Running " + name + " =====---"
	fmt.Println()
	fmt.Println(announce)

	run()

	fmt.Println("---" + strings.Repeat("=", len(announce)-6) + "---")
	fmt.Println()
}

func main() {
	announceAndRun("goTour02.RunExample08", goTour02.RunExample08)
	announceAndRun("goTour02.RunExample09", goTour02.RunExample09)
	announceAndRun("goTour02.RunExample12", goTour02.RunExample12)

	announceAndRun("goTour03.RunExercise18", goTour03.RunExercise18)
	announceAndRun("goTour03.RunExercise23", goTour03.RunExercise23)
	announceAndRun("goTour03.RunExercise26", goTour03.RunExercise26)

	announceAndRun("goTour04.RunExample00_03", goTour04.RunExample00_03)
	announceAndRun("goTour04.RunExample04_08", goTour04.RunExample04_08)
	announceAndRun("goTour04.RunExample09_14", goTour04.RunExample09_14)
	announceAndRun("goTour04.RunExample15_16", goTour04.RunExample15_16)
	announceAndRun("goTour04.RunExercise18", goTour04.RunExercise18)
	announceAndRun("goTour04.RunExercise20", goTour04.RunExercise20)
	announceAndRun("goTour04.RunExercises22_23", goTour04.RunExercises22_23)
	announceAndRun("goTour04.RunExercise25", goTour04.RunExercise25)

	announceAndRun("goTour05.RunExercise", goTour05.RunExercise)

	announceAndRun("goTour06.RunExample01", goTour06.RunExample01)
	announceAndRun("goTour06.RunExample02_04", goTour06.RunExample02_04)
	announceAndRun("goTour06.RunExample05_06", goTour06.RunExample05_06)
	announceAndRun("goTour06.RunExercise07_08", goTour06.RunExercise07_08)
	announceAndRun("goTour06.RunExample09", goTour06.RunExample09)
	announceAndRun("goTour06.RunExample10", goTour06.RunExample10)

	announceAndRun("albumRestApi.RunAlbumApi", albumRestApi.RunAlbumApi)
}
