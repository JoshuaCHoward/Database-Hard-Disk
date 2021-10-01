package LSM

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/google/uuid"
	gosql "home"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

// memoryCell is the underlying storage for the in-memory backend
// implementation. Each supported datatype can be mapped to and from
// this byte array.

//func NewMemoryBackend() *MemoryBackend {
//	return &MemoryBackend{
//		tables: map[string]*table{},
//	}
//}
type table struct {
	name    string
	columns []string
}

type MemoryBackend struct {
	tables       map[string]string
	tableBuffers map[string]*LinkedList
	bufferSize   int
	gosql.EmptyBackend
}

type DatabaseItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (eb *MemoryBackend) CreateTable(crt *gosql.CreateTableStatement) error {
	newPath := filepath.Join(".", "repo")
	_ = os.MkdirAll(newPath, os.ModePerm)
	//newpath = filepath.Join(".","repo" ,fmt.Sprintf("%s.txt",crt.Name.Value))
	newPath = filepath.Join(".", "repo", crt.Name.Value)
	_ = os.MkdirAll(newPath, os.ModePerm)
	newPath = filepath.Join(".", "repo", crt.Name.Value, fmt.Sprintf("%s.tree", crt.Name.Value))
	os.Create(newPath)
	eb.tables[crt.Name.Value] = filepath.Join(".", "repo", crt.Name.Value)
	return nil
}

func (eb *MemoryBackend) Select(crt *gosql.SelectStatement) (*gosql.Results, error) {

	println(crt.From.Value)
	table := (crt.From.Value)
	var text []gosql.ResultColumn
	for _, actualValue := range *crt.Item {
		keyValue, _ := strconv.Atoi(actualValue.Exp.Literal.Value)
		value, _ := eb.tableBuffers[table].tree.Get(keyValue)
		stringValue := fmt.Sprintf("%v", value)
		file, _ := os.Open(stringValue)
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			pageLine := scanner.Text()[0:1]
			if pageLine == strconv.Itoa(keyValue) {
				result := gosql.ResultColumn{
					Name: scanner.Text()[2:],
				}
				text = append(text, result)

				println(result.Name)

			}
		}
	}

	return &gosql.Results{Columns: text}, nil
}

//SELECT 1 FROM users
//func (eb *MemoryBackend) Sort(jsonArray []DatabaseItem){
//	for i, json := range jsonArray{
//		println(json.Key)
//		if (json.Key==jsonArray[len(jsonArray)-1].Key){
//			jsonArray[i]=jsonArray[len(jsonArray)-1]
//			jsonArray=jsonArray[:len(jsonArray)-1]
//			break
//		if (jsonArray[len(jsonArray)-1].Key>json.Key){
//			continue
//		} else {
//			jsonA
//
//		}
//
//		}
//	}
//}

//func (eb *MemoryBackend) Insert(crt *gosql.InsertStatement) error {
//	currentBuffer:=  eb.tableBuffers[crt.Table.Value]
//	println((*crt.Values)[0].Literal.Value)
//	currentItem := DatabaseItem{(*crt.Values)[0].Literal.Value,
//		(*crt.Values)[1].Literal.Value,
//		}
//	//stringifiedItem, _ :=json.Marshal(currentItem)
//	//fmt.Println(len(stringifiedItem))
//	eb.tableBuffers[crt.Table.Value]=append(currentBuffer,currentItem)
//	eb.Sort(currentBuffer)
//	//println(currentBuffer)
//	//println(currentBuffer.Cap())
//	//if (currentBuffer.Len()+len(stringifiedItem)<currentBuffer.Cap()){
//	//	//currentBuffer.Write([]byte(fmt.Sprintf("%s\n",stringifiedItem)))
//	//} else {
//	//	//sortedStrings := []string{}
//	//	//while true{}
//	//	//bufferedString,err:=currentBuffer.ReadString(byte('\n'))
//	//	//fmt.Println(currentBuffer.ReadString(byte('\n')),"YEah")
//	//	//currentBuffer.Reset()
//	//
//	//}
//
//	return nil
//
//}
//
//
//func NewMemoryBackend() *MemoryBackend {
//	files, _ := ioutil.ReadDir("./repo")
//	backend := &MemoryBackend{
//		tables:       map[string]*table{},
//		tableBuffers: map[string][]DatabaseItem{},
//		bufferSize: 200,
//	}
//	for _, f := range files {
//		backend.tableBuffers[f.Name()]=[]DatabaseItem{}
//
//	}
//	return backend
//}

