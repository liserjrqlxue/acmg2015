package evidence

import (
	"strconv"

	simple_util "github.com/liserjrqlxue/simple-util"
)

var BP3Func = map[string]bool{
	"cds-del":   true,
	"cds-ins":   true,
	"cds-indel": true,
}

// ture	:	"1"
// flase:	"0"
func CheckBP3(item map[string]string) string {
	if BP3Func[item["Function"]] && item["VarType"] != "snv" {
		if item["RepeatTag"] == "" || item["RepeatTag"] == "." {
			return "0"
		} else {
			if item["cHGVS_org"] == "" {
				item["cHGVS_org"] = item["cHGVS"]
			}
			subMatch := repeatSeq.FindStringSubmatch(item["cHGVS_org"])
			if len(subMatch) > 1 {
				dupCount, err := strconv.Atoi(subMatch[1])
				simple_util.CheckErr(err)
				if dupCount < 10 {
					return "0"
				} else {
					return "1"
				}
			} else {
				return "1"
			}
		}
	}
	return "0"
}

func CompareBP3(item map[string]string) {
	rule := "BP3"
	val := CheckBP3(item)
	if val != item[rule] {
		if item[rule] == "0" && val == "" {
		} else {
			PrintConflict(item, rule, val, "Function", "RepeatTag", "VarType")
		}
	}
}
