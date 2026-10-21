package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 bool; F1 int32; F2 int64 }
type GenS4 struct { F0 bool; F1 uint64; F2 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64; F1 bool; F2 int32 }
type GenS7 struct { F0 int32; F1 uint8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int8 }
func genFn0(p0 uint32, p1 int32, p2 bool) int64 {
    p2 = (true && false)
    fmt.Println(p0, p2)
    var v5 uint32 = (p0 * p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 266
}
func genFn1() uint {
    if (2147483647 >= 0) {
        var v6 float32 = 1.5
        if true {
            if true {
                v7 := GenS1{F0: 149}
                _ = v7
            } else {
                switch ("\n\t" + "a") {
                case "foo bar":
                    v6 += v6
                case "a":
                    v6 = v6
                }
            }
            v6 = 1.5
            fmt.Println(v6)
        } else {
            if (1.5 <= 3.14159) {
                switch ("foo bar" + "a") {
                case "a":
                    if (true && (!false)) {
                        var v8 uint64 = 0
                        if (!false) {
                            var v9 int16 = 774
                            v10 := []int{32767, 794}
                            _ = v10
                            _ = v9
                        } else {
                            v6 -= ((1e10 - (float32(1.5) - v6)) - v6)
                            for fi11 := 0; fi11 < 3; fi11++ {
                                for fi12 := 0; fi12 < 2; fi12++ {
                                    switch ("a" + "") {
                                    case "":
                                        v13 := GenS1{F0: 0}
                                        _ = v13
                                    case "a":
                                        v14 := []int{}
                                        _ = v14
                                    }
                                    fmt.Println(fi11, fi12)
                                    switch "hello" {
                                    case "":
                                        var v15 int64 = -644
                                        _ = v15
                                    case "hello":
                                        v16 := []int{(fi11 & fi12)}
                                        _ = v16
                                    default:
                                        fmt.Println(fi11)
                                    }
                                    var v17 int8 = 127
                                    v18 := []int{42, (fi11 % ((32768) | 1)), (fi12 % ((fi12) | 1)), 9223372036854775807}
                                    switch "\n\t" {
                                    case "a":
                                        switch 0 {
                                        case 981:
                                            if len(v18) > 0 {
                                                v18[0] = (fi11 / (((fi11 + int(v8))) | 1))
                                            }
                                        default:
                                            v19 := GenS1{F0: -2147483648}
                                            _ = v19
                                        }
                                    case "foo bar":
                                        v18 = append(v18, fi11)
                                    case "\n\t":
                                        var v20 uint64 = v8
                                        _ = v20
                                    default:
                                        switch (fi11 << ((uint(v8)) & 63)) {
                                        case -1:
                                            switch (fi11 % ((fi11) | 1)) {
                                            case 4294967295:
                                                var v21 int = (fi12 & (fi12 >> ((515) & 63)))
                                                _ = v21
                                            case -602:
                                                if true {
                                                    v17 += (v17 >> ((7) & 63))
                                                    fmt.Println(fi12)
                                                    v17 = v17
                                                    if len(v18) > 0 {
                                                        v18[0] = 182
                                                    }
                                                    fmt.Println(v8, fi11, fi11)
                                                } else {
                                                    var v22 uint64 = (^(v8))
                                                    fmt.Println(v8, v22)
                                                    v23 := GenS2{F0: 1, F1: 128}
                                                    v6 = v6
                                                    v8 = v8
                                                    _ = v22
                                                    _ = v23
                                                }
                                            default:
                                                v8 = (^(v8))
                                            }
                                        }
                                    }
                                    _ = fi12
                                    _ = v17
                                    _ = v18
                                }
                                v24 := []int{32768, fi11, (fi11 ^ 128), 127}
                                _ = fi11
                                _ = v24
                            }
                            if (false || true) {
                                fmt.Println(v6, v6)
                            } else {
                                v25 := GenS3{F0: (true || false), F1: 42, F2: -1}
                                v8 -= 42
                                var v26 int64 = -1
                                _ = v25
                                _ = v26
                            }
                        }
                        for fi27 := 0; fi27 < 11; fi27++ {
                            switch fi27 {
                            case 2147483647:
                                for fi28 := 0; fi28 < 2; fi28++ {
                                    fmt.Println(v6)
                                    v29 := []int{(fi28 | (^(fi28))), fi28, (fi27 * 32768), (fi27 % (((-(fi28))) | 1))}
                                    v6 = v6
                                    _ = fi28
                                    _ = v29
                                }
                            default:
                                var v30 string = (("\n\t" + "hello") + "")
                                _ = v30
                            }
                            _ = fi27
                        }
                        _ = v8
                    }
                case "hello":
                    var v31 uint8 = 127
                    _ = v31
                }
                var v32 string = ""
                _ = v32
            } else {
                v33 := GenS1{F0: -75}
                v34 := GenS3{F0: false, F1: 44, F2: 0}
                _ = v33
                _ = v34
            }
            v6 += 1e10
        }
        _ = v6
    }
    for fi35 := 0; fi35 < 8; fi35++ {
        var v36 uint = 18446744073709551615
        _ = fi35
        _ = v36
    }
    return 4294967295
}

func main() {
    var v1 int32 = 1
    v2 := []string{"", "\n\t"}
    v3 := []int{9223372036854775807, -9223372036854775808, int(v1)}
    var v4 uint32 = 723
    v1 += v1
    v4 = 1
    var v37 uint64 = 18446744073709551615
    if false {
        if false {
            v4 = 2
        }
        v4 += 274
        if (32768 < 42) {
            v4 += (^(v4))
            fmt.Println(v4)
            if false {
                var v38 bool = true
                var v39 int = 2147483647
                _ = v38
                _ = v39
            }
        } else {
            v40 := GenS1{F0: (^(v1))}
            v4 = (v4 - (^(v4)))
            var v41 float32 = (3.14159 - -2.25)
            _ = v40
            _ = v41
        }
    }
    for fi42 := 0; fi42 < 14; fi42++ {
        switch (^(fi42)) {
        case 128:
            v4 += v4
        case 32768:
            var v43 uint16 = uint16(v4)
            _ = v43
        case 2147483647:
            if true {
                v44 := GenS2{F0: 1, F1: 127}
                for fi45 := 0; fi45 < 4; fi45++ {
                    if false {
                        var v46 float64 = (float64((3.14159 + -2.25)) + (0.1 - (((1e10 - (0.0 - 1.5)) / -2.25) + -2.25)))
                        if len(v3) > 0 {
                            v3[0] = 750
                        }
                        _ = v46
                    }
                    _ = fi45
                }
                _ = v44
            }
        }
        _ = fi42
    }
    var v47 int16 = -1
    var v48 uint8 = 221
    v49 := GenS2{F0: 10, F1: 127}
    if false {
        if (!(!true)) {
            v50 := GenS4{F0: (1 >= 21), F1: (v37 ^ (v37 % (((v37 + (v37 | 2))) | 1))), F2: (("" + ("" + ("a" + "a"))) + "hello")}
            _ = v50
        } else {
            for fi51 := 0; fi51 < 10; fi51++ {
                v4 = (^(v4))
                v2 = append(v2, "\n\t")
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi51
            }
        }
        v48 = v48
    } else {
        fmt.Println(v48, v37, v37)
        switch "hello" {
        case "":
            v52 := GenS4{F0: false, F1: v37, F2: "\n\t"}
            _ = v52
        case "hello":
            var v53 int16 = (^(v47))
            _ = v53
        case "foo bar":
            v47 -= 235
        default:
            if (!(v1 >= 1)) {
                v1 -= 2147483647
                var v54 int16 = v47
                var v55 string = "hello"
                _ = v54
                _ = v55
            }
        }
        v1 = (v1 + v1)
    }
    if (215 < -637) {
        for fi56 := 0; fi56 < 6; fi56++ {
            v48 += 0
            if len(v3) > 0 {
                v3[0] = fi56
            }
            _ = len(v2)
            _ = fi56
        }
        if (!(!true)) {
            v57 := GenS5{F0: v4}
            _ = v57
        } else {
            for fi58 := 0; fi58 < 10; fi58++ {
                v48 += (v48 + v48)
                switch "hello" {
                case "foo bar":
                    v3 = append(v3, 42)
                case "":
                    var v59 int64 = -9223372036854775808
                    _ = v59
                case "a":
                    v60 := GenS6{F0: v37, F1: false, F2: (^(v1))}
                    _ = v60
                }
                _ = len(v3)
                _ = fi58
            }
            if (!(float32(3.14159) < (0.1 + 0.0))) {
                fmt.Println(v1, v48)
            }
        }
        for fi61 := 0; fi61 < 7; fi61++ {
            for fi62 := 0; fi62 < 8; fi62++ {
                var v63 bool = true
                var v64 string = "foo bar"
                _ = fi62
                _ = v63
                _ = v64
            }
            switch (fi61 + (-(fi61))) {
            case 127:
                v48 += uint8(v4)
            case -974:
                v4 -= (^(v4))
            case 1:
                for fi65 := 0; fi65 < 6; fi65++ {
                    switch int(v47) {
                    case 0:
                        var v66 int8 = -128
                        _ = v66
                    case 128:
                        switch fi61 {
                        case 127:
                            v67 := GenS7{F0: (v1 & v1), F1: uint8(v37)}
                            v68 := GenS3{F0: ((!true) && true), F1: (-(v1)), F2: -1}
                            _ = v67
                            _ = v68
                        case 521:
                            v4 += 340
                        }
                    default:
                        if (!((v4 + (v4 * v4)) <= v4)) {
                            v69 := GenS5{F0: 810}
                            _ = v69
                        } else {
                            switch (fi65 >> ((0) & 63)) {
                            case 127:
                                v37 = v37
                            case 725:
                                var v70 int = fi61
                                _ = v70
                            default:
                                v71 := GenS8{F0: (v1 ^ 245)}
                                _ = v71
                            }
                            if true {
                                v47 = (^(v47))
                                if (0.0 < (-2.25 / 1e10)) {
                                    fmt.Println(v47, fi61, v48)
                                    if true {
                                        v72 := GenS7{F0: v1, F1: 0}
                                        _ = v72
                                    }
                                    v73 := GenS9{F0: v47}
                                    _ = v73
                                } else {
                                    v4 = v4
                                    fmt.Println(v47, fi65, v48)
                                    v74 := GenS9{F0: (v47 + v47)}
                                    _ = v74
                                }
                                var v75 float64 = ((float64(((float32(((-2.25 - (3.14159 + ((1e10 + (float64(1e10) + (((1e10 - -2.25) - 1.5) * 0.1))) / ((0.1 + 0.1) - 1.5)))) - 1e10)) + float32((float64(-2.25) + 1e10))) + (-2.25 - (((((((1.5 + (3.14159 - 1.5)) - (((1e10 * (0.1 * -2.25)) * 0.0) + float32(float64(-2.25)))) * (1e10 + 1e10)) - 0.0) + 0.1) * 0.0) * 0.0)))) + 3.14159) + 0.0)
                                _ = v75
                            }
                            fmt.Println(fi61, v4, v48)
                        }
                    }
                    var v76 uint = 954
                    _ = fi65
                    _ = v76
                }
            }
            _ = fi61
        }
    }
    v77 := GenS8{F0: -2147483648}
    var v78 uint = 0
    var v79 uint32 = v4
    v80 := GenS10{F0: -61}
    var v81 uint32 = v79
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v77
    _ = v78
    _ = v79
    _ = v80
    _ = v81
}
