package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 uint64 }

func main() {
    v1 := []uint32{}
    v2 := []uint32{65535}
    var v3 float64 = (-2.25 * 1.5)
    v4 := []string{"foo bar", "hello"}
    v5 := []int{}
    switch "a" {
    case "":
        if false {
            var v6 int = 1
            var v7 uint = 65535
            var v8 float64 = v3
            _ = v6
            _ = v7
            _ = v8
        }
    case "a":
        fmt.Println(v3)
    case "foo bar":
        switch "a" {
        case "a":
            v9 := GenS1{F0: 1, F1: ("foo bar" + ("a" + "foo bar"))}
            _ = v9
        }
    default:
        var v10 int32 = 0
        _ = v10
    }
    fmt.Println(v3)
    v3 = (0.0 * (v3 + (v3 / v3)))
    v11 := GenS1{F0: -896, F1: "hello"}
    fmt.Println(v3)
    var v12 float64 = (v3 - 0.0)
    fmt.Println(v12)
    v3 = (0.1 + 1e10)
    for fi13 := 0; fi13 < 13; fi13++ {
        v14 := GenS2{F0: 100}
        v1 = append(v1, 2147483647)
        v15 := GenS3{F0: 1, F1: 0}
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
