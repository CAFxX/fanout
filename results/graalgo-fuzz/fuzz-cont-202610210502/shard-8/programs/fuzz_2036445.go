package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 int64; F2 bool }
type GenS5 struct { F0 int8; F1 uint; F2 int32 }
func genFn0() string {
    var v7 float64 = 0.1
    var v8 string = "foo bar"
    v9 := GenS1{F0: 0, F1: (v8 + "")}
    v7 = v7
    _ = v7
    _ = v8
    _ = v9
    return "foo bar"
}
func genFn1(p0 string, p1 bool, p2 string) int32 {
    v10 := []int{-2147483648, -649, -666}
    if true {
        for fi11 := 0; fi11 < 15; fi11++ {
            switch 525 {
            case -9223372036854775808:
                v12 := GenS2{F0: 1}
                _ = v12
            case 684:
                p1 = p1
            default:
                var v13 int8 = 42
                _ = v13
            }
            for fi14 := 0; fi14 < 16; fi14++ {
                var v15 uint8 = 18
                var v16 uint32 = 2147483647
                _ = fi14
                _ = v15
                _ = v16
            }
            fmt.Println(p2)
            _ = fi11
        }
        var v17 int16 = -32768
        v10 = append(v10, 0)
        _ = v17
    } else {
        var v18 int16 = 42
        var v19 string = ((p0 + (("\n\t" + ("a" + "a")) + p0)) + "\n\t")
        var v20 uint64 = 1
        _ = v18
        _ = v19
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return -529
}

func main() {
    var v1 int = -261
    var v2 bool = false
    v3 := []string{("foo bar" + "hello"), "foo bar", "hello"}
    var v4 float64 = 1.5
    var v5 float64 = v4
    var v6 string = "\n\t"
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 string = (((v6 + (("a" + v6) + ((v6 + ((v6 + (("hello" + (("\n\t" + ("foo bar" + (("" + v6) + (v6 + ("\n\t" + v6))))) + v6)) + (v6 + v6))) + (v6 + "\n\t"))) + "foo bar"))) + v6) + ((v6 + v6) + v6))
        if (false || false) {
            switch (v6 + v22) {
            case "a":
                if (!v2) {
                    v5 += (v4 * (-2.25 * (1e10 + float64(-2.25))))
                    if true {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        for fi23 := 0; fi23 < 5; fi23++ {
                            switch (fi23 - (fi21 - fi23)) {
                            case 2147483647:
                                var v24 uint8 = 2
                                _ = v24
                            case 127:
                                v22 = ((v6 + v22) + (("foo bar" + v6) + ("\n\t" + v6)))
                            }
                            _ = fi23
                        }
                        if v2 {
                            var v25 uint64 = 813
                            for fi26 := 0; fi26 < 14; fi26++ {
                                fmt.Println(v6)
                                switch ((v22 + v22) + (v22 + v6)) {
                                case "":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                default:
                                    v27 := GenS3{F0: v25}
                                    _ = v27
                                }
                                _ = fi26
                            }
                            switch 32767 {
                            case 805:
                                v3 = append(v3, "foo bar")
                            case -1:
                                if false {
                                    for fi28 := 0; fi28 < 15; fi28++ {
                                        v6 = v6
                                        var v29 uint32 = 2147483648
                                        v3 = append(v3, "")
                                        _ = fi28
                                        _ = v29
                                    }
                                    fmt.Println(v1)
                                    fmt.Println(v6, fi21, v25)
                                }
                            case -425:
                                for fi30 := 0; fi30 < 8; fi30++ {
                                    v25 = 560
                                    _ = fi30
                                }
                            default:
                                for fi31 := 0; fi31 < 15; fi31++ {
                                    if (!(0 == 863)) {
                                        v32 := GenS1{F0: -125, F1: "foo bar"}
                                        var v33 int = (fi31 * fi21)
                                        _ = v32
                                        _ = v33
                                    } else {
                                        var v34 uint16 = 65535
                                        v35 := GenS1{F0: -1, F1: "foo bar"}
                                        if ((v2 || v2) || false) {
                                            if false {
                                                var v36 float64 = v5
                                                v37 := GenS2{F0: 838}
                                                v38 := GenS1{F0: -1, F1: (("hello" + ("" + v22)) + (v22 + "hello"))}
                                                _ = v36
                                                _ = v37
                                                _ = v38
                                            } else {
                                                switch (fi31 >> ((2) & 63)) {
                                                case 889:
                                                    var v39 uint = 2
                                                    _ = v39
                                                default:
                                                    var v40 uint64 = v25
                                                    _ = v40
                                                }
                                            }
                                            if false {
                                                v41 := GenS1{F0: -672, F1: v22}
                                                var v42 uint8 = 0
                                                fmt.Println(v34)
                                                _ = v41
                                                _ = v42
                                            } else {
                                                v4 = float64(-2.25)
                                                var v43 float64 = (0.1 - v4)
                                                if (!v2) {
                                                    v25 += (v25 | 63)
                                                    if len(v3) > 0 {
                                                        v3[0] = v6
                                                    }
                                                    var v44 uint = uint(v25)
                                                    _ = v44
                                                }
                                                var v45 int64 = -152
                                                _ = v43
                                                _ = v45
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = v22
                                            }
                                            var v46 string = (v22 + (v22 + (v22 + v22)))
                                            if false {
                                                var v47 int = fi21
                                                v5 -= v4
                                                _ = v47
                                            }
                                            _ = v46
                                        }
                                        _ = v34
                                        _ = v35
                                    }
                                    v48 := GenS2{F0: genFn1(v22, (!(v2 && (v2 || false))), v6)}
                                    _ = fi31
                                    _ = v48
                                }
                            }
                            _ = v25
                        }
                    } else {
                        var v49 uint = 566
                        v4 = v4
                        for fi50 := 0; fi50 < 2; fi50++ {
                            v51 := GenS4{F0: "hello", F1: 9223372036854775807, F2: v2}
                            if len(v3) > 0 {
                                v3[0] = (v6 + v6)
                            }
                            _ = fi50
                            _ = v51
                        }
                        _ = v49
                    }
                    v52 := GenS2{F0: -707}
                    _ = v52
                }
            case "foo bar":
                fmt.Println(v22, v2)
            }
            v4 -= 0.0
            for fi53 := 0; fi53 < 16; fi53++ {
                v5 = (v4 * v4)
                _ = fi53
            }
        } else {
            v3 = append(v3, v6)
            switch (v1 & v1) {
            case 128:
                if true {
                    var v54 string = ((((v22 + v22) + (v6 + "hello")) + (v22 + "")) + "hello")
                    var v55 string = v22
                    var v56 uint64 = 18446744073709551615
                    _ = v54
                    _ = v55
                    _ = v56
                }
            case 9223372036854775807:
                if v2 {
                    var v57 int32 = 1
                    _ = v57
                }
            }
        }
        v6 = "foo bar"
        _ = fi21
        _ = v22
    }
    v2 = (v2 || ((93 == 128) && (((false && (true || ((2 > 98) || (-128 >= -128)))) && (!v2)) && v2)))
    v5 = (0.0 * v4)
    var v58 float32 = -2.25
    fmt.Println(v5, v2)
    v59 := GenS4{F0: "\n\t", F1: -9223372036854775808, F2: v2}
    for fi60 := 0; fi60 < 13; fi60++ {
        if len(v3) > 0 {
            v3[0] = v6
        }
        for fi61 := 0; fi61 < 16; fi61++ {
            fmt.Println(v2, v6, v2)
            v4 += ((v5 * v4) - 0.1)
            _ = fi61
        }
        for fi62 := 0; fi62 < 13; fi62++ {
            if (!v2) {
                v63 := GenS5{F0: 42, F1: 42, F2: 1}
                switch (v6 + (v6 + (v6 + v6))) {
                case "":
                    var v64 int16 = 608
                    _ = v64
                case "hello":
                    if false {
                        v1 += (v1 - (fi60 + (fi62 - fi60)))
                    }
                case "foo bar":
                    fmt.Println(v4)
                default:
                    if (!false) {
                        if (0.0 > v4) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = len(v3)
                            v65 := GenS2{F0: 339}
                            _ = v65
                        } else {
                            v2 = (-2147483648 < -1)
                        }
                        var v66 uint16 = 385
                        v6 = "\n\t"
                        fmt.Println(fi60, fi60, fi62)
                        var v67 int64 = 0
                        _ = v66
                        _ = v67
                    }
                }
                var v68 string = v6
                _ = v63
                _ = v68
            } else {
                switch (fi60 ^ 32768) {
                case 42:
                    fmt.Println(fi60)
                default:
                    fmt.Println(fi60, v1)
                }
            }
            _ = fi62
        }
        _ = fi60
    }
    v69 := GenS2{F0: -1}
    v3 = append(v3, "foo bar")
    switch ("a" + v6) {
    case "a":
        var v70 float64 = 1.5
        _ = v70
    default:
        v58 += 3.14159
    }
    if v2 {
        for fi71 := 0; fi71 < 12; fi71++ {
            v72 := GenS4{F0: v6, F1: -1, F2: false}
            for fi73 := 0; fi73 < 8; fi73++ {
                var v74 bool = (v58 < (v58 / (v58 + (v58 + (((1e10 + v58) + -2.25) * v58)))))
                v2 = (v58 >= v58)
                fmt.Println(v2, fi71, fi71)
                _ = fi73
                _ = v74
            }
            _ = fi71
            _ = v72
        }
        var v75 uint32 = 542
        switch (v1 - v1) {
        case -825:
            switch (v1 / (((-(v1))) | 1)) {
            case -9223372036854775808:
                var v76 int8 = -128
                _ = v76
            case 1:
                v2 = v2
            }
        case -2147483648:
            fmt.Println(v5, v5)
        default:
            if v2 {
                var v77 int16 = -32768
                _ = len(v3)
                for fi78 := 0; fi78 < 19; fi78++ {
                    if len(v3) > 0 {
                        v3[0] = v6
                    }
                    fmt.Println(v4, v1, v58)
                    switch ((v6 + v6) + "a") {
                    case "hello":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    _ = fi78
                }
                _ = v77
            }
        }
        _ = v75
    } else {
        var v79 float32 = (v58 * 0.0)
        _ = v79
    }
    var v80 int8 = -128
    switch (-(v1)) {
    case 9223372036854775807:
        for fi81 := 0; fi81 < 18; fi81++ {
            v58 = (-2.25 * (float32(v5) * 0.1))
            _ = fi81
        }
    case 32767:
        v82 := GenS4{F0: v6, F1: -76, F2: (-90 == -32768)}
        _ = v82
    case -9223372036854775808:
        v83 := GenS1{F0: -1, F1: (v6 + "a")}
        _ = v83
    }
    if v2 {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v69
    _ = v80
}
