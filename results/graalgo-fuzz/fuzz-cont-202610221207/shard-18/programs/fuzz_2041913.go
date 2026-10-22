package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 int32; F2 int16 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 uint32; F1 uint16 }
type GenS7 struct { F0 int8; F1 int8; F2 int8 }
type GenS8 struct { F0 int8; F1 uint64; F2 int }
func genFn0() uint16 {
    v7 := []int{32768}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = len(v7)
    _ = v7
    return 452
}
func genFn1(p0 int8, p1 int64) string {
    var v8 int = 42
    v9 := []int{(^(v8)), (^(v8))}
    for fi10 := 0; fi10 < 15; fi10++ {
        if false {
            v8 += (v8 * fi10)
        } else {
            if (1 != 99) {
                v9 = append(v9, (^(fi10)))
                v11 := GenS1{F0: 0, F1: -1}
                v9 = append(v9, 42)
                _ = v11
            }
            fmt.Println(p0, fi10)
            fmt.Println(fi10, fi10, p1)
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return (("" + "foo bar") + "a")
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{2147483647}
    var v3 float32 = 0.1
    v4 := make(map[int]bool)
    v5 := []string{"\n\t", "foo bar", ("foo bar" + ("a" + "a")), "hello"}
    var v6 string = ((("\n\t" + "foo bar") + "") + ("\n\t" + ("hello" + "hello")))
    for fi12 := 0; fi12 < 15; fi12++ {
        switch (fi12 * (fi12 * (fi12 - (fi12 * (fi12 / ((9223372036854775807) | 1)))))) {
        case 128:
            if true {
                var v13 float64 = -2.25
                if false {
                    if false {
                        v2 = append(v2, -2147483648)
                    } else {
                        v14 := GenS1{F0: -2147483648, F1: 127}
                        if true {
                            if false {
                                fmt.Println(fi12, v3, v3)
                                var v15 string = "hello"
                                _ = v15
                            }
                        } else {
                            v2 = append(v2, fi12)
                            var v16 int = -9223372036854775808
                            _ = v16
                        }
                        if ((!false) && ((!(-1 >= 109)) && (!false))) {
                            v6 = (("a" + "") + v6)
                        }
                        _ = v14
                    }
                    if true {
                        fmt.Println(fi12, v6)
                        v13 += 1.5
                    } else {
                        var v17 float64 = ((0.0 + (v13 + (((v13 / 1e10) + v13) + float64(1e10)))) + -2.25)
                        _ = v17
                    }
                }
                _ = v13
            } else {
                switch fi12 {
                case 467:
                    if false {
                        var v18 bool = (1 <= 0)
                        var v19 int32 = 42
                        var v20 float32 = 0.0
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                case 1:
                    switch v6 {
                    case "\n\t":
                        v5 = append(v5, (v6 + v6))
                    case "a":
                        if false {
                            if (50 > 42) {
                                v21 := GenS1{F0: 560, F1: -13}
                                v6 = v6
                                _ = v21
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = (fi12 / (((fi12 / (((fi12 ^ 293)) | 1))) | 1))
                                }
                                var v22 uint32 = 65535
                                switch (fi12 & 4294967295) {
                                case 558:
                                    for fi23 := 0; fi23 < 14; fi23++ {
                                        v24 := v4[(fi23 - (fi12 + 4294967295))]
                                        var v25 float32 = float32(1e10)
                                        _ = fi23
                                        _ = v24
                                        _ = v25
                                    }
                                }
                                _ = v22
                            }
                            v26 := v4[(fi12 / ((9223372036854775807) | 1))]
                            for fi27 := 0; fi27 < 1; fi27++ {
                                if (!("foo bar" != (((v6 + (v6 + ("foo bar" + "\n\t"))) + v6) + (((v6 + v6) + "\n\t") + v6)))) {
                                    fmt.Println(fi27, v6)
                                    var v28 int = 0
                                    v29 := GenS2{F0: 65535}
                                    _ = v28
                                    _ = v29
                                }
                                v3 = -2.25
                                var v30 uint16 = 307
                                _ = fi27
                                _ = v30
                            }
                            _ = v26
                        } else {
                            v31 := v4[(-(fi12))]
                            v3 += (v3 / 1.0)
                            if v31 {
                                v32 := GenS1{F0: 2147483647, F1: -128}
                                v6 = ("hello" + (("hello" + v6) + v6))
                                v33 := v1[v6]
                                _ = v32
                                _ = v33
                            }
                            _ = v31
                        }
                    default:
                        if true {
                            fmt.Println(fi12, v3, fi12)
                            if (fi12 < -417) {
                                var v34 bool = (0 == -32768)
                                switch -1 {
                                case 128:
                                    var v35 int16 = -298
                                    _ = v35
                                case 4294967295:
                                    v3 -= (((((v3 * (1e10 + v3)) + ((float32((-2.25 + 0.0)) - v3) / 1.0)) / v3) * 0.0) / v3)
                                default:
                                    var v36 string = "a"
                                    _ = v36
                                }
                                _ = v34
                            } else {
                                var v37 int = 32767
                                _ = len(v2)
                                switch "\n\t" {
                                case "hello":
                                    delete(v1, ((v6 + (v6 + v6)) + v6))
                                case "a":
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                default:
                                    v6 = (((((v6 + v6) + "foo bar") + "a") + "\n\t") + v6)
                                }
                                _ = v37
                            }
                            if false {
                                v38 := GenS3{F0: 65535, F1: -2147483648, F2: 32767}
                                v6 = "hello"
                                _ = v38
                            } else {
                                if (((v3 * v3) - (3.14159 / v3)) != ((v3 - v3) - (1.5 + (v3 / v3)))) {
                                    for fi39 := 0; fi39 < 3; fi39++ {
                                        v6 = v6
                                        var v40 string = v6
                                        _ = fi39
                                        _ = v40
                                    }
                                    if true {
                                        var v41 uint8 = 128
                                        v42 := v4[2147483647]
                                        fmt.Println(v6)
                                        _ = v41
                                        _ = v42
                                    } else {
                                        v3 = v3
                                    }
                                    v43 := GenS4{F0: 0, F1: true}
                                    _ = v43
                                }
                            }
                        }
                    }
                }
            }
        case 235:
            var v44 uint32 = 495
            _ = v44
        case 212:
            v6 = v6
        }
        v45 := GenS2{F0: 306}
        switch (fi12 - fi12) {
        case 313:
            for fi46 := 0; fi46 < 4; fi46++ {
                v47 := GenS5{F0: 65535, F1: -2147483648}
                var v48 uint64 = 534
                if (1000 >= 251) {
                    if (!(32767 <= (fi12 ^ fi46))) {
                        delete(v1, ((v6 + v6) + v6))
                        v48 -= (v48 + v48)
                        v49 := GenS4{F0: 650, F1: true}
                        _ = v49
                    }
                }
                var v50 int64 = -1
                if (!(!(4294967295 == 4294967295))) {
                    v3 = float32(0.1)
                    v48 = v48
                }
                _ = fi46
                _ = v47
                _ = v48
                _ = v50
            }
        }
        _ = fi12
        _ = v45
    }
    for fi51 := 0; fi51 < 1; fi51++ {
        if (!(!false)) {
            fmt.Println(v3, v6, v3)
            _ = len(v1)
            var v52 int32 = -2147483648
            _ = v52
        } else {
            if (-623 != 0) {
                var v53 bool = true
                var v54 uint = 467
                _ = v53
                _ = v54
            }
            var v55 bool = false
            for fi56 := 0; fi56 < 18; fi56++ {
                if (v6 == v6) {
                    if (0 < -787) {
                        if (!false) {
                            v57 := GenS3{F0: 488, F1: 682, F2: -32768}
                            _ = v57
                        } else {
                            var v58 int16 = 0
                            if true {
                                v5 = append(v5, (v6 + (v6 + v6)))
                                _ = len(v5)
                                v59 := GenS2{F0: 936}
                                if v55 {
                                    v60 := GenS3{F0: 311, F1: -2147483648, F2: (v58 | 722)}
                                    var v61 string = ""
                                    v58 -= (v58 + 1000)
                                    v55 = v55
                                    if len(v2) > 0 {
                                        v2[0] = (fi51 * fi51)
                                    }
                                    var v62 int8 = -1
                                    _ = v60
                                    _ = v61
                                    _ = v62
                                } else {
                                    v3 -= ((((0.1 * ((v3 - 1.5) * v3)) / 0.1) + 1e10) / (v3 / 1.0))
                                    v1["foo bar"] = fi51
                                }
                                v63 := GenS6{F0: 0, F1: genFn0()}
                                _ = v59
                                _ = v63
                            } else {
                                var v64 uint8 = 1
                                if (!true) {
                                    var v65 uint = 1
                                    v66 := GenS2{F0: 2}
                                    _ = v65
                                    _ = v66
                                } else {
                                    if ((v64 / ((129) | 1)) >= (v64 | v64)) {
                                        var v67 uint8 = 1
                                        _ = len(v4)
                                        _ = v67
                                    }
                                    v68 := v4[(fi51 ^ (fi56 - fi51))]
                                    v55 = v55
                                    _ = v68
                                }
                                switch (-(fi56)) {
                                case -1:
                                    _ = len(v4)
                                default:
                                    delete(v1, "a")
                                }
                                delete(v1, "a")
                                var v69 int64 = 1
                                _ = v64
                                _ = v69
                            }
                            _ = v58
                        }
                    }
                    var v70 int8 = -128
                    var v71 int32 = -770
                    _ = v70
                    _ = v71
                }
                if ("foo bar" != "hello") {
                    var v72 uint = 703
                    v73 := GenS7{F0: 100, F1: -128, F2: 42}
                    _ = v72
                    _ = v73
                }
                if true {
                    fmt.Println(v6, v3, v3)
                }
                _ = fi56
            }
            _ = v55
        }
        var v74 bool = (false && (42 < 127))
        _ = fi51
        _ = v74
    }
    if (0 == 100) {
        for fi75 := 0; fi75 < 6; fi75++ {
            v6 = v6
            v4[2147483647] = (564 < 1)
            _ = fi75
        }
        v6 = v6
    } else {
        switch (v6 + v6) {
        case "":
            for fi76 := 0; fi76 < 11; fi76++ {
                for fi77 := 0; fi77 < 3; fi77++ {
                    var v78 string = "foo bar"
                    v79 := GenS4{F0: 1, F1: (((true && false) || true) && (false || (127 < 128)))}
                    _ = fi77
                    _ = v78
                    _ = v79
                }
                if ("" != v6) {
                    for fi80 := 0; fi80 < 3; fi80++ {
                        var v81 float32 = -2.25
                        if (v3 > float32(3.14159)) {
                            v3 -= -2.25
                        }
                        var v82 float64 = (0.0 / 0.1)
                        var v83 int16 = 825
                        _ = fi80
                        _ = v81
                        _ = v82
                        _ = v83
                    }
                    _ = len(v1)
                    for fi84 := 0; fi84 < 6; fi84++ {
                        fmt.Println(fi76, v6, v6)
                        _ = fi84
                    }
                } else {
                    switch (fi76 | fi76) {
                    case 0:
                        fmt.Println(v6, v6, fi76)
                    case 145:
                        v3 -= ((v3 / float32(1e10)) / 1.0)
                    default:
                        delete(v1, v6)
                    }
                }
                _ = fi76
            }
        default:
            var v85 uint32 = 4294967295
            _ = v85
        }
        v86 := GenS2{F0: 255}
        var v87 uint32 = 504
        _ = v86
        _ = v87
    }
    switch -244 {
    case 32767:
        if (!true) {
            fmt.Println(v6, v3)
        }
    case 9223372036854775807:
        var v88 bool = (((v6 + v6) + (v6 + v6)) == v6)
        _ = v88
    case 0:
        var v89 int8 = 26
        _ = v89
    }
    var v90 uint32 = 65535
    if len(v2) > 0 {
        v2[0] = 18
    }
    delete(v4, 42)
    fmt.Println(v3)
    var v91 bool = false
    fmt.Println(v90)
    switch ("a" + v6) {
    case "hello":
        var v92 int64 = -14
        _ = v92
    }
    if (!(((2 < 0) && v91) && (188 == 127))) {
        switch v6 {
        case "":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "hello":
            _ = len(v2)
        case "a":
            v91 = v91
        }
    } else {
        if false {
            v3 = v3
            var v93 uint64 = 18446744073709551615
            if (42 == 42) {
                for fi94 := 0; fi94 < 17; fi94++ {
                    delete(v1, v6)
                    _ = fi94
                }
                v3 += v3
            }
            _ = v93
        } else {
            v95 := GenS8{F0: -69, F1: 2, F2: 32768}
            for fi96 := 0; fi96 < 8; fi96++ {
                v91 = ((!true) && (42 != 42))
                v91 = true
                _ = fi96
            }
            v3 -= float32((-2.25 * (0.0 + ((float64(v3) + 1.5) / 0.1))))
            _ = v95
        }
    }
    var v97 float64 = 1.5
    v98 := v4[317]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v90
    _ = v91
    _ = v97
    _ = v98
}
