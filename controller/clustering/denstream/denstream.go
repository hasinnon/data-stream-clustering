package denstream

import (
	"math"
	dbscan "my/ar/399/datastream/controller/clustering/dbscan"
	"my/ar/399/datastream/controller/clustering/dbscan/point"
)

// Point is a sample , a record of input
type Point map[string]interface{}

// Assign tow point
func (p1 Point) Assign(p2 Point) {
	for key, val := range p2 {
		var v = val
		p1[key] = v
	}
}

func (p1 Point) ToFloat() *[]float64 {
	rp := make([]float64, len(p1))
	i := 0
	for k, v := range p1 {
		switch p1[k].(type) {
		case float64:
			rp[i] = v.(float64)
		case int:
			rp[i] = float64(v.(int))
		}
		i++

	}
	return &rp
}

// Pow2 point
func (p1 Point) Pow2() *Point {
	rp := make(Point)
	// rp := p1
	for k, v := range p1 {
		switch v.(type) {
		case float64:
			rp[k] = v.(float64) * v.(float64)
		case int:
			rp[k] = v.(int) * v.(int)
		}

	}
	return &rp
}

// Add tow point
func (p1 Point) Add(newP Point) *Point {
	rp := make(Point)
	//rp.Assign(p1)
	// rp := p1
	for k, v := range p1 {
		switch v.(type) {
		case float64:
			switch newP[k].(type) {
			case float64:
				rp[k] = v.(float64) + newP[k].(float64)
			case int:
				rp[k] = v.(float64) + float64(newP[k].(int))
			case string:
				//
			}
		case int:
			switch newP[k].(type) {
			case float64:
				rp[k] = float64(v.(int)) + newP[k].(float64)
			case int:
				rp[k] = v.(int) + newP[k].(int)
			case string:
				//
			}
		case string:
			//
		}
	}
	return &rp
}

// Sub tow point
func (p1 Point) Sub(newP Point) *Point {
	rp := make(Point)
	//rp.Assign(p1)
	// rp := p1
	for k, v := range p1 {
		switch v.(type) {
		case float64:
			switch newP[k].(type) {
			case float64:
				rp[k] = v.(float64) - newP[k].(float64)
			case int:
				rp[k] = v.(float64) - float64(newP[k].(int))
			case string:
				//
			}
		case int:
			switch newP[k].(type) {
			case float64:
				rp[k] = float64(v.(int)) - newP[k].(float64)
			case int:
				rp[k] = v.(int) - newP[k].(int)
			case string:
				//
			}
		case string:
			//
		}
	}
	return &rp
}

// Mult tow point
func (p1 Point) Mult(z float64) *Point {
	rp := make(Point)
	//rp.Assign(p1)
	// rp := p1
	for k, v := range p1 {
		switch v.(type) {
		case float64:
			rp[k] = v.(float64) * z
		case int:
			rp[k] = z * float64(v.(int))
		case string:
			//
		}
	}
	return &rp
}

// MultP tow point
func (p1 Point) MultP(newP Point) *Point {
	rp := make(Point)
	//rp.Assign(p1)
	// rp := p1
	for k, v := range p1 {
		switch v.(type) {
		case float64:
			switch newP[k].(type) {
			case float64:
				rp[k] = v.(float64) * newP[k].(float64)
			case int:
				rp[k] = v.(float64) * float64(newP[k].(int))
			case string:
				//
			}
		case int:
			switch newP[k].(type) {
			case float64:
				rp[k] = float64(v.(int)) * newP[k].(float64)
			case int:
				rp[k] = v.(int) * newP[k].(int)
			case string:
				//
			}
		case string:
			//
		}
	}
	return &rp
}

// DenPoint stores DenStream Point
type DenPoint struct {
	Value     Point
	timestamp int32
	covered   bool
}

// Assign tow DenStream point
func (p1 *DenPoint) Assign(p2 DenPoint) {
	p1.Value = make(Point)
	p1.timestamp = p2.timestamp
	p1.covered = p2.covered
	p1.Value.Assign(p2.Value)
}

