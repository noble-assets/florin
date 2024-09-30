// Code generated by generate.go. DO NOT EDIT.

package staticcheck

import (
	"honnef.co/go/tools/analysis/lint"
	"honnef.co/go/tools/staticcheck/sa1000"
	"honnef.co/go/tools/staticcheck/sa1001"
	"honnef.co/go/tools/staticcheck/sa1002"
	"honnef.co/go/tools/staticcheck/sa1003"
	"honnef.co/go/tools/staticcheck/sa1004"
	"honnef.co/go/tools/staticcheck/sa1005"
	"honnef.co/go/tools/staticcheck/sa1006"
	"honnef.co/go/tools/staticcheck/sa1007"
	"honnef.co/go/tools/staticcheck/sa1008"
	"honnef.co/go/tools/staticcheck/sa1010"
	"honnef.co/go/tools/staticcheck/sa1011"
	"honnef.co/go/tools/staticcheck/sa1012"
	"honnef.co/go/tools/staticcheck/sa1013"
	"honnef.co/go/tools/staticcheck/sa1014"
	"honnef.co/go/tools/staticcheck/sa1015"
	"honnef.co/go/tools/staticcheck/sa1016"
	"honnef.co/go/tools/staticcheck/sa1017"
	"honnef.co/go/tools/staticcheck/sa1018"
	"honnef.co/go/tools/staticcheck/sa1019"
	"honnef.co/go/tools/staticcheck/sa1020"
	"honnef.co/go/tools/staticcheck/sa1021"
	"honnef.co/go/tools/staticcheck/sa1023"
	"honnef.co/go/tools/staticcheck/sa1024"
	"honnef.co/go/tools/staticcheck/sa1025"
	"honnef.co/go/tools/staticcheck/sa1026"
	"honnef.co/go/tools/staticcheck/sa1027"
	"honnef.co/go/tools/staticcheck/sa1028"
	"honnef.co/go/tools/staticcheck/sa1029"
	"honnef.co/go/tools/staticcheck/sa1030"
	"honnef.co/go/tools/staticcheck/sa1031"
	"honnef.co/go/tools/staticcheck/sa1032"
	"honnef.co/go/tools/staticcheck/sa2000"
	"honnef.co/go/tools/staticcheck/sa2001"
	"honnef.co/go/tools/staticcheck/sa2002"
	"honnef.co/go/tools/staticcheck/sa2003"
	"honnef.co/go/tools/staticcheck/sa3000"
	"honnef.co/go/tools/staticcheck/sa3001"
	"honnef.co/go/tools/staticcheck/sa4000"
	"honnef.co/go/tools/staticcheck/sa4001"
	"honnef.co/go/tools/staticcheck/sa4003"
	"honnef.co/go/tools/staticcheck/sa4004"
	"honnef.co/go/tools/staticcheck/sa4005"
	"honnef.co/go/tools/staticcheck/sa4006"
	"honnef.co/go/tools/staticcheck/sa4008"
	"honnef.co/go/tools/staticcheck/sa4009"
	"honnef.co/go/tools/staticcheck/sa4010"
	"honnef.co/go/tools/staticcheck/sa4011"
	"honnef.co/go/tools/staticcheck/sa4012"
	"honnef.co/go/tools/staticcheck/sa4013"
	"honnef.co/go/tools/staticcheck/sa4014"
	"honnef.co/go/tools/staticcheck/sa4015"
	"honnef.co/go/tools/staticcheck/sa4016"
	"honnef.co/go/tools/staticcheck/sa4017"
	"honnef.co/go/tools/staticcheck/sa4018"
	"honnef.co/go/tools/staticcheck/sa4019"
	"honnef.co/go/tools/staticcheck/sa4020"
	"honnef.co/go/tools/staticcheck/sa4021"
	"honnef.co/go/tools/staticcheck/sa4022"
	"honnef.co/go/tools/staticcheck/sa4023"
	"honnef.co/go/tools/staticcheck/sa4024"
	"honnef.co/go/tools/staticcheck/sa4025"
	"honnef.co/go/tools/staticcheck/sa4026"
	"honnef.co/go/tools/staticcheck/sa4027"
	"honnef.co/go/tools/staticcheck/sa4028"
	"honnef.co/go/tools/staticcheck/sa4029"
	"honnef.co/go/tools/staticcheck/sa4030"
	"honnef.co/go/tools/staticcheck/sa4031"
	"honnef.co/go/tools/staticcheck/sa4032"
	"honnef.co/go/tools/staticcheck/sa5000"
	"honnef.co/go/tools/staticcheck/sa5001"
	"honnef.co/go/tools/staticcheck/sa5002"
	"honnef.co/go/tools/staticcheck/sa5003"
	"honnef.co/go/tools/staticcheck/sa5004"
	"honnef.co/go/tools/staticcheck/sa5005"
	"honnef.co/go/tools/staticcheck/sa5007"
	"honnef.co/go/tools/staticcheck/sa5008"
	"honnef.co/go/tools/staticcheck/sa5009"
	"honnef.co/go/tools/staticcheck/sa5010"
	"honnef.co/go/tools/staticcheck/sa5011"
	"honnef.co/go/tools/staticcheck/sa5012"
	"honnef.co/go/tools/staticcheck/sa6000"
	"honnef.co/go/tools/staticcheck/sa6001"
	"honnef.co/go/tools/staticcheck/sa6002"
	"honnef.co/go/tools/staticcheck/sa6003"
	"honnef.co/go/tools/staticcheck/sa6005"
	"honnef.co/go/tools/staticcheck/sa6006"
	"honnef.co/go/tools/staticcheck/sa9001"
	"honnef.co/go/tools/staticcheck/sa9002"
	"honnef.co/go/tools/staticcheck/sa9003"
	"honnef.co/go/tools/staticcheck/sa9004"
	"honnef.co/go/tools/staticcheck/sa9005"
	"honnef.co/go/tools/staticcheck/sa9006"
	"honnef.co/go/tools/staticcheck/sa9007"
	"honnef.co/go/tools/staticcheck/sa9008"
	"honnef.co/go/tools/staticcheck/sa9009"
)

