package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 bool; F2 string }
type GenS4 struct { F0 int64; F1 uint32; F2 string }
type GenS5 struct { F0 uint8 }
func genFn0(p0 int32, p1 int16, p2 int) uint16 {
    var v6 float32 = 0.1
    if (((p2 + (p2 >> ((771) & 63))) > p2) || true) {
        v7 := []int{-711, (p2 << ((836) & 63)), p2, (^(p2))}
        for fi8 := 0; fi8 < 4; fi8++ {
            for fi9 := 0; fi9 < 13; fi9++ {
                if (0.1 == (0.1 * (1e10 + 1.5))) {
                    p2 = fi8
                    switch fi8 {
                    case 128:
                        var v10 float64 = float64(v6)
                        _ = v10
                    case -465:
                        var v11 uint64 = 18446744073709551615
                        _ = v11
                    case 32768:
                        v12 := GenS1{F0: 29, F1: 255}
                        _ = v12
                    }
                }
                v13 := GenS1{F0: 67, F1: 4294967295}
                _ = fi9
                _ = v13
            }
            v6 = (v6 + (((((0.1 / -2.25) - ((float32(float64(v6)) - float32(0.1)) * 0.1)) + 1.5) * (0.1 / v6)) + v6))
            _ = len(v7)
            _ = fi8
        }
        if false {
            switch "\n\t" {
            case "\n\t":
                v7 = append(v7, 32768)
            case "foo bar":
                if false {
                    p1 += -1
                    if false {
                        p2 -= (p2 / ((p2) | 1))
                        v14 := GenS1{F0: -1, F1: 1}
                        _ = v14
                    }
                }
            }
        }
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}
func genFn1(p0 int64) uint32 {
    v15 := GenS2{F0: (p0 << ((0) & 63))}
    v16 := GenS3{F0: 780, F1: (!(509 <= 641)), F2: ((("foo bar" + "") + "") + "foo bar")}
    _ = p0
    _ = v15
    _ = v16
    return 65535
}

func main() {
    var v1 uint16 = 396
    var v2 uint8 = 127
    v3 := []string{("\n\t" + ("a" + "a")), "a", "hello"}
    v4 := []int{4294967295, -261, 141}
    var v5 int32 = 42
    for fi17 := 0; fi17 < 3; fi17++ {
        if (!false) {
            if (false && ((v1 + 2) != 65535)) {
                fmt.Println(v2)
            } else {
                if (false && true) {
                    if (0 <= int8(fi17)) {
                        var v18 uint = 615
                        _ = v18
                    }
                }
            }
        } else {
            fmt.Println(v1, v1)
        }
        _ = fi17
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        if true {
            var v20 bool = (-1 == 42)
            _ = v20
        } else {
            var v21 uint64 = 2
            if (0 >= -9223372036854775808) {
                _ = len(v4)
            }
            switch (^(fi19)) {
            case 4294967295:
                if true {
                    for fi22 := 0; fi22 < 11; fi22++ {
                        var v23 int = 42
                        var v24 int32 = (^(v5))
                        var v25 int8 = -128
                        _ = fi22
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                } else {
                    switch 4294967295 {
                    case 448:
                        fmt.Println(v2, v1)
                    case 128:
                        var v26 uint16 = v1
                        _ = v26
                    case 763:
                        switch (fi19 * fi19) {
                        case 9223372036854775807:
                            v3 = append(v3, "")
                        }
                    default:
                        switch fi19 {
                        case 42:
                            switch "" {
                            case "a":
                                if true {
                                    v3 = append(v3, "\n\t")
                                    if true {
                                        if (true && true) {
                                            var v27 uint32 = 221
                                            var v28 uint = 4294967295
                                            v21 -= 42
                                            _ = v27
                                            _ = v28
                                        } else {
                                            if false {
                                                v2 = (v2 + 42)
                                            }
                                            if false {
                                                v1 = (^(v1))
                                                var v29 uint32 = 4294967295
                                                _ = v29
                                            } else {
                                                if (!(48 > 4294967295)) {
                                                    v5 += -2147483648
                                                    var v30 int16 = -1
                                                    _ = v30
                                                }
                                            }
                                        }
                                    }
                                }
                            default:
                                var v31 string = ("foo bar" + "foo bar")
                                _ = v31
                            }
                        default:
                            if true {
                                v32 := GenS4{F0: 1, F1: 820, F2: ("a" + ((((("" + "\n\t") + (("hello" + "a") + "hello")) + ("a" + ("hello" + "foo bar"))) + ((("a" + "hello") + ("" + ("a" + ("hello" + "hello")))) + "")) + (("foo bar" + "foo bar") + "\n\t")))}
                                fmt.Println(v1, fi19)
                                var v33 uint = 19
                                _ = v32
                                _ = v33
                            }
                        }
                    }
                }
            case 127:
                if len(v3) > 0 {
                    v3[0] = "foo bar"
                }
            case 32767:
                v2 += (v2 & v2)
            default:
                fmt.Println(v1, v1, v21)
            }
            _ = v21
        }
        for fi34 := 0; fi34 < 17; fi34++ {
            var v35 string = "foo bar"
            _ = fi34
            _ = v35
        }
        _ = fi19
    }
    if (!(32767 > -677)) {
        v2 += v2
    } else {
        _ = len(v3)
    }
    v36 := GenS5{F0: (v2 ^ (^(v2)))}
    var v37 uint32 = 1
    if (1e10 < 1e10) {
        switch -229 {
        case 536:
            v37 -= v37
        default:
            v5 -= v5
        }
    } else {
        v38 := GenS5{F0: (v2 + (v2 / ((128) | 1)))}
        if ((v2 | v2) < (v2 * v2)) {
            switch "\n\t" {
            case "foo bar":
                v4 = append(v4, -320)
            case "\n\t":
                var v39 uint = 18446744073709551615
                _ = v39
            case "hello":
                for fi40 := 0; fi40 < 7; fi40++ {
                    if len(v4) > 0 {
                        v4[0] = (fi40 >> ((4294967295) & 63))
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi40
                }
            default:
                var v41 float32 = 1e10
                _ = v41
            }
            v42 := GenS5{F0: 255}
            _ = v42
        }
        var v43 int16 = 42
        _ = v38
        _ = v43
    }
    var v44 int64 = -9223372036854775808
    var v45 uint32 = (v37 & 2)
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v44
    _ = v45
    _ = v5
}
