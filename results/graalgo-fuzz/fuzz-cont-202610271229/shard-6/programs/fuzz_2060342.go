package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int8 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 int32; F1 int; F2 int16 }
type GenS5 struct { F0 int8; F1 uint64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int64; F1 int64; F2 int64 }
type GenS9 struct { F0 int16; F1 int; F2 uint64 }
type GenS10 struct { F0 uint64; F1 uint8 }
type GenS11 struct { F0 bool; F1 bool; F2 uint64 }
type GenS12 struct { F0 int32; F1 string; F2 int8 }
func genFn0(p0 uint32, p1 int16) uint8 {
    if false {
        p0 += (p0 >> ((817) & 63))
        p0 -= 2147483647
    }
    for fi4 := 0; fi4 < 6; fi4++ {
        var v5 uint8 = 255
        fmt.Println(fi4, fi4)
        _ = fi4
        _ = v5
    }
    switch ("" + "") {
    case "":
        v6 := GenS1{F0: 42, F1: -128, F2: -128}
        _ = v6
    case "hello":
        v7 := []int{-2147483648, 1, 1}
        _ = v7
    case "\n\t":
        v8 := GenS2{F0: (^(p0)), F1: 998}
        _ = v8
    }
    _ = p0
    _ = p1
    return 141
}
func genFn1() uint32 {
    var v9 int8 = 1
    v9 += v9
    switch "foo bar" {
    case "":
        v10 := GenS1{F0: 950, F1: v9, F2: v9}
        _ = v10
    case "\n\t":
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 string = "foo bar"
            _ = fi11
            _ = v12
        }
    }
    _ = v9
    return 65535
}