var Analyzers = []*lint.Analyzer{
	sa1000.SCAnalyzer,
	sa1001.SCAnalyzer,
	sa1002.SCAnalyzer,
	sa1003.SCAnalyzer,
	sa1004.SCAnalyzer,
	sa1005.SCAnalyzer,
	sa1006.SCAnalyzer,
	sa1007.SCAnalyzer,
	sa1008.SCAnalyzer,
	sa1010.SCAnalyzer,
	sa1011.SCAnalyzer,
	sa1012.SCAnalyzer,
	sa1013.SCAnalyzer,
	sa1014.SCAnalyzer,
	sa1015.SCAnalyzer,
	sa1016.SCAnalyzer,
	sa1017.SCAnalyzer,
	sa1018.SCAnalyzer,
	sa1019.SCAnalyzer,
	sa1020.SCAnalyzer,
	sa1021.SCAnalyzer,
	sa1023.SCAnalyzer,
	sa1024.SCAnalyzer,
	sa1025.SCAnalyzer,
	sa1026.SCAnalyzer,
	sa1027.SCAnalyzer,
	sa1028.SCAnalyzer,
	sa1029.SCAnalyzer,
	sa1030.SCAnalyzer,
	sa1031.SCAnalyzer,
	sa1032.SCAnalyzer,
	sa2000.SCAnalyzer,
	sa2001.SCAnalyzer,
	sa2002.SCAnalyzer,
	sa2003.SCAnalyzer,
	sa3000.SCAnalyzer,
	sa3001.SCAnalyzer,
	sa4000.SCAnalyzer,
	sa4001.SCAnalyzer,
	sa4003.SCAnalyzer,
	sa4004.SCAnalyzer,
	sa4005.SCAnalyzer,
	sa4006.SCAnalyzer,
	sa4008.SCAnalyzer,
	sa4009.SCAnalyzer,
	sa4010.SCAnalyzer,
	sa4011.SCAnalyzer,
	sa4012.SCAnalyzer,
	sa4013.SCAnalyzer,
	sa4014.SCAnalyzer,
	sa4015.SCAnalyzer,
	sa4016.SCAnalyzer,
	sa4017.SCAnalyzer,
	sa4018.SCAnalyzer,
	sa4019.SCAnalyzer,
	sa4020.SCAnalyzer,
	sa4021.SCAnalyzer,
	sa4022.SCAnalyzer,
	sa4023.SCAnalyzer,
	sa4024.SCAnalyzer,
	sa4025.SCAnalyzer,
	sa4026.SCAnalyzer,
	sa4027.SCAnalyzer,
	sa4028.SCAnalyzer,
	sa4029.SCAnalyzer,
	sa4030.SCAnalyzer,
	sa4031.SCAnalyzer,
	sa4032.SCAnalyzer,
	sa5000.SCAnalyzer,
	sa5001.SCAnalyzer,
	sa5002.SCAnalyzer,
	sa5003.SCAnalyzer,
	sa5004.SCAnalyzer,
	sa5005.SCAnalyzer,
	sa5007.SCAnalyzer,
	sa5008.SCAnalyzer,
	sa5009.SCAnalyzer,
	sa5010.SCAnalyzer,
	sa5011.SCAnalyzer,
	sa5012.SCAnalyzer,
	sa6000.SCAnalyzer,
	sa6001.SCAnalyzer,
	sa6002.SCAnalyzer,
	sa6003.SCAnalyzer,
	sa6005.SCAnalyzer,
	sa6006.SCAnalyzer,
	sa9001.SCAnalyzer,
	sa9002.SCAnalyzer,
	sa9003.SCAnalyzer,
	sa9004.SCAnalyzer,
	sa9005.SCAnalyzer,
	sa9006.SCAnalyzer,
	sa9007.SCAnalyzer,
	sa9008.SCAnalyzer,
	sa9009.SCAnalyzer,
}