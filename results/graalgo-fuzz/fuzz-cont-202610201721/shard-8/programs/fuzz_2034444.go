package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int64 }
type GenS2 struct { F0 string; F1 uint; F2 string }
type GenS3 struct { F0 bool; F1 bool; F2 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 uint64; F1 int16 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int8; F1 uint; F2 int8 }
type GenS11 struct { F0 int16; F1 int64 }
type GenS12 struct { F0 uint8; F1 int }
type GenS13 struct { F0 uint16; F1 bool }
func genFn0(p0 int64, p1 bool) int32 {
    for fi7 := 0; fi7 < 18; fi7++ {
        fmt.Println(p1, p1)
        v8 := GenS1{F0: 205, F1: 0, F2: -9223372036854775808}
        if p1 {
            for fi9 := 0; fi9 < 15; fi9++ {
                var v10 uint16 = 42
                var v11 float64 = (-2.25 - 3.14159)
                _ = fi9
                _ = v10
                _ = v11
            }
            fmt.Println(p1, p0, p1)
            v12 := GenS2{F0: "foo bar", F1: 42, F2: ("a" + "\n\t")}
            _ = v12
        } else {
            v13 := GenS1{F0: 42, F1: 486, F2: (-(p0))}
            switch fi7 {
            case -27:
                v14 := GenS3{F0: (!p1), F1: (p0 <= -1), F2: int(p0)}
                _ = v14
            }
            _ = v13
        }
        _ = fi7
        _ = v8
    }
    if true {
        fmt.Println(p1)
    } else {
        p0 = p0
        v15 := GenS1{F0: 18446744073709551615, F1: 42, F2: 1}
        p1 = (!p1)
        _ = v15
    }
    _ = p0
    _ = p1
    return 849
}

