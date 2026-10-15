package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 int; F1 uint16; F2 string }
type GenS5 struct { F0 string; F1 int64; F2 int32 }
type GenS6 struct { F0 uint8; F1 uint32; F2 string }
type GenS7 struct { F0 int8; F1 int; F2 int64 }
type GenS8 struct { F0 bool; F1 uint32 }
type GenS9 struct { F0 uint; F1 int64 }

func main() {
    var v1 bool = false
    v2 := make(map[uint32]string)
    var v3 uint = 278
    v4 := []int{929, -254, 32768, -9223372036854775808}
    v5 := make(map[string]int)
    var v6 uint16 = 65535
    var v7 float32 = -2.25
    fmt.Println(v6, v7, v1)
    var v8 int64 = 9223372036854775807
    v9 := GenS1{F0: 960}
    if (!v1) {
        switch -758 {
        case -1:
            v6 += (^(v6))
        case 192:
            v10 := GenS1{F0: 2147483647}
            _ = v10
        }
        v6 = 42
        var v11 string = ""
        _ = v11
    }
    if v1 {
        if v1 {
            if v1 {
                var v12 uint64 = 221
                for fi13 := 0; fi13 < 9; fi13++ {
                    _ = len(v4)
                    _ = fi13
                }
                _ = v12
            }
            var v14 uint16 = v6
            _ = v14
        } else {
            if (!(!(!v1))) {
                var v15 int32 = -2147483648
                switch 127 {
                case 2147483647:
                    v16 := GenS1{F0: 32767}
                    _ = v16
                default:
                    v7 = v7
                }
                if v1 {
                    switch "foo bar" {
                    case "foo bar":
                        if false {
                            v7 -= 1e10
                            v3 = (v3 >> ((v3) & 63))
                            v17 := GenS1{F0: 9223372036854775807}
                            _ = v17
                        } else {
                            v3 -= (^(v3))
                            v8 = v8
                        }
                    }
                } else {
                    switch "" {
                    case "\n\t":
                        v18 := GenS1{F0: 42}
                        _ = v18
                    case "foo bar":
                        fmt.Println(v3, v6)
                    }
                    v19 := GenS2{F0: 29, F1: v3}
                    var v20 int16 = 1
                    _ = v19
                    _ = v20
                }
                _ = v15
            }
            v1 = v1
            v21 := GenS1{F0: -2147483648}
            _ = v21
        }
        v7 += 3.14159
        v1 = v1
    } else {
        var v22 int16 = 32767
        var v23 uint32 = 2
        v22 = (v22 + 774)
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 12; fi24++ {
        v3 -= 1
        v25 := GenS2{F0: 0, F1: (v3 ^ (v3 << (((v3 * 18446744073709551615)) & 63)))}
        _ = fi24
        _ = v25
    }
    switch 127 {
    case -987:
        switch (("hello" + (("" + "a") + ("\n\t" + "\n\t"))) + "hello") {
        case "a":
            for fi26 := 0; fi26 < 8; fi26++ {
                v3 += (v3 * v3)
                if (!v1) {
                    v5["foo bar"] = fi26
                    switch fi26 {
                    case 0:
                        if v1 {
                            v4 = append(v4, (fi26 >> ((v3) & 63)))
                            _ = len(v4)
                        }
                    case 2147483647:
                        if ((v1 && (1 <= 1)) && (1 > 4294967295)) {
                            v27 := GenS3{F0: 100, F1: (fi26 ^ fi26)}
                            for fi28 := 0; fi28 < 4; fi28++ {
                                v7 += 1.5
                                fmt.Println(v7, v1)
                                _ = fi28
                            }
                            _ = v27
                        } else {
                            if false {
                                var v29 int32 = 0
                                if false {
                                    switch 289 {
                                    case -589:
                                        switch 1 {
                                        case -53:
                                            v30 := GenS4{F0: int(v6), F1: (v6 ^ (v6 * 1)), F2: "hello"}
                                            _ = v30
                                        case 1:
                                            if v1 {
                                                v3 += v3
                                                if (v1 && false) {
                                                    for fi31 := 0; fi31 < 15; fi31++ {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        _ = fi31
                                                    }
                                                } else {
                                                    v32 := GenS1{F0: fi26}
                                                    if v1 {
                                                        v33 := GenS4{F0: (fi26 * fi26), F1: (v6 >> ((v3) & 63)), F2: ""}
                                                        var v34 int32 = v29
                                                        var v35 float32 = (0.1 * v7)
                                                        _ = v33
                                                        _ = v34
                                                        _ = v35
                                                    } else {
                                                        v5[""] = (fi26 | -1)
                                                        var v36 bool = false
                                                        var v37 uint8 = 6
                                                        _ = v36
                                                        _ = v37
                                                    }
                                                    if false {
                                                        var v38 bool = (v3 <= v3)
                                                        v7 -= v7
                                                        _ = v38
                                                    } else {
                                                        v39 := v2[2147483648]
                                                        _ = v39
                                                    }
                                                    _ = v32
                                                }
                                                var v40 uint32 = 42
                                                _ = v40
                                            }
                                        default:
                                            if v1 {
                                                if len(v4) > 0 {
                                                    v4[0] = 42
                                                }
                                                var v41 int8 = 75
                                                var v42 uint16 = (v6 | v6)
                                                _ = v41
                                                _ = v42
                                            } else {
                                                v4 = append(v4, fi26)
                                                v43 := GenS1{F0: fi26}
                                                _ = v43
                                            }
                                        }
                                    case 490:
                                        fmt.Println(v7)
                                    case 1:
                                        fmt.Println(v8, v29)
                                    }
                                    if v1 {
                                        var v44 uint = 255
                                        _ = v44
                                    }
                                } else {
                                    var v45 int32 = v29
                                    if (27 >= 24) {
                                        fmt.Println(v3, v45)
                                        for fi46 := 0; fi46 < 12; fi46++ {
                                            if (!(0 >= 1)) {
                                                var v47 int64 = 0
                                                _ = v47
                                            } else {
                                                v6 -= v6
                                                v29 -= (v45 | (v29 % ((-110) | 1)))
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                v48 := GenS5{F0: "hello", F1: (v8 * 0), F2: 2147483647}
                                                _ = v48
                                            }
                                            v49 := GenS5{F0: "a", F1: (v8 - (v8 ^ -1)), F2: -2147483648}
                                            _ = fi46
                                            _ = v49
                                        }
                                    }
                                    _ = v45
                                }
                                v8 = (-(v8))
                                _ = v29
                            }
                            var v50 uint16 = 292
                            _ = v50
                        }
                    case 9223372036854775807:
                        var v51 int32 = 42
                        _ = v51
                    default:
                        for fi52 := 0; fi52 < 6; fi52++ {
                            v8 = v8
                            _ = fi52
                        }
                    }
                }
                if false {
                    var v53 uint8 = 1
                    _ = v53
                }
                _ = fi26
            }
        case "\n\t":
            _ = len(v2)
        case "hello":
            if (v1 && v1) {
                switch 4294967295 {
                case -944:
                    if (0.0 > (1e10 / 1.0)) {
                        var v54 int8 = -1
                        _ = v54
                    }
                case 32767:
                    var v55 float32 = v7
                    _ = v55
                case -896:
                    if true {
                        v56 := GenS6{F0: uint8(v3), F1: 2, F2: "hello"}
                        if v1 {
                            var v57 uint32 = 216
                            v57 = v57
                            var v58 uint = (^(v3))
                            _ = v57
                            _ = v58
                        } else {
                            if v1 {
                                v59 := GenS3{F0: -112, F1: -9223372036854775808}
                                switch "\n\t" {
                                case "":
                                    delete(v5, ("hello" + "hello"))
                                case "\n\t":
                                    var v60 uint16 = v6
                                    _ = v60
                                default:
                                    v3 = (v3 / ((uint(v6)) | 1))
                                }
                                _ = v59
                            } else {
                                v61 := GenS7{F0: 124, F1: 127, F2: (^(v8))}
                                v62 := GenS3{F0: -23, F1: 0}
                                var v63 int64 = v8
                                _ = v61
                                _ = v62
                                _ = v63
                            }
                            var v64 uint = v3
                            _ = v64
                        }
                        _ = v56
                    } else {
                        fmt.Println(v3, v1)
                        fmt.Println(v1)
                    }
                }
                var v65 uint16 = (v6 + (v6 & (v6 / ((42) | 1))))
                _ = v65
            } else {
                var v66 uint = v3
                _ = v66
            }
        default:
            if v1 {
                for fi67 := 0; fi67 < 8; fi67++ {
                    v68 := v2[65535]
                    v1 = (false && v1)
                    fmt.Println(v8, v6, v3)
                    _ = fi67
                    _ = v68
                }
            } else {
                for fi69 := 0; fi69 < 13; fi69++ {
                    v2[453] = "\n\t"
                    for fi70 := 0; fi70 < 4; fi70++ {
                        v71 := GenS4{F0: (fi70 % ((fi69) | 1)), F1: 312, F2: "\n\t"}
                        _ = fi70
                        _ = v71
                    }
                    var v72 int8 = -86
                    _ = fi69
                    _ = v72
                }
                for fi73 := 0; fi73 < 9; fi73++ {
                    v3 = (v3 - v3)
                    for fi74 := 0; fi74 < 14; fi74++ {
                        var v75 uint = v3
                        _ = fi74
                        _ = v75
                    }
                    v3 = 152
                    _ = fi73
                }
                for fi76 := 0; fi76 < 18; fi76++ {
                    for fi77 := 0; fi77 < 6; fi77++ {
                        fmt.Println(v1, fi76, fi76)
                        _ = fi77
                    }
                    _ = len(v4)
                    for fi78 := 0; fi78 < 11; fi78++ {
                        switch "hello" {
                        case "foo bar":
                            fmt.Println(fi76)
                        case "":
                            var v79 uint8 = 1
                            _ = v79
                        default:
                            v3 = (v3 & (^(v3)))
                        }
                        v80 := GenS8{F0: ((704 >= 255) && false), F1: 2147483648}
                        _ = fi78
                        _ = v80
                    }
                    _ = fi76
                }
            }
        }
    }
    fmt.Println(v7)
    v2[uint32(v3)] = ""
    for fi81 := 0; fi81 < 19; fi81++ {
        switch "" {
        case "hello":
            v7 -= v7
        }
        var v82 uint32 = 2147483648
        v1 = v1
        _ = fi81
        _ = v82
    }
    var v83 int16 = -204
    for fi84 := 0; fi84 < 7; fi84++ {
        v85 := GenS9{F0: (v3 + (v3 >> (((v3 * (v3 | v3))) & 63))), F1: (v8 + v8)}
        var v86 int8 = 120
        v7 = (v7 - 1e10)
        _ = fi84
        _ = v85
        _ = v86
    }
    var v87 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v83:", v83)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v83
    _ = v87
    _ = v9
}
