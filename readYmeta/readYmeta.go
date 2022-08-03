/*
readYmeta.go reading and converting Yoda metadata
Author: Brett G. Olivier
email: @bgoli
licence: BSD 3 Clause
version: 0.3
(C) Brett G. Olivier, Vrije Universiteit Amsterdam, Amsterdam, The Netherlands, 2022
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

// Define global constants here?
const _VERSION_ = "0.3"

// Vanilla Yoda metadata struct
type Yoda18Metadata struct {
	Links []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
	Discipline []string `json:"Discipline"`
	Language   string   `json:"Language"`
	Collected  struct {
		StartDate string `json:"Start_Date"`
		EndDate   string `json:"End_Date"`
	} `json:"Collected"`
	CoveredGeolocationPlace []string `json:"Covered_Geolocation_Place"`
	CoveredPeriod           struct {
		StartDate string `json:"Start_Date"`
		EndDate   string `json:"End_Date"`
	} `json:"Covered_Period"`
	Tag                []string `json:"Tag"`
	RelatedDatapackage []struct {
		PersistentIdentifier struct {
			IdentifierScheme string `json:"Identifier_Scheme"`
			Identifier       string `json:"Identifier"`
		} `json:"Persistent_Identifier,omitempty"`
		RelationType          string `json:"Relation_Type"`
		Title                 string `json:"Title"`
		PersistentIdentifier0 struct {
			Identifier string `json:"Identifier"`
		} `json:"Persistent_Identifier,omitempty"`
	} `json:"Related_Datapackage"`
	RetentionPeriod  int    `json:"Retention_Period"`
	DataType         string `json:"Data_Type"`
	FundingReference []struct {
		FunderName  string `json:"Funder_Name"`
		AwardNumber string `json:"Award_Number"`
	} `json:"Funding_Reference"`
	Creator []struct {
		Name struct {
			GivenName  string `json:"Given_Name"`
			FamilyName string `json:"Family_Name"`
		} `json:"Name"`
		Affiliation      []string `json:"Affiliation"`
		PersonIdentifier []struct {
			NameIdentifierScheme string `json:"Name_Identifier_Scheme"`
			NameIdentifier       string `json:"Name_Identifier"`
		} `json:"Person_Identifier"`
	} `json:"Creator"`
	Contributor []struct {
		Name struct {
			GivenName  string `json:"Given_Name"`
			FamilyName string `json:"Family_Name"`
		} `json:"Name"`
		Affiliation      []string `json:"Affiliation"`
		PersonIdentifier []struct {
			NameIdentifierScheme string `json:"Name_Identifier_Scheme"`
			NameIdentifier       string `json:"Name_Identifier"`
		} `json:"Person_Identifier"`
		ContributorType string `json:"Contributor_Type"`
	} `json:"Contributor"`
	DataAccessRestriction string `json:"Data_Access_Restriction"`
	Title                 string `json:"Title"`
	Description           string `json:"Description"`
	Version               string `json:"Version"`
	RetentionInformation  string `json:"Retention_Information"`
	EmbargoEndDate        string `json:"Embargo_End_Date"`
	DataClassification    string `json:"Data_Classification"`
	CollectionName        string `json:"Collection_Name"`
	Remarks               string `json:"Remarks"`
	License               string `json:"License"`
}

// Yoda metadata struct with advanced options
type Yoda18MetadataV2 struct {
	Links []struct {
		Rel  string `json:"rel,omitempty"`
		Href string `json:"href,omitempty"`
	} `json:"links,omitempty"`
	Discipline []string `json:"Discipline,omitempty"`
	Language   string   `json:"Language,omitempty"`
	Collected  struct {
		StartDate string `json:"Start_Date,omitempty"`
		EndDate   string `json:"End_Date,omitempty"`
	} `json:"Collected,omitempty"`
	CoveredGeolocationPlace []string `json:"Covered_Geolocation_Place,omitempty"`
	CoveredPeriod           struct {
		StartDate string `json:"Start_Date,omitempty"`
		EndDate   string `json:"End_Date,omitempty"`
	} `json:"Covered_Period,omitempty"`
	Tag                []string `json:"Tag,omitempty"`
	RelatedDatapackage []struct {
		PersistentIdentifier struct {
			IdentifierScheme string `json:"Identifier_Scheme,omitempty"`
			Identifier       string `json:"Identifier,omitempty"`
		} `json:"Persistent_Identifier,omitempty"`
		RelationType          string `json:"Relation_Type,omitempty"`
		Title                 string `json:"Title,omitempty"`
		PersistentIdentifier0 struct {
			Identifier string `json:"Identifier,omitempty"`
		} `json:"Persistent_Identifier,omitempty"`
	} `json:"Related_Datapackage,omitempty"`
	RetentionPeriod  int    `json:"Retention_Period,omitempty"`
	DataType         string `json:"Data_Type,omitempty"`
	FundingReference []struct {
		FunderName  string `json:"Funder_Name,omitempty"`
		AwardNumber string `json:"Award_Number,omitempty"`
	} `json:"Funding_Reference,omitempty"`
	Creator []struct {
		Name struct {
			GivenName  string `json:"Given_Name,omitempty"`
			FamilyName string `json:"Family_Name,omitempty"`
		} `json:"Name,omitempty"`
		Affiliation      []string `json:"Affiliation,omitempty"`
		PersonIdentifier []struct {
			NameIdentifierScheme string `json:"Name_Identifier_Scheme,omitempty"`
			NameIdentifier       string `json:"Name_Identifier,omitempty"`
		} `json:"Person_Identifier,omitempty"`
	} `json:"Creator,omitempty"`
	Contributor []struct {
		Name struct {
			GivenName  string `json:"Given_Name,omitempty"`
			FamilyName string `json:"Family_Name,omitempty"`
		} `json:"Name,omitempty"`
		Affiliation      []string `json:"Affiliation,omitempty"`
		PersonIdentifier []struct {
			NameIdentifierScheme string `json:"Name_Identifier_Scheme,omitempty"`
			NameIdentifier       string `json:"Name_Identifier,omitempty"`
		} `json:"Person_Identifier,omitempty"`
		ContributorType string `json:"Contributor_Type,omitempty"`
	} `json:"Contributor,omitempty"`
	DataAccessRestriction string `json:"Data_Access_Restriction,omitempty"`
	Title                 string `json:"Title,omitempty"`
	Description           string `json:"Description,omitempty"`
	Version               string `json:"Version,omitempty"`
	RetentionInformation  string `json:"Retention_Information,omitempty"`
	EmbargoEndDate        string `json:"Embargo_End_Date,omitempty"`
	DataClassification    string `json:"Data_Classification,omitempty"`
	CollectionName        string `json:"Collection_Name,omitempty"`
	Remarks               string `json:"Remarks,omitempty"`
	License               string `json:"License,omitempty"`
}

func errcntrl(e error) {
	if e != nil {
		panic(e)
	}
}

func pdf_create_and_dump(fname string, sarr []string) {

	// Do things

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	//m.SetBorder(true)

	m.Row(10, func() {
		m.Col(4, func() {
			m.Text("Written by readYmeta, the Yoda Metadata converter - v"+_VERSION_, props.Text{
				Top:         0,
				Size:        16,
				Extrapolate: true,
			})
		})
		m.ColSpace(4)
	})
	// m.Row(10, func() {})
	m.Line(10)

	for idx, ele := range sarr {
		fmt.Println("Index :", idx, " Element :", ele)
		// pdf_write_row(m, fmt.Sprintln("Index :", idx, " Element :", ele))
		pdf_write_row(m, fmt.Sprintln(ele))
	}

	err := m.OutputFileAndClose(fmt.Sprintf("%s.pdf", fname))
	errcntrl(err)
}

func pdf_write_row(m pdf.Maroto, line string) {
	m.Row(6, func() {
		m.Col(4, func() {
			m.Text(line, props.Text{
				Top:         0,
				Size:        10,
				Extrapolate: true,
			})
		})
		//m.ColSpace(12)
	})
}

func main() {

	msg := "Welcome to the Yoda metadata translator\n(C)Brett G. Olivier, Vrije Universiteit Amsterdam, 2022"
	fmt.Println(msg)

	// read metadata file
	json_file, err1 := os.ReadFile("yoda-metadata.json")
	errcntrl(err1)

	// print the file cast as string
	fmt.Print(string(json_file))

	// create metadata struct and fill it with file data
	var json_dat Yoda18Metadata
	err2 := json.Unmarshal(json_file, &json_dat)
	errcntrl(err2)

	// print struct and explore new options
	/*
		fmt.Println(" ")
		fmt.Println(json_dat)
		fmt.Println(reflect.TypeOf(json_dat))
		fmt.Println(len(json_dat.Contributor))
	*/
	// lets do something more useful
	fmt.Printf("\n\n----------------\n\n")
	var basic_info_str []string
	basic_info_str = get_basic_data(json_dat)

	// Random checks
	fmt.Println(basic_info_str)
	fmt.Println(reflect.TypeOf(basic_info_str))
	fmt.Println(basic_info_str[0])

	// lets play with dumping to PDF
	pdf_create_and_dump("dumpfile", basic_info_str)
}

