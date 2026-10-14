package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32; F1 uint }
func genFn0(p0 int32, p1 int64, p2 int) uint16 {
    p0 = p0
    p2 = (p2 >> ((255) & 63))
    for fi4 := 0; fi4 < 14; fi4++ {
        v5 := GenS1{F0: (p1 * 0)}
        p2 = fi4
        _ = fi4
        _ = v5
    }
    v6 := []int{-29}
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 int = p2
            v10 := GenS1{F0: p1}
            if false {
                if len(v6) > 0 {
                    v6[0] = (v9 | v9)
                }
                v11 := GenS2{F0: 0, F1: 65535}
                _ = v11
            } else {
                fmt.Println(v9)
                v12 := GenS3{F0: 924, F1: ""}
                var v13 int8 = 1
                _ = v12
                _ = v13
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}

func main() {
    var v1 uint = 255
    v2 := []uint32{316, 2147483648, 0}
    var v3 int64 = 42
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    v14 := GenS4{F0: v1}
    var v15 uint8 = 128
    var v16 float32 = (1e10 - 1e10)
    if true {
        if true {
            v17 := GenS3{F0: 130, F1: ""}
            var v18 bool = true
            v19 := GenS3{F0: 0, F1: (("foo bar" + "") + "foo bar")}
            _ = v17
            _ = v18
            _ = v19
        } else {
            for fi20 := 0; fi20 < 19; fi20++ {
                if false {
                    var v21 string = "foo bar"
                    _ = v21
                }
                fmt.Println(v1, fi20)
                switch 32768 {
                case 4294967295:
                    var v22 int = (fi20 ^ (-(fi20)))
                    _ = v22
                case 128:
                    for fi23 := 0; fi23 < 14; fi23++ {
                        if len(v2) > 0 {
                            v2[0] = 2
                        }
                        var v24 int8 = 100
                        var v25 uint8 = v15
                        if true {
                            fmt.Println(fi20)
                        }
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                case -345:
                    if ((1 < v3) || (((^(v1)) != v1) && true)) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        fmt.Println(v1)
                        if (!(-881 <= 446)) {
                            fmt.Println(v15)
                            var v26 bool = (v16 > v16)
                            _ = v26
                        } else {
                            v27 := GenS1{F0: (v3 % ((1) | 1))}
                            var v28 float32 = v16
                            _ = v27
                            _ = v28
                        }
                        v29 := GenS1{F0: v3}
                        _ = v29
                    }
                }
                _ = fi20
            }
        }
        if true {
            v16 = v16
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi30 := 0; fi30 < 16; fi30++ {
                switch fi30 {
                case 778:
                    for fi31 := 0; fi31 < 13; fi31++ {
                        var v32 float64 = (3.14159 * (float64(3.14159) / 1e10))
                        if false {
                            var v33 uint8 = v15
                            _ = v33
                        }
                        _ = fi31
                        _ = v32
                    }
                case -1:
                    var v34 int32 = 0
                    _ = v34
                default:
                    v3 = -1
                }
                if true {
                    if (2 > 65535) {
                        for fi35 := 0; fi35 < 4; fi35++ {
                            v2 = append(v2, 0)
                            _ = fi35
                        }
                        for fi36 := 0; fi36 < 12; fi36++ {
                            if (!true) {
                                v15 += 255
                                v16 = float32(((float64(((v16 * (v16 + ((v16 + v16) * v16))) + v16)) * (-2.25 * 3.14159)) + (1.5 * 3.14159)))
                                var v37 uint32 = 42
                                _ = v37
                            } else {
                                v2 = append(v2, 8)
                                if true {
                                    v38 := GenS2{F0: 409, F1: 798}
                                    _ = v38
                                }
                            }
                            v2 = append(v2, 475)
                            _ = fi36
                        }
                        switch (("hello" + (("\n\t" + "\n\t") + "hello")) + "a") {
                        case "foo bar":
                            for fi39 := 0; fi39 < 14; fi39++ {
                                if (0 <= 42) {
                                    switch fi30 {
                                    case -2147483648:
                                        if false {
                                            v2 = append(v2, 361)
                                            var v40 int64 = (v3 * (v3 << (((v1 | (v1 / (((v1 << ((v1) & 63))) | 1)))) & 63)))
                                            v1 = 0
                                            _ = v40
                                        }
                                    case -162:
                                        fmt.Println(fi30, fi39)
                                    default:
                                        if (2147483647 > -2147483648) {
                                            var v41 int16 = -32768
                                            _ = v41
                                        }
                                    }
                                    v2 = append(v2, 0)
                                }
                                var v42 uint = v1
                                _ = fi39
                                _ = v42
                            }
                        case "":
                            fmt.Println(v16)
                        case "a":
                            var v43 uint64 = 0
                            _ = v43
                        default:
                            for fi44 := 0; fi44 < 14; fi44++ {
                                var v45 int8 = -1
                                v1 += v1
                                _ = fi44
                                _ = v45
                            }
                        }
                    }
                    if (v1 >= uint(v15)) {
                        var v46 float32 = (float32(float64((v16 + (((v16 / 1.0) * 1.5) + float32(float64((v16 * (0.1 - 3.14159)))))))) / ((0.1 - (((v16 * ((v16 + v16) * v16)) * v16) / (1e10 - v16))) - v16))
                        for fi47 := 0; fi47 < 16; fi47++ {
                            switch (("\n\t" + ("hello" + "hello")) + "foo bar") {
                            case "hello":
                                v48 := GenS5{F0: (v15 & 0)}
                                _ = v48
                            case "\n\t":
                                v46 += 3.14159
                            case "":
                                var v49 float64 = 1.5
                                _ = v49
                            default:
                                if (1 < 1) {
                                    var v50 uint8 = v15
                                    _ = len(v2)
                                    var v51 uint64 = 18446744073709551615
                                    _ = v50
                                    _ = v51
                                }
                            }
                            _ = fi47
                        }
                        _ = v46
                    }
                    v3 = -1
                } else {
                    v16 += 1e10
                }
                v1 = v1
                _ = fi30
            }
        }
    }
    if (1000 <= -1) {
        var v52 int8 = 100
        v52 = (v52 / (((v52 << (((v1 & 549)) & 63))) | 1))
        _ = v52
    } else {
        var v53 int8 = -128
        v15 -= (v15 & (v15 / ((v15) | 1)))
        _ = v53
    }
    var v54 uint = (v1 << (((^(v1))) & 63))
    var v55 uint = v54
    var v56 string = ("hello" + ("" + "foo bar"))
    v1 = 42
    var v57 string = v56
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v2 = append(v2, 0)
    if len(v2) > 0 {
        v2[0] = 212
    }
    if true {
        if (uint64(v15) >= 4294967295) {
            v16 = -2.25
            var v58 int32 = -2147483648
            _ = v58
        } else {
            fmt.Println(v3, v1)
            if false {
                switch "\n\t" {
                case "a":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case "foo bar":
                    v59 := GenS6{F0: 1, F1: 4294967295}
                    _ = v59
                }
                v1 = v55
            }
            v3 += (v3 << ((2) & 63))
        }
    }
    v3 += 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v54
    _ = v55
    _ = v56
    _ = v57
}
