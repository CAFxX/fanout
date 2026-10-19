package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint; F2 int8 }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 int32; F1 string; F2 uint }
type GenS5 struct { F0 int64; F1 uint }
func genFn0() uint8 {
    var v7 float64 = 1e10
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 int = fi8
        switch 32768 {
        case -9223372036854775808:
            for fi10 := 0; fi10 < 17; fi10++ {
                v11 := GenS1{F0: 228}
                var v12 uint = 491
                switch "a" {
                case "\n\t":
                    v12 -= (v12 - 699)
                case "":
                    v12 = 51
                case "a":
                    fmt.Println(fi10, fi10, v12)
                default:
                    var v13 float64 = 1e10
                    _ = v13
                }
                if ("foo bar" == ("\n\t" + ("a" + "foo bar"))) {
                    v14 := GenS2{F0: 2147483648, F1: (^(v12)), F2: 0}
                    var v15 bool = true
                    v15 = (2147483647 != 65535)
                    _ = v14
                    _ = v15
                } else {
                    switch fi10 {
                    case 0:
                        fmt.Println(v12, v7)
                    case -762:
                        if false {
                            v12 = (v12 >> ((v12) & 63))
                            fmt.Println(v9, v9, v12)
                            v7 += 0.0
                            if false {
                                if true {
                                    switch fi10 {
                                    case 1:
                                        v16 := []int{}
                                        _ = v16
                                    }
                                    v12 = v12
                                    if (!(853 <= 32767)) {
                                        v17 := GenS1{F0: 1000}
                                        var v18 uint32 = 59
                                        fmt.Println(fi8)
                                        var v19 uint = (v12 << (((v12 ^ v12)) & 63))
                                        fmt.Println(v18, v7, v12)
                                        v20 := []int{}
                                        _ = v17
                                        _ = v18
                                        _ = v19
                                        _ = v20
                                    }
                                } else {
                                    v21 := GenS1{F0: 1000}
                                    v22 := []int{42, 1, fi10, (fi8 - fi8)}
                                    _ = v21
                                    _ = v22
                                }
                                var v23 string = "foo bar"
                                v23 = (((v23 + (v23 + v23)) + v23) + ("a" + v23))
                                _ = v23
                            } else {
                                v9 -= -841
                            }
                            var v24 int32 = 251
                            var v25 float64 = (v7 * float64((0.1 + (float32(((v7 + v7) - v7)) / float32(1e10)))))
                            var v26 uint32 = 557
                            fmt.Println(v24, v24)
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                    default:
                        switch fi8 {
                        case 32767:
                            switch 820 {
                            case 2147483647:
                                v12 -= (v12 & (v12 / ((v12) | 1)))
                            case -2147483648:
                                v27 := []int{}
                                v12 = 0
                                _ = v27
                            default:
                                var v28 uint = 830
                                _ = v28
                            }
                        case 2147483647:
                            v29 := []int{}
                            _ = v29
                        }
                    }
                    v30 := GenS3{F0: 255, F1: 954}
                    fmt.Println(v12)
                    switch 127 {
                    case 426:
                        if false {
                            if true {
                                v31 := GenS1{F0: -744}
                                _ = v31
                            }
                            var v32 float32 = float32((v7 + v7))
                            _ = v32
                        }
                    }
                    _ = v30
                }
                _ = fi10
                _ = v11
                _ = v12
            }
        case 717:
            v9 = (v9 * v9)
        case 4294967295:
            if true {
                var v33 uint8 = 255
                _ = v33
            }
        default:
            v9 = (fi8 * v9)
        }
        _ = fi8
        _ = v9
    }
    if false {
        v34 := []int{-845}
        v7 = v7
        var v35 int16 = 590
        _ = v34
        _ = v35
    }
    if (255 >= 314) {
        v7 = 3.14159
        var v36 float64 = ((v7 * float64((1e10 * (-2.25 + 3.14159)))) * 0.0)
        for fi37 := 0; fi37 < 13; fi37++ {
            v36 = v36
            v38 := []int{fi37}
            v39 := GenS1{F0: -32768}
            _ = fi37
            _ = v38
            _ = v39
        }
        _ = v36
    }
    for fi40 := 0; fi40 < 13; fi40++ {
        v41 := GenS4{F0: -767, F1: "a", F2: 468}
        if (v7 <= 3.14159) {
            v42 := []int{4294967295, 889}
            _ = v42
        }
        var v43 string = ("foo bar" + "a")
        _ = fi40
        _ = v41
        _ = v43
    }
    _ = v7
    return 84
}
func genFn1(p0 string, p1 int8, p2 uint32) int16 {
    v44 := []int{128, 557}
    v45 := GenS1{F0: 0}
    p1 = (p1 / ((p1) | 1))
    var v46 uint16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v44
    _ = v45
    _ = v46
    return -671
}

