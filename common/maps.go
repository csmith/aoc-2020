package common

// Tile is a single space within a 2D map, as read by ReadFileAsMap.
type Tile rune

// Map is a 2d map.
type Map [][]Tile

// ReadFileAsMap reads all lines from the given path and returns them as a two-dimensional map.
// If an error occurs, the function will panic.
func ReadFileAsMap(path string) Map {
	var res [][]Tile
	lines := ReadFileAsStrings(path)
	for i := range lines {
		res = append(res, []Tile(lines[i]))
	}
	return res
}

// TileAt returns the tile at the given co-ordinates in the map, wrapping around on both axes.
func (m Map) TileAt(row, col int) Tile {
	line := m[row % len(m)]
	return line[col % len(line)]
}
