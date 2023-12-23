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

func (res RangeElements) String() string {
	str := ""
	for _, re := range res {
		str += re.String() + "\n"
	}
	return str
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

func SplitSeeds(seeds []int) [][]int {
	seedGroups := [][]int{}
	for i := 0; i < len(seeds); i += 2 {
		seedGroups = append(seedGroups, seeds[i:i+2])
	}
	return seedGroups
}

func CalculateMinLocation(seeds []int, ranges map[string]RangeElements) int {
	seedToSoilRange := ranges["seedToSoilRange"]
	soilToFertilizerRange := ranges["soilToFertilizerRange"]
	fertilizerToWaterRange := ranges["fertilizerToWaterRange"]
	waterToLightRange := ranges["waterToLightRange"]
	lightToTemperatureRange := ranges["lightToTemperatureRange"]
	temperatureToHumidityRange := ranges["temperatureToHumidityRange"]
	humidityToLocationRange := ranges["humidityToLocationRange"]

	minLocation := math.MaxInt64
	tmp := 0

	for _, seed := range seeds {
		tmp = seedToSoilRange.GetDestination(seed)
		tmp = soilToFertilizerRange.GetDestination(tmp)
		tmp = fertilizerToWaterRange.GetDestination(tmp)
		tmp = waterToLightRange.GetDestination(tmp)
		tmp = lightToTemperatureRange.GetDestination(tmp)
		tmp = temperatureToHumidityRange.GetDestination(tmp)
		location := humidityToLocationRange.GetDestination(tmp)
		if location < minLocation {
			minLocation = location
		}
	}

	return minLocation
}

func Part1(lines []string) int {
	seeds, ranges := Parse(lines)
	return CalculateMinLocation(seeds, ranges)
}

func Part2(lines []string) int {
	seeds, ranges := Parse(lines)

	seedPairs := SplitSeeds(seeds)
	newSeeds := []int{}
	for _, seedPair := range seedPairs {
		seedStart := seedPair[0]
		seedRange := seedPair[1]
		for i := 0; i < seedRange; i++ {
			newSeeds = append(newSeeds, seedStart+i)
		}
	}

	return CalculateMinLocation(newSeeds, ranges)
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