func main() {
    v1 := []int{}
    var v2 uint32 = 629
    v3 := []uint32{}
    v4 := []int{}
    v5 := make(map[string]uint32)
    v6 := []uint32{(v2 << ((65535) & 63))}
    var v47 int = 42
    v48 := GenS1{F0: genFn1(("\n\t" + ""), -54, (v2 | (v2 >> ((18446744073709551615) & 63))))}
    var v49 uint8 = 58
    switch (v47 - v47) {
    case 9223372036854775807:
        if true {
            for fi50 := 0; fi50 < 8; fi50++ {
                v47 = fi50
                v49 += (v49 << ((1) & 63))
                _ = fi50
            }
            var v51 int = (v47 | (v47 & (v47 - v47)))
            _ = v51
        }
    case 1:
        if len(v6) > 0 {
            v6[0] = (v2 % (((v2 - (v2 + (v2 + v2)))) | 1))
        }
    case -819:
        v52 := GenS2{F0: v2, F1: 0, F2: -128}
        _ = v52
    default:
        v53 := GenS5{F0: 969, F1: 42}
        _ = v53
    }
    v54 := GenS3{F0: 32767, F1: -228}
    fmt.Println(v49)
    if ((!(!(!true))) && true) {
        var v55 bool = true
        _ = v55
    } else {
        switch int(v49) {
        case 9223372036854775807:
            switch (v47 | v47) {
            case 1:
                for fi56 := 0; fi56 < 14; fi56++ {
                    switch fi56 {
                    case -1:
                        fmt.Println(v49, v2)
                    default:
                        if (255 != 255) {
                            v47 = (v47 % ((-908) | 1))
                        } else {
                            if ((v49 ^ 155) != (v49 % ((v49) | 1))) {
                                fmt.Println(v2, v2)
                                fmt.Println(fi56)
                            } else {
                                switch ("a" + "foo bar") {
                                case "\n\t":
                                    switch "foo bar" {
                                    case "foo bar":
                                        for fi57 := 0; fi57 < 4; fi57++ {
                                            if len(v4) > 0 {
                                                v4[0] = (fi56 % ((fi56) | 1))
                                            }
                                            _ = fi57
                                        }
                                    case "\n\t":
                                        var v58 string = ""
                                        _ = v58
                                    case "":
                                        v59 := GenS5{F0: -9223372036854775808, F1: 243}
                                        _ = v59
                                    default:
                                        var v60 float32 = -2.25
                                        _ = v60
                                    }
                                case "":
                                    var v61 float32 = 1.5
                                    _ = v61
                                case "hello":
                                    if true {
                                        delete(v5, "foo bar")
                                        v47 += fi56
                                    }
                                default:
                                    switch fi56 {
                                    case -1:
                                        v47 += (-(v47))
                                    case -505:
                                        _ = len(v3)
                                    case 42:
                                        if false {
                                            fmt.Println(v49, fi56)
                                            switch "\n\t" {
                                            case "":
                                                v47 = fi56
                                            case "foo bar":
                                                var v62 string = ("a" + "hello")
                                                _ = v62
                                            case "a":
                                                var v63 bool = true
                                                _ = v63
                                            }
                                        }
                                    default:
                                        for fi64 := 0; fi64 < 14; fi64++ {
                                            var v65 uint16 = 42
                                            var v66 int16 = 18
                                            _ = fi64
                                            _ = v65
                                            _ = v66
                                        }
                                    }
                                }
                                for fi67 := 0; fi67 < 3; fi67++ {
                                    v2 = uint32(v49)
                                    var v68 int8 = 42
                                    var v69 int = v47
                                    _ = fi67
                                    _ = v68
                                    _ = v69
                                }
                            }
                        }
                    }
                    _ = fi56
                }
            case -609:
                switch v47 {
                case 32768:
                    var v70 int32 = 1
                    _ = v70
                default:
                    var v71 string = "hello"
                    _ = v71
                }
            case 774:
                fmt.Println(v49, v2, v49)
            default:
                v2 -= v2
            }
        default:
            if true {
                fmt.Println(v2)
                _ = len(v5)
            }
        }
        v72 := GenS3{F0: 991, F1: v47}
        switch (v47 - v47) {
        case -2147483648:
            var v73 uint32 = (v2 ^ (v2 - uint32(v49)))
            _ = v73
        case 0:
            var v74 int = v47
            _ = v74
        case 300:
            if (genFn1("a", 100, (v2 + 4294967295)) > 1000) {
                fmt.Println(v2, v47, v49)
                if len(v6) > 0 {
                    v6[0] = v2
                }
            }
        default:
            _ = len(v6)
        }
        _ = v72
    }
    var v75 int8 = 100
    v2 += v2
    var v76 int16 = -1
    v77 := GenS4{F0: 42, F1: ("" + "foo bar"), F2: 65535}
    var v78 int = (v47 * v47)
    v78 = 33
    switch ("" + "") {
    case "hello":
        if (65535 == 4294967295) {
            v79 := GenS3{F0: 993, F1: (v47 * (v47 / ((128) | 1)))}
            _ = v79
        }
    case "\n\t":
        if true {
            v80 := GenS5{F0: 1, F1: 294}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v80
        }
    default:
        for fi81 := 0; fi81 < 14; fi81++ {
            v2 = 781
            for fi82 := 0; fi82 < 14; fi82++ {
                var v83 int64 = 0
                _ = fi82
                _ = v83
            }
            _ = fi81
        }
    }
    var v84 uint64 = 286
    var v85 string = ""
    v49 += (v49 % (((^(v49))) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v78:", v78)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v54
    _ = v6
    _ = v75
    _ = v76
    _ = v77
    _ = v78
    _ = v84
    _ = v85
}
