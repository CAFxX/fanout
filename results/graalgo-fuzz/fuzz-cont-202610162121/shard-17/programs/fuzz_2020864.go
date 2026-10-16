package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 int; F2 bool }
type GenS4 struct { F0 uint16; F1 bool; F2 uint8 }
func genFn0(p0 bool, p1 string) int32 {
    p0 = p0
    for fi4 := 0; fi4 < 16; fi4++ {
        v5 := GenS1{F0: 89, F1: -856, F2: 0}
        _ = fi4
        _ = v5
    }
    var v6 bool = false
    _ = p0
    _ = p1
    _ = v6
    return -1
}

func main() {
    var v1 uint16 = 2
    var v2 uint32 = 65535
    var v3 int64 = 0
    var v7 int = -9223372036854775808
    var v8 int64 = v3
    for fi9 := 0; fi9 < 16; fi9++ {
        v7 -= (v7 + (v7 >> ((1) & 63)))
        v3 += 9223372036854775807
        v10 := GenS2{F0: 127}
        _ = fi9
        _ = v10
    }
    fmt.Println(v7, v1)
    v11 := []int{v7, v7, (v7 << ((766) & 63)), (v7 >> ((42) & 63))}
    v1 = 229
    fmt.Println(v1, v8)
    var v12 uint8 = 13
    v8 += 9223372036854775807
    v13 := GenS3{F0: (v2 >> ((235) & 63)), F1: v7, F2: (1.5 == float64((float32(0.0) / 1.5)))}
    switch "hello" {
    case "a":
        fmt.Println(v1, v2)
    }
    v14 := GenS4{F0: v1, F1: false, F2: 42}
    var v15 uint16 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
