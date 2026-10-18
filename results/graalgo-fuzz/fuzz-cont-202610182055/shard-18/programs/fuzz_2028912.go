package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int64 }
type GenS2 struct { F0 int8; F1 int; F2 uint16 }
type GenS3 struct { F0 uint32; F1 int; F2 int16 }
type GenS4 struct { F0 string; F1 uint16; F2 int8 }
func genFn0(p0 uint64, p1 bool, p2 string) int32 {
    v6 := GenS1{F0: p0, F1: -519, F2: 1}
    if p1 {
        for fi7 := 0; fi7 < 14; fi7++ {
            for fi8 := 0; fi8 < 13; fi8++ {
                v9 := GenS2{F0: -1, F1: 4294967295, F2: 570}
                var v10 uint = 2
                _ = fi8
                _ = v10
                _ = v9
            }
            fmt.Println(p0, p1)
            _ = fi7
        }
        v11 := []int{127, 545, 886}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 689
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = (-2147483648 <= -9223372036854775808)
    v3 := make(map[int]string)
    v4 := []uint32{}
    var v5 int16 = 42
    delete(v1, 503)
    if len(v4) > 0 {
        v4[0] = 1
    }
    v12 := GenS3{F0: 42, F1: 301, F2: (v5 + -367)}
    var v13 int8 = 1
    fmt.Println(v5)
    for fi14 := 0; fi14 < 18; fi14++ {
        delete(v1, fi14)
        _ = fi14
    }
    var v15 int16 = v5
    v16 := GenS4{F0: "foo bar", F1: 32767, F2: v13}
    var v17 bool = (!(32767 >= 1))
    v17 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
