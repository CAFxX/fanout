package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 int16; F2 string }
func genFn0(p0 uint64, p1 string, p2 uint16) uint32 {
    p1 = "\n\t"
    var v4 uint64 = p0
    if false {
        v5 := []int{-321}
        fmt.Println(v4, v4, p1)
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}
func genFn1(p0 int16, p1 int8) string {
    v6 := GenS1{F0: 42, F1: 18446744073709551615, F2: -463}
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            if ((!true) && true) {
                switch (fi8 % ((fi8) | 1)) {
                case -2147483648:
                    switch ("foo bar" + "foo bar") {
                    case "\n\t":
                        v9 := []int{fi8, (-(fi8))}
                        _ = v9
                    case "hello":
                        v10 := []int{}
                        _ = v10
                    default:
                        fmt.Println(fi8, fi7)
                    }
                case 2147483647:
                    var v11 float64 = 0.1
                    _ = v11
                default:
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(fi7, fi7, p1)
                    default:
                        if false {
                            v12 := []int{(fi8 - (fi7 & (fi7 * (^(fi8))))), fi7, (fi8 * (fi8 ^ fi7)), 0}
                            fmt.Println(fi8, fi8)
                            _ = v12
                        } else {
                            p1 = 26
                            var v13 uint = 939
                            _ = v13
                        }
                    }
                }
                switch (fi7 % (((fi8 ^ fi7)) | 1)) {
                case 792:
                    if false {
                        if (!(!(true && (false && true)))) {
                            p0 = (p0 / ((p0) | 1))
                            p1 = int8(fi8)
                        }
                        fmt.Println(fi8, p0)
                        var v14 int = fi8
                        _ = v14
                    }
                case 61:
                    if true {
                        p0 += (p0 >> ((65535) & 63))
                    } else {
                        switch fi7 {
                        case 128:
                            var v15 uint32 = 0
                            _ = v15
                        case -623:
                            v16 := []int{fi8}
                            _ = v16
                        case 127:
                            p1 += p1
                        default:
                            p0 += (p0 >> ((4294967295) & 63))
                        }
                        p1 = p1
                        v17 := []int{}
                        _ = v17
                    }
                case -957:
                    v18 := []int{(fi7 >> ((18446744073709551615) & 63)), (fi7 ^ 32767)}
                    _ = v18
                default:
                    if (false && true) {
                        var v19 uint = 0
                        var v20 float64 = (3.14159 * -2.25)
                        _ = v19
                        _ = v20
                    } else {
                        if true {
                            var v21 int64 = -9223372036854775808
                            _ = v21
                        } else {
                            p1 = p1
                            var v22 int32 = 1
                            var v23 int16 = p0
                            _ = v22
                            _ = v23
                        }
                        v24 := []int{(fi8 / ((fi7) | 1)), fi8, -1}
                        _ = len(v24)
                        _ = v24
                    }
                }
                fmt.Println(p0, p1, fi8)
                v25 := []int{fi8}
                _ = v25
            } else {
                var v26 uint16 = 32767
                _ = v26
            }
            p1 -= (-(p1))
            _ = fi8
        }
        p1 = (p1 ^ (p1 << ((255) & 63)))
        _ = fi7
    }
    p1 -= int8(p0)
    _ = p0
    _ = p1
    _ = v6
    return (("" + "\n\t") + "a")
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 117
    v3 := []uint32{295, 400, 2147483648, 96}
    v2 += (v2 / ((v2) | 1))
    var v27 float32 = (float32(0.0) - -2.25)
    var v28 int = 4294967295
    v27 = float32((3.14159 / 1e10))
    v2 += v2
    var v29 bool = true
    if (32767 == -1) {
        fmt.Println(v2)
        fmt.Println(v28, v29)
        var v30 uint16 = 42
        _ = v30
    } else {
        var v31 uint8 = (v2 >> ((42) & 63))
        for fi32 := 0; fi32 < 16; fi32++ {
            if (v27 == (v27 / (v27 - 1.5))) {
                for fi33 := 0; fi33 < 11; fi33++ {
                    var v34 int8 = -128
                    _ = fi33
                    _ = v34
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                var v35 bool = ("" == ("" + "\n\t"))
                for fi36 := 0; fi36 < 5; fi36++ {
                    var v37 bool = true
                    _ = fi36
                    _ = v37
                }
                _ = v35
            }
            _ = fi32
        }
        if v29 {
            v3 = append(v3, 42)
            var v38 uint16 = 0
            v1[354] = "foo bar"
            _ = v38
        } else {
            v2 = v31
        }
        _ = v31
    }
    v27 = (3.14159 / v27)
    var v39 int32 = -582
    if len(v3) > 0 {
        v3[0] = 786
    }
    for fi40 := 0; fi40 < 12; fi40++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v3 = append(v3, 2)
        v28 = fi40
        _ = fi40
    }
    v3 = append(v3, 2)
    fmt.Println(v28, v29)
    v29 = v29
    var v41 uint64 = 18446744073709551615
    switch "\n\t" {
    case "":
        fmt.Println(v29, v29, v29)
    }
    if false {
        if (100 > -128) {
            var v42 uint16 = 65535
            var v43 uint8 = (v2 / (((v2 >> ((431) & 63))) | 1))
            var v44 float64 = (((float64(v27) + (float64(v27) - ((((1e10 / 0.1) * 0.0) / float64(((v27 * v27) - v27))) - (1e10 / 3.14159)))) / float64(v27)) + 1.5)
            _ = v42
            _ = v43
            _ = v44
        }
        var v45 uint = 65535
        if (0 != 1) {
            fmt.Println(v27, v29, v27)
            fmt.Println(v39)
        } else {
            switch v28 {
            case 317:
                if false {
                    if (v41 != v41) {
                        v46 := GenS2{F0: ""}
                        _ = v46
                    } else {
                        if len(v3) > 0 {
                            v3[0] = 570
                        }
                    }
                    var v47 int16 = 915
                    fmt.Println(v45, v29, v27)
                    _ = v47
                }
            case -574:
                var v48 int16 = 123
                _ = v48
            case 323:
                if (v29 || (!v29)) {
                    v28 += 128
                    var v49 uint16 = 393
                    _ = v49
                } else {
                    if (!v29) {
                        if (-724 != 0) {
                            v27 += (v27 - 0.1)
                            v3 = append(v3, 42)
                        } else {
                            if v29 {
                                if len(v3) > 0 {
                                    v3[0] = 4294967295
                                }
                            } else {
                                for fi50 := 0; fi50 < 6; fi50++ {
                                    var v51 int64 = -9223372036854775808
                                    var v52 uint16 = 2
                                    _ = fi50
                                    _ = v51
                                    _ = v52
                                }
                            }
                        }
                    } else {
                        fmt.Println(v39, v41, v2)
                        var v53 int = 1
                        fmt.Println(v39)
                        _ = v53
                    }
                }
            default:
                v54 := GenS2{F0: "foo bar"}
                _ = v54
            }
            v55 := GenS3{F0: v39, F1: 1, F2: "\n\t"}
            var v56 int = (v28 / ((v28) | 1))
            _ = v55
            _ = v56
        }
        _ = v45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v39
    _ = v41
}
