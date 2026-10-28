package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16; F1 int16; F2 int64 }
type GenS6 struct { F0 uint; F1 int; F2 int16 }
type GenS7 struct { F0 uint; F1 int32 }
func genFn0(p0 uint16) uint32 {
    var v5 int8 = 127
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := []int{fi6}
        var v8 int8 = v5
        if true {
            var v9 int8 = (v5 >> ((4294967295) & 63))
            for fi10 := 0; fi10 < 8; fi10++ {
                p0 -= (p0 & 748)
                var v11 uint16 = p0
                fmt.Println(fi10, v5)
                _ = fi10
                _ = v11
            }
            if (true || (p0 >= (p0 * p0))) {
                p0 += p0
            }
            _ = v9
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v5
    return 381
}
func genFn1(p0 uint64, p1 int, p2 int32) uint8 {
    for fi12 := 0; fi12 < 8; fi12++ {
        for fi13 := 0; fi13 < 7; fi13++ {
            if true {
                var v14 uint = 65535
                fmt.Println(p0, fi12, p0)
                switch "\n\t" {
                case "a":
                    v15 := GenS1{F0: fi12, F1: 175}
                    _ = v15
                default:
                    p1 = (fi13 | -1)
                }
                _ = v14
            }
            p0 = (p0 / ((18446744073709551615) | 1))
            if ((!true) && (!((true && (false || false)) || (1 > 39)))) {
                p1 = (fi12 * (p1 | fi13))
                v16 := GenS1{F0: 0, F1: -220}
                _ = v16
            } else {
                v17 := GenS2{F0: 962}
                p0 = (p0 / ((p0) | 1))
                _ = v17
            }
            _ = fi13
        }
        p1 -= -1
        if (-128 >= -1) {
            var v18 uint32 = 902
            p2 += (^(p2))
            v19 := []int{-2147483648, (p1 >> ((2) & 63))}
            _ = v18
            _ = v19
        }
        _ = fi12
    }
    if true {
        for fi20 := 0; fi20 < 7; fi20++ {
            var v21 uint16 = 370
            var v22 int32 = 2147483647
            _ = fi20
            _ = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 16; fi23++ {
            if (18446744073709551615 == 4294967295) {
                if true {
                    var v24 int16 = -32768
                    _ = v24
                }
            }
            v25 := GenS2{F0: 42}
            _ = fi23
            _ = v25
        }
        v26 := GenS3{F0: 65535}
        _ = v26
    } else {
        p1 = (-(p1))
        p0 += (p0 << ((2) & 63))
        p2 = 1
    }
    var v27 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return 42
}

func main() {
    var v1 int64 = 1
    v2 := make(map[int]uint32)
    var v3 int = 4294967295
    var v4 int8 = 127
    for fi28 := 0; fi28 < 13; fi28++ {
        v29 := GenS4{F0: false}
        _ = fi28
        _ = v29
    }
    v1 = 826
    fmt.Println(v1, v1, v4)
    v30 := GenS5{F0: 0, F1: 270, F2: v1}
    v3 -= v3
    var v31 int16 = -32768
    for fi32 := 0; fi32 < 2; fi32++ {
        var v33 float64 = float64(0.0)
        var v34 bool = true
        v35 := GenS5{F0: 0, F1: v31, F2: 9223372036854775807}
        _ = fi32
        _ = v33
        _ = v34
        _ = v35
    }
    v3 += 1
    v36 := GenS6{F0: 253, F1: v3, F2: v31}
    v37 := GenS7{F0: 937, F1: 0}
    v3 -= (v3 - v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v37
    _ = v4
}
