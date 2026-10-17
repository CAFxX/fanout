package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint; F2 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint64; F1 int64 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint8; F1 uint16; F2 int32 }

func main() {
    var v1 uint16 = 655
    var v2 bool = true
    v3 := []string{"a", ("" + "hello")}
    var v4 string = "\n\t"
    var v5 bool = (-684 != -1)
    v6 := GenS1{F0: 0, F1: 29, F2: ((("a" + "a") + ("a" + v4)) + ("\n\t" + (v4 + ((v4 + "a") + "a"))))}
    v1 = v1
    switch v4 {
    case "foo bar":
        var v7 int32 = -1
        _ = v7
    case "":
        v1 += (v1 % ((v1) | 1))
    case "hello":
        if v5 {
            v4 = (("a" + "foo bar") + "")
            if (!v2) {
                if (uint32(v1) > 65535) {
                    var v8 int = -1
                    _ = v8
                }
                v4 = ("\n\t" + v4)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                var v9 float32 = (1e10 - -2.25)
                v5 = ((153 != 18446744073709551615) || v5)
                _ = v9
            }
        } else {
            if v5 {
                for fi10 := 0; fi10 < 1; fi10++ {
                    for fi11 := 0; fi11 < 20; fi11++ {
                        v12 := GenS1{F0: 0, F1: 2, F2: "a"}
                        var v13 uint = 255
                        fmt.Println(fi10)
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    _ = fi10
                }
            }
        }
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 string = v4
        _ = fi14
        _ = v15
    }
    if len(v3) > 0 {
        v3[0] = "a"
    }
    v3 = append(v3, ("hello" + ""))
    for fi16 := 0; fi16 < 17; fi16++ {
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 int8 = -115
            switch v4 {
            case "foo bar":
                if v5 {
                    if len(v3) > 0 {
                        v3[0] = ("foo bar" + v4)
                    }
                    var v19 uint64 = 569
                    if v2 {
                        if v5 {
                            v3 = append(v3, "foo bar")
                            var v20 bool = (!false)
                            if ((fi16 >> ((476) & 63)) < 854) {
                                var v21 uint8 = 171
                                var v22 uint8 = v21
                                _ = v21
                                _ = v22
                            }
                            _ = v20
                        } else {
                            if (uint8(v1) < 0) {
                                v4 = v4
                                v18 += 1
                                if (true && false) {
                                    if len(v3) > 0 {
                                        v3[0] = v4
                                    }
                                    var v23 bool = (true && (v1 > v1))
                                    if v23 {
                                        switch v4 {
                                        case "foo bar":
                                            var v24 int64 = 9223372036854775807
                                            _ = v24
                                        case "\n\t":
                                            if (int32(fi17) == 436) {
                                                var v25 int64 = 9223372036854775807
                                                var v26 int16 = -843
                                                _ = v25
                                                _ = v26
                                            } else {
                                                var v27 string = v4
                                                _ = v27
                                            }
                                        default:
                                            v28 := GenS2{F0: v1}
                                            _ = v28
                                        }
                                        var v29 float64 = ((1.5 / (-2.25 / 1.0)) - float64(0.0))
                                        _ = v29
                                    }
                                    var v30 uint32 = 2147483647
                                    v31 := GenS2{F0: 1}
                                    v30 -= 810
                                    _ = v23
                                    _ = v30
                                    _ = v31
                                }
                            } else {
                                v18 -= (v18 & 100)
                                var v32 uint = 692
                                v2 = (!v2)
                                switch v4 {
                                case "\n\t":
                                    switch 9223372036854775807 {
                                    case 32768:
                                        if false {
                                            if true {
                                                var v33 int32 = 0
                                                v34 := GenS2{F0: (v1 % ((42) | 1))}
                                                _ = v33
                                                _ = v34
                                            } else {
                                                v19 += v19
                                            }
                                            var v35 uint32 = 42
                                            if len(v3) > 0 {
                                                v3[0] = ("a" + v4)
                                            }
                                            _ = v35
                                        }
                                    case -2147483648:
                                        switch 128 {
                                        case 0:
                                            _ = len(v3)
                                        }
                                    }
                                case "foo bar":
                                    if v2 {
                                        switch (fi17 ^ fi17) {
                                        case 843:
                                            var v36 uint16 = v1
                                            _ = v36
                                        case -9223372036854775808:
                                            v37 := GenS3{F0: 4294967295, F1: (v32 - (^(v32))), F2: 42}
                                            _ = v37
                                        }
                                    }
                                case "a":
                                    if v2 {
                                        v38 := GenS3{F0: v19, F1: (v32 % ((v32) | 1)), F2: -32768}
                                        fmt.Println(v5, v18)
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v39 := GenS4{F0: 42}
                                        _ = v38
                                        _ = v39
                                    }
                                }
                                _ = v32
                            }
                        }
                        switch fi16 {
                        case 32767:
                            v4 = "a"
                        case 1:
                            switch v4 {
                            case "hello":
                                if (((-2.25 / ((1.5 / 1e10) - 1e10)) / 1.0) <= 1e10) {
                                    if len(v3) > 0 {
                                        v3[0] = ((v4 + v4) + "a")
                                    }
                                    var v40 int16 = 32767
                                    var v41 float64 = 0.0
                                    _ = v40
                                    _ = v41
                                } else {
                                    v1 = uint16(v19)
                                    v42 := GenS4{F0: 2}
                                    fmt.Println(v19, fi17)
                                    _ = v42
                                }
                            case "foo bar":
                                switch (fi17 + 1000) {
                                case -9223372036854775808:
                                    v5 = (v2 && v2)
                                }
                            }
                        case 128:
                            var v43 float64 = ((0.1 + 1e10) + 1.5)
                            v3 = append(v3, v4)
                            _ = v43
                        default:
                            fmt.Println(v19)
                        }
                        v5 = v5
                    }
                    if (((v2 && v2) && v5) || v5) {
                        var v44 uint16 = (^(v1))
                        if true {
                            switch v4 {
                            case "":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case "hello":
                                var v45 string = v4
                                _ = v45
                            }
                        } else {
                            switch v4 {
                            case "foo bar":
                                v46 := GenS5{F0: 1, F1: 42}
                                _ = v46
                            default:
                                if v5 {
                                    v18 = (v18 & (v18 / (((v18 % (((^(v18))) | 1))) | 1)))
                                    var v47 float64 = (float64(0.1) - -2.25)
                                    v48 := GenS1{F0: 4294967295, F1: 216, F2: v4}
                                    _ = v47
                                    _ = v48
                                }
                            }
                            if (255 >= 42) {
                                if v5 {
                                    var v49 bool = (true || v2)
                                    _ = v49
                                }
                                if true {
                                    if (255 != 4294967295) {
                                        fmt.Println(v1)
                                    }
                                    var v50 uint32 = 946
                                    _ = v50
                                } else {
                                    var v51 bool = ((v18 | (v18 ^ (v18 - v18))) <= (^(v18)))
                                    _ = v51
                                }
                                var v52 uint8 = 1
                                v53 := GenS5{F0: 252, F1: 9223372036854775807}
                                _ = v52
                                _ = v53
                            } else {
                                switch fi16 {
                                case -2147483648:
                                    if len(v3) > 0 {
                                        v3[0] = "hello"
                                    }
                                case 0:
                                    v19 += v19
                                default:
                                    v3 = append(v3, "foo bar")
                                }
                                if true {
                                    var v54 uint64 = v19
                                    _ = v54
                                } else {
                                    v44 -= v1
                                    v3 = append(v3, "hello")
                                }
                            }
                        }
                        _ = v44
                    }
                    _ = v19
                }
            case "a":
                if (v5 && (v18 <= 38)) {
                    if (0 == 2147483647) {
                        v4 = "hello"
                    } else {
                        v4 = ""
                        var v55 int = (^(fi16))
                        _ = v55
                    }
                    v56 := GenS4{F0: 2}
                    v18 = (v18 ^ (v18 ^ (v18 << ((255) & 63))))
                    _ = v56
                } else {
                    var v57 uint32 = 2147483647
                    if ((0 > -2147483648) || v5) {
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                    } else {
                        if (v1 >= v1) {
                            if v2 {
                                _ = len(v3)
                            } else {
                                v5 = (!(127 < 42))
                                if (-1 == 329) {
                                    var v58 uint16 = v1
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v59 := GenS6{F0: -820}
                                    _ = v58
                                    _ = v59
                                } else {
                                    var v60 uint64 = 18446744073709551615
                                    _ = v60
                                }
                                v18 = 100
                                v61 := GenS7{F0: 1, F1: -1}
                                _ = v61
                            }
                            v62 := GenS2{F0: (v1 ^ 65535)}
                            var v63 float32 = 1.5
                            _ = v62
                            _ = v63
                        } else {
                            v4 = v4
                            _ = len(v3)
                            if true {
                                var v64 bool = (!(42 == 1))
                                var v65 uint8 = 128
                                var v66 string = "\n\t"
                                switch fi16 {
                                case 593:
                                    var v67 int8 = (v18 & (v18 ^ (v18 + v18)))
                                    _ = v67
                                }
                                v2 = v64
                                fmt.Println(v2, v1)
                                _ = v64
                                _ = v65
                                _ = v66
                            }
                        }
                        v68 := GenS8{F0: -926}
                        var v69 uint8 = 0
                        var v70 uint16 = v1
                        _ = v68
                        _ = v69
                        _ = v70
                    }
                    switch v4 {
                    case "":
                        var v71 uint8 = 229
                        _ = v71
                    case "foo bar":
                        switch v4 {
                        case "foo bar":
                            var v72 int8 = v18
                            _ = v72
                        }
                    case "hello":
                        v4 = ("" + "foo bar")
                    }
                    if (true || (255 > 18446744073709551615)) {
                        if (v57 >= (v57 >> ((18446744073709551615) & 63))) {
                            switch v4 {
                            case "hello":
                                switch 127 {
                                case 1:
                                    _ = len(v3)
                                case -9223372036854775808:
                                    switch v4 {
                                    case "foo bar":
                                        if ((v4 + "foo bar") != (("\n\t" + v4) + "")) {
                                            _ = len(v3)
                                            var v73 int8 = (v18 / (((v18 + (v18 * (^(v18))))) | 1))
                                            v3 = append(v3, "hello")
                                            _ = v73
                                        }
                                    }
                                case -35:
                                    if false {
                                        if ("" >= v4) {
                                            fmt.Println(v2, v18, v1)
                                        } else {
                                            v1 += 163
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            switch "" {
                                            case "foo bar":
                                                switch "hello" {
                                                case "":
                                                    switch v4 {
                                                    case "foo bar":
                                                        v1 += 0
                                                    case "a":
                                                        if v2 {
                                                            var v74 uint16 = v1
                                                            fmt.Println(fi17, v18)
                                                            _ = v74
                                                        }
                                                    case "hello":
                                                        var v75 uint16 = v1
                                                        _ = v75
                                                    default:
                                                        v4 = (v4 + "a")
                                                    }
                                                case "hello":
                                                    var v76 float32 = 0.0
                                                    _ = v76
                                                default:
                                                    _ = len(v3)
                                                }
                                                var v77 string = (("\n\t" + "\n\t") + ((v4 + (v4 + "hello")) + ""))
                                                v77 = v77
                                                _ = v77
                                            default:
                                                if (0.0 <= (-2.25 + ((-2.25 + 1e10) / -2.25))) {
                                                    fmt.Println(fi17, v18, fi17)
                                                } else {
                                                    v78 := GenS3{F0: 18446744073709551615, F1: 4294967295, F2: 19}
                                                    _ = v78
                                                }
                                            }
                                        }
                                        v4 = ("foo bar" + v4)
                                    } else {
                                        v1 -= (v1 % ((v1) | 1))
                                        var v79 int64 = -1
                                        _ = v79
                                    }
                                default:
                                    v80 := GenS2{F0: uint16(v57)}
                                    _ = v80
                                }
                            default:
                                var v81 uint16 = (v1 & 567)
                                _ = v81
                            }
                            fmt.Println(fi17)
                        } else {
                            var v82 int = int(v18)
                            switch 32767 {
                            case -83:
                                switch -2147483648 {
                                case 9223372036854775807:
                                    v57 = (v57 * v57)
                                case -1:
                                    v83 := GenS9{F0: 169, F1: (v1 % (((^(v1))) | 1)), F2: -1}
                                    _ = v83
                                default:
                                    v5 = v5
                                }
                            default:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            var v84 uint64 = 928
                            _ = v82
                            _ = v84
                        }
                    } else {
                        var v85 uint16 = v1
                        v3 = append(v3, v4)
                        _ = v85
                    }
                    _ = v57
                }
                switch (v4 + "foo bar") {
                case "foo bar":
                    switch -1 {
                    case 9223372036854775807:
                        var v86 float32 = 1e10
                        _ = v86
                    }
                case "\n\t":
                    v87 := GenS5{F0: 32767, F1: -9223372036854775808}
                    _ = v87
                default:
                    var v88 float32 = (0.1 / float32(((1.5 + 3.14159) / ((3.14159 + (0.0 + -2.25)) / 0.1))))
                    _ = v88
                }
            }
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    v1 = (v1 & (v1 >> ((865) & 63)))
    v89 := GenS7{F0: 42, F1: 1}
    var v90 uint8 = 74
    var v91 int = 0
    if ((v1 + v1) == (v1 * v1)) {
        for fi92 := 0; fi92 < 19; fi92++ {
            if (2 < 255) {
                var v93 int16 = 32767
                v94 := GenS5{F0: v93, F1: 9223372036854775807}
                _ = v93
                _ = v94
            }
            _ = fi92
        }
        v2 = (!v5)
    } else {
        var v95 int32 = -183
        _ = v95
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v89
    _ = v90
    _ = v91
}