func NewDenPoint(arg Point) *DenPoint {
	var p DenPoint
	p.Value = make(Point)
	p.Value.Assign(arg)
	return &p
}

func newDenPointmulti(arg ...interface{}) *DenPoint {
	if len(arg) == 0 {
		return &DenPoint{}
	} else if len(arg) == 1 {
		return &DenPoint{
			Value: arg[0].(Point),
		}
	}
	return &DenPoint{
		Value:     arg[0].(map[string]interface{}),
		timestamp: arg[1].(int32),
	}

}

// Distance return distance of tow point
func Distance(a DenPoint, b DenPoint) float64 {
	var sumSquaredDiffs float64 = 0
	// s := make(Point)
	// s.Assign(*a.Value.Sub(b.Value))
	// p := make(Point)
	// p.Assign(*s.Pow2())
	s := a.Value.Sub(b.Value)
	p := s.Pow2()

	for _, value := range *p {
		switch value.(type) {
		case float32:
			sumSquaredDiffs += float64(value.(float32))
		case float64:
			sumSquaredDiffs += float64(value.(float64))
		case int:
			sumSquaredDiffs += float64(value.(int))
		case int32:
			sumSquaredDiffs += float64(value.(int32))
		case int64:
			sumSquaredDiffs += float64(value.(int64))
		}
	}
	return math.Sqrt(sumSquaredDiffs)
}

func (p1 *DenPoint) setTimestamp(timestamp int32) {
	p1.timestamp = timestamp
}

// MicroCluster is struct of micro cluster of DenStream Algorithm
type MicroCluster struct {
	LS                Point //????float64
	SS                Point //??????float64
	N                 int
	LastEditT         int32 //??????? = -1;
	Lambda            float64
	CurrentTimestamp  int32
	CreationTimestamp int32 //????????? = -1;
	Weight            float64
	Radius            float64
}

func newMicroCluster(center Point, dimensions int, creationTimestamp int32, lambda float64, currentTimestamp int32) *MicroCluster {
	mCluster := MicroCluster{CreationTimestamp: creationTimestamp}
	mCluster.LastEditT = creationTimestamp
	mCluster.Lambda = lambda
	mCluster.CurrentTimestamp = currentTimestamp
	mCluster.N = 1
	mCluster.LS = make(Point)
	mCluster.LS.Assign(center)
	mCluster.SS = make(Point, dimensions)
	mCluster.SS = *center.Pow2()
	mCluster.Weight = 0
	mCluster.Radius = 0
	return &mCluster
}

func (mCluster *MicroCluster) insert(point DenPoint, timstamp int32) {
	mCluster.N++
	mCluster.LastEditT = timstamp

	// Update Radius
	mCluster.Radius = (float64(mCluster.N)*mCluster.Radius+Distance(point, DenPoint{Value: mCluster.LS}))/float64(mCluster.N) + 1
	oldWeight := mCluster.Weight
	mCluster.Weight = mCluster.Weight*math.Pow(2, -1*mCluster.Lambda) + 1
	// oldLS11 := make(Point)
	// oldLS1.Assign(mCluster.LS)

	mCluster.LS = *mCluster.LS.Add(*point.Value.Sub(mCluster.LS).Mult(1.00 / float64(mCluster.N)))
	// newLS1 := make(Point)
	// newLS.Assign(mCluster.LS)
	// newLS1 := make(Point)
	// oldSS1 := make(Point)
	// oldSS.Assign(mCluster.SS)
	mCluster.SS = *mCluster.SS.Mult(((mCluster.Weight - float64(timstamp)) / oldWeight)).Add(*(point.Value.Sub(mCluster.LS)).MultP(*(point.Value.Sub(mCluster.SS))).Mult(mCluster.Weight))
	//mCluster.SS = point.Value.Add()
}

// GetCenter return center a micro cluster
func (mCluster MicroCluster) GetCenter(arg ...int32) *DenPoint {
	if len(arg) > 0 {
		return mCluster.getCenter(arg[0])
	}
	return mCluster.getCenter(mCluster.CurrentTimestamp)
}

