package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint32 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 uint8; F1 int8; F2 uint8 }
func genFn0(p0 int64, p1 int16) uint16 {
    v5 := []int{127, 0, 9223372036854775807}
    var v6 string = "foo bar"
    v6 = v6
    var v7 bool = true
    v8 := GenS1{F0: -88, F1: p1, F2: 249}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 693
}
func genFn1(p0 int64) int {
    v9 := []int{-9223372036854775808, -9223372036854775808}
    var v10 string = (("\n\t" + ((((("foo bar" + "a") + "") + "\n\t") + "a") + "a")) + "a")
    _ = p0
    _ = v10
    _ = v9
    return 128
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 2
    v3 := []string{("\n\t" + "\n\t")}
    var v4 float64 = 3.14159
    v2 = (^(v2))
    var v11 uint16 = 0
    v2 = (v2 + v2)
    var v12 uint8 = 255
    var v13 uint16 = v11
    _ = len(v1)
    if true {
        var v14 bool = true
        _ = v14
    } else {
        fmt.Println(v2, v2, v13)
        if (true && false) {
            if (float64(float32(v4)) == 0.1) {
                v15 := GenS1{F0: -27, F1: -1, F2: 0}
                for fi16 := 0; fi16 < 11; fi16++ {
                    fmt.Println(v4, v2, fi16)
                    v13 = (v13 + genFn0(-754, 1000))
                    _ = fi16
                }
                v3 = append(v3, "a")
                _ = v15
            } else {
                var v17 int = genFn1(1)
                _ = v17
            }
        }
        var v18 bool = (44 < (v13 * v13))
        _ = v18
    }
    switch -1 {
    case 127:
        if ("foo bar" < "") {
            v19 := GenS2{F0: ((((0.1 * (0.0 - 1e10)) + -2.25) + 1.5) != 1e10), F1: (false || (((9223372036854775807 != 548) && (uint32(v2) == 1)) || false))}
            _ = v19
        } else {
            v3 = append(v3, "a")
            fmt.Println(v11, v12)
        }
    default:
        if false {
            for fi20 := 0; fi20 < 1; fi20++ {
                v21 := v1["foo bar"]
                if (2147483648 == 0) {
                    v13 -= (v11 - (v11 + v13))
                    if true {
                        for fi22 := 0; fi22 < 6; fi22++ {
                            var v23 float32 = 3.14159
                            if ((true && (false && (v23 == 1e10))) && false) {
                                v12 = (v12 % (((v12 ^ v12)) | 1))
                                if true {
                                    var v24 uint64 = uint64(v13)
                                    fmt.Println(fi22)
                                    _ = len(v1)
                                    _ = v24
                                }
                            } else {
                                fmt.Println(v13, v4, v2)
                            }
                            if len(v3) > 0 {
                                v3[0] = "a"
                            }
                            _ = fi22
                            _ = v23
                        }
                        v11 += v13
                    } else {
                        v11 -= (v11 & genFn0(306, -304))
                    }
                } else {
                    v2 -= v2
                }
                _ = fi20
                _ = v21
            }
            switch "foo bar" {
            case "foo bar":
                if false {
                    if false {
                        for fi25 := 0; fi25 < 5; fi25++ {
                            var v26 uint = 287
                            var v27 bool = (v2 < (v2 / ((v2) | 1)))
                            if len(v3) > 0 {
                                v3[0] = "\n\t"
                            }
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                        v28 := GenS1{F0: -1, F1: 32767, F2: 402}
                        var v29 uint = 903
                        _ = v28
                        _ = v29
                    }
                    var v30 int8 = 0
                    v2 -= 2
                    _ = v30
                } else {
                    var v31 bool = ((42 <= 0) && false)
                    v32 := GenS3{F0: (v12 / ((v12) | 1)), F1: 100, F2: v12}
                    _ = v31
                    _ = v32
                }
            case "hello":
                v33 := GenS1{F0: -31, F1: 42, F2: 65535}
                _ = v33
            default:
                if false {
                    if (v4 == 3.14159) {
                        v34 := GenS3{F0: (^(v12)), F1: 127, F2: (v12 >> ((376) & 63))}
                        _ = v34
                    }
                }
            }
            v35 := GenS1{F0: 127, F1: -32768, F2: 4294967295}
            _ = v35
        }
    }
    if ((uint16(v2) == genFn0(0, 602)) || true) {
        var v36 int16 = -550
        _ = v36
    } else {
        var v37 int16 = 1
        v38 := GenS1{F0: 42, F1: (v37 << ((18446744073709551615) & 63)), F2: 381}
        _ = v37
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
