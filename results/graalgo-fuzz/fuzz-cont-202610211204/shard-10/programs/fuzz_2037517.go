package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 uint64; F1 int }
type GenS7 struct { F0 string; F1 uint64 }
type GenS8 struct { F0 bool }
func genFn0(p0 uint16, p1 uint16, p2 int) uint64 {
    var v4 float32 = float32(1e10)
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := []int{p2, p2}
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}

func main() {
    v1 := []string{}
    var v2 uint64 = 18446744073709551615
    var v3 uint = 369
    for fi7 := 0; fi7 < 3; fi7++ {
        v2 -= (v2 & v2)
        var v8 uint = 4294967295
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 941}
    if true {
        for fi10 := 0; fi10 < 4; fi10++ {
            var v11 float64 = (1.5 - -2.25)
            for fi12 := 0; fi12 < 12; fi12++ {
                v1 = append(v1, (((("foo bar" + "") + "hello") + ("" + (("" + ("a" + "")) + "\n\t"))) + ""))
                var v13 int8 = int8(v3)
                v11 -= (1e10 + v11)
                fmt.Println(fi12)
                _ = fi12
                _ = v13
            }
            v3 = (v3 / (((v3 + (^(v3)))) | 1))
            _ = fi10
            _ = v11
        }
    }
    if (!(!(-2147483648 <= -911))) {
        for fi14 := 0; fi14 < 4; fi14++ {
            v15 := GenS1{F0: 32767}
            v16 := GenS2{F0: fi14}
            fmt.Println(v2, fi14)
            _ = fi14
            _ = v15
            _ = v16
        }
    } else {
        switch -796 {
        case 127:
            if true {
                var v17 bool = ("hello" < "a")
                v18 := GenS3{F0: (-2.25 == 0.0)}
                _ = v17
                _ = v18
            } else {
                if ((523 < 834) || true) {
                    var v19 float64 = 0.1
                    var v20 bool = false
                    for fi21 := 0; fi21 < 18; fi21++ {
                        v19 = 1e10
                        for fi22 := 0; fi22 < 14; fi22++ {
                            var v23 float32 = 0.1
                            fmt.Println(v2, v2, v23)
                            if v20 {
                                v24 := GenS4{F0: uint(v2), F1: -2147483648}
                                if ((!((("foo bar" + "foo bar") + "a") == "")) && true) {
                                    _ = len(v1)
                                    v19 -= v19
                                    var v25 float32 = (-2.25 - 0.0)
                                    _ = v25
                                } else {
                                    var v26 uint8 = 42
                                    fmt.Println(v23, v3, fi21)
                                    _ = v26
                                }
                                v2 += 50
                                var v27 uint8 = 255
                                var v28 uint8 = v27
                                _ = v24
                                _ = v27
                                _ = v28
                            }
                            _ = fi22
                            _ = v23
                        }
                        _ = fi21
                    }
                    _ = v19
                    _ = v20
                }
                if len(v1) > 0 {
                    v1[0] = ("hello" + "\n\t")
                }
            }
        case 1:
            v2 = (^(v2))
        default:
            v29 := GenS5{F0: 42, F1: 0}
            _ = v29
        }
        for fi30 := 0; fi30 < 17; fi30++ {
            for fi31 := 0; fi31 < 2; fi31++ {
                switch fi31 {
                case 127:
                    var v32 int = fi31
                    _ = v32
                case 128:
                    var v33 int = (fi31 & fi30)
                    _ = v33
                default:
                    var v34 int8 = -128
                    _ = v34
                }
                v35 := GenS3{F0: (0 >= fi30)}
                v36 := GenS6{F0: 1, F1: (-(fi31))}
                v37 := GenS7{F0: "", F1: (v2 * 661)}
                var v38 bool = true
                v38 = (1e10 < 1.5)
                var v39 float64 = (0.0 / 1e10)
                v2 -= v2
                _ = fi31
                _ = v35
                _ = v36
                _ = v37
                _ = v38
                _ = v39
            }
            v1 = append(v1, "")
            switch ("a" + ("" + "")) {
            case "foo bar":
                if false {
                    for fi40 := 0; fi40 < 7; fi40++ {
                        var v41 int8 = 127
                        v41 += (v41 & v41)
                        _ = fi40
                        _ = v41
                    }
                    for fi42 := 0; fi42 < 7; fi42++ {
                        fmt.Println(v2, fi42, fi42)
                        v3 = uint(v2)
                        v2 += genFn0(255, 66, 127)
                        v1 = append(v1, "foo bar")
                        switch "hello" {
                        case "a":
                            var v43 bool = true
                            _ = v43
                        case "":
                            var v44 uint = 925
                            _ = v44
                        case "hello":
                            switch "\n\t" {
                            case "a":
                                v3 += 2
                            case "":
                                v45 := GenS8{F0: (false || (-2.25 >= 0.1))}
                                _ = v45
                            default:
                                v2 = 389
                            }
                        }
                        _ = fi42
                    }
                    v46 := GenS5{F0: 1, F1: 0}
                    _ = v46
                }
            case "":
                if true {
                    for fi47 := 0; fi47 < 7; fi47++ {
                        v2 = genFn0(255, 429, 728)
                        var v48 int8 = -11
                        if len(v1) > 0 {
                            v1[0] = "foo bar"
                        }
                        _ = fi47
                        _ = v48
                    }
                }
            default:
                for fi49 := 0; fi49 < 18; fi49++ {
                    if false {
                        if false {
                            if (true || true) {
                                v3 -= v3
                                var v50 uint = 42
                                fmt.Println(v50, v2, fi49)
                                _ = v50
                            } else {
                                v3 += (v3 / ((937) | 1))
                                var v51 uint = (^(v3))
                                if (false || true) {
                                    v52 := GenS8{F0: true}
                                    var v53 uint32 = 42
                                    _ = v52
                                    _ = v53
                                }
                                _ = v51
                            }
                            v1 = append(v1, ("hello" + "a"))
                            v54 := GenS5{F0: uint16(v2), F1: -9223372036854775808}
                            _ = v54
                        }
                        var v55 string = "foo bar"
                        v1 = append(v1, ((v55 + "") + v55))
                        _ = v55
                    } else {
                        if (!true) {
                            if true {
                                if (!(true && (788 != (^(v3))))) {
                                    if (!false) {
                                        if false {
                                            v2 = v2
                                            var v56 int64 = 9223372036854775807
                                            _ = v56
                                        } else {
                                            fmt.Println(fi30, fi30)
                                            fmt.Println(v3)
                                        }
                                    }
                                    fmt.Println(fi30, fi30, fi49)
                                    switch (("" + "") + "\n\t") {
                                    case "hello":
                                        var v57 int32 = -808
                                        _ = v57
                                    case "":
                                        v58 := GenS2{F0: fi30}
                                        _ = v58
                                    case "foo bar":
                                        var v59 uint8 = 69
                                        _ = v59
                                    default:
                                        var v60 uint8 = 128
                                        _ = v60
                                    }
                                }
                                var v61 uint32 = 2147483648
                                v3 = 162
                                _ = v61
                            } else {
                                var v62 uint8 = 1
                                _ = v62
                            }
                        }
                    }
                    _ = fi49
                }
            }
            _ = fi30
        }
    }
    fmt.Println(v2, v3)
    v2 += v2
    var v63 int8 = -2
    var v64 uint16 = 1
    if ((v63 & (^(v63))) < (v63 % ((v63) | 1))) {
        var v65 int = -2147483648
        v63 = v63
        for fi66 := 0; fi66 < 11; fi66++ {
            if (!((false || true) && (1 > 9223372036854775807))) {
                v65 = 1
            }
            if (!(((false || false) || (false && ((-2.25 == (((1.5 - 1.5) / 1.0) - ((float64((1.5 - -2.25)) * 0.0) - 0.1))) && (v2 >= genFn0(v64, (v64 >> ((65535) & 63)), (^(fi66))))))) && false)) {
                if false {
                    v63 = (v63 & (v63 + v63))
                    v2 = (v2 | (v2 - 931))
                } else {
                    var v67 float64 = (0.0 - 0.1)
                    _ = v67
                }
                v2 = v2
            }
            _ = fi66
        }
        _ = v65
    }
    switch "" {
    case "hello":
        var v68 float32 = 0.0
        _ = v68
    default:
        v69 := GenS3{F0: (32767 != 32767)}
        _ = v69
    }
    if (-1 < 42) {
        var v70 uint16 = uint16(v2)
        _ = v70
    }
    v64 += v64
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v63
    _ = v64
    _ = v9
}