func (mCluster MicroCluster) getCenter(timestamp int32) *DenPoint {
	//var dt float64 = float64(timestamp - mCluster.LastEditT)
	//var w float64 = mCluster.GetWeight(timestamp)
	var result DenPoint
	result.Value = make(Point)
	// result.Value.Assign(mCluster.LS)
	result.Value = mCluster.LS

	return &result
}

// GetWeight return weight a micro cluster
func (mCluster MicroCluster) GetWeight(arg ...int32) float64 {
	if len(arg) > 0 {
		return mCluster.getWeight(arg[0])
	}
	return mCluster.getWeight(mCluster.CurrentTimestamp)
}

func (mCluster MicroCluster) getWeight(timestamp int32) float64 {
	var dt float64 = float64(timestamp - mCluster.LastEditT)
	return float64(mCluster.N) * math.Pow(2, -1*mCluster.Lambda*dt)
}

func (mCluster MicroCluster) calcCF3() *[]float64 {
	var cf2 = make([]float64, len(mCluster.SS))
	i := 0 //??????????
	for _, v := range mCluster.SS {
		switch v.(type) {
		case float64:
			cf2[i] = v.(float64)
		case int:
			cf2[i] = float64(v.(int))
		case string:
		}
		i++
	}
	return &cf2
}

func (mCluster MicroCluster) getCreationTime() int32 {
	return mCluster.CreationTimestamp
}

// DenStream is struct of DenStream Algorithm
type DenStream struct {
	PmicroCluster    []MicroCluster
	OmicroCluster    []MicroCluster
	Timestamp        int32 // =0
	Mepsilon         float64
	DBepsilon        float64
	Beta             float64
	Mu               float64
	CurrentTimestamp int32
	Lambda           float64
	Tp               int64
	Initialized      bool
	MinPoints        int
	InitBuffer       []DenPoint
	InitN            int
}

// Start set attribute of DenStream
func (denSRM *DenStream) Start(eps, dbeps, lambda float64) {
	denSRM.setLearningPara(eps, dbeps, lambda)
}

func (denSRM *DenStream) StartDenStream(p Point) {
	p1 := NewDenPoint(p)

	denSRM.DenStreamAlgorithm(*p1)
}

func (denSRM *DenStream) setLearningPara(eps, dbeps, lambda float64) {
	denSRM.Mepsilon = eps //f;?????? =3
	denSRM.DBepsilon = dbeps
	denSRM.Beta = 0.4      //f;
	denSRM.Mu = 3          //f;
	denSRM.Lambda = lambda //f;
	denSRM.InitN = 30
	denSRM.MinPoints = 6

	denSRM.Initialized = false
	denSRM.Tp = int64(math.Round(1 / denSRM.Lambda * math.Log((denSRM.Beta*denSRM.Mu)/(denSRM.Beta*denSRM.Mu-1))))
}

func (denSRM *DenStream) initialDBScan() {
	for i := 0; i < len(denSRM.InitBuffer); i++ {
		point := &denSRM.InitBuffer[i]
		// point := denSRM.InitBuffer[i]
		if !point.covered {
			point.covered = true
			var neighbourhood []int = denSRM.getNeighbourhoodIDs(*point, denSRM.InitBuffer)
			if len(neighbourhood) >= denSRM.MinPoints {
				mc := newMicroCluster(point.Value, 2, denSRM.Timestamp, denSRM.Lambda, denSRM.CurrentTimestamp)
				denSRM.expandCluster(*mc, &denSRM.InitBuffer, neighbourhood)
				denSRM.PmicroCluster = append(denSRM.PmicroCluster, *mc)
			} else {
				point.covered = false
			}
		}

	}
}

