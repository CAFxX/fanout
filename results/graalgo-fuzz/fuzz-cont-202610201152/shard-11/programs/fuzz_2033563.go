package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int8; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint64; F1 int64; F2 bool }
func genFn0(p0 int16, p1 uint) uint16 {
    if (false || (398 >= 2)) {
        for fi4 := 0; fi4 < 13; fi4++ {
            p1 = (^(p1))
            var v5 int32 = -289
            _ = fi4
            _ = v5
        }
    } else {
        var v6 float64 = (0.1 - (float64(0.0) - 1.5))
        var v7 float64 = float64(0.0)
        _ = v6
        _ = v7
    }
    switch -2147483648 {
    case 32768:
        var v8 string = ((("\n\t" + ("a" + "\n\t")) + "a") + ("a" + "a"))
        _ = v8
    }
    v9 := GenS1{F0: 2147483647, F1: 42, F2: 2}
    for fi10 := 0; fi10 < 12; fi10++ {
        fmt.Println(p0, p1)
        if true {
            p0 = (^(p0))
        } else {
            var v11 uint16 = 661
            p1 -= p1
            v11 += v11
            _ = v11
        }
        v12 := GenS1{F0: 991, F1: 2, F2: 400}
        _ = fi10
        _ = v12
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v9
    return 150
}
func genFn1(p0 int, p1 int8, p2 uint16) int {
    p1 -= p1
    switch "" {
    case "foo bar":
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 int8 = (p1 & (-(p1)))
            v15 := GenS2{F0: v14, F1: 128, F2: -159}
            for fi16 := 0; fi16 < 9; fi16++ {
                v17 := []int{(fi16 | fi16)}
                var v18 float64 = 1.5
                var v19 int32 = -2147483648
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            _ = fi13
            _ = v14
            _ = v15
        }
    default:
        v20 := []int{375, (^(p0)), p0}
        _ = v20
    }
    p2 = 1
    var v21 float32 = 0.0
    switch 128 {
    case -9223372036854775808:
        p1 += (p1 / ((-1) | 1))
    case 1:
        if (!(-9223372036854775808 != 1)) {
            fmt.Println(p2)
            v22 := []int{}
            for fi23 := 0; fi23 < 4; fi23++ {
                p0 = p0
                switch 2147483647 {
                case 32768:
                    var v24 float32 = 3.14159
                    _ = v24
                case 0:
                    var v25 float32 = v21
                    _ = v25
                case 128:
                    for fi26 := 0; fi26 < 11; fi26++ {
                        switch fi26 {
                        case 128:
                            p1 += (p1 * p1)
                        }
                        v27 := GenS3{F0: 1, F1: 9223372036854775807, F2: (!(v21 < v21))}
                        var v28 uint16 = (p2 + (p2 << ((0) & 63)))
                        if len(v22) > 0 {
                            v22[0] = -1
                        }
                        v29 := GenS1{F0: -1, F1: genFn0(0, 4294967295), F2: (v28 % ((genFn0(42, 4294967295)) | 1))}
                        v22 = append(v22, 9223372036854775807)
                        var v30 int64 = 9223372036854775807
                        v28 = v28
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                }
                v31 := GenS1{F0: -2147483648, F1: (p2 + (p2 | p2)), F2: genFn0(-1, 65535)}
                _ = fi23
                _ = v31
            }
            _ = v22
        } else {
            var v32 uint = 42
            fmt.Println(v21, p0)
            v33 := GenS1{F0: 2147483647, F1: 933, F2: (p2 + (^(p2)))}
            _ = v32
            _ = v33
        }
    case 0:
        v21 -= v21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return -11
}

func main() {
    var v1 float64 = (0.0 + ((1e10 - 1e10) - ((0.1 + 1.5) - ((0.0 + (0.1 * (float64(3.14159) / 1.0))) / -2.25))))
    var v2 int8 = -92
    var v3 int64 = 1
    v34 := []int{-2147483648, -2147483648, 1}
    var v35 uint32 = 2147483647
    v35 = 181
    v3 -= v3
    v2 += v2
    var v36 uint = 954
    fmt.Println(v35, v3)
    v3 = (-(v3))
    var v37 int16 = 1000
    v38 := GenS1{F0: 42, F1: 2, F2: 102}
    v39 := GenS3{F0: 18446744073709551615, F1: (v3 | 457), F2: ((v35 % ((v35) | 1)) <= (v35 ^ (v35 & v35)))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
}
