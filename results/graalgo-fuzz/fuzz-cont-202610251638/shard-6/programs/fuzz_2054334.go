package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 bool; F2 int16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int; F1 string; F2 int8 }
type GenS7 struct { F0 uint8; F1 uint64; F2 bool }
type GenS8 struct { F0 int16; F1 int8; F2 string }
type GenS9 struct { F0 uint; F1 int; F2 uint8 }
type GenS10 struct { F0 int8; F1 int32 }
type GenS11 struct { F0 int8; F1 uint }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint; F1 bool }
type GenS14 struct { F0 string; F1 int8; F2 int16 }
type GenS15 struct { F0 int8; F1 int8; F2 string }
type GenS16 struct { F0 uint; F1 int8; F2 uint16 }
func genFn0(p0 int, p1 uint32, p2 string) int8 {
    fmt.Println(p2, p1, p1)
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 20
}
func genFn1(p0 bool, p1 uint64, p2 int) int32 {
    if p0 {
        var v8 float64 = float64(1e10)
        _ = v8
    }
    var v9 bool = ((p2 | p2) != p2)
    fmt.Println(p1)
    fmt.Println(p2, p0, p2)
    v10 := []int{p2, (p2 * p2), 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return -2147483648
}

func main() {
    var v1 float64 = (0.1 * ((3.14159 - 1.5) - (0.1 + (0.0 - 1e10))))
    var v2 int32 = 42
    var v3 uint64 = 4294967295
    v4 := []string{"foo bar", (((("" + "\n\t") + "a") + "") + (("\n\t" + "\n\t") + "hello"))}
    var v5 float32 = -2.25
    v6 := make(map[int]uint32)
    for fi11 := 0; fi11 < 14; fi11++ {
        switch ("\n\t" + (("" + (("hello" + ("\n\t" + "\n\t")) + ("a" + "a"))) + "foo bar")) {
        case "hello":
            if (-2.25 < v5) {
                switch fi11 {
                case 128:
                    v12 := v6[-93]
                    _ = v12
                }
                if false {
                    delete(v6, fi11)
                    switch "hello" {
                    case "a":
                        var v13 uint32 = 65535
                        _ = v13
                    }
                } else {
                    if true {
                        fmt.Println(v5)
                    }
                }
            }
        case "\n\t":
            var v14 uint = uint(v3)
            _ = v14
        default:
            v3 = 0
        }
        for fi15 := 0; fi15 < 17; fi15++ {
            if (((v1 * 1e10) == v1) && (true && true)) {
                if false {
                    if true {
                        var v16 string = "a"
                        v17 := GenS1{F0: 127}
                        v18 := GenS1{F0: 109}
                        switch v16 {
                        case "a":
                            v19 := GenS1{F0: 81}
                            _ = v19
                        case "":
                            var v20 int8 = 0
                            _ = v20
                        }
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    v2 = 42
                    if (!true) {
                        var v21 int8 = -72
                        if ((v3 > (v3 * (v3 & v3))) && true) {
                            if (false || true) {
                                var v22 uint = 255
                                v5 += v5
                                _ = v22
                            } else {
                                var v23 uint = 29
                                _ = v23
                            }
                            if false {
                                v3 = 590
                                v1 += float64(v5)
                            }
                            switch (fi15 ^ 127) {
                            case -1:
                                v24 := GenS1{F0: 128}
                                _ = v24
                            case 128:
                                var v25 int16 = -981
                                _ = v25
                            case -9223372036854775808:
                                _ = len(v6)
                            default:
                                v21 = 100
                            }
                            var v26 float32 = (v5 * 3.14159)
                            _ = v26
                        } else {
                            switch fi15 {
                            case 128:
                                if false {
                                    fmt.Println(v1, fi11)
                                }
                            case -630:
                                if true {
                                    if (v21 <= genFn0(fi15, 65535, "")) {
                                        switch fi15 {
                                        case 0:
                                            if true {
                                                var v27 int16 = -32768
                                                if (42 != genFn1((false && true), (v3 / (((v3 % ((389) | 1))) | 1)), fi15)) {
                                                    v3 -= (v3 << ((18446744073709551615) & 63))
                                                    v28 := GenS1{F0: 128}
                                                    fmt.Println(fi11, v5, v1)
                                                    switch (fi11 & fi11) {
                                                    case -7:
                                                        var v29 uint16 = 42
                                                        _ = v29
                                                    case 4294967295:
                                                        switch "" {
                                                        case "\n\t":
                                                            switch (fi15 | fi15) {
                                                            case 818:
                                                                fmt.Println(v3, v21)
                                                                v2 = -2147483648
                                                            }
                                                        case "":
                                                            if true {
                                                                v30 := GenS2{F0: 79, F1: "a"}
                                                                _ = len(v4)
                                                                var v31 string = "\n\t"
                                                                _ = v30
                                                                _ = v31
                                                            }
                                                            var v32 float32 = (3.14159 + 3.14159)
                                                            v32 += (0.1 - (0.0 * -2.25))
                                                            _ = v32
                                                        case "hello":
                                                            switch 32768 {
                                                            case -212:
                                                                if true {
                                                                    v3 -= v3
                                                                    v4 = append(v4, "foo bar")
                                                                    fmt.Println(v27, fi11)
                                                                }
                                                            case 4294967295:
                                                                v33 := GenS3{F0: v27}
                                                                _ = v33
                                                            case 0:
                                                                switch 607 {
                                                                case 315:
                                                                    var v34 int8 = v21
                                                                    _ = v34
                                                                case 131:
                                                                    var v35 uint16 = 740
                                                                    _ = v35
                                                                case 9223372036854775807:
                                                                    fmt.Println(fi15, v21)
                                                                default:
                                                                    if (494 != 0) {
                                                                        v36 := GenS1{F0: 168}
                                                                        _ = v36
                                                                    }
                                                                }
                                                            default:
                                                                fmt.Println(v1, v21)
                                                            }
                                                            var v37 uint8 = 1
                                                            _ = v37
                                                        }
                                                    case -9223372036854775808:
                                                        var v38 uint = 18446744073709551615
                                                        _ = v38
                                                    default:
                                                        v6[-676] = 0
                                                        var v39 int64 = 0
                                                        _ = v39
                                                    }
                                                    _ = v28
                                                } else {
                                                    var v40 uint32 = 4294967295
                                                    v27 += v27
                                                    _ = v40
                                                }
                                                v41 := GenS4{F0: 255, F1: false, F2: (v27 ^ v27)}
                                                _ = v27
                                                _ = v41
                                            }
                                        }
                                        v42 := GenS1{F0: 255}
                                        var v43 float32 = (0.1 + v5)
                                        _ = v42
                                        _ = v43
                                    } else {
                                        v2 = (v2 & v2)
                                    }
                                } else {
                                    switch fi11 {
                                    case 32767:
                                        v44 := GenS5{F0: (v21 - (-(v21)))}
                                        _ = v44
                                    }
                                    var v45 float64 = float64(v5)
                                    if ((fi11 >= fi15) || ((v1 - v1) == v1)) {
                                        fmt.Println(v21, v21)
                                        switch ("" + "a") {
                                        case "a":
                                            var v46 int32 = v2
                                            switch "foo bar" {
                                            case "\n\t":
                                                switch fi11 {
                                                case 32768:
                                                    fmt.Println(v3)
                                                case 948:
                                                    var v47 uint = 18446744073709551615
                                                    _ = v47
                                                case -253:
                                                    fmt.Println(v45, v46, v3)
                                                default:
                                                    v48 := GenS6{F0: (fi11 & (fi11 / ((fi11) | 1))), F1: "a", F2: genFn0((fi15 | (^(fi11))), 2, "foo bar")}
                                                    _ = v48
                                                }
                                            case "foo bar":
                                                fmt.Println(v21)
                                            default:
                                                v49 := GenS7{F0: 255, F1: v3, F2: (2 == 917)}
                                                _ = v49
                                            }
                                            _ = v46
                                        }
                                        v50 := GenS1{F0: 127}
                                        _ = v50
                                    } else {
                                        switch (("hello" + ("foo bar" + (("" + "") + ""))) + "\n\t") {
                                        case "\n\t":
                                            var v51 uint = 255
                                            _ = v51
                                        }
                                        var v52 bool = (-1 == -791)
                                        var v53 uint32 = 42
                                        _ = v52
                                        _ = v53
                                    }
                                    if true {
                                        v3 = 18446744073709551615
                                        var v54 bool = false
                                        v55 := GenS8{F0: -32768, F1: -128, F2: ("foo bar" + "")}
                                        _ = v54
                                        _ = v55
                                    } else {
                                        fmt.Println(v1, v21)
                                        v56 := GenS8{F0: 32767, F1: 0, F2: (("a" + "a") + "a")}
                                        _ = v56
                                    }
                                    fmt.Println(fi15, v3, v1)
                                    _ = v45
                                }
                            case 0:
                                v57 := GenS9{F0: 42, F1: fi15, F2: 255}
                                _ = v57
                            }
                            delete(v6, fi11)
                            v1 -= 1e10
                        }
                        if len(v4) > 0 {
                            v4[0] = "\n\t"
                        }
                        _ = v21
                    }
                    var v58 bool = (255 == 0)
                    delete(v6, 2147483647)
                    _ = v58
                } else {
                    v2 = v2
                    if (-128 >= -128) {
                        v2 -= 1
                    } else {
                        v1 -= v1
                        var v59 uint32 = 2147483647
                        _ = v59
                    }
                }
                var v60 uint = 18446744073709551615
                switch 32767 {
                case -245:
                    var v61 uint64 = v3
                    var v62 float32 = (v5 / (v5 / v5))
                    switch (fi15 >> ((v60) & 63)) {
                    case 42:
                        fmt.Println(v62, v1)
                    case 32767:
                        fmt.Println(v1, v3, v2)
                    case 127:
                        v62 -= 0.1
                        var v63 uint64 = (v3 | 2)
                        v64 := GenS2{F0: 255, F1: "\n\t"}
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if (v3 >= 1) {
                            v4 = append(v4, (("foo bar" + "\n\t") + "hello"))
                            if (100 <= 127) {
                                v65 := GenS10{F0: -128, F1: (v2 + (-(v2)))}
                                _ = v65
                            } else {
                                v60 += v60
                                if false {
                                    v62 -= v62
                                    v5 = (v62 + v62)
                                } else {
                                    var v66 bool = true
                                    var v67 int32 = 0
                                    _ = v66
                                    _ = v67
                                }
                                if true {
                                    fmt.Println(v61)
                                } else {
                                    v63 -= v3
                                }
                            }
                            v68 := GenS4{F0: uint(v61), F1: true, F2: 1000}
                            _ = v68
                        }
                        _ = v63
                        _ = v64
                    default:
                        v4 = append(v4, "a")
                    }
                    v69 := GenS11{F0: -1, F1: 436}
                    switch fi15 {
                    case 32767:
                        switch "\n\t" {
                        case "":
                            if (!(!false)) {
                                var v70 float64 = v1
                                _ = v70
                            } else {
                                v61 += uint64(v60)
                            }
                        case "hello":
                            switch fi15 {
                            case 1:
                                switch ("hello" + "hello") {
                                case "hello":
                                    v5 = 3.14159
                                case "foo bar":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                default:
                                    switch "" {
                                    case "a":
                                        v71 := GenS12{F0: 2}
                                        _ = v71
                                    default:
                                        var v72 uint8 = 127
                                        _ = v72
                                    }
                                }
                            case -2147483648:
                                v2 = (v2 << ((v60) & 63))
                            }
                        case "a":
                            v73 := GenS13{F0: (v60 % ((v60) | 1)), F1: true}
                            _ = v73
                        }
                    case 1:
                        var v74 uint = v60
                        _ = v74
                    }
                    _ = len(v6)
                    var v75 uint8 = 42
                    v62 = 0.0
                    _ = v61
                    _ = v62
                    _ = v69
                    _ = v75
                case 128:
                    var v76 uint16 = 42
                    if true {
                        var v77 uint16 = v76
                        _ = v77
                    }
                    var v78 uint64 = (v3 | (v3 << ((255) & 63)))
                    _ = v76
                    _ = v78
                default:
                    if len(v4) > 0 {
                        v4[0] = "\n\t"
                    }
                }
                _ = v60
            }
            var v79 uint16 = 0
            var v80 bool = (42 < 1)
            _ = fi15
            _ = v79
            _ = v80
        }
        _ = fi11
    }
    var v81 int64 = 9223372036854775807
    var v82 int32 = (^(v2))
    v3 = v3
    _ = len(v4)
    if true {
        v83 := GenS14{F0: "foo bar", F1: 127, F2: 1}
        v82 -= v82
        v84 := GenS15{F0: 42, F1: genFn0(42, 202, ("hello" + "\n\t")), F2: ""}
        _ = v83
        _ = v84
    } else {
        var v85 float32 = (v5 + float32((0.0 + v1)))
        v86 := GenS16{F0: 529, F1: genFn0(9223372036854775807, 120, "\n\t"), F2: 714}
        var v87 uint32 = 65535
        _ = v85
        _ = v86
        _ = v87
    }
    fmt.Println(v5)
    if (v5 <= float32(v1)) {
        v88 := GenS10{F0: 1, F1: (v2 ^ v2)}
        v89 := v6[127]
        if (v1 < v1) {
            var v90 string = "foo bar"
            var v91 uint = 975
            _ = v90
            _ = v91
        }
        _ = v88
        _ = v89
    }
    v82 = v2
    var v92 float64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v81
    _ = v82
    _ = v92
}
