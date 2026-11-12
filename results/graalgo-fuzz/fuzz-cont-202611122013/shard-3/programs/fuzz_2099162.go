package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 bool }
type GenS2 struct { F0 int32; F1 int32 }
func genFn0() int {
    v4 := []int{177, -9223372036854775808}
    var v5 string = "\n\t"
    switch -907 {
    case -9223372036854775808:
        v4 = append(v4, 42)
    case 966:
        v6 := GenS1{F0: 570, F1: (v5 == v5), F2: true}
        _ = v6
    }
    var v7 uint64 = 0
    var v8 int8 = -128
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    return -375
}

func main() {
    var v1 int = -2147483648
    var v2 uint16 = 65535
    v3 := []uint32{65535, 0, 92, 65535}
    for fi9 := 0; fi9 < 13; fi9++ {
        fmt.Println(fi9, v1, fi9)
        _ = fi9
    }
    var v10 int64 = 174
    v11 := GenS1{F0: 2147483648, F1: (717 >= 818), F2: true}
    var v12 uint16 = v2
    var v13 int8 = -128
    _ = len(v3)
    var v14 int = (v1 - (v1 | (v1 & v1)))
    if true {
        v15 := GenS1{F0: uint32(v12), F1: true, F2: false}
        _ = v15
    } else {
        var v16 float32 = (1.5 + float32(0.1))
        _ = v16
    }
    fmt.Println(v13)
    if true {
        _ = len(v3)
        fmt.Println(v12, v10)
        if false {
            var v17 uint16 = v12
            _ = v17
        }
    } else {
        var v18 bool = true
        switch v14 {
        case 1:
            v19 := GenS1{F0: 578, F1: (v18 || true), F2: ((("" + "\n\t") > "\n\t") || true)}
            _ = v19
        case 0:
            if (v18 && (1000 != 1)) {
                v20 := GenS2{F0: 42, F1: -134}
                fmt.Println(v18)
                _ = v20
            }
        default:
            fmt.Println(v1, v2, v13)
        }
        var v21 uint = 42
        _ = v18
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
