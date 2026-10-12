package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 int32; F1 int64 }
type GenS6 struct { F0 uint8; F1 int8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int; F1 int; F2 uint32 }
type GenS9 struct { F0 uint16; F1 int8 }
type GenS10 struct { F0 uint8; F1 uint8 }
type GenS11 struct { F0 int16; F1 int }
type GenS12 struct { F0 int8; F1 int8; F2 int8 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 bool; F1 uint16; F2 int8 }
type GenS15 struct { F0 uint32; F1 int16 }
type GenS16 struct { F0 bool }
func genFn0(p0 int8) int16 {
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := []int{fi7, 9223372036854775807}
        _ = fi7
        _ = v8
    }
    fmt.Println(p0)
    v9 := []int{-906, 1, 4294967295, -2147483648}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    var v10 int16 = 967
    _ = p0
    _ = v10
    _ = v9
    return v10
}

func main() {
    var v1 int8 = -128
    v2 := make(map[uint32]int)
    var v3 int16 = 1
    var v4 uint8 = 125
    v5 := make(map[string]uint32)
    v6 := make(map[string]bool)
    if (739 != 0) {
        switch "hello" {
        case "hello":
            for fi11 := 0; fi11 < 17; fi11++ {
                switch "" {
                case "hello":
                    v12 := []int{fi11, fi11}
                    _ = v12
                case "":
                    var v13 int16 = v3
                    _ = v13
                case "foo bar":
                    if false {
                        var v14 int16 = (v3 << ((1) & 63))
                        for fi15 := 0; fi15 < 3; fi15++ {
                            v16 := []int{fi11, (fi11 >> ((2) & 63)), -887}
                            var v17 int8 = v1
                            if true {
                                if ((false && (-821 < 1)) || false) {
                                    v18 := GenS1{F0: (^(v4)), F1: false}
                                    if (32767 >= 2) {
                                        v19 := GenS1{F0: v4, F1: false}
                                        _ = len(v16)
                                        _ = v19
                                    }
                                    v17 = v17
                                    var v20 int8 = 0
                                    v4 -= v4
                                    _ = v18
                                    _ = v20
                                }
                            } else {
                                var v21 float64 = (3.14159 * ((3.14159 * 1.5) * (3.14159 - float64((-2.25 - 0.1)))))
                                fmt.Println(v21)
                                v14 += (v3 * 846)
                                _ = v21
                            }
                            switch "" {
                            case "a":
                                v22 := GenS1{F0: (v4 / ((v4) | 1)), F1: true}
                                _ = v22
                            }
                            v23 := GenS2{F0: (v14 >> ((58) & 63))}
                            v1 -= (v1 ^ -100)
                            v24 := GenS3{F0: 42}
                            _ = fi15
                            _ = v16
                            _ = v17
                            _ = v23
                            _ = v24
                        }
                        v4 += (v4 >> ((4294967295) & 63))
                        _ = v14
                    } else {
                        v25 := GenS4{F0: (18446744073709551615 == 4294967295), F1: 2147483647}
                        v1 -= (v1 & v1)
                        v26 := GenS1{F0: 2, F1: (0 != 1)}
                        _ = v25
                        _ = v26
                    }
                default:
                    if true {
                        delete(v5, "hello")
                        v27 := GenS2{F0: genFn0((^(v1)))}
                        _ = v27
                    } else {
                        v28 := GenS5{F0: 42, F1: 1}
                        _ = v28
                    }
                }
                v1 = (v1 + (v1 + v1))
                v29 := []int{}
                _ = fi11
                _ = v29
            }
        case "":
            var v30 int32 = 1
            _ = v30
        case "\n\t":
            for fi31 := 0; fi31 < 6; fi31++ {
                if ((v1 >> ((18446744073709551615) & 63)) <= v1) {
                    v32 := GenS5{F0: 42, F1: 856}
                    v3 = (-(v3))
                    _ = v32
                } else {
                    var v33 uint = 965
                    v34 := GenS6{F0: (v4 & 103), F1: int8(v33)}
                    _ = v33
                    _ = v34
                }
                v35 := GenS6{F0: (^(v4)), F1: (v1 >> ((18446744073709551615) & 63))}
                var v36 uint32 = 0
                _ = fi31
                _ = v35
                _ = v36
            }
        }
        if (3.14159 <= 1.5) {
            switch "\n\t" {
            case "":
                v4 = v4
            case "hello":
                var v37 uint8 = v4
                _ = v37
            case "\n\t":
                for fi38 := 0; fi38 < 1; fi38++ {
                    var v39 uint32 = 1
                    _ = fi38
                    _ = v39
                }
            default:
                v40 := GenS2{F0: 957}
                _ = v40
            }
            var v41 float32 = (1e10 * 1e10)
            switch (("hello" + "hello") + ("hello" + "foo bar")) {
            case "\n\t":
                v4 -= (v4 + (v4 + (v4 - (v4 ^ (^(v4))))))
            case "":
                for fi42 := 0; fi42 < 8; fi42++ {
                    v43 := GenS7{F0: (!(1 > 18446744073709551615))}
                    var v44 int16 = (v3 & (v3 << ((623) & 63)))
                    switch -691 {
                    case 432:
                        if (!true) {
                            v45 := GenS8{F0: int(v44), F1: fi42, F2: 330}
                            for fi46 := 0; fi46 < 17; fi46++ {
                                v47 := GenS6{F0: (v4 ^ (v4 << ((2) & 63))), F1: (-(v1))}
                                var v48 int = -562
                                var v49 uint8 = v4
                                _ = fi46
                                _ = v47
                                _ = v48
                                _ = v49
                            }
                            _ = v45
                        }
                    case -2147483648:
                        fmt.Println(v3, v3)
                    case -721:
                        fmt.Println(v44, v41, v1)
                    }
                    _ = fi42
                    _ = v43
                    _ = v44
                }
            default:
                var v50 uint64 = 893
                _ = v50
            }
            _ = v41
        } else {
            if false {
                var v51 string = "\n\t"
                for fi52 := 0; fi52 < 9; fi52++ {
                    var v53 float32 = (1.5 * 0.0)
                    v54 := GenS5{F0: 2147483647, F1: -1}
                    for fi55 := 0; fi55 < 18; fi55++ {
                        var v56 uint64 = 0
                        if true {
                            v57 := GenS9{F0: 65535, F1: (v1 << ((0) & 63))}
                            v56 = (v56 << ((2) & 63))
                            v58 := []int{}
                            if true {
                                if false {
                                    var v59 string = v51
                                    var v60 bool = ((2 > 0) && (false && false))
                                    var v61 int32 = -978
                                    var v62 uint8 = v4
                                    _ = v59
                                    _ = v60
                                    _ = v61
                                    _ = v62
                                } else {
                                    var v63 uint8 = v4
                                    var v64 int = fi52
                                    var v65 int = fi55
                                    if true {
                                        v53 = 0.0
                                        var v66 bool = (!(true || ((false && true) && (1 == 2))))
                                        _ = v66
                                    }
                                    _ = v63
                                    _ = v64
                                    _ = v65
                                }
                            }
                            _ = v57
                            _ = v58
                        } else {
                            v53 += 0.1
                        }
                        delete(v2, 65535)
                        if false {
                            if (v1 >= -1) {
                                _ = len(v2)
                                v67 := GenS5{F0: 42, F1: 42}
                                v4 += (^(v4))
                                _ = v67
                            }
                            v6[v51] = false
                        }
                        _ = fi55
                        _ = v56
                    }
                    _ = fi52
                    _ = v53
                    _ = v54
                }
                for fi68 := 0; fi68 < 9; fi68++ {
                    if (2 > 18446744073709551615) {
                        for fi69 := 0; fi69 < 12; fi69++ {
                            v1 = v1
                            v51 = (v51 + v51)
                            v1 += v1
                            _ = fi69
                        }
                        v70 := []int{fi68}
                        v4 += 42
                        _ = v70
                    }
                    for fi71 := 0; fi71 < 7; fi71++ {
                        v72 := GenS10{F0: 1, F1: v4}
                        v73 := GenS6{F0: v4, F1: (v1 % ((int8(v4)) | 1))}
                        if ((v1 | v1) != v1) {
                            fmt.Println(fi71, fi71)
                        } else {
                            if false {
                                v4 = (v4 >> ((18446744073709551615) & 63))
                                if false {
                                    v4 += 90
                                    v74 := GenS3{F0: -782}
                                    v75 := []int{(-(fi71)), (fi68 ^ fi71)}
                                    _ = v74
                                    _ = v75
                                } else {
                                    v76 := []int{fi71, -205}
                                    _ = v76
                                }
                            }
                            v77 := []int{fi71, fi71}
                            _ = v77
                        }
                        _ = fi71
                        _ = v72
                        _ = v73
                    }
                    for fi78 := 0; fi78 < 4; fi78++ {
                        v3 += 0
                        _ = fi78
                    }
                    _ = fi68
                }
                _ = v51
            }
        }
    } else {
        for fi79 := 0; fi79 < 7; fi79++ {
            for fi80 := 0; fi80 < 12; fi80++ {
                v81 := []int{int(v4), int(v4), 1}
                var v82 int8 = 1
                v1 += v82
                _ = fi80
                _ = v81
                _ = v82
            }
            v4 = (^(v4))
            _ = fi79
        }
    }
    if true {
        switch ("\n\t" + "a") {
        case "":
            switch -1 {
            case 645:
                v83 := GenS4{F0: (!(327 < 255)), F1: 42}
                _ = v83
            case -746:
                if ((^(v4)) > (v4 >> ((178) & 63))) {
                    v4 = 201
                    var v84 bool = false
                    _ = v84
                } else {
                    v4 = v4
                }
            }
        default:
            for fi85 := 0; fi85 < 15; fi85++ {
                v86 := GenS11{F0: 0, F1: fi85}
                var v87 int16 = (^(v3))
                for fi88 := 0; fi88 < 16; fi88++ {
                    fmt.Println(v87, fi88)
                    if true {
                        fmt.Println(v4, v3, v1)
                        if false {
                            v1 = (v1 / (((^(v1))) | 1))
                            v89 := GenS12{F0: 1, F1: int8(fi88), F2: v1}
                            v90 := GenS13{F0: 100}
                            v3 += genFn0((v1 & (v1 ^ v1)))
                            var v91 string = "hello"
                            _ = v89
                            _ = v90
                            _ = v91
                        } else {
                            if true {
                                v92 := GenS14{F0: (!false), F1: 0, F2: v1}
                                v93 := GenS2{F0: (v87 << ((1) & 63))}
                                _ = v92
                                _ = v93
                            }
                            var v94 uint64 = uint64(v4)
                            var v95 float32 = 3.14159
                            _ = v94
                            _ = v95
                        }
                    }
                    delete(v5, ((("\n\t" + ("" + ("" + (("hello" + "hello") + "hello")))) + (("\n\t" + "") + "")) + "\n\t"))
                    _ = fi88
                }
                _ = fi85
                _ = v86
                _ = v87
            }
        }
    } else {
        v96 := GenS6{F0: (v4 - 68), F1: -66}
        v97 := v6["a"]
        if false {
            v98 := []int{9223372036854775807, 32767, 486}
            _ = len(v2)
            delete(v5, "foo bar")
            _ = v98
        } else {
            v99 := GenS6{F0: 42, F1: (v1 * v1)}
            _ = v99
        }
        _ = v96
        _ = v97
    }
    v1 = v1
    v1 = (v1 * int8(v3))
    if true {
        if false {
            var v100 int8 = 0
            fmt.Println(v4)
            _ = v100
        }
    } else {
        v101 := GenS15{F0: 2147483648, F1: (v3 >> ((uint(v4)) & 63))}
        var v102 int16 = v3
        _ = v101
        _ = v102
    }
    v103 := GenS6{F0: (v4 % ((128) | 1)), F1: v1}
    for fi104 := 0; fi104 < 6; fi104++ {
        v105 := GenS16{F0: ((true || true) || false)}
        for fi106 := 0; fi106 < 1; fi106++ {
            v5["\n\t"] = 2
            if (false || ((1.5 - 3.14159) == 0.0)) {
                var v107 string = "a"
                var v108 int = -9223372036854775808
                v107 = v107
                _ = v107
                _ = v108
            }
            _ = fi106
        }
        _ = fi104
        _ = v105
    }
    var v109 int8 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v109:", v109)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v103
    _ = v109
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
