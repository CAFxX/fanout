package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS4 struct { F0 int; F1 uint32; F2 int }
func genFn0(p0 uint16, p1 int32, p2 uint64) int {
    var v6 float64 = (1e10 * (((0.0 - 1.5) - (-2.25 * ((0.0 + (3.14159 / (1.5 * (-2.25 / (0.1 - (((1e10 - -2.25) / 3.14159) * (0.1 - -2.25))))))) / -2.25))) - 0.1))
    var v7 uint32 = 4294967295
    if true {
        var v8 uint32 = 489
        v9 := []int{729}
        _ = v8
        _ = v9
    } else {
        var v10 float64 = v6
        switch "\n\t" {
        case "":
            for fi11 := 0; fi11 < 16; fi11++ {
                var v12 int = (-(fi11))
                _ = fi11
                _ = v12
            }
        case "foo bar":
            for fi13 := 0; fi13 < 7; fi13++ {
                for fi14 := 0; fi14 < 1; fi14++ {
                    var v15 string = "a"
                    v7 = (v7 % ((uint32(p2)) | 1))
                    v16 := []int{}
                    p2 -= (p2 - p2)
                    if ((v10 * v10) <= 1e10) {
                        v16 = append(v16, (fi14 + 0))
                        fmt.Println(v10, v6)
                        switch 521 {
                        case 213:
                            switch v15 {
                            case "":
                                v17 := GenS1{F0: 127, F1: uint(p2)}
                                _ = v17
                            case "hello":
                                v18 := GenS2{F0: (fi13 & (fi13 * (fi14 ^ -1)))}
                                _ = v18
                            case "a":
                                p0 -= (^(p0))
                            }
                        case 2147483647:
                            v6 = 1e10
                        }
                        _ = len(v16)
                    }
                    v10 -= v6
                    p1 = p1
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                fmt.Println(v6, p1)
                _ = fi13
            }
        case "\n\t":
            p1 += p1
        default:
            switch 4294967295 {
            case -9223372036854775808:
                v19 := GenS3{F0: 0, F1: 0, F2: 42}
                _ = v19
            case 128:
                v7 = (v7 / (((^(v7))) | 1))
            default:
                for fi20 := 0; fi20 < 19; fi20++ {
                    fmt.Println(v10, v6)
                    var v21 int64 = 1
                    var v22 int64 = -976
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            }
        }
        v23 := GenS1{F0: 42, F1: 4294967295}
        _ = v10
        _ = v23
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        fmt.Println(p1)
        v25 := []int{-968, fi24}
        _ = fi24
        _ = v25
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -659
}
func genFn1(p0 uint, p1 uint64, p2 uint16) string {
    fmt.Println(p2)
    p2 -= p2
    v26 := GenS3{F0: p2, F1: 1, F2: -1}
    var v27 uint8 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v27
    return "foo bar"
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = 42
    var v3 uint16 = 371
    v4 := []string{"foo bar", "", "\n\t", "foo bar"}
    v5 := []uint32{80, 58, uint32(v3), 0}
    var v28 float32 = (1e10 * 1.5)
    var v29 float32 = v28
    v30 := GenS3{F0: (^(v3)), F1: 132, F2: -32768}
    var v31 bool = (-2.25 >= v28)
    if (!(((!v31) && true) && false)) {
        v1[("a" + "\n\t")] = true
        v1["hello"] = (v31 || true)
        v28 -= (v28 * 3.14159)
    }
    for fi32 := 0; fi32 < 10; fi32++ {
        fmt.Println(v2, v28, v2)
        if (v31 && (v31 && true)) {
            delete(v1, "\n\t")
            if (2 >= 2147483647) {
                v2 = fi32
                v31 = (v31 || v31)
            }
        } else {
            fmt.Println(v28, v28, v28)
            v33 := GenS4{F0: (v2 ^ 4294967295), F1: 715, F2: (v2 & 855)}
            var v34 uint32 = uint32(v3)
            _ = v33
            _ = v34
        }
        for fi35 := 0; fi35 < 5; fi35++ {
            var v36 int = genFn0(v3, -1, 21)
            v29 += v29
            _ = fi35
            _ = v36
        }
        _ = fi32
    }
    fmt.Println(v28, v2, v29)
    v31 = v31
    v29 += ((0.1 / v28) / 3.14159)
    fmt.Println(v28)
    for fi37 := 0; fi37 < 15; fi37++ {
        v3 = v3
        _ = fi37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
