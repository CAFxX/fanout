package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 string }
type GenS2 struct { F0 int8; F1 int64; F2 string }
func genFn0(p0 bool, p1 int16) uint32 {
    for fi4 := 0; fi4 < 11; fi4++ {
        var v5 uint32 = 600
        var v6 uint = 2
        _ = fi4
        _ = v5
        _ = v6
    }
    if p0 {
        var v7 uint = 1
        _ = v7
    } else {
        var v8 int64 = 1
        _ = v8
    }
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 bool = false
    var v2 int64 = 541
    var v3 uint32 = 1
    v9 := GenS1{F0: 1000, F1: 128, F2: ""}
    fmt.Println(v2, v2, v3)
    fmt.Println(v2)
    if (false || v1) {
        var v10 uint64 = 734
        _ = v10
    } else {
        v11 := GenS1{F0: -59, F1: 582, F2: ("hello" + "hello")}
        _ = v11
    }
    v12 := []int{128, 128}
    for fi13 := 0; fi13 < 5; fi13++ {
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        var v14 int16 = -32768
        _ = fi13
        _ = v14
    }
    _ = len(v12)
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    v1 = ((2 > 60) && (!(-128 < 100)))
    if (!(-1 != 0)) {
        fmt.Println(v1, v1, v2)
        v15 := GenS2{F0: -49, F1: (^(v2)), F2: "hello"}
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v9
}
