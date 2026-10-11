package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 int16; F1 int16; F2 int64 }
type GenS5 struct { F0 uint; F1 int64; F2 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32; F1 uint16 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint; F1 int64; F2 int32 }
type GenS10 struct { F0 uint8; F1 int; F2 int }
type GenS11 struct { F0 uint16 }
func genFn0(p0 uint32, p1 uint32) uint8 {
    if false {
        var v4 bool = (!(!true))
        switch "" {
        case "\n\t":
            for fi5 := 0; fi5 < 3; fi5++ {
                v6 := GenS1{F0: (fi5 - fi5), F1: 559}
                _ = fi5
                _ = v6
            }
        case "hello":
            fmt.Println(p0, v4, v4)
        default:
            p1 -= (p0 % ((624) | 1))
        }
        v7 := GenS1{F0: 42, F1: 0}
        _ = v4
        _ = v7
    }
    switch 32768 {
    case 915:
        v8 := GenS1{F0: 128, F1: 0}
        _ = v8
    }
    _ = p0
    _ = p1
    return 141
}
func genFn1(p0 int16, p1 uint8) uint16 {
    v9 := GenS2{F0: (p0 >> ((4294967295) & 63)), F1: p0}
    v10 := []int{2147483647}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return uint16(p1)
}

func main() {
    var v1 int = -1
    v2 := make(map[uint32]uint32)
    var v3 int8 = 0
    var v11 bool = (1 < 735)
    if v11 {
        v12 := GenS3{F0: 42, F1: (("a" + ("\n\t" + "hello")) + "hello")}
        _ = v12
    } else {
        var v13 string = (("" + ("\n\t" + "")) + "foo bar")
        v14 := GenS3{F0: 42, F1: v13}
        if (genFn1(-32768, 237) != genFn1(969, 1)) {
            if v11 {
                v13 = (((v13 + (v13 + "")) + v13) + (v13 + ""))
                for fi15 := 0; fi15 < 5; fi15++ {
                    switch ((v13 + v13) + ("a" + v13)) {
                    case "":
                        v2[42] = 4294967295
                    case "a":
                        fmt.Println(fi15)
                    case "hello":
                        v16 := GenS4{F0: 1000, F1: 307, F2: 97}
                        _ = v16
                    }
                    if v11 {
                        v17 := GenS3{F0: -460, F1: "a"}
                        if (!true) {
                            for fi18 := 0; fi18 < 12; fi18++ {
                                v19 := GenS5{F0: 42, F1: -1, F2: (v3 * (v3 & v3))}
                                if v11 {
                                    if (!((3.14159 / (0.1 / 3.14159)) == float32(0.0))) {
                                        v3 -= (v3 + (v3 / (((v3 & (v3 * -40))) | 1)))
                                    } else {
                                        v20 := GenS6{F0: 4294967295}
                                        v11 = v11
                                        if (!(179 > 18446744073709551615)) {
                                            _ = len(v2)
                                        }
                                        _ = v20
                                    }
                                } else {
                                    if (!true) {
                                        var v21 uint32 = 2147483648
                                        v22 := GenS5{F0: 919, F1: 42, F2: v3}
                                        v23 := GenS6{F0: v21}
                                        switch (v13 + v13) {
                                        case "":
                                            v24 := []int{fi15, 4294967295, v1, 32768}
                                            _ = v24
                                        case "foo bar":
                                            v25 := GenS3{F0: -2147483648, F1: v13}
                                            _ = v25
                                        case "a":
                                            if (0 < (^(v3))) {
                                                v3 += (v3 >> ((1) & 63))
                                                switch v13 {
                                                case "foo bar":
                                                    var v26 int = -138
                                                    _ = v26
                                                case "\n\t":
                                                    var v27 uint = 98
                                                    _ = v27
                                                }
                                            }
                                        }
                                        _ = v21
                                        _ = v22
                                        _ = v23
                                    }
                                    var v28 uint64 = 647
                                    _ = v28
                                }
                                switch (fi15 / ((fi18) | 1)) {
                                case 9223372036854775807:
                                    var v29 int8 = (^(v3))
                                    _ = v29
                                default:
                                    switch v1 {
                                    case 340:
                                        var v30 float32 = 1.5
                                        _ = v30
                                    case 884:
                                        if false {
                                            var v31 int64 = 1
                                            switch (v13 + ("" + ((("a" + v13) + (v13 + "foo bar")) + (v13 + "hello")))) {
                                            case "hello":
                                                v11 = false
                                            case "":
                                                v13 = "a"
                                            case "a":
                                                v3 = v3
                                            }
                                            _ = v31
                                        }
                                        v32 := []int{(fi18 / ((fi18) | 1)), (fi18 & (fi18 << ((18446744073709551615) & 63))), fi15}
                                        v3 = v3
                                        _ = v32
                                    default:
                                        if (0 < 0) {
                                            v33 := []int{127, v1, (fi15 + (-(fi15))), (fi15 + 128)}
                                            fmt.Println(v11, v13, v11)
                                            _ = v33
                                        }
                                    }
                                }
                                _ = fi18
                                _ = v19
                            }
                        } else {
                            for fi34 := 0; fi34 < 13; fi34++ {
                                v35 := GenS6{F0: 42}
                                if v11 {
                                    delete(v2, 0)
                                    var v36 int32 = -36
                                    _ = v36
                                }
                                _ = fi34
                                _ = v35
                            }
                        }
                        _ = v17
                    }
                    _ = fi15
                }
                v11 = v11
            } else {
                var v37 int32 = 42
                v37 = v37
                if (-1 < 1) {
                    v37 += (v37 - v37)
                    switch (v1 * v1) {
                    case 4294967295:
                        fmt.Println(v1)
                    case 9223372036854775807:
                        v13 = v13
                    case 32768:
                        fmt.Println(v11, v13, v11)
                    default:
                        v13 = v13
                    }
                    v38 := []int{-2147483648, 42, v1}
                    _ = v38
                } else {
                    switch v13 {
                    case "\n\t":
                        v11 = (484 >= (v1 - v1))
                    }
                    var v39 bool = true
                    _ = v39
                }
                _ = v37
            }
            v40 := GenS7{F0: 585, F1: 65535}
            for fi41 := 0; fi41 < 2; fi41++ {
                switch 32768 {
                case -9223372036854775808:
                    fmt.Println(v13, v13)
                case -518:
                    if false {
                        fmt.Println(fi41)
                    } else {
                        v42 := GenS4{F0: 32767, F1: -729, F2: 1}
                        if ("foo bar" == (v13 + ("\n\t" + v13))) {
                            fmt.Println(v3)
                            for fi43 := 0; fi43 < 12; fi43++ {
                                v1 = 1
                                if (!(!(v3 <= (v3 ^ (^(v3)))))) {
                                    var v44 uint16 = 255
                                    switch 2147483647 {
                                    case -893:
                                        v11 = false
                                        fmt.Println(fi43, v1, v11)
                                        v44 = v44
                                        v45 := v2[42]
                                        v46 := GenS7{F0: 884, F1: v44}
                                        _ = v45
                                        _ = v46
                                    case 32767:
                                        if v11 {
                                            v3 += v3
                                            switch (^(fi43)) {
                                            case 40:
                                                switch -2147483648 {
                                                case -852:
                                                    var v47 int8 = v3
                                                    _ = v47
                                                case 2147483647:
                                                    var v48 int64 = 793
                                                    v3 = -128
                                                    _ = v48
                                                case 128:
                                                    fmt.Println(v3)
                                                }
                                            default:
                                                switch "\n\t" {
                                                case "\n\t":
                                                    if (1e10 <= 3.14159) {
                                                        var v49 int = 0
                                                        var v50 uint64 = 657
                                                        fmt.Println(v13, v44)
                                                        _ = v49
                                                        _ = v50
                                                    }
                                                default:
                                                    if v11 {
                                                        fmt.Println(fi41, v3, v3)
                                                        switch v13 {
                                                        case "hello":
                                                            v51 := v2[924]
                                                            _ = v51
                                                        case "a":
                                                            fmt.Println(fi41)
                                                        case "\n\t":
                                                            v13 = (v13 + "\n\t")
                                                        }
                                                    }
                                                }
                                            }
                                        } else {
                                            if (7 == (v3 << ((367) & 63))) {
                                                v52 := GenS8{F0: 127}
                                                _ = v52
                                            } else {
                                                v53 := []int{(v1 ^ fi41), (fi43 + (fi43 ^ v1))}
                                                v54 := GenS2{F0: 0, F1: -1}
                                                var v55 int64 = 0
                                                _ = v53
                                                _ = v54
                                                _ = v55
                                            }
                                            if v11 {
                                                _ = len(v2)
                                                if v11 {
                                                    v56 := GenS9{F0: 4294967295, F1: 1, F2: -1}
                                                    var v57 uint16 = v44
                                                    switch v1 {
                                                    case 128:
                                                        v57 += (v44 & (v44 % ((32767) | 1)))
                                                    case 42:
                                                        switch fi41 {
                                                        case 609:
                                                            v13 = ""
                                                            v11 = true
                                                            v44 = (v57 - v44)
                                                        case 32768:
                                                            v58 := GenS10{F0: 218, F1: (fi41 * -616), F2: 1}
                                                            _ = v58
                                                        case 128:
                                                            switch fi43 {
                                                            case 32768:
                                                                v59 := []int{(v1 | v1), (v1 ^ (fi41 ^ (v1 & v1)))}
                                                                _ = v59
                                                            case 9223372036854775807:
                                                                delete(v2, 65535)
                                                            case 42:
                                                                v60 := GenS6{F0: 1}
                                                                _ = v60
                                                            }
                                                            v1 = (^(fi43))
                                                            var v61 int = v1
                                                            _ = v61
                                                        }
                                                    case 2147483647:
                                                        var v62 uint = 255
                                                        _ = v62
                                                    }
                                                    _ = v56
                                                    _ = v57
                                                }
                                                var v63 int32 = 42
                                                _ = v63
                                            } else {
                                                fmt.Println(fi41, v11)
                                            }
                                            _ = len(v2)
                                        }
                                    default:
                                        var v64 int = (fi41 + (v1 | fi41))
                                        _ = v64
                                    }
                                    _ = v44
                                } else {
                                    var v65 int = (v1 / ((v1) | 1))
                                    switch v13 {
                                    case "a":
                                        var v66 uint16 = 0
                                        _ = v66
                                    case "hello":
                                        var v67 bool = (965 >= 42)
                                        _ = v67
                                    case "\n\t":
                                        v68 := GenS7{F0: 42, F1: 833}
                                        _ = v68
                                    }
                                    v11 = false
                                    _ = v65
                                }
                                fmt.Println(fi43, v13, v1)
                                v69 := GenS1{F0: fi43, F1: 629}
                                _ = fi43
                                _ = v69
                            }
                        }
                        v11 = ("" != "")
                        _ = v42
                    }
                case 1:
                    for fi70 := 0; fi70 < 14; fi70++ {
                        var v71 uint = 230
                        if v11 {
                            var v72 float64 = 1.5
                            var v73 int64 = -942
                            _ = v72
                            _ = v73
                        } else {
                            v11 = (-2.25 <= -2.25)
                        }
                        _ = fi70
                        _ = v71
                    }
                }
                _ = len(v2)
                v2[151] = 404
                _ = fi41
            }
            _ = v40
        } else {
            for fi74 := 0; fi74 < 19; fi74++ {
                var v75 uint16 = 2
                _ = fi74
                _ = v75
            }
        }
        _ = v13
        _ = v14
    }
    if ((!v11) && false) {
        v11 = false
        v1 += (v1 ^ 22)
    }
    v76 := GenS11{F0: 0}
    var v77 int16 = 876
    var v78 float32 = 0.0
    v1 -= -2147483648
    v79 := v2[65535]
    fmt.Println(v11)
    v11 = v11
    if v11 {
        var v80 int32 = -1
        v81 := GenS2{F0: (v77 ^ 42), F1: (v77 | (v77 ^ -1))}
        var v82 float32 = v78
        _ = v80
        _ = v81
        _ = v82
    } else {
        v83 := GenS8{F0: v3}
        _ = v83
    }
    if ((((-9223372036854775808 == -1) || (false && v11)) && true) || (v11 || (18446744073709551615 > 4294967295))) {
        var v84 uint8 = 42
        v85 := GenS7{F0: 489, F1: 427}
        _ = v84
        _ = v85
    } else {
        _ = len(v2)
        v2[609] = (v79 | (v79 * 2147483648))
        if (1 < 270) {
            v2[(v79 - 42)] = v79
            v1 = (v1 ^ (v1 + (v1 >> ((228) & 63))))
        }
    }
    var v86 int = (v1 ^ v1)
    v87 := []int{v86, v1, v86}
    var v88 uint = 558
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v86:", v86)
    fmt.Println("slicelen v87:", len(v87), v87)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v76
    _ = v77
    _ = v78
    _ = v79
    _ = v86
    _ = v87
    _ = v88
}
