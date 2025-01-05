package util_string

import (
	"archive/zip"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type UtilsFile struct{}

func UseFileHelper() UtilsFile {
	return UtilsFile{}
}

func (u UtilsFile) GetPosrKMZ(filspath string) (*float64, *float64, error) {
	type Placemark struct {
		Name        string `xml:"name"`
		Coordinates string `xml:"Point>coordinates"`
	}
	type KMLData struct {
		Placemarks []Placemark `xml:"Document>Placemark"`
	}

	r, err := zip.OpenReader(filspath)
	if err != nil {
		log.Println("Error opening .kmz file:", err)
		return nil, nil, err
	}
	defer r.Close()

	var kmlContent []byte
	for _, file := range r.File {
		if strings.HasSuffix(file.Name, ".kml") {
			rc, err := file.Open()
			if err != nil {
				log.Println("Error opening .kml file inside .kmz:", err)
				return nil, nil, err
			}
			defer rc.Close()

			kmlContent, err = io.ReadAll(rc)
			if err != nil {
				log.Println("Error reading .kml file:", err)
				return nil, nil, err
			}
			break
		}
	}

	if kmlContent == nil {
		log.Println("No .kml file found inside .kmz")
		return nil, nil, err
	}

	var kmlData KMLData
	err = xml.Unmarshal(kmlContent, &kmlData)
	if err != nil {
		log.Println("Error parsing .kml file:", err)
		return nil, nil, err
	}

	for _, placemark := range kmlData.Placemarks {
		// log.Println("1")
		fmt.Printf("Name: %s\n", placemark.Name)
		coords := strings.Split(placemark.Coordinates, ",")
		if len(coords) >= 2 {

			latitude, err := strconv.ParseFloat(coords[1], 64)
			if err != nil {
				return nil, nil, err
			}
			longitude, err := strconv.ParseFloat(coords[0], 64)
			if err != nil {
				return nil, nil, err
			}
			return &latitude, &longitude, nil
		} else {
			return nil, nil, errors.New("[Error] : Position data is null")
		}
	}

	return nil, nil, errors.New("[Error] : invalid process")
}
