package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 uint32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 uint16; F1 int16 }
type GenS9 struct { F0 bool; F1 int8; F2 uint32 }
type GenS10 struct { F0 uint; F1 int8 }
func genFn0(p0 uint8, p1 int8) uint8 {
    v7 := GenS1{F0: 18446744073709551615}
    p0 += (p0 % (((p0 / ((p0) | 1))) | 1))
    _ = p0
    _ = p1
    _ = v7
    return 1
}
func genFn1(p0 string, p1 int8) uint8 {
    for fi8 := 0; fi8 < 7; fi8++ {
        p1 = (p1 + 1)
        _ = fi8
    }
    v9 := GenS2{F0: 0, F1: true}
    _ = p0
    _ = p1
    _ = v9
    return 2
}

func main() {
    var v1 int32 = -714
    v2 := make(map[string]string)
    var v3 int16 = 42
    v4 := make(map[uint32]int)
    var v5 uint64 = 2
    var v6 string = "foo bar"
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := GenS3{F0: v5}
        _ = fi10
        _ = v11
    }
    var v12 int64 = 42
    v13 := GenS4{F0: false}
    v6 = v6
    v14 := GenS5{F0: 150, F1: 771}
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 float32 = ((float32(-2.25) + -2.25) * (0.1 / float32(1e10)))
        var v17 bool = true
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 int16 = (v3 << ((4294967295) & 63))
        if true {
            var v20 int8 = 100
            v6 = v6
            delete(v4, 4294967295)
            _ = v20
        }
        _ = fi18
        _ = v19
    }
    if false {
        var v21 int32 = -58
        _ = v21
    }
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v5)
    default:
        v22 := GenS3{F0: (v5 & (^(v5)))}
        _ = v22
    }
    v23 := []int{-2147483648, 4294967295, -2147483648}
    for fi24 := 0; fi24 < 2; fi24++ {
        var v25 uint16 = 571
        var v26 float64 = 1.5
        if true {
            var v27 uint16 = v25
            _ = v27
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    v12 += -9223372036854775808
    v28 := GenS6{F0: true}
    for fi29 := 0; fi29 < 19; fi29++ {
        var v30 float64 = 0.0
        switch "" {
        case "\n\t":
            switch v6 {
            case "foo bar":
                if true {
                    v1 = 1
                    var v31 float64 = float64(float32((v30 / 1e10)))
                    _ = v31
                } else {
                    v30 = v30
                }
            case "":
                if false {
                    v3 = v3
                } else {
                    for fi32 := 0; fi32 < 17; fi32++ {
                        fmt.Println(v6, fi29)
                        v33 := GenS2{F0: 255, F1: false}
                        if true {
                            v34 := GenS4{F0: true}
                            _ = v34
                        }
                        _ = fi32
                        _ = v33
                    }
                }
            }
        case "a":
            var v35 int = 42
            _ = v35
        }
        var v36 uint8 = 42
        _ = fi29
        _ = v30
        _ = v36
    }
    v1 = 1
    for fi37 := 0; fi37 < 9; fi37++ {
        if false {
            var v38 bool = (729 < 667)
            for fi39 := 0; fi39 < 19; fi39++ {
                switch "foo bar" {
                case "":
                    switch v6 {
                    case "hello":
                        v5 = 155
                    case "a":
                        switch (v6 + v6) {
                        case "foo bar":
                            v40 := GenS7{F0: v38, F1: v1}
                            _ = v40
                        case "a":
                            delete(v4, 4294967295)
                        case "\n\t":
                            switch "hello" {
                            case "":
                                v3 -= 500
                            case "a":
                                if (v5 == v5) {
                                    v38 = (((1.5 == -2.25) && ((v38 && (100 > int8(v5))) && v38)) || ((!v38) && true))
                                    v5 += (v5 ^ (v5 ^ (v5 - v5)))
                                    var v41 int64 = v12
                                    _ = v41
                                } else {
                                    var v42 float32 = 3.14159
                                    if (fi37 <= (fi39 + fi39)) {
                                        var v43 int8 = 100
                                        var v44 uint32 = 368
                                        delete(v2, "")
                                        _ = v43
                                        _ = v44
                                    } else {
                                        if (v38 || v38) {
                                            v5 += (v5 % ((v5) | 1))
                                            v6 = v6
                                        }
                                        var v45 int8 = 127
                                        switch (fi39 ^ (fi39 ^ 127)) {
                                        case 4294967295:
                                            if (true && false) {
                                                var v46 uint64 = (v5 / (((v5 | v5)) | 1))
                                                var v47 int16 = v3
                                                var v48 uint32 = 2147483647
                                                _ = v46
                                                _ = v47
                                                _ = v48
                                            }
                                        }
                                        var v49 uint64 = 184
                                        if (v38 || v38) {
                                            v38 = (true && (((v5 % ((v5) | 1)) == (v5 % ((v5) | 1))) || (((v6 + ("a" + "foo bar")) + (v6 + v6)) != "a")))
                                            if (42 > 0) {
                                                var v50 int64 = (v12 * 42)
                                                _ = v50
                                            } else {
                                                if v38 {
                                                    v42 -= (((((3.14159 / (0.0 - (v42 / v42))) * v42) + 0.1) * float32(float64(float32(3.14159)))) - float32(-2.25))
                                                    delete(v2, v6)
                                                    var v51 int = fi39
                                                    _ = v51
                                                }
                                                if (!(v1 <= 342)) {
                                                    v23 = append(v23, fi39)
                                                    fmt.Println(v42, fi39, fi39)
                                                    if (false || false) {
                                                        v52 := GenS1{F0: (v49 ^ (v5 / (((v49 & 55)) | 1)))}
                                                        v53 := v4[449]
                                                        if (0 > v12) {
                                                            v38 = ((v5 << ((1) & 63)) == 4294967295)
                                                            var v54 int = 32767
                                                            v38 = (v38 || v38)
                                                            _ = v54
                                                        } else {
                                                            fmt.Println(fi37, v12)
                                                            if (v49 <= v5) {
                                                                fmt.Println(v53)
                                                                if v38 {
                                                                    v5 -= 313
                                                                } else {
                                                                    var v55 uint = 458
                                                                    _ = len(v23)
                                                                    if false {
                                                                        v1 += int32(v55)
                                                                        v56 := GenS2{F0: 2, F1: (v38 && (v12 >= int64(v55)))}
                                                                        _ = v56
                                                                    } else {
                                                                        v23 = append(v23, v53)
                                                                        v55 = v55
                                                                    }
                                                                    _ = v55
                                                                }
                                                            } else {
                                                                var v57 int = -375
                                                                v6 = "\n\t"
                                                                _ = v57
                                                            }
                                                        }
                                                        _ = v52
                                                        _ = v53
                                                    } else {
                                                        v42 -= ((v42 + 1e10) - v42)
                                                    }
                                                } else {
                                                    v12 -= 0
                                                    _ = len(v23)
                                                    var v58 uint16 = 32767
                                                    _ = v58
                                                }
                                            }
                                            if v38 {
                                                var v59 uint16 = 65535
                                                if v38 {
                                                    v45 = (v45 >> ((65535) & 63))
                                                    v60 := GenS8{F0: (^(v59)), F1: v3}
                                                    var v61 uint32 = 185
                                                    _ = v60
                                                    _ = v61
                                                }
                                                _ = v59
                                            }
                                        }
                                        if v38 {
                                            if (v1 == (v1 - v1)) {
                                                var v62 uint16 = 1
                                                v12 += (v12 | v12)
                                                fmt.Println(v5, v3)
                                                _ = v62
                                            } else {
                                                var v63 int64 = int64(v3)
                                                v12 -= (v12 + v63)
                                                var v64 int16 = v3
                                                _ = v63
                                                _ = v64
                                            }
                                            if false {
                                                fmt.Println(v49, v5, v1)
                                            } else {
                                                v65 := GenS9{F0: v38, F1: 1, F2: 2147483647}
                                                if false {
                                                    fmt.Println(v3)
                                                    var v66 int8 = (^(v45))
                                                    _ = v66
                                                }
                                                _ = v65
                                            }
                                            var v67 uint64 = 0
                                            _ = v67
                                        }
                                        _ = v45
                                        _ = v49
                                    }
                                    _ = v42
                                }
                            case "\n\t":
                                delete(v2, v6)
                            default:
                                switch v6 {
                                case "hello":
                                    if (v38 && (float64(0.1) != (((-2.25 / 1.5) * float64(float32(-2.25))) + 1.5))) {
                                        if true {
                                            var v68 int16 = v3
                                            switch fi37 {
                                            case -9223372036854775808:
                                                v6 = v6
                                            case -2147483648:
                                                var v69 float64 = -2.25
                                                _ = v69
                                            case 124:
                                                _ = len(v2)
                                            }
                                            v68 = v68
                                            v68 = 876
                                            _ = v68
                                        } else {
                                            switch fi37 {
                                            case 2147483647:
                                                v6 = v6
                                            case -572:
                                                var v70 int64 = v12
                                                _ = v70
                                            case 4294967295:
                                                if true {
                                                    fmt.Println(fi39, v38)
                                                    fmt.Println(v3, v12)
                                                    if false {
                                                        v5 += v5
                                                        switch fi37 {
                                                        case 315:
                                                            if true {
                                                                var v71 float64 = -2.25
                                                                _ = v71
                                                            } else {
                                                                v3 = v3
                                                                v72 := GenS5{F0: 0, F1: 65535}
                                                                _ = v72
                                                            }
                                                        case 32767:
                                                            v73 := GenS4{F0: v38}
                                                            _ = v73
                                                        case 971:
                                                            if (false && (v6 < "hello")) {
                                                                v74 := GenS10{F0: 18446744073709551615, F1: 1}
                                                                v75 := v2[v6]
                                                                _ = v74
                                                                _ = v75
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    fmt.Println(v3, fi39, v6)
                                                    var v76 float64 = -2.25
                                                    _ = v76
                                                }
                                            default:
                                                var v77 bool = (4294967295 != 255)
                                                _ = v77
                                            }
                                            if false {
                                                var v78 uint8 = 127
                                                _ = v78
                                            }
                                        }
                                    }
                                    var v79 float32 = (1.5 * 0.1)
                                    _ = len(v23)
                                    _ = v79
                                case "\n\t":
                                    if len(v23) >= 2 {
                                        v23 = v23[1:len(v23)]
                                    }
                                    fmt.Println(v12, v12, v12)
                                    switch fi37 {
                                    case 32767:
                                        v80 := GenS10{F0: 18446744073709551615, F1: -107}
                                        _ = v80
                                    case 32768:
                                        var v81 int16 = v3
                                        _ = v81
                                    case 128:
                                        delete(v2, v6)
                                    }
                                case "foo bar":
                                    v12 = (v12 * (v12 | v12))
                                }
                            }
                        default:
                            if (!((fi37 + fi39) >= (fi39 ^ 9223372036854775807))) {
                                switch v6 {
                                case "a":
                                    fmt.Println(fi39, v6)
                                }
                                var v82 float32 = ((((0.1 - 0.1) * 0.1) + ((-2.25 / (float32((0.1 / 1.0)) * 3.14159)) / 1.0)) / (1e10 + (1.5 / 1e10)))
                                var v83 int8 = -128
                                if len(v23) >= 2 {
                                    v23 = v23[1:len(v23)]
                                }
                                _ = v82
                                _ = v83
                            }
                        }
                    case "":
                        switch (fi37 >> ((0) & 63)) {
                        case 0:
                            v84 := GenS8{F0: 65535, F1: 0}
                            _ = v84
                        case 972:
                            fmt.Println(v1)
                        case 2147483647:
                            v5 = 961
                        default:
                            v1 = 1
                        }
                        if true {
                            if (-2147483648 > (v1 * v1)) {
                                v6 = ("" + v6)
                                var v85 float64 = 3.14159
                                var v86 uint8 = 42
                                var v87 bool = (v38 || (!(v86 >= genFn0((v86 / ((0) | 1)), 1))))
                                var v88 uint64 = v5
                                _ = v85
                                _ = v86
                                _ = v87
                                _ = v88
                            } else {
                                v38 = (!(-85 <= (-(v12))))
                            }
                        }
                    }
                default:
                    _ = len(v23)
                }
                v5 -= 1
                _ = fi39
            }
            var v89 int32 = (v1 | v1)
            _ = v38
            _ = v89
        }
        v4[151] = 4
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
