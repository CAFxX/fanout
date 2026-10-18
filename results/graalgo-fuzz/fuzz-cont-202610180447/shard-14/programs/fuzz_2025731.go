package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint }
type GenS2 struct { F0 uint; F1 uint32; F2 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 int8; F1 string; F2 int32 }
type GenS6 struct { F0 int32; F1 bool; F2 int32 }
type GenS7 struct { F0 bool }
func genFn0() int {
    if (42 == -963) {
        if false {
            var v7 int = 42
            v8 := []int{}
            _ = v7
            _ = v8
        } else {
            var v9 float32 = ((-2.25 - 3.14159) * 1e10)
            v10 := []int{128, 32768, -414}
            _ = v10
            _ = v9
        }
        v11 := GenS1{F0: 127, F1: 255, F2: 2}
        _ = v11
    }
    switch -1 {
    case 9223372036854775807:
        v12 := GenS2{F0: 18446744073709551615, F1: 0, F2: ("" + "foo bar")}
        _ = v12
    }
    v13 := []int{9223372036854775807, 0, -517}
    for fi14 := 0; fi14 < 2; fi14++ {
        switch "a" {
        case "a":
            for fi15 := 0; fi15 < 2; fi15++ {
                if ((0 == (fi15 / ((fi14) | 1))) && true) {
                    var v16 float64 = (1e10 / (0.0 - float64(1.5)))
                    v17 := GenS2{F0: 1, F1: 337, F2: "\n\t"}
                    var v18 int8 = 126
                    _ = v16
                    _ = v17
                    _ = v18
                } else {
                    fmt.Println(fi14)
                    var v19 uint16 = 0
                    var v20 uint = 18446744073709551615
                    _ = v19
                    _ = v20
                }
                _ = fi15
            }
        case "foo bar":
            switch "a" {
            case "a":
                for fi21 := 0; fi21 < 6; fi21++ {
                    if len(v13) > 0 {
                        v13[0] = (^(fi21))
                    }
                    _ = fi21
                }
            case "foo bar":
                for fi22 := 0; fi22 < 13; fi22++ {
                    var v23 int8 = 0
                    _ = fi22
                    _ = v23
                }
            case "hello":
                fmt.Println(fi14)
            default:
                for fi24 := 0; fi24 < 7; fi24++ {
                    var v25 string = "hello"
                    _ = fi24
                    _ = v25
                }
            }
        default:
            var v26 uint64 = 42
            _ = v26
        }
        _ = fi14
    }
    var v27 uint16 = 1
    _ = v13
    _ = v27
    return -9223372036854775808
}
func genFn1(p0 int64) int32 {
    var v28 uint8 = 37
    v28 = v28
    _ = p0
    _ = v28
    return -2147483648
}

func main() {
    var v1 string = "\n\t"
    var v2 float32 = ((((1.5 / (float32(((1e10 * float64(0.1)) * 0.1)) - (0.0 + -2.25))) - -2.25) / (1.5 / 1.0)) / 3.14159)
    var v3 uint64 = 0
    var v4 uint = 65535
    var v5 uint = v4
    v6 := make(map[int]bool)
    delete(v6, 274)
    for fi29 := 0; fi29 < 17; fi29++ {
        v30 := GenS3{F0: uint(v3)}
        for fi31 := 0; fi31 < 16; fi31++ {
            v32 := GenS4{F0: "foo bar", F1: -419}
            _ = fi31
            _ = v32
        }
        for fi33 := 0; fi33 < 2; fi33++ {
            if ((!(9223372036854775807 <= -9223372036854775808)) && true) {
                v3 = 2
            } else {
                var v34 uint16 = 255
                _ = len(v6)
                _ = v34
            }
            v35 := GenS5{F0: 111, F1: v1, F2: -2147483648}
            var v36 int8 = -1
            _ = fi33
            _ = v35
            _ = v36
        }
        _ = fi29
        _ = v30
    }
    var v37 uint16 = 1
    var v38 int64 = -953
    v1 = "\n\t"
    for fi39 := 0; fi39 < 4; fi39++ {
        v37 = (v37 << (((v4 ^ (v5 >> ((uint(v37)) & 63)))) & 63))
        if true {
            if false {
                if (!true) {
                    for fi40 := 0; fi40 < 11; fi40++ {
                        if true {
                            var v41 uint8 = 1
                            _ = v41
                        }
                        var v42 int32 = -2147483648
                        _ = len(v6)
                        _ = fi40
                        _ = v42
                    }
                } else {
                    var v43 uint32 = 460
                    _ = len(v6)
                    _ = v43
                }
                v44 := v6[fi39]
                _ = v44
            }
            v45 := GenS6{F0: 2147483647, F1: false, F2: -1}
            for fi46 := 0; fi46 < 6; fi46++ {
                if false {
                    v37 = v37
                    delete(v6, -9223372036854775808)
                    v4 = v4
                    var v47 float32 = v2
                    v5 = v5
                    _ = v47
                } else {
                    v48 := []int{(fi46 & fi46)}
                    v49 := GenS7{F0: false}
                    _ = v48
                    _ = v49
                }
                _ = fi46
            }
            _ = v45
        } else {
            fmt.Println(v5, v1)
            var v50 uint64 = 2
            _ = v50
        }
        _ = fi39
    }
    fmt.Println(v4, v3, v3)
    if false {
        v51 := []int{127}
        _ = v51
    } else {
        var v52 uint32 = uint32(v5)
        fmt.Println(v4, v52, v1)
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
