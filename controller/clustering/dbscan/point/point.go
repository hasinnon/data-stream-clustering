// Package cluster implements DBScan clustering on (lat, lon) using K-D Tree
package point

// Point is longitue, latittude
type Point []float64

// PointList is a slice of Points
type PointList []Point

// Cluster is a result of DBScan work
type Cluster struct {
	C      int
	Points []int
}

// sqDist returns squared (w/o sqrt & normalization) distance between two points
func (a *Point) sqDist(b *Point) float64 {
	return DistanceNormal(a, b)
}

// LessEq - a <= b
func (a *Point) LessEq(b *Point) bool {
	for i, v := range *a {
		if v > (*b)[i] {
			return false
		}
	}
	return true
}

// GreaterEq - a >= b
func (a *Point) GreaterEq(b *Point) bool {
	for i, v := range *a {
		if v < (*b)[i] {
			return false
		}
	}
	return true
}

// Equal - a == b
func (a *Point) Equal(b *Point) bool {
	for i, v := range *a {
		if v != (*b)[i] {
			return false
		}
	}
	return true
}

// CentroidAndBounds calculates center and cluster bounds
func (c *Cluster) CentroidAndBounds(points PointList) (center, min, max Point) {
	if len(c.Points) == 0 {
		panic("empty cluster")
	}

	min = points[0]
	max = points[0]

	for _, i := range c.Points {
		pt := points[i]

		for j := range pt {
			center[j] += pt[j]

			if pt[j] < min[j] {
				min[j] = pt[j]
			}
			if pt[j] > max[j] {
				max[j] = pt[j]
			}
		}
	}

	for j := range center {
		center[j] /= float64(len(c.Points))
	}

	return
}

// Inside checks if (innerMin, innerMax) rectangle is inside (outerMin, outMax) rectangle
func Inside(innerMin, innerMax, outerMin, outerMax *Point) bool {
	return innerMin.GreaterEq(outerMin) && innerMax.LessEq(outerMax)
}
