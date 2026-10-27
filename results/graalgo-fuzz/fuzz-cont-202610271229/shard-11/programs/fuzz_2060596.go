package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 uint64; F2 int32 }
func genFn0(p0 int, p1 uint64) uint {
    var v7 int16 = 42
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            v7 = (v7 << ((1) & 63))
            p1 += (^(p1))
            switch "" {
            case "":
                p0 = fi9
            case "foo bar":
                v10 := GenS1{F0: 1}
                v11 := GenS1{F0: 2147483647}
                if (1 == -2147483648) {
                    switch ("" + "foo bar") {
                    case "a":
                        var v12 float64 = (0.0 - 1e10)
                        _ = v12
                    case "hello":
                        if ((^(fi8)) != fi8) {
                            p1 = (p1 << ((914) & 63))
                            p0 -= 42
                            var v13 float64 = 1.5
                            _ = v13
                        }
                    default:
                        v14 := GenS1{F0: 1}
                        var v15 uint8 = 0
                        v7 = (v7 + v7)
                        _ = v14
                        _ = v15
                    }
                    v16 := GenS2{F0: 28, F1: 42}
                    _ = v16
                }
                _ = v10
                _ = v11
            default:
                p0 = fi9
            }
            switch "foo bar" {
            case "a":
                v17 := GenS1{F0: 2147483647}
                _ = v17
            case "":
                var v18 string = ("foo bar" + "foo bar")
                _ = v18
            }
            _ = fi9
        }
        _ = fi8
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        v20 := GenS3{F0: 65535}
        var v21 bool = false
        _ = fi19
        _ = v20
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v7
    return 1
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[string]bool)
    var v3 float64 = -2.25
    v4 := make(map[int]uint32)
    var v5 int = 917
    var v6 float32 = (((0.1 * (0.0 * 1e10)) + (1.5 * (float32(1e10) / 3.14159))) + (3.14159 - 1e10))
    var v22 uint8 = 128
    v23 := GenS1{F0: 679}
    fmt.Println(v5, v5)
    _ = len(v2)
    v22 = v22
    if false {
        v24 := GenS3{F0: 2}
        v25 := GenS4{F0: 68, F1: v1, F2: -1}
        if true {
            v26 := []int{}
            v3 = v3
            _ = v26
        } else {
            v1 -= (v1 ^ v1)
            var v27 uint16 = 934
            _ = v27
        }
        _ = v24
        _ = v25
    } else {
        switch v5 {
        case -9223372036854775808:
            var v28 uint = 492
            _ = v28
        default:
            v29 := GenS3{F0: 944}
            _ = v29
        }
        var v30 float64 = 3.14159
        switch "hello" {
        case "foo bar":
            var v31 int8 = 127
            _ = v31
        case "\n\t":
            for fi32 := 0; fi32 < 1; fi32++ {
                var v33 int8 = -26
                _ = fi32
                _ = v33
            }
        }
        _ = v30
    }
    v34 := GenS1{F0: 1}
    v35 := []int{(v5 % (((v5 + v5)) | 1)), v5}
    var v36 float64 = (v3 * (v3 - v3))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
