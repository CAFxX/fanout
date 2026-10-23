package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int8 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 int8; F1 uint; F2 int }
type GenS7 struct { F0 uint32; F1 string }
type GenS8 struct { F0 int8; F1 int32 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int64; F1 int }
func genFn0(p0 bool, p1 int8) bool {
    for fi7 := 0; fi7 < 10; fi7++ {
        p0 = false
        p1 = 42
        _ = fi7
    }
    v8 := []int{}
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 string = ""
        _ = fi9
        _ = v10
    }
    if p0 {
        var v11 uint8 = 42
        _ = v11
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        p1 = (p1 % (((p1 - (p1 / ((-1) | 1)))) | 1))
        var v13 string = "hello"
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v8
    return (false || true)
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{913}
    v3 := make(map[int]bool)
    var v4 int32 = 129
    v5 := []string{"hello"}
    v6 := []int{-989, 127, -283}
    for fi14 := 0; fi14 < 8; fi14++ {
        for fi15 := 0; fi15 < 6; fi15++ {
            if (18446744073709551615 <= 0) {
                var v16 int32 = (v4 * v4)
                var v17 float32 = 0.1
                var v18 int64 = -9223372036854775808
                _ = v16
                _ = v17
                _ = v18
            } else {
                v19 := GenS1{F0: 1, F1: (1 == 2), F2: -1}
                _ = v19
            }
            fmt.Println(fi14, v4, fi14)
            _ = fi15
        }
        var v20 uint64 = 0
        fmt.Println(v4, v20)
        _ = fi14
        _ = v20
    }
    v4 -= (v4 >> ((0) & 63))
    v21 := GenS2{F0: 37, F1: 879}
    if false {
        v4 -= v4
        var v22 int = 1
        _ = v22
    }
    var v23 uint16 = 65535
    var v24 bool = (-2.25 != (3.14159 - ((0.1 + 1e10) - 3.14159)))
    switch 0 {
    case 2147483647:
        if v24 {
            v4 = 0
        }
    case 9223372036854775807:
        for fi25 := 0; fi25 < 11; fi25++ {
            v26 := GenS2{F0: -1, F1: 0}
            for fi27 := 0; fi27 < 15; fi27++ {
                v24 = true
                if v24 {
                    var v28 uint = 42
                    switch "\n\t" {
                    case "a":
                        if (100 != 100) {
                            v4 = v4
                        }
                    case "\n\t":
                        switch "\n\t" {
                        case "hello":
                            if (-54 != 100) {
                                fmt.Println(v28, v23, fi27)
                            } else {
                                var v29 int = int(v23)
                                v28 = v28
                                v30 := GenS3{F0: "\n\t"}
                                fmt.Println(v24, fi25, fi25)
                                _ = v29
                                _ = v30
                            }
                        case "a":
                            v28 -= (v28 & (v28 * v28))
                        case "\n\t":
                            var v31 int64 = 0
                            _ = v31
                        }
                    }
                    _ = v28
                } else {
                    _ = len(v1)
                    switch "" {
                    case "foo bar":
                        if (!v24) {
                            switch -1 {
                            case 32768:
                                switch fi25 {
                                case 32768:
                                    v32 := GenS4{F0: 1, F1: 4294967295}
                                    _ = v32
                                case 127:
                                    var v33 float64 = 0.1
                                    var v34 uint64 = 18446744073709551615
                                    if ("foo bar" <= "\n\t") {
                                        if false {
                                            switch 84 {
                                            case 9223372036854775807:
                                                v35 := GenS5{F0: 737, F1: (v4 ^ 0)}
                                                _ = v35
                                            case 32768:
                                                v24 = v24
                                            }
                                            v1["hello"] = 285
                                        } else {
                                            var v36 uint16 = 42
                                            _ = v36
                                        }
                                    } else {
                                        if (v4 == v4) {
                                            var v37 int64 = -9223372036854775808
                                            switch 1 {
                                            case 32768:
                                                switch fi25 {
                                                case 967:
                                                    if (false && (!(v23 <= v23))) {
                                                        var v38 uint16 = (v23 << ((42) & 63))
                                                        _ = v38
                                                    } else {
                                                        var v39 int64 = (v37 % ((v37) | 1))
                                                        v4 = (v4 * 1)
                                                        _ = v39
                                                    }
                                                }
                                            default:
                                                v5 = append(v5, "a")
                                            }
                                            v33 = 1.5
                                            _ = v37
                                        } else {
                                            v1["hello"] = 4294967295
                                            switch ("" + "foo bar") {
                                            case "foo bar":
                                                v40 := GenS6{F0: -44, F1: 343, F2: (fi27 - (fi27 / ((fi27) | 1)))}
                                                _ = v40
                                            case "a":
                                                if v24 {
                                                    var v41 float32 = 1.5
                                                    _ = v41
                                                } else {
                                                    var v42 int64 = 1
                                                    fmt.Println(v24)
                                                    _ = v42
                                                }
                                            }
                                        }
                                        v43 := GenS7{F0: 42, F1: ("hello" + "")}
                                        v4 = v4
                                        _ = v43
                                    }
                                    var v44 int32 = (v4 + v4)
                                    _ = v33
                                    _ = v34
                                    _ = v44
                                case 682:
                                    var v45 float64 = 3.14159
                                    _ = v45
                                }
                            default:
                                switch "" {
                                case "hello":
                                    fmt.Println(v23, v24, v24)
                                case "a":
                                    _ = len(v6)
                                case "":
                                    v46 := GenS5{F0: v4, F1: v4}
                                    _ = v46
                                default:
                                    var v47 string = "\n\t"
                                    _ = v47
                                }
                            }
                        }
                    case "a":
                        fmt.Println(v24, fi25, v24)
                    default:
                        switch "\n\t" {
                        case "\n\t":
                            v23 += v23
                            _ = len(v6)
                            v48 := GenS7{F0: 2, F1: "hello"}
                            v5 = append(v5, ("\n\t" + ""))
                            _ = v48
                        case "a":
                            v23 = v23
                        case "foo bar":
                            if v24 {
                                var v49 int32 = v4
                                switch "\n\t" {
                                case "a":
                                    v23 = 65535
                                case "foo bar":
                                    fmt.Println(v24, fi25)
                                case "\n\t":
                                    var v50 uint = 0
                                    _ = v50
                                }
                                var v51 uint32 = 2147483648
                                if (false && true) {
                                    switch 127 {
                                    case 9223372036854775807:
                                        switch 128 {
                                        case -9223372036854775808:
                                            if (649 > 255) {
                                                if (1 >= 255) {
                                                    var v52 int8 = 127
                                                    _ = v52
                                                } else {
                                                    v53 := GenS8{F0: -105, F1: (v49 * v4)}
                                                    v49 += v4
                                                    var v54 bool = (v24 && v24)
                                                    _ = v53
                                                    _ = v54
                                                }
                                                fmt.Println(v49, v49)
                                            } else {
                                                fmt.Println(v4)
                                            }
                                            v51 = v51
                                        case 706:
                                            if ((v23 & (v23 | (v23 / ((v23) | 1)))) >= 255) {
                                                var v55 int32 = (-(v49))
                                                _ = v55
                                            }
                                        case 128:
                                            var v56 int = (fi27 + fi25)
                                            _ = v56
                                        default:
                                            if (!true) {
                                                var v57 string = "\n\t"
                                                switch 42 {
                                                case 1:
                                                    var v58 uint16 = v23
                                                    _ = v58
                                                case 4294967295:
                                                    switch 683 {
                                                    case -9223372036854775808:
                                                        if false {
                                                            switch ("a" + v57) {
                                                            case "hello":
                                                                var v59 uint8 = 255
                                                                _ = v59
                                                            case "\n\t":
                                                                v49 = -1
                                                            default:
                                                                if ((3.14159 < (3.14159 * 1e10)) || (-1 < -32768)) {
                                                                    v24 = (true || v24)
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                    v3[fi25] = true
                                                                } else {
                                                                    v24 = (0.1 < 0.0)
                                                                    v51 = (^(v51))
                                                                }
                                                            }
                                                            if v24 {
                                                                var v60 float64 = -2.25
                                                                var v61 uint = 65535
                                                                v24 = (!v24)
                                                                v49 = v4
                                                                _ = v60
                                                                _ = v61
                                                            } else {
                                                                if v24 {
                                                                    if v24 {
                                                                        var v62 int64 = -157
                                                                        switch ((v57 + v57) + v57) {
                                                                        case "a":
                                                                            v51 = (v51 & v51)
                                                                        case "\n\t":
                                                                            v57 = "foo bar"
                                                                        case "":
                                                                            v63 := GenS9{F0: v62}
                                                                            _ = v63
                                                                        default:
                                                                            v24 = false
                                                                        }
                                                                        _ = v62
                                                                    }
                                                                    v64 := GenS8{F0: 0, F1: 0}
                                                                    _ = v64
                                                                }
                                                            }
                                                            if (!(-611 == -9223372036854775808)) {
                                                                v65 := GenS9{F0: -1}
                                                                _ = v65
                                                            } else {
                                                                v66 := GenS10{F0: (fi27 | fi25)}
                                                                _ = v66
                                                            }
                                                        }
                                                    }
                                                default:
                                                    v23 = (v23 * 65535)
                                                }
                                                v67 := GenS8{F0: 100, F1: v49}
                                                _ = v57
                                                _ = v67
                                            }
                                        }
                                    }
                                    v24 = v24
                                }
                                _ = v49
                                _ = v51
                            } else {
                                fmt.Println(v4)
                            }
                        }
                    }
                    v68 := v3[(fi25 << ((32) & 63))]
                    _ = len(v6)
                    if v24 {
                        var v69 bool = (v4 < v4)
                        _ = v69
                    } else {
                        if len(v2) > 0 {
                            v2[0] = 993
                        }
                        _ = len(v5)
                        if (((136 != 5) && v24) && v68) {
                            v23 += (^(v23))
                            var v70 uint16 = 270
                            _ = v70
                        }
                    }
                    var v71 bool = v24
                    v72 := GenS9{F0: -9223372036854775808}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v68
                    _ = v71
                    _ = v72
                }
                _ = fi27
            }
            v23 -= v23
            _ = fi25
            _ = v26
        }
    }
    if (1 >= -1) {
        var v73 float64 = ((((0.0 - 3.14159) - -2.25) - -2.25) - ((1.5 - (1.5 / (((0.1 + -2.25) - 1e10) - -2.25))) / float64(((1e10 + (1.5 * 1.5)) / ((3.14159 / 1e10) - ((1.5 + -2.25) / 1e10))))))
        _ = v73
    } else {
        v23 = v23
    }
    if v24 {
        if false {
            delete(v1, "a")
            fmt.Println(v24)
        } else {
            v24 = (3.14159 >= (3.14159 * (0.0 - 1e10)))
            v74 := GenS5{F0: v4, F1: -968}
            _ = v74
        }
        fmt.Println(v24, v23)
    }
    if false {
        var v75 uint64 = 975
        for fi76 := 0; fi76 < 5; fi76++ {
            v24 = v24
            var v77 int8 = 42
            var v78 int16 = 32767
            _ = fi76
            _ = v77
            _ = v78
        }
        _ = v75
    } else {
        switch (("a" + "") + ("a" + "\n\t")) {
        case "a":
            var v79 int32 = -1
            _ = v79
        case "":
            switch "\n\t" {
            case "a":
                v23 -= v23
            case "foo bar":
                if (-2.25 >= float32((0.1 / 0.1))) {
                    var v80 int64 = -1
                    _ = v80
                } else {
                    v81 := GenS11{F0: 0, F1: 4294967295}
                    _ = v81
                }
            case "":
                v82 := GenS4{F0: 42, F1: 913}
                _ = v82
            default:
                v83 := GenS7{F0: 980, F1: ("" + (("hello" + (("" + ("hello" + ("foo bar" + ((("a" + ("a" + "hello")) + "") + ("\n\t" + "hello"))))) + "foo bar")) + ("hello" + "foo bar")))}
                _ = v83
            }
        default:
            switch "\n\t" {
            case "hello":
                if (843 >= 849) {
                    fmt.Println(v23, v23)
                }
            case "":
                if len(v2) > 0 {
                    v2[0] = 9223372036854775807
                }
            case "a":
                if (v4 > 42) {
                    var v84 bool = v24
                    _ = v84
                } else {
                    if false {
                        for fi85 := 0; fi85 < 13; fi85++ {
                            var v86 float32 = 3.14159
                            _ = fi85
                            _ = v86
                        }
                    }
                    delete(v3, -934)
                    v1["a"] = 2147483647
                }
            }
        }
        v87 := GenS4{F0: 1, F1: 609}
        v88 := v3[127]
        _ = v87
        _ = v88
    }
    var v89 bool = (!(255 <= v23))
    var v90 int32 = 0
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v89
    _ = v90
}
