package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint8; F1 uint; F2 string }
type GenS4 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint32 }
func genFn0(p0 int, p1 int, p2 uint8) int32 {
    p2 += (p2 % ((p2) | 1))
    var v5 float64 = ((3.14159 / (1e10 / 1e10)) / -2.25)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return int32(p2)
}
func genFn1(p0 bool, p1 uint8) string {
    var v6 uint8 = (p1 + p1)
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := []int{(fi8 + (fi8 >> ((924) & 63))), fi7, (-(fi8))}
            var v10 uint32 = 42
            v11 := GenS1{F0: "foo bar", F1: -32768, F2: (v10 * v10)}
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        for fi12 := 0; fi12 < 12; fi12++ {
            var v13 int32 = -2147483648
            _ = fi12
            _ = v13
        }
        if (-32768 < 32767) {
            for fi14 := 0; fi14 < 16; fi14++ {
                if (("" + "hello") < "") {
                    fmt.Println(p0)
                    v15 := []int{}
                    switch -103 {
                    case 32767:
                        p0 = p0
                        fmt.Println(v6, v6)
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                    case 9223372036854775807:
                        var v16 uint16 = 0
                        _ = v16
                    }
                    _ = v15
                } else {
                    v17 := GenS2{F0: -122, F1: 0}
                    switch 127 {
                    case 9223372036854775807:
                        if p0 {
                            var v18 int8 = -1
                            switch (fi14 & fi14) {
                            case 0:
                                switch "" {
                                case "foo bar":
                                    if (!(!((("hello" + "") + ("\n\t" + "hello")) == ""))) {
                                        var v19 int = 1
                                        _ = v19
                                    }
                                default:
                                    fmt.Println(fi7, p0)
                                }
                                fmt.Println(p0, fi14, p0)
                                v20 := []int{fi14, (fi14 & (fi7 / ((fi14) | 1))), 42, 278}
                                _ = v20
                            default:
                                p1 = 0
                            }
                            if (!(!(p0 || (4294967295 == 496)))) {
                                var v21 float32 = ((3.14159 - (0.1 + 0.1)) / (-2.25 * -2.25))
                                var v22 uint16 = 1
                                _ = v21
                                _ = v22
                            } else {
                                v23 := GenS3{F0: (^(p1)), F1: 729, F2: "foo bar"}
                                _ = v23
                            }
                            _ = v18
                        } else {
                            v24 := GenS4{F0: 251, F1: p1, F2: 0}
                            _ = v24
                        }
                    }
                    v25 := []int{fi7, fi7, (fi14 ^ fi14), fi14}
                    _ = v17
                    _ = v25
                }
                p0 = (false || true)
                _ = fi14
            }
        }
        _ = fi7
    }
    v26 := GenS2{F0: -32768, F1: 945}
    p1 = v6
    fmt.Println(v6, p0)
    _ = p0
    _ = p1
    _ = v26
    _ = v6
    return "hello"
}

