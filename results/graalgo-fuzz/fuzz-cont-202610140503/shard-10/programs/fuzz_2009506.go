package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int8; F2 uint8 }
type GenS3 struct { F0 int }
func genFn0(p0 uint8, p1 uint64, p2 int32) int8 {
    p0 = p0
    var v7 uint = 2
    fmt.Println(p2, p1)
    var v8 int16 = 860
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 82
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{42, 2}
    v3 := make(map[string]string)
    v4 := make(map[int]bool)
    v5 := make(map[string]uint32)
    var v6 float64 = 1e10
    var v9 uint32 = 65535
    fmt.Println(v6, v9)
    v6 = float64(((((1e10 - 1e10) + (float32((v6 - ((1.5 - 0.1) - (0.1 / v6)))) + (1.5 * float32(v6)))) * -2.25) - 0.0))
    v10 := GenS1{F0: 127}
    v6 = -2.25
    fmt.Println(v9, v9)
    v6 += (((3.14159 / (0.0 + v6)) + (float64(float32(v6)) * (((((v6 - (v6 / (((1.5 - (0.1 + v6)) / v6) - (-2.25 * 0.1)))) / -2.25) + 0.0) * v6) * v6))) * 3.14159)
    v11 := GenS1{F0: -434}
    var v12 int32 = 0
    for fi13 := 0; fi13 < 6; fi13++ {
        if ((1e10 + 1.5) != float32(-2.25)) {
            v14 := GenS1{F0: fi13}
            for fi15 := 0; fi15 < 12; fi15++ {
                var v16 int = (fi15 << ((255) & 63))
                v16 += (-(v16))
                _ = fi15
                _ = v16
            }
            v17 := GenS2{F0: (v12 | 2147483647), F1: 0, F2: 0}
            _ = v14
            _ = v17
        } else {
            var v18 uint16 = 785
            _ = v18
        }
        switch (fi13 * (fi13 % (((fi13 % ((fi13) | 1))) | 1))) {
        case 4294967295:
            v12 -= v12
        default:
            v6 -= v6
        }
        var v19 float32 = float32(v6)
        _ = fi13
        _ = v19
    }
    v9 = v9
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 int64 = -417
        fmt.Println(v12, fi20, v6)
        _ = fi20
        _ = v21
    }
    if (1000 <= -404) {
        for fi22 := 0; fi22 < 11; fi22++ {
            switch fi22 {
            case 1:
                v12 = (v12 & (v12 ^ v12))
            case 9223372036854775807:
                v9 = 0
            case -1:
                fmt.Println(v9, v9)
            default:
                fmt.Println(v12)
            }
            _ = fi22
        }
    } else {
        v9 = (^(v9))
        for fi23 := 0; fi23 < 20; fi23++ {
            var v24 string = "foo bar"
            if false {
                for fi25 := 0; fi25 < 18; fi25++ {
                    fmt.Println(v24, fi23)
                    var v26 int = fi23
                    var v27 float64 = v6
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                v28 := GenS2{F0: v12, F1: -18, F2: 145}
                _ = v28
            } else {
                var v29 int = (fi23 << ((18446744073709551615) & 63))
                for fi30 := 0; fi30 < 3; fi30++ {
                    v31 := GenS3{F0: fi23}
                    _ = fi30
                    _ = v31
                }
                v6 += v6
                _ = v29
            }
            _ = fi23
            _ = v24
        }
        for fi32 := 0; fi32 < 10; fi32++ {
            v6 = -2.25
            _ = fi32
        }
    }
    var v33 int8 = -1
    var v34 uint32 = (v9 << ((503) & 63))
    switch "\n\t" {
    case "a":
        switch ("a" + "\n\t") {
        case "foo bar":
            var v35 uint64 = 2
            _ = v35
        case "hello":
            v36 := GenS3{F0: -346}
            _ = v36
        case "\n\t":
            v37 := GenS3{F0: int(v33)}
            _ = v37
        }
    case "foo bar":
        for fi38 := 0; fi38 < 7; fi38++ {
            var v39 uint8 = 127
            for fi40 := 0; fi40 < 20; fi40++ {
                v41 := GenS1{F0: (fi40 << ((700) & 63))}
                _ = fi40
                _ = v41
            }
            _ = fi38
            _ = v39
        }
    }
    var v42 uint16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v9
}
