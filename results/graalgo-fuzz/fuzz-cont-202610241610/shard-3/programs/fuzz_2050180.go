package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 int16; F2 int32 }
type GenS4 struct { F0 uint64; F1 int16; F2 bool }
func genFn0(p0 int32) int {
    v6 := []int{127, -9223372036854775808, 2147483647}
    if (p0 == (p0 >> ((2) & 63))) {
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := GenS1{F0: ((42 == 42) || (2 <= 32767))}
            p0 = 0
            var v9 int = (fi7 << ((42) & 63))
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    fmt.Println(p0)
    v6 = append(v6, -2147483648)
    _ = p0
    _ = v6
    return -9223372036854775808
}
func genFn1() int {
    switch "\n\t" {
    case "foo bar":
        v10 := []int{32768}
        _ = v10
    case "a":
        if true {
            var v11 uint64 = 2
            v11 = 42
            switch "a" {
            case "\n\t":
                v12 := GenS1{F0: true}
                _ = v12
            default:
                v13 := []int{4294967295, -363}
                _ = v13
            }
            _ = v11
        }
    case "\n\t":
        v14 := []int{301}
        _ = v14
    default:
        var v15 int8 = -30
        _ = v15
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        fmt.Println(fi16)
        _ = fi16
    }
    v17 := GenS1{F0: (!(!(true && false)))}
    _ = v17
    return 621
}

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = 0.0
    var v3 int8 = -11
    var v4 float32 = -2.25
    v5 := []uint32{2147483647, 42}
    v18 := GenS1{F0: ((false && (!(1 <= 251))) && (-395 < 0))}
    v2 += (v2 * (0.0 + -2.25))
    v19 := GenS2{F0: -9223372036854775808}
    v3 = -101
    for fi20 := 0; fi20 < 4; fi20++ {
        _ = len(v1)
        var v21 int8 = v3
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        v23 := GenS2{F0: -927}
        v3 = (v3 / (((v3 | v3)) | 1))
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 4; fi24++ {
        for fi25 := 0; fi25 < 9; fi25++ {
            v3 = v3
            _ = fi25
        }
        for fi26 := 0; fi26 < 18; fi26++ {
            if true {
                var v27 int64 = -9223372036854775808
                fmt.Println(fi26, v2)
                _ = v27
            } else {
                v28 := GenS1{F0: (3.14159 > (v2 - float64(((0.1 - 3.14159) * v4))))}
                var v29 uint64 = 275
                v30 := GenS2{F0: -1}
                v4 -= (0.0 * v4)
                var v31 int8 = 100
                _ = v28
                _ = v29
                _ = v30
                _ = v31
            }
            v1[(fi24 % ((fi24) | 1))] = (!((v2 > v2) || false))
            _ = fi26
        }
        fmt.Println(v4, v4, v2)
        _ = fi24
    }
    v32 := GenS3{F0: 829, F1: 1000, F2: -1}
    switch -416 {
    case 32768:
        var v33 uint64 = 4294967295
        _ = v33
    }
    if false {
        v34 := GenS4{F0: 754, F1: -32768, F2: (568 <= 42)}
        _ = v34
    } else {
        if true {
            var v35 uint = 255
            for fi36 := 0; fi36 < 19; fi36++ {
                var v37 float32 = 1.5
                var v38 int32 = 232
                _ = fi36
                _ = v37
                _ = v38
            }
            v39 := GenS2{F0: 380}
            _ = v35
            _ = v39
        } else {
            switch 42 {
            case 127:
                var v40 int16 = 679
                _ = v40
            case 0:
                v41 := GenS2{F0: -1}
                _ = v41
            }
            v42 := v1[-1]
            _ = v42
        }
        for fi43 := 0; fi43 < 20; fi43++ {
            var v44 int8 = v3
            v3 = 76
            var v45 uint32 = 1
            _ = fi43
            _ = v44
            _ = v45
        }
        for fi46 := 0; fi46 < 13; fi46++ {
            v4 += float32(((1.5 / -2.25) + v2))
            _ = fi46
        }
    }
    v4 = (((((v4 / (3.14159 * v4)) / v4) / v4) - v4) - float32((float64(((v4 * v4) * v4)) + float64(v4))))
    if false {
        if (-453 < 378) {
            for fi47 := 0; fi47 < 19; fi47++ {
                for fi48 := 0; fi48 < 20; fi48++ {
                    v2 -= 1.5
                    _ = fi48
                }
                _ = fi47
            }
        }
        v1[997] = (((0.0 * ((float32(float64(v4)) + (float32(float64(v4)) - float32((1.5 / (v2 * 0.1))))) + v4)) * v4) > v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
