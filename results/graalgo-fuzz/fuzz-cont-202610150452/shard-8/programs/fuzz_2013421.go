package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int64 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 int16; F2 uint64 }
type GenS5 struct { F0 int32; F1 uint8 }
type GenS6 struct { F0 string; F1 uint8; F2 uint64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int16; F1 int32; F2 bool }
type GenS9 struct { F0 bool; F1 bool }
type GenS10 struct { F0 uint16; F1 int8; F2 uint }
func genFn0() int8 {
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := []int{-830, fi6, fi6, -392}
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 128, F1: (true || (!((false && ((true && (42 >= 368)) || ((0.1 >= 0.0) && (212 < 0)))) && true))), F2: -9223372036854775808}
    var v9 int64 = 21
    var v10 uint32 = 2
    var v11 uint16 = 32767
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 71
}

func main() {
    v1 := []int{}
    v2 := []uint32{4294967295, 65535, 2147483647, 2147483648}
    v3 := make(map[string]string)
    v4 := []string{}
    var v5 int32 = -548
    for fi12 := 0; fi12 < 17; fi12++ {
        for fi13 := 0; fi13 < 1; fi13++ {
            if true {
                var v14 uint64 = 4294967295
                fmt.Println(fi12)
                _ = v14
            } else {
                switch ("foo bar" + ((("hello" + "") + "a") + "\n\t")) {
                case "\n\t":
                    v5 = v5
                case "a":
                    var v15 uint16 = 1
                    _ = v15
                case "foo bar":
                    var v16 uint16 = 589
                    _ = v16
                default:
                    fmt.Println(fi13, fi13)
                    v17 := GenS1{F0: 0, F1: false, F2: -9223372036854775808}
                    v18 := GenS1{F0: 2, F1: ("a" < "a"), F2: 755}
                    _ = v17
                    _ = v18
                }
                if ((4294967295 <= 403) && false) {
                    switch (^(fi13)) {
                    case -2147483648:
                        var v19 float32 = (-2.25 / (3.14159 / 0.1))
                        _ = v19
                    default:
                        if (42 <= 128) {
                            if (1 > 42) {
                                var v20 int = fi13
                                switch "" {
                                case "\n\t":
                                    v21 := GenS2{F0: 42, F1: 4294967295}
                                    _ = v21
                                case "":
                                    var v22 uint32 = 2147483648
                                    _ = v22
                                case "a":
                                    v2 = append(v2, 866)
                                }
                                v20 += fi12
                                var v23 int64 = -9223372036854775808
                                _ = v20
                                _ = v23
                            } else {
                                var v24 uint32 = 2
                                _ = v24
                            }
                        }
                    }
                } else {
                    v5 += (v5 >> ((1) & 63))
                    v5 -= (v5 | (v5 << ((42) & 63)))
                    fmt.Println(v5)
                }
                switch fi12 {
                case 127:
                    v25 := GenS2{F0: 1, F1: 1}
                    _ = v25
                case -943:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case 2147483647:
                    var v26 int32 = -2147483648
                    if true {
                        v5 -= v26
                        v26 += 0
                        if true {
                            var v27 float32 = 1e10
                            if true {
                                if len(v1) > 0 {
                                    v1[0] = 32768
                                }
                            } else {
                                v28 := v3["foo bar"]
                                _ = v28
                            }
                            v29 := GenS3{F0: 42}
                            _ = v27
                            _ = v29
                        }
                    }
                    v26 = (v26 + -2147483648)
                    _ = v26
                default:
                    if true {
                        if (true || (429 >= 4294967295)) {
                            v5 -= (v5 * -33)
                            var v30 int = fi12
                            _ = v30
                        }
                    } else {
                        v31 := GenS1{F0: 0, F1: true, F2: 0}
                        if (!true) {
                            v32 := v3["\n\t"]
                            switch ((v32 + (((("hello" + (("hello" + ("foo bar" + ((v32 + "foo bar") + (("\n\t" + v32) + v32)))) + v32)) + "hello") + v32) + "hello")) + "\n\t") {
                            case "hello":
                                var v33 float32 = float32(0.0)
                                if false {
                                    if (!(235 >= 222)) {
                                        v5 = v5
                                        v33 = 0.0
                                        v32 = (v32 + v32)
                                    } else {
                                        fmt.Println(fi12, fi13)
                                        if (1 >= 42) {
                                            var v34 int = (fi12 >> ((65535) & 63))
                                            switch 778 {
                                            case 128:
                                                var v35 int64 = -470
                                                _ = v35
                                            }
                                            _ = v34
                                        }
                                        v2 = append(v2, 543)
                                    }
                                    switch ("hello" + v32) {
                                    case "foo bar":
                                        v36 := GenS3{F0: 18446744073709551615}
                                        _ = v36
                                    case "a":
                                        fmt.Println(v5, v5)
                                    default:
                                        fmt.Println(v33)
                                    }
                                    v4 = append(v4, "")
                                } else {
                                    fmt.Println(fi12, v33)
                                    v37 := GenS1{F0: 127, F1: (v32 >= "foo bar"), F2: 0}
                                    var v38 uint16 = 255
                                    if (9 >= 824) {
                                        if len(v1) > 0 {
                                            v1[0] = (fi12 ^ (fi12 << ((0) & 63)))
                                        }
                                        v5 -= v5
                                    }
                                    _ = v37
                                    _ = v38
                                }
                                fmt.Println(fi12)
                                _ = v33
                            case "foo bar":
                                var v39 uint16 = 913
                                _ = v39
                            case "":
                                v1 = append(v1, (^(fi12)))
                            default:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                            if true {
                                switch v32 {
                                case "a":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                case "":
                                    var v40 uint32 = 432
                                    _ = v40
                                }
                            }
                            var v41 uint = 536
                            v32 = ((v32 + v32) + (("hello" + v32) + v32))
                            _ = v32
                            _ = v41
                        } else {
                            var v42 int = (fi12 + fi13)
                            var v43 string = ("hello" + "\n\t")
                            _ = v42
                            _ = v43
                        }
                        var v44 uint = 255
                        _ = v31
                        _ = v44
                    }
                }
            }
            _ = fi13
        }
        switch 32767 {
        case -803:
            var v45 int32 = -518
            _ = v45
        case 4294967295:
            if (false && ((3.14159 >= 3.14159) || true)) {
                for fi46 := 0; fi46 < 12; fi46++ {
                    if (1 != 781) {
                        v5 = (v5 << ((42) & 63))
                    }
                    if (0 > 1) {
                        fmt.Println(fi46, fi12)
                        var v47 uint32 = 65535
                        if false {
                            var v48 uint8 = 204
                            var v49 int64 = 19
                            var v50 int8 = -128
                            var v51 float32 = (1e10 - (float32(1.5) / (3.14159 / 1.5)))
                            if ((!(!true)) || true) {
                                fmt.Println(fi46, v50)
                            } else {
                                var v52 int64 = v49
                                v50 = v50
                                _ = len(v3)
                                _ = v52
                            }
                            _ = v48
                            _ = v49
                            _ = v50
                            _ = v51
                        } else {
                            switch "" {
                            case "":
                                fmt.Println(fi12)
                            case "\n\t":
                                v47 += (^(v47))
                            default:
                                v3["a"] = ("" + "\n\t")
                            }
                            v47 -= (v47 >> ((958) & 63))
                            fmt.Println(v5)
                            var v53 int = fi46
                            v54 := GenS4{F0: 0, F1: 42, F2: 772}
                            _ = v53
                            _ = v54
                        }
                        _ = v47
                    }
                    _ = fi46
                }
                switch "a" {
                case "a":
                    v5 = v5
                }
                v55 := GenS5{F0: (v5 >> ((0) & 63)), F1: 20}
                _ = v55
            } else {
                v1 = append(v1, 32767)
                _ = len(v3)
                var v56 uint8 = 128
                _ = v56
            }
        }
        _ = fi12
    }
    v57 := GenS6{F0: "a", F1: 159, F2: 2}
    v3["\n\t"] = ("" + "\n\t")
    switch 127 {
    case 2147483647:
        for fi58 := 0; fi58 < 5; fi58++ {
            if false {
                switch "\n\t" {
                case "hello":
                    switch "hello" {
                    case "hello":
                        var v59 bool = (true || false)
                        _ = v59
                    case "a":
                        var v60 int64 = 521
                        _ = v60
                    default:
                        var v61 uint8 = 127
                        _ = v61
                    }
                case "foo bar":
                    v5 += (v5 ^ 0)
                case "a":
                    v62 := GenS4{F0: 0, F1: 467, F2: 383}
                    _ = v62
                }
                fmt.Println(v5, v5)
                var v63 uint = 767
                _ = v63
            }
            v64 := GenS7{F0: "foo bar"}
            var v65 uint8 = 2
            _ = fi58
            _ = v64
            _ = v65
        }
    default:
        v5 += (v5 + -1)
    }
    for fi66 := 0; fi66 < 20; fi66++ {
        v5 += (v5 + (-(v5)))
        var v67 int = fi66
        var v68 uint = 255
        _ = fi66
        _ = v67
        _ = v68
    }
    fmt.Println(v5)
    var v69 int8 = -121
    var v70 int = -271
    var v71 bool = (-1 >= 1)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v72 := GenS4{F0: -32768, F1: int16(v70), F2: 0}
    var v73 string = ""
    v74 := v3["hello"]
    for fi75 := 0; fi75 < 9; fi75++ {
        for fi76 := 0; fi76 < 5; fi76++ {
            v3[v74] = (("a" + v74) + (v74 + v74))
            v77 := GenS5{F0: v5, F1: 2}
            var v78 uint16 = 1
            v79 := GenS8{F0: 0, F1: (v5 ^ (v5 * (v5 - v5))), F2: false}
            var v80 string = "foo bar"
            _ = fi76
            _ = v77
            _ = v78
            _ = v79
            _ = v80
        }
        for fi81 := 0; fi81 < 9; fi81++ {
            switch v73 {
            case "hello":
                v82 := GenS9{F0: (!(193 >= 128)), F1: v71}
                switch ("\n\t" + "foo bar") {
                case "foo bar":
                    v5 -= v5
                case "hello":
                    if (1000 >= 42) {
                        v70 = 0
                        v83 := GenS3{F0: 636}
                        v84 := GenS1{F0: 0, F1: (4294967295 > 264), F2: 0}
                        if len(v2) > 0 {
                            v2[0] = 2
                        }
                        var v85 int32 = 511
                        _ = v83
                        _ = v84
                        _ = v85
                    }
                case "":
                    v86 := GenS10{F0: 42, F1: (v69 | (v69 * genFn0())), F2: 1}
                    _ = v86
                default:
                    switch v70 {
                    case -635:
                        _ = len(v3)
                    default:
                        v87 := GenS8{F0: 32767, F1: (v5 - (-(v5))), F2: ((!false) || v71)}
                        _ = v87
                    }
                }
                _ = v82
            }
            if (1.5 > ((float64(1.5) + 0.1) / (float64(1.5) - 1e10))) {
                v69 = (v69 | genFn0())
                _ = len(v3)
                v70 = fi81
            } else {
                var v88 float64 = -2.25
                _ = v88
            }
            _ = fi81
        }
        _ = fi75
    }
    v4 = append(v4, v74)
    v5 = 42
    fmt.Println(v69, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v69
    _ = v70
    _ = v71
    _ = v72
    _ = v73
    _ = v74
}
