package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 uint; F2 uint32 }
func genFn0(p0 uint16, p1 int32, p2 int32) string {
    switch 42 {
    case 582:
        var v7 float64 = -2.25
        _ = v7
    default:
        p1 += (^(p2))
    }
    if (!((((128 < 1) || (p0 >= 370)) || (781 <= 0)) && true)) {
        v8 := []int{986}
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 int64 = 781
            var v11 int32 = (p1 / (((-(p2))) | 1))
            _ = fi9
            _ = v10
            _ = v11
        }
        p0 = p0
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return "a"
}
func genFn1(p0 bool, p1 int64) bool {
    for fi12 := 0; fi12 < 8; fi12++ {
        fmt.Println(p0, fi12)
        for fi13 := 0; fi13 < 12; fi13++ {
            switch -267 {
            case -2147483648:
                v14 := []int{}
                _ = v14
            default:
                if p0 {
                    switch (-(fi12)) {
                    case 32767:
                        var v15 int64 = (^(p1))
                        _ = v15
                    case 4294967295:
                        v16 := []int{(fi12 % ((fi13) | 1)), -9223372036854775808, (fi13 | (fi12 | fi12)), fi12}
                        _ = v16
                    case -2147483648:
                        fmt.Println(fi13, fi12)
                    }
                    var v17 int16 = -262
                    if (255 >= 0) {
                        v17 += v17
                    } else {
                        if (2 != 2) {
                            p0 = (true || ((v17 & -32768) < -1))
                            if false {
                                var v18 uint = 825
                                _ = v18
                            } else {
                                v19 := []int{}
                                _ = v19
                            }
                        } else {
                            var v20 int32 = -668
                            if false {
                                if (true || p0) {
                                    p0 = p0
                                    switch -488 {
                                    case -1:
                                        var v21 float32 = (float32(float64((((float32(((-2.25 - (1.5 * 3.14159)) * 1e10)) - (1e10 + 1.5)) * float32(((1e10 + 1.5) / 0.1))) + 3.14159))) / 3.14159)
                                        _ = v21
                                    case 917:
                                        var v22 uint = 82
                                        _ = v22
                                    }
                                    var v23 uint16 = 157
                                    _ = v23
                                }
                                v24 := []int{}
                                _ = v24
                            }
                            _ = v20
                        }
                        if p0 {
                            if (255 <= 42) {
                                p1 = p1
                            }
                            switch (^(fi12)) {
                            case 32768:
                                if true {
                                    var v25 int = (fi12 / ((fi12) | 1))
                                    _ = v25
                                }
                            default:
                                var v26 uint64 = 42
                                _ = v26
                            }
                        } else {
                            v17 -= (v17 + (v17 >> ((885) & 63)))
                        }
                        v27 := []int{fi12, (fi12 + (fi12 / (((fi12 * -306)) | 1))), fi12}
                        v28 := GenS1{F0: 1}
                        _ = v27
                        _ = v28
                    }
                    _ = v17
                }
            }
            v29 := GenS2{F0: 0}
            fmt.Println(fi13, fi12, fi13)
            _ = fi13
            _ = v29
        }
        _ = fi12
    }
    fmt.Println(p1, p1)
    switch "hello" {
    case "\n\t":
        v30 := GenS1{F0: -717}
        _ = v30
    default:
        var v31 uint8 = 128
        _ = v31
    }
    _ = p0
    _ = p1
    return p0
}

func main() {
    var v1 string = "\n\t"
    var v2 uint32 = 4294967295
    v3 := make(map[string]int)
    v4 := make(map[string]uint32)
    var v5 float32 = 0.0
    var v6 uint = 18446744073709551615
    fmt.Println(v5)
    v32 := GenS3{F0: v2, F1: (v6 + v6), F2: 42}
    for fi33 := 0; fi33 < 16; fi33++ {
        v34 := []int{541, 193}
        _ = fi33
        _ = v34
    }
    for fi35 := 0; fi35 < 4; fi35++ {
        v4[(v1 + v1)] = (^(v2))
        var v36 int8 = 1
        v37 := GenS3{F0: 0, F1: 4294967295, F2: (^(v2))}
        _ = fi35
        _ = v36
        _ = v37
    }
    var v38 float64 = ((0.1 - 1.5) + (-2.25 + 1.5))
    if (-32768 == -32768) {
        v38 -= 0.1
    }
    v38 = float64((v5 * (-2.25 + (v5 + (v5 * v5)))))
    v1 = v1
    var v39 int64 = -9223372036854775808
    if (false || true) {
        v1 = ""
    }
    for fi40 := 0; fi40 < 9; fi40++ {
        switch (("" + v1) + (v1 + "hello")) {
        case "\n\t":
            var v41 uint = v6
            _ = v41
        case "foo bar":
            v2 += v2
        case "hello":
            var v42 float64 = v38
            _ = v42
        }
        _ = fi40
    }
    var v43 int64 = 1
    v44 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v38
    _ = v39
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
