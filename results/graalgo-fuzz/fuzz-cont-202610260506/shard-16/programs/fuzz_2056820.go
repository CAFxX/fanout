package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 int8 }

func main() {
    var v1 float64 = 0.1
    var v2 float64 = v1
    v3 := []int{}
    v2 = (1.5 + v2)
    var v4 string = "\n\t"
    if len(v3) > 0 {
        v3[0] = 32767
    }
    v5 := GenS1{F0: (968 < 370)}
    var v6 bool = true
    v4 = "foo bar"
    switch "foo bar" {
    case "foo bar":
        var v7 string = v4
        _ = v7
    case "\n\t":
        if len(v3) > 0 {
            v3[0] = 32767
        }
    case "a":
        v8 := GenS2{F0: 42, F1: 100}
        _ = v8
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
