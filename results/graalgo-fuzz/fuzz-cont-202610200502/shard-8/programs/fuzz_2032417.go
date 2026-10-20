package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int8; F1 bool; F2 uint8 }
type GenS8 struct { F0 uint64; F1 int }
func genFn0(p0 int8, p1 uint32, p2 uint8) uint8 {
    if true {
        p1 = 2147483647
        p0 = (p0 << ((962) & 63))
        var v6 uint = 4294967295
        _ = v6
    } else {
        var v7 string = ("" + "foo bar")
        _ = v7
    }
    var v8 float64 = float64((1e10 + -2.25))
    var v9 uint8 = (p2 * 255)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 28
}
func genFn1(p0 int16, p1 int64, p2 int8) int16 {
    switch "" {
    case "foo bar":
        var v10 int64 = 864
        _ = v10
    case "a":
        p0 = 0
    case "\n\t":
        v11 := GenS1{F0: (p1 - p1), F1: false}
        _ = v11
    default:
        v12 := []int{-9223372036854775808, 459, -1}
        _ = v12
    }
    if (463 > 2147483648) {
        p1 -= 9223372036854775807
        v13 := []int{745, 4294967295, 127}
        _ = v13
    }
    if true {
        var v14 uint64 = 820
        _ = v14
    } else {
        var v15 string = "hello"
        v15 = v15
        switch 128 {
        case 32767:
            var v16 bool = (((true || true) || (!false)) || (!false))
            _ = v16
        case 9223372036854775807:
            p1 -= p1
        case 326:
            if true {
                var v17 int = -9223372036854775808
                v18 := []int{(v17 - int(p2)), (v17 & v17), (v17 << ((759) & 63))}
                if len(v18) > 0 {
                    v18[0] = v17
                }
                _ = v17
                _ = v18
            }
        default:
            p0 = p0
        }
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    return (p0 >> ((175) & 63))
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{("\n\t" + "foo bar"), ""}
    var v3 int32 = 2147483647
    var v4 bool = false
    var v5 uint16 = 32767
    for fi19 := 0; fi19 < 17; fi19++ {
        v20 := GenS1{F0: -1, F1: (933 <= 0)}
        var v21 float64 = 1.5
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        for fi23 := 0; fi23 < 1; fi23++ {
            v24 := v1[("a" + ((("a" + "hello") + ("\n\t" + ("" + "hello"))) + ""))]
            var v25 uint64 = 4294967295
            v3 = v3
            _ = fi23
            _ = v24
            _ = v25
        }
        if (v4 && (((1000 != -32768) && v4) && v4)) {
            if (!false) {
                switch "" {
                case "hello":
                    v26 := GenS2{F0: 2147483648, F1: ""}
                    _ = v26
                case "":
                    var v27 int32 = v3
                    _ = v27
                case "foo bar":
                    fmt.Println(v3)
                default:
                    v28 := GenS3{F0: 2147483648, F1: 65535}
                    _ = v28
                }
                v3 = 1
            }
            if (3.14159 >= -2.25) {
                var v29 uint32 = 42
                fmt.Println(v3, fi22)
                _ = len(v1)
                _ = v29
            } else {
                v4 = false
                if ((!((((!v4) && v4) && ((1000 > -32768) && true)) || v4)) && (((-(v3)) != (v3 >> ((2) & 63))) || (((fi22 & 1) <= (fi22 >> ((2) & 63))) && v4))) {
                    if false {
                        var v30 uint16 = v5
                        _ = v30
                    } else {
                        v2 = append(v2, "")
                        switch fi22 {
                        case 32767:
                            if v4 {
                                v1[("hello" + "\n\t")] = (v4 && v4)
                            }
                        case -874:
                            for fi31 := 0; fi31 < 15; fi31++ {
                                _ = len(v1)
                                v1["a"] = (!((v5 << ((275) & 63)) != v5))
                                var v32 float32 = (0.1 + (-2.25 / (3.14159 - float32(1.5))))
                                _ = fi31
                                _ = v32
                            }
                        default:
                            if (!(("a" < ("a" + "a")) && (!true))) {
                                switch fi22 {
                                case 9223372036854775807:
                                    v33 := GenS4{F0: 1, F1: 2}
                                    _ = v33
                                case 680:
                                    var v34 uint8 = genFn0(1, 2147483648, genFn0(-128, 2, 42))
                                    _ = v34
                                default:
                                    v3 -= v3
                                }
                            }
                        }
                        for fi35 := 0; fi35 < 3; fi35++ {
                            if v4 {
                                if (-128 != 1) {
                                    v36 := GenS1{F0: 656, F1: v4}
                                    delete(v1, (("a" + (("a" + "foo bar") + "a")) + ("a" + ((("foo bar" + "hello") + "foo bar") + "\n\t"))))
                                    v4 = (fi22 > (fi35 << ((0) & 63)))
                                    _ = v36
                                } else {
                                    var v37 int32 = (v3 | v3)
                                    var v38 bool = (65535 == 625)
                                    var v39 uint8 = 17
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                }
                            } else {
                                _ = len(v2)
                            }
                            v5 = (^(v5))
                            _ = fi35
                        }
                    }
                    var v40 uint8 = 117
                    _ = v40
                }
            }
            var v41 int16 = -1
            _ = v41
        }
        for fi42 := 0; fi42 < 3; fi42++ {
            if (v3 <= (v3 % (((v3 * v3)) | 1))) {
                switch ((("" + "a") + "\n\t") + "") {
                case "foo bar":
                    if len(v2) > 0 {
                        v2[0] = "hello"
                    }
                }
                switch fi42 {
                case 32768:
                    v43 := GenS5{F0: -769, F1: 18446744073709551615}
                    _ = v43
                case -549:
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                case -707:
                    if (!(-128 <= 127)) {
                        v44 := GenS4{F0: (-(fi42)), F1: 2}
                        _ = v44
                    } else {
                        fmt.Println(v3, v5)
                        if (100 >= 127) {
                            v45 := v1["\n\t"]
                            fmt.Println(fi22, v5, v3)
                            v5 = (v5 + v5)
                            _ = v45
                        }
                        var v46 int32 = (-(v3))
                        _ = v46
                    }
                }
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
            } else {
                var v47 float64 = 0.1
                if (331 <= 18446744073709551615) {
                    v4 = v4
                    v47 = v47
                    v4 = (!true)
                    var v48 uint32 = 2147483648
                    var v49 float32 = float32(v47)
                    v2 = append(v2, (("a" + "hello") + "\n\t"))
                    var v50 uint64 = 66
                    _ = v48
                    _ = v49
                    _ = v50
                }
                _ = v47
            }
            var v51 bool = false
            if (v4 && v4) {
                if (!true) {
                    v52 := GenS6{F0: (v3 % (((v3 >> ((65535) & 63))) | 1))}
                    switch "a" {
                    case "a":
                        v1["\n\t"] = v4
                    default:
                        v2 = append(v2, "foo bar")
                    }
                    v3 -= v3
                    _ = v52
                } else {
                    var v53 uint32 = 2147483648
                    _ = v53
                }
                v51 = false
                v5 += (v5 + v5)
            }
            _ = fi42
            _ = v51
        }
        _ = fi22
    }
    var v54 string = "hello"
    fmt.Println(v4)
    if (v4 && (false && true)) {
        v5 += v5
        var v55 uint = 18446744073709551615
        _ = v55
    } else {
        for fi56 := 0; fi56 < 16; fi56++ {
            v54 = v54
            if true {
                v3 = v3
                if true {
                    if true {
                        v5 += (^(v5))
                        if v4 {
                            if false {
                                if (((0.0 >= 0.1) || (!v4)) && (!(713 <= 659))) {
                                    fmt.Println(v5)
                                }
                            } else {
                                for fi57 := 0; fi57 < 15; fi57++ {
                                    v58 := GenS7{F0: 127, F1: ((fi56 >> ((0) & 63)) < fi57), F2: 245}
                                    var v59 bool = ((1 == int(v5)) && true)
                                    var v60 uint = 255
                                    _ = fi57
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                }
                                delete(v1, v54)
                                var v61 float32 = float32(1e10)
                                _ = v61
                            }
                            v54 = v54
                            v3 = -2147483648
                        }
                    } else {
                        for fi62 := 0; fi62 < 13; fi62++ {
                            var v63 uint8 = 255
                            v3 += (v3 << ((65535) & 63))
                            var v64 uint64 = 2
                            _ = fi62
                            _ = v63
                            _ = v64
                        }
                    }
                    var v65 int64 = 9223372036854775807
                    _ = v65
                }
                var v66 int16 = 42
                _ = v66
            }
            _ = fi56
        }
    }
    var v67 float32 = 1.5
    v67 = float32(float64(1e10))
    for fi68 := 0; fi68 < 6; fi68++ {
        v4 = false
        v4 = v4
        v54 = (((v54 + "a") + "") + "")
        _ = fi68
    }
    var v69 uint32 = 2147483647
    v70 := GenS8{F0: 2, F1: 128}
    var v71 uint64 = 2
    v5 -= v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v67:", v67)
    fmt.Println("v69:", v69)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v67
    _ = v69
    _ = v70
    _ = v71
}
