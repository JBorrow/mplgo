package mplgo

import (
	"bytes"
	"image/color"
	"os/exec"
	"strconv"
	"strings"
)

var PY_EXTRACTOR = "import matplotlib.pyplot as plt; print(plt.get_cmap('CMAP_NAME')([float(x) / float(STEPS - 1) for x in range(STEPS)]))"

func GetCmap(cmapName string, steps int) (ColorMap, error) {
	return GetCmapCustom(PY_EXTRACTOR, cmapName, steps, 4)
}

func GetCmapCustom(baseExtractor string, cmapName string, steps int, cols int) (ColorMap, error) {
	colorMap := ColorMap{steps: steps, fsteps: float64(steps), name: cmapName}

	stdout, _, err := runPython(baseExtractor, cmapName, steps)

	if err != nil {
		return colorMap, err
	}

	floatMap, err := rgbaArrayFromStdout(stdout, steps, cols)

	if err != nil {
		return colorMap, err
	}

	intMap := make([]color.RGBA, steps)

	for i := range intMap {
		intMap[i] = color.RGBA{
			uint8(floatMap[i][0] * 255.0),
			uint8(floatMap[i][1] * 255.0),
			uint8(floatMap[i][2] * 255.0),
			uint8(floatMap[i][3] * 255.0),
		}
	}

	colorMap.data = intMap

	return colorMap, nil
}

func runPython(baseExtractor string, cmapName string, steps int) (string, string, error) {
	extractor := strings.ReplaceAll(
		strings.ReplaceAll(
			baseExtractor, "STEPS", strconv.Itoa(steps),
		), "CMAP_NAME", cmapName,
	)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("python", "-c", extractor)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}

func rgbaArrayFromLine(line string, length int) ([]float64, error) {
	output := make([]float64, length)

	sanitizedLine := strings.ReplaceAll(strings.ReplaceAll(line, "[", ""), "]", "")

	for i, number := range strings.Fields(sanitizedLine) {
		value, err := strconv.ParseFloat(number, 64)

		if err != nil {
			return output, err
		}

		output[i] = float64(value)
	}

	return output, nil
}

func rgbaArrayFromStdout(stdout string, steps int, cols int) ([][]float64, error) {
	output := make([][]float64, steps)
	var err error

	for i, line := range strings.Split(strings.TrimSuffix(stdout, "\n"), "\n") {
		output[i], err = rgbaArrayFromLine(line, cols)

		if err != nil {
			return output, err
		}
	}

	return output, nil
}
