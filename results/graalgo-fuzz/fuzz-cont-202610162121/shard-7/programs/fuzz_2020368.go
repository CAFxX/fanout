package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 int; F1 int64; F2 int16 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 int8; F1 int16; F2 int8 }
type GenS6 struct { F0 int8; F1 int32; F2 int16 }
type GenS7 struct { F0 int; F1 uint32; F2 string }
type GenS8 struct { F0 int; F1 int }
type GenS9 struct { F0 int8; F1 uint; F2 int }
func genFn0(p0 int8, p1 string, p2 int16) int8 {
    for fi5 := 0; fi5 < 18; fi5++ {
        v6 := []int{fi5, (fi5 << ((4294967295) & 63))}
        var v7 float64 = -2.25
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 int = -185
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 string = (p1 + p1)
        _ = fi9
        _ = v10
    }
    v11 := []int{v8, v8, -2147483648}
    for fi12 := 0; fi12 < 13; fi12++ {
        _ = len(v11)
        p2 = p2
        v11 = append(v11, (v8 >> ((4294967295) & 63)))
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v8
    return 127
}
func genFn1(p0 uint16, p1 uint64, p2 bool) uint32 {
    p0 += (p0 % (((p0 | 32767)) | 1))
    if true {
        p1 += p1
    } else {
        p0 -= p0
        var v13 float32 = 0.0
        if (9223372036854775807 < 824) {
            if p2 {
                switch -9223372036854775808 {
                case 32767:
                    switch 1 {
                    case -436:
                        v13 = (0.0 - 0.1)
                    case 9223372036854775807:
                        for fi14 := 0; fi14 < 3; fi14++ {
                            v15 := []int{fi14, fi14}
                            _ = fi14
                            _ = v15
                        }
                    }
                case 113:
                    for fi16 := 0; fi16 < 3; fi16++ {
                        var v17 uint32 = 1
                        v18 := []int{}
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                }
            } else {
                var v19 uint64 = (p1 * 0)
                _ = v19
            }
            p1 += p1
            for fi20 := 0; fi20 < 1; fi20++ {
                v21 := GenS1{F0: -9223372036854775808, F1: int8(p0)}
                v22 := []int{fi20}
                if len(v22) > 0 {
                    v22[0] = (fi20 - fi20)
                }
                _ = fi20
                _ = v21
                _ = v22
            }
        } else {
            if (-263 < -1) {
                for fi23 := 0; fi23 < 18; fi23++ {
                    for fi24 := 0; fi24 < 11; fi24++ {
                        var v25 uint64 = (p1 << ((42) & 63))
                        _ = fi24
                        _ = v25
                    }
                    for fi26 := 0; fi26 < 19; fi26++ {
                        v27 := []int{(fi26 | (fi23 + fi23)), (fi26 - (fi26 % ((fi23) | 1))), (fi23 + 1), (fi23 | (-(fi26)))}
                        p0 -= (p0 % ((1) | 1))
                        var v28 uint64 = p1
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                    v29 := GenS1{F0: 42, F1: -52}
                    _ = fi23
                    _ = v29
                }
            }
            if (33 <= 2) {
                v30 := GenS1{F0: 0, F1: 42}
                for fi31 := 0; fi31 < 6; fi31++ {
                    v32 := []int{fi31, fi31}
                    _ = fi31
                    _ = v32
                }
                fmt.Println(v13, p1, p2)
                _ = v30
            } else {
                for fi33 := 0; fi33 < 7; fi33++ {
                    for fi34 := 0; fi34 < 2; fi34++ {
                        if false {
                            v13 = ((-2.25 + v13) + ((v13 / (((-2.25 - v13) + (v13 - v13)) + float32((-2.25 + 0.1)))) * v13))
                        } else {
                            var v35 uint16 = p0
                            fmt.Println(fi34)
                            _ = v35
                        }
                        var v36 bool = (v13 >= 3.14159)
                        v37 := GenS2{F0: 2147483647, F1: 2}
                        v38 := []int{(fi34 % ((fi34) | 1)), fi33, (fi33 << ((65535) & 63)), -1}
                        _ = fi34
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                    var v39 float32 = v13
                    _ = fi33
                    _ = v39
                }
                if ((-2.25 / 3.14159) <= 1.5) {
                    p2 = false
                    var v40 int8 = genFn0(114, "hello", 1)
                    v13 -= -2.25
                    _ = v40
                }
                for fi41 := 0; fi41 < 5; fi41++ {
                    var v42 int64 = -1
                    var v43 int16 = 42
                    _ = fi41
                    _ = v42
                    _ = v43
                }
            }
        }
        _ = v13
    }
    fmt.Println(p2)
    var v44 int64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v44
    return 65535
}

func main() {
    v1 := []string{("" + ("" + "foo bar")), "\n\t", (("hello" + ("hello" + "a")) + "foo bar")}
    var v2 bool = false
    var v3 uint32 = 2147483647
    var v4 bool = v2
    fmt.Println(v3, v3, v2)
    for fi45 := 0; fi45 < 13; fi45++ {
        switch (fi45 * (-(fi45))) {
        case 4294967295:
            v3 += 790
        case -754:
            var v46 float32 = 3.14159
            _ = v46
        case 127:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            for fi47 := 0; fi47 < 5; fi47++ {
                var v48 uint16 = 42
                var v49 float32 = (1.5 + 0.1)
                fmt.Println(v2)
                _ = fi47
                _ = v48
                _ = v49
            }
        }
        _ = fi45
    }
    if (65535 < 255) {
        for fi50 := 0; fi50 < 7; fi50++ {
            if (!(v2 && (92 <= 127))) {
                fmt.Println(v4)
                for fi51 := 0; fi51 < 16; fi51++ {
                    var v52 string = ("\n\t" + ("a" + ("\n\t" + "")))
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if v4 {
                        v53 := GenS3{F0: fi50, F1: -9223372036854775808, F2: 71}
                        _ = v53
                    } else {
                        if (!v2) {
                            var v54 uint = 4294967295
                            var v55 int = (fi50 >> ((v54) & 63))
                            _ = len(v1)
                            fmt.Println(v3)
                            var v56 uint = v54
                            v3 = v3
                            v55 -= -9223372036854775808
                            v4 = false
                            _ = v54
                            _ = v55
                            _ = v56
                        }
                        switch ("foo bar" + ("hello" + v52)) {
                        case "\n\t":
                            var v57 uint16 = 42
                            _ = v57
                        case "":
                            v58 := GenS4{F0: -801, F1: 109}
                            _ = v58
                        }
                        v2 = true
                    }
                    _ = fi51
                    _ = v52
                }
            }
            fmt.Println(fi50)
            _ = fi50
        }
    }
    var v59 float64 = 0.1
    fmt.Println(v3, v4, v3)
    if false {
        v60 := GenS5{F0: 42, F1: 0, F2: -128}
        if false {
            v2 = (32767 != 438)
            v61 := GenS6{F0: genFn0(100, ("a" + "\n\t"), 32767), F1: -1, F2: 474}
            v62 := GenS7{F0: -661, F1: genFn1(42, 627, (1000 == 32767)), F2: "foo bar"}
            _ = v61
            _ = v62
        } else {
            for fi63 := 0; fi63 < 2; fi63++ {
                _ = len(v1)
                _ = fi63
            }
        }
        _ = v60
    }
    for fi64 := 0; fi64 < 11; fi64++ {
        var v65 float64 = (float64(((float32((0.0 - v59)) / ((3.14159 / ((-2.25 + 3.14159) - 1.5)) + 1e10)) / 1.0)) + v59)
        _ = fi64
        _ = v65
    }
    if (!(v4 || (!(-128 != genFn0(genFn0(1, "", -1), "\n\t", 1))))) {
        v66 := GenS6{F0: 127, F1: -430, F2: 42}
        v3 -= 437
        if false {
            if (167 != 9) {
                var v67 float64 = 0.1
                _ = v67
            }
        } else {
            v4 = (false && (255 < 2))
        }
        _ = v66
    }
    if true {
        var v68 float32 = (0.0 + ((1e10 * (1e10 * 0.1)) / float32((v59 / ((float64(-2.25) - v59) / ((v59 * v59) * 0.1))))))
        for fi69 := 0; fi69 < 8; fi69++ {
            var v70 uint8 = 2
            if v2 {
                v4 = (2 < genFn1(42, 279, v4))
                if (!true) {
                    for fi71 := 0; fi71 < 11; fi71++ {
                        var v72 int64 = int64(v3)
                        fmt.Println(v68, fi69, v70)
                        _ = fi71
                        _ = v72
                    }
                } else {
                    for fi73 := 0; fi73 < 8; fi73++ {
                        var v74 int32 = 42
                        _ = fi73
                        _ = v74
                    }
                }
                v75 := GenS8{F0: fi69, F1: (fi69 * 0)}
                _ = v75
            }
            var v76 int8 = genFn0(0, "\n\t", 740)
            _ = fi69
            _ = v70
            _ = v76
        }
        _ = v68
    } else {
        fmt.Println(v3, v4)
        var v77 float64 = float64(3.14159)
        fmt.Println(v3, v4)
        _ = v77
    }
    var v78 uint = 255
    if (false || (42 <= 208)) {
        if v2 {
            for fi79 := 0; fi79 < 19; fi79++ {
                for fi80 := 0; fi80 < 19; fi80++ {
                    v2 = (-931 <= 543)
                    fmt.Println(v78, v78, v59)
                    _ = fi80
                }
                var v81 uint = v78
                v4 = (false && false)
                _ = fi79
                _ = v81
            }
            var v82 uint = v78
            _ = v82
        } else {
            if len(v1) > 0 {
                v1[0] = "hello"
            }
        }
        fmt.Println(v3, v59)
        if (uint(v3) <= (v78 & v78)) {
            fmt.Println(v59)
        } else {
            var v83 uint32 = 345
            _ = v83
        }
    }
    for fi84 := 0; fi84 < 15; fi84++ {
        var v85 int = (fi84 & (fi84 * (fi84 - fi84)))
        var v86 uint8 = 172
        _ = fi84
        _ = v85
        _ = v86
    }
    v1 = append(v1, "")
    var v87 float32 = float32((float64((1e10 + 1.5)) - 3.14159))
    if v2 {
        if (false || ((v3 - 4294967295) < (v3 >> (((v78 | v78)) & 63)))) {
            fmt.Println(v78)
            _ = len(v1)
            v88 := GenS9{F0: -1, F1: 255, F2: -2147483648}
            _ = v88
        } else {
            switch 127 {
            case 128:
                v1 = append(v1, ("\n\t" + "foo bar"))
            }
            switch 1 {
            case -802:
                for fi89 := 0; fi89 < 10; fi89++ {
                    var v90 int = fi89
                    _ = fi89
                    _ = v90
                }
            default:
                var v91 int16 = 42
                _ = v91
            }
        }
        v78 = (^(v78))
        var v92 uint8 = 127
        _ = v92
    }
    v4 = true
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v93 uint64 = 18446744073709551615
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v59:", v59)
    fmt.Println("v78:", v78)
    fmt.Println("v87:", v87)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v59
    _ = v78
    _ = v87
    _ = v93
}
