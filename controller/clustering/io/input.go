package iopoint

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"log"
	"math/rand"
	"my/ar/399/datastream/controller/clustering/dbscan/point"
	den "my/ar/399/datastream/controller/clustering/denstream"
	"my/ar/399/datastream/datalayer"
	"strconv"
	"strings"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const debug = true

func appendXY(pts *plotter.XYs, po []float64) {
	var pt plotter.XY
	if len(po) == 2 {
		pt.X = po[0]
		pt.Y = po[1]
		*pts = append(*pts, pt)
	} else if len(po) == 1 {
		pt.X = po[0]
		pt.Y = 0
		*pts = append(*pts, pt)
	}
}

func toXYs(p [][][]float64) []plotter.XYs {
	var cls []plotter.XYs

	for i := range p {
		var pts plotter.XYs
		for j := range p[i] {
			var pt plotter.XY
			if len(p[i][j]) == 2 {
				pt.X = p[i][j][0]
				pt.Y = p[i][j][1]
				pts = append(pts, pt)
			} else if len(p[i][j]) == 1 {
				pt.X = p[i][j][0]
				pt.Y = 0
				pts = append(pts, pt)
			}
		}
		cls = append(cls, pts)
	}
	return cls
}

func toXYZs(p [][][]float64, r float64) []plotter.XYZs {
	var cls []plotter.XYZs

	for i := range p {
		var pts plotter.XYZs
		for j := range p[i] {
			var pt plotter.XYZ
			if len(p[i][j]) == 2 {
				pt.X = p[i][j][0]
				pt.Y = p[i][j][1]
				pt.Z = r
				pts = append(pts, pt)
			} else if len(p[i][j]) == 2 {
				pt.X = p[i][j][0]
				pt.Y = 0
				pt.Z = r
				pts = append(pts, pt)
			}
		}
		cls = append(cls, pts)
	}
	return cls
}

func toPlot3D(p point.PointList, r float64) plotter.XYZs {
	var pts plotter.XYZs

	for _, point := range p {

		var pt plotter.XYZ
		if len(point) == 2 {
			pt.X = point[0]
			pt.Y = point[1]
			pt.Z = r
			pts = append(pts, pt)
		} else if len(point) == 1 {
			pt.X = point[0]
			pt.Y = 0
			pt.Z = r
			pts = append(pts, pt)
		}

	}
	return pts
}

func toPlot2D(p point.PointList) plotter.XYs {
	var pts plotter.XYs

	for _, point := range p {

		var pt plotter.XY
		if len(point) == 2 {
			pt.X = point[0]
			pt.Y = point[1]
			pts = append(pts, pt)
		} else if len(point) == 1 {
			pt.X = point[0]
			pt.Y = 0
			pts = append(pts, pt)
		}

	}
	return pts
}

func InfoCSV(cFile io.Reader) (map[int][]string, int, int, error) {
	const numFRow = 10
	var rowNum int = 0
	var colNum int = 0
	table := make(map[int][]string)
	var msgerr error

	// Create a new reader.
	r := csv.NewReader(cFile)
	for row := 1; ; row++ {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			rowNum = row - 1
			break
		}
		if row > numFRow {
			rowNum = row - 1
			break
		}

		if err != nil {
			msgerr = err
		}
		// append to table
		i := 0
		for i = 0; i < len(record); i++ {
			table[row] = append(table[row], record[i])
		}
		if i > colNum {
			colNum = i
		}
	}
	return table, rowNum, colNum, msgerr

}

// Clustering CSV file

var Progressbar map[string]chan int = make(map[string]chan int)

type record []string

type InfoTable struct {
	Header   map[int]string
	Type     []string
	Check    []bool
	HdrCheck bool
	Sname    string
	ServKey  string
	UID      uint
	Meps     float64
	DBeps    float64
}

type errPoints struct {
	plotter.XYs
	plotter.YErrors
	plotter.XErrors
}

