package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint8; F1 int32; F2 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint; F1 uint32 }
type GenS6 struct { F0 uint8; F1 uint8 }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int; F1 string }
func genFn0(p0 int16, p1 int) int16 {
    p1 = p1
    v5 := GenS1{F0: true, F1: (402 < -951)}
    var v6 string = "hello"
    v7 := GenS1{F0: true, F1: true}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return p0
}

func main() {
    v1 := []int{}
    v2 := make(map[string]uint32)
    var v3 int64 = 0
    var v4 uint64 = 18446744073709551615
    if false {
        if ((false || true) && true) {
            if true {
                _ = len(v1)
                v8 := GenS1{F0: false, F1: false}
                var v9 string = "foo bar"
                _ = v8
                _ = v9
            }
            var v10 uint32 = 1
            _ = v10
        }
        switch "hello" {
        case "hello":
            v11 := GenS1{F0: false, F1: ((float64((0.1 / 1.5)) / 1.5) < ((3.14159 - float64(1.5)) * 1e10))}
            _ = v11
        case "foo bar":
            if (true && true) {
                var v12 int16 = 1
                if (0 == 187) {
                    if (-2147483648 != 127) {
                        v12 = v12
                        if (!(-142 >= -1)) {
                            var v13 float32 = (1.5 * -2.25)
                            _ = v13
                        } else {
                            for fi14 := 0; fi14 < 19; fi14++ {
                                if (true || false) {
                                    if false {
                                        var v15 float32 = ((0.0 / 1.5) + 3.14159)
                                        v2[""] = 4294967295
                                        _ = v15
                                    } else {
                                        for fi16 := 0; fi16 < 3; fi16++ {
                                            _ = len(v2)
                                            _ = len(v1)
                                            _ = fi16
                                        }
                                    }
                                    v12 = v12
                                }
                                v12 = v12
                                _ = fi14
                            }
                            v17 := v2[((("\n\t" + (("\n\t" + "") + "foo bar")) + "\n\t") + "foo bar")]
                            _ = v17
                        }
                    } else {
                        switch "a" {
                        case "foo bar":
                            v4 += 373
                        }
                        v18 := GenS2{F0: 0, F1: -1, F2: v3}
                        _ = v18
                    }
                }
                switch "a" {
                case "hello":
                    if false {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v12 = v12
                    } else {
                        var v19 uint64 = 0
                        if true {
                            var v20 uint16 = 763
                            if ((((1.5 + (float32(0.0) - (((-2.25 + 1e10) - 1e10) + (1.5 + 0.1)))) + 1.5) / 0.1) <= 3.14159) {
                                if true {
                                    if (2147483647 < 874) {
                                        v21 := GenS3{F0: -128}
                                        _ = v21
                                    } else {
                                        v22 := v2[""]
                                        _ = v22
                                    }
                                    fmt.Println(v19)
                                    v23 := GenS4{F0: 0}
                                    _ = v23
                                }
                                fmt.Println(v4, v20, v4)
                            } else {
                                v2["a"] = 1
                            }
                            var v24 uint16 = v20
                            _ = v20
                            _ = v24
                        } else {
                            switch -9223372036854775808 {
                            case 127:
                                v25 := GenS5{F0: 1, F1: 978}
                                _ = v25
                            case 0:
                                for fi26 := 0; fi26 < 16; fi26++ {
                                    for fi27 := 0; fi27 < 6; fi27++ {
                                        var v28 float64 = 0.1
                                        _ = fi27
                                        _ = v28
                                    }
                                    v4 = 18446744073709551615
                                    v12 = v12
                                    _ = fi26
                                }
                            }
                        }
                        _ = v19
                    }
                case "":
                    if len(v1) > 0 {
                        v1[0] = -1
                    }
                default:
                    var v29 uint64 = (v4 - (v4 | v4))
                    _ = v29
                }
                _ = v12
            }
        }
    } else {
        var v30 uint32 = 564
        _ = v30
    }
    _ = len(v1)
    v31 := GenS1{F0: (255 != 32767), F1: (4294967295 <= 65535)}
    for fi32 := 0; fi32 < 7; fi32++ {
        fmt.Println(v3)
        if false {
            switch "a" {
            case "a":
                v3 = v3
            case "foo bar":
                fmt.Println(v4, v4)
            case "hello":
                var v33 int = (fi32 << ((2) & 63))
                _ = v33
            }
        } else {
            v4 = (v4 ^ (v4 & v4))
            v4 = v4
            switch "" {
            case "a":
                var v34 float64 = 0.1
                _ = v34
            case "\n\t":
                switch "" {
                case "":
                    v35 := GenS6{F0: 127, F1: 127}
                    _ = v35
                case "a":
                    v4 = (v4 + v4)
                }
            case "":
                var v36 bool = (-128 > 100)
                _ = v36
            default:
                v3 -= (v3 + (v3 | v3))
            }
        }
        _ = fi32
    }
    v37 := GenS4{F0: -525}
    switch "\n\t" {
    case "\n\t":
        var v38 int32 = 42
        _ = v38
    }
    v3 = v3
    var v39 float64 = (0.1 - ((1.5 / 1e10) - 0.1))
    v40 := GenS7{F0: 65535, F1: 9223372036854775807}
    v4 = 0
    switch 2147483647 {
    case 106:
        if (false || true) {
            if (float64((1.5 - float32((0.0 - v39)))) >= 1.5) {
                v4 = (^(v4))
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v41 uint16 = 32767
            _ = v41
        } else {
            if true {
                if (false && (false && (!false))) {
                    if ("hello" == "foo bar") {
                        v39 += (3.14159 + (v39 / 0.1))
                    } else {
                        if ((v3 - v3) == v3) {
                            for fi42 := 0; fi42 < 6; fi42++ {
                                if (!false) {
                                    if (!(2 <= 42)) {
                                        v2["a"] = 1
                                    } else {
                                        if (false && (922 != 2147483648)) {
                                            fmt.Println(v4)
                                            v43 := GenS8{F0: (v4 | (^(v4)))}
                                            _ = v43
                                        }
                                        if false {
                                            var v44 int64 = (v3 << ((2) & 63))
                                            if false {
                                                v4 = 427
                                            }
                                            v1 = append(v1, 32767)
                                            _ = v44
                                        } else {
                                            var v45 int = fi42
                                            if false {
                                                switch ("foo bar" + "a") {
                                                case "foo bar":
                                                    v39 -= v39
                                                case "hello":
                                                    for fi46 := 0; fi46 < 12; fi46++ {
                                                        delete(v2, "foo bar")
                                                        v47 := GenS3{F0: -40}
                                                        var v48 uint = 0
                                                        v49 := v2[("a" + "foo bar")]
                                                        _ = fi46
                                                        _ = v47
                                                        _ = v48
                                                        _ = v49
                                                    }
                                                case "\n\t":
                                                    switch 32767 {
                                                    case -729:
                                                        var v50 int16 = 1000
                                                        _ = v50
                                                    case -9223372036854775808:
                                                        var v51 uint8 = 128
                                                        _ = v51
                                                    case 830:
                                                        v3 = v3
                                                    default:
                                                        if false {
                                                            fmt.Println(v3)
                                                            fmt.Println(v45, v45)
                                                        }
                                                    }
                                                }
                                                if len(v1) > 0 {
                                                    v1[0] = (v45 * (-(v45)))
                                                }
                                            } else {
                                                for fi52 := 0; fi52 < 16; fi52++ {
                                                    var v53 int = 904
                                                    v4 += v4
                                                    _ = fi52
                                                    _ = v53
                                                }
                                            }
                                            _ = v45
                                        }
                                        if (!true) {
                                            for fi54 := 0; fi54 < 14; fi54++ {
                                                var v55 string = "\n\t"
                                                _ = fi54
                                                _ = v55
                                            }
                                            var v56 uint8 = 0
                                            _ = v56
                                        }
                                    }
                                } else {
                                    if false {
                                        for fi57 := 0; fi57 < 15; fi57++ {
                                            var v58 float32 = (1.5 * ((((1e10 + float32(0.0)) * 0.0) * 3.14159) - (3.14159 * 3.14159)))
                                            fmt.Println(fi42)
                                            _ = len(v1)
                                            _ = fi57
                                            _ = v58
                                        }
                                        var v59 uint8 = 2
                                        fmt.Println(v59, fi42, v3)
                                        _ = v59
                                    }
                                }
                                var v60 uint64 = (v4 / (((v4 / ((2) | 1))) | 1))
                                var v61 uint = 2
                                _ = fi42
                                _ = v60
                                _ = v61
                            }
                            v62 := GenS3{F0: 100}
                            _ = v62
                        } else {
                            switch "hello" {
                            case "":
                                v63 := GenS9{F0: -1, F1: "a"}
                                _ = v63
                            }
                            v64 := GenS8{F0: (^(v4))}
                            _ = v64
                        }
                        switch 127 {
                        case -1:
                            v39 += v39
                        case 556:
                            var v65 uint = 65535
                            _ = v65
                        case 0:
                            v39 = float64(0.0)
                        default:
                            var v66 uint = 2
                            _ = v66
                        }
                    }
                }
                fmt.Println(v4, v3)
            }
        }
    case -367:
        var v67 float32 = -2.25
        _ = v67
    }
    fmt.Println(v3, v39)
    v3 = v3
    v4 += (v4 % (((v4 * v4)) | 1))
    v68 := GenS8{F0: (v4 >> ((255) & 63))}
    v69 := GenS6{F0: 127, F1: 2}
    fmt.Println(v4, v3, v39)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v37
    _ = v39
    _ = v4
    _ = v40
    _ = v68
    _ = v69
}
