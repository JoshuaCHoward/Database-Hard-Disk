package LSM

import (
	"encoding/json"
	"github.com/emirpasic/gods/trees/btree"
	"os"
	"strconv"
)

type LinkedList struct {
	length       int
	List         *Node
	fileLocation string
	treeFile     string
	tree         *btree.Tree
}

type Node struct {
	left  *Node
	right *Node
	value int
	spare interface{}
}

type Item struct {
	Value struct {
		Int    int
		String string
	}
}

//func (i Item) Less(than btree.Item) bool {
//	println("Bruh")
//	println(i.Value.Int)
//	if (than==nil){
//		return false
//	}
//	println("Rengoku")
//
//	x := than.(Item)
//	println("Rengoku")
//	if i.Value.Int < x.Value.Int {
//		return true
//	} else {
//		return false
//	}
//	return true
//}

//func (i Item) Less(than btree.Item) bool {
//	println("Bruh")
//	println(i.Value.Int)
//	println(than)
//	println("Rengoku")
//	if than==nil{
//		println("DUBMDUMB")
//		return false
//	}
//	x := than.(Item)
//	println("Rengoku")
//	if i.Value.Int < x.Value.Int {
//		return true
//	} else {
//		return false
//	}
//	return true
//}

//INSERT INTO users VALUES (1, 'Admin');
//func (i Item) Less(than btree.Item) bool {
//	println("+++++++++++++++++++")
//	print(than)
//	//var p btree.Item = (*Item)(nil)
//	x:=than.(*Item)
//	if n, ok := than.(*Item); ok {
//		fmt.Printf("n=%#v\n", n)
//		return i.Value.Int < n.Value.Int
//
//	}
//	return false
//	//x := than.(Item)
//}

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

	i := btree.NewWithIntComparator(3)
	body, _ := os.ReadFile(buffer.treeFile)
	if len(body) == 0 {
		buffer.tree = btree.NewWithIntComparator(3)
	} else {
		c := make(map[string]string)

		json.Unmarshal([]byte(`{"1":"repo\\users\\068e985f-0b6d-4311-8867-a148c9371a87","2":"repo\\users\\068e985f-0b6d-4311-8867-a148c9371a87","3":"repo\\users\\068e985f-0b6d-4311-8867-a148c9371a87"}`), &c)
		_ = FromJSON(i, body)
		buffer.tree = i

	}

}
func FromJSON(tree *btree.Tree, data []byte) error {
	elements := make(map[string]string)
	err := json.Unmarshal(data, &elements)
	if err == nil {
		tree.Clear()

		for key, value := range elements {
			intKey, _ := strconv.Atoi(key)
			tree.Put(intKey, value)
		}
	}
	return err
}
func UpdateTreeFile(buffer *LinkedList) {
	body, _ := os.Create(buffer.treeFile)
	println("...............")
	println(buffer.tree.String())
	treeJSON, _ := buffer.tree.ToJSON()
	//println(string(treeJSON))
	body.Write(treeJSON)
	println("----------")
	body.Close()

}

//func AddNode
