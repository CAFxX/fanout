package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint32 }
func genFn0(p0 uint8) int {
    fmt.Println(p0)
    v6 := GenS1{F0: 1, F1: 0}
    v7 := GenS1{F0: -77, F1: 1000}
    _ = p0
    _ = v6
    _ = v7
    return 553
}
func genFn1(p0 string) int8 {
    var v8 int = 9223372036854775807
    p0 = p0
    _ = p0
    _ = v8
    return 0
}

func main() {
    var v1 uint = 255
    var v2 uint = 42
    v3 := make(map[int]string)
    v4 := make(map[int]bool)
    v5 := []int{4294967295, 42, -9223372036854775808}
    for fi9 := 0; fi9 < 17; fi9++ {
        fmt.Println(v1, v1)
        fmt.Println(v1, fi9, v2)
        var v10 uint32 = 459
        _ = fi9
        _ = v10
    }
    v11 := v3[275]
    v2 = 1
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 int8 = int8(fi12)
        var v14 bool = (true && ((v13 != (^(v13))) && true))
        _ = fi12
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: 253, F1: -658}
    if len(v5) > 0 {
        v5[0] = -9223372036854775808
    }
    v11 = (v11 + ("hello" + v11))
    switch v11 {
    case "\n\t":
        var v16 bool = (494 == 0)
        _ = v16
    }
    if false {
        v2 -= v1
        var v17 int64 = -1
        if false {
            v18 := GenS2{F0: true, F1: 65535, F2: 966}
            _ = v18
        }
        _ = v17
    } else {
        var v19 bool = false
        _ = v19
    }
    fmt.Println(v11, v11, v11)
    var v20 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