func Append(buffer *LinkedList, item DatabaseItem) *LinkedList {
	num, _ := strconv.Atoi(item.Key)
	newNode := &Node{
		value: num,
		spare: item.Value,
	}
	bufferNode := buffer.List
	for true {

		number, _ := strconv.Atoi(item.Key)
		if number == bufferNode.value {
			bufferNode.value = number
			buffer.length -= 1
			break
		}
		if bufferNode.right == nil {
			bufferNode.right = newNode
			newNode.left = bufferNode
			break
		}
		if number < bufferNode.right.value {
			bufferNode.right.left = newNode
			newNode.right = bufferNode.right
			newNode.left = bufferNode
			bufferNode.right = newNode
			break
		}
		bufferNode = bufferNode.right

	}
	buffer.length += 1
	return nil
}

func (eb *MemoryBackend) Insert(crt *gosql.InsertStatement) error {
	currentBuffer := eb.tableBuffers[crt.Table.Value]
	println((*crt.Values)[0].Literal.Value)
	currentItem := DatabaseItem{(*crt.Values)[0].Literal.Value,
		(*crt.Values)[1].Literal.Value,
	}
	//stringifiedItem, _ :=json.Marshal(currentItem)
	//fmt.Println(len(stringifiedItem))
	Append(currentBuffer, currentItem)
	if eb.bufferSize <= currentBuffer.length {
		Write(currentBuffer)
	}
	//Print(currentBuffer)
	//println(currentBuffer.length)
	//eb.Sort(currentBuffer)
	//println(currentBuffer)
	//println(currentBuffer.Cap())
	//if (currentBuffer.Len()+len(stringifiedItem)<currentBuffer.Cap()){
	//	//currentBuffer.Write([]byte(fmt.Sprintf("%s\n",stringifiedItem)))
	//} else {
	//	//sortedStrings := []string{}
	//	//while true{}
	//	//bufferedString,err:=currentBuffer.ReadString(byte('\n'))
	//	//fmt.Println(currentBuffer.ReadString(byte('\n')),"YEah")
	//	//currentBuffer.Reset()
	//
	//}
	return nil
}

func Write(buffer *LinkedList) {
	bytesBuffer := bytes.Buffer{}
	bufferNode := buffer.List.right
	filePath := filepath.Join(buffer.fileLocation, uuid.NewString())
	for true {
		bytesBuffer.WriteString(fmt.Sprintf("%d %v\n", bufferNode.value, bufferNode.spare))
		treeItem := Item{
			Value: struct {
				Int    int
				String string
			}{bufferNode.value, filePath},
		}
		println(treeItem.Value.Int)
		buffer.tree.Put(treeItem.Value.Int, treeItem.Value.String)
		if bufferNode.right == nil {
			break
		}
		bufferNode = bufferNode.right
	}

	newFile, _ := os.Create(filePath)
	bytesBuffer.WriteTo(newFile)
	buffer.length = 0
	newFile.Close()
	UpdateTreeFile(buffer)
	bytesBuffer.Reset()
	buffer.List.right = nil
}

func Print(buffer *LinkedList) {
	bufferNode := buffer.List
	println("------------")
	for true {
		println(bufferNode.value)
		if bufferNode.right == nil {
			break
		}
		bufferNode = bufferNode.right
	}
	return
}

func NewMemoryBackend() *MemoryBackend {
	files, _ := ioutil.ReadDir("./repo")
	backend := &MemoryBackend{
		tables:       map[string]string{},
		tableBuffers: map[string]*LinkedList{},
		bufferSize:   3,
	}
	for _, f := range files {
		backend.tableBuffers[f.Name()] = &LinkedList{List: &Node{
			left:  nil,
			right: nil,
			value: 0,
		},
			fileLocation: filepath.Join(".", "repo", f.Name()),
			treeFile:     filepath.Join(".", "repo", f.Name(), fmt.Sprintf("%s.tree", f.Name())),
			tree:         nil,
		}
		backend.tables[f.Name()] = filepath.Join(".", "repo", f.Name())
		DecodeTreeFile(backend.tableBuffers[f.Name()])

	}
	return backend
}
