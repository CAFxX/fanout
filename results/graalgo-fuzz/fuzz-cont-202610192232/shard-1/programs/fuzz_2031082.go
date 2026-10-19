package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int8 }
type GenS2 struct { F0 int64; F1 int16; F2 uint }
type GenS3 struct { F0 uint8; F1 bool; F2 uint }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 int }
func genFn0() bool {
    for fi6 := 0; fi6 < 3; fi6++ {
        if true {
            var v7 float32 = ((1e10 + (1.5 * 0.0)) + 1.5)
            v7 = v7
            _ = v7
        }
        _ = fi6
    }
    if (255 >= 42) {
        if true {
            v8 := GenS1{F0: 42, F1: 255, F2: -96}
            v9 := GenS1{F0: 0, F1: 255, F2: -92}
            _ = v8
            _ = v9
        }
        var v10 int16 = 42
        _ = v10
    } else {
        var v11 bool = true
        if (288 > 837) {
            var v12 uint = 757
            v11 = (0.0 == 1.5)
            _ = v12
        } else {
            var v13 int32 = 262
            _ = v13
        }
        v11 = v11
        _ = v11
    }
    return (42 > 0)
}
func genFn1(p0 uint32, p1 uint32, p2 uint8) int32 {
    for fi14 := 0; fi14 < 4; fi14++ {
        v15 := GenS2{F0: int64(p1), F1: -836, F2: 255}
        var v16 int32 = 0
        v17 := []int{}
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    if (!false) {
        p0 = p0
        var v18 string = ""
        p0 = p0
        _ = v18
    }
    var v19 uint = 2
    v20 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    return 0
}

func main() {
    var v1 int = -871
    var v2 uint64 = 42
    var v3 float64 = (1e10 - float64(0.1))
    var v4 uint16 = 532
    var v5 uint8 = 127
    for fi21 := 0; fi21 < 20; fi21++ {
        v22 := []int{v1, (v1 | -24), (fi21 & (fi21 & 42))}
        v2 += 798
        v5 = 128
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        if (v5 <= v5) {
            fmt.Println(v5, v1, v1)
        }
        _ = fi23
    }
    var v24 uint = 479
    v25 := GenS1{F0: 65, F1: (^(v24)), F2: -1}
    v26 := GenS3{F0: uint8(v4), F1: (664 <= 4294967295), F2: v24}
    v1 = int(v4)
    var v27 int16 = -1
    if true {
        for fi28 := 0; fi28 < 19; fi28++ {
            v29 := GenS1{F0: int8(v1), F1: v24, F2: 127}
            _ = fi28
            _ = v29
        }
        if (-1 == -1) {
            v24 += (v24 + v24)
        } else {
            fmt.Println(v5, v2, v27)
            var v30 int64 = -473
            v31 := GenS4{F0: 0, F1: (v24 - v24)}
            _ = v30
            _ = v31
        }
        for fi32 := 0; fi32 < 11; fi32++ {
            v33 := GenS2{F0: -9223372036854775808, F1: (v27 << (((^(v24))) & 63)), F2: 2}
            v34 := GenS3{F0: v5, F1: (!(3.14159 > 1e10)), F2: 172}
            _ = fi32
            _ = v33
            _ = v34
        }
    }
    switch v1 {
    case 2147483647:
        v35 := GenS4{F0: (v5 * (^(v5))), F1: 693}
        _ = v35
    default:
        v36 := GenS5{F0: v1}
        _ = v36
    }
    v5 -= v5
    v37 := GenS3{F0: (v5 % ((128) | 1)), F1: (false && (true && (-128 <= 0))), F2: v24}
    switch v1 {
    case 42:
        v38 := []int{(-(v1)), (v1 * 4294967295), (v1 + 2147483647)}
        _ = v38
    }
    var v39 int64 = 24
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v37
    _ = v39
    _ = v4
    _ = v5
}
