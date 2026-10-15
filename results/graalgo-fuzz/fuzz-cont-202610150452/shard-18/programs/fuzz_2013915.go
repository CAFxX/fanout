package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 int8 }

func main() {
    var v1 int16 = -32768
    var v2 int64 = -488
    var v3 float64 = ((0.1 / ((float64(0.0) - 0.1) * (0.0 - 0.1))) + 1e10)
    v4 := make(map[uint32]int)
    v5 := make(map[string]bool)
    v2 += v2
    v6 := GenS1{F0: 640, F1: 2147483648, F2: 330}
    fmt.Println(v2, v3)
    v3 = v3
    v7 := GenS1{F0: 857, F1: 4294967295, F2: 161}
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := GenS2{F0: -2147483648}
        v10 := GenS1{F0: 2, F1: 0, F2: 0}
        v2 = v2
        _ = fi8
        _ = v10
        _ = v9
    }
    v2 = v2
    switch ("a" + ("a" + "hello")) {
    case "\n\t":
        switch 0 {
        case -479:
            v11 := []int{}
            _ = v11
        }
    case "":
        var v12 int = 0
        _ = v12
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        fmt.Println(fi13)
        fmt.Println(v2, v3, v1)
        _ = fi13
    }
    v14 := GenS3{F0: 501, F1: -128}
    fmt.Println(v2, v1, v1)
    fmt.Println(v3)
    fmt.Println(v1, v3, v1)
    var v15 int64 = int64(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
