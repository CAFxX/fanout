package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint }
type GenS4 struct { F0 uint64; F1 uint16; F2 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int; F1 bool; F2 uint }
func genFn0(p0 int, p1 int16, p2 int16) uint8 {
    p2 -= p2
    if (806 != 0) {
        fmt.Println(p1)
    }
    v7 := []int{p0, (-(p0)), p0, (p0 + (^(p0)))}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}
func genFn1(p0 uint64) bool {
    var v8 float64 = (0.0 - 0.0)
    p0 = 2
    v9 := []int{}
    switch ("hello" + "foo bar") {
    case "hello":
        for fi10 := 0; fi10 < 14; fi10++ {
            v8 += 0.1
            p0 += 411
            _ = fi10
        }
    }
    _ = p0
    _ = v8
    _ = v9
    return (-9223372036854775808 <= -84)
}

func main() {
    var v1 string = "foo bar"
    v2 := []uint32{2, 65535}
    v3 := []string{v1, (v1 + v1), v1, v1}
    var v4 string = v1
    var v5 float32 = -2.25
    var v6 int16 = 42
    var v11 float64 = ((0.0 + -2.25) / 1.0)
    var v12 int64 = -9223372036854775808
    v11 = float64((v5 * v5))
    switch (v1 + "hello") {
    case "a":
        var v13 int8 = 0
        _ = v13
    case "hello":
        if (true || true) {
            for fi14 := 0; fi14 < 3; fi14++ {
                _ = len(v2)
                v15 := GenS1{F0: 2, F1: 4294967295}
                _ = fi14
                _ = v15
            }
            if len(v3) > 0 {
                v3[0] = v1
            }
        } else {
            fmt.Println(v12, v11, v1)
            v6 = v6
            if len(v2) > 0 {
                v2[0] = 0
            }
        }
    default:
        var v16 uint32 = 42
        _ = v16
    }
    if false {
        v5 += v5
        v17 := GenS1{F0: 246, F1: 1}
        _ = v17
    }
    if (65535 >= 44) {
        var v18 int8 = -1
        _ = v18
    }
    if false {
        var v19 int16 = v6
        switch v4 {
        case "foo bar":
            fmt.Println(v19)
        default:
            for fi20 := 0; fi20 < 2; fi20++ {
                switch v1 {
                case "hello":
                    for fi21 := 0; fi21 < 4; fi21++ {
                        if (!(true || false)) {
                            v3 = append(v3, ("hello" + v1))
                            if len(v2) > 0 {
                                v2[0] = 0
                            }
                            v22 := GenS2{F0: -1, F1: true}
                            _ = v22
                        }
                        fmt.Println(v6, fi20, v11)
                        _ = fi21
                    }
                case "a":
                    var v23 bool = false
                    _ = v23
                default:
                    for fi24 := 0; fi24 < 18; fi24++ {
                        v25 := GenS3{F0: 201, F1: 576, F2: 1}
                        switch v1 {
                        case "\n\t":
                            switch "foo bar" {
                            case "foo bar":
                                v4 = "foo bar"
                            case "\n\t":
                                v26 := GenS4{F0: 1, F1: 785, F2: ""}
                                _ = v26
                            }
                        case "a":
                            v2 = append(v2, 2147483648)
                        default:
                            v27 := GenS4{F0: 942, F1: 519, F2: (v4 + ((("a" + ("hello" + v4)) + ("a" + (v1 + v4))) + v4))}
                            _ = v27
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        fmt.Println(v4, v12)
                        _ = fi24
                        _ = v25
                    }
                }
                _ = fi20
            }
        }
        _ = v19
    } else {
        var v28 uint8 = 44
        var v29 float32 = 0.1
        for fi30 := 0; fi30 < 3; fi30++ {
            var v31 int8 = 42
            var v32 uint32 = 839
            for fi33 := 0; fi33 < 11; fi33++ {
                if true {
                    v3 = append(v3, "a")
                } else {
                    v6 = 42
                    v31 = 42
                    v29 = ((3.14159 / (v5 - v5)) / (0.1 - 3.14159))
                    switch 32767 {
                    case -945:
                        v34 := GenS1{F0: 18446744073709551615, F1: (v32 - (v32 * v32))}
                        _ = v34
                    case 42:
                        v5 -= (((v5 * (v5 * ((float32(((v11 * ((v11 + 0.0) * v11)) - v11)) * v5) - ((v5 - v5) - v29)))) + v29) - (((v29 * (1e10 / 1.5)) + v5) - ((3.14159 * v29) / v29)))
                    case 379:
                        switch (("\n\t" + "\n\t") + "foo bar") {
                        case "a":
                            v28 = (v28 % (((v28 % (((^(v28))) | 1))) | 1))
                        case "foo bar":
                            v32 += v32
                        case "hello":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        default:
                            v35 := GenS4{F0: 0, F1: 1, F2: v4}
                            _ = v35
                        }
                    default:
                        var v36 int8 = 42
                        _ = v36
                    }
                    if false {
                        v1 = ("a" + v4)
                        v32 += v32
                    }
                    v2 = append(v2, (v32 / ((2147483648) | 1)))
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = fi33
            }
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = v28
        _ = v29
    }
    if (18446744073709551615 >= 279) {
        v12 += v12
    }
    var v37 string = v4
    if (!false) {
        v37 = (("a" + (v37 + v37)) + v37)
    }
    v38 := GenS4{F0: 42, F1: 748, F2: "a"}
    for fi39 := 0; fi39 < 6; fi39++ {
        if true {
            switch "hello" {
            case "":
                fmt.Println(v5)
            case "\n\t":
                var v40 uint8 = 2
                _ = v40
            }
            var v41 int8 = 55
            _ = v41
        }
        _ = fi39
    }
    var v42 string = (v4 + "")
    for fi43 := 0; fi43 < 8; fi43++ {
        var v44 uint32 = 2147483647
        v45 := GenS5{F0: 1}
        v46 := GenS6{F0: -536, F1: true, F2: 870}
        _ = fi43
        _ = v44
        _ = v45
        _ = v46
    }
    fmt.Println(v11, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
