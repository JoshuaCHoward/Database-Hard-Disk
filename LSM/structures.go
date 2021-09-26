package LSM

import (
	"bytes"
	"encoding/gob"
	"github.com/google/btree"
	"os"
)

type LinkedList struct {
	length       int
	List         *Node
	fileLocation string
	treeFile     string
	tree         *btree.BTree
}

type Node struct {
	left  *Node
	right *Node
	value int
}

//func (i Item) Less(than Item) bool {
//	if i.value<than.value {
//		return true
//	} else {
//		return false
//	}
//}

//func DecodeTreeFile(buffer *LinkedList) {
//	//body,_:=ioutil.ReadFile(buffer.treeFile)
//	file,_:=os.Create("NoNameText")
//	//buf := bytes.NewBuffer([]byte("Hero"))
//	dec := gob.NewEncoder(file)
//	dec.Encode("Hero")
//	//var i interface{}
//	var i string
//	i=""
//	buf := new(bytes.Buffer)
//	body,_:=os.ReadFile("NoNameText")
//	println(string(body))
//	buf= bytes.NewBuffer(body)
//	println(buf.String())
//	deca:= gob.NewDecoder(buf)
//	deca.Decode(&i)
//	//println(err.Error())
//	println(i)
//}

//func DecodeTreeFile(buffer *LinkedList) {
//
//	var i string
//	i=""
//	buf := new(bytes.Buffer)
//	body,_:=os.ReadFile(buffer.treeFile)
//	println(string(body))
//	buf= bytes.NewBuffer(body)
//	println(buf.String())
//	deca:= gob.NewDecoder(buf)
//	deca.Decode(&i)
//	//println(err.Error())
//	println(i)
//}

func DecodeTreeFile(buffer *LinkedList) {

	var i *btree.BTree
	i = new(btree.BTree)
	buf := new(bytes.Buffer)
	body, _ := os.ReadFile(buffer.treeFile)
	println(string(body))
	if len(body) == 0 {
		buffer.tree = btree.New(3)
	} else {
		buf = bytes.NewBuffer(body)
		println(buf.String())
		deca := gob.NewDecoder(buf)
		deca.Decode(i)
		//println(err.Error())
		buffer.tree = i
		println(i)
	}

}

//func AddNode
