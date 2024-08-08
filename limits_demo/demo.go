package main

import (
	"fmt"
	"log"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/toolutils"
	"github.com/google/go-jsonnet/ast"
)
/*
func SafeVM(stack int) *jsonnet.VM{
	defaultImporter := &jsonnet.MemoryImporter{map[string]jsonnet.Contents{}}
	return &jsonnet.VM{
		MaxStack:       stack,
		ext:            make(jsonnet.vmExtMap),
		tla:            make(jsonnet.vmExtMap),
		nativeFuncs:    make(map[string]*jsonnet.NativeFunction),
		ErrorFormatter: &jsonnet.termErrorFormatter{pretty: false, maxStackTraceSize: 20},
		importer:       defaultImporter,
		importCache:    jsonnet.makeImportCache(defaultImporter),
		traceOut:       os.Stderr,
	}
}
*/
func printNode(a ast.Node){
	fmt.Printf("%T:",a)
	switch node := a.(type) {
	case *ast.Apply:
		fmt.Printf("%+v,\n %+v\n====\n",node.Target, node.Arguments)
		printNode(node.Target)
		for _,expr := range node.Arguments.Positional{
			printNode(expr.Expr)
		}
	case *ast.Index:
		fmt.Printf("%+v,\n %+v\n====\n",node.Target, node.Index)
		fmt.Printf("tgt:")
		printNode(node.Target)
		fmt.Printf("idx:")
		printNode(node.Index)
		

	default:
		fmt.Printf("%+v\n", a)
	}
}
func traverse(tree ast.Node){
	
	//fmt.Printf("%T::\t",tree)	
	printNode(tree)
	for _, sub := range toolutils.Children(tree){
	fmt.Printf("%T\t",sub)
	}
	fmt.Printf("\n")

	for _, sub := range toolutils.Children(tree){
	traverse(sub)
	}
}


func main() {
	vm := jsonnet.MakeVM()
	vm.Importer(&jsonnet.MemoryImporter{map[string]jsonnet.Contents{}})
	vm.MaxStack = 10
	vm.MaxSteps = 400 // During testing, it looks like setting up the environment and evaluating std costs 105 steps.
	vm.MaxAllocations = 400 // Change this around while changing how big the range allocated in the program is to see it work!
	snippet := `{
		data: std.range(1,300)
	}`

	jsonStr, err := vm.EvaluateAnonymousSnippet("example1.jsonnet", snippet)	
	if err != nil {
		log.Fatal(err)
	}
	tree,err := jsonnet.SnippetToAST("foo", snippet)
	traverse(tree)	
	fmt.Println(jsonStr)
	//fmt.Println(vm.Interpreter)

	/*
	   {
	     "person1": {
	         "name": "Alice",
	         "welcome": "Hello Alice!"
	     },
	     "person2": {
	         "name": "Bob",
	         "welcome": "Hello Bob!"
	     }
	   }
	*/
}