func (rec record) toPoint(info InfoTable) *den.Point {
	var result = make(den.Point)
	header := info.Header
	/*
		if len(rec) <= len(header) {
			i := 0
			for i = 0; i < len(rec); i++ {
				val, _ := strconv.ParseFloat(rec[i], 64)
				result[header[i]] = val //????????????
			}
			for i = len(rec); i < len(header); i++ {
				result[header[i]] = ""
			}
		} else {
			i := 0
			for i = 0; i < len(header); i++ {
				val, _ := strconv.ParseFloat(rec[i], 64)
				result[header[i]] = val //????????????
			}
			for i = len(header); i < len(rec); i++ {
				val, _ := strconv.ParseFloat(rec[i], 64)
				result["col___"+fmt.Sprint(i)] = val //????????????
			}
		}*/
	recSize := len(rec)
	for i, h := range header {
		if i < recSize {
			val, parserr := strconv.ParseFloat(rec[i], 64)
			if parserr != nil {
				noSpaceString := strings.ReplaceAll(rec[i], " ", "")
				val1, parserr1 := strconv.ParseFloat(noSpaceString, 64)
				if parserr1 != nil {
					fmt.Println(">>" + rec[i])
				}
				result[h] = val1
			}
			result[h] = val
		}
	}
	return &result
}

