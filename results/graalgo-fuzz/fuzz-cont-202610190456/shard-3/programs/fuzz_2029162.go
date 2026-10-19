package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int64 }
type GenS2 struct { F0 int; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64; F1 int32; F2 int }
type GenS7 struct { F0 int; F1 int16 }
type GenS8 struct { F0 int16; F1 bool }
type GenS9 struct { F0 int16; F1 uint8; F2 int32 }
type GenS10 struct { F0 int64; F1 uint8; F2 string }
func genFn0(p0 int16, p1 uint64) uint32 {
    var v7 string = ("foo bar" + "hello")
    p1 = (p1 << ((4294967295) & 63))
    p0 = 1000
    _ = p0
    _ = p1
    _ = v7
    return 2
}
func genFn1(p0 uint, p1 string) int64 {
    var v8 float64 = 0.1
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 int32 = -713
        _ = fi9
        _ = v10
    }
    switch p1 {
    case "":
        v11 := GenS1{F0: -83, F1: 1, F2: -307}
        _ = v11
    case "foo bar":
        switch 0 {
        case -2147483648:
            p0 = 4294967295
        default:
            for fi12 := 0; fi12 < 19; fi12++ {
                if (!((!false) || (255 < 42))) {
                    v13 := []int{fi12, (fi12 >> (((p0 << (((p0 >> (((p0 & (p0 ^ 1))) & 63))) & 63))) & 63)), 127}
                    if true {
                        switch 4294967295 {
                        case -194:
                            fmt.Println(v8, p0, fi12)
                        case 458:
                            switch (fi12 >> ((0) & 63)) {
                            case 2147483647:
                                v8 = v8
                            case 32767:
                                for fi14 := 0; fi14 < 12; fi14++ {
                                    switch ("\n\t" + (p1 + ("hello" + (p1 + ((p1 + p1) + p1))))) {
                                    case "hello":
                                        _ = len(v13)
                                    case "":
                                        v13 = append(v13, 32768)
                                    }
                                    _ = fi14
                                }
                            }
                        case 1:
                            if true {
                                p0 += p0
                                v8 -= v8
                            }
                        }
                        v8 = 0.1
                    } else {
                        v8 = (((v8 * v8) - (v8 + (v8 + (v8 + v8)))) + v8)
                        var v15 int64 = -497
                        _ = v15
                    }
                    v13 = append(v13, fi12)
                    _ = v13
                } else {
                    p1 = ""
                    v8 -= v8
                    for fi16 := 0; fi16 < 6; fi16++ {
                        if false {
                            p0 -= p0
                        } else {
                            v17 := GenS1{F0: 1, F1: fi12, F2: 1}
                            _ = v17
                        }
                        switch ("foo bar" + (p1 + (p1 + ((((p1 + p1) + "") + ("hello" + p1)) + ("a" + ("foo bar" + p1)))))) {
                        case "foo bar":
                            v8 = v8
                        case "":
                            var v18 float64 = float64(0.1)
                            _ = v18
                        case "a":
                            var v19 int16 = 308
                            _ = v19
                        default:
                            v20 := []int{(fi16 | fi16), fi16, 0}
                            _ = v20
                        }
                        var v21 uint64 = 4294967295
                        _ = fi16
                        _ = v21
                    }
                }
                v22 := GenS1{F0: -128, F1: fi12, F2: -9223372036854775808}
                _ = fi12
                _ = v22
            }
        }
    }
    switch p1 {
    case "\n\t":
        for fi23 := 0; fi23 < 8; fi23++ {
            if (false || (true || false)) {
                fmt.Println(p1, v8, fi23)
            } else {
                switch ((p1 + p1) + "foo bar") {
                case "a":
                    for fi24 := 0; fi24 < 1; fi24++ {
                        v25 := GenS1{F0: -33, F1: fi23, F2: -1}
                        p1 = ("\n\t" + p1)
                        switch p1 {
                        case "":
                            switch (fi24 * 126) {
                            case -9223372036854775808:
                                var v26 uint16 = 560
                                _ = v26
                            case -535:
                                var v27 uint32 = 2147483647
                                _ = v27
                            default:
                                v8 = -2.25
                            }
                        case "a":
                            var v28 int64 = -9223372036854775808
                            _ = v28
                        default:
                            if true {
                                var v29 uint = (p0 % ((255) | 1))
                                p1 = ("a" + "foo bar")
                                p1 = p1
                                _ = v29
                            } else {
                                var v30 uint32 = 2
                                if (!(!(!false))) {
                                    v8 -= v8
                                    if (!false) {
                                        v31 := GenS2{F0: (-(fi23)), F1: 92, F2: 128}
                                        var v32 int = 32768
                                        v33 := GenS1{F0: 127, F1: 9223372036854775807, F2: 9223372036854775807}
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    }
                                    var v34 uint = 18446744073709551615
                                    _ = v34
                                }
                                if true {
                                    var v35 int8 = 127
                                    _ = v35
                                } else {
                                    v8 = 1e10
                                    if false {
                                        var v36 uint16 = 65535
                                        _ = v36
                                    } else {
                                        p1 = ""
                                        if (2 <= 701) {
                                            v8 += (v8 * (v8 / v8))
                                        }
                                    }
                                }
                                _ = v30
                            }
                        }
                        _ = fi24
                        _ = v25
                    }
                case "foo bar":
                    for fi37 := 0; fi37 < 19; fi37++ {
                        switch p1 {
                        case "hello":
                            v38 := []int{(^(fi37)), 534, (fi23 ^ fi37)}
                            _ = v38
                        }
                        if (0.0 >= ((v8 / v8) - 0.0)) {
                            v39 := GenS3{F0: 255}
                            v40 := GenS3{F0: 128}
                            v41 := GenS4{F0: 123, F1: (("hello" + "foo bar") + ("\n\t" + p1))}
                            switch ((((p1 + p1) + (p1 + "")) + p1) + ("foo bar" + (p1 + p1))) {
                            case "foo bar":
                                p1 = p1
                            default:
                                var v42 float32 = 3.14159
                                _ = v42
                            }
                            fmt.Println(fi37, p1)
                            if (747 >= 1) {
                                p0 = 1
                                p1 = p1
                                v8 = (-2.25 * (-2.25 + float64(float32(float64((-2.25 * float32(v8)))))))
                            } else {
                                var v43 uint16 = 0
                                v44 := []int{}
                                _ = v43
                                _ = v44
                            }
                            p0 = p0
                            var v45 uint16 = 42
                            _ = v39
                            _ = v40
                            _ = v41
                            _ = v45
                        }
                        if false {
                            var v46 bool = (false && false)
                            _ = v46
                        }
                        _ = fi37
                    }
                }
                var v47 bool = ((("hello" + p1) >= p1) && (!false))
                v47 = true
                _ = v47
            }
            fmt.Println(p0, p1)
            _ = fi23
        }
    case "foo bar":
        for fi48 := 0; fi48 < 3; fi48++ {
            var v49 int = (fi48 % (((^(fi48))) | 1))
            _ = fi48
            _ = v49
        }
    case "":
        p1 = p1
    default:
        if false {
            for fi50 := 0; fi50 < 3; fi50++ {
                for fi51 := 0; fi51 < 20; fi51++ {
                    fmt.Println(fi51, fi50, fi51)
                    fmt.Println(v8, p0)
                    if (289 > -1) {
                        if true {
                            p0 -= (p0 & p0)
                        } else {
                            switch "hello" {
                            case "foo bar":
                                var v52 float64 = v8
                                _ = v52
                            default:
                                var v53 int64 = 0
                                _ = v53
                            }
                            v54 := GenS5{F0: 0}
                            _ = v54
                        }
                    }
                    v55 := GenS6{F0: 1, F1: 1, F2: -1}
                    var v56 int64 = -9223372036854775808
                    v57 := GenS7{F0: (fi51 ^ (fi50 * (fi51 << ((p0) & 63)))), F1: int16(p0)}
                    _ = fi51
                    _ = v55
                    _ = v56
                    _ = v57
                }
                v58 := []int{fi50}
                for fi59 := 0; fi59 < 14; fi59++ {
                    v60 := GenS8{F0: 42, F1: false}
                    fmt.Println(fi59, p1, v8)
                    _ = fi59
                    _ = v60
                }
                _ = fi50
                _ = v58
            }
            var v61 uint = p0
            _ = v61
        }
    }
    _ = p0
    _ = p1
    _ = v8
    return 1
}

