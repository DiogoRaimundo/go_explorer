package main

import (
	goTour02 "go_explorer/02_flow_control"
	goTour03 "go_explorer/03_more_types"
	goTour04 "go_explorer/04_methods_and_interfaces"
	goTour05 "go_explorer/05_generics"
	goTour06 "go_explorer/06_concurrency"
	albumRestApi "go_explorer/album_rest_api"
	fileReadWrite "go_explorer/file_read_write"
	nRequestPerMinuteExercise "go_explorer/n_request_per_minute_exercise"
)

const (
	GoTour02_Example08 int = iota
	GoTour02_Example09
	GoTour02_Example12

	GoTour03_Exercise18
	GoTour03_Exercise23
	GoTour03_Exercise26

	GoTour04_Example00_03
	GoTour04_Example04_08
	GoTour04_Example09_14
	GoTour04_Example15_16
	GoTour04_Exercise18
	GoTour04_Exercise20
	GoTour04_Exercise22_23
	GoTour04_Exercise25

	GoTour05_Exercise

	GoTour06_Example01
	GoTour06_Example02_04
	GoTour06_Example05_06
	GoTour06_Exercise07_08
	GoTour06_Example09
	GoTour06_Example10

	AlbumRestApi_AlbumGinApi
	AlbumRestApi_AlbumStandardApi

	FileReadWrite

	NRequestPerMinuteExercise

	All_Options
)

type Runnable struct {
	Name string
	Run  func()
}

var options = []Runnable{
	{"goTour02.RunExample08", goTour02.RunExample08},
	{"goTour02.RunExample09", goTour02.RunExample09},
	{"goTour02.RunExample12", goTour02.RunExample12},

	{"goTour03.RunExercise18", goTour03.RunExercise18},
	{"goTour03.RunExercise23", goTour03.RunExercise23},
	{"goTour03.RunExercise26", goTour03.RunExercise26},

	{"goTour04.RunExample00_03", goTour04.RunExample00_03},
	{"goTour04.RunExample04_08", goTour04.RunExample04_08},
	{"goTour04.RunExample09_14", goTour04.RunExample09_14},
	{"goTour04.RunExample15_16", goTour04.RunExample15_16},
	{"goTour04.RunExercise18", goTour04.RunExercise18},
	{"goTour04.RunExercise20", goTour04.RunExercise20},
	// {"goTour04.RunExercises22_23", goTour04.RunExercises22_23},
	{"goTour04.RunExercises22_23", printUnableToCompile},
	{"goTour04.RunExercise25", goTour04.RunExercise25},

	{"goTour05.RunExercise", goTour05.RunExercise},

	{"goTour06.RunExample01", goTour06.RunExample01},
	{"goTour06.RunExample02_04", goTour06.RunExample02_04},
	{"goTour06.RunExample05_06", goTour06.RunExample05_06},
	{"goTour06.RunExercise07_08", goTour06.RunExercise07_08},
	{"goTour06.RunExample09", goTour06.RunExample09},
	{"goTour06.RunExample10", goTour06.RunExample10},

	{"albumRestApi.RunAlbumGinApi", albumRestApi.RunAlbumGinApi},
	{"albumRestApi.RunAlbumStandardApi", albumRestApi.RunAlbumStandardApi},

	{"fileReadWrite.Run", fileReadWrite.Run},

	{"nRequestPerMinuteExercise.Run", nRequestPerMinuteExercise.Run},
}