// Clusteringfile is function for Clustering a csv file by get file and info
func Clusteringfile(dbhandler datalayer.MyDB, cFile io.Reader, inp InfoTable) (int64, error) {
	const numFRow = 10
	var row int = 0
	var info InfoTable
	// Create a new reader.
	r := csv.NewReader(cFile)
	var pts plotter.XYs
	//var pCluster plotter.XYs
	var lenH int = -1

	// Read header
	if !inp.HdrCheck {
		hrec, herr := r.Read()
		if herr != nil {
			return -10, herr
		}
		temphdr := map[int]string{}
		for c := 0; c < len(inp.Check); c++ {
			if inp.Check[c] {
				temphdr[c] = hrec[c]
			}
		}
		info.Header = temphdr
		lenH = len(temphdr)
	}

	// Read first record
	var frec record
	frec, ferr := r.Read()
	if ferr != nil {
		return -11, ferr
	}
	fpoint := *frec.toPoint(info)
	if lenH == 2 || lenH == 1 {
		appendXY(&pts, *fpoint.ToFloat())
	}
	servkey := inp.ServKey
	select {
	case Progressbar[servkey] <- 0:
	default:
	}
	// Start algoritm
	var serv den.DenStream
	serv.Start(inp.Meps, inp.DBeps, 0.00000001)
	serv.StartDenStream(fpoint)

	for row = 1; ; row++ {
		var record record
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		if err != nil {
			return -12, err
		}
		//
		rPoint := *record.toPoint(info)
		if lenH == 2 || lenH == 1 {
			appendXY(&pts, *rPoint.ToFloat())
		}
		dPoint := den.NewDenPoint(rPoint)

		serv.DenStreamAlgorithm(*dPoint)

		select {
		case Progressbar[servkey] <- row:
		default:
		}
	}
	serv.Arrange()
	clusters := serv.Offline()

	//insert to database
	dbserv := datalayer.Service{
		Sid: sql.NullInt32{
			Int32: -1,
			Valid: false,
		},
		Name: sql.NullString{
			String: inp.Sname,
			Valid:  true,
		},
		Created: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		Credit: sql.NullInt32{
			Int32: -1,
			Valid: true,
		},
		HDcount: sql.NullInt32{
			Int32: int32(lenH),
			Valid: true,
		},
		Deleted: sql.NullBool{
			Bool:  false,
			Valid: true,
		},
		Type: sql.NullBool{
			Bool:  false,
			Valid: true,
		},
		UID: sql.NullInt32{
			Int32: int32(inp.UID),
			Valid: true,
		},
	}
	fclu := den.USresult{
		Clusters: *clusters,
		Eps:      serv.DBepsilon,
		NumData:  row,
		NumMicro: len(serv.PmicroCluster),
		HdrName:  info.Header,
	}
	jsclu, _ := json.Marshal(fclu)
	dbclu := datalayer.ServiceClusters{
		Clusters: sql.NullString{
			String: string(jsclu),
			Valid:  true,
		},
	}
	lastid, err := dbhandler.InsertService(dbserv, dbclu)
	if err != nil {
		return lastid, err
	}
	if lenH == 2 || lenH == 1 {
		//------------------------------------- all point

		// Create a new plot, set its title and
		// axis labels.
		p1 := plot.New()
		// if err != nil {
		// 	panic(err)
		// }
		p1.Title.Text = "BADADE <all records>"
		p1.X.Label.Text = info.Header[0]
		p1.Y.Label.Text = info.Header[1]
		// Draw a grid behind the data
		p1.Add(plotter.NewGrid())

		// Make a scatter plotter and set its style.
		s, err := plotter.NewScatter(pts)
		if err != nil {
			panic(err)
		}
		s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}

		// Add the plotters to the plot, with a legend
		// entry for each
		p1.Add(s)
		p1.Legend.Add("Point", s)

		// Save the plot to a PNG file.
		if err := p1.Save(20*vg.Inch, 20*vg.Inch, "datalayer/plot/a-"+fmt.Sprint(lastid)+".png"); err != nil {
			panic(err)
		}
		if debug {

			//------------------------------------- all clusters
			listpoint := serv.MicroClusterPoints()
			allXYs1 := toPlot2D(listpoint)
			allXYZs1 := toPlot3D(listpoint, serv.Mepsilon)

			// Create a new plot, set its title and

			// axis labels.
			p3 := plot.New()
			// if err != nil {
			// 	panic(err)
			// }
			p3.Title.Text = "BADADE <all micro Clusters>"
			p3.X.Label.Text = info.Header[0]
			p3.Y.Label.Text = info.Header[1]
			// Draw a grid behind the data
			p3.Add(plotter.NewGrid())

			cluColor := color.RGBA{R: uint8(rand.Intn(254)), G: uint8(rand.Intn(254)), B: uint8(rand.Intn(254)), A: 255}

			c, err := plotter.NewScatter(allXYZs1)
			if err != nil {
				panic(err)
			}
			c.Color = cluColor
			n := len(allXYs1)
			perr := make(plotter.Errors, n)
			for i := range perr {
				perr[i].Low = serv.Mepsilon
				perr[i].High = serv.Mepsilon
			}
			data := errPoints{
				XYs:     allXYs1,
				YErrors: plotter.YErrors(perr),
				XErrors: plotter.XErrors(perr),
			}
			xerrs, err := plotter.NewXErrorBars(data)
			if err != nil {
				log.Panic(err)
			}
			yerrs, err := plotter.NewYErrorBars(data)
			if err != nil {
				log.Panic(err)
			}

			xerrs.Color = cluColor
			yerrs.Color = cluColor

			p3.Add(xerrs, yerrs)
			p3.Legend.Add("all micro cluster ", c)

			// Save the plot to a PNG file.
			if err := p3.Save(20*vg.Inch, 20*vg.Inch, "datalayer/plot/m-"+fmt.Sprint(lastid)+".png"); err != nil {
				panic(err)
			}
		}
		//------------------------------------- clusters
		allXYs := toXYs(*clusters)
		allXYZs := toXYZs(*clusters, serv.Mepsilon)

		// Create a new plot, set its title and

		// axis labels.
		p2 := plot.New()
		// if err != nil {
		// 	panic(err)
		// }
		p2.Title.Text = "BADADE <Clusters>"
		p2.X.Label.Text = info.Header[0]
		p2.Y.Label.Text = info.Header[1]
		// Draw a grid behind the data
		p2.Add(plotter.NewGrid())

		for i := 0; i < len(allXYZs); i++ {

			cluColor := color.RGBA{R: uint8(rand.Intn(254)), G: uint8(rand.Intn(254)), B: uint8(rand.Intn(254)), A: 255}

			c, err := plotter.NewScatter(allXYZs[i])
			if err != nil {
				panic(err)
			}
			c.Color = cluColor
			n := len(allXYs[i])
			perr := make(plotter.Errors, n)
			for i := range perr {
				perr[i].Low = serv.Mepsilon
				perr[i].High = serv.Mepsilon
			}
			data := errPoints{
				XYs:     allXYs[i],
				YErrors: plotter.YErrors(perr),
				XErrors: plotter.XErrors(perr),
			}
			xerrs, err := plotter.NewXErrorBars(data)
			if err != nil {
				log.Panic(err)
			}
			yerrs, err := plotter.NewYErrorBars(data)
			if err != nil {
				log.Panic(err)
			}

			xerrs.Color = cluColor
			yerrs.Color = cluColor

			p2.Add(xerrs, yerrs)
			p2.Legend.Add("cluster "+fmt.Sprint(i+1), c)

		}
		// Save the plot to a PNG file.
		if err := p2.Save(20*vg.Inch, 20*vg.Inch, "datalayer/plot/c-"+fmt.Sprint(lastid)+".png"); err != nil {
			panic(err)
		}

	}

	select {
	case Progressbar[servkey] <- -1000:
	default:
	}
	return lastid, err
}

/*
func appendXYZ(pts *plotter.XYZs, po []float64) {
	var pt plotter.XYZ
	if len(po) == 2 {
		pt.X = po[0]
		pt.Y = po[1]
		pt.Z = 25
		*pts = append(*pts, pt)
	}
}
*/