func main() {
    v1 := []uint32{289}
    var v2 uint = 4294967295
    var v3 float64 = 0.1
    v4 := make(map[string]int)
    var v5 int64 = -579
    var v6 string = "foo bar"
    if ((true || (true || true)) || false) {
        for fi62 := 0; fi62 < 12; fi62++ {
            v5 -= 1
            v63 := GenS8{F0: 554, F1: true}
            _ = fi62
            _ = v63
        }
        for fi64 := 0; fi64 < 17; fi64++ {
            for fi65 := 0; fi65 < 10; fi65++ {
                if false {
                    fmt.Println(fi65)
                }
                v66 := GenS9{F0: -1, F1: 255, F2: -877}
                _ = fi65
                _ = v66
            }
            v67 := GenS6{F0: 2, F1: -2147483648, F2: fi64}
            var v68 uint16 = 32767
            _ = fi64
            _ = v67
            _ = v68
        }
        if (false || (!(true || (42 != -374)))) {
            var v69 int16 = 42
            _ = v69
        } else {
            v3 = v3
        }
    } else {
        v2 -= v2
        v5 = v5
    }
    fmt.Println(v5, v2, v3)
    if (((0.1 - (-2.25 - ((((3.14159 + 3.14159) * 0.1) / (3.14159 * 1.5)) - float32((-2.25 * float64(1.5)))))) - (-2.25 - 1.5)) >= 1.5) {
        v1 = append(v1, 290)
        for fi70 := 0; fi70 < 17; fi70++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v71 int8 = 108
            _ = fi70
            _ = v71
        }
    } else {
        var v72 bool = (-1 >= 1)
        switch 525 {
        case 128:
            v2 = (v2 & 0)
        case 222:
            v72 = (-1 <= 1)
        }
        var v73 uint = 42
        _ = v72
        _ = v73
    }
    var v74 int16 = 0
    for fi75 := 0; fi75 < 7; fi75++ {
        var v76 float32 = ((1e10 * (0.0 + 1.5)) / (1e10 - 0.1))
        v77 := GenS2{F0: 127, F1: v74, F2: 1}
        _ = fi75
        _ = v76
        _ = v77
    }
    fmt.Println(v74)
    fmt.Println(v74)
    var v78 int = 2147483647
    if false {
        v79 := GenS10{F0: v5, F1: 42, F2: v6}
        _ = v79
    }
    var v80 uint16 = 330
    var v81 bool = true
    fmt.Println(v81, v2, v81)
    v5 -= -9223372036854775808
    v82 := v4[v6]
    var v83 int = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v74:", v74)
    fmt.Println("v78:", v78)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v74
    _ = v78
    _ = v80
    _ = v81
    _ = v82
    _ = v83
}
