package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS5 struct { F0 int32; F1 int8; F2 uint }
func genFn0(p0 uint16) string {
    var v5 int = -2147483648
    v6 := GenS1{F0: "a"}
    var v7 int8 = 25
    if (uint8(p0) != 255) {
        switch "a" {
        case "a":
            switch ("a" + ("hello" + "hello")) {
            case "a":
                for fi8 := 0; fi8 < 20; fi8++ {
                    if false {
                        p0 = 2
                    } else {
                        var v9 float32 = (float32(float64(3.14159)) * 1.5)
                        switch "" {
                        case "":
                            v10 := []int{-420}
                            _ = v10
                        case "foo bar":
                            v11 := GenS2{F0: 2, F1: 587}
                            _ = v11
                        }
                        _ = v9
                    }
                    v5 = (v5 * fi8)
                    v7 = v7
                    _ = fi8
                }
            case "foo bar":
                for fi12 := 0; fi12 < 13; fi12++ {
                    var v13 float64 = (1.5 / 3.14159)
                    v14 := []int{fi12, v5, v5, (^(fi12))}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            case "\n\t":
                var v15 uint16 = p0
                _ = v15
            default:
                if ((255 < 1) && (true || (v7 < 127))) {
                    for fi16 := 0; fi16 < 11; fi16++ {
                        fmt.Println(p0, p0)
                        if (!true) {
                            switch "foo bar" {
                            case "\n\t":
                                if ((false && (18446744073709551615 <= 2)) && (3.14159 >= (0.0 - 3.14159))) {
                                    var v17 int32 = 44
                                    v18 := GenS2{F0: 709, F1: 379}
                                    _ = v17
                                    _ = v18
                                }
                            case "hello":
                                v7 -= v7
                            case "a":
                                p0 -= (p0 % ((0) | 1))
                            }
                        }
                        for fi19 := 0; fi19 < 9; fi19++ {
                            fmt.Println(v7, v5, fi19)
                            _ = fi19
                        }
                        _ = fi16
                    }
                    v7 += v7
                    v5 -= v5
                }
            }
        }
    } else {
        var v20 int8 = -1
        v5 = (v5 * -80)
        for fi21 := 0; fi21 < 1; fi21++ {
            for fi22 := 0; fi22 < 12; fi22++ {
                var v23 int32 = 2147483647
                var v24 int32 = (v23 - (^(v23)))
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = fi21
        }
        _ = v20
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return ("\n\t" + "\n\t")
}
func genFn1(p0 string, p1 int16, p2 string) int8 {
    p0 = (p2 + (p0 + "a"))
    if true {
        p1 = (p1 / (((^(p1))) | 1))
        var v25 uint = 65535
        _ = v25
    }
    v26 := []int{490, -915}
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    return 1
}

func main() {
    var v1 int16 = 418
    var v2 int8 = -92
    var v3 uint = 18446744073709551615
    var v4 int16 = (v1 + (v1 ^ (v1 + v1)))
    var v27 int16 = v1
    if (!false) {
        if false {
            fmt.Println(v3)
            v28 := []int{32767, -639}
            var v29 int8 = (v2 ^ int8(v3))
            _ = v28
            _ = v29
        }
        v30 := []int{0, 417}
        _ = v30
    } else {
        v3 = v3
        var v31 int64 = -181
        _ = v31
    }
    v2 -= (v2 * v2)
    var v32 int8 = (v2 ^ (v2 % ((v2) | 1)))
    var v33 uint16 = 42
    v33 = 42
    v34 := []int{971, 42}
    v35 := GenS3{F0: 132}
    v36 := GenS2{F0: (^(v3)), F1: 42}
    v4 = (v1 ^ 0)
    v37 := GenS3{F0: v3}
    v38 := GenS1{F0: ""}
    for fi39 := 0; fi39 < 6; fi39++ {
        v32 += -122
        _ = fi39
    }
    switch 128 {
    case 4294967295:
        if ("a" >= ("\n\t" + "a")) {
            var v40 int32 = 2147483647
            v41 := GenS4{F0: -152, F1: 4294967295, F2: uint64(v3)}
            _ = v40
            _ = v41
        }
    case 127:
        var v42 string = ("" + (("" + "hello") + "foo bar"))
        _ = v42
    case 32767:
        v3 -= (^(v3))
    }
    v43 := GenS5{F0: 42, F1: 1, F2: (v3 / ((v3) | 1))}
    switch "hello" {
    case "a":
        if len(v34) > 0 {
            v34[0] = -119
        }
    }
    var v44 string = (("" + "") + ("\n\t" + ("\n\t" + "foo bar")))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v43
    _ = v44
}
