package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint32; F1 uint32; F2 uint64 }
func genFn0(p0 int8, p1 uint8, p2 uint8) uint64 {
    for fi4 := 0; fi4 < 6; fi4++ {
        var v5 uint32 = 2147483648
        _ = fi4
        _ = v5
    }
    p0 += (p0 ^ (p0 | 42))
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := []int{(fi6 & -9223372036854775808), 616, (fi6 << ((355) & 63))}
        var v8 bool = true
        fmt.Println(p2, p1, p1)
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 int64 = -9223372036854775808
        fmt.Println(fi9, p0, p2)
        if (!false) {
            p1 = p2
            var v11 bool = false
            v12 := GenS1{F0: 11}
            _ = v11
            _ = v12
        } else {
            v13 := []int{}
            _ = v13
        }
        _ = fi9
        _ = v10
    }
    var v14 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 614
}
func genFn1(p0 uint64, p1 uint8) uint8 {
    for fi15 := 0; fi15 < 17; fi15++ {
        fmt.Println(p0)
        _ = fi15
    }
    v16 := []int{0, -156}
    _ = p0
    _ = p1
    _ = v16
    return 255
}

func main() {
    v1 := make(map[int]string)
    var v2 float32 = -2.25
    v3 := []int{127, -2147483648, -318}
    v2 -= v2
    v2 = v2
    var v17 int32 = -1
    v2 += -2.25
    v2 += (float32(((float64((v2 * float32(float64(0.0)))) + ((1.5 / float64(v2)) / (1.5 / 3.14159))) - 0.0)) - (-2.25 - v2))
    switch "\n\t" {
    case "hello":
        v1[969] = ("hello" + ("hello" + ("\n\t" + ("" + (("foo bar" + ("foo bar" + (("" + ("hello" + "\n\t")) + (("foo bar" + "a") + "foo bar")))) + "foo bar")))))
    case "\n\t":
        switch "hello" {
        case "":
            v17 += 2147483647
        case "foo bar":
            v18 := GenS1{F0: 0}
            _ = v18
        }
    case "foo bar":
        v2 = float32(0.0)
    }
    switch 4294967295 {
    case -156:
        _ = len(v3)
    }
    if len(v3) > 0 {
        v3[0] = 133
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        v20 := GenS2{F0: 0}
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        switch -1 {
        case -9223372036854775808:
            if true {
                v22 := v1[fi21]
                for fi23 := 0; fi23 < 10; fi23++ {
                    if (-1 < 100) {
                        if (genFn0(28, 42, 209) > 211) {
                            if true {
                                fmt.Println(fi23, fi21)
                                if len(v3) > 0 {
                                    v3[0] = (fi21 & (fi21 + 127))
                                }
                                v17 -= v17
                            } else {
                                var v24 int16 = 0
                                _ = v24
                            }
                            var v25 int8 = -128
                            _ = v25
                        }
                    } else {
                        _ = len(v1)
                    }
                    switch fi21 {
                    case -2147483648:
                        if (!(!false)) {
                            if ((!true) || false) {
                                delete(v1, 1)
                            }
                            fmt.Println(fi23)
                            fmt.Println(fi21, fi23, fi21)
                            v26 := GenS1{F0: 255}
                            _ = v26
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if true {
                                v27 := GenS3{F0: 128, F1: v17}
                                _ = v27
                            } else {
                                v28 := GenS1{F0: 32767}
                                _ = v28
                            }
                            fmt.Println(v22)
                        }
                    case 1:
                        fmt.Println(v22, v2, v22)
                    default:
                        v29 := GenS4{F0: 2147483647, F1: 2147483647, F2: 1}
                        _ = v29
                    }
                    _ = fi23
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v22
            }
        }
        fmt.Println(fi21, fi21, v17)
        _ = fi21
    }
    var v30 bool = (1 > 0)
    v30 = (!false)
    for fi31 := 0; fi31 < 5; fi31++ {
        delete(v1, (fi31 ^ -9223372036854775808))
        _ = fi31
    }
    fmt.Println(v17)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v30
}
