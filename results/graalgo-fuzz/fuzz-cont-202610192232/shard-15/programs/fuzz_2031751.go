package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int32 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS4 struct { F0 int; F1 bool; F2 bool }
type GenS5 struct { F0 bool }
func genFn0(p0 uint32) int {
    p0 = p0
    v4 := GenS1{F0: true, F1: 736, F2: -50}
    var v5 uint32 = (p0 + p0)
    var v6 uint64 = 1
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 585
}

func main() {
    var v1 uint64 = 520
    var v2 uint32 = 2
    var v3 float64 = (0.0 * 1e10)
    var v7 uint64 = v1
    var v8 string = "a"
    switch 4294967295 {
    case -950:
        v9 := GenS2{F0: (false || true), F1: 4294967295}
        _ = v9
    case 127:
        v10 := GenS2{F0: (!false), F1: v2}
        _ = v10
    case -9223372036854775808:
        v11 := []int{genFn0((v2 / ((4294967295) | 1))), 947}
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        fmt.Println(v2)
        v7 += 1
        var v13 int8 = -1
        _ = fi12
        _ = v13
    }
    v14 := GenS3{F0: -128, F1: v7, F2: (^(v1))}
    if false {
        v3 = v3
    }
    var v15 string = v8
    var v16 bool = false
    var v17 uint16 = 32767
    v7 = v7
    if false {
        for fi18 := 0; fi18 < 17; fi18++ {
            var v19 int64 = 0
            _ = fi18
            _ = v19
        }
        if v16 {
            for fi20 := 0; fi20 < 4; fi20++ {
                v8 = (v15 + v15)
                _ = fi20
            }
        } else {
            var v21 int8 = 42
            for fi22 := 0; fi22 < 18; fi22++ {
                switch genFn0(v2) {
                case 4294967295:
                    var v23 int16 = 0
                    _ = v23
                default:
                    for fi24 := 0; fi24 < 4; fi24++ {
                        var v25 int64 = -541
                        fmt.Println(v1, v3, v15)
                        v8 = v8
                        _ = fi24
                        _ = v25
                    }
                }
                v8 = "hello"
                _ = fi22
            }
            var v26 uint16 = 654
            _ = v21
            _ = v26
        }
    }
    var v27 float64 = 3.14159
    for fi28 := 0; fi28 < 19; fi28++ {
        v29 := GenS4{F0: fi28, F1: v16, F2: v16}
        v30 := []int{genFn0(v2)}
        if v16 {
            var v31 int32 = 490
            _ = v31
        }
        _ = fi28
        _ = v29
        _ = v30
    }
    if v16 {
        var v32 uint64 = (v7 << ((255) & 63))
        for fi33 := 0; fi33 < 9; fi33++ {
            var v34 int = fi33
            _ = fi33
            _ = v34
        }
        _ = v32
    } else {
        v7 -= (v7 >> ((9) & 63))
        v3 += (((v27 / v27) / v3) + 1.5)
        for fi35 := 0; fi35 < 16; fi35++ {
            if v16 {
                for fi36 := 0; fi36 < 7; fi36++ {
                    var v37 uint32 = v2
                    _ = fi36
                    _ = v37
                }
                if (!v16) {
                    v38 := []int{fi35, genFn0(915), fi35, (fi35 * fi35)}
                    v3 = v3
                    for fi39 := 0; fi39 < 12; fi39++ {
                        var v40 uint16 = v17
                        _ = fi39
                        _ = v40
                    }
                    _ = v38
                }
            }
            v41 := GenS5{F0: v16}
            _ = fi35
            _ = v41
        }
    }
    if v16 {
        v3 = (0.1 - -2.25)
        fmt.Println(v27)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v7
    _ = v8
}
