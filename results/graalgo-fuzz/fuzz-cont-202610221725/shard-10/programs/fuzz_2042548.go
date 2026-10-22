package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int16 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 uint8; F1 string; F2 uint16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS8 struct { F0 uint16; F1 uint32; F2 int }
type GenS9 struct { F0 int; F1 int64; F2 int8 }
func genFn0(p0 int64, p1 int, p2 int64) int32 {
    v7 := []int{p1}
    if len(v7) > 0 {
        v7[0] = p1
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 int64 = (p0 | 9223372036854775807)
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 249
}
func genFn1(p0 uint) bool {
    switch 128 {
    case 128:
        p0 -= p0
    case 32767:
        for fi10 := 0; fi10 < 19; fi10++ {
            p0 -= (p0 / (((p0 - 2)) | 1))
            v11 := GenS1{F0: 32767, F1: -100, F2: 42}
            _ = fi10
            _ = v11
        }
    default:
        if true {
            p0 = p0
            v12 := GenS2{F0: 1, F1: 255}
            var v13 uint8 = 127
            _ = v12
            _ = v13
        } else {
            var v14 int16 = -32768
            v15 := GenS2{F0: 2147483647, F1: 65535}
            for fi16 := 0; fi16 < 6; fi16++ {
                v17 := []int{(^(fi16)), fi16}
                v18 := GenS3{F0: genFn0(0, fi16, -1), F1: fi16}
                fmt.Println(fi16, p0)
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = v14
            _ = v15
        }
    }
    v19 := GenS1{F0: 255, F1: 127, F2: 336}
    if true {
        v20 := []int{-1, 1, 128, 1}
        _ = v20
    }
    _ = p0
    _ = v19
    return true
}

func main() {
    var v1 uint32 = 869
    var v2 int = 42
    v3 := make(map[string]string)
    v4 := make(map[string]string)
    var v5 int8 = 110
    v6 := make(map[int]uint32)
    var v21 bool = false
    var v22 uint64 = 0
    var v23 bool = v21
    switch (v2 | int(v5)) {
    case 4294967295:
        v24 := GenS2{F0: -2147483648, F1: 32767}
        _ = v24
    case 2147483647:
        switch v2 {
        case -936:
            switch v2 {
            case -563:
                for fi25 := 0; fi25 < 19; fi25++ {
                    v26 := GenS1{F0: 1, F1: -6, F2: 767}
                    _ = fi25
                    _ = v26
                }
            default:
                fmt.Println(v5, v2)
            }
        }
    case 32768:
        var v27 uint = 42
        _ = v27
    default:
        var v28 int64 = 0
        _ = v28
    }
    v2 -= v2
    v29 := GenS4{F0: 836, F1: 880}
    fmt.Println(v21, v2, v22)
    if v21 {
        switch ("hello" + "foo bar") {
        case "\n\t":
            v30 := []int{1, (^(v2))}
            _ = v30
        case "foo bar":
            for fi31 := 0; fi31 < 7; fi31++ {
                if v21 {
                    v32 := []int{(v2 ^ fi31), -9223372036854775808, (fi31 | (v2 >> ((200) & 63))), v2}
                    switch "\n\t" {
                    case "hello":
                        if (-2.25 > ((0.1 * 0.1) * (1e10 + 3.14159))) {
                            var v33 bool = (!(621 > 4294967295))
                            var v34 string = ("a" + "foo bar")
                            _ = v33
                            _ = v34
                        }
                    case "":
                        var v35 float64 = ((-2.25 + -2.25) * (-2.25 / 1.0))
                        _ = v35
                    }
                    fmt.Println(v1)
                    _ = v32
                } else {
                    v5 += v5
                    for fi36 := 0; fi36 < 16; fi36++ {
                        v2 -= 128
                        _ = fi36
                    }
                    v2 += 32767
                }
                v2 = (v2 / ((4294967295) | 1))
                v5 += (v5 + v5)
                _ = fi31
            }
        case "a":
            v21 = (4294967295 != (v2 + -2147483648))
        default:
            var v37 uint64 = 794
            _ = v37
        }
    } else {
        var v38 float32 = (1.5 - 0.1)
        _ = v38
    }
    _ = len(v3)
    if (1.5 != -2.25) {
        v39 := GenS5{F0: uint8(v1), F1: "", F2: 520}
        _ = v39
    } else {
        if (v1 <= (^(v1))) {
            for fi40 := 0; fi40 < 6; fi40++ {
                var v41 int16 = 1
                var v42 uint = 0
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(v22)
                case "":
                    if v21 {
                        v43 := []int{}
                        var v44 float32 = (1e10 - 1.5)
                        fmt.Println(v2)
                        _ = v43
                        _ = v44
                    } else {
                        if (!v23) {
                            if v21 {
                                fmt.Println(v22, v2, v41)
                                if false {
                                    var v45 uint64 = 42
                                    v45 += 4294967295
                                    var v46 uint64 = uint64(v1)
                                    _ = v45
                                    _ = v46
                                }
                                for fi47 := 0; fi47 < 20; fi47++ {
                                    var v48 int8 = (v5 / (((v5 & v5)) | 1))
                                    var v49 float64 = 1.5
                                    v50 := []int{(-(fi47)), 0}
                                    var v51 int16 = int16(v48)
                                    _ = fi47
                                    _ = v48
                                    _ = v49
                                    _ = v50
                                    _ = v51
                                }
                            }
                        } else {
                            v52 := v3[("foo bar" + "a")]
                            if v21 {
                                v53 := GenS4{F0: 42, F1: v22}
                                if (-585 == 1) {
                                    delete(v4, (v52 + (("" + "\n\t") + (v52 + (("foo bar" + "a") + v52)))))
                                    var v54 uint = (v42 + (v42 / (((^(v42))) | 1)))
                                    _ = v54
                                }
                                if v23 {
                                    v55 := []int{(v2 & 42), fi40}
                                    fmt.Println(v42, v2, v23)
                                    _ = v55
                                } else {
                                    v56 := v6[fi40]
                                    switch 2147483647 {
                                    case 42:
                                        var v57 int64 = -162
                                        _ = v57
                                    case 0:
                                        if true {
                                            if v23 {
                                                v21 = v21
                                                v22 = 18446744073709551615
                                                v58 := GenS1{F0: 142, F1: (v5 + 40), F2: 528}
                                                _ = v58
                                            }
                                            var v59 int64 = 0
                                            _ = v59
                                        } else {
                                            if (!v21) {
                                                if false {
                                                    for fi60 := 0; fi60 < 11; fi60++ {
                                                        if (v21 || (false && (!(-2147483648 == -2147483648)))) {
                                                            v61 := v4["foo bar"]
                                                            _ = v61
                                                        }
                                                        _ = fi60
                                                    }
                                                    var v62 uint = v42
                                                    var v63 int = (v2 >> (((v62 >> ((v42) & 63))) & 63))
                                                    _ = v62
                                                    _ = v63
                                                }
                                                var v64 float64 = float64(float32(-2.25))
                                                _ = v64
                                            }
                                            delete(v6, 422)
                                        }
                                    default:
                                        v1 -= 0
                                    }
                                    v65 := GenS1{F0: 32767, F1: (v5 / ((int8(v22)) | 1)), F2: 126}
                                    _ = v56
                                    _ = v65
                                }
                                _ = v53
                            }
                            _ = v52
                        }
                        v66 := []int{int(v5), (fi40 % ((42) | 1)), 9223372036854775807}
                        switch fi40 {
                        case 128:
                            switch ("foo bar" + ("foo bar" + (("\n\t" + "\n\t") + "hello"))) {
                            case "hello":
                                if (!(2 >= 223)) {
                                    v2 -= v2
                                }
                            }
                        case 42:
                            fmt.Println(v2)
                        case 2147483647:
                            v1 = v1
                        default:
                            v66 = append(v66, 843)
                        }
                        _ = v66
                    }
                case "foo bar":
                    for fi67 := 0; fi67 < 6; fi67++ {
                        v68 := GenS4{F0: -377, F1: 5}
                        _ = fi67
                        _ = v68
                    }
                default:
                    fmt.Println(v5, v2)
                }
                _ = fi40
                _ = v41
                _ = v42
            }
        } else {
            v69 := GenS6{F0: -132}
            fmt.Println(v21, v22, v22)
            if true {
                fmt.Println(v5, v21)
            }
            _ = v69
        }
        v23 = true
        var v70 uint16 = 65535
        _ = v70
    }
    for fi71 := 0; fi71 < 3; fi71++ {
        v5 = (v5 >> ((255) & 63))
        _ = fi71
    }
    switch v2 {
    case -9223372036854775808:
        v2 += (v2 << ((1) & 63))
    }
    if v21 {
        if v21 {
            var v72 uint8 = 42
            _ = v72
        }
        var v73 float32 = (float32(1.5) * (1e10 + 3.14159))
        _ = v73
    }
    switch v2 {
    case -534:
        switch v2 {
        case -2147483648:
            for fi74 := 0; fi74 < 2; fi74++ {
                v23 = v23
                if ("foo bar" != "hello") {
                    var v75 uint32 = v1
                    fmt.Println(v2)
                    _ = v75
                } else {
                    var v76 uint = 42
                    var v77 float32 = float32(1e10)
                    fmt.Println(v22, v21)
                    _ = v76
                    _ = v77
                }
                for fi78 := 0; fi78 < 20; fi78++ {
                    switch (fi78 << ((4294967295) & 63)) {
                    case 872:
                        if false {
                            v21 = true
                            switch v2 {
                            case -1:
                                var v79 uint8 = 2
                                _ = v79
                            case 32767:
                                switch 744 {
                                case -9223372036854775808:
                                    v1 -= 680
                                case 4294967295:
                                    var v80 uint16 = 2
                                    _ = v80
                                case 495:
                                    v81 := []int{(fi78 + v2)}
                                    _ = v81
                                default:
                                    if true {
                                        _ = len(v6)
                                        v2 = v2
                                    } else {
                                        var v82 int = (fi74 + fi74)
                                        v1 -= (v1 / (((v1 + v1)) | 1))
                                        v83 := GenS7{F0: 166, F1: 127, F2: -733}
                                        if true {
                                            v2 = -473
                                            var v84 uint32 = 711
                                            _ = v84
                                        }
                                        var v85 bool = false
                                        _ = v82
                                        _ = v83
                                        _ = v85
                                    }
                                }
                            }
                            var v86 uint8 = 107
                            _ = v86
                        } else {
                            _ = len(v4)
                        }
                    case 790:
                        if (true && ((v2 & (-(fi78))) <= int(v1))) {
                            v87 := []int{}
                            _ = v87
                        }
                        v88 := GenS8{F0: 744, F1: 680, F2: (v2 + v2)}
                        _ = v88
                    }
                    v89 := GenS9{F0: v2, F1: -494, F2: (v5 + (-(v5)))}
                    _ = fi78
                    _ = v89
                }
                _ = fi74
            }
        case 4294967295:
            v2 += v2
        case 9223372036854775807:
            var v90 uint32 = 215
            _ = v90
        default:
            v3["\n\t"] = ""
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
