package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }

func main() {
    var v1 float64 = -2.25
    var v2 uint32 = 4294967295
    v3 := make(map[uint32]string)
    if ("a" == "\n\t") {
        var v4 int = -932
        v5 := v3[(v2 & (v2 ^ (v2 << ((4294967295) & 63))))]
        _ = v4
        _ = v5
    } else {
        var v6 uint16 = 2
        _ = v6
    }
    fmt.Println(v2, v1)
    var v7 float64 = float64(((3.14159 + 1.5) - 1e10))
    fmt.Println(v2)
    _ = len(v3)
    fmt.Println(v7, v7, v1)
    v1 = v7
    v8 := GenS1{F0: true, F1: -54}
    for fi9 := 0; fi9 < 6; fi9++ {
        fmt.Println(fi9, fi9)
        _ = fi9
    }
    if true {
        v10 := GenS1{F0: true, F1: -1}
        if true {
            v11 := GenS1{F0: (!(1 > -768)), F1: 573}
            v12 := GenS1{F0: (1 <= -9223372036854775808), F1: 1}
            var v13 string = ""
            _ = v11
            _ = v12
            _ = v13
        }
        _ = v10
    } else {
        v14 := []int{-688, 2147483647, 1, 9223372036854775807}
        _ = len(v14)
        _ = v14
    }
    switch 9223372036854775807 {
    case 543:
        var v15 uint64 = 42
        _ = v15
    }
    var v16 float32 = 1e10
    v17 := []int{-9223372036854775808, 9223372036854775807, 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
