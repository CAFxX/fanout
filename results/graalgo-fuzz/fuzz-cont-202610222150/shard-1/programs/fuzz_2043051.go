package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 bool; F2 int8 }
type GenS3 struct { F0 int32; F1 int; F2 uint8 }
type GenS4 struct { F0 uint8; F1 uint; F2 uint }
type GenS5 struct { F0 uint32; F1 int64; F2 string }
type GenS6 struct { F0 uint32; F1 int8 }
type GenS7 struct { F0 uint64; F1 int }
type GenS8 struct { F0 uint16; F1 bool; F2 uint }
func genFn0(p0 bool) uint {
    var v7 uint32 = 1
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 10; fi9++ {
            v7 += v7
            if ((1.5 - (float64(((float32(-2.25) * 1.5) - (1.5 - 1.5))) - ((float64(0.1) / 3.14159) - 0.0))) <= (1.5 - 0.1)) {
                if false {
                    var v10 int64 = 0
                    if true {
                        if p0 {
                            var v11 string = ("\n\t" + (("foo bar" + ("\n\t" + "")) + "a"))
                            _ = v11
                        }
                        switch 32767 {
                        case -1:
                            v12 := GenS1{F0: v7}
                            _ = v12
                        case 0:
                            v13 := GenS2{F0: (fi8 / (((fi9 << ((383) & 63))) | 1)), F1: (p0 && false), F2: 42}
                            _ = v13
                        }
                    } else {
                        fmt.Println(v7)
                        v14 := []int{fi8, fi8}
                        _ = v14
                    }
                    var v15 int8 = -75
                    _ = v10
                    _ = v15
                } else {
                    var v16 bool = true
                    _ = v16
                }
                v7 -= v7
            }
            _ = fi9
        }
        var v17 uint8 = 159
        _ = fi8
        _ = v17
    }
    _ = p0
    _ = v7
    return 0
}
func genFn1(p0 int, p1 uint64) uint32 {
    for fi18 := 0; fi18 < 19; fi18++ {
        v19 := []int{}
        _ = fi18
        _ = v19
    }
    if (true || false) {
        var v20 int32 = -504
        v20 += (v20 * (^(v20)))
        _ = v20
    }
    var v21 uint32 = 2147483647
    if true {
        switch p0 {
        case -981:
            var v22 bool = false
            _ = v22
        default:
            var v23 uint64 = 4294967295
            _ = v23
        }
        if (!("" < "foo bar")) {
            for fi24 := 0; fi24 < 16; fi24++ {
                if false {
                    switch (-(p0)) {
                    case 208:
                        var v25 int16 = -32768
                        _ = v25
                    default:
                        for fi26 := 0; fi26 < 13; fi26++ {
                            v27 := []int{(fi24 - (p0 + fi24)), fi24, p0}
                            fmt.Println(fi24)
                            v28 := GenS3{F0: -38, F1: fi24, F2: 156}
                            v29 := GenS3{F0: 1, F1: (fi24 - fi24), F2: 127}
                            fmt.Println(p1)
                            var v30 uint8 = uint8(v21)
                            p0 = (fi26 - int(p1))
                            _ = fi26
                            _ = v27
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                    }
                    v31 := GenS1{F0: (^(v21))}
                    if true {
                        var v32 uint8 = 38
                        if false {
                            v21 -= (v21 ^ v21)
                            v33 := GenS3{F0: -683, F1: p0, F2: (^(v32))}
                            for fi34 := 0; fi34 < 4; fi34++ {
                                fmt.Println(fi24, p1)
                                v35 := GenS3{F0: 2147483647, F1: p0, F2: v32}
                                v36 := []int{(p0 & -2147483648), p0}
                                fmt.Println(fi34, v21)
                                _ = fi34
                                _ = v35
                                _ = v36
                            }
                            _ = v33
                        } else {
                            if false {
                                var v37 float32 = 1.5
                                v38 := GenS4{F0: (^(v32)), F1: genFn0(true), F2: 42}
                                for fi39 := 0; fi39 < 13; fi39++ {
                                    p0 += (fi39 >> ((genFn0(true)) & 63))
                                    var v40 uint64 = (p1 + 1)
                                    v41 := []int{fi39, (fi24 ^ fi39), p0}
                                    _ = fi39
                                    _ = v40
                                    _ = v41
                                }
                                _ = v37
                                _ = v38
                            } else {
                                var v42 string = ("hello" + "foo bar")
                                var v43 uint = genFn0(false)
                                _ = v42
                                _ = v43
                            }
                        }
                        v32 = (^(v32))
                        _ = v32
                    }
                    _ = v31
                }
                _ = fi24
            }
            v44 := []int{370, (^(p0)), p0, p0}
            for fi45 := 0; fi45 < 8; fi45++ {
                v44 = append(v44, fi45)
                p1 -= (^(p1))
                _ = fi45
            }
            _ = v44
        } else {
            p0 += (p0 << ((18446744073709551615) & 63))
            var v46 float64 = 1.5
            if true {
                v47 := GenS5{F0: v21, F1: -645, F2: "foo bar"}
                if true {
                    var v48 int16 = 0
                    var v49 int8 = 63
                    _ = v48
                    _ = v49
                } else {
                    fmt.Println(v21, p1, v21)
                    switch (p0 & 42) {
                    case 32767:
                        v50 := GenS5{F0: uint32(p1), F1: -9223372036854775808, F2: ""}
                        _ = v50
                    }
                }
                _ = v47
            } else {
                var v51 string = "foo bar"
                if false {
                    if (false || false) {
                        for fi52 := 0; fi52 < 5; fi52++ {
                            var v53 string = v51
                            if (109 >= -128) {
                                v54 := GenS2{F0: -106, F1: true, F2: 42}
                                if false {
                                    for fi55 := 0; fi55 < 12; fi55++ {
                                        var v56 uint = 2
                                        var v57 int32 = 2147483647
                                        var v58 uint32 = v21
                                        _ = fi55
                                        _ = v56
                                        _ = v57
                                        _ = v58
                                    }
                                    for fi59 := 0; fi59 < 6; fi59++ {
                                        v60 := []int{36, p0}
                                        fmt.Println(v46)
                                        _ = fi59
                                        _ = v60
                                    }
                                    for fi61 := 0; fi61 < 1; fi61++ {
                                        if (!false) {
                                            if true {
                                                if true {
                                                    if (false && (42 < 81)) {
                                                        switch (p0 + (-(p0))) {
                                                        case 625:
                                                            switch 32767 {
                                                            case -9223372036854775808:
                                                                fmt.Println(p1, fi52)
                                                            default:
                                                                fmt.Println(v51, fi52)
                                                            }
                                                        case -9223372036854775808:
                                                            fmt.Println(v51, v51, p1)
                                                        case 128:
                                                            v46 -= v46
                                                        }
                                                    } else {
                                                        switch (p0 & fi61) {
                                                        case 9223372036854775807:
                                                            if true {
                                                                if (true && (22 <= 69)) {
                                                                    v62 := GenS3{F0: 2147483647, F1: fi52, F2: 127}
                                                                    fmt.Println(p0, v53)
                                                                    v63 := GenS6{F0: (v21 ^ (v21 >> ((0) & 63))), F1: int8(p0)}
                                                                    _ = v62
                                                                    _ = v63
                                                                }
                                                            } else {
                                                                switch 32768 {
                                                                case 963:
                                                                    var v64 int16 = -1
                                                                    _ = v64
                                                                case 164:
                                                                    v65 := GenS5{F0: v21, F1: -9223372036854775808, F2: "\n\t"}
                                                                    _ = v65
                                                                case 743:
                                                                    v53 = ""
                                                                }
                                                                switch "hello" {
                                                                case "hello":
                                                                    var v66 bool = (!true)
                                                                    _ = v66
                                                                }
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    v53 = "\n\t"
                                                    v67 := GenS7{F0: (^(p1)), F1: 13}
                                                    _ = v67
                                                }
                                                p0 += fi52
                                                if true {
                                                    v68 := []int{fi52, (fi52 | (p0 | (-(fi61))))}
                                                    _ = v68
                                                }
                                                switch v53 {
                                                case "a":
                                                    v69 := []int{2147483647}
                                                    _ = v69
                                                case "\n\t":
                                                    v53 = (((v53 + v51) + v53) + ("" + v51))
                                                case "":
                                                    switch p0 {
                                                    case -9223372036854775808:
                                                        v70 := []int{128, fi52}
                                                        _ = v70
                                                    case 169:
                                                        var v71 bool = true
                                                        _ = v71
                                                    case -1:
                                                        v72 := []int{p0}
                                                        _ = v72
                                                    }
                                                }
                                                v73 := GenS7{F0: 0, F1: fi61}
                                                _ = v73
                                            } else {
                                                fmt.Println(fi52, p1, v51)
                                            }
                                        }
                                        var v74 int64 = -1
                                        _ = fi61
                                        _ = v74
                                    }
                                }
                                _ = v54
                            }
                            for fi75 := 0; fi75 < 19; fi75++ {
                                v51 = ""
                                v76 := GenS1{F0: (v21 * v21)}
                                var v77 int8 = 100
                                _ = fi75
                                _ = v76
                                _ = v77
                            }
                            _ = fi52
                            _ = v53
                        }
                        var v78 uint64 = 18446744073709551615
                        _ = v78
                    }
                }
                _ = v51
            }
            _ = v46
        }
        for fi79 := 0; fi79 < 19; fi79++ {
            var v80 uint8 = 128
            if (!false) {
                var v81 int32 = 457
                _ = v81
            } else {
                v82 := GenS5{F0: v21, F1: -1, F2: "\n\t"}
                v83 := []int{(fi79 << ((791) & 63)), -305}
                v84 := GenS7{F0: (^(p1)), F1: (fi79 / (((-(fi79))) | 1))}
                _ = v82
                _ = v83
                _ = v84
            }
            _ = fi79
            _ = v80
        }
    }
    for fi85 := 0; fi85 < 10; fi85++ {
        v86 := []int{p0, (fi85 & (p0 >> ((1) & 63))), fi85}
        for fi87 := 0; fi87 < 16; fi87++ {
            p0 = (fi85 % ((1) | 1))
            var v88 int8 = 1
            switch "" {
            case "foo bar":
                if len(v86) >= 2 {
                    v86 = v86[1:len(v86)]
                }
            case "a":
                _ = len(v86)
            }
            if false {
                var v89 int16 = 1000
                v90 := GenS8{F0: 42, F1: (v89 >= (v89 * v89)), F2: 255}
                var v91 int32 = 0
                v92 := GenS5{F0: (v21 & v21), F1: int64(fi85), F2: (("hello" + ("hello" + "")) + "")}
                _ = v89
                _ = v90
                _ = v91
                _ = v92
            }
            _ = fi87
            _ = v88
        }
        _ = fi85
        _ = v86
    }
    _ = p0
    _ = p1
    _ = v21
    return (v21 & 4294967295)
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]uint32)
    var v3 uint32 = 0
    v4 := make(map[int]int)
    var v5 int16 = 42
    var v6 uint = 42
    var v93 int = -316
    for fi94 := 0; fi94 < 20; fi94++ {
        v3 = 2
        _ = fi94
    }
    fmt.Println(v3, v93)
    v95 := []int{(v93 >> ((v6) & 63))}
    var v96 int8 = 100
    fmt.Println(v5, v96)
    for fi97 := 0; fi97 < 18; fi97++ {
        v98 := GenS3{F0: 0, F1: v93, F2: 18}
        v4[fi97] = 44
        _ = fi97
        _ = v98
    }
    if len(v95) >= 2 {
        v95 = v95[1:len(v95)]
    }
    v3 += v3
    var v99 string = ("foo bar" + "foo bar")
    for fi100 := 0; fi100 < 12; fi100++ {
        if ((9223372036854775807 >= -9223372036854775808) && ((^(v6)) < v6)) {
            v101 := GenS3{F0: -741, F1: fi100, F2: 42}
            v96 += (v96 & v96)
            _ = v101
        } else {
            var v102 int16 = 1
            _ = v102
        }
        _ = fi100
    }
    v103 := GenS4{F0: 0, F1: 237, F2: (^(v6))}
    _ = len(v95)
    if true {
        for fi104 := 0; fi104 < 16; fi104++ {
            v93 = (v93 >> ((4294967295) & 63))
            v95 = append(v95, -1)
            _ = fi104
        }
    } else {
        _ = len(v95)
        v105 := GenS5{F0: v3, F1: 9223372036854775807, F2: "\n\t"}
        v95 = append(v95, v93)
        _ = v105
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v93:", v93)
    fmt.Println("slicelen v95:", len(v95), v95)
    fmt.Println("v96:", v96)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v103
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v93
    _ = v95
    _ = v96
    _ = v99
}