func get_basic_data(doc Yoda18Metadata) []string {
	var output []string
	output = append(output, fmt.Sprintf("Title: %s", doc.Title))
	output = append(output, fmt.Sprintf("Description: %s", doc.Description))
	output = append(output, fmt.Sprintf("Version: %s", doc.Version))
	output = append(output, fmt.Sprintf("Licence: %s", doc.License))
	output = append(output, fmt.Sprintf("Language: %s", doc.Language))
	output = append(output, fmt.Sprintf("Rentention_Period: %d", doc.RetentionPeriod))
	output = append(output, fmt.Sprintf("Data_Type: %s", doc.DataType))
	output = append(output, fmt.Sprintf("Data_Access_Restriction: %s", doc.DataAccessRestriction))
	output = append(output, fmt.Sprintf("Retention_Information: %s", doc.RetentionInformation))
	output = append(output, fmt.Sprintf("Embargo_End_Date: %s", doc.EmbargoEndDate))
	output = append(output, fmt.Sprintf("Data_Classification: %s", doc.DataClassification))
	output = append(output, fmt.Sprintf("Collection_Name: %s", doc.CollectionName))
	output = append(output, fmt.Sprintf("Remarks: %s", doc.Remarks))
	return output
}

/*
Seems to work!
Funky GO template builder: https://mholt.github.io/json-to-go/
*/

// learning how not to do stuff
/*
func read_json_branch(json_map map[string]interface{}) {
	fmt.Println(" ")
	fmt.Println("json_map")
	fmt.Println(json_map)
	fmt.Println(reflect.TypeOf(json_map))

	for k, v := range json_map {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			fmt.Println(reflect.TypeOf(k))
			fmt.Println(reflect.TypeOf(v))
			fmt.Println(reflect.TypeOf(vv))
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}


*/
