package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64 }
func genFn0(p0 uint32, p1 uint32, p2 uint64) uint16 {
    v5 := []int{-296, 1, -748, -577}
    var v6 uint8 = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 2
}

func main() {
    var v1 float64 = 0.0
    var v2 uint = 712
    var v3 int16 = -1
    var v4 uint16 = 478
    switch "foo bar" {
    case "hello":
        v7 := GenS1{F0: -35, F1: 1, F2: (^(v2))}
        _ = v7
    default:
        v8 := GenS2{F0: 857}
        _ = v8
    }
    var v9 int16 = (v3 >> (((v2 >> (((v2 >> ((v2) & 63))) & 63))) & 63))
    v10 := []int{128, 128}
    for fi11 := 0; fi11 < 6; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            var v13 uint64 = 42
            v14 := GenS3{F0: v4, F1: 94}
            _ = fi12
            _ = v13
            _ = v14
        }
        v3 = 137
        if true {
            v15 := GenS1{F0: 0, F1: 42, F2: 0}
            v10 = append(v10, fi11)
            var v16 string = "\n\t"
            _ = v15
            _ = v16
        }
        _ = fi11
    }
    if ((true || true) && false) {
        switch "foo bar" {
        case "hello":
            if true {
                var v17 string = "hello"
                _ = v17
            }
        default:
            var v18 uint32 = 65535
            _ = v18
        }
        if (!true) {
            v19 := GenS4{F0: (^(v3))}
            _ = v19
        } else {
            fmt.Println(v4, v3)
            v20 := GenS3{F0: (v4 % ((v4) | 1)), F1: int8(v9)}
            _ = v20
        }
        for fi21 := 0; fi21 < 19; fi21++ {
            v22 := GenS5{F0: 42}
            var v23 int16 = (-(v3))
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    v3 -= v9
    for fi24 := 0; fi24 < 10; fi24++ {
        var v25 bool = (uint64(v2) != 4294967295)
        fmt.Println(v2, v9, v9)
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        v1 = (float64(1e10) - (float64((1e10 / 0.1)) / v1))
        for fi27 := 0; fi27 < 17; fi27++ {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            var v28 int8 = 100
            _ = fi27
            _ = v28
        }
        var v29 int64 = 1
        _ = fi26
        _ = v29
    }
    if true {
        if true {
            if false {
                fmt.Println(v3)
            }
            var v30 uint = 255
            v31 := GenS1{F0: 7, F1: -75, F2: v2}
            _ = v30
            _ = v31
        }
        _ = len(v10)
        if (false || true) {
            for fi32 := 0; fi32 < 5; fi32++ {
                v3 = 42
                var v33 bool = true
                for fi34 := 0; fi34 < 18; fi34++ {
                    v9 -= (v9 ^ (v9 & v3))
                    _ = fi34
                }
                _ = fi32
                _ = v33
            }
        } else {
            v10 = append(v10, -551)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
