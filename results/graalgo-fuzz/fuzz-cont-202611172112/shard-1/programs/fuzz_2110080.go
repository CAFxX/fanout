package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int16 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8; F1 uint8 }
type GenS7 struct { F0 bool; F1 uint32; F2 int32 }
type GenS8 struct { F0 int; F1 int16; F2 uint }
func genFn0(p0 int32) int64 {
    v6 := []int{}
    p0 -= (-(p0))
    _ = p0
    _ = v6
    return -524
}

func main() {
    v1 := []string{((("foo bar" + "\n\t") + ("\n\t" + "hello")) + "hello"), ((("a" + "a") + ((("" + "foo bar") + "\n\t") + "hello")) + (("" + "foo bar") + "\n\t")), (("\n\t" + "") + ("a" + "foo bar")), "hello"}
    var v2 float32 = (0.1 / 1.0)
    var v3 float64 = ((0.1 - 1e10) - 0.1)
    var v4 int16 = -976
    var v5 int = 128
    v3 -= (0.0 * 1.5)
    v5 = -386
    var v7 float64 = v3
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v5 += v5
    switch v5 {
    case 42:
        v4 -= 1000
    case 2147483647:
        if len(v1) > 0 {
            v1[0] = ("\n\t" + "hello")
        }
    }
    v1 = append(v1, "foo bar")
    _ = len(v1)
    switch (("\n\t" + "a") + "") {
    case "a":
        var v8 uint64 = 18446744073709551615
        _ = v8
    case "\n\t":
        for fi9 := 0; fi9 < 11; fi9++ {
            for fi10 := 0; fi10 < 1; fi10++ {
                _ = len(v1)
                v2 += ((3.14159 / 0.1) / ((v2 / (3.14159 - v2)) - v2))
                switch ("\n\t" + "") {
                case "\n\t":
                    v3 -= 0.0
                case "foo bar":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "hello":
                    var v11 float32 = (v2 / -2.25)
                    _ = v11
                default:
                    var v12 int = v5
                    var v13 float64 = (v7 / v3)
                    _ = v12
                    _ = v13
                }
                if (42 > 42) {
                    if (!false) {
                        if (true || true) {
                            fmt.Println(v4, v2)
                            if (false && true) {
                                if false {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = len(v1)
                                    v5 -= fi10
                                } else {
                                    var v14 string = ("\n\t" + "a")
                                    switch "a" {
                                    case "hello":
                                        var v15 string = v14
                                        _ = v15
                                    default:
                                        var v16 uint32 = 42
                                        _ = v16
                                    }
                                    _ = v14
                                }
                                fmt.Println(fi10, v7, v5)
                            }
                        } else {
                            switch v5 {
                            case 4294967295:
                                v1 = append(v1, "")
                            case 945:
                                v17 := GenS1{F0: 0, F1: (false || ("" <= "hello")), F2: 1000}
                                _ = v17
                            default:
                                var v18 int16 = (v4 >> ((0) & 63))
                                _ = v18
                            }
                            v7 += (v3 + v7)
                        }
                    } else {
                        fmt.Println(fi10, fi9)
                    }
                    var v19 bool = true
                    var v20 uint = 670
                    _ = v19
                    _ = v20
                } else {
                    v3 = float64((1.5 + v2))
                    var v21 uint8 = 132
                    _ = v21
                }
                if (277 >= 323) {
                    fmt.Println(fi10, v4)
                    if (genFn0(-961) != -299) {
                        v1 = append(v1, "foo bar")
                        v3 = v3
                    } else {
                        if true {
                            var v22 int = fi9
                            v7 = 1.5
                            _ = v22
                        }
                    }
                    var v23 int32 = 2147483647
                    _ = v23
                }
                v24 := GenS1{F0: v4, F1: false, F2: (^(v4))}
                v5 = fi10
                _ = fi10
                _ = v24
            }
            v7 -= (3.14159 / (v7 / v3))
            switch "" {
            case "\n\t":
                fmt.Println(fi9)
            case "hello":
                switch "a" {
                case "foo bar":
                    v1 = append(v1, (("\n\t" + "foo bar") + "a"))
                }
            case "":
                if (false || false) {
                    v7 -= (((1.5 / 0.1) / (((3.14159 / v3) / -2.25) - (v3 - 1e10))) - (((v3 * v3) - v3) - v3))
                }
            }
            _ = fi9
        }
    default:
        if (702 != 2) {
            var v25 uint64 = 1
            v26 := GenS1{F0: 1, F1: (true && (v5 < int(v25))), F2: int16(v5)}
            var v27 int32 = int32(v5)
            _ = v25
            _ = v26
            _ = v27
        } else {
            switch "" {
            case "foo bar":
                switch "" {
                case "foo bar":
                    switch (v5 & (v5 % ((127) | 1))) {
                    case 42:
                        var v28 string = ((("hello" + "hello") + ("a" + "foo bar")) + (("" + "hello") + "hello"))
                        _ = v28
                    }
                }
            case "\n\t":
                v29 := GenS1{F0: v4, F1: false, F2: int16(v5)}
                _ = v29
            case "a":
                var v30 int64 = genFn0(2147483647)
                _ = v30
            default:
                switch "hello" {
                case "foo bar":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "hello":
                    v31 := GenS2{F0: 1, F1: 2147483647}
                    _ = v31
                case "":
                    if ((!true) || true) {
                        v32 := GenS1{F0: (v4 + (v4 & 887)), F1: true, F2: -417}
                        _ = v32
                    }
                default:
                    if true {
                        if true {
                            v33 := GenS2{F0: 0, F1: 1}
                            _ = v33
                        } else {
                            var v34 int32 = 42
                            v5 -= 941
                            _ = v34
                        }
                        v35 := GenS1{F0: -32768, F1: (42 != 285), F2: (^(v4))}
                        _ = v35
                    } else {
                        var v36 float64 = float64((0.0 + (float32((v3 / v7)) + (v2 / float32(float64(v2))))))
                        _ = v36
                    }
                }
            }
        }
    }
    fmt.Println(v7)
    v37 := GenS1{F0: (v4 ^ 32767), F1: true, F2: (v4 & v4)}
    if true {
        if (127 != 127) {
            var v38 bool = true
            _ = v38
        }
        v4 = v4
        for fi39 := 0; fi39 < 8; fi39++ {
            v3 = v7
            fmt.Println(fi39, fi39, v3)
            _ = fi39
        }
    }
    if false {
        v40 := GenS2{F0: 255, F1: int32(v4)}
        if false {
            if true {
                for fi41 := 0; fi41 < 4; fi41++ {
                    var v42 string = ("" + ("" + ("foo bar" + (("hello" + ("hello" + ("foo bar" + "foo bar"))) + ("a" + "hello")))))
                    v43 := GenS2{F0: 1, F1: -560}
                    v44 := GenS3{F0: 42}
                    _ = fi41
                    _ = v42
                    _ = v43
                    _ = v44
                }
                switch v5 {
                case -9223372036854775808:
                    var v45 uint8 = 128
                    _ = v45
                }
                if (2 <= 2) {
                    for fi46 := 0; fi46 < 19; fi46++ {
                        var v47 int64 = -9223372036854775808
                        _ = fi46
                        _ = v47
                    }
                }
            } else {
                for fi48 := 0; fi48 < 14; fi48++ {
                    var v49 uint64 = 1
                    v50 := GenS1{F0: (v4 / ((v4) | 1)), F1: (true && true), F2: v4}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi48
                    _ = v49
                    _ = v50
                }
            }
        }
        if false {
            var v51 uint32 = 938
            if false {
                if (18446744073709551615 >= 42) {
                    if true {
                        if (true || (95 >= 42)) {
                            var v52 uint = 713
                            _ = v52
                        } else {
                            var v53 int = v5
                            if (128 >= 1) {
                                v2 = -2.25
                                v1 = append(v1, "")
                                v54 := GenS4{F0: 1}
                                _ = v54
                            } else {
                                fmt.Println(v53, v2)
                            }
                            _ = v53
                        }
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v55 float32 = 3.14159
                        for fi56 := 0; fi56 < 12; fi56++ {
                            switch (fi56 + 32767) {
                            case -166:
                                switch "" {
                                case "a":
                                    v57 := GenS5{F0: v4}
                                    _ = v57
                                case "hello":
                                    switch v5 {
                                    case -727:
                                        if false {
                                            v58 := GenS1{F0: v4, F1: (-448 <= 32768), F2: (-(v4))}
                                            if false {
                                                fmt.Println(v4, fi56, v5)
                                            } else {
                                                v4 = int16(v51)
                                                fmt.Println(v51, v4, v51)
                                                var v59 uint16 = 2
                                                _ = v59
                                            }
                                            var v60 int64 = 9223372036854775807
                                            _ = v58
                                            _ = v60
                                        }
                                    }
                                case "foo bar":
                                    fmt.Println(v4)
                                default:
                                    v1 = append(v1, "\n\t")
                                }
                            default:
                                fmt.Println(v3, v4, v4)
                            }
                            var v61 bool = ((0 != int64(v4)) || (int(v51) != v5))
                            _ = fi56
                            _ = v61
                        }
                        _ = v55
                    }
                }
                var v62 uint16 = 2
                fmt.Println(v62, v5)
                _ = v62
            } else {
                for fi63 := 0; fi63 < 18; fi63++ {
                    if true {
                        v64 := GenS6{F0: 40, F1: 255}
                        for fi65 := 0; fi65 < 17; fi65++ {
                            if false {
                                var v66 uint16 = 1
                                _ = v66
                            } else {
                                _ = len(v1)
                                fmt.Println(v2, fi63, v3)
                                v7 = v7
                            }
                            v2 = v2
                            if true {
                                v5 = (fi65 * (fi65 * fi65))
                                var v67 uint16 = 65535
                                _ = v67
                            } else {
                                v7 = 1.5
                                var v68 uint = 2
                                _ = v68
                            }
                            _ = fi65
                        }
                        if true {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            v69 := GenS1{F0: v4, F1: false, F2: -1}
                            _ = v69
                        }
                        _ = v64
                    } else {
                        if true {
                            v70 := GenS1{F0: int16(v51), F1: false, F2: 0}
                            fmt.Println(v3, v7, fi63)
                            _ = v70
                        }
                    }
                    if len(v1) > 0 {
                        v1[0] = "hello"
                    }
                    var v71 bool = (!true)
                    _ = fi63
                    _ = v71
                }
                if (43 <= 18446744073709551615) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if ("" < "") {
                        if true {
                            v72 := GenS7{F0: (!false), F1: 509, F2: -108}
                            v2 = 1e10
                            switch "\n\t" {
                            case "\n\t":
                                v3 = v3
                            case "hello":
                                if len(v1) > 0 {
                                    v1[0] = ""
                                }
                            case "":
                                v73 := GenS8{F0: (v5 + (v5 ^ (v5 >> ((665) & 63)))), F1: (v4 << ((65535) & 63)), F2: 1}
                                _ = v73
                            }
                            _ = v72
                        } else {
                            for fi74 := 0; fi74 < 1; fi74++ {
                                v7 = float64(v2)
                                for fi75 := 0; fi75 < 8; fi75++ {
                                    _ = len(v1)
                                    if len(v1) > 0 {
                                        v1[0] = ("hello" + "a")
                                    }
                                    _ = fi75
                                }
                                _ = fi74
                            }
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v2 = float32((((1.5 + 1.5) + v7) * v3))
                    }
                    var v76 int32 = 1
                    _ = v76
                } else {
                    _ = len(v1)
                }
                var v77 int32 = 42
                _ = v77
            }
            for fi78 := 0; fi78 < 6; fi78++ {
                v4 -= (-(v4))
                _ = fi78
            }
            _ = v51
        } else {
            var v79 int32 = 42
            v3 += 1e10
            for fi80 := 0; fi80 < 3; fi80++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi80
            }
            _ = v79
        }
        _ = v40
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v7
}
