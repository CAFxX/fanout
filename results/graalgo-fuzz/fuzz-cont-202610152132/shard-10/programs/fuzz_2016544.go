package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint64 }
type GenS2 struct { F0 int; F1 int32 }
func genFn0(p0 uint, p1 uint8, p2 int32) uint32 {
    v4 := []int{32767, -2147483648}
    fmt.Println(p0, p0, p0)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(p1, p0, p1)
    var v5 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 2
}

func main() {
    var v1 uint8 = 93
    var v2 uint32 = 4294967295
    v3 := []int{}
    var v6 string = ""
    v7 := GenS1{F0: 0, F1: -820, F2: 42}
    if true {
        v6 = ""
        if len(v3) > 0 {
            v3[0] = 42
        }
    }
    v6 = v6
    var v8 int8 = 100
    var v9 uint16 = 946
    switch (v6 + "a") {
    case "\n\t":
        var v10 uint64 = 4294967295
        _ = v10
    case "":
        v9 -= v9
    default:
        v11 := GenS2{F0: 360, F1: 543}
        _ = v11
    }
    v8 = v8
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
