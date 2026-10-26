package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }

func main() {
    v1 := []string{("foo bar" + "\n\t"), (("foo bar" + "") + ""), "hello", ("hello" + "hello")}
    var v2 string = ("" + "\n\t")
    var v3 uint = 42
    var v4 uint16 = 0
    var v5 int16 = 1
    v1 = append(v1, v2)
    var v6 string = ("hello" + "a")
    v7 := GenS1{F0: 2, F1: v2}
    fmt.Println(v5)
    switch "hello" {
    case "hello":
        v2 = v2
    }
    v1 = append(v1, ("a" + "foo bar"))
    var v8 float32 = 0.1
    v4 -= 42
    if len(v1) > 0 {
        v1[0] = v6
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
