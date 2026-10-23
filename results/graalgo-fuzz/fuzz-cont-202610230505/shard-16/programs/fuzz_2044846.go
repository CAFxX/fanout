package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 string }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 int16; F1 bool; F2 uint8 }
type GenS5 struct { F0 int; F1 int16 }
func genFn0(p0 uint64, p1 int) uint32 {
    p1 -= p1
    var v4 int32 = 872
    p0 -= p0
    for fi5 := 0; fi5 < 11; fi5++ {
        v4 += -2147483648
        p0 -= p0
        fmt.Println(p1)
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = v4
    return 42
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]string)
    v3 := []string{"a", ("hello" + ("hello" + "\n\t")), "hello"}
    if (!((false || false) || true)) {
        if (((-1 >= 860) || (-2147483648 >= 119)) && (((((-2.25 - 0.0) / 0.1) / 1.5) / 3.14159) > ((float32((-2.25 * 0.1)) / 1e10) * float32((-2.25 * 0.1))))) {
            var v6 int = -1
            for fi7 := 0; fi7 < 8; fi7++ {
                v6 = fi7
                v6 += (fi7 / ((v6) | 1))
                var v8 int16 = 32767
                _ = fi7
                _ = v8
            }
            _ = v6
        }
    } else {
        switch 453 {
        case -9223372036854775808:
            v9 := GenS1{F0: "hello", F1: 32768, F2: "foo bar"}
            _ = v9
        }
        var v10 int64 = 9223372036854775807
        _ = v10
    }
    switch "hello" {
    case "a":
        var v11 int64 = -1
        _ = v11
    case "":
        var v12 string = "a"
        _ = v12
    case "foo bar":
        var v13 uint16 = 32767
        _ = v13
    default:
        _ = len(v3)
    }
    var v14 bool = (!false)
    var v15 int16 = 864
    for fi16 := 0; fi16 < 14; fi16++ {
        switch 127 {
        case 118:
            var v17 int64 = 156
            _ = v17
        case 309:
            for fi18 := 0; fi18 < 15; fi18++ {
                if ((((((0.0 * 1.5) / -2.25) * (((1.5 - (1e10 * 1e10)) - ((float32(3.14159) / 3.14159) - 3.14159)) + 1.5)) / 1.0) != 0.1) || false) {
                    switch -9223372036854775808 {
                    case 32768:
                        switch "a" {
                        case "":
                            if true {
                                _ = len(v3)
                                switch fi16 {
                                case 32767:
                                    v3 = append(v3, (("hello" + "foo bar") + ("" + "a")))
                                case 9223372036854775807:
                                    switch "a" {
                                    case "hello":
                                        var v19 string = "\n\t"
                                        _ = v19
                                    default:
                                        var v20 int = fi18
                                        _ = v20
                                    }
                                case 0:
                                    switch "\n\t" {
                                    case "":
                                        v21 := GenS2{F0: 4294967295, F1: fi18}
                                        _ = v21
                                    case "foo bar":
                                        if (-553 > fi18) {
                                            fmt.Println(v15, v15)
                                            var v22 float64 = -2.25
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            _ = v22
                                        } else {
                                            var v23 float64 = 0.1
                                            v23 -= v23
                                            _ = v23
                                        }
                                    case "a":
                                        v15 = v15
                                    }
                                }
                            }
                        case "foo bar":
                            _ = len(v3)
                        default:
                            v24 := GenS1{F0: "foo bar", F1: fi18, F2: "hello"}
                            _ = v24
                        }
                    case 9223372036854775807:
                        if (v14 || (77 < 127)) {
                            v14 = (749 != 65535)
                            fmt.Println(fi18, fi18)
                        } else {
                            var v25 uint16 = 584
                            _ = v25
                        }
                        var v26 int8 = -119
                        var v27 int64 = 42
                        v15 = -925
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v26
                        _ = v27
                    }
                } else {
                    fmt.Println(v14, v15, fi18)
                    v3 = append(v3, "a")
                    v15 += (v15 ^ v15)
                }
                v15 += v15
                switch fi18 {
                case 349:
                    if v14 {
                        v15 = v15
                    } else {
                        var v28 int8 = 76
                        _ = v28
                    }
                default:
                    fmt.Println(fi18, fi16)
                }
                _ = fi18
            }
        case -523:
            for fi29 := 0; fi29 < 20; fi29++ {
                if (false && (fi29 >= (fi16 * (-(fi16))))) {
                    var v30 uint8 = 69
                    if false {
                        if false {
                            v3 = append(v3, (("a" + "hello") + ("" + ("hello" + "foo bar"))))
                            switch "\n\t" {
                            case "a":
                                var v31 int32 = -512
                                _ = v31
                            case "foo bar":
                                v32 := GenS2{F0: 774, F1: (fi29 & (-(fi16)))}
                                _ = v32
                            case "":
                                var v33 uint64 = 2
                                _ = v33
                            }
                            fmt.Println(v14)
                            v3 = append(v3, (("foo bar" + "hello") + "hello"))
                        } else {
                            if (2147483648 == genFn0(42, (fi29 | (-(fi16))))) {
                                switch 4294967295 {
                                case 32767:
                                    fmt.Println(v15, v30)
                                    v34 := GenS1{F0: "hello", F1: fi29, F2: "foo bar"}
                                    if v14 {
                                        v15 = -221
                                        v30 = v30
                                        fmt.Println(fi29)
                                    }
                                    _ = v34
                                case -9223372036854775808:
                                    var v35 int8 = -128
                                    _ = v35
                                case 2147483647:
                                    if v14 {
                                        v15 += -702
                                        var v36 float32 = -2.25
                                        var v37 uint32 = 2147483648
                                        _ = v36
                                        _ = v37
                                    } else {
                                        v38 := GenS2{F0: 18446744073709551615, F1: (fi16 / ((fi16) | 1))}
                                        v15 += 42
                                        v15 = 42
                                        _ = v38
                                    }
                                default:
                                    var v39 int64 = 9223372036854775807
                                    _ = v39
                                }
                                switch ("\n\t" + "a") {
                                case "hello":
                                    delete(v2, ("" + ("foo bar" + ("foo bar" + ""))))
                                case "\n\t":
                                    var v40 float32 = 3.14159
                                    _ = v40
                                case "foo bar":
                                    v41 := GenS3{F0: (v15 + v15), F1: 255}
                                    _ = v41
                                }
                                var v42 float64 = ((-2.25 + float64(0.0)) * 1e10)
                                _ = v42
                            }
                            if v14 {
                                v15 -= v15
                            }
                        }
                        fmt.Println(fi29, fi16, fi16)
                    } else {
                        switch 1 {
                        case 2147483647:
                            var v43 string = "a"
                            _ = v43
                        default:
                            v44 := GenS4{F0: v15, F1: v14, F2: (v30 ^ v30)}
                            _ = v44
                        }
                    }
                    fmt.Println(v14, fi16, v15)
                    _ = v30
                }
                if (42 > 42) {
                    delete(v2, "a")
                } else {
                    if v14 {
                        if (0 <= 255) {
                            v15 = (v15 & v15)
                            fmt.Println(v15)
                        }
                        fmt.Println(v15)
                        if v14 {
                            v45 := GenS2{F0: 18446744073709551615, F1: fi16}
                            delete(v2, ("foo bar" + "\n\t"))
                            if true {
                                switch (((("hello" + "a") + "hello") + "") + ("\n\t" + "a")) {
                                case "hello":
                                    if len(v3) > 0 {
                                        v3[0] = (("a" + "") + "\n\t")
                                    }
                                case "foo bar":
                                    v14 = (!v14)
                                default:
                                    if v14 {
                                        var v46 int16 = v15
                                        _ = v46
                                    }
                                }
                            } else {
                                _ = len(v3)
                                if (42 <= 0) {
                                    v47 := GenS1{F0: ("a" + ("a" + "hello")), F1: (fi29 | fi16), F2: "hello"}
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v48 := GenS5{F0: fi29, F1: (v15 * (v15 * -662))}
                                    _ = v47
                                    _ = v48
                                } else {
                                    var v49 int32 = 0
                                    _ = v49
                                }
                                v50 := GenS4{F0: (v15 << ((281) & 63)), F1: v14, F2: 127}
                                _ = v50
                            }
                            var v51 uint64 = 18446744073709551615
                            v51 = 2
                            _ = v45
                            _ = v51
                        }
                        _ = len(v3)
                        v52 := GenS2{F0: 375, F1: (^(fi16))}
                        var v53 uint16 = 858
                        _ = v52
                        _ = v53
                    } else {
                        switch fi16 {
                        case 2147483647:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case -1:
                            var v54 int32 = -452
                            _ = v54
                        case -32:
                            v15 += (v15 % ((int16(fi16)) | 1))
                        }
                        var v55 uint32 = 2
                        v15 -= v15
                        if (48 < 255) {
                            switch fi16 {
                            case -9223372036854775808:
                                var v56 int64 = -249
                                _ = v56
                            default:
                                if (int32(fi16) < -2147483648) {
                                    var v57 int64 = 1
                                    _ = v57
                                } else {
                                    fmt.Println(v55)
                                }
                            }
                            var v58 float32 = 1e10
                            var v59 uint32 = (v55 >> ((1) & 63))
                            _ = v58
                            _ = v59
                        }
                        var v60 int8 = -1
                        if v14 {
                            if v14 {
                                v3 = append(v3, "foo bar")
                                v55 += (v55 / (((v55 | (v55 | (v55 ^ 1)))) | 1))
                                v15 += (v15 ^ v15)
                            }
                            v61 := GenS4{F0: v15, F1: ((v55 + genFn0(0, 32767)) == 840), F2: 156}
                            v1[(fi16 - (fi16 >> ((683) & 63)))] = ((v15 % ((1000) | 1)) > v15)
                            _ = v61
                        }
                        _ = v55
                        _ = v60
                    }
                }
                fmt.Println(v15, v15)
                _ = fi29
            }
        }
        _ = fi16
    }
    v62 := GenS2{F0: 2, F1: -372}
    for fi63 := 0; fi63 < 12; fi63++ {
        var v64 bool = (42 > 550)
        v65 := GenS1{F0: "", F1: (fi63 ^ (fi63 + -1)), F2: "foo bar"}
        v15 = (v15 % ((v15) | 1))
        _ = fi63
        _ = v64
        _ = v65
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    _ = len(v1)
    v15 += (v15 << ((80) & 63))
    switch 9223372036854775807 {
    case 128:
        for fi66 := 0; fi66 < 11; fi66++ {
            v14 = (556 < 199)
            _ = fi66
        }
    default:
        if (int64(v15) < -9223372036854775808) {
            var v67 int8 = -128
            if (!(42 > 1)) {
                delete(v2, "")
            } else {
                v15 += v15
                v15 = 0
            }
            _ = v67
        }
    }
    if (!(!(float64(float32((1e10 - (0.0 - 0.1)))) <= 0.0))) {
        var v68 uint64 = 692
        var v69 int16 = 0
        var v70 int32 = 637
        _ = v68
        _ = v69
        _ = v70
    }
    var v71 int8 = 1
    var v72 uint32 = 140
    v14 = true
    var v73 int = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v62
    _ = v71
    _ = v72
    _ = v73
}