func main() {
    var v1 uint8 = 2
    v2 := make(map[uint32]string)
    var v3 int8 = 0
    v4 := []int{0, -9223372036854775808, -793, -9223372036854775808}
    if false {
        _ = len(v2)
        fmt.Println(v1, v3)
        v1 += (v1 / ((42) | 1))
    } else {
        v3 = 0
    }
    switch "a" {
    case "\n\t":
        var v27 uint64 = 673
        _ = v27
    }
    v3 = 1
    for fi28 := 0; fi28 < 4; fi28++ {
        var v29 int16 = -32768
        if false {
            var v30 uint8 = v1
            _ = v30
        }
        _ = fi28
        _ = v29
    }
    v31 := GenS3{F0: v1, F1: 580, F2: ""}
    switch "\n\t" {
    case "":
        v2[153] = "\n\t"
    case "a":
        var v32 float32 = 1.5
        _ = v32
    case "foo bar":
        _ = len(v2)
    default:
        for fi33 := 0; fi33 < 10; fi33++ {
            if true {
                v34 := GenS4{F0: 0, F1: 128, F2: 82}
                for fi35 := 0; fi35 < 4; fi35++ {
                    v3 += -1
                    var v36 uint64 = 42
                    if false {
                        fmt.Println(v36, fi33, v36)
                    }
                    _ = fi35
                    _ = v36
                }
                _ = v34
            } else {
                switch fi33 {
                case 591:
                    if len(v4) > 0 {
                        v4[0] = fi33
                    }
                case -2147483648:
                    var v37 string = "hello"
                    _ = v37
                default:
                    v38 := GenS5{F0: 0}
                    _ = v38
                }
                fmt.Println(v3, v3, v1)
                switch fi33 {
                case 1:
                    var v39 uint32 = 0
                    _ = v39
                case 674:
                    for fi40 := 0; fi40 < 20; fi40++ {
                        fmt.Println(v3, v3, fi40)
                        v1 -= (^(v1))
                        _ = fi40
                    }
                case 0:
                    fmt.Println(fi33, v1)
                }
            }
            v3 -= (v3 & v3)
            var v41 int8 = (-(v3))
            _ = fi33
            _ = v41
        }
    }
    v1 = (v1 / (((v1 % ((85) | 1))) | 1))
    if true {
        v42 := GenS6{F0: (v3 & (-(v3)))}
        v3 = v3
        delete(v2, 2147483647)
        _ = v42
    } else {
        if (((0 < 32768) && (20 != 42)) && (!((!(539 == 65535)) && (42 < -76)))) {
            fmt.Println(v3)
        } else {
            if true {
                switch 9223372036854775807 {
                case 0:
                    if false {
                        if false {
                            if (false && true) {
                                _ = len(v4)
                            }
                            v43 := GenS1{F0: ("\n\t" + ("a" + (("\n\t" + (("\n\t" + "a") + "a")) + "\n\t"))), F1: 42, F2: 42}
                            _ = v43
                        } else {
                            v44 := GenS5{F0: 42}
                            _ = v44
                        }
                        var v45 int8 = 1
                        _ = v45
                    } else {
                        for fi46 := 0; fi46 < 4; fi46++ {
                            if ((1e10 + 3.14159) < 3.14159) {
                                v3 = v3
                                switch (("a" + "\n\t") + "") {
                                case "foo bar":
                                    if ("hello" <= "foo bar") {
                                        v1 += v1
                                        var v47 int = (fi46 * (fi46 & 1))
                                        _ = v47
                                    }
                                case "\n\t":
                                    var v48 uint8 = v1
                                    _ = v48
                                }
                            }
                            _ = fi46
                        }
                    }
                case -9223372036854775808:
                    if true {
                        switch -9223372036854775808 {
                        case -9223372036854775808:
                            switch 391 {
                            case 32767:
                                v1 -= (v1 - (v1 * (v1 / ((v1) | 1))))
                            default:
                                v1 += 128
                            }
                        }
                        var v49 uint16 = 661
                        v50 := GenS4{F0: (v49 + v49), F1: v1, F2: 0}
                        _ = v49
                        _ = v50
                    }
                }
                var v51 float32 = 0.0
                _ = v51
            } else {
                switch "a" {
                case "a":
                    for fi52 := 0; fi52 < 7; fi52++ {
                        v3 -= 42
                        v53 := v2[2147483647]
                        _ = fi52
                        _ = v53
                    }
                case "":
                    fmt.Println(v3, v1)
                }
                v54 := GenS7{F0: 2}
                _ = v54
            }
            if true {
                _ = len(v2)
                v1 -= 255
            }
        }
        switch 32767 {
        case 1:
            v2[826] = ""
        case -615:
            v1 -= (^(v1))
        }
    }
    for fi55 := 0; fi55 < 9; fi55++ {
        v56 := GenS1{F0: "", F1: 0, F2: 2147483647}
        fmt.Println(v3, v3)
        _ = fi55
        _ = v56
    }
    fmt.Println(v1, v3)
    fmt.Println(v3, v3)
    for fi57 := 0; fi57 < 3; fi57++ {
        if len(v4) > 0 {
            v4[0] = 127
        }
        _ = fi57
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
}
