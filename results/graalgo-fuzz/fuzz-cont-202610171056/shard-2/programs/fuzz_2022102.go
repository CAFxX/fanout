package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 uint64; F2 int16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string; F1 uint64; F2 int64 }
type GenS6 struct { F0 int64; F1 string }

func main() {
    var v1 int32 = 0
    var v2 uint16 = 704
    var v3 int64 = 1
    v4 := []uint32{0, 483}
    v5 := make(map[string]string)
    v6 := []uint32{}
    var v7 uint64 = 42
    v5[(("a" + ("hello" + "a")) + "\n\t")] = ""
    var v8 float64 = 1e10
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v6) > 0 {
        v6[0] = 886
    }
    v9 := v5["a"]
    delete(v5, "hello")
    var v10 uint = 65535
    if true {
        v7 += (^(v7))
        for fi11 := 0; fi11 < 18; fi11++ {
            var v12 uint32 = 0
            for fi13 := 0; fi13 < 2; fi13++ {
                v14 := GenS1{F0: 32767}
                _ = fi13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
        var v15 float64 = 0.1
        _ = v15
    }
    switch 9223372036854775807 {
    case 4294967295:
        var v16 bool = false
        _ = v16
    case 128:
        var v17 int64 = 9223372036854775807
        _ = v17
    case -1:
        v18 := v5[("foo bar" + (v9 + "hello"))]
        _ = v18
    }
    if false {
        if (false || true) {
            for fi19 := 0; fi19 < 4; fi19++ {
                v8 = 1e10
                switch fi19 {
                case 4294967295:
                    var v20 float64 = (((v8 / v8) - (1e10 * 0.1)) * v8)
                    _ = v20
                case -9223372036854775808:
                    var v21 uint8 = 128
                    _ = v21
                }
                _ = fi19
            }
        } else {
            v10 = (v10 * (^(v10)))
            var v22 float32 = 0.1
            _ = v22
        }
        for fi23 := 0; fi23 < 13; fi23++ {
            if (127 >= 216) {
                if true {
                    var v24 float64 = ((v8 * 0.0) + v8)
                    if true {
                        var v25 int32 = -1
                        var v26 int8 = -96
                        if (!(!(v24 == v24))) {
                            for fi27 := 0; fi27 < 18; fi27++ {
                                switch (fi27 ^ (-(fi27))) {
                                case 678:
                                    if (42 <= 2) {
                                        v7 += (v7 - 4294967295)
                                        var v28 uint16 = v2
                                        _ = v28
                                    } else {
                                        v24 = (1.5 + ((v24 + v24) * 3.14159))
                                        delete(v5, "a")
                                    }
                                }
                                _ = fi27
                            }
                        } else {
                            var v29 int64 = (v3 & (v3 | v3))
                            _ = v29
                        }
                        _ = v25
                        _ = v26
                    }
                    _ = v24
                } else {
                    var v30 string = "foo bar"
                    var v31 string = ((("foo bar" + v30) + v30) + (v9 + (v30 + "hello")))
                    _ = v30
                    _ = v31
                }
            } else {
                if (-177 >= (v3 & (v3 ^ -1))) {
                    fmt.Println(v2, v8)
                    var v32 string = v9
                    v1 += v1
                    _ = v32
                } else {
                    v7 = v7
                    if true {
                        var v33 int = -9223372036854775808
                        for fi34 := 0; fi34 < 3; fi34++ {
                            var v35 uint64 = 18446744073709551615
                            if false {
                                v36 := GenS1{F0: int(v2)}
                                _ = v36
                            } else {
                                fmt.Println(v8)
                                var v37 int8 = int8(v2)
                                _ = v37
                            }
                            var v38 int16 = 32767
                            _ = fi34
                            _ = v35
                            _ = v38
                        }
                        _ = v33
                    }
                }
                var v39 int16 = -336
                v1 = v1
                _ = v39
            }
            _ = fi23
        }
        v8 += 1.5
    }
    var v40 float64 = 1.5
    v3 = 9223372036854775807
    v4 = append(v4, 42)
    for fi41 := 0; fi41 < 17; fi41++ {
        v1 = 703
        if (v8 < v40) {
            v40 -= v8
            v4 = append(v4, uint32(v10))
        } else {
            if ((!(((v9 <= (("a" + v9) + v9)) || ((v10 >> (((v10 | v10)) & 63)) == (v10 * v10))) && false)) || (1 > 2)) {
                for fi42 := 0; fi42 < 11; fi42++ {
                    v43 := GenS1{F0: (-(fi42))}
                    _ = fi42
                    _ = v43
                }
            } else {
                if len(v4) > 0 {
                    v4[0] = 1
                }
                if (!((v2 / ((1) | 1)) > (v2 | v2))) {
                    fmt.Println(v8)
                } else {
                    v3 += v3
                    var v44 uint8 = uint8(v10)
                    var v45 string = ("" + "")
                    _ = v44
                    _ = v45
                }
            }
        }
        _ = fi41
    }
    switch ("a" + "foo bar") {
    case "foo bar":
        v10 -= v10
    case "\n\t":
        fmt.Println(v1)
    default:
        switch v9 {
        case "\n\t":
            var v46 int16 = int16(v2)
            _ = v46
        case "":
            v47 := GenS1{F0: int(v2)}
            _ = v47
        }
    }
    for fi48 := 0; fi48 < 6; fi48++ {
        for fi49 := 0; fi49 < 10; fi49++ {
            if (-12 != 99) {
                if false {
                    fmt.Println(v3)
                    if (((49 != 255) || false) || true) {
                        if false {
                            if (!((v10 & (v10 - uint(v2))) <= (v10 - 286))) {
                                switch (fi49 + 128) {
                                case 4294967295:
                                    v50 := GenS2{F0: (v9 + "")}
                                    _ = v50
                                case 2147483647:
                                    switch "\n\t" {
                                    case "":
                                        switch v9 {
                                        case "\n\t":
                                            if false {
                                                v8 = -2.25
                                                fmt.Println(v10, v1, v9)
                                                if false {
                                                    switch ((((("hello" + v9) + "\n\t") + v9) + "foo bar") + "hello") {
                                                    case "a":
                                                        v51 := GenS2{F0: ("" + "\n\t")}
                                                        _ = v51
                                                    case "":
                                                        v1 += (-(v1))
                                                        fmt.Println(fi48)
                                                    case "foo bar":
                                                        fmt.Println(v2, v7)
                                                    }
                                                }
                                                var v52 uint16 = 325
                                                v3 -= -9223372036854775808
                                                _ = v52
                                            }
                                        case "a":
                                            v5[v9] = v9
                                        }
                                    case "a":
                                        _ = len(v4)
                                    default:
                                        var v53 int16 = 819
                                        _ = v53
                                    }
                                }
                                switch fi48 {
                                case -2147483648:
                                    fmt.Println(v40, v10)
                                }
                                fmt.Println(v10)
                                var v54 int8 = 1
                                _ = v54
                            } else {
                                var v55 int64 = v3
                                v1 += -1
                                var v56 uint = v10
                                _ = v55
                                _ = v56
                            }
                            fmt.Println(v7, fi49)
                        }
                        switch fi48 {
                        case -1:
                            if (false && false) {
                                fmt.Println(v8)
                                if false {
                                    v2 = v2
                                    var v57 float32 = 0.0
                                    if (v8 < ((-2.25 + v40) * (v40 + v40))) {
                                        var v58 float64 = (1e10 - 1e10)
                                        v8 -= v58
                                        _ = len(v5)
                                        _ = v58
                                    }
                                    v59 := GenS1{F0: fi49}
                                    v8 -= 1e10
                                    switch (fi48 | (fi49 / (((^(fi48))) | 1))) {
                                    case -642:
                                        var v60 int8 = 0
                                        _ = v60
                                    case -925:
                                        switch -711 {
                                        case 451:
                                            switch v9 {
                                            case "":
                                                var v61 uint = 18446744073709551615
                                                _ = v61
                                            case "hello":
                                                switch fi49 {
                                                case 32767:
                                                    switch 2147483647 {
                                                    case -717:
                                                        var v62 string = (((((v9 + v9) + v9) + v9) + ("" + v9)) + ("\n\t" + v9))
                                                        _ = v62
                                                    case 457:
                                                        switch (v9 + v9) {
                                                        case "":
                                                            v63 := v5[v9]
                                                            _ = v63
                                                        case "a":
                                                            var v64 int32 = v1
                                                            _ = v64
                                                        case "foo bar":
                                                            v65 := GenS2{F0: "a"}
                                                            _ = v65
                                                        }
                                                    default:
                                                        if len(v6) > 0 {
                                                            v6[0] = uint32(v10)
                                                        }
                                                    }
                                                case 23:
                                                    fmt.Println(v10)
                                                default:
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                }
                                            case "foo bar":
                                                v66 := GenS2{F0: ((v9 + v9) + "foo bar")}
                                                _ = v66
                                            }
                                            var v67 int = (fi49 + (fi49 | fi49))
                                            v7 -= (v7 + (v7 % ((0) | 1)))
                                            _ = v67
                                        case 142:
                                            v9 = (v9 + (v9 + v9))
                                        }
                                    case -1:
                                        switch v9 {
                                        case "":
                                            var v68 int64 = v3
                                            _ = v68
                                        case "foo bar":
                                            v69 := GenS2{F0: ((v9 + v9) + ("foo bar" + ("foo bar" + v9)))}
                                            _ = v69
                                        case "a":
                                            if false {
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                v70 := GenS3{F0: 1, F1: (^(v7)), F2: 42}
                                                _ = v70
                                            } else {
                                                switch 2147483647 {
                                                case 1:
                                                    v71 := GenS4{F0: (-282 <= 0)}
                                                    v5["\n\t"] = "hello"
                                                    var v72 uint8 = 127
                                                    _ = v71
                                                    _ = v72
                                                case 492:
                                                    if (false && false) {
                                                        v40 -= float64(3.14159)
                                                        v57 = float32((v8 - ((float64(((3.14159 - ((1e10 + (v57 - float32(((v8 - v8) * v40)))) - (((v57 - 0.0) * float32((v8 / v40))) * (v57 / float32((v40 - float64(((((v57 / v57) - (v57 / v57)) - (1.5 / (v57 + 3.14159))) * ((((3.14159 - ((float32(float64(0.1)) * float32(v8)) * float32((0.1 / v8)))) * ((1.5 + v57) * (0.0 * v57))) + (v57 - 0.0)) / (3.14159 / (v57 - v57))))))))))) - v57)) - 1e10) / (1.5 / ((0.1 - v40) * v40)))))
                                                        if false {
                                                            var v73 bool = true
                                                            _ = len(v5)
                                                            _ = v73
                                                        }
                                                        var v74 uint32 = 2147483648
                                                        v75 := GenS5{F0: v9, F1: v7, F2: (v3 >> ((v10) & 63))}
                                                        v5[v9] = v9
                                                        _ = v74
                                                        _ = v75
                                                    }
                                                case 541:
                                                    v8 += v8
                                                default:
                                                    var v76 uint16 = uint16(v10)
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    _ = v76
                                                }
                                                var v77 uint32 = 42
                                                var v78 int8 = 127
                                                v79 := GenS6{F0: (-(v3)), F1: "hello"}
                                                v40 = v40
                                                v80 := GenS2{F0: "foo bar"}
                                                _ = v77
                                                _ = v78
                                                _ = v79
                                                _ = v80
                                            }
                                        default:
                                            var v81 float64 = v40
                                            v82 := GenS1{F0: fi48}
                                            _ = v81
                                            _ = v82
                                        }
                                    }
                                    _ = v57
                                    _ = v59
                                } else {
                                    v40 += v40
                                }
                            } else {
                                var v83 int = fi49
                                _ = v83
                            }
                        default:
                            if false {
                                v84 := GenS4{F0: true}
                                _ = v84
                            } else {
                                if true {
                                    fmt.Println(v10, v1)
                                } else {
                                    var v85 uint = (v10 | v10)
                                    _ = v85
                                }
                                v3 += (v3 * 1)
                            }
                        }
                    }
                }
                var v86 float64 = 0.0
                _ = v86
            } else {
                v87 := GenS5{F0: "hello", F1: uint64(v10), F2: (v3 | -9223372036854775808)}
                _ = v87
            }
            _ = fi49
        }
        var v88 int32 = (^(v1))
        for fi89 := 0; fi89 < 20; fi89++ {
            if len(v4) > 0 {
                v4[0] = 0
            }
            _ = fi89
        }
        _ = fi48
        _ = v88
    }
    var v90 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    _ = v90
}
