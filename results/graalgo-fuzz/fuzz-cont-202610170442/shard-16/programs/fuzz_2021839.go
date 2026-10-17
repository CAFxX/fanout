package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint64; F1 uint; F2 int16 }
type GenS3 struct { F0 int; F1 int32; F2 int }
type GenS4 struct { F0 int16 }
func genFn0(p0 int8, p1 uint) uint16 {
    p0 -= -1
    var v4 uint = 18446744073709551615
    var v5 int16 = 468
    if false {
        v6 := []int{-674, 9223372036854775807, -884, 32767}
        if (1 <= 122) {
            for fi7 := 0; fi7 < 14; fi7++ {
                if (!true) {
                    v5 = int16(p0)
                    var v8 int16 = v5
                    fmt.Println(v4)
                    _ = v8
                }
                v4 = p1
                for fi9 := 0; fi9 < 7; fi9++ {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    var v10 int32 = 607
                    switch "hello" {
                    case "\n\t":
                        var v11 int16 = 390
                        _ = v11
                    case "a":
                        v12 := GenS1{F0: 255, F1: (v5 % ((int16(p0)) | 1)), F2: 32767}
                        v13 := GenS1{F0: 552, F1: v5, F2: 162}
                        fmt.Println(p1)
                        _ = v12
                        _ = v13
                    case "foo bar":
                        if true {
                            if len(v6) > 0 {
                                v6[0] = (fi7 * (fi9 / ((fi9) | 1)))
                            }
                            v10 += (v10 | (v10 + (-(v10))))
                        }
                        v14 := GenS1{F0: 65535, F1: 0, F2: 42}
                        var v15 uint64 = uint64(v4)
                        v15 -= 1
                        v16 := GenS1{F0: 42, F1: v5, F2: 938}
                        _ = v14
                        _ = v15
                        _ = v16
                    default:
                        var v17 float64 = 0.0
                        _ = v17
                    }
                    _ = fi9
                    _ = v10
                }
                _ = fi7
            }
            v18 := GenS2{F0: 716, F1: p1, F2: 0}
            var v19 int8 = p0
            _ = v18
            _ = v19
        } else {
            v20 := GenS3{F0: -9223372036854775808, F1: -2147483648, F2: -447}
            _ = v20
        }
        for fi21 := 0; fi21 < 3; fi21++ {
            v22 := GenS3{F0: -1, F1: -2147483648, F2: (-(fi21))}
            _ = len(v6)
            for fi23 := 0; fi23 < 9; fi23++ {
                if len(v6) > 0 {
                    v6[0] = 128
                }
                fmt.Println(fi23, v5, v4)
                var v24 uint64 = 0
                _ = fi23
                _ = v24
            }
            _ = fi21
            _ = v22
        }
        _ = v6
    } else {
        switch ("hello" + "foo bar") {
        case "\n\t":
            fmt.Println(p0, p1, p1)
        default:
            v25 := []int{-30, 4294967295, 32768, 4294967295}
            _ = v25
        }
        p1 = 193
    }
    v26 := GenS3{F0: 42, F1: 0, F2: 32768}
    _ = p0
    _ = p1
    _ = v26
    _ = v4
    _ = v5
    return 2
}
func genFn1(p0 uint16) uint32 {
    v27 := GenS1{F0: 1, F1: 660, F2: p0}
    p0 = p0
    var v28 uint = 833
    for fi29 := 0; fi29 < 5; fi29++ {
        v30 := GenS4{F0: 0}
        var v31 int = fi29
        var v32 int64 = 1
        _ = fi29
        _ = v30
        _ = v31
        _ = v32
    }
    _ = p0
    _ = v27
    _ = v28
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 18446744073709551615
    var v3 int64 = 0
    if false {
        if ((v2 & (v2 / (((v2 >> ((187) & 63))) | 1))) < (v2 / ((v2) | 1))) {
            for fi33 := 0; fi33 < 7; fi33++ {
                v34 := GenS2{F0: v2, F1: 42, F2: -1}
                _ = fi33
                _ = v34
            }
            v1["hello"] = ((87 >= -1) && false)
        } else {
            var v35 uint = 42
            v3 = (v3 | v3)
            _ = v35
        }
    } else {
        v2 -= v2
        v3 -= (v3 | (-(v3)))
    }
    for fi36 := 0; fi36 < 4; fi36++ {
        var v37 uint32 = 65535
        var v38 int64 = v3
        for fi39 := 0; fi39 < 18; fi39++ {
            v40 := []int{fi36}
            switch (("a" + "") + "") {
            case "hello":
                v3 -= v3
            default:
                v37 = (v37 ^ genFn1(uint16(v2)))
            }
            _ = fi39
            _ = v40
        }
        _ = fi36
        _ = v37
        _ = v38
    }
    fmt.Println(v2)
    var v41 uint = 235
    v2 += (v2 ^ (v2 + 0))
    var v42 int64 = v3
    var v43 float64 = float64((float32((((0.1 - (float64(-2.25) - 0.1)) + -2.25) * float64(float32((1.5 * (1.5 / 3.14159)))))) + 1e10))
    delete(v1, "\n\t")
    v41 = v41
    v44 := []int{32767, -342, 1, -1}
    fmt.Println(v41)
    if len(v44) >= 2 {
        v44 = v44[1:len(v44)]
    }
    v45 := GenS3{F0: -1, F1: 638, F2: 127}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v44:", len(v44), v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}
