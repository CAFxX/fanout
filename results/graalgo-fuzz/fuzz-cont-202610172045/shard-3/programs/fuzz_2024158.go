package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 int; F1 int; F2 uint64 }
func genFn0(p0 int, p1 int) int8 {
    var v7 uint8 = 2
    v7 -= v7
    v8 := GenS1{F0: false, F1: 42}
    v9 := []int{}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return -1
}
func genFn1(p0 uint, p1 int16) bool {
    if false {
        v10 := GenS1{F0: (!((p1 >> (((p0 / ((0) | 1))) & 63)) < p1)), F1: 2147483647}
        fmt.Println(p0, p0)
        var v11 int64 = 9223372036854775807
        _ = v10
        _ = v11
    } else {
        fmt.Println(p0)
        var v12 uint16 = 0
        _ = v12
    }
    switch ("foo bar" + (("a" + ((("hello" + ("a" + "\n\t")) + "hello") + ("hello" + "hello"))) + "\n\t")) {
    case "a":
        p0 = 261
    case "hello":
        var v13 uint16 = 435
        _ = v13
    case "\n\t":
        var v14 int = 128
        _ = v14
    default:
        if (false && false) {
            var v15 int32 = 0
            _ = v15
        }
    }
    if (!true) {
        if (-2.25 >= 3.14159) {
            p1 += p1
            fmt.Println(p0)
        }
    }
    _ = p0
    _ = p1
    return false
}

func main() {
    var v1 float64 = 1.5
    v2 := make(map[string]string)
    v3 := []int{}
    var v4 int32 = 1
    var v5 int8 = -92
    v6 := make(map[uint32]int)
    v16 := GenS1{F0: (-30 > 0), F1: -665}
    _ = len(v3)
    fmt.Println(v1, v4, v1)
    switch 43 {
    case -1:
        v4 += (v4 >> ((255) & 63))
    case -238:
        fmt.Println(v5, v1, v4)
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 uint = 4294967295
        v1 = (3.14159 / v1)
        v4 = (v4 ^ v4)
        _ = fi17
        _ = v18
    }
    v19 := GenS2{F0: 250, F1: 493}
    var v20 bool = (false || (481 >= 547))
    if true {
        if true {
            v5 = v5
            for fi21 := 0; fi21 < 10; fi21++ {
                if v20 {
                    var v22 int64 = -1
                    _ = v22
                }
                for fi23 := 0; fi23 < 3; fi23++ {
                    var v24 bool = false
                    _ = fi23
                    _ = v24
                }
                v25 := GenS1{F0: v20, F1: v4}
                _ = fi21
                _ = v25
            }
            if (v20 && false) {
                var v26 bool = true
                v26 = (v5 <= v5)
                v4 += (v4 / (((v4 >> ((18446744073709551615) & 63))) | 1))
                _ = v26
            } else {
                v20 = v20
                var v27 int16 = -141
                v3 = append(v3, 370)
                _ = v27
            }
        } else {
            for fi28 := 0; fi28 < 9; fi28++ {
                for fi29 := 0; fi29 < 6; fi29++ {
                    if v20 {
                        v4 = (v4 >> ((18446744073709551615) & 63))
                    }
                    var v30 uint = 65535
                    _ = fi29
                    _ = v30
                }
                _ = fi28
            }
            _ = len(v6)
            var v31 uint = 42
            _ = v31
        }
        _ = len(v3)
    } else {
        if (false && (1 < 293)) {
            var v32 int16 = 42
            _ = v32
        } else {
            switch "" {
            case "":
                v1 += 1.5
            case "\n\t":
                v33 := v2[("a" + "foo bar")]
                _ = v33
            case "a":
                if false {
                    var v34 float32 = (3.14159 / float32(v1))
                    switch 2147483647 {
                    case 414:
                        v3 = append(v3, 42)
                    case 1:
                        var v35 float64 = float64(((3.14159 * float32((((-2.25 + ((float64(v34) + (v1 - v1)) + v1)) / ((0.0 * (v1 / (v1 * 3.14159))) - v1)) * 0.1))) * v34))
                        _ = v35
                    case 9223372036854775807:
                        for fi36 := 0; fi36 < 19; fi36++ {
                            v5 = genFn0(-395, 42)
                            for fi37 := 0; fi37 < 4; fi37++ {
                                if (1000 == 1) {
                                    v38 := GenS1{F0: (v20 || (false && (("\n\t" == "\n\t") || (12 <= 0)))), F1: (v4 << ((255) & 63))}
                                    switch ("" + "foo bar") {
                                    case "a":
                                        v4 = 795
                                    case "hello":
                                        var v39 uint = 65535
                                        _ = v39
                                    case "foo bar":
                                        var v40 bool = true
                                        _ = v40
                                    default:
                                        v41 := GenS3{F0: (fi36 + 42), F1: -127, F2: 1}
                                        _ = v41
                                    }
                                    v34 -= float32(((float64(0.0) / v1) * 0.0))
                                    _ = v38
                                } else {
                                    _ = len(v3)
                                    var v42 int64 = -696
                                    var v43 string = "a"
                                    _ = v42
                                    _ = v43
                                }
                                _ = fi37
                            }
                            v44 := GenS1{F0: v20, F1: (-(v4))}
                            _ = fi36
                            _ = v44
                        }
                    }
                    var v45 bool = v20
                    _ = v34
                    _ = v45
                }
            }
        }
    }
    v46 := GenS2{F0: 962, F1: 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
