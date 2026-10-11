package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]uint32)
    var v4 int = 766
    v5 := []uint32{2, 2, 1}
    v6 := make(map[int]uint32)
    var v7 bool = (!(42 < 1))
    fmt.Println(v7, v4)
    v8 := GenS1{F0: 42}
    fmt.Println(v7)
    if true {
        v7 = (42 <= 9223372036854775807)
    } else {
        v7 = (1 < 42)
    }
    v7 = (!(-129 < v4))
    v7 = v7
    for fi9 := 0; fi9 < 5; fi9++ {
        if (1 > 2) {
            var v10 uint64 = 998
            v11 := GenS1{F0: 1}
            v4 += (^(fi9))
            _ = v10
            _ = v11
        }
        _ = fi9
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        fmt.Println(fi12, v4)
        fmt.Println(v4, v4, fi12)
        var v13 int32 = 650
        _ = fi12
        _ = v13
    }
    if v7 {
        v7 = v7
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 uint = 723
            var v16 uint64 = 1
            var v17 float64 = 0.0
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
    }
    v7 = (v7 && (65535 == 4294967295))
    v7 = ("" <= "")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
