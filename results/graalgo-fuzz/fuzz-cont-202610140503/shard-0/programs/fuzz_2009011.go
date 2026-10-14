package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int8; F1 int32; F2 string }
type GenS4 struct { F0 uint8; F1 string; F2 int8 }
type GenS5 struct { F0 uint64; F1 int; F2 int16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 bool; F1 int8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int8; F1 uint8; F2 int32 }
type GenS11 struct { F0 int32; F1 uint64; F2 int16 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint; F1 int32; F2 bool }
type GenS14 struct { F0 uint16; F1 int64 }
func genFn0(p0 int64, p1 int64, p2 uint8) int32 {
    p0 = p0
    fmt.Println(p0, p1, p0)
    for fi5 := 0; fi5 < 13; fi5++ {
        fmt.Println(p1, p0)
        var v6 int32 = 2147483647
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 uint32 = 2147483648
            _ = fi7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    v9 := GenS1{F0: 42, F1: (!true)}
    var v10 uint8 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 2147483647
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 18446744073709551615
    v3 := []int{}
    var v4 bool = true
    _ = len(v1)
    fmt.Println(v4)
    var v11 int = 4294967295
    v4 = v4
    v2 += 142
    v12 := v1[(("hello" + "") + "foo bar")]
    for fi13 := 0; fi13 < 13; fi13++ {
        if true {
            v12 = "hello"
        }
        if (((v4 && (3.14159 > 0.0)) || v4) && (32767 < 0)) {
            var v14 string = "a"
            v15 := GenS2{F0: (^(v2)), F1: 656, F2: 128}
            switch fi13 {
            case 128:
                for fi16 := 0; fi16 < 16; fi16++ {
                    v1["a"] = v14
                    var v17 int = (fi13 ^ fi13)
                    v18 := GenS1{F0: 623, F1: (!v4)}
                    v19 := GenS1{F0: 0, F1: (!false)}
                    switch fi13 {
                    case 127:
                        v20 := GenS2{F0: (v2 / ((850) | 1)), F1: 4294967295, F2: 187}
                        if false {
                            if false {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            v12 = v12
                            v3 = append(v3, -2147483648)
                            var v21 int8 = -67
                            _ = v21
                        }
                        _ = v20
                    case 42:
                        _ = len(v1)
                    case -817:
                        var v22 uint64 = 204
                        _ = v22
                    }
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                }
            }
            _ = v14
            _ = v15
        }
        for fi23 := 0; fi23 < 10; fi23++ {
            v24 := GenS3{F0: 44, F1: 2147483647, F2: v12}
            var v25 float64 = -2.25
            if (0.1 != 3.14159) {
                var v26 int = fi13
                var v27 int32 = 1
                switch v12 {
                case "foo bar":
                    fmt.Println(v11, fi23, v27)
                case "a":
                    _ = len(v1)
                case "hello":
                    var v28 string = ((v12 + v12) + v12)
                    _ = v28
                default:
                    var v29 int16 = -551
                    fmt.Println(fi23)
                    _ = v29
                }
                var v30 uint16 = 32767
                _ = v26
                _ = v27
                _ = v30
            }
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi13
    }
    if v4 {
        v31 := GenS1{F0: 2147483647, F1: (v4 || true)}
        if v4 {
            if v4 {
                v11 += (v11 | -713)
                v2 += (v2 | (v2 * 65535))
            } else {
                if (42 >= 32767) {
                    if false {
                        v1[v12] = "foo bar"
                    }
                    var v32 float32 = float32((-2.25 * (1.5 + 0.1)))
                    if v4 {
                        for fi33 := 0; fi33 < 15; fi33++ {
                            if v4 {
                                v34 := GenS4{F0: 42, F1: v12, F2: -83}
                                v11 = v11
                                _ = v34
                            }
                            _ = fi33
                        }
                    } else {
                        switch v12 {
                        case "foo bar":
                            fmt.Println(v4, v2)
                        case "":
                            v12 = v12
                        case "\n\t":
                            v32 += v32
                        }
                    }
                    _ = v32
                }
                if ((0.0 == 0.0) && (v11 != (-(v11)))) {
                    delete(v1, v12)
                    if (!(0.0 > 1.5)) {
                        var v35 int16 = 820
                        if (3.14159 == -2.25) {
                            var v36 uint = 42
                            _ = len(v3)
                            v37 := GenS5{F0: 42, F1: v11, F2: (-(v35))}
                            _ = v36
                            _ = v37
                        } else {
                            if (false || true) {
                                if v4 {
                                    v38 := GenS1{F0: 2, F1: ((v2 << (((v2 % ((v2) | 1))) & 63)) == (^(v2)))}
                                    if ((v11 * v11) != -267) {
                                        var v39 int32 = 931
                                        if (738 == (v35 | 0)) {
                                            var v40 uint8 = 0
                                            _ = v40
                                        }
                                        _ = v39
                                    }
                                    _ = v38
                                } else {
                                    v12 = "hello"
                                }
                                v2 -= v2
                            }
                            v35 = (v35 | 301)
                        }
                        v41 := GenS6{F0: (v11 / (((-(v11))) | 1))}
                        _ = v35
                        _ = v41
                    }
                    v42 := GenS2{F0: (v2 | (v2 % ((v2) | 1))), F1: 2, F2: 229}
                    _ = v42
                }
            }
            if (665 > 1) {
                v11 += (v11 << ((v2) & 63))
            }
            fmt.Println(v12, v2, v12)
        } else {
            var v43 uint16 = 65535
            _ = v43
        }
        v12 = (((v12 + v12) + ((v12 + v12) + v12)) + "\n\t")
        _ = v31
    } else {
        v2 += v2
        var v44 uint64 = 42
        switch v12 {
        case "hello":
            if v4 {
                if (61 > 1) {
                    var v45 uint64 = v44
                    switch (v12 + "foo bar") {
                    case "":
                        if (v4 || v4) {
                            for fi46 := 0; fi46 < 7; fi46++ {
                                v1[""] = (v12 + "hello")
                                v47 := GenS7{F0: -274}
                                var v48 int8 = 0
                                _ = fi46
                                _ = v47
                                _ = v48
                            }
                            for fi49 := 0; fi49 < 10; fi49++ {
                                if ((uint32(v44) < uint32(v44)) || false) {
                                    if (!(!(-116 > v11))) {
                                        fmt.Println(v44, v12, v4)
                                    }
                                    var v50 uint = 1
                                    _ = v50
                                } else {
                                    for fi51 := 0; fi51 < 13; fi51++ {
                                        v52 := GenS2{F0: v2, F1: (v45 - (v44 / (((v45 + (v44 | (^(v44))))) | 1))), F2: 0}
                                        var v53 int32 = 2147483647
                                        if (!false) {
                                            fmt.Println(v53, fi49, fi49)
                                            v54 := GenS8{F0: v4, F1: 53}
                                            _ = v54
                                        } else {
                                            var v55 uint64 = (v44 | (v44 / (((v45 % ((18446744073709551615) | 1))) | 1)))
                                            _ = v55
                                        }
                                        _ = fi51
                                        _ = v52
                                        _ = v53
                                    }
                                    var v56 uint64 = (v45 << ((18446744073709551615) & 63))
                                    var v57 int8 = 100
                                    _ = v56
                                    _ = v57
                                }
                                v2 = 2
                                switch 421 {
                                case 941:
                                    if false {
                                        _ = len(v3)
                                        var v58 int16 = 42
                                        switch -1 {
                                        case 32768:
                                            if (2147483647 < 173) {
                                                switch "\n\t" {
                                                case "hello":
                                                    var v59 int8 = -128
                                                    _ = v59
                                                case "a":
                                                    if (1e10 != 0.1) {
                                                        v60 := GenS9{F0: (^(v2))}
                                                        _ = v60
                                                    } else {
                                                        v61 := GenS4{F0: 0, F1: ("\n\t" + v12), F2: -29}
                                                        _ = v61
                                                    }
                                                case "\n\t":
                                                    var v62 int = -776
                                                    _ = v62
                                                }
                                                fmt.Println(v11, v11)
                                            } else {
                                                var v63 uint16 = 191
                                                var v64 int8 = -128
                                                var v65 uint8 = 112
                                                _ = v63
                                                _ = v64
                                                _ = v65
                                            }
                                        default:
                                            v45 = v45
                                        }
                                        _ = v58
                                    } else {
                                        var v66 int64 = 9223372036854775807
                                        _ = v66
                                    }
                                }
                                _ = fi49
                            }
                            if len(v3) > 0 {
                                v3[0] = v11
                            }
                        } else {
                            for fi67 := 0; fi67 < 8; fi67++ {
                                if (!true) {
                                    var v68 int8 = 127
                                    var v69 int = 127
                                    _ = v68
                                    _ = v69
                                }
                                v12 = "a"
                                var v70 uint16 = 255
                                _ = fi67
                                _ = v70
                            }
                            v44 -= (^(v44))
                            var v71 uint32 = 2147483647
                            _ = v71
                        }
                    }
                    switch ("" + "hello") {
                    case "a":
                        if len(v3) > 0 {
                            v3[0] = v11
                        }
                    case "":
                        if v4 {
                            fmt.Println(v44)
                            v72 := GenS10{F0: 0, F1: 1, F2: 1}
                            _ = v72
                        }
                    case "foo bar":
                        v73 := GenS11{F0: 218, F1: v45, F2: 42}
                        _ = v73
                    default:
                        v2 = v2
                    }
                    _ = v45
                }
            }
        case "a":
            if v4 {
                v12 = "\n\t"
                var v74 uint = 403
                for fi75 := 0; fi75 < 5; fi75++ {
                    if v4 {
                        var v76 uint32 = 685
                        v76 += 285
                        var v77 int64 = 1
                        _ = v76
                        _ = v77
                    }
                    _ = fi75
                }
                _ = v74
            }
        }
        _ = v44
    }
    v4 = (65535 == 32767)
    v4 = v4
    v2 = 4294967295
    var v78 int16 = -1
    v3 = append(v3, (v11 * v11))
    _ = len(v3)
    if (-128 != -128) {
        v3 = append(v3, 32767)
        if v4 {
            v3 = append(v3, v11)
            fmt.Println(v78)
        } else {
            v79 := GenS4{F0: 206, F1: ("foo bar" + (v12 + v12)), F2: 127}
            v80 := GenS12{F0: 0}
            _ = v79
            _ = v80
        }
        v81 := GenS3{F0: 127, F1: 152, F2: v12}
        _ = v81
    }
    for fi82 := 0; fi82 < 20; fi82++ {
        switch v12 {
        case "a":
            if v4 {
                for fi83 := 0; fi83 < 18; fi83++ {
                    switch "a" {
                    case "":
                        v3 = append(v3, -878)
                    case "a":
                        v78 = (v78 / (((v78 << (((v2 * (v2 & v2))) & 63))) | 1))
                        v84 := GenS13{F0: 2, F1: -846, F2: (100 <= 44)}
                        if (!true) {
                            v85 := GenS14{F0: 65535, F1: 42}
                            _ = v85
                        }
                        _ = v84
                    case "hello":
                        switch "foo bar" {
                        case "a":
                            v78 += v78
                        case "":
                            _ = len(v1)
                        case "foo bar":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    default:
                        fmt.Println(v11, v12)
                    }
                    var v86 uint64 = 4294967295
                    fmt.Println(fi83, v78)
                    fmt.Println(v86, v11, fi82)
                    v87 := GenS11{F0: genFn0(9223372036854775807, -1, 93), F1: v86, F2: v78}
                    _ = fi83
                    _ = v86
                    _ = v87
                }
            }
        }
        _ = fi82
    }
    var v88 float64 = ((1.5 - 1e10) - (1e10 * (3.14159 - (3.14159 + -2.25))))
    v4 = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v78:", v78)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v78
    _ = v88
}