func (denSRM *DenStream) initWithoutDBScan() {
	sel := make([]int, len(denSRM.InitBuffer))
	for i := 0; i < len(sel); i++ {
		sel[i] = i
	}
	sels := make([][]int, 1)
	combinations(sel, denSRM.MinPoints, 0, make([]int, denSRM.MinPoints), &sels)
	minR := 0.0
	fmc := newMicroCluster(denSRM.InitBuffer[sels[1][0]].Value, 2, denSRM.Timestamp, denSRM.Lambda, denSRM.CurrentTimestamp)
	for j := 1; j < len(sels[1]); j++ {
		fmc.insert(denSRM.InitBuffer[sels[1][j]], int32(j+1))
	}
	minR = fmc.Radius

	for i := 1; i < len(sels); i++ {
		mc := newMicroCluster(denSRM.InitBuffer[sels[i][0]].Value, 2, denSRM.Timestamp, denSRM.Lambda, denSRM.CurrentTimestamp)
		for j := 1; j < len(sels[i]); j++ {
			mc.insert(denSRM.InitBuffer[sels[i][j]], 2)
			if mc.Radius > minR {
				continue
			}
		}
		if mc.Radius < minR {
			minR = mc.Radius
		}
	}

	denSRM.Mepsilon = minR
}

func (denSRM *DenStream) getNeighbourhoodIDs(point DenPoint, points []DenPoint) []int {
	var neighbourIDs []int
	for i := 0; i < len(points); i++ {
		var testPoint DenPoint
		testPoint.Assign(points[i])
		// testPoint := points[i]
		if !testPoint.covered {
			var dist float64 = Distance(testPoint, point)
			if dist < denSRM.Mepsilon {
				neighbourIDs = append(neighbourIDs, i)
			}
		}
	}
	return neighbourIDs
}

func (denSRM *DenStream) expandCluster(mc MicroCluster, points *[]DenPoint, neighbourhood []int) {
	for i := 0; i < len(neighbourhood); i++ {
		testPoint := &(*points)[neighbourhood[i]]
		if !(*testPoint).covered {
			(*testPoint).covered = true
			mc.insert(*testPoint, denSRM.Timestamp)
			var neighbourhood2 []int = denSRM.getNeighbourhoodIDs((*testPoint), denSRM.InitBuffer)
			if len(neighbourhood) >= denSRM.MinPoints {
				denSRM.expandCluster(mc, points, neighbourhood2)
			}
		}
	}
}

// DenStreamAlgorithm run Algorithm by get point
func (denSRM *DenStream) DenStreamAlgorithm(point DenPoint) {
	denSRM.Timestamp++
	point.setTimestamp(denSRM.Timestamp)
	if !denSRM.Initialized {
		denSRM.InitBuffer = append(denSRM.InitBuffer, point)
		if len(denSRM.InitBuffer) >= denSRM.InitN {
			if denSRM.Mepsilon <= 0 {
				denSRM.initWithoutDBScan()
			} else {
				denSRM.initialDBScan()
			}
			denSRM.Initialized = true
		}

	} else {
		denSRM.Merging(point, denSRM.Timestamp)
	}

	if int64(denSRM.Timestamp)%denSRM.Tp == 0 {
		denSRM.Arrange()
	}
}

// Arrange of PmicroCluster and OmicroCluster
func (denSRM *DenStream) Arrange() {

	var removalListP []int
	var removalListO []int
	for i := 0; i < len(denSRM.PmicroCluster); i++ {
		if denSRM.PmicroCluster[i].GetWeight() < denSRM.Beta*denSRM.Mu {
			removalListP = append(removalListP, i)
		}
	}

	for i := len(removalListP) - 1; i >= 0; i-- {
		// Remove
		lenp := len(denSRM.PmicroCluster)
		denSRM.PmicroCluster[removalListP[i]] = denSRM.PmicroCluster[lenp-1]
		denSRM.PmicroCluster = denSRM.PmicroCluster[:lenp-1]
	}

	for i := 0; i < len(denSRM.OmicroCluster); i++ {
		var t0 int32 = denSRM.OmicroCluster[i].getCreationTime()
		var xsi1 float64 = math.Pow(2, (-denSRM.Lambda*(float64(int64(denSRM.Timestamp-t0)+denSRM.Tp)))) - 1
		var xsi2 float64 = math.Pow(2, (denSRM.Lambda*float64(denSRM.Tp*-1))) - 1
		var xsi float64 = xsi1 / xsi2

		if denSRM.OmicroCluster[i].GetWeight() < xsi {
			removalListO = append(removalListO, i)
		}
	}

	for i := len(removalListO) - 1; i >= 0; i-- {
		// Remove
		leno := len(denSRM.OmicroCluster)
		denSRM.OmicroCluster[removalListO[i]] = denSRM.OmicroCluster[leno-1]
		denSRM.OmicroCluster = denSRM.OmicroCluster[:leno-1]
	}
}

