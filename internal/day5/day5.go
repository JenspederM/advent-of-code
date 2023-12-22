package day5

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/cloudflare/cfssl/log"
	"github.com/jenspederm/advent-of-code/internal/utils"
)

type RangeElements []RangeElement
type RangeElement struct {
	sourceStart      int
	destinationStart int
	rangeLength      int
}

func (res RangeElements) GetDestination(source int) int {
	for _, re := range res {
		if re.InRange(source) {
			return re.destinationStart + source - re.sourceStart
		}
	}
	return source
}

func (re RangeElement) InRange(source int) bool {
	return re.sourceStart <= source && source <= re.sourceStart+re.rangeLength
}

func (s RangeElement) String() string {
	return strconv.Itoa(s.sourceStart) + " -> " + strconv.Itoa(s.destinationStart) + " (" + strconv.Itoa(s.rangeLength) + ")"
}

func rangeFromLines(prefix string, lines []string, i int) (RangeElements, int, error) {
	if !strings.HasPrefix(lines[i], prefix) {
		return nil, i, errors.New("Prefix unmatched at line " + strconv.Itoa(i))
	}
	log.Debug("Found " + prefix)
	i++
	partOfRange := true
	elements := RangeElements{}
	for partOfRange && i < len(lines) {
		line := lines[i]
		if line == "" {
			partOfRange = false
			break
		} else {
			parts := strings.Split(line, " ")
			if len(parts) != 3 {
				return nil, i, errors.New("Invalid line at " + strconv.Itoa(i))
			}
			destination, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, i, errors.New("Invalid destination at " + strconv.Itoa(i))
			}
			source, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, i, errors.New("Invalid source at " + strconv.Itoa(i))
			}
			rangeLength, err := strconv.Atoi(parts[2])
			if err != nil {
				return nil, i, errors.New("Invalid rangeLength at " + strconv.Itoa(i))
			}

			elements = append(elements, RangeElement{sourceStart: source, destinationStart: destination, rangeLength: rangeLength})
		}
		i++
	}

	return elements, i, nil
}

func Parse(lines []string) ([]int, map[string]RangeElements) {
	var seeds []string
	var seedToSoilRange RangeElements
	var soilToFertilizerRange RangeElements
	var fertilizerToWaterRange RangeElements
	var waterToLightRange RangeElements
	var lightToTemperatureRange RangeElements
	var temperatureToHumidityRange RangeElements
	var humidityToLocationRange RangeElements

	i := 0
	for i < len(lines) {
		line := lines[i]
		if line == "" {
			i++
			continue
		}
		if strings.HasPrefix(line, "seeds:") {
			parts := strings.Split(line, ":")
			foundSeeds := strings.Split(parts[1], " ")
			for _, seed := range foundSeeds {
				trimmedSeed := strings.TrimSpace(seed)
				if trimmedSeed == "" {
					continue
				}
				seeds = append(seeds, trimmedSeed)
			}
		} else if _range, newI, err := rangeFromLines("seed-to-soil map:", lines, i); err == nil {
			seedToSoilRange = _range
			i = newI
		} else if _range, newI, err := rangeFromLines("soil-to-fertilizer map:", lines, i); err == nil {
			soilToFertilizerRange = _range
			i = newI
		} else if _range, newI, err := rangeFromLines("fertilizer-to-water map:", lines, i); err == nil {
			fertilizerToWaterRange = _range
			i = newI
		} else if _range, newI, err := rangeFromLines("water-to-light map:", lines, i); err == nil {
			waterToLightRange = _range
			i = newI
		} else if _range, newI, err := rangeFromLines("light-to-temperature map:", lines, i); err == nil {
			lightToTemperatureRange = _range
			i = newI
		} else if _range, newI, err := rangeFromLines("temperature-to-humidity map:", lines, i); err == nil {
			temperatureToHumidityRange = _range
			i = newI
		} else if _range, newI, err := rangeFromLines("humidity-to-location map:", lines, i); err == nil {
			humidityToLocationRange = _range
			i = newI
		} else {
			println("Unknown line: " + line)
		}
		i++
	}

	parsedSeeds := []int{}

	for _, seed := range seeds {
		seedInt, err := strconv.Atoi(seed)
		if err != nil {
			panic(err)
		}
		parsedSeeds = append(parsedSeeds, seedInt)
	}

	return parsedSeeds, map[string]RangeElements{
		"seedToSoilRange":            seedToSoilRange,
		"soilToFertilizerRange":      soilToFertilizerRange,
		"fertilizerToWaterRange":     fertilizerToWaterRange,
		"waterToLightRange":          waterToLightRange,
		"lightToTemperatureRange":    lightToTemperatureRange,
		"temperatureToHumidityRange": temperatureToHumidityRange,
		"humidityToLocationRange":    humidityToLocationRange,
	}

}

func Part1(lines []string) int {
	seeds, ranges := Parse(lines)
	seedToSoilRange := ranges["seedToSoilRange"]
	soilToFertilizerRange := ranges["soilToFertilizerRange"]
	fertilizerToWaterRange := ranges["fertilizerToWaterRange"]
	waterToLightRange := ranges["waterToLightRange"]
	lightToTemperatureRange := ranges["lightToTemperatureRange"]
	temperatureToHumidityRange := ranges["temperatureToHumidityRange"]
	humidityToLocationRange := ranges["humidityToLocationRange"]

	humidities := []int{}
	soil := 0
	for _, seed := range seeds {
		soil = seedToSoilRange.GetDestination(seed)
		soil = soilToFertilizerRange.GetDestination(soil)
		soil = fertilizerToWaterRange.GetDestination(soil)
		soil = waterToLightRange.GetDestination(soil)
		soil = lightToTemperatureRange.GetDestination(soil)
		soil = temperatureToHumidityRange.GetDestination(soil)
		soil = humidityToLocationRange.GetDestination(soil)
		humidities = append(humidities, soil)
	}

	minHumidity := math.MaxInt64
	for _, humidity := range humidities {
		if humidity < minHumidity {
			minHumidity = humidity
		}
	}

	return minHumidity
}

func Part2(lines []string) int {
	return 0
}

func Run() {
	lines := utils.LoadText("./data/day5.txt")

	println()
	println("##############################")
	println("#            Day 5           #")
	println("##############################")
	println()
	println("Part 1")
	println(Part1(lines))

	println("Part 2")
	println(Part2(lines))
}
