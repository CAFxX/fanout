package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint32; F1 int64; F2 int16 }
type GenS4 struct { F0 bool; F1 uint8; F2 int }
type GenS5 struct { F0 int32; F1 int16; F2 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint16; F1 string }
type GenS8 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS9 struct { F0 uint64; F1 uint32 }
type GenS10 struct { F0 int32 }
func genFn0(p0 int8, p1 uint8, p2 bool) uint64 {
    for fi4 := 0; fi4 < 8; fi4++ {
        v5 := []int{(fi4 ^ fi4), (-(fi4)), fi4, fi4}
        p2 = (!p2)
        _ = fi4
        _ = v5
    }
    switch "\n\t" {
    case "foo bar":
        switch 807 {
        case 0:
            v6 := GenS1{F0: 0}
            _ = v6
        default:
            switch 32767 {
            case -9223372036854775808:
                v7 := GenS1{F0: 65535}
                _ = v7
            case -89:
                p1 += p1
            default:
                var v8 int8 = (p0 >> ((18446744073709551615) & 63))
                _ = v8
            }
        }
    case "a":
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 uint8 = (p1 * 2)
            for fi11 := 0; fi11 < 14; fi11++ {
                var v12 int8 = -128
                var v13 float64 = 3.14159
                v14 := []int{fi11}
                var v15 string = "a"
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
                _ = v15
            }
            _ = fi9
            _ = v10
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]uint32)
    var v16 uint8 = 255
    switch 2147483647 {
    case 32768:
        var v17 uint = 1
        _ = v17
    case -554:
        v18 := GenS1{F0: 800}
        _ = v18
    }
    if true {
        v16 += (^(v16))
    } else {
        if (true && true) {
            if true {
                switch 42 {
                case 2147483647:
                    v19 := GenS2{F0: 1, F1: 332}
                    _ = v19
                case 9223372036854775807:
                    fmt.Println(v16)
                case 4294967295:
                    for fi20 := 0; fi20 < 4; fi20++ {
                        for fi21 := 0; fi21 < 9; fi21++ {
                            v16 = v16
                            if (1.5 >= (-2.25 / 1.0)) {
                                var v22 int = fi20
                                if (18446744073709551615 != 117) {
                                    if (3.14159 < 0.1) {
                                        switch (fi21 >> ((42) & 63)) {
                                        case 32768:
                                            v23 := GenS3{F0: 2147483648, F1: -9223372036854775808, F2: 32767}
                                            _ = v23
                                        default:
                                            var v24 int16 = 0
                                            _ = v24
                                        }
                                        fmt.Println(fi20, fi21)
                                        v2[268] = 582
                                    } else {
                                        var v25 int32 = -926
                                        _ = v25
                                    }
                                    fmt.Println(fi21)
                                    v26 := []int{(-(v22))}
                                    _ = v26
                                } else {
                                    if true {
                                        if (fi21 < v22) {
                                            var v27 uint32 = 642
                                            if true {
                                                v22 = 127
                                                var v28 int = -295
                                                _ = v28
                                            } else {
                                                var v29 float32 = ((1.5 * -2.25) / 0.1)
                                                v30 := GenS3{F0: v27, F1: -348, F2: 32767}
                                                v31 := GenS4{F0: (65535 < 32767), F1: uint8(v27), F2: 463}
                                                _ = v29
                                                _ = v30
                                                _ = v31
                                            }
                                            _ = v27
                                        }
                                    } else {
                                        fmt.Println(v16)
                                    }
                                    v16 += v16
                                }
                                _ = v22
                            }
                            _ = fi21
                        }
                        fmt.Println(v16, fi20)
                        _ = fi20
                    }
                }
                v32 := GenS5{F0: 42, F1: 1, F2: -2147483648}
                var v33 int16 = -1
                _ = v32
                _ = v33
            }
        } else {
            for fi34 := 0; fi34 < 9; fi34++ {
                fmt.Println(fi34)
                _ = fi34
            }
            v35 := GenS6{F0: 10}
            for fi36 := 0; fi36 < 7; fi36++ {
                switch fi36 {
                case 9223372036854775807:
                    v37 := GenS6{F0: 604}
                    _ = v37
                case -9223372036854775808:
                    v38 := GenS5{F0: -2147483648, F1: 32767, F2: fi36}
                    _ = v38
                case -903:
                    v16 = (v16 | (v16 & 0))
                }
                if (false && (18446744073709551615 >= 18446744073709551615)) {
                    v16 = (v16 ^ (v16 - (^(v16))))
                    v2[251] = 42
                } else {
                    var v39 int = fi36
                    if false {
                        v40 := GenS5{F0: -1, F1: -32768, F2: 128}
                        for fi41 := 0; fi41 < 6; fi41++ {
                            var v42 float64 = (1.5 / 1.5)
                            _ = fi41
                            _ = v42
                        }
                        v43 := GenS7{F0: 65535, F1: (("" + "hello") + "\n\t")}
                        _ = v40
                        _ = v43
                    } else {
                        if false {
                            var v44 uint = 4294967295
                            v44 += v44
                            var v45 int64 = 219
                            _ = v44
                            _ = v45
                        }
                        if (-2.25 < 1.5) {
                            v39 += 0
                        }
                    }
                    fmt.Println(v16)
                    _ = v39
                }
                _ = fi36
            }
            _ = v35
        }
    }
    var v46 string = ("foo bar" + "foo bar")
    var v47 uint8 = (v16 | (v16 / ((0) | 1)))
    fmt.Println(v16, v46)
    if ("\n\t" < v46) {
        var v48 uint32 = 42
        for fi49 := 0; fi49 < 15; fi49++ {
            for fi50 := 0; fi50 < 7; fi50++ {
                if (!false) {
                    v51 := GenS3{F0: v48, F1: 0, F2: 1}
                    var v52 float32 = 0.1
                    switch fi49 {
                    case 9223372036854775807:
                        v53 := GenS8{F0: 81, F1: 1, F2: 1}
                        _ = v53
                    case 32767:
                        delete(v1, "a")
                    }
                    _ = v51
                    _ = v52
                } else {
                    var v54 int = (fi50 >> ((65535) & 63))
                    if (0 < 2) {
                        var v55 float32 = 1.5
                        _ = v55
                    } else {
                        v56 := []int{(v54 >> ((1) & 63)), fi50, v54, v54}
                        var v57 uint = 400
                        _ = v56
                        _ = v57
                    }
                    _ = v54
                }
                if (("foo bar" <= "\n\t") && false) {
                    v58 := GenS3{F0: 42, F1: 245, F2: 422}
                    v16 -= v16
                    if true {
                        var v59 float32 = (1.5 * -2.25)
                        var v60 uint16 = 255
                        v47 = v47
                        _ = v59
                        _ = v60
                    } else {
                        switch "foo bar" {
                        case "a":
                            var v61 bool = (42 >= 42)
                            _ = v61
                        case "":
                            v62 := GenS7{F0: 721, F1: "hello"}
                            _ = v62
                        case "\n\t":
                            var v63 float32 = -2.25
                            _ = v63
                        default:
                            v64 := GenS9{F0: 4294967295, F1: v48}
                            _ = v64
                        }
                        v65 := GenS1{F0: (v48 & (v48 - (v48 + v48)))}
                        _ = v65
                    }
                    v16 = (v16 / (((^(v47))) | 1))
                    switch v46 {
                    case "hello":
                        fmt.Println(v46, v46)
                    case "a":
                        v66 := v1[("\n\t" + ("a" + v46))]
                        _ = v66
                    case "\n\t":
                        if false {
                            v67 := GenS6{F0: 1}
                            var v68 int = fi49
                            v48 = (^(v48))
                            _ = v67
                            _ = v68
                        }
                        v69 := GenS10{F0: -1}
                        var v70 int = (fi49 | (fi50 % (((fi49 << ((4294967295) & 63))) | 1)))
                        _ = v69
                        _ = v70
                    default:
                        v48 = 2147483648
                    }
                    var v71 int16 = 42
                    v71 += (v71 + v71)
                    _ = v58
                    _ = v71
                } else {
                    v72 := GenS10{F0: -853}
                    if false {
                        v47 -= (v47 + (v47 & uint8(v48)))
                        v73 := GenS5{F0: 1, F1: -193, F2: fi50}
                        var v74 int = 32768
                        var v75 uint64 = 940
                        var v76 uint8 = (v47 - v16)
                        _ = v73
                        _ = v74
                        _ = v75
                        _ = v76
                    }
                    _ = v72
                }
                switch fi50 {
                case 42:
                    switch v46 {
                    case "\n\t":
                        fmt.Println(v16)
                    case "hello":
                        var v77 uint64 = 0
                        _ = v77
                    default:
                        var v78 int8 = 127
                        _ = v78
                    }
                case 128:
                    var v79 int16 = 1000
                    _ = v79
                }
                _ = fi50
            }
            _ = fi49
        }
        _ = v48
    }
    if (false || (false || (true && (-1 > 0)))) {
        fmt.Println(v46, v46, v16)
        fmt.Println(v47, v47)
    }
    for fi80 := 0; fi80 < 15; fi80++ {
        v46 = v46
        v81 := GenS7{F0: 42, F1: ""}
        for fi82 := 0; fi82 < 7; fi82++ {
            v16 += v16
            _ = fi82
        }
        _ = fi80
        _ = v81
    }
    v3[174] = 979
    var v83 uint64 = 1
    fmt.Println(v83)
    if false {
        var v84 uint8 = v47
        var v85 float32 = (3.14159 + float32(1e10))
        _ = len(v1)
        _ = v84
        _ = v85
    } else {
        v16 -= (^(v16))
        if (!((951 <= -2147483648) && false)) {
            v46 = (("foo bar" + ("\n\t" + "hello")) + ((("a" + v46) + "foo bar") + ("hello" + v46)))
        }
    }
    var v86 uint8 = v16
    fmt.Println(v83)
    switch v46 {
    case "a":
        if false {
            fmt.Println(v47)
            if true {
                v87 := []int{508}
                _ = len(v87)
                v16 = 128
                _ = v87
            } else {
                var v88 int32 = 42
                _ = v88
            }
            v86 += (v16 >> ((41) & 63))
        } else {
            if (796 <= 1000) {
                v83 = genFn0(100, (v16 << ((0) & 63)), true)
                v89 := v3[42]
                _ = v89
            } else {
                var v90 uint16 = 2
                v91 := v3[65535]
                _ = v90
                _ = v91
            }
        }
    case "hello":
        v92 := v3[2]
        _ = v92
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v83:", v83)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v46
    _ = v47
    _ = v83
    _ = v86
}
