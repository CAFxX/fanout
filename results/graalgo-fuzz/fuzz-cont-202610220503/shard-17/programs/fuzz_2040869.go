package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 uint64; F1 uint; F2 int8 }
type GenS5 struct { F0 int64; F1 uint8 }
func genFn0(p0 int64, p1 int) uint32 {
    v5 := []int{-9223372036854775808, p1, p1}
    if len(v5) > 0 {
        v5[0] = p1
    }
    _ = p0
    _ = p1
    _ = v5
    return 1
}

func main() {
    var v1 int8 = 1
    v2 := []int{-306}
    var v3 int = -2147483648
    var v4 float32 = float32(((float64(((0.0 + 3.14159) - 1e10)) - float64(-2.25)) * (1e10 + 0.0)))
    if false {
        var v6 uint64 = 0
        switch "a" {
        case "foo bar":
            if ((v1 | v1) != v1) {
                var v7 int16 = 32767
                _ = v7
            } else {
                fmt.Println(v1, v3, v4)
            }
        case "a":
            var v8 int64 = -1
            _ = v8
        case "hello":
            fmt.Println(v4, v1, v6)
        }
        _ = v6
    }
    v9 := GenS1{F0: -1}
    v1 = (v1 / ((v1) | 1))
    var v10 float64 = float64(v4)
    if (true || (((1.5 * v10) != (v10 + (0.1 + v10))) || true)) {
        if (54 >= 0) {
            if (true || false) {
                v4 -= ((float32(float64(3.14159)) / ((v4 + v4) - (((v4 + (v4 / (float32(((v10 - v10) * v10)) / v4))) + 3.14159) * v4))) * v4)
            } else {
                switch ("" + "a") {
                case "foo bar":
                    for fi11 := 0; fi11 < 1; fi11++ {
                        v3 -= fi11
                        _ = fi11
                    }
                }
                var v12 float32 = v4
                _ = v12
            }
            if (418 <= 9223372036854775807) {
                v13 := GenS2{F0: 65535}
                var v14 uint32 = 697
                v4 = float32(v10)
                _ = v13
                _ = v14
            }
        } else {
            fmt.Println(v1)
            var v15 bool = (v1 >= v1)
            var v16 float32 = (v4 * v4)
            _ = v15
            _ = v16
        }
    } else {
        for fi17 := 0; fi17 < 6; fi17++ {
            switch "" {
            case "hello":
                v3 = -1
            default:
                v18 := GenS2{F0: 255}
                _ = v18
            }
            var v19 float32 = float32((v10 - v10))
            _ = fi17
            _ = v19
        }
    }
    v4 = (1e10 / (v4 / (v4 - float32(v10))))
    for fi20 := 0; fi20 < 19; fi20++ {
        switch 128 {
        case 128:
            var v21 uint8 = 128
            _ = v21
        case -19:
            _ = len(v2)
        }
        if len(v2) > 0 {
            v2[0] = (v3 >> ((2) & 63))
        }
        fmt.Println(v10)
        _ = fi20
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        for fi23 := 0; fi23 < 11; fi23++ {
            v2 = append(v2, (fi22 % ((fi23) | 1)))
            _ = fi23
        }
        for fi24 := 0; fi24 < 8; fi24++ {
            switch "hello" {
            case "":
                switch "a" {
                case "a":
                    if false {
                        fmt.Println(fi24, v4, fi22)
                        if true {
                            var v25 int = (fi24 * v3)
                            v3 = (v3 % (((fi24 & fi24)) | 1))
                            _ = v25
                        }
                    }
                default:
                    fmt.Println(v10, fi22, v4)
                }
            case "\n\t":
                _ = len(v2)
            case "foo bar":
                v4 -= v4
            }
            var v26 bool = false
            if v26 {
                var v27 uint32 = 374
                var v28 uint8 = 0
                var v29 float32 = v4
                _ = v27
                _ = v28
                _ = v29
            } else {
                v2 = append(v2, fi22)
            }
            _ = fi24
            _ = v26
        }
        if (65535 > 492) {
            v30 := GenS1{F0: 1}
            _ = v30
        } else {
            v4 += 1e10
        }
        _ = fi22
    }
    if (!true) {
        v10 = v10
        if true {
            var v31 int = -1
            v31 = -1
            _ = v31
        } else {
            for fi32 := 0; fi32 < 1; fi32++ {
                v3 = -965
                var v33 uint8 = 36
                v34 := GenS3{F0: 419, F1: 255}
                _ = fi32
                _ = v33
                _ = v34
            }
            v35 := GenS4{F0: 0, F1: 892, F2: v1}
            _ = v35
        }
        var v36 int64 = -1
        _ = v36
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (false || (4294967295 != 0)) {
        v37 := GenS5{F0: 1, F1: 127}
        v1 += (v1 - 1)
        _ = v37
    }
    fmt.Println(v1)
    for fi38 := 0; fi38 < 8; fi38++ {
        v4 = v4
        _ = fi38
    }
    var v39 uint64 = 1
    v40 := GenS4{F0: v39, F1: uint(v39), F2: 25}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v9
}
