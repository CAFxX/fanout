package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint; F1 int }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint16; F1 uint8; F2 uint }
type GenS12 struct { F0 uint8 }
func genFn0(p0 int32, p1 int) int32 {
    var v5 int32 = (p0 & 763)
    switch "" {
    case "foo bar":
        var v6 int8 = 100
        _ = v6
    case "\n\t":
        fmt.Println(p1, p1, v5)
    case "":
        switch "\n\t" {
        case "\n\t":
            fmt.Println(v5, p1, v5)
        case "foo bar":
            p0 += v5
        case "a":
            for fi7 := 0; fi7 < 18; fi7++ {
                for fi8 := 0; fi8 < 4; fi8++ {
                    if false {
                        v5 = p0
                        if (true || (true || (!true))) {
                            var v9 int8 = -128
                            var v10 int16 = -785
                            _ = v10
                            _ = v9
                        } else {
                            switch (^(fi7)) {
                            case 0:
                                var v11 float64 = 1.5
                                _ = v11
                            case 32767:
                                p0 = (^(v5))
                            case -862:
                                switch (-(fi7)) {
                                case 32767:
                                    p0 += v5
                                case 9223372036854775807:
                                    v5 = -2147483648
                                case -89:
                                    var v12 uint64 = 889
                                    _ = v12
                                default:
                                    p1 = (p1 >> ((2) & 63))
                                }
                            }
                            p0 -= (p0 % ((p0) | 1))
                            v13 := GenS1{F0: true, F1: 1}
                            _ = v13
                        }
                        if false {
                            p1 = 1
                            fmt.Println(fi7, p1)
                        } else {
                            if false {
                                v14 := GenS2{F0: 0, F1: 42}
                                p0 -= (v5 * 2147483647)
                                fmt.Println(p0, p1, fi7)
                                _ = v14
                            }
                        }
                        p1 -= fi8
                        if (0 > -450) {
                            var v15 uint8 = 255
                            _ = v15
                        }
                    }
                    fmt.Println(v5, fi8)
                    switch (p1 - (fi8 * (p1 % (((fi7 ^ (fi7 | p1))) | 1)))) {
                    case 127:
                        var v16 int = (p1 * -9223372036854775808)
                        _ = v16
                    case -901:
                        v17 := GenS3{F0: (false && true), F1: 1000}
                        _ = v17
                    case 885:
                        if true {
                            var v18 bool = (288 != 471)
                            _ = v18
                        }
                        switch "hello" {
                        case "foo bar":
                            fmt.Println(fi8, v5, v5)
                        case "":
                            switch "\n\t" {
                            case "hello":
                                v19 := GenS3{F0: false, F1: 32767}
                                _ = v19
                            case "\n\t":
                                p0 = 1
                            case "":
                                switch (("a" + "\n\t") + "") {
                                case "":
                                    if false {
                                        v20 := GenS2{F0: 1, F1: 42}
                                        v21 := GenS4{F0: 42}
                                        var v22 int64 = 9223372036854775807
                                        _ = v20
                                        _ = v21
                                        _ = v22
                                    } else {
                                        switch "" {
                                        case "a":
                                            switch "a" {
                                            case "hello":
                                                fmt.Println(fi7)
                                            case "a":
                                                switch ("a" + "a") {
                                                case "a":
                                                    fmt.Println(p0, p0, p1)
                                                case "\n\t":
                                                    p0 = p0
                                                default:
                                                    v23 := GenS5{F0: "a", F1: 322}
                                                    _ = v23
                                                }
                                            case "\n\t":
                                                p0 = v5
                                                v24 := []int{42, fi7, fi7, -677}
                                                p0 = 0
                                                v25 := GenS2{F0: 207, F1: 65535}
                                                switch "\n\t" {
                                                case "\n\t":
                                                    if (-9223372036854775808 != 42) {
                                                        var v26 int8 = 0
                                                        v27 := GenS4{F0: 0}
                                                        _ = v26
                                                        _ = v27
                                                    } else {
                                                        var v28 uint = 2
                                                        _ = len(v24)
                                                        _ = v28
                                                    }
                                                default:
                                                    switch ("hello" + ("foo bar" + "")) {
                                                    case "foo bar":
                                                        p0 = 2147483647
                                                    case "\n\t":
                                                        fmt.Println(fi8, fi7)
                                                    }
                                                }
                                                _ = v24
                                                _ = v25
                                            }
                                        case "foo bar":
                                            v29 := GenS6{F0: 1, F1: 2, F2: 2}
                                            _ = v29
                                        }
                                        var v30 int16 = 0
                                        _ = v30
                                    }
                                case "\n\t":
                                    v31 := GenS7{F0: 18446744073709551615}
                                    _ = v31
                                default:
                                    v32 := GenS8{F0: 703}
                                    _ = v32
                                }
                            }
                        default:
                            var v33 uint16 = 65535
                            _ = v33
                        }
                    default:
                        v34 := []int{}
                        _ = v34
                    }
                    _ = fi8
                }
                v35 := []int{}
                _ = fi7
                _ = v35
            }
        default:
            fmt.Println(v5, v5, p1)
        }
    }
    for fi36 := 0; fi36 < 20; fi36++ {
        v37 := GenS9{F0: 255, F1: fi36}
        if false {
            var v38 float32 = (-2.25 + -2.25)
            var v39 int8 = -128
            _ = v38
            _ = v39
        } else {
            v40 := []int{}
            if len(v40) > 0 {
                v40[0] = 4294967295
            }
            _ = v40
        }
        p1 = 933
        _ = fi36
        _ = v37
    }
    for fi41 := 0; fi41 < 12; fi41++ {
        var v42 string = "hello"
        _ = fi41
        _ = v42
    }
    switch "" {
    case "hello":
        for fi43 := 0; fi43 < 17; fi43++ {
            v5 = v5
            _ = fi43
        }
    case "\n\t":
        var v44 uint = 4294967295
        _ = v44
    case "foo bar":
        var v45 uint64 = 978
        _ = v45
    default:
        v46 := []int{(p1 - (p1 ^ (p1 << ((18446744073709551615) & 63)))), (p1 >> ((18446744073709551615) & 63))}
        _ = v46
    }
    _ = p0
    _ = p1
    _ = v5
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 854
    v3 := []uint32{(^(v2))}
    v4 := []string{}
    switch 32767 {
    case -304:
        if true {
            v47 := GenS10{F0: 1}
            _ = v47
        } else {
            v3 = append(v3, 0)
            v48 := GenS9{F0: 86, F1: 79}
            v2 = (v2 & (v2 * v2))
            _ = v48
        }
    case 42:
        v1[v2] = (v2 ^ (v2 >> ((0) & 63)))
    case 9223372036854775807:
        if true {
            v2 = (^(v2))
        } else {
            if true {
                var v49 uint8 = 36
                _ = v49
            } else {
                var v50 uint32 = v2
                _ = v50
            }
            for fi51 := 0; fi51 < 19; fi51++ {
                if true {
                    if (42 < 42) {
                        var v52 int8 = 127
                        v53 := GenS10{F0: 1}
                        v54 := GenS4{F0: 206}
                        _ = v52
                        _ = v53
                        _ = v54
                    } else {
                        if (!(2 > 0)) {
                            for fi55 := 0; fi55 < 18; fi55++ {
                                v2 = (v2 * (^(v2)))
                                v2 += 4294967295
                                _ = fi55
                            }
                            for fi56 := 0; fi56 < 6; fi56++ {
                                var v57 uint64 = 450
                                _ = fi56
                                _ = v57
                            }
                            delete(v1, v2)
                        } else {
                            var v58 int8 = 64
                            _ = v58
                        }
                        var v59 uint8 = 255
                        var v60 int = fi51
                        _ = v59
                        _ = v60
                    }
                    for fi61 := 0; fi61 < 11; fi61++ {
                        fmt.Println(v2, v2)
                        if (!(false || ((-128 < 42) || (0 != 42)))) {
                            fmt.Println(v2, v2, fi61)
                            if false {
                                var v62 int = 961
                                var v63 int8 = int8(v62)
                                _ = v62
                                _ = v63
                            } else {
                                var v64 uint64 = 2
                                _ = v64
                            }
                            v2 = v2
                        } else {
                            if ((false || true) || true) {
                                var v65 string = "a"
                                _ = v65
                            }
                            fmt.Println(fi61, fi61, fi61)
                            if false {
                                var v66 float32 = (1.5 - ((1.5 * ((1e10 - -2.25) / -2.25)) * float32(0.1)))
                                var v67 int = (^(fi51))
                                var v68 bool = false
                                _ = v66
                                _ = v67
                                _ = v68
                            } else {
                                if true {
                                    if (893 <= 156) {
                                        switch "foo bar" {
                                        case "\n\t":
                                            switch (fi51 - -9223372036854775808) {
                                            case -294:
                                                var v69 float32 = ((-2.25 + -2.25) + float32((0.0 / ((0.1 + (1e10 / (1e10 * 0.1))) - (3.14159 + 1.5)))))
                                                v2 = v2
                                                _ = v69
                                            case -961:
                                                switch (fi51 >> ((42) & 63)) {
                                                case 9223372036854775807:
                                                    switch "hello" {
                                                    case "hello":
                                                        fmt.Println(fi51)
                                                    case "\n\t":
                                                        var v70 int32 = -1
                                                        _ = v70
                                                    }
                                                case 1:
                                                    if len(v4) > 0 {
                                                        v4[0] = ("\n\t" + "")
                                                    }
                                                    var v71 int8 = 17
                                                    if true {
                                                        if (!false) {
                                                            fmt.Println(fi61, v2)
                                                            v2 += v2
                                                        }
                                                    }
                                                    _ = v71
                                                case 776:
                                                    v2 += 42
                                                }
                                            }
                                        default:
                                            v4 = append(v4, "hello")
                                        }
                                        var v72 bool = (false && (2 > 255))
                                        _ = v72
                                    } else {
                                        _ = len(v1)
                                    }
                                }
                            }
                        }
                        _ = fi61
                    }
                    var v73 uint = 180
                    _ = v73
                } else {
                    v1[2147483647] = v2
                }
                var v74 uint32 = (v2 * 42)
                _ = fi51
                _ = v74
            }
        }
    }
    v75 := GenS4{F0: 544}
    switch "foo bar" {
    case "a":
        v76 := GenS11{F0: 65535, F1: 0, F2: 2}
        _ = v76
    default:
        switch 0 {
        case -9223372036854775808:
            var v77 uint16 = 32767
            _ = v77
        }
    }
    var v78 string = "\n\t"
    v78 = "foo bar"
    for fi79 := 0; fi79 < 8; fi79++ {
        v80 := GenS10{F0: 156}
        v81 := GenS12{F0: 202}
        var v82 int64 = -1
        _ = fi79
        _ = v80
        _ = v81
        _ = v82
    }
    v78 = v78
    switch 42 {
    case 2147483647:
        for fi83 := 0; fi83 < 9; fi83++ {
            var v84 bool = false
            _ = fi83
            _ = v84
        }
    default:
        switch 4294967295 {
        case 9223372036854775807:
            var v85 uint16 = 2
            _ = v85
        case 117:
            var v86 uint = 982
            _ = v86
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v75
    _ = v78
}
