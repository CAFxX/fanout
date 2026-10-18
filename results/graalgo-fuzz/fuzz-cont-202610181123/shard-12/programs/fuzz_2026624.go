package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 string }

func main() {
    var v1 uint16 = 841
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    v4 := GenS1{F0: "hello", F1: 0, F2: "\n\t"}
    for fi5 := 0; fi5 < 13; fi5++ {
        v1 = (^(v1))
        var v6 uint32 = 1
        _ = fi5
        _ = v6
    }
    v7 := []int{4294967295, 2147483647, 286}
    v8 := GenS1{F0: ("hello" + ("a" + "")), F1: -211, F2: "a"}
    var v9 float64 = (0.1 - 1e10)
    fmt.Println(v1, v9)
    var v10 int8 = 127
    v9 += v9
    v11 := GenS1{F0: "\n\t", F1: -637, F2: ("" + "hello")}
    v7 = append(v7, -812)
    v10 = 1
    var v12 bool = true
    _ = len(v7)
    var v13 uint64 = 0
    switch "hello" {
    case "foo bar":
        var v14 bool = false
        _ = v14
    case "a":
        fmt.Println(v13)
    }
    switch "foo bar" {
    case "a":
        for fi15 := 0; fi15 < 19; fi15++ {
            v1 = v1
            v9 -= (0.1 * (1e10 - (-2.25 + -2.25)))
            var v16 uint = 677
            _ = fi15
            _ = v16
        }
    case "foo bar":
        var v17 int16 = -32768
        _ = v17
    }
    var v18 uint = 4294967295
    v10 = (v10 ^ v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
