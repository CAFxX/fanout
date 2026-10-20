package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }

func main() {
    v1 := []uint32{2, 0, 1}
    var v2 int64 = 976
    var v3 uint16 = 1
    v4 := []string{("" + "")}
    var v5 bool = false
    v5 = ((818 <= 850) && (v5 && (30 < 1)))
    fmt.Println(v2)
    v2 += v2
    if v5 {
        var v6 int32 = -2147483648
        _ = v6
    } else {
        v7 := GenS1{F0: 42, F1: 1}
        var v8 uint = 2
        _ = v7
        _ = v8
    }
    if v5 {
        fmt.Println(v5, v2)
    } else {
        _ = len(v4)
        v3 -= 65535
        v5 = false
    }
    v1 = append(v1, 0)
    v1 = append(v1, 289)
    v5 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
