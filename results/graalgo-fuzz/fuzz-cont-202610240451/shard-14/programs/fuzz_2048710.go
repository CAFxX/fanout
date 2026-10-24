package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int64 }
func genFn0(p0 uint) int64 {
    fmt.Println(p0)
    for fi6 := 0; fi6 < 11; fi6++ {
        switch (fi6 - -9223372036854775808) {
        case -1:
            fmt.Println(p0, fi6)
        }
        var v7 string = ""
        _ = fi6
        _ = v7
    }
    _ = p0
    return 1
}
func genFn1() string {
    var v8 uint64 = 2
    v8 -= (^(v8))
    v8 += 18446744073709551615
    _ = v8
    return ("" + "a")
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    v3 := make(map[int]bool)
    v4 := make(map[string]int)
    v5 := []uint32{42, 2147483647, 2, 4294967295}
    var v9 uint8 = 128
    v10 := GenS1{F0: 318}
    var v11 int32 = 1
    switch -459 {
    case 128:
        v9 = 82
    case 831:
        v4["hello"] = 1
    case 4294967295:
        var v12 float64 = ((float64(0.1) * (3.14159 / 0.1)) * 1.5)
        _ = v12
    default:
        fmt.Println(v9, v11)
    }
    v3[-433] = true
    _ = len(v4)
    var v13 int = 209
    v13 = v13
    v14 := GenS2{F0: 4294967295, F1: -9223372036854775808}
    v15 := GenS1{F0: (v13 + v13)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