// Merging is merging a point to a micro cluster
func (denSRM *DenStream) Merging(point DenPoint, timestamp int32) {
	denSRM.CurrentTimestamp = timestamp
	var merged bool = false
	if len(denSRM.PmicroCluster) != 0 {
		x := denSRM.nearestCluster(point, "P")
		//???var xCopy MicroCluster = *x.copy()
		xCopy := denSRM.PmicroCluster[x]
		xCopy.insert(point, timestamp)

		// if xCopy.getRadius(timestamp) <= denSRM.Mepsilon {
		if xCopy.Radius <= denSRM.Mepsilon {
			denSRM.PmicroCluster[x].insert(point, timestamp)
			merged = true
		}
	}

	if !merged && (len(denSRM.OmicroCluster) != 0) {
		x := denSRM.nearestCluster(point, "O")
		//var xCopy MicroCluster = *x.copy()
		xCopy := denSRM.OmicroCluster[x]
		xOrginal := denSRM.OmicroCluster[x]
		xCopy.insert(point, timestamp)

		if xCopy.Radius <= denSRM.Mepsilon {
			denSRM.OmicroCluster[x].insert(point, timestamp)
			merged = true
			if denSRM.OmicroCluster[x].GetWeight() > denSRM.Beta*denSRM.Mu {
				// Remove
				leno := len(denSRM.OmicroCluster)
				denSRM.OmicroCluster[x] = denSRM.OmicroCluster[leno-1]
				denSRM.OmicroCluster = denSRM.OmicroCluster[:leno-1]
				// Append
				denSRM.PmicroCluster = append(denSRM.PmicroCluster, xOrginal)
			}
		}
	}

	if !merged {
		denSRM.OmicroCluster = append(denSRM.OmicroCluster, *newMicroCluster(point.Value, 2, timestamp, denSRM.Lambda, denSRM.CurrentTimestamp))
	}
}

func (denSRM *DenStream) nearestCluster(point DenPoint, clu string) int {
	var min int
	var check bool = true
	var minDist float64 = 0
	var cl []MicroCluster

	if clu == "P" || clu == "p" {
		cl = denSRM.PmicroCluster
	} else if clu == "O" || clu == "o" {
		cl = denSRM.OmicroCluster
	}

	for c := 0; c < len(cl); c++ {
		var cluster MicroCluster = cl[c]
		if check {
			min = c
			check = false
		}
		var dist float64
		center := *cluster.GetCenter()
		dist = Distance(point, center)
		dist -= cluster.Radius

		if c == 0 {
			minDist = dist
		} else if dist < minDist {
			minDist = dist
			min = c

		}
	}

	return min
}

// MicroClusterPoints return PmicroCluster of DenStream
func (denSRM *DenStream) MicroClusterPoints() point.PointList {
	var result point.PointList
	var point point.Point
	if len(denSRM.PmicroCluster) != 0 {
		for _, v := range denSRM.PmicroCluster {
			point = *v.GetCenter().Value.ToFloat()
			/*for k, val := range *v.GetCenter().Value.toFloat() {
				point[k] = val

			}*/
			result = append(result, point)
		}
	}
	return result
}

// Offline is offline part of DenStream Algorithm
func (denSRM *DenStream) Offline() *[][][]float64 {
	Points := denSRM.MicroClusterPoints()
	eps := denSRM.DBepsilon
	if eps <= 0 || eps < denSRM.Mepsilon*1.5 {
		eps = denSRM.Mepsilon * 3.5
	}
	dbclusters, _ := dbscan.DBScan(Points, eps, 3)
	result := make([][][]float64, len(dbclusters))

	for i := range result {
		result[i] = make([][]float64, len(dbclusters[i].Points))
		for j := range result[i] {
			z := len(Points[dbclusters[i].Points[j]])
			result[i][j] = make([]float64, z)
			for k := 0; k < z; k++ {
				result[i][j][k] = Points[dbclusters[i].Points[j]][k]
			}
		}
	}
	return &result
}

