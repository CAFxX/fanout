package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int; F2 uint32 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS6 struct { F0 bool; F1 uint32; F2 int }
type GenS7 struct { F0 int32; F1 bool }
type GenS8 struct { F0 uint32; F1 int64 }
type GenS9 struct { F0 uint16; F1 string }
func genFn0() uint32 {
    var v7 bool = (-873 == 1)
    if (v7 && v7) {
        v7 = false
        var v8 bool = v7
        _ = v8
    } else {
        fmt.Println(v7)
        var v9 uint8 = 0
        var v10 int64 = 0
        _ = v10
        _ = v9
    }
    _ = v7
    return 65535
}
func genFn1(p0 bool, p1 int32, p2 bool) int8 {
    var v11 uint32 = 2147483648
    var v12 int64 = 1
    var v13 uint64 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    return 42
}

func main() {
    v1 := []string{("a" + ("" + "hello")), ((("hello" + ("" + ("" + (((("\n\t" + "a") + "\n\t") + "a") + "")))) + "a") + "\n\t"), ("\n\t" + (("hello" + "") + "hello")), "foo bar"}
    v2 := make(map[uint32]string)
    v3 := []uint32{944, 1, 0, 104}
    v4 := make(map[uint32]string)
    var v5 uint16 = 65535
    var v6 uint32 = 776
    if false {
        _ = len(v4)
    } else {
        fmt.Println(v5)
        fmt.Println(v5, v6)
        var v14 int32 = int32(v6)
        _ = v14
    }
    fmt.Println(v6, v6)
    switch 32767 {
    case 1:
        v15 := GenS1{F0: 1, F1: 42}
        _ = v15
    }
    v5 -= v5
    var v16 int = 1
    if (v16 > -563) {
        v17 := v4[v6]
        var v18 uint = 2
        for fi19 := 0; fi19 < 16; fi19++ {
            for fi20 := 0; fi20 < 11; fi20++ {
                v17 = v17
                if (1 != 42) {
                    if true {
                        v21 := GenS1{F0: 11, F1: 18}
                        _ = v21
                    } else {
                        v1 = append(v1, ("hello" + ""))
                        v16 = fi19
                        if ((!(-32768 >= 42)) || true) {
                            v17 = ((((v17 + v17) + "foo bar") + v17) + (((v17 + "\n\t") + "a") + v17))
                            var v22 float32 = 1.5
                            fmt.Println(fi20)
                            _ = v22
                        } else {
                            if false {
                                v6 -= (v6 * 2147483647)
                                var v23 int64 = -1
                                _ = v23
                            }
                            fmt.Println(v16, fi20, v16)
                            v17 = (v17 + v17)
                        }
                        switch "\n\t" {
                        case "hello":
                            v18 -= 255
                            v24 := GenS1{F0: 201, F1: 164}
                            var v25 int32 = -859
                            v26 := GenS1{F0: 21, F1: 10}
                            if (-2.25 <= 1e10) {
                                var v27 uint32 = (v6 >> ((v18) & 63))
                                if (32767 != -32768) {
                                    v27 -= v27
                                    var v28 uint16 = (v5 ^ (v5 << ((1) & 63)))
                                    var v29 float32 = -2.25
                                    _ = v28
                                    _ = v29
                                } else {
                                    fmt.Println(fi19)
                                }
                                v30 := GenS1{F0: 2, F1: 48}
                                _ = v27
                                _ = v30
                            } else {
                                v3 = append(v3, (^(v6)))
                                var v31 uint8 = uint8(v5)
                                if ((v5 + 0) == v5) {
                                    v6 = (^(v6))
                                } else {
                                    if true {
                                        var v32 float64 = 3.14159
                                        v16 = 2147483647
                                        _ = v32
                                    }
                                }
                                fmt.Println(v31, v31, fi20)
                                var v33 int32 = 1
                                _ = v31
                                _ = v33
                            }
                            if (false || (548 <= (v18 & 0))) {
                                v5 = v5
                            }
                            _ = v24
                            _ = v25
                            _ = v26
                        default:
                            switch 32768 {
                            case -126:
                                var v34 uint32 = 4294967295
                                _ = v34
                            case 127:
                                fmt.Println(v18)
                            case 979:
                                if (121 <= 1) {
                                    var v35 uint8 = 128
                                    _ = v35
                                }
                            default:
                                v1 = append(v1, ("" + (("a" + ("hello" + (v17 + "\n\t"))) + "hello")))
                            }
                        }
                    }
                    var v36 uint16 = (v5 | 42)
                    _ = v36
                } else {
                    switch (fi20 % ((4294967295) | 1)) {
                    case 2147483647:
                        v3 = append(v3, v6)
                    case 128:
                        if (!true) {
                            if (-128 != 42) {
                                v5 = (v5 | (v5 / ((v5) | 1)))
                                if (0 < genFn1(((^(v6)) <= 65535), 983, true)) {
                                    v5 = (v5 >> (((v18 >> ((0) & 63))) & 63))
                                    switch "hello" {
                                    case "\n\t":
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                }
                                var v37 int16 = 42
                                _ = v37
                            }
                            v38 := GenS2{F0: 0, F1: -534, F2: v6}
                            fmt.Println(fi19, v17)
                            _ = v38
                        } else {
                            v18 -= (^(v18))
                        }
                    }
                    v17 = v17
                    fmt.Println(v17, v17)
                }
                if false {
                    delete(v4, 65535)
                    v1 = append(v1, ((v17 + ("" + v17)) + v17))
                    var v39 int32 = 1
                    _ = v39
                } else {
                    var v40 string = ("" + "\n\t")
                    var v41 uint16 = (v5 + v5)
                    if (true && (42 == 0)) {
                        switch v16 {
                        case 127:
                            if true {
                                v6 -= (^(v6))
                                v5 -= 2
                            } else {
                                if false {
                                    if true {
                                        v42 := GenS3{F0: fi20, F1: 1000}
                                        var v43 int64 = 0
                                        var v44 int32 = -2147483648
                                        _ = v42
                                        _ = v43
                                        _ = v44
                                    } else {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        switch (v16 / ((fi20) | 1)) {
                                        case 128:
                                            if len(v3) > 0 {
                                                v3[0] = genFn0()
                                            }
                                        case 9223372036854775807:
                                            v17 = "\n\t"
                                        case 32768:
                                            var v45 float32 = 3.14159
                                            switch int(v6) {
                                            case -318:
                                                fmt.Println(v41)
                                            case -692:
                                                var v46 uint = 18446744073709551615
                                                _ = v46
                                            case 486:
                                                var v47 uint16 = v41
                                                _ = v47
                                            }
                                            _ = v45
                                        }
                                    }
                                    if (-128 < 42) {
                                        v2[(v6 & v6)] = v17
                                    }
                                    fmt.Println(v18)
                                    if true {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    } else {
                                        v48 := v4[(v6 % (((v6 % ((v6) | 1))) | 1))]
                                        _ = v48
                                    }
                                    v49 := GenS4{F0: 42, F1: -72}
                                    _ = v49
                                } else {
                                    if ("" > v40) {
                                        fmt.Println(fi20, v6)
                                    } else {
                                        var v50 float32 = -2.25
                                        var v51 string = v40
                                        _ = v50
                                        _ = v51
                                    }
                                }
                                if true {
                                    switch "\n\t" {
                                    case "hello":
                                        _ = len(v4)
                                    case "":
                                        if len(v3) > 0 {
                                            v3[0] = v6
                                        }
                                    default:
                                        if ((3.14159 == 0.0) || true) {
                                            var v52 int8 = 89
                                            if (false || false) {
                                                var v53 float32 = 3.14159
                                                _ = v53
                                            }
                                            v54 := GenS5{F0: 42, F1: 2, F2: -717}
                                            _ = v52
                                            _ = v54
                                        } else {
                                            var v55 string = ("" + "\n\t")
                                            _ = v55
                                        }
                                    }
                                }
                                delete(v2, (v6 | v6))
                            }
                        case 32767:
                            if (v18 != v18) {
                                var v56 int64 = -130
                                var v57 int64 = v56
                                v56 = v56
                                v40 = v17
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v56
                                _ = v57
                            } else {
                                v1 = append(v1, (("hello" + ("\n\t" + (((v17 + ((v17 + v40) + v17)) + v17) + v40))) + "\n\t"))
                            }
                        case -2147483648:
                            fmt.Println(v18, v6, v18)
                        }
                    } else {
                        var v58 int16 = 44
                        var v59 bool = (2 >= 1)
                        if v59 {
                            fmt.Println(v6, v59)
                            v60 := GenS6{F0: (1.5 > 3.14159), F1: uint32(v18), F2: v16}
                            _ = v60
                        }
                        _ = v58
                        _ = v59
                    }
                    switch v16 {
                    case -375:
                        switch ("hello" + ("hello" + (v17 + ""))) {
                        case "foo bar":
                            switch (v17 + v17) {
                            case "":
                                fmt.Println(fi20, v41, v41)
                            case "foo bar":
                                v18 = (v18 % ((758) | 1))
                            case "hello":
                                if (9223372036854775807 >= 42) {
                                    fmt.Println(fi19, fi20, fi20)
                                    var v61 uint16 = v41
                                    _ = v61
                                } else {
                                    switch (v40 + "\n\t") {
                                    case "hello":
                                        var v62 uint = v18
                                        _ = v62
                                    }
                                    fmt.Println(fi19)
                                    fmt.Println(fi20, fi20, v5)
                                }
                            default:
                                v6 += (v6 / (((v6 ^ 800)) | 1))
                            }
                        default:
                            v63 := GenS7{F0: 1, F1: (4294967295 == 9223372036854775807)}
                            if (false && ((v18 ^ v18) > 2)) {
                                if false {
                                    v64 := GenS8{F0: 0, F1: 9223372036854775807}
                                    v65 := GenS9{F0: (v5 << ((65535) & 63)), F1: v40}
                                    _ = v64
                                    _ = v65
                                } else {
                                    var v66 float32 = (-2.25 / 3.14159)
                                    var v67 bool = true
                                    v16 = (fi20 | (v16 + 127))
                                    _ = v66
                                    _ = v67
                                }
                                if len(v1) > 0 {
                                    v1[0] = "\n\t"
                                }
                            }
                            var v68 uint32 = 2147483647
                            v4[genFn0()] = v17
                            var v69 int32 = -1
                            _ = v63
                            _ = v68
                            _ = v69
                        }
                    case -1:
                        v41 += v5
                    }
                    _ = v40
                    _ = v41
                }
                _ = fi20
            }
            _ = fi19
        }
        _ = v17
        _ = v18
    }
    var v70 int64 = -916
    v71 := GenS7{F0: -636, F1: false}
    for fi72 := 0; fi72 < 4; fi72++ {
        for fi73 := 0; fi73 < 10; fi73++ {
            var v74 int16 = 32767
            _ = fi73
            _ = v74
        }
        _ = fi72
    }
    var v75 uint8 = 125
    if (!(731 == 93)) {
        v76 := GenS8{F0: (v6 & genFn0()), F1: int64(v5)}
        _ = v76
    }
    _ = len(v2)
    if (((true && (!false)) || (v16 != (^(v16)))) && true) {
        switch v16 {
        case 42:
            v77 := GenS5{F0: (^(v6)), F1: (v6 | (^(v6))), F2: (v70 - v70)}
            _ = v77
        case 128:
            for fi78 := 0; fi78 < 2; fi78++ {
                v16 = 1
                _ = fi78
            }
        case 238:
            if true {
                var v79 bool = (!true)
                var v80 int = (v16 / ((v16) | 1))
                _ = v79
                _ = v80
            } else {
                if false {
                    if (false || true) {
                        _ = len(v2)
                    }
                    if false {
                        if false {
                            for fi81 := 0; fi81 < 9; fi81++ {
                                v16 = 32767
                                if true {
                                    v82 := GenS7{F0: -230, F1: ((v70 % (((v70 ^ int64(v5))) | 1)) == v70)}
                                    var v83 string = "\n\t"
                                    _ = v82
                                    _ = v83
                                } else {
                                    fmt.Println(v70, v75)
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    var v84 uint16 = (v5 & (v5 ^ v5))
                                    _ = v84
                                }
                                v85 := GenS7{F0: -2147483648, F1: ((304 == 0) && (!(!(0.1 >= 1.5))))}
                                _ = fi81
                                _ = v85
                            }
                            var v86 uint16 = v5
                            _ = v86
                        } else {
                            for fi87 := 0; fi87 < 7; fi87++ {
                                var v88 float32 = ((1e10 / -2.25) / ((((1e10 + 0.0) * 0.1) - (1.5 - 0.0)) * 0.1))
                                fmt.Println(fi87, v70, v75)
                                _ = fi87
                                _ = v88
                            }
                        }
                    } else {
                        var v89 uint8 = v75
                        v90 := GenS5{F0: (v6 / ((v6) | 1)), F1: (v6 + genFn0()), F2: (v70 / (((v70 / ((0) | 1))) | 1))}
                        _ = v89
                        _ = v90
                    }
                    fmt.Println(v5, v6, v70)
                }
                fmt.Println(v6, v5)
            }
        default:
            v91 := v2[0]
            _ = v91
        }
        var v92 bool = (v16 != (^(v16)))
        _ = v92
    }
    v6 = v6
    if len(v3) > 0 {
        v3[0] = (v6 + (v6 | (v6 + 2147483647)))
    }
    v70 = (v70 | (v70 ^ v70))
    if ((^(v6)) != v6) {
        switch ("foo bar" + "hello") {
        case "":
            fmt.Println(v75, v5, v5)
        case "\n\t":
            delete(v2, (v6 ^ 0))
        case "hello":
            v6 = v6
        }
        var v93 int64 = v70
        _ = v93
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v70
    _ = v71
    _ = v75
}
