package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64; F1 string; F2 uint }
type GenS6 struct { F0 uint; F1 uint32; F2 bool }
type GenS7 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS8 struct { F0 uint8; F1 bool }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int8; F1 string }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int; F1 string }
type GenS13 struct { F0 uint16; F1 int16 }

func main() {
    v1 := []string{}
    v2 := []string{"a", (((("foo bar" + "foo bar") + "foo bar") + ((("a" + "\n\t") + ("a" + "\n\t")) + ("" + "a"))) + "foo bar")}
    var v3 string = "a"
    v4 := []string{("foo bar" + (v3 + "a")), v3}
    var v5 uint8 = 42
    var v6 uint16 = 255
    switch "\n\t" {
    case "a":
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
    default:
        var v7 int8 = 0
        _ = v7
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        if false {
            v5 -= v5
            v6 -= v6
            v5 += (v5 >> ((18446744073709551615) & 63))
        }
        _ = fi8
    }
    fmt.Println(v5, v3, v5)
    v9 := GenS1{F0: 4294967295, F1: 65535, F2: -164}
    v1 = append(v1, "\n\t")
    if len(v2) > 0 {
        v2[0] = v3
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        switch -1 {
        case -1:
            if false {
                v3 = v3
                var v11 string = v3
                _ = v11
            } else {
                for fi12 := 0; fi12 < 8; fi12++ {
                    switch (fi10 << ((678) & 63)) {
                    case 0:
                        fmt.Println(fi12, fi10, v5)
                    case 127:
                        v3 = (v3 + v3)
                        var v13 string = ("hello" + v3)
                        if true {
                            var v14 uint32 = 207
                            if false {
                                switch (v13 + v3) {
                                case "\n\t":
                                    v5 += v5
                                case "a":
                                    var v15 int8 = -1
                                    _ = v15
                                case "":
                                    var v16 int64 = 1
                                    _ = v16
                                }
                                v13 = ""
                                var v17 bool = (234 <= v5)
                                _ = v17
                            } else {
                                if (!(2147483647 == v14)) {
                                    v18 := GenS1{F0: 2, F1: 18446744073709551615, F2: 255}
                                    v19 := GenS1{F0: 0, F1: 0, F2: -9223372036854775808}
                                    _ = v18
                                    _ = v19
                                }
                                v13 = "hello"
                                if (!false) {
                                    if false {
                                        fmt.Println(v6, fi10, v14)
                                        var v20 uint16 = 469
                                        var v21 int64 = -9223372036854775808
                                        _ = v20
                                        _ = v21
                                    } else {
                                        if false {
                                            v22 := GenS1{F0: 675, F1: 819, F2: 9223372036854775807}
                                            if (v3 >= v13) {
                                                var v23 uint64 = 18446744073709551615
                                                fmt.Println(fi10)
                                                _ = v23
                                            }
                                            if (!false) {
                                                v24 := GenS2{F0: 0}
                                                _ = v24
                                            }
                                            v25 := GenS3{F0: 1, F1: 0}
                                            var v26 int8 = -128
                                            _ = v22
                                            _ = v25
                                            _ = v26
                                        } else {
                                            v13 = v3
                                            var v27 int32 = 42
                                            v28 := GenS3{F0: fi10, F1: 120}
                                            _ = v27
                                            _ = v28
                                        }
                                    }
                                } else {
                                    var v29 uint16 = (v6 >> ((0) & 63))
                                    _ = v29
                                }
                            }
                            if true {
                                v3 = "\n\t"
                            } else {
                                if false {
                                    switch v13 {
                                    case "hello":
                                        if true {
                                            if (!false) {
                                                v14 += v14
                                                v6 -= (v6 * (v6 & v6))
                                            }
                                            v30 := GenS4{F0: 547}
                                            var v31 int16 = 32767
                                            _ = v30
                                            _ = v31
                                        } else {
                                            if (0 > 9223372036854775807) {
                                                if (0 < 65535) {
                                                    switch v3 {
                                                    case "foo bar":
                                                        v5 -= (^(v5))
                                                    case "a":
                                                        var v32 int32 = int32(v6)
                                                        _ = v32
                                                    }
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    fmt.Println(fi10, v14, v3)
                                                    if (0 < 522) {
                                                        var v33 uint64 = 1
                                                        _ = v33
                                                    }
                                                    fmt.Println(v13, v14, v13)
                                                } else {
                                                    if (1 >= 35) {
                                                        var v34 int64 = -892
                                                        v35 := GenS5{F0: 4294967295, F1: v13, F2: 584}
                                                        if (((v13 + v3) + (("foo bar" + "foo bar") + v13)) <= "\n\t") {
                                                            v14 += v14
                                                            if false {
                                                                v4 = append(v4, (("hello" + v13) + v13))
                                                                if (!false) {
                                                                    v36 := GenS5{F0: 915, F1: v3, F2: 4294967295}
                                                                    if len(v1) > 0 {
                                                                        v1[0] = (((v13 + "hello") + ("a" + v13)) + "hello")
                                                                    }
                                                                    if (v3 > "foo bar") {
                                                                        var v37 uint8 = (v5 * (v5 | (v5 * (v5 + (v5 / (((^(v5))) | 1))))))
                                                                        if ((fi12 != fi10) && true) {
                                                                            var v38 uint8 = uint8(v14)
                                                                            _ = v38
                                                                        } else {
                                                                            v4 = append(v4, v3)
                                                                            var v39 int8 = 42
                                                                            _ = v39
                                                                        }
                                                                        switch v13 {
                                                                        case "":
                                                                            var v40 uint = 814
                                                                            _ = v40
                                                                        }
                                                                        v14 += 299
                                                                        var v41 uint16 = v6
                                                                        v34 = v34
                                                                        v42 := GenS5{F0: 18446744073709551615, F1: ("a" + "hello"), F2: 0}
                                                                        _ = v37
                                                                        _ = v41
                                                                        _ = v42
                                                                    } else {
                                                                        v34 += (v34 % ((-9223372036854775808) | 1))
                                                                        if (1 >= 989) {
                                                                            v13 = ((((((((("foo bar" + v3) + (((v3 + "a") + v3) + "foo bar")) + v13) + (v3 + ("\n\t" + v3))) + v13) + (((v13 + v3) + ((v13 + (v3 + (("hello" + "a") + v13))) + v13)) + (v13 + "hello"))) + (v3 + (v13 + (("foo bar" + (v3 + v13)) + (v3 + v3))))) + ((v13 + v13) + v13)) + v13)
                                                                        }
                                                                        v34 -= -1
                                                                    }
                                                                    _ = v36
                                                                } else {
                                                                    v43 := GenS6{F0: 65535, F1: (^(v14)), F2: false}
                                                                    fmt.Println(v14, v13)
                                                                    _ = v43
                                                                }
                                                                if (1e10 > ((0.1 / 1.0) - 0.1)) {
                                                                    v3 = v13
                                                                    var v44 uint16 = v6
                                                                    v14 -= (v14 - 839)
                                                                    _ = v44
                                                                }
                                                            }
                                                        } else {
                                                            v6 = (v6 ^ v6)
                                                            var v45 uint = 1
                                                            _ = v45
                                                        }
                                                        _ = v34
                                                        _ = v35
                                                    } else {
                                                        var v46 int64 = 42
                                                        var v47 int64 = -298
                                                        v1 = append(v1, v3)
                                                        _ = v46
                                                        _ = v47
                                                    }
                                                    var v48 uint32 = uint32(v5)
                                                    _ = v48
                                                }
                                                var v49 float32 = (0.1 * ((-2.25 + (3.14159 * 1e10)) / -2.25))
                                                _ = v49
                                            }
                                            var v50 uint8 = v5
                                            if ((-2.25 > 0.1) || (-1 <= -1)) {
                                                var v51 bool = true
                                                fmt.Println(v13, v13, v3)
                                                _ = len(v2)
                                                _ = v51
                                            } else {
                                                v52 := GenS7{F0: v6, F1: 758, F2: -573}
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                v3 = ("\n\t" + v13)
                                                _ = v52
                                            }
                                            _ = v50
                                        }
                                    case "foo bar":
                                        if false {
                                            if true {
                                                if (!(true && ((v5 >> ((65535) & 63)) >= (v5 << ((860) & 63))))) {
                                                    switch (v13 + (v13 + v3)) {
                                                    case "foo bar":
                                                        var v53 uint32 = v14
                                                        _ = v53
                                                    case "a":
                                                        v13 = v3
                                                    case "hello":
                                                        v54 := GenS8{F0: v5, F1: (((-1 >= -128) && true) && (!false))}
                                                        _ = v54
                                                    default:
                                                        if true {
                                                            v6 = 32767
                                                            v55 := GenS9{F0: 18446744073709551615}
                                                            var v56 float64 = -2.25
                                                            _ = v55
                                                            _ = v56
                                                        } else {
                                                            fmt.Println(v13, v14)
                                                            _ = len(v2)
                                                            var v57 int32 = 42
                                                            _ = v57
                                                        }
                                                    }
                                                    v14 += (v14 % (((^(v14))) | 1))
                                                }
                                                var v58 uint16 = 1
                                                _ = v58
                                            } else {
                                                v3 = ("a" + "hello")
                                            }
                                            var v59 int32 = 0
                                            fmt.Println(v3)
                                            v14 = (v14 + v14)
                                            if true {
                                                var v60 float64 = (1e10 * -2.25)
                                                v61 := GenS10{F0: 74, F1: ((((v3 + (v13 + v3)) + ((v13 + v3) + ("hello" + ("a" + (v3 + ""))))) + "hello") + v13)}
                                                v5 -= 210
                                                _ = v60
                                                _ = v61
                                            }
                                            switch ("foo bar" + v13) {
                                            case "hello":
                                                switch 1 {
                                                case 566:
                                                    if false {
                                                        _ = len(v1)
                                                    }
                                                }
                                            default:
                                                v62 := GenS6{F0: 18446744073709551615, F1: (v14 + (v14 % ((v14) | 1))), F2: false}
                                                _ = v62
                                            }
                                            _ = v59
                                        } else {
                                            v63 := GenS11{F0: v14}
                                            if true {
                                                switch "" {
                                                case "":
                                                    switch (fi10 << ((255) & 63)) {
                                                    case 402:
                                                        v3 = v13
                                                    case -170:
                                                        v3 = (v13 + (("a" + ("a" + "a")) + ("" + v13)))
                                                    }
                                                    v6 = v6
                                                    v3 = v3
                                                }
                                                var v64 int32 = 2147483647
                                                _ = v64
                                            }
                                            _ = v63
                                        }
                                    }
                                }
                                v65 := GenS2{F0: 909}
                                if (1 >= (v14 % ((v14) | 1))) {
                                    if len(v2) > 0 {
                                        v2[0] = v3
                                    }
                                    v14 -= uint32(v5)
                                }
                                var v66 int64 = -9223372036854775808
                                var v67 int8 = -40
                                _ = v65
                                _ = v66
                                _ = v67
                            }
                            _ = v14
                        } else {
                            switch (-(fi10)) {
                            case 105:
                                switch fi12 {
                                case 128:
                                    v13 = "hello"
                                    v13 = v13
                                    fmt.Println(fi10, v6)
                                default:
                                    v68 := GenS12{F0: -1, F1: "a"}
                                    _ = v68
                                }
                            default:
                                v6 += (v6 / (((v6 << ((329) & 63))) | 1))
                            }
                            var v69 uint8 = (v5 >> ((0) & 63))
                            if true {
                                v70 := GenS4{F0: 0}
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v70
                            } else {
                                fmt.Println(fi10)
                                if len(v4) > 0 {
                                    v4[0] = (v13 + "foo bar")
                                }
                            }
                            _ = v69
                        }
                        _ = v13
                    case 32767:
                        v2 = append(v2, "\n\t")
                    }
                    var v71 string = ""
                    _ = fi12
                    _ = v71
                }
                var v72 int8 = -107
                v5 += v5
                _ = v72
            }
        case -540:
            switch fi10 {
            case 2147483647:
                var v73 uint32 = 0
                _ = v73
            default:
                var v74 int16 = 1000
                _ = v74
            }
        case -2147483648:
            for fi75 := 0; fi75 < 5; fi75++ {
                if len(v4) > 0 {
                    v4[0] = v3
                }
                v76 := GenS6{F0: 1, F1: 0, F2: true}
                var v77 string = v3
                _ = fi75
                _ = v76
                _ = v77
            }
        default:
            fmt.Println(v6, v6, v6)
        }
        v78 := GenS10{F0: 62, F1: (("hello" + "hello") + v3)}
        v79 := GenS13{F0: uint16(v5), F1: 0}
        _ = fi10
        _ = v78
        _ = v79
    }
    fmt.Println(v6, v6, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
