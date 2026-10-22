package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int8; F1 int32; F2 int16 }
type GenS8 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint32; F1 int16 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]bool)
    var v3 uint = 42
    v4 := make(map[string]uint32)
    v5 := make(map[string]string)
    var v6 uint64 = 2
    var v7 uint16 = 637
    v8 := GenS1{F0: 42, F1: -2, F2: 1000}
    v9 := GenS1{F0: 42, F1: -79, F2: 0}
    if (127 == -128) {
        if (2 >= 545) {
            switch 348 {
            case 717:
                var v10 bool = (3.14159 == float32(-2.25))
                _ = v10
            case 4294967295:
                v11 := GenS1{F0: 1, F1: 127, F2: 1000}
                _ = v11
            }
            var v12 uint = 0
            _ = v12
        } else {
            v13 := GenS2{F0: 2147483648}
            v14 := []int{-9223372036854775808, 4294967295, 9223372036854775807, 127}
            switch "\n\t" {
            case "hello":
                for fi15 := 0; fi15 < 11; fi15++ {
                    fmt.Println(v6, v3, v3)
                    v6 = v6
                    _ = fi15
                }
            default:
                var v16 bool = (-149 < 127)
                _ = v16
            }
            _ = v13
            _ = v14
        }
    } else {
        for fi17 := 0; fi17 < 6; fi17++ {
            switch "a" {
            case "hello":
                if (v3 <= (v3 & v3)) {
                    var v18 int16 = 32767
                    var v19 float64 = 3.14159
                    _ = v18
                    _ = v19
                }
            }
            _ = fi17
        }
    }
    v20 := v2[-2147483648]
    for fi21 := 0; fi21 < 15; fi21++ {
        switch 4294967295 {
        case 280:
            v3 = v3
        }
        var v22 int16 = -1
        if (true || v20) {
            for fi23 := 0; fi23 < 6; fi23++ {
                _ = len(v4)
                var v24 int32 = int32(v7)
                switch "foo bar" {
                case "":
                    if (-1 >= 810) {
                        var v25 int32 = v24
                        _ = v25
                    } else {
                        v26 := []int{int(v6)}
                        if len(v26) > 0 {
                            v26[0] = int(v22)
                        }
                        if v20 {
                            fmt.Println(v20, v22)
                            v27 := GenS2{F0: 539}
                            _ = v27
                        } else {
                            if (v20 || (v7 < 2)) {
                                if len(v26) > 0 {
                                    v26[0] = (fi23 / (((^(fi21))) | 1))
                                }
                                v26 = append(v26, fi23)
                                v3 = (v3 << (((v3 + 255)) & 63))
                            }
                        }
                        _ = v26
                    }
                    v28 := GenS3{F0: 533}
                    if v20 {
                        switch (("" + (("a" + ("\n\t" + "foo bar")) + "hello")) + "a") {
                        case "":
                            v3 = v3
                        case "\n\t":
                            v7 = v7
                        default:
                            _ = len(v4)
                        }
                        var v29 uint8 = 1
                        v29 -= 75
                        if (100 <= 127) {
                            var v30 int64 = 375
                            switch ("\n\t" + "hello") {
                            case "a":
                                v22 = 1
                                switch ("" + "hello") {
                                case "\n\t":
                                    if v20 {
                                        v31 := []int{(fi23 | (fi23 & fi21)), (fi21 - (fi23 * (fi23 % (((fi23 | fi23)) | 1)))), (fi23 * fi23), 962}
                                        switch fi23 {
                                        case -442:
                                            if ((fi23 + (fi21 >> ((2) & 63))) > (fi21 - (-(fi21)))) {
                                                fmt.Println(v3)
                                                delete(v2, 141)
                                            } else {
                                                v32 := GenS4{F0: 128, F1: 830, F2: v29}
                                                v31 = append(v31, (fi23 * fi23))
                                                _ = v32
                                            }
                                        case -2147483648:
                                            if v20 {
                                                var v33 uint16 = v7
                                                var v34 string = (("foo bar" + ("hello" + ((("foo bar" + "foo bar") + "") + "\n\t"))) + "hello")
                                                if (!v20) {
                                                    var v35 int8 = 127
                                                    var v36 int16 = (v22 * v22)
                                                    v37 := GenS2{F0: 693}
                                                    _ = v35
                                                    _ = v36
                                                    _ = v37
                                                } else {
                                                    var v38 int = 128
                                                    var v39 int32 = (v24 * v24)
                                                    switch fi21 {
                                                    case -327:
                                                        if ((v39 - (v39 << (((v3 % ((v3) | 1))) & 63))) == int32(v38)) {
                                                            v22 -= v22
                                                            var v40 int64 = v30
                                                            if ((1e10 - 3.14159) < 1e10) {
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    if (0 == 1) {
                                                                        if false {
                                                                            if len(v31) >= 2 {
                                                                                v31 = v31[1:len(v31)]
                                                                            }
                                                                        }
                                                                        if v20 {
                                                                            switch v38 {
                                                                            case -918:
                                                                                if (true || true) {
                                                                                    if v20 {
                                                                                        v30 = (v40 << (((v3 % (((v3 ^ uint(v29))) | 1))) & 63))
                                                                                        v3 = (v3 * 65535)
                                                                                    }
                                                                                    var v41 int8 = 42
                                                                                    _ = v41
                                                                                }
                                                                            case 2147483647:
                                                                                if v20 {
                                                                                    switch -132 {
                                                                                    case -133:
                                                                                        v42 := GenS5{F0: (fi21 | fi21)}
                                                                                        _ = v42
                                                                                    }
                                                                                    if v20 {
                                                                                        v30 -= int64(v39)
                                                                                        switch (^(v38)) {
                                                                                        case 569:
                                                                                            fmt.Println(v34)
                                                                                        case 0:
                                                                                            _ = len(v4)
                                                                                        case 1:
                                                                                            v31 = append(v31, (fi21 % ((fi23) | 1)))
                                                                                        default:
                                                                                            v29 += v29
                                                                                        }
                                                                                        v43 := GenS1{F0: v40, F1: -128, F2: (v22 * (^(v22)))}
                                                                                        _ = v43
                                                                                    } else {
                                                                                        if false {
                                                                                            var v44 int32 = (^(v24))
                                                                                            var v45 bool = true
                                                                                            if len(v31) > 0 {
                                                                                                v31[0] = fi23
                                                                                            }
                                                                                            v3 = (v3 | (v3 * (^(v3))))
                                                                                            v46 := GenS3{F0: (v33 << (((v3 * (v3 << ((337) & 63)))) & 63))}
                                                                                            _ = v44
                                                                                            _ = v45
                                                                                            _ = v46
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    v47 := GenS6{F0: (fi21 ^ (fi23 % (((v38 << ((v3) & 63))) | 1)))}
                                                                                    fmt.Println(v33, v7)
                                                                                    if ((v33 >= v33) || v20) {
                                                                                        fmt.Println(v22, v33, v40)
                                                                                    }
                                                                                    _ = v47
                                                                                }
                                                                            case 32768:
                                                                                v34 = "foo bar"
                                                                            default:
                                                                                v48 := GenS1{F0: v40, F1: 100, F2: v22}
                                                                                _ = v48
                                                                            }
                                                                            _ = len(v5)
                                                                            var v49 int32 = v24
                                                                            _ = v49
                                                                        }
                                                                    } else {
                                                                        switch "" {
                                                                        case "":
                                                                            var v50 uint64 = (v6 - 42)
                                                                            _ = v50
                                                                        case "foo bar":
                                                                            if (65535 <= (v33 << (((v3 / ((v3) | 1))) & 63))) {
                                                                                v29 = 1
                                                                            }
                                                                        }
                                                                    }
                                                                case "hello":
                                                                    if (((v3 | 0) <= v3) || (((!v20) && ((v33 * v33) <= uint16(v3))) && v20)) {
                                                                        v51 := GenS1{F0: v40, F1: 0, F2: (v22 & v22)}
                                                                        switch "a" {
                                                                        case "hello":
                                                                            v52 := v1[1]
                                                                            _ = v52
                                                                        case "\n\t":
                                                                            var v53 int64 = int64(v33)
                                                                            _ = v53
                                                                        default:
                                                                            var v54 float64 = 3.14159
                                                                            _ = v54
                                                                        }
                                                                        v55 := GenS3{F0: v7}
                                                                        _ = v51
                                                                        _ = v55
                                                                    } else {
                                                                        switch fi23 {
                                                                        case 9223372036854775807:
                                                                            v20 = v20
                                                                        case 42:
                                                                            var v56 uint16 = v33
                                                                            _ = v56
                                                                        case -1:
                                                                            delete(v2, (fi23 ^ v38))
                                                                        }
                                                                    }
                                                                case "a":
                                                                    var v57 uint64 = (v6 + (^(v6)))
                                                                    _ = v57
                                                                }
                                                                delete(v2, (-(fi23)))
                                                                delete(v2, (fi23 << ((255) & 63)))
                                                            }
                                                            switch 2147483647 {
                                                            case -1:
                                                                var v58 bool = true
                                                                _ = v58
                                                            case 385:
                                                                v7 += (v33 / (((v33 * (v33 | v33))) | 1))
                                                                var v59 uint = v3
                                                                _ = v59
                                                            }
                                                            _ = v40
                                                        } else {
                                                            var v60 string = v34
                                                            fmt.Println(v6, v24)
                                                            _ = v60
                                                        }
                                                    case 0:
                                                        if len(v31) > 0 {
                                                            v31[0] = fi23
                                                        }
                                                    case -980:
                                                        v38 -= 2147483647
                                                    }
                                                    _ = v38
                                                    _ = v39
                                                }
                                                _ = v33
                                                _ = v34
                                            } else {
                                                var v61 int16 = v22
                                                v24 = (v24 >> ((142) & 63))
                                                if len(v31) > 0 {
                                                    v31[0] = (fi23 & (^(fi23)))
                                                }
                                                _ = v61
                                            }
                                        case 2147483647:
                                            v29 += 127
                                        default:
                                            v6 += (v6 - uint64(v29))
                                        }
                                        _ = v31
                                    }
                                default:
                                    var v62 uint16 = 121
                                    _ = v62
                                }
                                var v63 int = (fi23 * fi21)
                                _ = v63
                            case "hello":
                                v64 := GenS7{F0: 127, F1: v24, F2: (v22 << (((^(v3))) & 63))}
                                _ = v64
                            }
                            v65 := []int{fi23, (^(fi23))}
                            switch ("" + "foo bar") {
                            case "\n\t":
                                switch (fi21 & fi21) {
                                case -510:
                                    if (v20 && v20) {
                                        v66 := GenS7{F0: -1, F1: v24, F2: (-(v22))}
                                        v3 += v3
                                        _ = v66
                                    }
                                case 1:
                                    fmt.Println(fi21, fi23, v29)
                                case 0:
                                    if (0 == v7) {
                                        fmt.Println(v7)
                                        fmt.Println(v22, fi21)
                                        var v67 string = "\n\t"
                                        _ = v67
                                    }
                                default:
                                    v68 := GenS8{F0: 4294967295, F1: v30, F2: v6}
                                    _ = v68
                                }
                            }
                            var v69 float64 = 1.5
                            switch "foo bar" {
                            case "foo bar":
                                var v70 bool = v20
                                _ = v70
                            }
                            _ = v30
                            _ = v65
                            _ = v69
                        } else {
                            if v20 {
                                v71 := GenS2{F0: 2}
                                if (!(v24 > v24)) {
                                    fmt.Println(v29, v20, v29)
                                    var v72 float32 = (1.5 + (float32((3.14159 * 0.0)) - 1.5))
                                    _ = v72
                                } else {
                                    var v73 bool = v20
                                    v74 := GenS9{F0: v20}
                                    _ = v73
                                    _ = v74
                                }
                                if v20 {
                                    var v75 int32 = (v24 & v24)
                                    _ = v75
                                } else {
                                    var v76 uint16 = (v7 & v7)
                                    _ = v76
                                }
                                _ = v71
                            }
                            v77 := GenS10{F0: (v24 + -718)}
                            _ = v77
                        }
                        _ = v29
                    } else {
                        var v78 uint = 485
                        switch (fi23 % ((fi21) | 1)) {
                        case 4294967295:
                            v79 := []int{(fi23 << (((v3 ^ v3)) & 63)), 42}
                            _ = v79
                        }
                        delete(v2, (fi21 & (fi23 & (fi21 | (fi21 % ((fi21) | 1))))))
                        v80 := GenS5{F0: -1}
                        v81 := v2[(fi23 * (fi23 / ((fi21) | 1)))]
                        v22 = v22
                        var v82 uint = v3
                        v83 := GenS10{F0: v24}
                        _ = v78
                        _ = v80
                        _ = v81
                        _ = v82
                        _ = v83
                    }
                    _ = v28
                }
                _ = fi23
                _ = v24
            }
            var v84 float32 = float32(-2.25)
            _ = v84
        }
        _ = fi21
        _ = v22
    }
    var v85 int8 = 42
    v86 := GenS11{F0: 594, F1: 452}
    var v87 float64 = ((float64((1.5 - 1e10)) - 0.1) - -2.25)
    var v88 string = "foo bar"
    var v89 float64 = v87
    v6 -= 34
    v90 := GenS2{F0: 2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v85:", v85)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v85
    _ = v86
    _ = v87
    _ = v88
    _ = v89
    _ = v9
    _ = v90
}
