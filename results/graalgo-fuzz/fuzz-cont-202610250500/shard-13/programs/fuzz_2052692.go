package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int64 }
func genFn0() uint16 {
    var v7 float32 = (3.14159 - 1e10)
    var v8 string = ""
    var v9 int = 473
    v10 := GenS1{F0: 65535, F1: 209, F2: (v9 % (((v9 % (((v9 << ((415) & 63))) | 1))) | 1))}
    fmt.Println(v9, v8)
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 46
}

func main() {
    var v1 int32 = -685
    var v2 string = "\n\t"
    var v3 float32 = 1e10
    v4 := make(map[int]bool)
    var v5 string = (v2 + "foo bar")
    var v6 uint8 = 42
    var v11 uint8 = (v6 >> ((0) & 63))
    for fi12 := 0; fi12 < 14; fi12++ {
        fmt.Println(fi12, v5, v1)
        var v13 uint32 = 65535
        if (true && false) {
            var v14 uint16 = 557
            var v15 uint8 = v11
            var v16 int = fi12
            _ = v14
            _ = v15
            _ = v16
        }
        _ = fi12
        _ = v13
    }
    var v17 uint16 = 32767
    _ = len(v4)
    for fi18 := 0; fi18 < 16; fi18++ {
        v4[int(v1)] = false
        switch v5 {
        case "a":
            switch (fi18 & 32768) {
            case -567:
                fmt.Println(fi18, v5)
            default:
                v4[(fi18 - (fi18 / (((fi18 + int(v1))) | 1)))] = false
            }
        case "\n\t":
            v5 = v2
        default:
            var v19 uint8 = (^(v11))
            _ = v19
        }
        if (!true) {
            fmt.Println(v17, v5)
            switch fi18 {
            case 955:
                var v20 uint16 = 2
                _ = v20
            case -151:
                v5 = v5
            default:
                var v21 int64 = 1
                _ = v21
            }
            v11 -= 42
        } else {
            v22 := GenS2{F0: 2147483648}
            v23 := v4[610]
            _ = v22
            _ = v23
        }
        _ = fi18
    }
    v1 = (-(v1))
    v24 := GenS3{F0: 2, F1: -746}
    fmt.Println(v6, v3, v11)
    var v25 float32 = -2.25
    var v26 bool = (false && (!true))
    var v27 uint = 850
    if v26 {
        v3 = ((v25 / (v3 / 1.5)) / (v3 / v3))
        var v28 uint64 = 1
        v29 := GenS1{F0: (v27 * 61), F1: v6, F2: 32767}
        _ = v28
        _ = v29
    }
    v30 := []int{1, 128, 0}
    v11 = v11
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
