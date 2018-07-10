package apps

var MiddlewaresComposer = []string{

	// curd interseptor
	"store@Curd#/api/data/test/,test",

	"middwares@Demo#root", // pkg@method#param1,param2
}

var PrefixOfRoot = "/"
var MethodOfRoot = "GET"
