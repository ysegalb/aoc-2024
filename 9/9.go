package main

import (
	"aoc2024/utils"
	"fmt"
	"strconv"
)

type Drive struct {
	Content []string
}

func (d *Drive) SwapSectorContent(i, j int) {
	d.Content[i], d.Content[j] = d.Content[j], d.Content[i]
}

func (d *Drive) GetFirstFreeSpaceIndex() int {
	for i := 0; i < len(d.Content); i++ {
		if d.Content[i] == "." {
			return i
		}
	}
	return -1
}

func (d *Drive) GetLastDataIndex() int {
	for i := len(d.Content) - 1; i > 0; i-- {
		if d.Content[i] != "." {
			return i
		}
	}
	return -1
}

func (d *Drive) GetLastDataIndexWithChangeDetection() int {
	data := false
	for i := len(d.Content) - 1; i > 0; i-- {
		if d.Content[i] != "." {
			data = true
		}
		if d.Content[i-1] != d.Content[i] && data {
			return i
		}
	}
	return -1
}

func (d *Drive) Compact() {
	var lastFreeSpaceIndex, lastDataIndex int
	for {
		freeSpaceIndex := d.GetFirstFreeSpaceIndex()
		dataIndex := d.GetLastDataIndex()

		if lastFreeSpaceIndex == freeSpaceIndex || lastDataIndex == dataIndex || freeSpaceIndex > dataIndex {
			break
		}
		d.SwapSectorContent(freeSpaceIndex, dataIndex)
		lastFreeSpaceIndex = freeSpaceIndex
		lastDataIndex = dataIndex
	}
}

type SectorizedDrive struct {
	Index   int
	Size    int
	Content string
}

func fromDrive(d *Drive) []SectorizedDrive {
	sd := make([]SectorizedDrive, 0)
	prevContent := d.Content[0]
	size := 0
	idx := 0
	for i, content := range d.Content {
		if content == prevContent {
			size++
		} else {
			sd = append(sd, SectorizedDrive{
				Index:   idx,
				Size:    size,
				Content: prevContent,
			})
			prevContent = content
			idx = i
			size = 1
		}
	}
	sd = append(sd, SectorizedDrive{
		Index:   idx,
		Size:    size,
		Content: prevContent,
	})
	return sd
}

func getEmptySectors(sd []SectorizedDrive) []SectorizedDrive {
	emptySectors := make([]SectorizedDrive, 0)
	for _, sector := range sd {
		if sector.Content == "." {
			emptySectors = append(emptySectors, sector)
		}
	}
	return emptySectors
}

func getEmptySectorsIndexFrom(sd []SectorizedDrive, index int) []SectorizedDrive {
	emptySectors := make([]SectorizedDrive, 0)
	for _, sector := range sd {
		if sector.Content == "." && sector.Index < index {
			emptySectors = append(emptySectors, sector)
		}
	}
	return emptySectors
}

func getDataSectors(sd []SectorizedDrive) []SectorizedDrive {
	dataSectors := make([]SectorizedDrive, 0)
	for _, sector := range sd {
		if sector.Content != "." {
			dataSectors = append(dataSectors, sector)
		}
	}
	return dataSectors
}

func swapDataBlock(d *Drive, from, to int, size int) {
	for i := 0; i < size; i++ {
		d.SwapSectorContent(to+i, from+i)
	}
}

func (d *Drive) Defrag() {
	dataSectors := getDataSectors(fromDrive(d))
	emptySectors := getEmptySectors(fromDrive(d))

	i := len(dataSectors) - 1
	for {
		if len(dataSectors) <= 0 {
			break
		}
		for j := 0; j < len(emptySectors); j++ {
			if emptySectors[j].Index > dataSectors[i].Index {
				break
			}
			if dataSectors[i].Size <= emptySectors[j].Size {
				swapDataBlock(d, dataSectors[i].Index, emptySectors[j].Index, dataSectors[i].Size)
				break
			}
		}
		dataSectors = dataSectors[:i]
		if len(dataSectors) <= 0 {
			break
		}
		emptySectors = getEmptySectorsIndexFrom(fromDrive(d), dataSectors[len(dataSectors)-1].Index)
		i--
	}
}

func (d *Drive) calculateChecksum() int {
	totalChecksum := 0
	for i, content := range d.Content {
		if content != "." {
			blockContent, _ := strconv.Atoi(content)
			totalChecksum += i * blockContent
		}
	}
	return totalChecksum
}

func GetCompactedChecksum(contentFile string, wholeSector bool) int {
	lines := utils.ReadLines(contentFile)
	drive := processLine(lines[0])
	if wholeSector {
		drive.Defrag()
	} else {
		drive.Compact()
	}
	return drive.calculateChecksum()
}

func processLine(line string) Drive {
	drive := Drive{Content: make([]string, 0)}
	id := 0
	for index, ch := range line {
		limit, _ := strconv.Atoi(string(ch))
		var content string
		for i := 0; i < limit; i++ {
			if index%2 == 0 {
				content = fmt.Sprintf("%d", id)
			} else {
				content = "."
			}
			drive.Content = append(drive.Content, content)
		}
		if "." != content && string(ch) != "0" {
			id++
		}
	}
	return drive
}
