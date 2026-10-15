package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 bool }
type GenS2 struct { F0 uint64; F1 int8; F2 bool }
type GenS3 struct { F0 int64; F1 uint64; F2 string }
type GenS4 struct { F0 uint32; F1 uint8; F2 string }
type GenS5 struct { F0 uint8; F1 string; F2 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint64; F1 int }
type GenS9 struct { F0 int64; F1 int32; F2 uint16 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 bool; F1 int; F2 uint }
type GenS12 struct { F0 uint; F1 uint8 }
type GenS13 struct { F0 uint8 }
func genFn0(p0 int64, p1 int8) string {
    if (-1 <= 0) {
        v5 := []int{int(p1), 2147483647, -2147483648}
        if len(v5) > 0 {
            v5[0] = -9223372036854775808
        }
        var v6 int16 = -1
        _ = v5
        _ = v6
    } else {
        v7 := GenS1{F0: -9223372036854775808, F1: 65535, F2: false}
        fmt.Println(p0, p0)
        _ = v7
    }
    v8 := []int{32768, 0, -2147483648}
    v9 := GenS1{F0: 128, F1: 1, F2: (true || true)}
    switch ("\n\t" + "hello") {
    case "\n\t":
        if ((1.5 * -2.25) < (float64(0.0) / 1.5)) {
            v10 := GenS2{F0: 88, F1: p1, F2: (!true)}
            _ = v10
        }
    case "hello":
        for fi11 := 0; fi11 < 18; fi11++ {
            v12 := GenS3{F0: 42, F1: 712, F2: "\n\t"}
            _ = len(v8)
            fmt.Println(fi11, fi11)
            _ = fi11
            _ = v12
        }
    default:
        switch "hello" {
        case "":
            v13 := GenS3{F0: p0, F1: 0, F2: ("foo bar" + "foo bar")}
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return "hello"
}

func main() {
    var v1 uint = 1
    v2 := make(map[int]string)
    v3 := []string{("a" + ("hello" + "")), "\n\t", "\n\t"}
    v4 := []int{-1, 9223372036854775807, 4294967295}
    if true {
        switch "\n\t" {
        case "\n\t":
            v14 := GenS4{F0: 42, F1: 2, F2: "\n\t"}
            _ = v14
        case "hello":
            v1 = 18446744073709551615
        case "foo bar":
            if false {
                if false {
                    switch "hello" {
                    case "hello":
                        _ = len(v2)
                    case "foo bar":
                        for fi15 := 0; fi15 < 1; fi15++ {
                            if true {
                                v1 = 42
                                for fi16 := 0; fi16 < 8; fi16++ {
                                    var v17 uint64 = 4294967295
                                    var v18 bool = false
                                    if true {
                                        if v18 {
                                            v18 = ((fi16 - fi15) != (fi16 << ((65535) & 63)))
                                            if (fi16 != fi16) {
                                                v2[(fi16 / ((fi15) | 1))] = "\n\t"
                                            } else {
                                                var v19 uint32 = 65535
                                                if (2 == 127) {
                                                    _ = len(v3)
                                                    v17 -= (^(v17))
                                                } else {
                                                    v18 = (!false)
                                                    v3 = append(v3, ("" + "a"))
                                                    v20 := GenS2{F0: (^(v17)), F1: int8(v17), F2: ("a" > ("" + ("hello" + ("foo bar" + ""))))}
                                                    _ = v20
                                                }
                                                _ = v19
                                            }
                                        } else {
                                            var v21 int32 = -1
                                            v1 = (v1 << ((v1) & 63))
                                            fmt.Println(v18, fi16)
                                            _ = v21
                                        }
                                        if v18 {
                                            v18 = v18
                                            if true {
                                                v22 := GenS2{F0: (v17 & v17), F1: 36, F2: v18}
                                                var v23 int = fi15
                                                v18 = (((true && false) && v18) && (-32768 != 1))
                                                var v24 bool = v18
                                                _ = v22
                                                _ = v23
                                                _ = v24
                                            } else {
                                                var v25 int64 = -829
                                                _ = v25
                                            }
                                            v18 = true
                                            var v26 bool = v18
                                            fmt.Println(v17)
                                            _ = v26
                                        }
                                        v27 := GenS3{F0: -9223372036854775808, F1: v17, F2: (("\n\t" + ("foo bar" + "a")) + "hello")}
                                        _ = v27
                                    } else {
                                        switch "a" {
                                        case "\n\t":
                                            var v28 uint16 = 0
                                            _ = v28
                                        case "foo bar":
                                            if v18 {
                                                var v29 string = "a"
                                                if (65535 > 1) {
                                                    v30 := GenS5{F0: 1, F1: "\n\t", F2: (fi16 & 4294967295)}
                                                    v17 = (v17 / (((^(v17))) | 1))
                                                    _ = v30
                                                }
                                                _ = v29
                                            } else {
                                                switch 3 {
                                                case 355:
                                                    if (!((42 < 2147483647) || v18)) {
                                                        v17 = uint64(v1)
                                                    } else {
                                                        var v31 uint8 = 1
                                                        _ = v31
                                                    }
                                                    if (0 >= -1) {
                                                        v18 = v18
                                                        var v32 int8 = 106
                                                        _ = v32
                                                    }
                                                    v17 += v17
                                                }
                                                fmt.Println(v18, fi16, fi15)
                                            }
                                        case "a":
                                            v17 -= (v17 | (v17 + 1))
                                        }
                                        fmt.Println(v17, v1)
                                        v33 := GenS4{F0: 827, F1: 1, F2: "hello"}
                                        _ = v33
                                    }
                                    _ = fi16
                                    _ = v17
                                    _ = v18
                                }
                            } else {
                                v1 -= v1
                                if true {
                                    if true {
                                        v3 = append(v3, "")
                                    } else {
                                        if (1e10 == 3.14159) {
                                            for fi34 := 0; fi34 < 8; fi34++ {
                                                var v35 int = fi15
                                                v2[(fi15 / (((v35 / (((v35 & (v35 << ((42) & 63)))) | 1))) | 1))] = "a"
                                                switch "hello" {
                                                case "a":
                                                    v36 := GenS6{F0: 2147483648}
                                                    _ = v36
                                                case "\n\t":
                                                    switch 42 {
                                                    case 4294967295:
                                                        v35 = -2147483648
                                                    case 171:
                                                        v35 += 32767
                                                    default:
                                                        if true {
                                                            v35 -= (fi34 << (((v1 / ((v1) | 1))) & 63))
                                                            v37 := GenS5{F0: 254, F1: "hello", F2: 4294967295}
                                                            _ = v37
                                                        }
                                                    }
                                                default:
                                                    if true {
                                                        if (128 >= 2) {
                                                            if (!((fi15 & v35) > 4294967295)) {
                                                                v1 += v1
                                                                fmt.Println(fi34, v1, fi34)
                                                            } else {
                                                                var v38 int8 = 100
                                                                _ = v38
                                                            }
                                                            v2[v35] = "hello"
                                                        }
                                                        v39 := GenS2{F0: 42, F1: -1, F2: false}
                                                        if true {
                                                            if ((v1 >> ((2) & 63)) > (v1 >> (((v1 | v1)) & 63))) {
                                                                if (v35 <= fi34) {
                                                                    fmt.Println(fi34, v35)
                                                                } else {
                                                                    if (1 != 2) {
                                                                        if (!(false && (false && ((^(v1)) > v1)))) {
                                                                            delete(v2, fi34)
                                                                            fmt.Println(v35, fi15)
                                                                            delete(v2, 4294967295)
                                                                            v35 += fi34
                                                                            v35 -= fi15
                                                                        } else {
                                                                            var v40 uint64 = 1
                                                                            _ = v40
                                                                        }
                                                                        v35 = fi15
                                                                        fmt.Println(v1)
                                                                    }
                                                                    v41 := v2[fi15]
                                                                    _ = v41
                                                                }
                                                            }
                                                            v42 := GenS7{F0: (18446744073709551615 < 46)}
                                                            v1 += (^(v1))
                                                            _ = v42
                                                        } else {
                                                            v43 := GenS3{F0: int64(v35), F1: 42, F2: ("hello" + "foo bar")}
                                                            _ = v43
                                                        }
                                                        _ = v39
                                                    }
                                                }
                                                var v44 uint8 = 113
                                                if false {
                                                    switch ("\n\t" + "") {
                                                    case "a":
                                                        v45 := GenS5{F0: (v44 % ((v44) | 1)), F1: "foo bar", F2: 42}
                                                        _ = v45
                                                    }
                                                    v44 = 127
                                                }
                                                var v46 bool = false
                                                _ = fi34
                                                _ = v35
                                                _ = v44
                                                _ = v46
                                            }
                                            v1 -= v1
                                        }
                                        var v47 uint8 = 0
                                        _ = v47
                                    }
                                    switch fi15 {
                                    case 0:
                                        fmt.Println(v1, fi15)
                                    case 32767:
                                        if true {
                                            fmt.Println(v1, fi15)
                                            if ((-(fi15)) > fi15) {
                                                delete(v2, (fi15 ^ (fi15 | 9223372036854775807)))
                                                _ = len(v2)
                                                v48 := v2[32768]
                                                _ = v48
                                            }
                                        } else {
                                            if len(v3) > 0 {
                                                v3[0] = "\n\t"
                                            }
                                        }
                                    default:
                                        v49 := GenS8{F0: 42, F1: fi15}
                                        _ = v49
                                    }
                                    var v50 int64 = 369
                                    _ = v50
                                }
                                switch fi15 {
                                case -2147483648:
                                    var v51 int32 = 42
                                    _ = v51
                                case 2147483647:
                                    switch fi15 {
                                    case -2147483648:
                                        delete(v2, fi15)
                                    }
                                default:
                                    switch fi15 {
                                    case 32768:
                                        v52 := GenS8{F0: 18446744073709551615, F1: (fi15 << (((v1 << ((v1) & 63))) & 63))}
                                        _ = v52
                                    case -2147483648:
                                        for fi53 := 0; fi53 < 4; fi53++ {
                                            _ = len(v4)
                                            if true {
                                                _ = len(v4)
                                            } else {
                                                var v54 float32 = 1.5
                                                _ = v54
                                            }
                                            delete(v2, (fi53 ^ fi53))
                                            _ = fi53
                                        }
                                    case -658:
                                        switch (fi15 >> (((v1 >> ((v1) & 63))) & 63)) {
                                        case 128:
                                            fmt.Println(v1)
                                        case -2147483648:
                                            if false {
                                                var v55 bool = ((42 != 65535) && false)
                                                v55 = false
                                                _ = v55
                                            } else {
                                                var v56 uint = (v1 * v1)
                                                if true {
                                                    for fi57 := 0; fi57 < 8; fi57++ {
                                                        v58 := GenS6{F0: 1}
                                                        fmt.Println(fi15, fi57, fi15)
                                                        _ = fi57
                                                        _ = v58
                                                    }
                                                } else {
                                                    for fi59 := 0; fi59 < 4; fi59++ {
                                                        var v60 int16 = 267
                                                        _ = fi59
                                                        _ = v60
                                                    }
                                                    for fi61 := 0; fi61 < 18; fi61++ {
                                                        fmt.Println(fi61, v56)
                                                        if ((!(!false)) && ((1.5 - (-2.25 / 3.14159)) > 3.14159)) {
                                                            switch -9223372036854775808 {
                                                            case -831:
                                                                v56 += v1
                                                            case 815:
                                                                if (!(!false)) {
                                                                    var v62 uint64 = 905
                                                                    v63 := GenS7{F0: true}
                                                                    fmt.Println(v1, fi15, v1)
                                                                    _ = v62
                                                                    _ = v63
                                                                }
                                                            default:
                                                                v2[fi15] = ("foo bar" + "foo bar")
                                                            }
                                                        } else {
                                                            var v64 int32 = 2147483647
                                                            var v65 uint32 = 2147483647
                                                            var v66 uint8 = 218
                                                            _ = v64
                                                            _ = v65
                                                            _ = v66
                                                        }
                                                        _ = fi61
                                                    }
                                                    var v67 uint32 = 65535
                                                    _ = v67
                                                }
                                                v68 := GenS9{F0: 42, F1: -896, F2: 65535}
                                                _ = v56
                                                _ = v68
                                            }
                                        }
                                    default:
                                        for fi69 := 0; fi69 < 2; fi69++ {
                                            var v70 uint32 = 65535
                                            if true {
                                                v70 = (v70 & (v70 - v70))
                                                v71 := GenS5{F0: 190, F1: "a", F2: (fi69 & (fi69 << (((v1 % ((v1) | 1))) & 63)))}
                                                _ = v71
                                            } else {
                                                var v72 bool = false
                                                _ = v72
                                            }
                                            fmt.Println(v70, fi15)
                                            v73 := GenS10{F0: fi15}
                                            _ = fi69
                                            _ = v70
                                            _ = v73
                                        }
                                    }
                                }
                            }
                            var v74 int32 = 1
                            fmt.Println(fi15)
                            _ = fi15
                            _ = v74
                        }
                    default:
                        var v75 int64 = -580
                        _ = v75
                    }
                    if (-1 >= 2147483647) {
                        v76 := GenS11{F0: false, F1: -9223372036854775808, F2: v1}
                        _ = v76
                    } else {
                        fmt.Println(v1)
                        v77 := GenS12{F0: v1, F1: 86}
                        var v78 uint8 = 42
                        _ = v77
                        _ = v78
                    }
                }
                v1 -= (v1 - 255)
                v79 := GenS5{F0: 255, F1: ("" + "hello"), F2: -657}
                _ = v79
            } else {
                for fi80 := 0; fi80 < 11; fi80++ {
                    for fi81 := 0; fi81 < 15; fi81++ {
                        switch fi80 {
                        case 127:
                            var v82 uint = v1
                            _ = v82
                        case 2147483647:
                            _ = len(v2)
                        default:
                            if true {
                                _ = len(v3)
                                v83 := GenS13{F0: 2}
                                _ = v83
                            }
                        }
                        v84 := GenS3{F0: 0, F1: 282, F2: ("a" + ("" + "\n\t"))}
                        v1 += v1
                        _ = fi81
                        _ = v84
                    }
                    _ = fi80
                }
                if len(v3) > 0 {
                    v3[0] = ("foo bar" + "foo bar")
                }
            }
        }
        v85 := GenS6{F0: uint32(v1)}
        _ = v85
    }
    _ = len(v2)
    if true {
        v1 += v1
        for fi86 := 0; fi86 < 19; fi86++ {
            v1 = v1
            _ = fi86
        }
        var v87 float64 = 0.0
        _ = v87
    }
    var v88 int = 2147483647
    var v89 uint = v1
    for fi90 := 0; fi90 < 13; fi90++ {
        v88 += 0
        v4 = append(v4, (fi90 & v88))
        for fi91 := 0; fi91 < 7; fi91++ {
            v88 = fi90
            v89 = v1
            delete(v2, v88)
            v88 = 32768
            _ = fi91
        }
        _ = fi90
    }
    var v92 uint64 = 1
    fmt.Println(v92, v1)
    var v93 int32 = -564
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v88
    _ = v89
    _ = v92
    _ = v93
}