func main() {
    var v1 int16 = -1
    v2 := []int{-2147483648, 4294967295, -9223372036854775808}
    v3 := []string{"\n\t", "foo bar"}
    v13 := GenS1{F0: 1, F1: 116, F2: 42}
    if len(v2) > 0 {
        v2[0] = 32767
    }
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    fmt.Println(v1)
    switch "a" {
    case "foo bar":
        v14 := GenS3{F0: 2147483647, F1: (v1 - (v1 ^ 0))}
        _ = v14
    case "hello":
        switch -1 {
        case 811:
            for fi15 := 0; fi15 < 7; fi15++ {
                if true {
                    fmt.Println(v1)
                } else {
                    fmt.Println(fi15, fi15)
                }
                fmt.Println(v1)
                _ = fi15
            }
        }
    default:
        for fi16 := 0; fi16 < 4; fi16++ {
            switch 1 {
            case 128:
                var v17 int16 = (v1 | 32767)
                _ = v17
            case 32767:
                switch "a" {
                case "\n\t":
                    var v18 int8 = 1
                    _ = v18
                case "foo bar":
                    v1 -= v1
                case "a":
                    for fi19 := 0; fi19 < 5; fi19++ {
                        if false {
                            switch "" {
                            case "foo bar":
                                switch 1 {
                                case 128:
                                    if false {
                                        var v20 bool = true
                                        v20 = false
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = v20
                                    }
                                default:
                                    v1 = (-(v1))
                                }
                            default:
                                v1 = (v1 / (((v1 * -32768)) | 1))
                            }
                            if false {
                                var v21 uint = 255
                                v22 := GenS4{F0: -1, F1: fi19, F2: int16(fi19)}
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v23 uint16 = 42
                                if len(v3) > 0 {
                                    v3[0] = "\n\t"
                                }
                                _ = v21
                                _ = v22
                                _ = v23
                            }
                            switch (("\n\t" + ("" + ("a" + ""))) + "hello") {
                            case "a":
                                var v24 float32 = 1.5
                                _ = v24
                            case "":
                                var v25 uint16 = 65535
                                _ = v25
                            }
                        }
                        var v26 string = ("hello" + "")
                        _ = fi19
                        _ = v26
                    }
                default:
                    v1 = 0
                }
            default:
                v1 = 1000
            }
            switch 1 {
            case 32767:
                for fi27 := 0; fi27 < 11; fi27++ {
                    switch fi27 {
                    case 413:
                        var v28 uint8 = 2
                        _ = v28
                    case 1:
                        var v29 uint64 = 119
                        _ = v29
                    case 128:
                        if true {
                            switch ("a" + ((("" + ("\n\t" + "\n\t")) + "\n\t") + "a")) {
                            case "a":
                                if false {
                                    switch "\n\t" {
                                    case "":
                                        if false {
                                            if true {
                                                switch ("hello" + ("\n\t" + "hello")) {
                                                case "hello":
                                                    fmt.Println(fi27)
                                                case "a":
                                                    var v30 uint32 = genFn1()
                                                    _ = v30
                                                default:
                                                    var v31 uint = 285
                                                    _ = v31
                                                }
                                                var v32 uint64 = 468
                                                var v33 int16 = v1
                                                _ = v32
                                                _ = v33
                                            }
                                        }
                                    case "a":
                                        if true {
                                            fmt.Println(fi16, fi27, v1)
                                        } else {
                                            if (false && false) {
                                                fmt.Println(v1, v1, fi27)
                                                var v34 uint64 = 867
                                                _ = v34
                                            } else {
                                                if true {
                                                    v35 := GenS5{F0: 1, F1: 655}
                                                    _ = v35
                                                } else {
                                                    var v36 uint8 = genFn0(genFn1(), v1)
                                                    if false {
                                                        var v37 bool = (!((-2.25 - (float32(3.14159) * -2.25)) >= -2.25))
                                                        v37 = true
                                                        v37 = (!(true || (956 >= 2147483647)))
                                                        fmt.Println(v36, fi16, fi27)
                                                        if false {
                                                            var v38 float32 = (0.0 - 0.1)
                                                            _ = v38
                                                        }
                                                        _ = v37
                                                    }
                                                    _ = len(v2)
                                                    _ = v36
                                                }
                                                if (float64(-2.25) == 1.5) {
                                                    v1 = v1
                                                    fmt.Println(fi16, v1, fi27)
                                                    switch (((("\n\t" + "foo bar") + "hello") + ("\n\t" + ("a" + "hello"))) + "\n\t") {
                                                    case "a":
                                                        fmt.Println(fi16, fi16)
                                                    case "":
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    default:
                                                        fmt.Println(fi16, fi16)
                                                        if false {
                                                            var v39 uint32 = 65535
                                                            var v40 int64 = 642
                                                            _ = v39
                                                            _ = v40
                                                        }
                                                        v1 = -454
                                                    }
                                                } else {
                                                    fmt.Println(v1, fi16)
                                                    var v41 int32 = -2147483648
                                                    if false {
                                                        if false {
                                                            v41 = v41
                                                            if (false || (-278 < -9223372036854775808)) {
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    v41 = 686
                                                                case "":
                                                                    var v42 int32 = -1
                                                                    _ = v42
                                                                }
                                                                if (!false) {
                                                                    _ = len(v2)
                                                                    v41 = v41
                                                                }
                                                                v2 = append(v2, 32768)
                                                            }
                                                            var v43 uint = 18446744073709551615
                                                            v44 := GenS6{F0: 0}
                                                            _ = v43
                                                            _ = v44
                                                        } else {
                                                            v45 := GenS7{F0: ("" + "hello")}
                                                            _ = v45
                                                        }
                                                        var v46 float32 = (-2.25 / -2.25)
                                                        var v47 uint32 = 42
                                                        _ = v46
                                                        _ = v47
                                                    }
                                                    _ = v41
                                                }
                                                var v48 int8 = -1
                                                _ = v48
                                            }
                                        }
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v49 uint16 = 32767
                                    if (4294967295 > 0) {
                                        fmt.Println(v1)
                                        if ((v49 & (v49 | v49)) > (^(v49))) {
                                            if len(v3) > 0 {
                                                v3[0] = ("foo bar" + ("hello" + "hello"))
                                            }
                                        } else {
                                            v49 -= 65535
                                            switch (fi16 - (fi16 >> ((948) & 63))) {
                                            case 128:
                                                v50 := GenS8{F0: 100, F1: 997, F2: 0}
                                                fmt.Println(fi16, v49)
                                                v51 := GenS9{F0: v1, F1: (fi16 << ((254) & 63)), F2: 29}
                                                _ = v50
                                                _ = v51
                                            case -586:
                                                fmt.Println(v49, v1, fi27)
                                            case 42:
                                                switch "" {
                                                case "a":
                                                    fmt.Println(fi16, v49)
                                                    if false {
                                                        if (-663 == -2147483648) {
                                                            fmt.Println(fi27, v1, v1)
                                                        } else {
                                                            var v52 int32 = 2147483647
                                                            var v53 uint16 = v49
                                                            _ = v52
                                                            _ = v53
                                                        }
                                                    } else {
                                                        v49 = 420
                                                    }
                                                    fmt.Println(fi27, fi16)
                                                    if (1 != 127) {
                                                        fmt.Println(v49)
                                                        v1 -= (^(v1))
                                                        v3 = append(v3, "hello")
                                                    }
                                                    var v54 string = "foo bar"
                                                    _ = v54
                                                case "foo bar":
                                                    v49 = v49
                                                }
                                            }
                                            var v55 int16 = (v1 | v1)
                                            fmt.Println(v55, v49, v1)
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            _ = v55
                                        }
                                        fmt.Println(v49)
                                        if (true || ((^(v1)) != 987)) {
                                            if len(v2) > 0 {
                                                v2[0] = (fi27 % ((fi16) | 1))
                                            }
                                        } else {
                                            if (2 == 0) {
                                                var v56 uint8 = 190
                                                var v57 int64 = 42
                                                _ = v56
                                                _ = v57
                                            }
                                            switch (fi16 - (fi16 | (fi27 / ((32767) | 1)))) {
                                            case 4294967295:
                                                if (false || true) {
                                                    v49 += v49
                                                }
                                            default:
                                                if (713 <= 649) {
                                                    var v58 float32 = 0.0
                                                    switch (("" + ((("a" + "a") + "foo bar") + "foo bar")) + "\n\t") {
                                                    case "a":
                                                        _ = len(v3)
                                                    case "hello":
                                                        switch "a" {
                                                        case "a":
                                                            var v59 uint8 = 142
                                                            _ = v59
                                                        }
                                                    case "":
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    default:
                                                        if len(v2) > 0 {
                                                            v2[0] = fi16
                                                        }
                                                    }
                                                    fmt.Println(fi16)
                                                    _ = v58
                                                }
                                            }
                                        }
                                    }
                                    var v60 int16 = -1
                                    v61 := GenS3{F0: -17, F1: v60}
                                    v60 += int16(fi27)
                                    _ = v49
                                    _ = v60
                                    _ = v61
                                } else {
                                    var v62 float64 = ((1e10 / 3.14159) / 1.0)
                                    if ((v62 * 1e10) <= v62) {
                                        var v63 float32 = (((0.1 * float32(((float64(1.5) - 0.0) / v62))) - -2.25) + -2.25)
                                        _ = v63
                                    } else {
                                        v62 = float64(0.0)
                                        v62 += float64(float32(v62))
                                        v62 -= v62
                                    }
                                    var v64 int16 = (v1 / (((-(v1))) | 1))
                                    if false {
                                        v1 = v1
                                        var v65 string = (("" + "hello") + "hello")
                                        _ = v65
                                    } else {
                                        v62 = (v62 * float64((3.14159 * 1e10)))
                                        v62 = (0.1 + v62)
                                    }
                                    _ = v62
                                    _ = v64
                                }
                            case "foo bar":
                                v66 := GenS8{F0: 0, F1: 9223372036854775807, F2: 9223372036854775807}
                                _ = v66
                            case "hello":
                                v1 = (-(v1))
                                var v67 uint64 = 0
                                switch (-(fi16)) {
                                case -340:
                                    var v68 int64 = 811
                                    _ = v68
                                }
                                v69 := GenS10{F0: v67, F1: 0}
                                _ = v67
                                _ = v69
                            }
                            v2 = append(v2, (-(fi16)))
                        }
                    default:
                        v1 = v1
                    }
                    switch "" {
                    case "":
                        switch (fi16 * -2147483648) {
                        case 32767:
                            var v70 float32 = 1.5
                            _ = v70
                        }
                    case "\n\t":
                        v3 = append(v3, "hello")
                    case "foo bar":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    switch "a" {
                    case "foo bar":
                        v3 = append(v3, "a")
                    case "\n\t":
                        var v71 uint32 = 678
                        _ = v71
                    case "":
                        var v72 string = "foo bar"
                        _ = v72
                    }
                    _ = fi27
                }
            case -1:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case -863:
                v73 := GenS11{F0: (fi16 <= fi16), F1: (((true && ((1e10 * -2.25) > ((-2.25 * 1.5) - 1.5))) || (!true)) && ((!false) || false)), F2: 18446744073709551615}
                _ = v73
            default:
                v1 = (v1 / ((831) | 1))
            }
            _ = fi16
        }
    }
    var v74 bool = (0 == -2147483648)
    v75 := GenS8{F0: -1, F1: 9223372036854775807, F2: -1}
    if (v74 || true) {
        v3 = append(v3, ("a" + "a"))
    } else {
        v1 -= (-(v1))
        fmt.Println(v1, v1)
    }
    switch 2147483647 {
    case 127:
        v1 += (v1 << ((759) & 63))
    case 2147483647:
        v3 = append(v3, "")
    }
    v1 += 1000
    v76 := GenS12{F0: 303, F1: "hello", F2: 42}
    _ = len(v3)
    for fi77 := 0; fi77 < 4; fi77++ {
        fmt.Println(v1, fi77, v1)
        _ = fi77
    }
    var v78 int16 = (v1 * v1)
    if ("a" < "") {
        v74 = v74
        var v79 uint = 4294967295
        _ = v79
    } else {
        fmt.Println(v78, v78, v74)
        v80 := GenS1{F0: 0, F1: 100, F2: 127}
        _ = v80
    }
    for fi81 := 0; fi81 < 6; fi81++ {
        v82 := GenS12{F0: -2147483648, F1: "foo bar", F2: 100}
        _ = fi81
        _ = v82
    }
    var v83 string = ("\n\t" + "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v74:", v74)
    fmt.Println("v78:", v78)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v74
    _ = v75
    _ = v76
    _ = v78
    _ = v83
}
