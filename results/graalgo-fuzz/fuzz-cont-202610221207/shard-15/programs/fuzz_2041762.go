package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 bool; F2 int }
func genFn0(p0 string, p1 int, p2 uint8) int8 {
    p0 = "foo bar"
    var v7 bool = true
    p1 = p1
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 bool = v7
        v10 := GenS1{F0: 127}
        _ = fi8
        _ = v10
        _ = v9
    }
    p2 = 245
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 33
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 248
    v3 := make(map[int]string)
    var v4 uint = 2
    v5 := make(map[int]int)
    var v6 bool = (v4 >= v4)
    v11 := GenS1{F0: uint8(v4)}
    fmt.Println(v2)
    v12 := GenS2{F0: v6}
    if (v6 || ("a" <= "\n\t")) {
        v2 -= v2
        v4 -= (v4 + (v4 ^ 255))
    }
    var v13 int = 4294967295
    v14 := GenS1{F0: 0}
    v15 := GenS3{F0: 233, F1: v6, F2: (v13 ^ -661)}
    v2 = (^(v2))
    v16 := []int{}
    var v17 int = (v13 ^ v13)
    if (v6 && v6) {
        fmt.Println(v13, v17, v4)
        for fi18 := 0; fi18 < 14; fi18++ {
            for fi19 := 0; fi19 < 15; fi19++ {
                v20 := GenS3{F0: 127, F1: v6, F2: -285}
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
    }
    if true {
        delete(v5, 4294967295)
    } else {
        var v21 float64 = float64(0.1)
        v13 = (v13 << ((v4) & 63))
        _ = v21
    }
    var v22 uint64 = 42
    for fi23 := 0; fi23 < 7; fi23++ {
        _ = len(v1)
        _ = fi23
    }
    for fi24 := 0; fi24 < 15; fi24++ {
        if len(v16) > 0 {
            v16[0] = (v17 ^ fi24)
        }
        _ = fi24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
