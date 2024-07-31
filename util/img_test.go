package util

import (
	"fmt"
	"testing"
)

func TestResize(t *testing.T) {
	//resizeImg("../img/goblin/run/run_000.png", "./resizeRun000.png")
	resizeImg2("../img/goblin/run/run_000.png", "./resizeRun000.png")
}

func TestBatchResize(t *testing.T) {
	for i := 0; i < 40; i++ {
		inputPath := fmt.Sprintf("../img/goblin/died/died_%03d.png", i)
		outputPath := fmt.Sprintf("../img/goblin/died/smallDied_%03d.png", i)
		resizeImg2(inputPath, outputPath)
	}
}
