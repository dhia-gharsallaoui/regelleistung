/*
IP RL BSP API

IP RL BSP API for participation in capacity/energy market tenders.

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ProductName Representation of a product timeslice.
type ProductName string

// List of ProductName
const (
	NEG_00_04 ProductName = "NEG_00_04"
	NEG_04_08 ProductName = "NEG_04_08"
	NEG_08_12 ProductName = "NEG_08_12"
	NEG_12_16 ProductName = "NEG_12_16"
	NEG_16_20 ProductName = "NEG_16_20"
	NEG_20_24 ProductName = "NEG_20_24"
	POS_00_04 ProductName = "POS_00_04"
	POS_04_08 ProductName = "POS_04_08"
	POS_08_12 ProductName = "POS_08_12"
	POS_12_16 ProductName = "POS_12_16"
	POS_16_20 ProductName = "POS_16_20"
	POS_20_24 ProductName = "POS_20_24"
	NEGPOS_00_04 ProductName = "NEGPOS_00_04"
	NEGPOS_04_08 ProductName = "NEGPOS_04_08"
	NEGPOS_08_12 ProductName = "NEGPOS_08_12"
	NEGPOS_12_16 ProductName = "NEGPOS_12_16"
	NEGPOS_16_20 ProductName = "NEGPOS_16_20"
	NEGPOS_20_24 ProductName = "NEGPOS_20_24"
	NEGPOS_00_24 ProductName = "NEGPOS_00_24"
	NEG_001 ProductName = "NEG_001"
	NEG_002 ProductName = "NEG_002"
	NEG_003 ProductName = "NEG_003"
	NEG_004 ProductName = "NEG_004"
	NEG_005 ProductName = "NEG_005"
	NEG_006 ProductName = "NEG_006"
	NEG_007 ProductName = "NEG_007"
	NEG_008 ProductName = "NEG_008"
	NEG_009 ProductName = "NEG_009"
	NEG_010 ProductName = "NEG_010"
	NEG_011 ProductName = "NEG_011"
	NEG_012 ProductName = "NEG_012"
	NEG_013 ProductName = "NEG_013"
	NEG_014 ProductName = "NEG_014"
	NEG_015 ProductName = "NEG_015"
	NEG_016 ProductName = "NEG_016"
	NEG_017 ProductName = "NEG_017"
	NEG_018 ProductName = "NEG_018"
	NEG_019 ProductName = "NEG_019"
	NEG_020 ProductName = "NEG_020"
	NEG_021 ProductName = "NEG_021"
	NEG_022 ProductName = "NEG_022"
	NEG_023 ProductName = "NEG_023"
	NEG_024 ProductName = "NEG_024"
	NEG_025 ProductName = "NEG_025"
	NEG_026 ProductName = "NEG_026"
	NEG_027 ProductName = "NEG_027"
	NEG_028 ProductName = "NEG_028"
	NEG_029 ProductName = "NEG_029"
	NEG_030 ProductName = "NEG_030"
	NEG_031 ProductName = "NEG_031"
	NEG_032 ProductName = "NEG_032"
	NEG_033 ProductName = "NEG_033"
	NEG_034 ProductName = "NEG_034"
	NEG_035 ProductName = "NEG_035"
	NEG_036 ProductName = "NEG_036"
	NEG_037 ProductName = "NEG_037"
	NEG_038 ProductName = "NEG_038"
	NEG_039 ProductName = "NEG_039"
	NEG_040 ProductName = "NEG_040"
	NEG_041 ProductName = "NEG_041"
	NEG_042 ProductName = "NEG_042"
	NEG_043 ProductName = "NEG_043"
	NEG_044 ProductName = "NEG_044"
	NEG_045 ProductName = "NEG_045"
	NEG_046 ProductName = "NEG_046"
	NEG_047 ProductName = "NEG_047"
	NEG_048 ProductName = "NEG_048"
	NEG_049 ProductName = "NEG_049"
	NEG_050 ProductName = "NEG_050"
	NEG_051 ProductName = "NEG_051"
	NEG_052 ProductName = "NEG_052"
	NEG_053 ProductName = "NEG_053"
	NEG_054 ProductName = "NEG_054"
	NEG_055 ProductName = "NEG_055"
	NEG_056 ProductName = "NEG_056"
	NEG_057 ProductName = "NEG_057"
	NEG_058 ProductName = "NEG_058"
	NEG_059 ProductName = "NEG_059"
	NEG_060 ProductName = "NEG_060"
	NEG_061 ProductName = "NEG_061"
	NEG_062 ProductName = "NEG_062"
	NEG_063 ProductName = "NEG_063"
	NEG_064 ProductName = "NEG_064"
	NEG_065 ProductName = "NEG_065"
	NEG_066 ProductName = "NEG_066"
	NEG_067 ProductName = "NEG_067"
	NEG_068 ProductName = "NEG_068"
	NEG_069 ProductName = "NEG_069"
	NEG_070 ProductName = "NEG_070"
	NEG_071 ProductName = "NEG_071"
	NEG_072 ProductName = "NEG_072"
	NEG_073 ProductName = "NEG_073"
	NEG_074 ProductName = "NEG_074"
	NEG_075 ProductName = "NEG_075"
	NEG_076 ProductName = "NEG_076"
	NEG_077 ProductName = "NEG_077"
	NEG_078 ProductName = "NEG_078"
	NEG_079 ProductName = "NEG_079"
	NEG_080 ProductName = "NEG_080"
	NEG_081 ProductName = "NEG_081"
	NEG_082 ProductName = "NEG_082"
	NEG_083 ProductName = "NEG_083"
	NEG_084 ProductName = "NEG_084"
	NEG_085 ProductName = "NEG_085"
	NEG_086 ProductName = "NEG_086"
	NEG_087 ProductName = "NEG_087"
	NEG_088 ProductName = "NEG_088"
	NEG_089 ProductName = "NEG_089"
	NEG_090 ProductName = "NEG_090"
	NEG_091 ProductName = "NEG_091"
	NEG_092 ProductName = "NEG_092"
	NEG_093 ProductName = "NEG_093"
	NEG_094 ProductName = "NEG_094"
	NEG_095 ProductName = "NEG_095"
	NEG_096 ProductName = "NEG_096"
	NEG_097 ProductName = "NEG_097"
	NEG_098 ProductName = "NEG_098"
	NEG_099 ProductName = "NEG_099"
	NEG_100 ProductName = "NEG_100"
	POS_001 ProductName = "POS_001"
	POS_002 ProductName = "POS_002"
	POS_003 ProductName = "POS_003"
	POS_004 ProductName = "POS_004"
	POS_005 ProductName = "POS_005"
	POS_006 ProductName = "POS_006"
	POS_007 ProductName = "POS_007"
	POS_008 ProductName = "POS_008"
	POS_009 ProductName = "POS_009"
	POS_010 ProductName = "POS_010"
	POS_011 ProductName = "POS_011"
	POS_012 ProductName = "POS_012"
	POS_013 ProductName = "POS_013"
	POS_014 ProductName = "POS_014"
	POS_015 ProductName = "POS_015"
	POS_016 ProductName = "POS_016"
	POS_017 ProductName = "POS_017"
	POS_018 ProductName = "POS_018"
	POS_019 ProductName = "POS_019"
	POS_020 ProductName = "POS_020"
	POS_021 ProductName = "POS_021"
	POS_022 ProductName = "POS_022"
	POS_023 ProductName = "POS_023"
	POS_024 ProductName = "POS_024"
	POS_025 ProductName = "POS_025"
	POS_026 ProductName = "POS_026"
	POS_027 ProductName = "POS_027"
	POS_028 ProductName = "POS_028"
	POS_029 ProductName = "POS_029"
	POS_030 ProductName = "POS_030"
	POS_031 ProductName = "POS_031"
	POS_032 ProductName = "POS_032"
	POS_033 ProductName = "POS_033"
	POS_034 ProductName = "POS_034"
	POS_035 ProductName = "POS_035"
	POS_036 ProductName = "POS_036"
	POS_037 ProductName = "POS_037"
	POS_038 ProductName = "POS_038"
	POS_039 ProductName = "POS_039"
	POS_040 ProductName = "POS_040"
	POS_041 ProductName = "POS_041"
	POS_042 ProductName = "POS_042"
	POS_043 ProductName = "POS_043"
	POS_044 ProductName = "POS_044"
	POS_045 ProductName = "POS_045"
	POS_046 ProductName = "POS_046"
	POS_047 ProductName = "POS_047"
	POS_048 ProductName = "POS_048"
	POS_049 ProductName = "POS_049"
	POS_050 ProductName = "POS_050"
	POS_051 ProductName = "POS_051"
	POS_052 ProductName = "POS_052"
	POS_053 ProductName = "POS_053"
	POS_054 ProductName = "POS_054"
	POS_055 ProductName = "POS_055"
	POS_056 ProductName = "POS_056"
	POS_057 ProductName = "POS_057"
	POS_058 ProductName = "POS_058"
	POS_059 ProductName = "POS_059"
	POS_060 ProductName = "POS_060"
	POS_061 ProductName = "POS_061"
	POS_062 ProductName = "POS_062"
	POS_063 ProductName = "POS_063"
	POS_064 ProductName = "POS_064"
	POS_065 ProductName = "POS_065"
	POS_066 ProductName = "POS_066"
	POS_067 ProductName = "POS_067"
	POS_068 ProductName = "POS_068"
	POS_069 ProductName = "POS_069"
	POS_070 ProductName = "POS_070"
	POS_071 ProductName = "POS_071"
	POS_072 ProductName = "POS_072"
	POS_073 ProductName = "POS_073"
	POS_074 ProductName = "POS_074"
	POS_075 ProductName = "POS_075"
	POS_076 ProductName = "POS_076"
	POS_077 ProductName = "POS_077"
	POS_078 ProductName = "POS_078"
	POS_079 ProductName = "POS_079"
	POS_080 ProductName = "POS_080"
	POS_081 ProductName = "POS_081"
	POS_082 ProductName = "POS_082"
	POS_083 ProductName = "POS_083"
	POS_084 ProductName = "POS_084"
	POS_085 ProductName = "POS_085"
	POS_086 ProductName = "POS_086"
	POS_087 ProductName = "POS_087"
	POS_088 ProductName = "POS_088"
	POS_089 ProductName = "POS_089"
	POS_090 ProductName = "POS_090"
	POS_091 ProductName = "POS_091"
	POS_092 ProductName = "POS_092"
	POS_093 ProductName = "POS_093"
	POS_094 ProductName = "POS_094"
	POS_095 ProductName = "POS_095"
	POS_096 ProductName = "POS_096"
	POS_097 ProductName = "POS_097"
	POS_098 ProductName = "POS_098"
	POS_099 ProductName = "POS_099"
	POS_100 ProductName = "POS_100"
)

// All allowed values of ProductName enum
var AllowedProductNameEnumValues = []ProductName{
	"NEG_00_04",
	"NEG_04_08",
	"NEG_08_12",
	"NEG_12_16",
	"NEG_16_20",
	"NEG_20_24",
	"POS_00_04",
	"POS_04_08",
	"POS_08_12",
	"POS_12_16",
	"POS_16_20",
	"POS_20_24",
	"NEGPOS_00_04",
	"NEGPOS_04_08",
	"NEGPOS_08_12",
	"NEGPOS_12_16",
	"NEGPOS_16_20",
	"NEGPOS_20_24",
	"NEGPOS_00_24",
	"NEG_001",
	"NEG_002",
	"NEG_003",
	"NEG_004",
	"NEG_005",
	"NEG_006",
	"NEG_007",
	"NEG_008",
	"NEG_009",
	"NEG_010",
	"NEG_011",
	"NEG_012",
	"NEG_013",
	"NEG_014",
	"NEG_015",
	"NEG_016",
	"NEG_017",
	"NEG_018",
	"NEG_019",
	"NEG_020",
	"NEG_021",
	"NEG_022",
	"NEG_023",
	"NEG_024",
	"NEG_025",
	"NEG_026",
	"NEG_027",
	"NEG_028",
	"NEG_029",
	"NEG_030",
	"NEG_031",
	"NEG_032",
	"NEG_033",
	"NEG_034",
	"NEG_035",
	"NEG_036",
	"NEG_037",
	"NEG_038",
	"NEG_039",
	"NEG_040",
	"NEG_041",
	"NEG_042",
	"NEG_043",
	"NEG_044",
	"NEG_045",
	"NEG_046",
	"NEG_047",
	"NEG_048",
	"NEG_049",
	"NEG_050",
	"NEG_051",
	"NEG_052",
	"NEG_053",
	"NEG_054",
	"NEG_055",
	"NEG_056",
	"NEG_057",
	"NEG_058",
	"NEG_059",
	"NEG_060",
	"NEG_061",
	"NEG_062",
	"NEG_063",
	"NEG_064",
	"NEG_065",
	"NEG_066",
	"NEG_067",
	"NEG_068",
	"NEG_069",
	"NEG_070",
	"NEG_071",
	"NEG_072",
	"NEG_073",
	"NEG_074",
	"NEG_075",
	"NEG_076",
	"NEG_077",
	"NEG_078",
	"NEG_079",
	"NEG_080",
	"NEG_081",
	"NEG_082",
	"NEG_083",
	"NEG_084",
	"NEG_085",
	"NEG_086",
	"NEG_087",
	"NEG_088",
	"NEG_089",
	"NEG_090",
	"NEG_091",
	"NEG_092",
	"NEG_093",
	"NEG_094",
	"NEG_095",
	"NEG_096",
	"NEG_097",
	"NEG_098",
	"NEG_099",
	"NEG_100",
	"POS_001",
	"POS_002",
	"POS_003",
	"POS_004",
	"POS_005",
	"POS_006",
	"POS_007",
	"POS_008",
	"POS_009",
	"POS_010",
	"POS_011",
	"POS_012",
	"POS_013",
	"POS_014",
	"POS_015",
	"POS_016",
	"POS_017",
	"POS_018",
	"POS_019",
	"POS_020",
	"POS_021",
	"POS_022",
	"POS_023",
	"POS_024",
	"POS_025",
	"POS_026",
	"POS_027",
	"POS_028",
	"POS_029",
	"POS_030",
	"POS_031",
	"POS_032",
	"POS_033",
	"POS_034",
	"POS_035",
	"POS_036",
	"POS_037",
	"POS_038",
	"POS_039",
	"POS_040",
	"POS_041",
	"POS_042",
	"POS_043",
	"POS_044",
	"POS_045",
	"POS_046",
	"POS_047",
	"POS_048",
	"POS_049",
	"POS_050",
	"POS_051",
	"POS_052",
	"POS_053",
	"POS_054",
	"POS_055",
	"POS_056",
	"POS_057",
	"POS_058",
	"POS_059",
	"POS_060",
	"POS_061",
	"POS_062",
	"POS_063",
	"POS_064",
	"POS_065",
	"POS_066",
	"POS_067",
	"POS_068",
	"POS_069",
	"POS_070",
	"POS_071",
	"POS_072",
	"POS_073",
	"POS_074",
	"POS_075",
	"POS_076",
	"POS_077",
	"POS_078",
	"POS_079",
	"POS_080",
	"POS_081",
	"POS_082",
	"POS_083",
	"POS_084",
	"POS_085",
	"POS_086",
	"POS_087",
	"POS_088",
	"POS_089",
	"POS_090",
	"POS_091",
	"POS_092",
	"POS_093",
	"POS_094",
	"POS_095",
	"POS_096",
	"POS_097",
	"POS_098",
	"POS_099",
	"POS_100",
}

func (v *ProductName) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProductName(value)
	for _, existing := range AllowedProductNameEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProductName", value)
}

// NewProductNameFromValue returns a pointer to a valid ProductName
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProductNameFromValue(v string) (*ProductName, error) {
	ev := ProductName(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProductName: valid values are %v", v, AllowedProductNameEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProductName) IsValid() bool {
	for _, existing := range AllowedProductNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductName value
func (v ProductName) Ptr() *ProductName {
	return &v
}

type NullableProductName struct {
	value *ProductName
	isSet bool
}

func (v NullableProductName) Get() *ProductName {
	return v.value
}

func (v *NullableProductName) Set(val *ProductName) {
	v.value = val
	v.isSet = true
}

func (v NullableProductName) IsSet() bool {
	return v.isSet
}

func (v *NullableProductName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductName(val *ProductName) *NullableProductName {
	return &NullableProductName{value: val, isSet: true}
}

func (v NullableProductName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
