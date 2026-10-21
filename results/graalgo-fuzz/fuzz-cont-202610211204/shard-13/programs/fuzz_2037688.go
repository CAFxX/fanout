package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint32, p1 uint16) int64 {
    fmt.Println(p1, p1)
    v5 := GenS1{F0: (!true), F1: uint16(p0)}
    _ = p0
    _ = p1
    _ = v5
    return -731
}
func genFn1(p0 uint32, p1 int) uint64 {
    if (false || (false || (42 <= 399))) {
        v6 := GenS1{F0: (32767 <= -571), F1: 0}
        var v7 bool = (1 <= -32768)
        for fi8 := 0; fi8 < 20; fi8++ {
            for fi9 := 0; fi9 < 6; fi9++ {
                fmt.Println(fi9, p0)
                v10 := []int{(fi9 + 127), (fi9 ^ (p1 / ((fi8) | 1))), 42, fi8}
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
        _ = v6
        _ = v7
    } else {
        v11 := []int{(p1 % ((p1) | 1)), -100}
        fmt.Println(p0)
        _ = v11
    }
    var v12 int32 = 42
    var v13 uint8 = 127
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 883
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 2
    var v3 uint64 = 497
    v4 := make(map[string]string)
    switch "" {
    case "foo bar":
        for fi14 := 0; fi14 < 8; fi14++ {
            var v15 uint64 = v3
            if true {
                if true {
                    fmt.Println(v15, v15, v2)
                    if (("" + (("hello" + "") + "hello")) == "\n\t") {
                        v16 := []int{fi14}
                        var v17 uint = 4294967295
                        var v18 uint16 = 42
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                }
            }
            delete(v4, "")
            _ = fi14
            _ = v15
        }
    case "a":
        var v19 uint8 = (v2 | (v2 ^ (^(v2))))
        _ = v19
    }
    v3 = genFn1(42, 32768)
    if false {
        v20 := []int{}
        v21 := GenS1{F0: false, F1: 32767}
        switch 127 {
        case 156:
            fmt.Println(v2, v2)
        case 32768:
            if true {
                v2 -= 56
                v3 = 2
                fmt.Println(v3, v3)
            } else {
                for fi22 := 0; fi22 < 10; fi22++ {
                    var v23 uint32 = 4294967295
                    _ = fi22
                    _ = v23
                }
                switch -2147483648 {
                case 32767:
                    for fi24 := 0; fi24 < 8; fi24++ {
                        var v25 uint64 = v3
                        _ = fi24
                        _ = v25
                    }
                case 32768:
                    v2 = v2
                case -137:
                    var v26 int8 = -21
                    _ = v26
                default:
                    for fi27 := 0; fi27 < 10; fi27++ {
                        v2 = 0
                        _ = len(v1)
                        _ = fi27
                    }
                }
                if ("" <= "hello") {
                    if true {
                        v28 := GenS1{F0: true, F1: 255}
                        _ = v28
                    }
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                } else {
                    switch "\n\t" {
                    case "a":
                        v2 -= v2
                    case "foo bar":
                        if false {
                            for fi29 := 0; fi29 < 8; fi29++ {
                                switch 42 {
                                case 413:
                                    fmt.Println(fi29)
                                case -2147483648:
                                    var v30 int64 = 42
                                    _ = v30
                                case 1:
                                    fmt.Println(v2, fi29, v3)
                                default:
                                    if len(v20) > 0 {
                                        v20[0] = (fi29 + 4294967295)
                                    }
                                }
                                if ((^(v3)) <= v3) {
                                    if len(v20) > 0 {
                                        v20[0] = fi29
                                    }
                                } else {
                                    v31 := GenS1{F0: (false && ((false || false) && false)), F1: 32767}
                                    v4[("\n\t" + "\n\t")] = "foo bar"
                                    delete(v1, 161)
                                    _ = v31
                                }
                                for fi32 := 0; fi32 < 11; fi32++ {
                                    v33 := v4["a"]
                                    if true {
                                        var v34 uint16 = 948
                                        v35 := GenS2{F0: (!(32767 <= -819))}
                                        _ = v34
                                        _ = v35
                                    } else {
                                        fmt.Println(fi29, v33)
                                        v2 = v2
                                    }
                                    v33 = ""
                                    _ = fi32
                                    _ = v33
                                }
                                _ = fi29
                            }
                            var v36 string = "foo bar"
                            var v37 uint16 = 2
                            _ = v36
                            _ = v37
                        }
                    }
                    v20 = append(v20, 4294967295)
                }
            }
        case -1:
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
        default:
            switch 32767 {
            case 637:
                var v38 float64 = (0.0 * 0.0)
                _ = v38
            case 9223372036854775807:
                v39 := GenS3{F0: 18446744073709551615}
                _ = v39
            case -9223372036854775808:
                var v40 float64 = 1e10
                _ = v40
            }
        }
        _ = v20
        _ = v21
    }
    var v41 float32 = (0.1 + float32((-2.25 * -2.25)))
    var v42 bool = (genFn0(uint32(v3), 0) == -1)
    v42 = v42
    fmt.Println(v42)
    delete(v4, "")
    if ((3.14159 * 1e10) >= (((0.0 - float64(v41)) + 3.14159) - -2.25)) {
        if v42 {
            v42 = v42
            if true {
                for fi43 := 0; fi43 < 3; fi43++ {
                    var v44 int32 = 2147483647
                    _ = fi43
                    _ = v44
                }
                v3 = v3
                var v45 int8 = -73
                _ = v45
            } else {
                if v42 {
                    v41 -= (3.14159 - v41)
                } else {
                    v42 = v42
                    v46 := GenS3{F0: 65535}
                    _ = v46
                }
                fmt.Println(v42)
                if ((("hello" + ("a" + (("hello" + "") + (("hello" + "hello") + "a")))) + ("hello" + ("hello" + ("a" + (("" + "a") + (("a" + ("a" + "\n\t")) + "\n\t")))))) <= "a") {
                    switch "a" {
                    case "hello":
                        var v47 float64 = 1.5
                        _ = v47
                    default:
                        var v48 int8 = 42
                        _ = v48
                    }
                    v2 = (v2 * 127)
                }
            }
        }
        v49 := v1[295]
        _ = v49
    } else {
        v50 := GenS1{F0: v42, F1: 388}
        if v42 {
            var v51 int64 = -923
            _ = v51
        } else {
            if true {
                var v52 float32 = 1.5
                var v53 int16 = -32768
                if v42 {
                    v3 += v3
                }
                _ = v52
                _ = v53
            }
        }
        v54 := GenS4{F0: 42}
        _ = v50
        _ = v54
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
}
