package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int; F1 int16 }
type GenS7 struct { F0 int64; F1 uint32; F2 bool }
func genFn0(p0 string) uint32 {
    if false {
        v4 := GenS1{F0: 18446744073709551615, F1: 9223372036854775807}
        if (65535 < 2) {
            v5 := []int{-345, -9223372036854775808}
            _ = v5
        } else {
            var v6 int32 = -1
            var v7 string = ((("hello" + (p0 + "")) + p0) + ("foo bar" + "a"))
            _ = v6
            _ = v7
        }
        p0 = p0
        _ = v4
    }
    v8 := GenS1{F0: 4294967295, F1: 2147483647}
    p0 = p0
    p0 = "foo bar"
    p0 = ("hello" + "\n\t")
    _ = p0
    _ = v8
    return 0
}
func genFn1(p0 uint, p1 int) int64 {
    v9 := []int{(p1 >> ((981) & 63)), p1}
    fmt.Println(p0)
    for fi10 := 0; fi10 < 18; fi10++ {
        fmt.Println(fi10, p0)
        if len(v9) > 0 {
            v9[0] = p1
        }
        var v11 int32 = 42
        _ = fi10
        _ = v11
    }
    if len(v9) > 0 {
        v9[0] = (p1 & (p1 << (((^(p0))) & 63)))
    }
    _ = p0
    _ = p1
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 int = 42
    v2 := make(map[uint32]string)
    v3 := []uint32{}
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 bool = false
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        v15 := GenS2{F0: ((100 == int8(v1)) && (((((!(!(2147483647 < 324))) && false) || ("hello" <= "a")) || false) || true))}
        switch "hello" {
        case "\n\t":
            var v16 int16 = -1
            _ = v16
        case "a":
            v1 -= (fi14 & 32768)
        default:
            v1 += (fi14 * fi14)
        }
        _ = fi14
        _ = v15
    }
    v17 := GenS2{F0: (!(false && ((!(!false)) && (false || (-2.25 <= (0.1 + 3.14159))))))}
    v1 = -650
    for fi18 := 0; fi18 < 14; fi18++ {
        _ = len(v3)
        if (100 == 78) {
            switch "" {
            case "":
                fmt.Println(v1, v1)
            default:
                v19 := GenS2{F0: (genFn1(65535, 202) <= -209)}
                _ = v19
            }
        }
        if ("foo bar" > "a") {
            if (255 == 18446744073709551615) {
                if (!true) {
                    fmt.Println(fi18)
                }
                switch fi18 {
                case 402:
                    var v20 float64 = 1e10
                    _ = v20
                default:
                    for fi21 := 0; fi21 < 9; fi21++ {
                        var v22 uint16 = 228
                        var v23 float32 = (0.0 + float32(1e10))
                        var v24 uint = 606
                        _ = fi21
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                }
                var v25 float64 = (1e10 / -2.25)
                _ = v25
            } else {
                for fi26 := 0; fi26 < 14; fi26++ {
                    fmt.Println(v1, v1, v1)
                    fmt.Println(v1)
                    v3 = append(v3, 2147483647)
                    _ = fi26
                }
                v3 = append(v3, 2147483648)
            }
        }
        _ = fi18
    }
    var v27 uint = 827
    if ((v1 << ((v27) & 63)) == v1) {
        v2[2147483648] = ("" + "\n\t")
        v28 := GenS3{F0: 950}
        _ = len(v3)
        _ = v28
    } else {
        var v29 int = (v1 + v1)
        fmt.Println(v29, v27)
        _ = v29
    }
    var v30 float64 = 0.0
    switch "\n\t" {
    case "\n\t":
        for fi31 := 0; fi31 < 14; fi31++ {
            var v32 uint64 = 42
            var v33 string = "foo bar"
            _ = fi31
            _ = v32
            _ = v33
        }
    default:
        var v34 string = ("" + "")
        _ = v34
    }
    v30 += 0.0
    switch 815 {
    case -2147483648:
        if (!false) {
            if (v30 <= v30) {
                if (!false) {
                    v30 = 1.5
                    if (v27 == 255) {
                        var v35 uint = 18446744073709551615
                        switch "foo bar" {
                        case "foo bar":
                            var v36 uint64 = 31
                            _ = v36
                        case "\n\t":
                            fmt.Println(v1)
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        _ = v35
                    } else {
                        v37 := GenS4{F0: 161}
                        v1 = -9223372036854775808
                        if ((357 == 42) && (0 > -675)) {
                            switch "foo bar" {
                            case "":
                                switch ("hello" + "") {
                                case "hello":
                                    var v38 float64 = (((0.0 + float64((float32(v30) * 0.1))) / v30) * v30)
                                    _ = v38
                                case "\n\t":
                                    v39 := GenS1{F0: (v27 % ((1) | 1)), F1: v1}
                                    _ = v39
                                case "foo bar":
                                    fmt.Println(v1)
                                default:
                                    fmt.Println(v1)
                                }
                            case "a":
                                var v40 int = v1
                                _ = v40
                            case "\n\t":
                                if true {
                                    var v41 uint = 255
                                    if true {
                                        v42 := GenS3{F0: 42}
                                        v41 = v27
                                        var v43 uint = 65535
                                        _ = v42
                                        _ = v43
                                    } else {
                                        var v44 uint64 = 408
                                        _ = v44
                                    }
                                    _ = v41
                                } else {
                                    v45 := GenS1{F0: v27, F1: (v1 << (((v27 >> (((v27 ^ v27)) & 63))) & 63))}
                                    v1 += (v1 + (v1 * (v1 + v1)))
                                    _ = v45
                                }
                            }
                            switch 345 {
                            case -947:
                                var v46 int = v1
                                _ = v46
                            default:
                                v27 -= (v27 / (((^(v27))) | 1))
                            }
                        }
                        _ = v37
                    }
                }
                var v47 int16 = 1000
                _ = v47
            } else {
                v30 += 1.5
            }
            v48 := v2[2147483648]
            _ = v48
        } else {
            if (true || (true || (4294967295 == 777))) {
                v1 = 127
            } else {
                switch "hello" {
                case "\n\t":
                    fmt.Println(v30, v1, v1)
                case "foo bar":
                    v1 = v1
                default:
                    var v49 string = ("foo bar" + ((("" + "a") + "a") + ""))
                    _ = v49
                }
                if true {
                    if true {
                        for fi50 := 0; fi50 < 16; fi50++ {
                            var v51 uint = (v27 / ((v27) | 1))
                            _ = fi50
                            _ = v51
                        }
                    } else {
                        v2[0] = ("\n\t" + "foo bar")
                    }
                    v1 = v1
                    v52 := GenS5{F0: false}
                    _ = v52
                } else {
                    var v53 string = ""
                    var v54 int64 = 9223372036854775807
                    switch "" {
                    case "\n\t":
                        v55 := GenS4{F0: 142}
                        _ = v55
                    case "foo bar":
                        var v56 uint64 = 548
                        _ = v56
                    }
                    _ = v53
                    _ = v54
                }
                v27 = 852
            }
            var v57 int64 = -1
            _ = v57
        }
    case -9223372036854775808:
        switch 4294967295 {
        case 127:
            v27 = v27
        case 128:
            fmt.Println(v30, v1)
        case 32767:
            var v58 int64 = 0
            _ = v58
        default:
            switch "hello" {
            case "a":
                v1 = (^(v1))
            default:
                var v59 int16 = 42
                _ = v59
            }
        }
    }
    if (9223372036854775807 != (^(v1))) {
        if (142 != 0) {
            for fi60 := 0; fi60 < 12; fi60++ {
                var v61 uint32 = 2
                switch ("hello" + "a") {
                case "foo bar":
                    if false {
                        var v62 uint16 = 32767
                        var v63 int64 = -746
                        fmt.Println(v63, v27, v62)
                        _ = v62
                        _ = v63
                    } else {
                        v2[(v61 >> ((v27) & 63))] = "\n\t"
                        if false {
                            var v64 int8 = 100
                            if (979 == genFn0("a")) {
                                var v65 uint = v27
                                _ = v65
                            }
                            _ = v64
                        } else {
                            fmt.Println(v1)
                        }
                    }
                case "a":
                    v66 := GenS3{F0: 255}
                    _ = v66
                case "":
                    var v67 uint8 = 255
                    _ = v67
                default:
                    v3 = append(v3, (v61 + (v61 % ((v61) | 1))))
                }
                fmt.Println(v1)
                _ = fi60
                _ = v61
            }
        }
    } else {
        v2[723] = ((("foo bar" + "a") + (("\n\t" + "foo bar") + "hello")) + "\n\t")
        var v68 uint64 = 1
        _ = v68
    }
    delete(v2, genFn0((((("hello" + "") + ((("hello" + ("hello" + "hello")) + ((("\n\t" + "\n\t") + "a") + "hello")) + "")) + ("" + "")) + (("a" + "\n\t") + ""))))
    fmt.Println(v27, v30, v27)
    switch (v1 * v1) {
    case -870:
        switch ((("a" + ("\n\t" + "a")) + "\n\t") + (((("" + "foo bar") + "\n\t") + "\n\t") + "\n\t")) {
        case "\n\t":
            if false {
                v2[2] = "\n\t"
                if true {
                    fmt.Println(v30)
                    v69 := GenS2{F0: (892 >= 220)}
                    var v70 int32 = 739
                    _ = v69
                    _ = v70
                } else {
                    v71 := GenS4{F0: 42}
                    switch (("foo bar" + "a") + "") {
                    case "":
                        var v72 uint = (^(v27))
                        _ = v72
                    default:
                        switch "hello" {
                        case "":
                            v27 += 0
                        case "\n\t":
                            v30 = -2.25
                        case "a":
                            switch (-(v1)) {
                            case -9223372036854775808:
                                var v73 float64 = (0.1 - v30)
                                _ = v73
                            case 42:
                                fmt.Println(v27, v27, v30)
                            case -1:
                                var v74 float64 = v30
                                _ = v74
                            default:
                                for fi75 := 0; fi75 < 9; fi75++ {
                                    switch -2147483648 {
                                    case -169:
                                        var v76 bool = false
                                        _ = v76
                                    case 9223372036854775807:
                                        for fi77 := 0; fi77 < 12; fi77++ {
                                            v27 -= v27
                                            if false {
                                                v2[0] = (((("\n\t" + "") + "hello") + "hello") + (("foo bar" + "a") + ""))
                                            }
                                            switch "hello" {
                                            case "\n\t":
                                                v78 := GenS1{F0: v27, F1: -199}
                                                _ = v78
                                            case "":
                                                v79 := GenS6{F0: v1, F1: 1}
                                                _ = v79
                                            case "foo bar":
                                                fmt.Println(v27)
                                            }
                                            _ = fi77
                                        }
                                    default:
                                        v30 -= float64(1.5)
                                    }
                                    var v80 uint = (v27 & 4294967295)
                                    _ = fi75
                                    _ = v80
                                }
                            }
                        }
                    }
                    _ = v71
                }
            } else {
                for fi81 := 0; fi81 < 2; fi81++ {
                    fmt.Println(v1)
                    _ = fi81
                }
            }
        }
    case -2147483648:
        v1 += (^(v1))
    case -146:
        switch "foo bar" {
        case "hello":
            if true {
                v82 := v2[385]
                var v83 bool = (35 <= 42)
                var v84 int64 = 9223372036854775807
                _ = v82
                _ = v83
                _ = v84
            }
        case "":
            v85 := GenS7{F0: -149, F1: genFn0(("hello" + "foo bar")), F2: true}
            _ = v85
        case "foo bar":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    fmt.Println(v27, v27)
    switch (("a" + "") + ("a" + "\n\t")) {
    case "a":
        fmt.Println(v27, v1, v30)
    case "hello":
        v86 := GenS5{F0: false}
        _ = v86
    default:
        for fi87 := 0; fi87 < 7; fi87++ {
            for fi88 := 0; fi88 < 7; fi88++ {
                v89 := GenS5{F0: true}
                var v90 int8 = 42
                _ = fi88
                _ = v89
                _ = v90
            }
            for fi91 := 0; fi91 < 17; fi91++ {
                v2[65535] = ("hello" + "a")
                fmt.Println(v1, fi87, v30)
                _ = fi91
            }
            _ = fi87
        }
    }
    v1 -= (v1 ^ v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v30
}
