package parser

import (
	"regexp"
	"strings"
)

// SsacliPhysDisk data structure for output
type SsacliPhysDisk struct {
	SsacliPhysDiskData []SsacliPhysDiskData
}

// SsacliPhysDiskData data structure for output
type SsacliPhysDiskData struct {
	Bay       string
	Status    string
	DriveType string
	IntType   string
	Size      string
	BlockSize string
	Speed     string
	Firmware  string
	SN        string
	WWID      string
	CurTemp   float64
	MaxTemp   float64
	Model     string
}

// ParseSsacliPhysDisk return specific metric
func ParseSsacliPhysDisk(s string) *SsacliPhysDisk {
	data := parseSsacliPhysDisk(s)

	return data
}

func parseSsacliPhysDisk(s string) *SsacliPhysDisk {

	var (
		tmp SsacliPhysDiskData
	)
	re := regexp.MustCompile(`(.+?)\: (.+)`)
	for _, line := range strings.Split(s, "\n") {
		kvs := strings.Trim(line, " \t")
		kv := re.FindStringSubmatch(kvs)

		/* The input looks like this:
		   physicaldrive 1I:1:2
		   Port: 1I
		   Box: 1
		   Bay: 2
		   Status: OK
		   Drive Type: Data Drive
		   Interface Type: SAS
		   Size: 72 GB
		   Drive exposed to OS: False
		   Logical/Physical Block Size: 512/512
		   Rotational Speed: 15000
		   Firmware Revision: HPD2 (FW update is recommended to minimum version: HPD7)
		   Serial Number: 3PD0EYNW00009744Q72N
		   WWID: 5000C5000522B3B1
		   Model: HP      DH072ABAA6
		   Current Temperature (C): 32
		   Maximum Temperature (C): 41
		   PHY Count: 1
		   PHY Transfer Rate: 3.0Gbps
		   PHY Physical Link Rate: Unknown
		   PHY Maximum Link Rate: Unknown
		   Sanitize Erase Supported: False
		   Shingled Magnetic Recording Support: None
		*/
		if len(kv) == 3 {
			key := kv[1]
			value := kv[2]
			switch key {
			case "Bay":
				tmp.Bay = value
			case "Serial Number":
				tmp.SN = value
			case "Status":
				tmp.Status = value
			case "Drive Type":
				tmp.DriveType = value
			case "Interface Type":
				tmp.IntType = value
			case "Size":
				tmp.Size = value
			case "Logical/Physical Block Size":
				tmp.BlockSize = value
			case "Rotational Speed":
				tmp.Speed = value
			case "Firmware Revision":
				tmp.Firmware = value
			case "WWID":
				tmp.WWID = value
			case "Model":
				tmp.Model = value
			case "Current Temperature (C)":
				tmp.CurTemp = toFLO(value)
			case "Maximum Temperature (C)":
				tmp.MaxTemp = toFLO(value)
			}
		}
	}

	data := SsacliPhysDisk{
		SsacliPhysDiskData: []SsacliPhysDiskData{
			tmp,
		},
	}
	return &data
}
