package apps

var MiddlewaresComposer = []string{

	// crud interseptor
	"store@Crud#/api/data/test/,test",

	"middwares@Demo#root", // pkg@method#param1,param2
}

var PrefixOfRoot = "/"
var MethodOfRoot = "GET"
