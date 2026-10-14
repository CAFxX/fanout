package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 string }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 string; F1 int }
type GenS8 struct { F0 int16; F1 int8 }
type GenS9 struct { F0 int64 }
func genFn0(p0 string) bool {
    for fi7 := 0; fi7 < 2; fi7++ {
        if true {
            p0 = p0
            v8 := []int{fi7, fi7, fi7}
            _ = v8
        }
        _ = fi7
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        p0 = "hello"
        var v10 uint8 = 213
        var v11 int64 = 1
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := []int{847, 127, 151, 203}
    _ = p0
    _ = v12
    return false
}
func genFn1(p0 int8, p1 int64) uint64 {
    p1 -= (-(p1))
    var v13 uint64 = 0
    v13 = v13
    v14 := []int{4294967295, -705}
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return v13
}

func main() {
    var v1 uint32 = 89
    v2 := []string{}
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]int)
    var v5 bool = true
    v6 := make(map[int]bool)
    v15 := GenS1{F0: 2, F1: 42, F2: (("" + "hello") + ("" + ("a" + "a")))}
    for fi16 := 0; fi16 < 13; fi16++ {
        v17 := GenS2{F0: (("foo bar" + "hello") + "hello"), F1: "foo bar"}
        fmt.Println(fi16)
        delete(v4, 4294967295)
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        fmt.Println(fi18, fi18, v5)
        for fi19 := 0; fi19 < 10; fi19++ {
            switch "foo bar" {
            case "hello":
                v20 := v6[9223372036854775807]
                _ = v20
            case "":
                v5 = v5
            case "a":
                v1 = (v1 << ((622) & 63))
            default:
                var v21 uint16 = 65535
                _ = v21
            }
            var v22 float32 = 1e10
            if len(v2) > 0 {
                v2[0] = ""
            }
            _ = fi19
            _ = v22
        }
        for fi23 := 0; fi23 < 2; fi23++ {
            v1 -= (^(v1))
            _ = fi23
        }
        _ = fi18
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        var v25 float32 = (((1.5 - -2.25) + (-2.25 * 0.0)) - 1e10)
        if (v5 || ((32767 <= -1) && (0 >= -1))) {
            fmt.Println(v25, fi24)
            fmt.Println(v5, v5)
            var v26 uint = 2
            _ = v26
        } else {
            var v27 int16 = 1
            if v5 {
                for fi28 := 0; fi28 < 3; fi28++ {
                    v29 := GenS3{F0: 128, F1: -2147483648, F2: 211}
                    _ = len(v4)
                    v27 = 1
                    _ = fi28
                    _ = v29
                }
                var v30 int = (fi24 * fi24)
                _ = v30
            }
            var v31 float32 = (((0.1 + v25) / (((v25 - v25) + float32(((float64(3.14159) / -2.25) / 1.5))) / (v25 * v25))) - float32(1e10))
            _ = v27
            _ = v31
        }
        _ = fi24
        _ = v25
    }
    fmt.Println(v1, v5)
    switch "foo bar" {
    case "":
        _ = len(v6)
    case "foo bar":
        var v32 int = -9223372036854775808
        _ = v32
    default:
        var v33 int32 = -413
        _ = v33
    }
    v1 -= v1
    var v34 int = 4294967295
    for fi35 := 0; fi35 < 14; fi35++ {
        switch ((("\n\t" + "foo bar") + "") + "hello") {
        case "\n\t":
            v1 -= 65535
        default:
            for fi36 := 0; fi36 < 8; fi36++ {
                if v5 {
                    v1 -= (^(v1))
                    fmt.Println(v34)
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi36
            }
        }
        _ = fi35
    }
    switch v34 {
    case -419:
        fmt.Println(v1)
    default:
        v37 := GenS4{F0: (18446744073709551615 > genFn1(3, 711))}
        _ = v37
    }
    switch "hello" {
    case "hello":
        if v5 {
            if (v5 || (((1.5 * (1e10 + (1.5 - 1e10))) * (3.14159 + (0.0 * (1.5 + 0.0)))) >= -2.25)) {
                if len(v2) > 0 {
                    v2[0] = (("" + "\n\t") + "")
                }
                v2 = append(v2, "foo bar")
            } else {
                v38 := GenS5{F0: 42, F1: 364}
                _ = v38
            }
            switch ("a" + ("a" + ("" + "a"))) {
            case "foo bar":
                for fi39 := 0; fi39 < 16; fi39++ {
                    for fi40 := 0; fi40 < 5; fi40++ {
                        var v41 string = "\n\t"
                        _ = fi40
                        _ = v41
                    }
                    v2 = append(v2, ("" + "foo bar"))
                    if (false && v5) {
                        for fi42 := 0; fi42 < 3; fi42++ {
                            var v43 uint64 = 4294967295
                            var v44 uint64 = (v43 / (((v43 ^ (v43 & (v43 / ((v43) | 1))))) | 1))
                            if false {
                                var v45 float64 = 0.0
                                var v46 uint = 836
                                _ = v45
                                _ = v46
                            }
                            switch "" {
                            case "a":
                                v44 = 2
                                if (true || v5) {
                                    switch -2147483648 {
                                    case 579:
                                        v6[2147483647] = true
                                        v47 := GenS6{F0: (v1 >> ((381) & 63))}
                                        var v48 string = ("hello" + "")
                                        _ = v47
                                        _ = v48
                                    case 1:
                                        v2 = append(v2, "\n\t")
                                    case 4294967295:
                                        switch (fi39 >> ((0) & 63)) {
                                        case 984:
                                            v1 = v1
                                            fmt.Println(v43, fi39)
                                            v1 = (v1 & v1)
                                        default:
                                            v49 := GenS3{F0: uint8(v1), F1: -2147483648, F2: 109}
                                            switch "foo bar" {
                                            case "\n\t":
                                                var v50 uint = 701
                                                _ = v50
                                            case "foo bar":
                                                v44 -= 4294967295
                                            case "a":
                                                v44 -= 4294967295
                                            default:
                                                v51 := GenS1{F0: v44, F1: -1, F2: ("a" + (("a" + "\n\t") + "hello"))}
                                                _ = v51
                                            }
                                            switch (("\n\t" + ("hello" + "hello")) + "hello") {
                                            case "hello":
                                                v52 := GenS1{F0: (v43 & v43), F1: 2147483647, F2: ("a" + "a")}
                                                var v53 uint16 = 137
                                                v5 = true
                                                _ = v52
                                                _ = v53
                                            case "\n\t":
                                                v54 := GenS7{F0: (("hello" + ("" + ((("foo bar" + "hello") + "hello") + "foo bar"))) + ""), F1: (v34 ^ fi39)}
                                                _ = v54
                                            case "a":
                                                var v55 string = "\n\t"
                                                _ = v55
                                            default:
                                                v56 := GenS6{F0: v1}
                                                _ = v56
                                            }
                                            _ = v49
                                        }
                                    default:
                                        v57 := GenS1{F0: genFn1(100, 0), F1: 0, F2: "hello"}
                                        _ = v57
                                    }
                                    delete(v4, (v1 * (v1 << ((65535) & 63))))
                                }
                                switch fi42 {
                                case 4294967295:
                                    _ = len(v2)
                                case -9223372036854775808:
                                    var v58 uint = 65535
                                    _ = v58
                                }
                                v59 := v3[v1]
                                v43 = v43
                                _ = v59
                            default:
                                v44 = (^(v43))
                            }
                            var v60 string = "\n\t"
                            _ = fi42
                            _ = v43
                            _ = v44
                            _ = v60
                        }
                        if (("a" + "hello") > "hello") {
                            var v61 int16 = 245
                            switch "foo bar" {
                            case "":
                                _ = len(v3)
                            case "\n\t":
                                if (v5 || false) {
                                    fmt.Println(v1, v61, v61)
                                    v61 += 1000
                                } else {
                                    if v5 {
                                        fmt.Println(v1, v5)
                                    }
                                }
                            case "foo bar":
                                for fi62 := 0; fi62 < 2; fi62++ {
                                    switch (("hello" + "") + "foo bar") {
                                    case "\n\t":
                                        var v63 int8 = 1
                                        _ = v63
                                    case "hello":
                                        var v64 uint8 = 127
                                        _ = v64
                                    default:
                                        v5 = true
                                    }
                                    if ((fi39 / ((32767) | 1)) <= fi62) {
                                        v5 = false
                                        v34 = 2147483647
                                        fmt.Println(fi39, v61)
                                        _ = len(v3)
                                        fmt.Println(v1, v5)
                                        v65 := GenS8{F0: 780, F1: 0}
                                        v66 := GenS2{F0: ("hello" + ("a" + ("hello" + "a"))), F1: "a"}
                                        v34 += fi39
                                        delete(v6, fi62)
                                        fmt.Println(v34)
                                        var v67 uint64 = 0
                                        _ = v65
                                        _ = v66
                                        _ = v67
                                    } else {
                                        if v5 {
                                            v68 := GenS7{F0: ((("hello" + ("foo bar" + "a")) + ("\n\t" + "")) + ("a" + (("foo bar" + "foo bar") + "foo bar"))), F1: 694}
                                            _ = v68
                                        }
                                        switch fi62 {
                                        case -1:
                                            var v69 float32 = 3.14159
                                            _ = v69
                                        case 32768:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        default:
                                            v70 := GenS4{F0: (!v5)}
                                            _ = v70
                                        }
                                    }
                                    v3[1] = v1
                                    _ = fi62
                                }
                            default:
                                if v5 {
                                    if (255 > 202) {
                                        v34 = (fi39 + v34)
                                        v5 = v5
                                        var v71 int16 = (v61 & (v61 ^ v61))
                                        _ = v71
                                    }
                                    var v72 float64 = 0.0
                                    _ = v72
                                } else {
                                    if (v1 <= (v1 >> ((0) & 63))) {
                                        v73 := GenS4{F0: (42 < -663)}
                                        _ = v73
                                    } else {
                                        fmt.Println(fi39, v1)
                                        if v5 {
                                            if v5 {
                                                v5 = ((v61 != (v61 * v61)) && v5)
                                                fmt.Println(v61, v61)
                                                var v74 string = ((("foo bar" + ("a" + ("" + "hello"))) + ("hello" + "hello")) + "a")
                                                _ = v74
                                            }
                                            if len(v2) > 0 {
                                                v2[0] = ("" + "foo bar")
                                            }
                                            for fi75 := 0; fi75 < 3; fi75++ {
                                                v2 = append(v2, "hello")
                                                if (((1.5 < 1.5) && v5) || v5) {
                                                    var v76 int32 = 42
                                                    if v5 {
                                                        v34 = fi75
                                                        fmt.Println(v1, v34)
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    } else {
                                                        var v77 string = "foo bar"
                                                        _ = v77
                                                    }
                                                    v78 := GenS7{F0: "\n\t", F1: (^(fi75))}
                                                    _ = v76
                                                    _ = v78
                                                } else {
                                                    var v79 uint64 = 1
                                                    v5 = (false || (42 <= 9223372036854775807))
                                                    _ = v79
                                                }
                                                fmt.Println(fi39, v34, v5)
                                                _ = fi75
                                            }
                                        } else {
                                            delete(v3, 2147483647)
                                            if (false || ((2 >= 0) || (v34 < 9223372036854775807))) {
                                                var v80 float64 = 3.14159
                                                v80 += ((0.1 - v80) * (v80 - ((float64(1.5) + ((0.0 * v80) + 3.14159)) + 0.1)))
                                                v2 = append(v2, ("a" + "foo bar"))
                                                _ = v80
                                            } else {
                                                fmt.Println(v61)
                                                for fi81 := 0; fi81 < 2; fi81++ {
                                                    var v82 uint64 = 42
                                                    if len(v2) > 0 {
                                                        v2[0] = "a"
                                                    }
                                                    var v83 int16 = -1
                                                    _ = fi81
                                                    _ = v82
                                                    _ = v83
                                                }
                                                v84 := v3[v1]
                                                _ = v84
                                            }
                                        }
                                        v85 := GenS1{F0: 2, F1: 1, F2: "\n\t"}
                                        _ = v85
                                    }
                                    var v86 int64 = 9223372036854775807
                                    _ = v86
                                }
                            }
                            _ = v61
                        } else {
                            fmt.Println(fi39, v5, v34)
                        }
                    }
                    _ = fi39
                }
            case "":
                v87 := GenS5{F0: 2, F1: 1}
                _ = v87
            default:
                for fi88 := 0; fi88 < 15; fi88++ {
                    if v5 {
                        if v5 {
                            var v89 int32 = 1
                            _ = v89
                        } else {
                            for fi90 := 0; fi90 < 6; fi90++ {
                                v34 += (-(v34))
                                if false {
                                    if len(v2) > 0 {
                                        v2[0] = "a"
                                    }
                                    if v5 {
                                        if len(v2) > 0 {
                                            v2[0] = (((("hello" + "foo bar") + ("foo bar" + "\n\t")) + "\n\t") + "\n\t")
                                        }
                                        v5 = ((v5 && true) && v5)
                                        v5 = (-850 < 42)
                                    } else {
                                        switch (fi90 ^ (fi88 >> ((2) & 63))) {
                                        case 195:
                                            if ((0.0 != ((-2.25 * 1.5) * (3.14159 - -2.25))) || (!(v34 >= -507))) {
                                                _ = len(v3)
                                                if false {
                                                    var v91 string = ""
                                                    switch (^(fi88)) {
                                                    case 42:
                                                        v4[v1] = fi90
                                                    case 32767:
                                                        var v92 string = v91
                                                        _ = v92
                                                    }
                                                    v34 -= (^(v34))
                                                    _ = v91
                                                } else {
                                                    fmt.Println(fi88)
                                                    v93 := GenS9{F0: 42}
                                                    v34 += (fi88 * (fi88 - (^(fi88))))
                                                    _ = v93
                                                }
                                                var v94 float32 = 1e10
                                                _ = v94
                                            }
                                        case 32767:
                                            fmt.Println(v5)
                                        }
                                    }
                                } else {
                                    v2 = append(v2, "hello")
                                    var v95 float64 = (1.5 - 1.5)
                                    var v96 float64 = 3.14159
                                    _ = v95
                                    _ = v96
                                }
                                _ = fi90
                            }
                            v97 := GenS5{F0: 1, F1: 0}
                            _ = v97
                        }
                        v1 = v1
                        v2 = append(v2, "a")
                    }
                    switch fi88 {
                    case -664:
                        v34 = fi88
                    }
                    fmt.Println(v1, fi88)
                    _ = fi88
                }
            }
            var v98 float32 = 1.5
            _ = v98
        }
    case "\n\t":
        v99 := v3[0]
        _ = v99
    case "":
        v100 := GenS9{F0: -9223372036854775808}
        _ = v100
    default:
        v101 := GenS2{F0: "foo bar", F1: "hello"}
        _ = v101
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
