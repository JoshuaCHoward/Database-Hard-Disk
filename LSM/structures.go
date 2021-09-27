package LSM

import (
	"bytes"
	"encoding/gob"
	btree "home/storable_btree"

	"encoding/json"
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

type Item struct {
	btree.Item
	Value struct {
		int
		string
	}
}

func (i Item) Less(than btree.Item) bool {
	x := than.(Item)
	if i.Value.int < x.Value.int {
		return true
	} else {
		return false
	}
}

//Type assertion are for stuff you know you have but won't know at run time
//Type conversion are stuff that's the same as an another type but the other type but have more stuff and can fit. Note that stuff currently being defined and setting the parameter there to be converted will not work since the interface contents are not actually currently there since they are being currently made

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
//
//func DecodeTreeFile(buffer *LinkedList) {
//
//	var i *btree.BTree
//	i=new(btree.BTree)
//	buf := new(bytes.Buffer)
//	body,_:=os.ReadFile(buffer.treeFile)
//	extra,_:=os.Open(buffer.treeFile)
//	println(string(body))
//	if (len(body)==0){
//		buffer.tree=btree.New(3)
//	} else {
//		println(string(body))
//
//		buf= bytes.NewBuffer(body)
//		println(buf.String())
//		deca:= gob.NewDecoder(extra)
//		deca.Decode(i)
//		//println(err.Error())
//		buffer.tree=i
//		println(i.Len())
//		println(i.Max())
//	}
//
//}
//
//func UpdateTreeFile(buffer *LinkedList) {
//
//
//	body,_:=os.Create(buffer.treeFile)
//	deca:= gob.NewEncoder(body)
//	deca.Encode(buffer.tree)
//	println("----------")
//	println(buffer.tree.Len())
//	body.Close()
//
//}

func DecodeTreeFile(buffer *LinkedList) {

	var i *btree.BTree
	i = new(btree.BTree)
	buf := new(bytes.Buffer)
	body, _ := os.ReadFile(buffer.treeFile)
	extra, _ := os.Open(buffer.treeFile)
	println(string(body))
	if len(body) == 0 {
		buffer.tree = btree.New(3)
	} else {
		println(string(body))

		buf = bytes.NewBuffer(body)
		println(buf.String())
		deca := gob.NewDecoder(extra)
		deca.Decode(i)
		//println(err.Error())
		buffer.tree = i
		println(i.Len())
		println(i.Max())
	}

}

func UpdateTreeFile(buffer *LinkedList) {
	body, _ := os.Create(buffer.treeFile)
	println("...............")
	println(buffer.tree.Len())
	treeJSON, _ := json.Marshal(buffer.tree)
	println(string(treeJSON))
	body.Write(treeJSON)
	println("----------")
	println(buffer.tree.Len())
	body.Close()

}

//func AddNode
