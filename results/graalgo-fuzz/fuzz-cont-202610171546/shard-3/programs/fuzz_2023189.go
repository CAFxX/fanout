package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint64; F1 bool; F2 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint64; F1 bool; F2 bool }

func main() {
    var v1 bool = false
    v2 := make(map[string]int)
    v3 := make(map[string]int)
    var v4 bool = (0 < 42)
    var v5 int8 = -23
    v6 := []string{"foo bar"}
    var v7 float32 = (-2.25 * float32(0.0))
    fmt.Println(v7, v1, v5)
    for fi8 := 0; fi8 < 16; fi8++ {
        if v4 {
            for fi9 := 0; fi9 < 19; fi9++ {
                if v1 {
                    v10 := GenS1{F0: 4294967295, F1: -32768}
                    v11 := GenS1{F0: 1, F1: 756}
                    delete(v3, "\n\t")
                    _ = v10
                    _ = v11
                } else {
                    fmt.Println(v5)
                    v12 := GenS1{F0: 2, F1: 42}
                    v7 = v7
                    _ = v12
                }
                var v13 bool = v4
                var v14 uint64 = 18446744073709551615
                _ = fi9
                _ = v13
                _ = v14
            }
        } else {
            var v15 uint8 = 140
            v4 = (v1 && (1 < uint16(v15)))
            _ = v15
        }
        v7 = (v7 + (float32(0.0) - float32((-2.25 * float64(-2.25)))))
        fmt.Println(v4, fi8, v7)
        _ = fi8
    }
    v2["foo bar"] = 9223372036854775807
    for fi16 := 0; fi16 < 14; fi16++ {
        if v1 {
            v1 = v4
            for fi17 := 0; fi17 < 5; fi17++ {
                v18 := GenS1{F0: 18446744073709551615, F1: 42}
                v19 := GenS1{F0: 1, F1: -32768}
                _ = fi17
                _ = v18
                _ = v19
            }
        }
        _ = fi16
    }
    v7 = 3.14159
    switch ("" + ("foo bar" + "hello")) {
    case "foo bar":
        v20 := v2["hello"]
        _ = v20
    case "\n\t":
        if v4 {
            v5 = v5
        }
    case "a":
        for fi21 := 0; fi21 < 18; fi21++ {
            var v22 int = fi21
            if true {
                v23 := GenS1{F0: 42, F1: 1}
                v24 := GenS1{F0: 2, F1: -32768}
                _ = v23
                _ = v24
            } else {
                v25 := GenS2{F0: 42, F1: true, F2: "\n\t"}
                _ = v25
            }
            switch (-(fi21)) {
            case 128:
                var v26 uint32 = 0
                _ = v26
            case 32768:
                if v1 {
                    for fi27 := 0; fi27 < 12; fi27++ {
                        v1 = (127 <= 106)
                        switch (fi27 % ((-961) | 1)) {
                        case 739:
                            if (18446744073709551615 <= 4294967295) {
                                var v28 uint8 = 128
                                if (fi21 > (-(v22))) {
                                    var v29 bool = false
                                    if (1.5 > (1e10 / 1e10)) {
                                        v30 := GenS3{F0: (-(fi21))}
                                        _ = v30
                                    }
                                    var v31 float64 = 3.14159
                                    _ = v29
                                    _ = v31
                                }
                                switch int(v5) {
                                case 870:
                                    if v1 {
                                        if (65535 > 1) {
                                            var v32 uint32 = 561
                                            switch (-(fi21)) {
                                            case -740:
                                                v22 -= fi21
                                                if true {
                                                    var v33 int16 = 1000
                                                    if ((^(v5)) > v5) {
                                                        fmt.Println(v1, v22)
                                                        var v34 string = "\n\t"
                                                        var v35 uint8 = v28
                                                        var v36 int32 = 99
                                                        _ = v34
                                                        _ = v35
                                                        _ = v36
                                                    }
                                                    if (698 > 2) {
                                                        var v37 uint16 = 32767
                                                        var v38 float64 = ((3.14159 / 1e10) * -2.25)
                                                        fmt.Println(v5, v28, fi27)
                                                        _ = v37
                                                        _ = v38
                                                    } else {
                                                        var v39 int8 = 1
                                                        if (v1 && v4) {
                                                            if (v1 || (!v1)) {
                                                                switch v22 {
                                                                case -40:
                                                                    v33 = (v33 | v33)
                                                                case 128:
                                                                    v40 := GenS2{F0: 2, F1: (-1 <= 0), F2: ""}
                                                                    _ = v40
                                                                case -429:
                                                                    v41 := GenS1{F0: 0, F1: int16(fi27)}
                                                                    _ = v41
                                                                }
                                                            }
                                                            var v42 int8 = v39
                                                            v39 += v39
                                                            _ = v42
                                                        } else {
                                                            v43 := GenS3{F0: 798}
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            _ = v43
                                                        }
                                                        var v44 int16 = (v33 * int16(v32))
                                                        _ = v39
                                                        _ = v44
                                                    }
                                                    _ = v33
                                                } else {
                                                    v7 += 1e10
                                                    if (((42 == 1) || ((-93 >= 0) && (613 <= (^(v32))))) && (-63 != v5)) {
                                                        var v45 int64 = 9223372036854775807
                                                        _ = v45
                                                    } else {
                                                        switch v22 {
                                                        case 2147483647:
                                                            var v46 uint16 = 781
                                                            _ = v46
                                                        default:
                                                            if v4 {
                                                                if (-720 > 1000) {
                                                                    if v4 {
                                                                        fmt.Println(fi21)
                                                                        v7 -= float32(1e10)
                                                                    }
                                                                    v4 = ((^(v32)) <= v32)
                                                                    if (193 < v32) {
                                                                        var v47 uint8 = 128
                                                                        _ = v47
                                                                    } else {
                                                                        if (("foo bar" + (("hello" + "\n\t") + "")) < "foo bar") {
                                                                            v4 = (128 > (^(v28)))
                                                                        }
                                                                    }
                                                                }
                                                                var v48 uint16 = 807
                                                                v5 -= (^(v5))
                                                                _ = v48
                                                            } else {
                                                                v7 = v7
                                                            }
                                                        }
                                                        if true {
                                                            v5 += -115
                                                            var v49 uint64 = 929
                                                            var v50 int8 = 100
                                                            _ = v49
                                                            _ = v50
                                                        }
                                                    }
                                                    var v51 uint16 = 794
                                                    _ = v51
                                                }
                                                if true {
                                                    var v52 int32 = 1
                                                    switch (fi27 % ((fi27) | 1)) {
                                                    case 128:
                                                        if false {
                                                            v2["foo bar"] = (fi21 ^ -255)
                                                            var v53 bool = true
                                                            var v54 int16 = 32767
                                                            _ = v53
                                                            _ = v54
                                                        }
                                                        v55 := GenS4{F0: 0}
                                                        v56 := GenS3{F0: (^(fi21))}
                                                        _ = v55
                                                        _ = v56
                                                    }
                                                    switch ("hello" + ("" + (("hello" + "a") + ("a" + "")))) {
                                                    case "foo bar":
                                                        switch "" {
                                                        case "foo bar":
                                                            if ((v4 && (v7 == 3.14159)) || (!(v32 == (v32 - (v32 | 2))))) {
                                                                v57 := GenS2{F0: 18446744073709551615, F1: (-281 <= 955), F2: "foo bar"}
                                                                _ = v57
                                                            }
                                                        case "hello":
                                                            var v58 float32 = v7
                                                            _ = v58
                                                        case "a":
                                                            v59 := GenS5{F0: (v52 << ((2) & 63)), F1: v52}
                                                            _ = v59
                                                        default:
                                                            var v60 uint8 = (v28 ^ 71)
                                                            v60 -= (v28 ^ (v60 % (((v28 & (v60 >> ((255) & 63)))) | 1)))
                                                            switch ("" + "") {
                                                            case "hello":
                                                                fmt.Println(v5)
                                                            default:
                                                                if ((v32 ^ (v32 % ((1) | 1))) == (^(v32))) {
                                                                    if v1 {
                                                                        var v61 int16 = 32767
                                                                        if (float64((v7 / 1.0)) <= (3.14159 + ((0.0 + 3.14159) / (0.0 - (1e10 / 0.1))))) {
                                                                            var v62 uint64 = 4294967295
                                                                            v22 = 128
                                                                            fmt.Println(v5, v1, v7)
                                                                            _ = v62
                                                                        } else {
                                                                            var v63 float32 = v7
                                                                            v63 += ((v63 / 1.5) / (v7 + v63))
                                                                            v64 := GenS5{F0: (v52 | v52), F1: (v52 | v52)}
                                                                            _ = v63
                                                                            _ = v64
                                                                        }
                                                                        _ = v61
                                                                    }
                                                                    switch "a" {
                                                                    case "a":
                                                                        if len(v6) >= 2 {
                                                                            v6 = v6[1:len(v6)]
                                                                        }
                                                                    default:
                                                                        fmt.Println(fi21, v5)
                                                                    }
                                                                } else {
                                                                    var v65 uint8 = v60
                                                                    _ = v65
                                                                }
                                                            }
                                                            _ = v60
                                                        }
                                                    case "hello":
                                                        v66 := GenS4{F0: 0}
                                                        _ = v66
                                                    case "":
                                                        switch (("hello" + ("a" + "hello")) + "a") {
                                                        case "foo bar":
                                                            v5 = v5
                                                        default:
                                                            var v67 int8 = (v5 ^ 127)
                                                            _ = v67
                                                        }
                                                    }
                                                    _ = v52
                                                } else {
                                                    if (!false) {
                                                        if (32767 < 1) {
                                                            v1 = (771 < 659)
                                                            var v68 int8 = -128
                                                            if false {
                                                                var v69 uint32 = (v32 ^ v32)
                                                                switch ("\n\t" + "\n\t") {
                                                                case "hello":
                                                                    v1 = ((^(v68)) < v5)
                                                                case "":
                                                                    fmt.Println(v69, v1)
                                                                default:
                                                                    if (fi21 == fi21) {
                                                                        v70 := GenS4{F0: int16(v22)}
                                                                        fmt.Println(v7, v69, v68)
                                                                        _ = v70
                                                                    }
                                                                }
                                                                fmt.Println(v32)
                                                                _ = v69
                                                            } else {
                                                                v71 := v2["foo bar"]
                                                                v22 = int(v28)
                                                                fmt.Println(v4)
                                                                _ = v71
                                                            }
                                                            _ = v68
                                                        }
                                                        var v72 string = ("hello" + "a")
                                                        _ = v72
                                                    }
                                                    if len(v6) > 0 {
                                                        v6[0] = "foo bar"
                                                    }
                                                    fmt.Println(fi27)
                                                }
                                            default:
                                                v6 = append(v6, ("a" + "foo bar"))
                                                fmt.Println(fi21)
                                            }
                                            _ = v32
                                        } else {
                                            var v73 float64 = 0.1
                                            var v74 int32 = 0
                                            _ = v73
                                            _ = v74
                                        }
                                        var v75 uint = 65535
                                        var v76 float32 = ((((v7 + v7) * v7) - 3.14159) - float32(0.0))
                                        _ = v75
                                        _ = v76
                                    } else {
                                        v1 = v1
                                        _ = len(v6)
                                        fmt.Println(v28)
                                    }
                                case 32768:
                                    if true {
                                        var v77 uint = 4294967295
                                        v78 := GenS5{F0: 746, F1: 2147483647}
                                        var v79 uint32 = 1
                                        _ = v77
                                        _ = v78
                                        _ = v79
                                    } else {
                                        var v80 float32 = ((1e10 * v7) + (v7 * 0.0))
                                        if (!v4) {
                                            v22 -= (fi21 + (fi21 & 9223372036854775807))
                                            var v81 string = "hello"
                                            _ = v81
                                        } else {
                                            v80 = (v7 / (0.1 + -2.25))
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            var v82 bool = v4
                                            _ = v82
                                        }
                                        _ = v80
                                    }
                                default:
                                    fmt.Println(fi27, v22, v22)
                                }
                                _ = v28
                            }
                            if v1 {
                                if v4 {
                                    switch (v22 << ((255) & 63)) {
                                    case 32768:
                                        if v1 {
                                            fmt.Println(v1, v22, v22)
                                            fmt.Println(v22, fi27, v4)
                                        } else {
                                            v83 := GenS6{F0: v5}
                                            _ = v83
                                        }
                                    case -9223372036854775808:
                                        v5 += (-(v5))
                                    default:
                                        var v84 int = fi27
                                        _ = v84
                                    }
                                    var v85 int64 = -9223372036854775808
                                    v1 = (18446744073709551615 == 97)
                                    _ = v85
                                }
                                var v86 uint8 = 128
                                _ = v86
                            }
                            var v87 bool = v1
                            _ = v87
                        case 502:
                            if v1 {
                                v3["a"] = -9223372036854775808
                                v5 = 1
                            } else {
                                v22 -= 0
                                v88 := GenS7{F0: 0, F1: v4, F2: v1}
                                fmt.Println(v1)
                                _ = v88
                            }
                        }
                        var v89 int8 = v5
                        var v90 uint16 = 843
                        _ = fi27
                        _ = v89
                        _ = v90
                    }
                    _ = len(v3)
                }
            case 321:
                v6 = append(v6, "a")
            }
            _ = fi21
            _ = v22
        }
    default:
        v91 := GenS3{F0: 1}
        _ = v91
    }
    var v92 int = -1
    v93 := GenS6{F0: (v5 / (((v5 * v5)) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v92
    _ = v93
}