func main() {
    var v1 uint8 = 255
    v2 := []int{}
    var v3 int16 = 0
    v4 := make(map[int]bool)
    var v5 bool = true
    var v6 uint16 = 42
    fmt.Println(v1)
    if ("\n\t" == "\n\t") {
        v16 := GenS1{F0: 382, F1: 18446744073709551615, F2: 9223372036854775807}
        v5 = (42 <= 18446744073709551615)
        var v17 uint8 = v1
        _ = v16
        _ = v17
    }
    if (956 <= 2147483647) {
        v18 := GenS1{F0: 2, F1: 4294967295, F2: 42}
        _ = v18
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        v2 = append(v2, (fi19 << ((18446744073709551615) & 63)))
        switch "hello" {
        case "":
            v2 = append(v2, fi19)
        case "\n\t":
            v1 = v1
        case "foo bar":
            var v20 string = ("\n\t" + (("hello" + "a") + "\n\t"))
            _ = v20
        default:
            for fi21 := 0; fi21 < 10; fi21++ {
                switch ("" + "") {
                case "":
                    v3 -= v3
                case "a":
                    switch ("hello" + (("\n\t" + "") + "hello")) {
                    case "":
                        v22 := GenS3{F0: v5, F1: false, F2: fi19}
                        _ = v22
                    }
                case "hello":
                    if (false && v5) {
                        var v23 int64 = 0
                        _ = v23
                    } else {
                        _ = len(v4)
                        _ = len(v2)
                    }
                }
                _ = fi21
            }
        }
        _ = fi19
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        v25 := GenS3{F0: v5, F1: v5, F2: -1}
        v6 -= v6
        switch "a" {
        case "foo bar":
            switch "foo bar" {
            case "\n\t":
                var v26 uint8 = 2
                _ = v26
            default:
                switch "\n\t" {
                case "hello":
                    var v27 float64 = float64(0.0)
                    _ = v27
                case "":
                    v1 -= (v1 | v1)
                case "a":
                    for fi28 := 0; fi28 < 12; fi28++ {
                        if (!false) {
                            var v29 uint16 = (v6 % ((v6) | 1))
                            v30 := v4[(fi28 * (fi28 & fi28))]
                            if len(v2) > 0 {
                                v2[0] = -1
                            }
                            _ = v29
                            _ = v30
                        }
                        if false {
                            var v31 bool = v5
                            switch "hello" {
                            case "hello":
                                switch -802 {
                                case -2147483648:
                                    v32 := GenS4{F0: v6}
                                    _ = v32
                                case 128:
                                    var v33 int = (fi24 << ((65535) & 63))
                                    _ = v33
                                case 127:
                                    var v34 uint32 = 4294967295
                                    _ = v34
                                }
                            case "foo bar":
                                fmt.Println(fi28, v5, v1)
                            case "a":
                                var v35 int32 = 1
                                _ = v35
                            }
                            if (fi28 < fi28) {
                                var v36 int = (^(fi28))
                                var v37 string = "a"
                                if true {
                                    if (2147483648 != 40) {
                                        if len(v2) > 0 {
                                            v2[0] = v36
                                        }
                                        if (0 == v1) {
                                            v38 := GenS1{F0: 1, F1: 65535, F2: 1}
                                            v37 = "hello"
                                            fmt.Println(fi24, fi24)
                                            _ = v38
                                        } else {
                                            var v39 bool = true
                                            v2 = append(v2, -927)
                                            _ = v39
                                        }
                                        var v40 float64 = (float64(-2.25) + (3.14159 / 1.0))
                                        if v31 {
                                            v41 := GenS5{F0: 0}
                                            fmt.Println(v40)
                                            _ = v41
                                        }
                                        fmt.Println(v40)
                                        v3 -= (v3 / ((v3) | 1))
                                        var v42 int64 = 1
                                        var v43 float64 = 3.14159
                                        var v44 int16 = 42
                                        _ = v40
                                        _ = v42
                                        _ = v43
                                        _ = v44
                                    } else {
                                        v36 = (fi28 | v36)
                                        if v31 {
                                            var v45 uint64 = 0
                                            v31 = v5
                                            var v46 string = ""
                                            _ = v45
                                            _ = v46
                                        }
                                        v47 := GenS5{F0: -714}
                                        _ = v47
                                    }
                                    if (-1 > 0) {
                                        var v48 int32 = 135
                                        v6 = (v6 + (v6 * v6))
                                        if (v31 && ((18446744073709551615 > 18446744073709551615) || (v5 || false))) {
                                            v6 -= v6
                                            v49 := GenS6{F0: v1, F1: "foo bar"}
                                            v36 += -9223372036854775808
                                            _ = v49
                                        } else {
                                            v50 := GenS7{F0: 4294967295, F1: (v3 | -335)}
                                            v51 := v4[32768]
                                            _ = v50
                                            _ = v51
                                        }
                                        _ = v48
                                    }
                                    v3 -= (v3 & -294)
                                    if false {
                                        fmt.Println(fi28, v3, v5)
                                        v52 := GenS8{F0: (^(fi24))}
                                        var v53 uint64 = 0
                                        _ = v52
                                        _ = v53
                                    }
                                    v54 := GenS9{F0: 2}
                                    v6 = 1
                                    var v55 uint32 = 2
                                    _ = v54
                                    _ = v55
                                } else {
                                    if false {
                                        switch (v36 | fi24) {
                                        case -2147483648:
                                            _ = len(v4)
                                        case 32767:
                                            switch 0 {
                                            case 474:
                                                var v56 int8 = 0
                                                _ = v56
                                            }
                                        }
                                        var v57 string = "\n\t"
                                        v57 = (((v37 + ("foo bar" + (("hello" + v37) + v37))) + v57) + v57)
                                        _ = v57
                                    } else {
                                        v31 = ((fi28 & (fi28 + -1)) >= (fi28 - -2147483648))
                                        var v58 uint16 = (v6 + 34)
                                        _ = v58
                                    }
                                    var v59 int = (fi24 / (((fi28 % ((fi24) | 1))) | 1))
                                    var v60 uint = 18446744073709551615
                                    var v61 int8 = 127
                                    v62 := GenS10{F0: -128, F1: v60, F2: v61}
                                    _ = v59
                                    _ = v60
                                    _ = v61
                                    _ = v62
                                }
                                _ = v36
                                _ = v37
                            } else {
                                switch "\n\t" {
                                case "\n\t":
                                    switch (fi28 * fi28) {
                                    case 42:
                                        v6 = v6
                                    }
                                case "a":
                                    v5 = v31
                                }
                                fmt.Println(v5, fi24)
                                if false {
                                    if true {
                                        if (2147483647 <= 2) {
                                            switch ("hello" + ("\n\t" + ("hello" + "a"))) {
                                            case "hello":
                                                v1 -= (v1 >> ((4294967295) & 63))
                                            default:
                                                fmt.Println(fi28, v3)
                                            }
                                        }
                                        if (127 >= uint8(v6)) {
                                            var v63 uint8 = v1
                                            _ = v63
                                        } else {
                                            fmt.Println(fi24, v3, v3)
                                            var v64 int32 = -910
                                            var v65 int = 42
                                            _ = v64
                                            _ = v65
                                        }
                                    } else {
                                        v1 = v1
                                    }
                                    v66 := GenS3{F0: false, F1: (!(9223372036854775807 <= -1)), F2: (fi28 ^ (^(fi24)))}
                                    _ = v66
                                }
                            }
                            _ = v31
                        }
                        var v67 float64 = 3.14159
                        _ = fi28
                        _ = v67
                    }
                }
            }
        case "hello":
            if ((385 == 9223372036854775807) || (!false)) {
                fmt.Println(v1, v6)
                var v68 string = ("hello" + ("hello" + "foo bar"))
                _ = v68
            }
        case "":
            if (3.14159 <= 0.0) {
                var v69 uint16 = v6
                _ = v69
            } else {
                _ = len(v2)
                for fi70 := 0; fi70 < 10; fi70++ {
                    fmt.Println(v3, v5, v6)
                    var v71 bool = true
                    if v5 {
                        var v72 string = "hello"
                        if true {
                            if v71 {
                                v5 = v5
                                switch fi24 {
                                case 42:
                                    var v73 uint32 = 897
                                    _ = v73
                                default:
                                    switch (v72 + "") {
                                    case "hello":
                                        var v74 uint16 = (v6 % ((v6) | 1))
                                        _ = v74
                                    }
                                }
                            }
                            if false {
                                switch (-(fi70)) {
                                case -2147483648:
                                    var v75 int64 = 42
                                    _ = v75
                                case 588:
                                    if v5 {
                                        fmt.Println(v5, fi70, fi24)
                                        v76 := GenS2{F0: "foo bar", F1: 42, F2: ("a" + "foo bar")}
                                        _ = v76
                                    }
                                }
                                v5 = ((v5 && false) || false)
                                v77 := GenS11{F0: -32768, F1: -359}
                                _ = v77
                            } else {
                                v6 = 65535
                                if ((v1 % ((v1) | 1)) != (v1 & (^(v1)))) {
                                    if v5 {
                                        fmt.Println(v1, v5)
                                        var v78 string = (((v72 + v72) + "a") + v72)
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v71 = ((v1 <= 2) && (404 != 18446744073709551615))
                                        v3 += v3
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        if v71 {
                                            v72 = (((v78 + (v72 + v72)) + ("foo bar" + (v78 + "a"))) + (("\n\t" + "\n\t") + (v72 + "hello")))
                                            v79 := GenS12{F0: (^(v1)), F1: fi70}
                                            _ = v79
                                        } else {
                                            var v80 int32 = -1
                                            if ((v78 + v78) == (v72 + v72)) {
                                                if (true || true) {
                                                    switch v72 {
                                                    case "a":
                                                        delete(v4, fi70)
                                                    case "foo bar":
                                                        fmt.Println(v80, fi70, fi24)
                                                    case "\n\t":
                                                        if true {
                                                            var v81 float64 = 0.0
                                                            var v82 int64 = 0
                                                            v83 := GenS10{F0: 120, F1: 969, F2: 48}
                                                            _ = v81
                                                            _ = v82
                                                            _ = v83
                                                        }
                                                    }
                                                } else {
                                                    _ = len(v4)
                                                }
                                            } else {
                                                fmt.Println(v80, v3, v1)
                                            }
                                            _ = v80
                                        }
                                        _ = v78
                                    } else {
                                        var v84 string = v72
                                        _ = v84
                                    }
                                }
                                var v85 uint32 = 2
                                _ = v85
                            }
                        }
                        _ = v72
                    } else {
                        var v86 uint = 477
                        _ = v86
                    }
                    _ = fi70
                    _ = v71
                }
                for fi87 := 0; fi87 < 3; fi87++ {
                    v3 -= (v3 << ((42) & 63))
                    var v88 int32 = -2147483648
                    if v5 {
                        v89 := GenS2{F0: (((("" + "") + "a") + (((("a" + ("foo bar" + "foo bar")) + (("a" + "hello") + "\n\t")) + "foo bar") + "")) + (("a" + ("foo bar" + "\n\t")) + "a")), F1: 118, F2: "\n\t"}
                        v5 = ((float32(-2.25) * (1e10 - 0.1)) < float32(0.0))
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v89
                    }
                    _ = fi87
                    _ = v88
                }
            }
        }
        _ = fi24
        _ = v25
    }
    v1 += v1
    v90 := GenS13{F0: (v6 * (v6 >> ((2) & 63))), F1: false}
    var v91 int64 = -1
    for fi92 := 0; fi92 < 1; fi92++ {
        v2 = append(v2, (fi92 / (((fi92 + 9223372036854775807)) | 1)))
        v2 = append(v2, -9223372036854775808)
        _ = fi92
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v90
    _ = v91
}
