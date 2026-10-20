package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint8 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int64; F1 uint64; F2 int }
func genFn0(p0 uint32) int32 {
    fmt.Println(p0)
    v7 := []int{1, -2147483648}
    if (1 <= 56) {
        v8 := GenS1{F0: "", F1: 1000, F2: 42}
        p0 += (^(p0))
        if (false || (p0 <= (p0 << ((1) & 63)))) {
            var v9 uint32 = p0
            var v10 uint = 1
            v11 := GenS1{F0: "", F1: -125, F2: 1}
            _ = v10
            _ = v11
            _ = v9
        } else {
            v12 := GenS2{F0: 94, F1: 17, F2: 1}
            _ = v12
        }
        _ = v8
    } else {
        p0 = p0
    }
    p0 = (p0 % ((806) | 1))
    _ = p0
    _ = v7
    return 1
}
func genFn1() uint16 {
    var v13 uint = 18446744073709551615
    var v14 int64 = 42
    switch int(v13) {
    case 42:
        for fi15 := 0; fi15 < 16; fi15++ {
            switch (fi15 >> ((v13) & 63)) {
            case 511:
                v13 -= v13
            }
            if false {
                v16 := GenS2{F0: -1, F1: 42, F2: 42}
                var v17 uint8 = 255
                if (-193 == 561) {
                    v14 -= (v14 << ((v13) & 63))
                }
                _ = v16
                _ = v17
            } else {
                var v18 bool = true
                _ = v18
            }
            _ = fi15
        }
    case -1:
        var v19 float64 = 0.0
        _ = v19
    case 4294967295:
        v13 = v13
    default:
        v20 := []int{253}
        _ = v20
    }
    v14 = v14
    v21 := []int{1}
    _ = v13
    _ = v14
    _ = v21
    return 0
}

func main() {
    var v1 uint16 = 0
    v2 := []int{757}
    v3 := make(map[int]uint32)
    v4 := []int{}
    v5 := make(map[int]bool)
    var v6 int16 = 1000
    var v22 int32 = -1
    v23 := GenS3{F0: -1, F1: 18446744073709551615, F2: 0}
    for fi24 := 0; fi24 < 16; fi24++ {
        switch (^(fi24)) {
        case 605:
            var v25 float64 = (float64((0.1 * 1.5)) + 0.0)
            _ = v25
        }
        _ = fi24
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        var v27 bool = false
        var v28 int64 = -9223372036854775808
        _ = fi26
        _ = v27
        _ = v28
    }
    switch 2147483647 {
    case 0:
        var v29 uint8 = 194
        _ = v29
    }
    if false {
        v1 += 1
    }
    v3[-917] = 2
    v6 = v6
    var v30 float64 = (0.0 - ((3.14159 / (1e10 / 0.1)) / 1.0))
    var v31 bool = true
    v31 = v31
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
