package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 int64; F1 uint8; F2 int }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 uint; F1 uint64 }
type GenS6 struct { F0 uint32; F1 bool }
type GenS7 struct { F0 uint32; F1 string }
type GenS8 struct { F0 int8; F1 uint }
func genFn0() uint64 {
    switch "a" {
    case "foo bar":
        switch "foo bar" {
        case "foo bar":
            var v7 int32 = -402
            _ = v7
        case "":
            if false {
                if (!(true && false)) {
                    var v8 uint64 = 42
                    var v9 float32 = (float32((3.14159 * (3.14159 + (float64(1e10) * 3.14159)))) + 0.0)
                    _ = v8
                    _ = v9
                } else {
                    v10 := []int{4294967295, 32767, -137}
                    v11 := GenS1{F0: 0, F1: 1}
                    _ = v10
                    _ = v11
                }
            }
        default:
            var v12 float32 = 1e10
            _ = v12
        }
    case "\n\t":
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 int = 885
            _ = fi13
            _ = v14
        }
    case "hello":
        for fi15 := 0; fi15 < 6; fi15++ {
            fmt.Println(fi15)
            _ = fi15
        }
    default:
        for fi16 := 0; fi16 < 10; fi16++ {
            fmt.Println(fi16)
            for fi17 := 0; fi17 < 5; fi17++ {
                var v18 int64 = -400
                _ = fi17
                _ = v18
            }
            for fi19 := 0; fi19 < 2; fi19++ {
                switch (fi16 % ((fi16) | 1)) {
                case -545:
                    var v20 float64 = float64((3.14159 * (-2.25 - 1e10)))
                    _ = v20
                case 9223372036854775807:
                    switch fi16 {
                    case -2147483648:
                        if (0.1 == (((0.1 - float32(0.0)) + float32(-2.25)) - float32(3.14159))) {
                            fmt.Println(fi19, fi19)
                            var v21 int32 = -649
                            _ = v21
                        } else {
                            v22 := []int{32767, 1, fi19}
                            var v23 uint32 = 0
                            var v24 bool = false
                            if false {
                                if false {
                                    if len(v22) >= 2 {
                                        v22 = v22[1:len(v22)]
                                    }
                                }
                                v24 = v24
                                if v24 {
                                    switch (-(fi19)) {
                                    case -639:
                                        if len(v22) > 0 {
                                            v22[0] = 1
                                        }
                                    case 2147483647:
                                        switch fi19 {
                                        case 42:
                                            v24 = (-2147483648 > 266)
                                        case 9223372036854775807:
                                            var v25 uint = 255
                                            _ = v25
                                        case -1:
                                            v24 = (91 <= 1)
                                        }
                                        switch ("a" + (("" + "") + "hello")) {
                                        case "a":
                                            v26 := GenS1{F0: 1, F1: 1}
                                            _ = v26
                                        }
                                        fmt.Println(v24, v23, v24)
                                        switch (fi16 - fi16) {
                                        case -1:
                                            v27 := GenS2{F0: 1, F1: 0}
                                            _ = v27
                                        case 32767:
                                            var v28 string = "foo bar"
                                            _ = v28
                                        }
                                        if (!v24) {
                                            var v29 bool = v24
                                            _ = v29
                                        } else {
                                            v24 = true
                                        }
                                        v23 -= v23
                                        v30 := GenS2{F0: 0, F1: 120}
                                        _ = v30
                                    case 404:
                                        v24 = v24
                                    }
                                    _ = len(v22)
                                    if true {
                                        fmt.Println(fi16)
                                    }
                                } else {
                                    if (fi16 != (^(fi16))) {
                                        _ = len(v22)
                                    }
                                }
                                var v31 int8 = -1
                                _ = v31
                            }
                            var v32 float32 = (((1.5 * -2.25) * float32(0.0)) / 1.0)
                            if true {
                                v33 := GenS3{F0: -9223372036854775808, F1: 211, F2: fi16}
                                _ = v33
                            } else {
                                if len(v22) >= 2 {
                                    v22 = v22[1:len(v22)]
                                }
                            }
                            _ = v22
                            _ = v23
                            _ = v24
                            _ = v32
                        }
                    case 2147483647:
                        v34 := GenS4{F0: int16(fi16), F1: 9223372036854775807}
                        _ = v34
                    case 42:
                        if true {
                            v35 := []int{42, (fi19 & 42)}
                            var v36 string = ""
                            _ = v35
                            _ = v36
                        }
                    }
                default:
                    switch fi19 {
                    case -532:
                        var v37 float64 = float64(((-2.25 + 0.0) / 3.14159))
                        _ = v37
                    case -9223372036854775808:
                        v38 := GenS3{F0: 0, F1: 255, F2: fi16}
                        _ = v38
                    case -2147483648:
                        var v39 uint64 = 0
                        _ = v39
                    }
                }
                if true {
                    var v40 bool = (244 > 32767)
                    _ = v40
                } else {
                    v41 := GenS5{F0: 65535, F1: 591}
                    var v42 int32 = 1
                    _ = v41
                    _ = v42
                }
                _ = fi19
            }
            _ = fi16
        }
    }
    switch -1 {
    case 32767:
        var v43 uint16 = 0
        _ = v43
    default:
        var v44 uint32 = 329
        _ = v44
    }
    return 393
}
func genFn1(p0 uint64) string {
    var v45 int64 = 769
    for fi46 := 0; fi46 < 16; fi46++ {
        switch 9223372036854775807 {
        case -245:
            var v47 string = "foo bar"
            _ = v47
        case 32767:
            for fi48 := 0; fi48 < 2; fi48++ {
                if ((((float64(float32(0.0)) - 1.5) / ((0.0 - 0.0) - 1.5)) / 1e10) <= (1.5 * 1e10)) {
                    if (571 >= 1) {
                        v49 := GenS6{F0: 311, F1: false}
                        if (!true) {
                            p0 = p0
                            v50 := GenS2{F0: 94, F1: 255}
                            var v51 int16 = 32767
                            _ = v50
                            _ = v51
                        } else {
                            var v52 bool = true
                            v52 = true
                            v53 := []int{(fi48 / (((fi46 & fi46)) | 1)), (^(fi46)), fi48, -9223372036854775808}
                            v53 = append(v53, (fi46 * fi46))
                            v54 := GenS6{F0: 4294967295, F1: false}
                            p0 += (p0 % ((genFn0()) | 1))
                            _ = v52
                            _ = v53
                            _ = v54
                        }
                        var v55 float32 = 0.0
                        var v56 uint8 = 255
                        _ = v49
                        _ = v55
                        _ = v56
                    }
                    if (("" + ("\n\t" + "\n\t")) <= "\n\t") {
                        var v57 int32 = -984
                        if (true && false) {
                            v45 = (v45 ^ -1)
                            var v58 uint = 637
                            _ = v58
                        } else {
                            p0 = (p0 * (p0 + genFn0()))
                            v45 = v45
                            fmt.Println(p0)
                        }
                        v59 := []int{60, fi46, 127}
                        _ = v57
                        _ = v59
                    }
                    switch fi48 {
                    case -406:
                        if false {
                            p0 += (^(p0))
                            if (false || false) {
                                v60 := GenS7{F0: uint32(p0), F1: ("a" + (("a" + ("a" + "\n\t")) + "hello"))}
                                _ = v60
                            }
                            p0 += (^(p0))
                        } else {
                            v61 := GenS2{F0: 20, F1: 429}
                            fmt.Println(fi48)
                            v62 := GenS5{F0: 291, F1: p0}
                            _ = v61
                            _ = v62
                        }
                    case 1:
                        fmt.Println(p0)
                    }
                } else {
                    var v63 string = "foo bar"
                    v64 := GenS1{F0: v45, F1: 0}
                    _ = v63
                    _ = v64
                }
                v45 -= (-(v45))
                fmt.Println(fi46)
                switch fi46 {
                case -2147483648:
                    v65 := []int{(-(fi48)), -492}
                    _ = v65
                default:
                    v66 := []int{970}
                    _ = v66
                }
                v45 -= (v45 / ((9223372036854775807) | 1))
                _ = fi48
            }
        }
        if false {
            p0 -= genFn0()
        }
        _ = fi46
    }
    _ = p0
    _ = v45
    return "foo bar"
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{"\n\t", ("foo bar" + ((("\n\t" + "hello") + ("hello" + (("\n\t" + "") + ""))) + ("foo bar" + "\n\t"))), "foo bar", "\n\t"}
    v3 := make(map[string]int)
    var v4 int32 = 1
    var v5 int8 = 1
    var v6 uint32 = 2147483647
    for fi67 := 0; fi67 < 19; fi67++ {
        fmt.Println(fi67, v5)
        v68 := GenS4{F0: -865, F1: -577}
        _ = fi67
        _ = v68
    }
    v69 := GenS7{F0: v6, F1: ("" + "\n\t")}
    var v70 uint64 = 4294967295
    switch 32767 {
    case -36:
        if false {
            for fi71 := 0; fi71 < 13; fi71++ {
                var v72 uint8 = 42
                v73 := GenS8{F0: v5, F1: 18446744073709551615}
                var v74 float64 = float64(0.0)
                _ = fi71
                _ = v72
                _ = v73
                _ = v74
            }
            v5 = v5
            var v75 int64 = 1
            _ = v75
        }
    }
    var v76 int = -1
    v77 := v3["\n\t"]
    v76 -= v77
    _ = len(v1)
    var v78 int16 = 1000
    var v79 bool = false
    switch (v76 & int(v70)) {
    case 1:
        var v80 uint64 = (v70 << ((255) & 63))
        _ = v80
    case 32767:
        switch v76 {
        case 4294967295:
            v78 -= v78
        default:
            for fi81 := 0; fi81 < 13; fi81++ {
                for fi82 := 0; fi82 < 13; fi82++ {
                    fmt.Println(fi81, v70, v5)
                    delete(v1, (fi82 % (((v77 - (fi81 ^ -926))) | 1)))
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi82
                }
                v79 = (!(v79 && (43 != 2)))
                _ = fi81
            }
        }
    }
    v5 += v5
    v78 = 0
    for fi83 := 0; fi83 < 15; fi83++ {
        _ = len(v3)
        var v84 int64 = -9223372036854775808
        var v85 uint8 = 128
        _ = fi83
        _ = v84
        _ = v85
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v70
    _ = v76
    _ = v77
    _ = v78
    _ = v79
}
