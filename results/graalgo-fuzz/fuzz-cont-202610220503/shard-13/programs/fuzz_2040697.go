package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 uint16 }
func genFn0(p0 uint16) uint8 {
    switch -787 {
    case 32767:
        var v5 uint32 = 2
        _ = v5
    case -2147483648:
        var v6 float64 = (1.5 * -2.25)
        _ = v6
    case 0:
        var v7 uint16 = 1
        _ = v7
    default:
        var v8 int32 = -2147483648
        _ = v8
    }
    if false {
        for fi9 := 0; fi9 < 10; fi9++ {
            p0 = 133
            var v10 int64 = -999
            v10 -= -367
            _ = fi9
            _ = v10
        }
        switch 128 {
        case 127:
            v11 := []int{}
            _ = v11
        case -1:
            var v12 int32 = 0
            _ = v12
        default:
            if false {
                var v13 int64 = -375
                _ = v13
            }
        }
    } else {
        v14 := GenS1{F0: 118}
        fmt.Println(p0)
        _ = v14
    }
    p0 = (p0 >> ((2) & 63))
    var v15 int8 = -128
    for fi16 := 0; fi16 < 10; fi16++ {
        var v17 uint16 = (p0 << ((869) & 63))
        if false {
            var v18 bool = (1e10 >= 3.14159)
            v15 -= -110
            fmt.Println(v15, fi16)
            _ = v18
        } else {
            v19 := GenS2{F0: 1, F1: 76, F2: 0}
            _ = v19
        }
        _ = fi16
        _ = v17
    }
    _ = p0
    _ = v15
    return 102
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 1
    var v3 uint32 = 42
    v4 := []string{}
    var v20 uint16 = 499
    var v21 uint = v2
    if ((^(v3)) >= (v3 % (((^(v3))) | 1))) {
        v22 := GenS1{F0: 127}
        v23 := GenS1{F0: 127}
        v24 := v1[546]
        _ = v22
        _ = v23
        _ = v24
    } else {
        v25 := GenS2{F0: 65535, F1: 203, F2: v20}
        _ = v25
    }
    v26 := GenS3{F0: v20, F1: 32767, F2: 2}
    v3 = v3
    var v27 int64 = 1
    var v28 int32 = 0
    if false {
        if len(v4) > 0 {
            v4[0] = "hello"
        }
        fmt.Println(v20, v27, v21)
        v27 = v27
    } else {
        var v29 bool = false
        v27 += v27
        if v29 {
            v3 = v3
            v21 -= v2
            v27 = (v27 * v27)
        }
        _ = v29
    }
    v30 := GenS4{F0: 0, F1: (true || (97 == 42))}
    var v31 float64 = ((float64(1.5) / -2.25) - -2.25)
    fmt.Println(v21, v31, v20)
    v31 += (v31 - ((((v31 + (3.14159 + 1e10)) - v31) / (-2.25 * v31)) - (float64((0.1 / ((0.1 * (((0.1 + 0.1) / 1.5) * 1e10)) + 1.5))) + float64(1e10))))
    var v32 uint = v21
    switch 2147483647 {
    case 877:
        switch -445 {
        case 32767:
            v21 = (v2 | v21)
        default:
            var v33 int8 = 0
            _ = v33
        }
    case -331:
        v31 -= (v31 * (v31 - v31))
    case 0:
        if (true && (-9223372036854775808 >= -9223372036854775808)) {
            _ = len(v4)
            v34 := GenS5{F0: 627}
            var v35 int16 = 153
            _ = v34
            _ = v35
        }
    default:
        var v36 float32 = 3.14159
        _ = v36
    }
    var v37 int = int(v21)
    var v38 int64 = v27
    var v39 uint16 = (v20 | v20)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
