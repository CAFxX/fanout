package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS5 struct { F0 int32; F1 uint16; F2 bool }
type GenS6 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS7 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS8 struct { F0 string; F1 int32; F2 uint8 }
func genFn0(p0 uint16, p1 int16) uint32 {
    var v5 bool = true
    if (42 == 65535) {
        fmt.Println(v5, v5, v5)
    }
    fmt.Println(p0, v5)
    if (!((("a" + "hello") >= ("hello" + ("foo bar" + ("" + "foo bar")))) && v5)) {
        v6 := GenS1{F0: p0, F1: 546, F2: (2147483647 >= 921)}
        for fi7 := 0; fi7 < 7; fi7++ {
            for fi8 := 0; fi8 < 2; fi8++ {
                if true {
                    v9 := GenS2{F0: 127}
                    switch int(p0) {
                    case 738:
                        var v10 uint8 = 2
                        _ = v10
                    case 0:
                        v11 := []int{(fi7 * fi8), (fi7 * fi7)}
                        _ = v11
                    case 32767:
                        switch (^(fi8)) {
                        case 128:
                            var v12 bool = v5
                            _ = v12
                        case 32768:
                            var v13 uint32 = 65535
                            _ = v13
                        default:
                            switch (fi8 >> ((381) & 63)) {
                            case -1:
                                var v14 bool = ((((("a" + "a") + "foo bar") + "foo bar") + "") >= (("hello" + (("\n\t" + "hello") + ("foo bar" + ""))) + ""))
                                _ = v14
                            case -9223372036854775808:
                                v5 = v5
                            case 42:
                                if false {
                                    v5 = v5
                                    var v15 bool = v5
                                    if (true && v15) {
                                        var v16 uint8 = 127
                                        v16 -= (v16 - 127)
                                        var v17 bool = (v15 && false)
                                        _ = v16
                                        _ = v17
                                    } else {
                                        switch 42 {
                                        case 42:
                                            var v18 int16 = p1
                                            _ = v18
                                        case 866:
                                            v19 := GenS2{F0: -128}
                                            _ = v19
                                        case -34:
                                            if ((0.0 * (-2.25 / -2.25)) == 1e10) {
                                                p1 = (p1 & (p1 << ((2) & 63)))
                                                p1 -= p1
                                            } else {
                                                v20 := []int{(-(fi7)), 393}
                                                if false {
                                                    v21 := GenS1{F0: (p0 + p0), F1: 2147483647, F2: ((!(1 == 249)) && (!true))}
                                                    fmt.Println(fi7)
                                                    var v22 float32 = 1e10
                                                    _ = v21
                                                    _ = v22
                                                }
                                                _ = v20
                                            }
                                        }
                                        fmt.Println(p0, v5)
                                        v23 := GenS3{F0: p1, F1: 255}
                                        var v24 float32 = 0.1
                                        _ = v23
                                        _ = v24
                                    }
                                    var v25 int16 = p1
                                    _ = v15
                                    _ = v25
                                } else {
                                    p1 -= (p1 | 0)
                                }
                            default:
                                var v26 uint32 = 4294967295
                                _ = v26
                            }
                        }
                    }
                    var v27 string = "hello"
                    if (!true) {
                        v27 = ("foo bar" + (v27 + ((v27 + v27) + v27)))
                        v27 = v27
                        var v28 uint64 = 0
                        _ = v28
                    } else {
                        if v5 {
                            var v29 int64 = 1
                            _ = v29
                        }
                        var v30 uint32 = 0
                        _ = v30
                    }
                    _ = v27
                    _ = v9
                } else {
                    v31 := GenS2{F0: 115}
                    fmt.Println(v5, p1, p0)
                    _ = v31
                }
                v32 := GenS3{F0: 1, F1: 31}
                _ = fi8
                _ = v32
            }
            fmt.Println(v5, v5)
            fmt.Println(v5, v5)
            _ = fi7
        }
        _ = v6
    } else {
        v5 = v5
        p0 -= p0
    }
    _ = p0
    _ = p1
    _ = v5
    return 65535
}
func genFn1(p0 uint64, p1 int) string {
    p1 -= 0
    switch (("hello" + "") + "foo bar") {
    case "":
        switch p1 {
        case 0:
            v33 := GenS4{F0: p0, F1: 0, F2: 1}
            _ = v33
        default:
            p1 -= p1
        }
    }
    v34 := GenS4{F0: (p0 >> ((883) & 63)), F1: p0, F2: -1}
    _ = p0
    _ = p1
    _ = v34
    return ("foo bar" + "hello")
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{9223372036854775807, -1, 1}
    var v3 int16 = 32767
    var v4 uint64 = 164
    for fi35 := 0; fi35 < 11; fi35++ {
        v36 := GenS5{F0: -1, F1: 65535, F2: false}
        fmt.Println(v3, v4, v4)
        var v37 int8 = -63
        _ = fi35
        _ = v36
        _ = v37
    }
    v38 := GenS6{F0: 128, F1: -9223372036854775808, F2: 1}
    for fi39 := 0; fi39 < 2; fi39++ {
        for fi40 := 0; fi40 < 3; fi40++ {
            if true {
                switch 32767 {
                case 0:
                    if (-2.25 < 3.14159) {
                        v3 += v3
                    } else {
                        v4 -= (v4 & (v4 + v4))
                    }
                case 42:
                    var v41 int16 = v3
                    _ = v41
                case -1:
                    v4 -= v4
                }
                _ = len(v2)
                if (fi40 < 9223372036854775807) {
                    if false {
                        v4 = 913
                        v3 += (v3 + (v3 & (v3 >> ((0) & 63))))
                    } else {
                        var v42 uint64 = 448
                        _ = v42
                    }
                    v43 := GenS7{F0: 561, F1: 675, F2: 156}
                    _ = v43
                } else {
                    var v44 int = fi39
                    fmt.Println(v44, v44)
                    v45 := GenS6{F0: 255, F1: 0, F2: 2}
                    _ = v44
                    _ = v45
                }
                v46 := GenS8{F0: "\n\t", F1: 360, F2: 2}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch ("hello" + "\n\t") {
                case "\n\t":
                    var v47 int64 = -1
                    _ = v47
                case "hello":
                    var v48 bool = (0 >= -128)
                    _ = v48
                default:
                    var v49 float64 = 1e10
                    _ = v49
                }
                _ = v46
            } else {
                v3 = v3
                v50 := GenS6{F0: 255, F1: 9223372036854775807, F2: 2}
                _ = v50
            }
            _ = fi40
        }
        fmt.Println(v4, v3, v4)
        _ = fi39
    }
    switch 0 {
    case 436:
        for fi51 := 0; fi51 < 13; fi51++ {
            var v52 uint64 = v4
            switch "foo bar" {
            case "":
                switch ("a" + "foo bar") {
                case "foo bar":
                    fmt.Println(fi51, fi51, v52)
                case "hello":
                    if false {
                        if true {
                            if len(v2) > 0 {
                                v2[0] = 1
                            }
                        }
                        if (223 == 127) {
                            for fi53 := 0; fi53 < 5; fi53++ {
                                _ = len(v1)
                                _ = fi53
                            }
                            var v54 uint = 1
                            _ = v54
                        } else {
                            v52 = 0
                        }
                        var v55 int16 = 1000
                        _ = v55
                    }
                case "":
                    v4 = 18446744073709551615
                default:
                    if true {
                        v56 := GenS8{F0: "foo bar", F1: -1, F2: 47}
                        _ = v56
                    } else {
                        if (1 <= (fi51 / ((0) | 1))) {
                            delete(v1, "")
                            var v57 uint16 = 0
                            _ = v57
                        } else {
                            var v58 uint8 = 1
                            v4 = 925
                            var v59 int16 = 1000
                            _ = v58
                            _ = v59
                        }
                    }
                }
            }
            _ = fi51
            _ = v52
        }
    default:
        switch ("" + "") {
        case "hello":
            v4 = v4
        case "\n\t":
            if false {
                v3 += v3
            } else {
                var v60 uint64 = v4
                v61 := GenS4{F0: v60, F1: (v60 + v4), F2: v3}
                _ = v60
                _ = v61
            }
        default:
            switch ("hello" + "a") {
            case "\n\t":
                v4 = v4
            case "a":
                v62 := GenS5{F0: -1, F1: 44, F2: (!true)}
                _ = v62
            default:
                if (float64(3.14159) != (1e10 * 3.14159)) {
                    var v63 int = 0
                    v4 = v4
                    _ = v63
                } else {
                    if (((float64((0.1 / (1e10 * 1.5))) - 1.5) / ((((1e10 / (((0.0 / ((1.5 + ((3.14159 + 1.5) + (0.0 - ((float64((0.1 * 0.1)) / 0.1) - -2.25)))) - -2.25)) + (1.5 * float64(1e10))) * (1.5 + 3.14159))) * -2.25) * (1e10 - 0.0)) + 3.14159)) != -2.25) {
                        v3 -= v3
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(v3, v4)
                        default:
                            switch "hello" {
                            case "a":
                                var v64 int32 = -737
                                _ = v64
                            case "foo bar":
                                var v65 int = 4294967295
                                _ = v65
                            case "hello":
                                switch "foo bar" {
                                case "\n\t":
                                    var v66 int32 = 0
                                    _ = v66
                                default:
                                    var v67 bool = true
                                    _ = v67
                                }
                            }
                        }
                    }
                    var v68 uint = 360
                    _ = v68
                }
            }
        }
    }
    v69 := GenS3{F0: v3, F1: 42}
    var v70 uint64 = v4
    var v71 string = "hello"
    switch v71 {
    case "\n\t":
        delete(v1, (v71 + "a"))
    case "hello":
        v71 = v71
    default:
        v70 -= v4
    }
    v72 := GenS4{F0: 18446744073709551615, F1: 151, F2: v3}
    v73 := v1["\n\t"]
    fmt.Println(v70, v73, v4)
    delete(v1, v71)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v69
    _ = v70
    _ = v71
    _ = v72
    _ = v73
}
