package comment

import (
	"reflect"

	"github.com/tenntenn/comment"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name:             "comment",
	Doc:              "create comment map",
	Run:              run,
	RunDespiteErrors: true,
	ResultType:       reflect.TypeOf(new(comment.Maps)),
}

func run(pass *analysis.Pass) (interface{}, error) {
	return comment.New(pass.Fset, pass.Files), nil
}
