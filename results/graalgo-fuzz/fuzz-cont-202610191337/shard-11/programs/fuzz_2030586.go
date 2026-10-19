package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint32; F1 int32; F2 int64 }
func genFn0(p0 int, p1 int) int32 {
    fmt.Println(p1, p1)
    v6 := []int{}
    fmt.Println(p1)
    var v7 int8 = 86
    for fi8 := 0; fi8 < 18; fi8++ {
        fmt.Println(p1)
        var v9 uint64 = 4294967295
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 94
}

func main() {
    var v1 int32 = -2147483648
    v2 := []string{}
    v3 := make(map[int]bool)
    v4 := make(map[uint32]bool)
    var v5 int = -371
    var v10 float64 = 1.5
    if false {
        var v11 uint32 = 1
        for fi12 := 0; fi12 < 7; fi12++ {
            fmt.Println(v1)
            _ = fi12
        }
        if (!false) {
            var v13 float32 = 0.1
            var v14 bool = (("hello" + "foo bar") < "\n\t")
            _ = v13
            _ = v14
        } else {
            if false {
                v1 -= genFn0((v5 ^ -755), 835)
            } else {
                for fi15 := 0; fi15 < 1; fi15++ {
                    v1 = v1
                    var v16 string = "\n\t"
                    _ = fi15
                    _ = v16
                }
                var v17 string = ""
                _ = v17
            }
            v5 -= 499
            v18 := GenS1{F0: v11}
            _ = v18
        }
        _ = v11
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        v20 := GenS2{F0: (v1 < v1), F1: 223, F2: -535}
        var v21 uint16 = 1
        _ = fi19
        _ = v20
        _ = v21
    }
    v10 = ((v10 / -2.25) + 3.14159)
    if ((0 == 32767) || (false && (1 != 42))) {
        var v22 uint64 = 42
        fmt.Println(v1, v1)
        _ = v22
    } else {
        switch -9223372036854775808 {
        case 9223372036854775807:
            var v23 float32 = (1.5 - 0.0)
            _ = v23
        default:
            var v24 uint = 636
            _ = v24
        }
        switch v5 {
        case 127:
            var v25 int64 = 420
            _ = v25
        }
        _ = len(v4)
    }
    var v26 float64 = 1e10
    v4[42] = true
    var v27 uint32 = 227
    switch (v5 << ((2) & 63)) {
    case 0:
        switch -286 {
        case 0:
            v1 = (v1 * v1)
        }
    case 4294967295:
        var v28 uint8 = 0
        _ = v28
    case 32768:
        var v29 int8 = 93
        _ = v29
    default:
        v30 := GenS3{F0: v27, F1: v1, F2: -783}
        _ = v30
    }
    var v31 float32 = 0.0
    fmt.Println(v5)
    _ = len(v2)
    var v32 int16 = 1000
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v31 -= (0.0 - (1.5 - 1.5))
    v33 := GenS2{F0: true, F1: 18446744073709551615, F2: 1000}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