type USresult struct {
	Clusters [][][]float64  `json:"clusters"`
	HdrName  map[int]string `json:"hdrname"`
	Eps      float64        `json:"eps"`
	NumData  int            `json:"countdata"`
	NumMicro int            `json:"countmicro"`
}

func (re USresult) FindClustersIDs(point []float64) [][2]int {
	var neighbourIDs [][2]int
	for c := 0; c < len(re.Clusters); c++ {
		for i := 0; i < len(re.Clusters[c]); i++ {
			testPoint := re.Clusters[c][i]

			sum := 0.0
			for index, value := range testPoint {
				temp := value - point[index]
				sum += temp * temp
			}
			if sum < re.Eps*re.Eps {
				var add = [2]int{c, i}
				neighbourIDs = append(neighbourIDs, add)
			}

		}
	}
	return neighbourIDs
}

func per(s []int, k int, n int, result [][]int) {
	if k == n {
		result = append(result, s)
	} else {
		for i := k; i <= n; i++ {
			c := s[i]
			s[i] = s[k]
			s[k] = c
			per(s, k+1, n, result)
		}
	}
}
func combinations(arr []int, alen int, startPosition int, result []int, output *[][]int) {
	if alen == 0 {
		r := make([]int, len(result))
		for j := 0; j < len(r); j++ {
			r[j] = result[j]
		}
		*output = append(*output, r)
		return
	}
	for i := startPosition; i <= len(arr)-alen; i++ {
		result[len(result)-alen] = arr[i]
		combinations(arr, alen-1, i+1, result, output)
	}
}

type PointVars struct {
	Vars map[string]bool `json:"vars"`
}

// func (mCluster MicroCluster) copy() *MicroCluster {
// 	copy := newMicroCluster(mCluster.LS, len(mCluster.LS), mCluster.getCreationTime(), mCluster.Lambda, mCluster.CurrentTimestamp)
// 	copy.N = mCluster.N
// 	copy.SS = mCluster.SS
// 	copy.LS = mCluster.LS
// 	copy.LastEditT = mCluster.LastEditT
// 	return copy
// }

// func remove111(a []MicroCluster, index int) []MicroCluster {
// 	x := a
// 	x[index] = x[len(x)-1]
// 	return x[:len(x)-1]
// }

/*
func (mCluster MicroCluster) GetRadius1(arg ...int32) float64 {
	if len(arg) > 0 {
		return mCluster.getRadius1(arg[0])
	}
	return mCluster.getRadius1(mCluster.CurrentTimestamp)
}

func (mCluster MicroCluster) getRadius1(timestamp int32) float64 {
	/*var dt int32 = timestamp - mCluster.LastEditT
	var cf1 []float64 = *mCluster.calcCF1(dt)
	var cf2 []float64 = *mCluster.calcCF2(dt)
	var w float64 = mCluster.getWeight(timestamp)
	var max float64 = 0
	for i := 0; i < len(mCluster.SS); i++ {
		var x1 float64 = cf2[i] / w
		var temp float64 = cf1[i] / w
		var x2 float64 = temp * temp
		//// float64 x3 = Mathf.Pow (cf1 [i] / w, 2);
		var diffSqrt float64 = math.Sqrt(x1 - x2)
		if diffSqrt > max {
			max = diffSqrt
		}
	}

	var cf3 []float64 = *mCluster.calcCF3()
	var max float64 = 0
	for i := 0; i < len(cf3); i++ {
		if cf3[i] > max {
			max = cf3[i]
		}
	}

	return max
}

func (mCluster MicroCluster) calcCF1(dt int32) *[]float64 {
	var cf1 = make([]float64, len(mCluster.LS))
	i := 0 //??????????
	for _, v := range mCluster.LS {
		switch v.(type) {
		case float64:
			cf1[i] = math.Pow(2, float64(-1*dt)*mCluster.Lambda) * v.(float64)
		case int:
			cf1[i] = math.Pow(2, float64(-1*dt)*mCluster.Lambda) * float64(v.(int))
		case string:
		}
		i++
	}
	return &cf1
}

func (mCluster MicroCluster) calcCF2(dt int32) *[]float64 {
	var cf2 = make([]float64, len(mCluster.SS))
	i := 0 //??????????
	for _, v := range mCluster.SS {
		switch v.(type) {
		case float64:
			cf2[i] = math.Pow(2, float64(-1*dt)*mCluster.Lambda) * v.(float64)
		case int:
			cf2[i] = math.Pow(2, float64(-1*dt)*mCluster.Lambda) * float64(v.(int))
		case string:
		}
		i++
	}
	return &cf2
}
*/

