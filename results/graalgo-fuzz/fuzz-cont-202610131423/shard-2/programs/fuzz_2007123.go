package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool; F1 int }
func genFn0(p0 int64, p1 uint64) string {
    fmt.Println(p1, p1)
    v5 := GenS1{F0: 32767}
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := []int{fi6, (fi6 | 1)}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return ""
}

func main() {
    var v1 uint16 = 0
    var v2 bool = false
    var v3 uint16 = (^(v1))
    var v4 int32 = 0
    var v8 int32 = (v4 * (v4 + (v4 & -2147483648)))
    for fi9 := 0; fi9 < 8; fi9++ {
        switch (^(fi9)) {
        case 860:
            var v10 int16 = -840
            _ = v10
        case -145:
            v11 := GenS2{F0: 501, F1: 0}
            _ = v11
        default:
            fmt.Println(v8, v1, v4)
        }
        v4 = (v4 * v8)
        _ = fi9
    }
    var v12 int32 = v4
    v13 := []int{-2147483648, -9223372036854775808}
    if v2 {
        switch "a" {
        case "foo bar":
            v3 = v1
        case "hello":
            v8 += (v8 + (-(v8)))
        case "":
            _ = len(v13)
        }
    } else {
        if v2 {
            var v14 uint = 18446744073709551615
            _ = v14
        }
        v15 := GenS1{F0: (v1 | 42)}
        _ = v15
    }
    fmt.Println(v8)
    if (v2 && v2) {
        if len(v13) > 0 {
            v13[0] = 466
        }
    } else {
        for fi16 := 0; fi16 < 4; fi16++ {
            v3 = v3
            var v17 int64 = 587
            _ = fi16
            _ = v17
        }
        if (-252 >= 32768) {
            v18 := GenS1{F0: 255}
            _ = len(v13)
            if v2 {
                v13 = append(v13, -68)
            } else {
                v19 := GenS2{F0: 18446744073709551615, F1: 287}
                fmt.Println(v2, v1)
                v1 -= (v3 - v1)
                _ = v19
            }
            _ = v18
        } else {
            var v20 bool = true
            var v21 uint64 = 2
            for fi22 := 0; fi22 < 15; fi22++ {
                if len(v13) > 0 {
                    v13[0] = (fi22 | fi22)
                }
                if v20 {
                    _ = len(v13)
                }
                _ = fi22
            }
            _ = v20
            _ = v21
        }
        fmt.Println(v4, v8, v1)
    }
    if false {
        for fi23 := 0; fi23 < 6; fi23++ {
            switch 32768 {
            case 92:
                var v24 int64 = -1
                _ = v24
            case 0:
                v8 = 342
            default:
                v8 -= (-(v8))
            }
            _ = fi23
        }
    } else {
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        var v25 uint8 = 0
        _ = v25
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        if (((v2 || v2) || true) || (1.5 == ((((0.1 * ((0.1 * 3.14159) + ((float64(1e10) / (3.14159 + 0.0)) - float64(float32(0.0))))) * 3.14159) - 0.1) * 3.14159))) {
            if ((394 != 489) || (1 == -128)) {
                v27 := GenS1{F0: (v1 + (v1 >> ((255) & 63)))}
                v28 := GenS3{F0: 42, F1: v3}
                var v29 string = "\n\t"
                _ = v27
                _ = v28
                _ = v29
            }
            var v30 uint16 = v1
            if (-958 < fi26) {
                if true {
                    var v31 float64 = (3.14159 / (1.5 / 3.14159))
                    v13 = append(v13, fi26)
                    _ = v31
                } else {
                    v32 := GenS2{F0: 1, F1: 2147483647}
                    if (0 != -32768) {
                        var v33 int64 = -787
                        if (2 != 75) {
                            var v34 bool = (2 == 2)
                            v13 = append(v13, (fi26 ^ -9223372036854775808))
                            _ = v34
                        } else {
                            for fi35 := 0; fi35 < 11; fi35++ {
                                fmt.Println(fi26, v3)
                                var v36 int16 = 42
                                _ = fi35
                                _ = v36
                            }
                            var v37 uint = 746
                            v38 := GenS4{F0: -85}
                            _ = v37
                            _ = v38
                        }
                        _ = len(v13)
                        _ = v33
                    } else {
                        var v39 float64 = 0.1
                        _ = v39
                    }
                    v40 := GenS5{F0: v2, F1: (fi26 + fi26)}
                    _ = v32
                    _ = v40
                }
                fmt.Println(v12)
            } else {
                var v41 uint = 901
                _ = v41
            }
            _ = v30
        }
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
