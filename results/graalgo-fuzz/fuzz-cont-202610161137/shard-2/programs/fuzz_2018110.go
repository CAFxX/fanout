package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint; F1 uint8; F2 int32 }
type GenS3 struct { F0 bool; F1 bool; F2 int8 }
func genFn0(p0 string, p1 uint64) int {
    v6 := []int{}
    switch ("hello" + p0) {
    case "\n\t":
        var v7 int8 = 42
        _ = v7
    case "hello":
        v8 := GenS1{F0: 2, F1: 731}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return 127
}

func main() {
    v1 := make(map[string]int)
    var v2 int64 = 702
    v3 := make(map[string]bool)
    var v4 uint = 0
    var v5 int64 = (v2 + (v2 << ((2) & 63)))
    v4 += v4
    v9 := []int{-2147483648, 48}
    v4 = v4
    v10 := GenS2{F0: v4, F1: 128, F2: 886}
    for fi11 := 0; fi11 < 6; fi11++ {
        fmt.Println(v5, v4)
        _ = fi11
    }
    v12 := GenS2{F0: (v4 | (v4 ^ v4)), F1: 42, F2: 0}
    var v13 float64 = float64(1.5)
    for fi14 := 0; fi14 < 13; fi14++ {
        fmt.Println(v2, v2, v13)
        _ = fi14
    }
    fmt.Println(v13, v5, v2)
    v15 := GenS3{F0: (true && false), F1: true, F2: -1}
    var v16 int8 = 127
    var v17 float64 = v13
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
