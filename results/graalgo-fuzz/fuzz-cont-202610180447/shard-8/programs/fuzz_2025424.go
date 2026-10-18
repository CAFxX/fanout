package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32; F1 uint8; F2 uint16 }
func genFn0(p0 string, p1 int64) int8 {
    p1 += p1
    p0 = "hello"
    v7 := []int{269, 42}
    var v8 int = -1
    switch (p0 + p0) {
    case "\n\t":
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    default:
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 int16 = 32767
            fmt.Println(p0, p0, fi9)
            for fi11 := 0; fi11 < 2; fi11++ {
                if (-70 <= (v10 | v10)) {
                    var v12 int16 = -32768
                    v7 = append(v7, (fi11 - fi9))
                    if len(v7) > 0 {
                        v7[0] = fi9
                    }
                    _ = v12
                }
                v13 := GenS1{F0: p0}
                v14 := GenS1{F0: (("foo bar" + p0) + ((("foo bar" + "hello") + p0) + p0))}
                _ = fi11
                _ = v13
                _ = v14
            }
            _ = fi9
            _ = v10
        }
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int = -799
    var v2 bool = (!false)
    var v3 uint64 = 1
    v4 := []int{}
    var v5 uint32 = 273
    var v6 bool = ((124 != 1) || (!(v2 && ((!false) && v2))))
    v5 = (^(v5))
    var v15 int32 = 535
    if ((v6 && ((!(v15 > 42)) || false)) || ((v3 >= 910) || (!v2))) {
        var v16 int32 = -1
        for fi17 := 0; fi17 < 11; fi17++ {
            if len(v4) > 0 {
                v4[0] = fi17
            }
            _ = fi17
        }
        switch (v1 / ((v1) | 1)) {
        case 32768:
            switch ("a" + (("" + "a") + "")) {
            case "a":
                var v18 string = "foo bar"
                _ = v18
            case "hello":
                for fi19 := 0; fi19 < 9; fi19++ {
                    var v20 float32 = (float32((-2.25 * float64(3.14159))) + float32(float64(-2.25)))
                    _ = fi19
                    _ = v20
                }
            case "foo bar":
                var v21 float64 = -2.25
                _ = v21
            }
        case 9223372036854775807:
            if v6 {
                _ = len(v4)
                if (1.5 <= 0.0) {
                    for fi22 := 0; fi22 < 4; fi22++ {
                        v23 := GenS2{F0: v16, F1: 0}
                        for fi24 := 0; fi24 < 8; fi24++ {
                            if (!true) {
                                if (((v3 & (v3 << ((2) & 63))) > (v3 & (v3 % (((v3 >> ((881) & 63))) | 1)))) && (v15 <= v15)) {
                                    if v6 {
                                        v5 -= v5
                                        var v25 uint32 = v5
                                        _ = v25
                                    }
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                if ((v6 || (42 <= 2147483647)) && false) {
                                    var v26 float64 = (float64((0.0 - -2.25)) - ((1e10 + (-2.25 + 0.0)) * (float64(0.0) * (((0.0 - (1.5 + 3.14159)) - (1.5 - 1e10)) - 1.5))))
                                    _ = v26
                                } else {
                                    switch "" {
                                    case "":
                                        fmt.Println(v1, v15)
                                    }
                                    var v27 uint = 825
                                    v1 -= 4294967295
                                    _ = v27
                                }
                                var v28 float64 = 1e10
                                _ = v28
                            } else {
                                var v29 uint16 = 1
                                var v30 int8 = 43
                                v31 := GenS2{F0: v15, F1: 60}
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                            var v32 int16 = 42
                            if (v6 || false) {
                                if (uint32(v3) < v5) {
                                    if (v5 == v5) {
                                        if (v2 || (v15 == (v16 & v16))) {
                                            if len(v4) > 0 {
                                                v4[0] = (fi24 & (-(v1)))
                                            }
                                            if (-91 == genFn0("\n\t", 1)) {
                                                var v33 uint16 = 487
                                                switch "" {
                                                case "hello":
                                                    v16 = v16
                                                    if (v32 < v32) {
                                                        switch -1 {
                                                        case -246:
                                                            v34 := GenS3{F0: 70}
                                                            _ = v34
                                                        case -565:
                                                            v3 -= 42
                                                        case 641:
                                                            var v35 uint64 = v3
                                                            var v36 string = (("\n\t" + ("hello" + "\n\t")) + "foo bar")
                                                            v2 = (!true)
                                                            _ = v35
                                                            _ = v36
                                                        default:
                                                            var v37 uint64 = 112
                                                            _ = v37
                                                        }
                                                    } else {
                                                        v15 = -148
                                                    }
                                                    v5 -= v5
                                                case "a":
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                case "foo bar":
                                                    fmt.Println(v5, v32)
                                                }
                                                v32 = -45
                                                _ = v33
                                            }
                                        } else {
                                            fmt.Println(fi22, v6, v3)
                                        }
                                        var v38 int16 = (v32 ^ (v32 | (v32 * (v32 ^ 32767))))
                                        var v39 float64 = 1e10
                                        _ = v38
                                        _ = v39
                                    } else {
                                        switch v1 {
                                        case -46:
                                            v40 := GenS4{F0: 468}
                                            _ = v40
                                        }
                                    }
                                    switch fi24 {
                                    case 32767:
                                        v4 = append(v4, 32767)
                                    default:
                                        v1 = (v1 << ((2) & 63))
                                    }
                                    if v6 {
                                        switch ("hello" + "\n\t") {
                                        case "hello":
                                            var v41 uint32 = v5
                                            _ = v41
                                        case "foo bar":
                                            v16 = 2147483647
                                        }
                                        fmt.Println(fi24, v32, fi22)
                                    } else {
                                        var v42 int64 = 118
                                        var v43 int16 = v32
                                        v15 += v15
                                        _ = v42
                                        _ = v43
                                    }
                                    var v44 float64 = (1.5 / 3.14159)
                                    var v45 uint16 = 376
                                    _ = v44
                                    _ = v45
                                }
                            } else {
                                var v46 uint = 2
                                _ = v46
                            }
                            if (-128 <= 104) {
                                if v2 {
                                    var v47 int = (v1 % ((2147483647) | 1))
                                    _ = v47
                                } else {
                                    fmt.Println(fi22, fi24)
                                    if ((v16 > (v15 >> ((589) & 63))) && (4294967295 >= 838)) {
                                        var v48 int8 = genFn0(("" + ("\n\t" + ("" + "\n\t"))), 42)
                                        _ = len(v4)
                                        _ = v48
                                    } else {
                                        var v49 uint32 = (v5 >> ((255) & 63))
                                        var v50 float32 = (1.5 / 1e10)
                                        _ = v49
                                        _ = v50
                                    }
                                }
                                if v2 {
                                    fmt.Println(v16)
                                } else {
                                    if (v6 && true) {
                                        v15 += 461
                                        v2 = (v16 < v15)
                                        v51 := GenS2{F0: (^(v16)), F1: 2}
                                        _ = v51
                                    }
                                    fmt.Println(v3, fi22)
                                }
                            } else {
                                v3 = (v3 >> ((1) & 63))
                                v52 := GenS1{F0: ("foo bar" + "a")}
                                _ = v52
                            }
                            _ = fi24
                            _ = v32
                        }
                        if (100 > -128) {
                            v53 := GenS5{F0: ("hello" + "\n\t")}
                            v2 = (v6 && v2)
                            _ = v53
                        } else {
                            var v54 string = "\n\t"
                            var v55 uint8 = 255
                            var v56 bool = ((v1 + v1) < (fi22 >> ((18446744073709551615) & 63)))
                            _ = v54
                            _ = v55
                            _ = v56
                        }
                        _ = fi22
                        _ = v23
                    }
                }
            }
        case 0:
            fmt.Println(v2, v16, v3)
        }
        _ = v16
    } else {
        v3 += v3
        fmt.Println(v6, v5, v1)
        v1 += -2147483648
    }
    v15 -= (v15 - v15)
    for fi57 := 0; fi57 < 4; fi57++ {
        var v58 uint64 = (v3 % (((^(v3))) | 1))
        for fi59 := 0; fi59 < 12; fi59++ {
            switch 9223372036854775807 {
            case 1:
                v58 = (v58 - (v58 & v3))
            case 32767:
                switch "" {
                case "a":
                    v60 := GenS2{F0: v15, F1: 255}
                    _ = v60
                case "foo bar":
                    fmt.Println(fi59, v58)
                default:
                    switch "hello" {
                    case "a":
                        fmt.Println(fi57, v15, fi59)
                    default:
                        if false {
                            var v61 int32 = (v15 / ((v15) | 1))
                            var v62 int16 = 1000
                            v63 := GenS2{F0: (v61 + v15), F1: 42}
                            _ = v61
                            _ = v62
                            _ = v63
                        } else {
                            var v64 int = fi57
                            _ = v64
                        }
                    }
                }
            }
            v65 := GenS6{F0: 1, F1: 1, F2: 32767}
            _ = fi59
            _ = v65
        }
        fmt.Println(v15, v3)
        _ = fi57
        _ = v58
    }
    var v66 uint16 = 1
    v6 = true
    v1 = v1
    v67 := GenS6{F0: 422, F1: 49, F2: v66}
    for fi68 := 0; fi68 < 15; fi68++ {
        fmt.Println(v1, v1, v3)
        _ = fi68
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v67
}