/*
// Assign tow point
func (p1 Point) Assign(p2 Point) {
	for key, val := range p2 {
		var v = val
		p1[key] = v
	}
}

func (p1 Point) toFloat() *[]float64 {
	rp := make([]float64, len(p1))
	i := 0
	for k, v := range p1 {
		switch p1[k].(type) {
		case float64:
			rp[i] = v.(float64)
		case int:
			rp[i] = float64(v.(int))
		}
		i++

	}
	return &rp
}

// Pow2 point
func (p1 Point) Pow2() *Point {
	rp := make(Point)
	rp.Assign(p1)
	// rp := p1
	for k, v := range p1 {
		switch rp[k].(type) {
		case float64:
			rp[k] = v.(float64) * v.(float64)
		case int:
			rp[k] = v.(int) * v.(int)
		}

	}
	return &rp
}

// Add tow point
func (p1 Point) Add(newP Point) *Point {
	rp := make(Point)
	rp.Assign(p1)
	// rp := p1
	for k, v := range p1 {
		switch rp[k].(type) {
		case float64:
			switch newP[k].(type) {
			case float64:
				rp[k] = v.(float64) + newP[k].(float64)
			case int:
				rp[k] = v.(float64) + float64(newP[k].(int))
			case string:
				//
			}
		case int:
			switch newP[k].(type) {
			case float64:
				rp[k] = float64(v.(int)) + newP[k].(float64)
			case int:
				rp[k] = v.(int) + newP[k].(int)
			case string:
				//
			}
		case string:
			//
		}
	}
	return &rp
}

// Sub tow point
func (p1 Point) Sub(newP Point) *Point {
	rp := make(Point)
	rp.Assign(p1)
	// rp := p1
	for k, v := range p1 {
		switch rp[k].(type) {
		case float64:
			switch newP[k].(type) {
			case float64:
				rp[k] = v.(float64) - newP[k].(float64)
			case int:
				rp[k] = v.(float64) - float64(newP[k].(int))
			case string:
				//
			}
		case int:
			switch newP[k].(type) {
			case float64:
				rp[k] = float64(v.(int)) - newP[k].(float64)
			case int:
				rp[k] = v.(int) - newP[k].(int)
			case string:
				//
			}
		case string:
			//
		}
	}
	return &rp
}

// Mult tow point
func (p1 Point) Mult(z float64) *Point {
	rp := make(Point)
	rp.Assign(p1)
	// rp := p1
	for k, v := range p1 {
		switch rp[k].(type) {
		case float64:
			rp[k] = v.(float64) * z
		case int:
			rp[k] = z * float64(v.(int))
		case string:
			//
		}
	}
	return &rp
}

// MultP tow point
func (p1 Point) MultP(newP Point) *Point {
	rp := make(Point)
	rp.Assign(p1)
	// rp := p1
	for k, v := range p1 {
		switch rp[k].(type) {
		case float64:
			switch newP[k].(type) {
			case float64:
				rp[k] = v.(float64) * newP[k].(float64)
			case int:
				rp[k] = v.(float64) * float64(newP[k].(int))
			case string:
				//
			}
		case int:
			switch newP[k].(type) {
			case float64:
				rp[k] = float64(v.(int)) * newP[k].(float64)
			case int:
				rp[k] = v.(int) * newP[k].(int)
			case string:
				//
			}
		case string:
			//
		}
	}
	return &rp
}

*/
