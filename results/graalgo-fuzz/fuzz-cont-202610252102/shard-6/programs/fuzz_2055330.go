package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS4 struct { F0 string }
func genFn0(p0 uint) bool {
    switch 127 {
    case 128:
        for fi7 := 0; fi7 < 17; fi7++ {
            v8 := GenS1{F0: 0, F1: (fi7 - (^(fi7))), F2: p0}
            for fi9 := 0; fi9 < 11; fi9++ {
                var v10 int32 = 2147483647
                v11 := []int{}
                v12 := GenS1{F0: 18446744073709551615, F1: (fi7 << (((p0 / (((p0 / ((p0) | 1))) | 1))) & 63)), F2: (p0 % ((0) | 1))}
                var v13 uint16 = 0
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
                _ = v13
            }
            var v14 int32 = 0
            _ = fi7
            _ = v14
            _ = v8
        }
    case 376:
        for fi15 := 0; fi15 < 17; fi15++ {
            p0 -= (p0 / (((^(p0))) | 1))
            p0 = p0
            _ = fi15
        }
    case -2147483648:
        var v16 string = ("a" + ("a" + ""))
        _ = v16
    }
    var v17 uint32 = 2147483648
    _ = p0
    _ = v17
    return true
}
func genFn1(p0 uint8) bool {
    if (1.5 >= 3.14159) {
        fmt.Println(p0)
    } else {
        p0 += p0
        fmt.Println(p0)
        v18 := GenS2{F0: -2147483648}
        _ = v18
    }
    fmt.Println(p0)
    fmt.Println(p0)
    var v19 float32 = 0.1
    switch "\n\t" {
    case "a":
        fmt.Println(p0, v19)
    default:
        var v20 uint = 4294967295
        _ = v20
    }
    _ = p0
    _ = v19
    return ((0 >= -128) || (false || (true || true)))
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 2147483647
    v3 := make(map[int]bool)
    var v4 uint8 = 0
    var v5 uint16 = 1
    v6 := make(map[uint32]int)
    switch "\n\t" {
    case "\n\t":
        v1[333] = (564 > 18446744073709551615)
    case "hello":
        var v21 int64 = 554
        _ = v21
    case "":
        if true {
            for fi22 := 0; fi22 < 4; fi22++ {
                v23 := []int{(fi22 >> ((4294967295) & 63)), 4294967295, (fi22 * fi22)}
                if (v5 == (^(v5))) {
                    if false {
                        fmt.Println(v5)
                        v4 += (^(v4))
                    } else {
                        var v24 uint16 = 535
                        _ = v24
                    }
                } else {
                    v25 := GenS3{F0: -190, F1: v5, F2: 4294967295}
                    if len(v23) >= 2 {
                        v23 = v23[1:len(v23)]
                    }
                    _ = v25
                }
                _ = fi22
                _ = v23
            }
        } else {
            _ = len(v6)
            _ = len(v3)
        }
    }
    if (0.1 != 0.0) {
        for fi26 := 0; fi26 < 7; fi26++ {
            var v27 int8 = 1
            var v28 int8 = (v27 * v27)
            fmt.Println(v28, v28, v28)
            _ = fi26
            _ = v27
            _ = v28
        }
        v29 := GenS2{F0: 42}
        var v30 int64 = 525
        _ = v29
        _ = v30
    } else {
        v31 := GenS4{F0: ((("\n\t" + ("" + ("foo bar" + ("" + ("\n\t" + "a"))))) + "") + "\n\t")}
        _ = v31
    }
    v2 += v2
    v32 := []int{42}
    var v33 uint32 = 705
    var v34 bool = false
    if v34 {
        if v34 {
            v33 -= 877
        } else {
            fmt.Println(v4, v34, v33)
            v34 = true
            for fi35 := 0; fi35 < 2; fi35++ {
                v34 = (3.14159 < ((((((((0.0 * float64(-2.25)) + (((float64(-2.25) * ((3.14159 + 0.1) + (1e10 - (0.0 - 1.5)))) + float64((0.0 + float32(1.5)))) - 1.5)) + ((-2.25 / 1e10) * (float64(float32(1.5)) - ((1e10 - (1e10 + 1.5)) * -2.25)))) - 1e10) - 3.14159) / (3.14159 * 0.1)) + 1.5) + (0.1 / float64(0.1))))
                v5 -= (v5 ^ 789)
                _ = fi35
            }
        }
        var v36 uint16 = v5
        v5 = v5
        _ = v36
    } else {
        var v37 bool = v34
        v38 := GenS3{F0: 304, F1: v5, F2: uint64(v2)}
        fmt.Println(v4)
        _ = v37
        _ = v38
    }
    v6[v2] = 127
    var v39 bool = false
    var v40 uint16 = uint16(v33)
    v41 := GenS3{F0: 1, F1: v40, F2: 549}
    switch 128 {
    case -2147483648:
        for fi42 := 0; fi42 < 8; fi42++ {
            if v34 {
                v39 = (32767 < -1)
                delete(v1, (^(fi42)))
            } else {
                var v43 int8 = 0
                _ = v43
            }
            var v44 uint64 = 4294967295
            _ = fi42
            _ = v44
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
