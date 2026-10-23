package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 bool }
type GenS2 struct { F0 int; F1 int64; F2 int64 }
func genFn0(p0 bool, p1 int16) uint {
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := GenS1{F0: p1, F1: 359, F2: (!p0)}
        p1 = p1
        _ = fi5
        _ = v6
    }
    var v7 int16 = p1
    _ = p0
    _ = p1
    _ = v7
    return 697
}

func main() {
    var v1 bool = (223 > 79)
    v2 := make(map[int]int)
    v3 := make(map[int]bool)
    var v4 uint16 = 32767
    var v8 uint32 = 2147483647
    var v9 int = 9223372036854775807
    var v10 float64 = 1.5
    v10 = v10
    for fi11 := 0; fi11 < 9; fi11++ {
        v1 = (!true)
        var v12 uint = 42
        _ = fi11
        _ = v12
    }
    if v1 {
        var v13 uint16 = (v4 | v4)
        v9 -= v9
        _ = v13
    }
    var v14 uint64 = 42
    if (1 == 211) {
        fmt.Println(v9, v9)
    } else {
        delete(v2, v9)
        if v1 {
            var v15 bool = (!(-32768 <= 502))
            var v16 float64 = v10
            _ = v15
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        v18 := []int{-147}
        fmt.Println(v4, v1)
        _ = fi17
        _ = v18
    }
    switch v9 {
    case 0:
        v19 := GenS2{F0: v9, F1: 42, F2: 0}
        _ = v19
    case -392:
        for fi20 := 0; fi20 < 15; fi20++ {
            for fi21 := 0; fi21 < 6; fi21++ {
                v22 := v2[(fi20 + v9)]
                if v1 {
                    v4 -= (v4 >> ((630) & 63))
                    v8 += 0
                }
                _ = fi21
                _ = v22
            }
            var v23 int32 = int32(v14)
            _ = fi20
            _ = v23
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
