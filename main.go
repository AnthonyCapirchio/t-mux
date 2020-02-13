package main

import (
  "fmt"
  "strings"
  "encoding/json"
)

const PathDelimiter = "/"
const VarPrefix = ":"
const Wildcard = "*"

type TreeNode struct {
  Handler interface{}
  Childs map[string]*TreeNode
  VarName string 
  Parent *TreeNode `json:"-"`
}

func NewTree() *TreeNode {
  return &TreeNode{
    Handler: nil,
    Childs: map[string]*TreeNode{},
  }
}

func (t *TreeNode) AddNode(path string, handler interface{}) {
  if strings.HasPrefix(path, PathDelimiter) {
    path = strings.TrimPrefix(path, PathDelimiter)
  }
  splitted := strings.Split(path, PathDelimiter)

  currentNode := t

  for i := 0; i < len(splitted); i++ {
    key := splitted[i]
    varName := ""

    if strings.HasPrefix(key, VarPrefix) {
      varName = strings.TrimPrefix(key, VarPrefix)
      key = Wildcard
    }

    if _, ok := currentNode.Childs[key]; !ok {
      currentNode.Childs[key] = &TreeNode{
        Handler: nil,
        Childs: map[string]*TreeNode{},
        Parent: currentNode,
      }
    }
    
    currentNode = currentNode.Childs[key]
    if varName != "" {
      currentNode.VarName = varName
    }

    if i == len(splitted)-1 {
      currentNode.Handler = handler
      return
    }
  }
}

func (t *TreeNode) Mount(path string, tree *TreeNode) {
  if strings.HasPrefix(path, PathDelimiter) {
    path = strings.TrimPrefix(path, PathDelimiter)
  }
  splitted := strings.Split(path, PathDelimiter)

  currentNode := t

  for i := 0; i < len(splitted); i++ {
    key := splitted[i]
    varName := ""

    if strings.HasPrefix(key, VarPrefix) {
      varName = strings.TrimPrefix(key, VarPrefix)
      key = Wildcard
    }

    if _, ok := currentNode.Childs[key]; !ok {
      currentNode.Childs[key] = &TreeNode{
        Handler: nil,
        Childs: map[string]*TreeNode{},
        Parent: currentNode,
      }
    }
    
    currentNode = currentNode.Childs[key]
    if varName != "" {
      currentNode.VarName = varName
    }

    if i == len(splitted)-1 {
      currentNode.Childs = tree.Childs
      return
    }
  }
}

func (t TreeNode) GetNode(path string) {
  if strings.HasPrefix(path, PathDelimiter) {
    path = strings.TrimPrefix(path, PathDelimiter)
  }
  splitted := strings.Split(path, PathDelimiter)
  params := map[string]string{}
  currentNode := &t

  for i := 0; i < len(splitted); i++ {
    //isWildcard := false
    key := splitted[i]

    if _, ok := currentNode.Childs[key]; !ok {
      if _, ok := currentNode.Childs[Wildcard]; !ok {
        fmt.Println("Not Found")
        return
      } else {
        //isWildcard = true
        params[currentNode.Childs[Wildcard].VarName] = key
        key = Wildcard
      }
    }

    if i == len(splitted)-1 {
      if currentNode.Childs[key].Handler == nil {
        fmt.Println("Not Found")
        return
      } else {
        fmt.Println("Handler: ", currentNode.Childs[key].Handler)
        fmt.Println("Params: ", params)
      }
      return
    }

    currentNode = currentNode.Childs[key]
  }
}

func (t *TreeNode) RemoveNode(path string) {
  if strings.HasPrefix(path, PathDelimiter) {
    path = strings.TrimPrefix(path, PathDelimiter)
  }
  splitted := strings.Split(path, PathDelimiter)

  currentNode := t

  for i := 0; i < len(splitted); i++ {
    key := splitted[i]

    if _, ok := currentNode.Childs[key]; !ok {
      fmt.Println("Not Found")
      return
    }

    if i == len(splitted)-1 {
      if len(currentNode.Childs[key].Childs) == 0 {
        delete(currentNode.Childs, key)
      } else {
        currentNode.Childs[key].Handler = nil
      }
      return
    }

    currentNode = currentNode.Childs[key]
  }
}

func (t *TreeNode) Flush() {
  b, _ := json.MarshalIndent(t, "", "  ")
  fmt.Println( string(b) )
}

func main() {

  handlerA := 17
  handlerB := 18
  handlerC := 19

  test := NewTree()

  test.AddNode("/a", handlerA)
  test.AddNode("/test/of/new/path", handlerB)
  test.AddNode("/test/of/new/:var", handlerC)
  test.GetNode("/test/of/new/test")
  test.GetNode("/test/of/new/plouf")
  test.GetNode("/test/of/new/path")
  test.GetNode("/test/of/new/haha/hoho")

  test2 := NewTree()

  test2.AddNode("/ping/1", 1)
  test2.AddNode("/ping/2", 2)
  test2.AddNode("/ping/3", 3)

  test.Mount("/:mounted", test2)

  test.GetNode("/fr/ping/2")
}