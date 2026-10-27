package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 uint32; F2 bool }
type GenS4 struct { F0 int32; F1 int64; F2 uint8 }
type GenS5 struct { F0 uint8; F1 uint; F2 uint32 }
func genFn0(p0 uint8, p1 int8) bool {
    p0 += 128
    switch ("hello" + "\n\t") {
    case "hello":
        if false {
            switch "\n\t" {
            case "hello":
                var v5 uint32 = 2147483648
                _ = v5
            default:
                v6 := []int{42, -9223372036854775808}
                _ = v6
            }
        } else {
            v7 := []int{-2147483648, -121}
            var v8 int = 0
            var v9 float64 = 0.1
            _ = v7
            _ = v8
            _ = v9
        }
    case "\n\t":
        if (true && true) {
            fmt.Println(p1, p1)
            switch 4294967295 {
            case 32767:
                if (42 == -171) {
                    switch "foo bar" {
                    case "\n\t":
                        if (false && false) {
                            for fi10 := 0; fi10 < 1; fi10++ {
                                var v11 string = "hello"
                                var v12 float64 = 1.5
                                _ = fi10
                                _ = v11
                                _ = v12
                            }
                            if (p0 == p0) {
                                var v13 float32 = ((0.0 - 3.14159) * 3.14159)
                                for fi14 := 0; fi14 < 14; fi14++ {
                                    v13 += v13
                                    var v15 int64 = -381
                                    _ = fi14
                                    _ = v15
                                }
                                _ = v13
                            }
                        }
                    case "":
                        p1 = p1
                    case "foo bar":
                        fmt.Println(p1, p1)
                    }
                    for fi16 := 0; fi16 < 3; fi16++ {
                        v17 := GenS1{F0: 42, F1: (fi16 / ((fi16) | 1)), F2: "\n\t"}
                        _ = fi16
                        _ = v17
                    }
                    fmt.Println(p1)
                } else {
                    var v18 int32 = 203
                    _ = v18
                }
            }
        } else {
            if false {
                var v19 uint16 = 0
                var v20 int32 = 629
                _ = v19
                _ = v20
            }
            if (p0 < p0) {
                var v21 int64 = 9223372036854775807
                _ = v21
            } else {
                v22 := GenS2{F0: ((!false) || false)}
                var v23 int8 = 100
                _ = v22
                _ = v23
            }
        }
    }
    if (1000 == 1) {
        fmt.Println(p1, p1)
        p0 -= (p0 - 128)
        if (!false) {
            v24 := []int{32768}
            p0 = p0
            _ = v24
        }
    }
    var v25 float32 = 0.1
    _ = p0
    _ = p1
    _ = v25
    return true
}
func genFn1(p0 uint64) int16 {
    var v26 uint64 = 654
    v27 := GenS3{F0: -128, F1: 1, F2: true}
    var v28 uint16 = 362
    v29 := []int{}
    _ = p0
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    return -37
}

func main() {
    v1 := []uint32{42, 2147483648, 4294967295, 909}
    var v2 int = 669
    var v3 float64 = float64((1.5 * float32(1.5)))
    v4 := make(map[string]uint32)
    if (255 == 255) {
        switch ("foo bar" + "a") {
        case "a":
            v3 += ((1.5 - 1e10) / (((v3 - (1.5 - (v3 / 0.1))) * ((((((1e10 + 1.5) - ((v3 + 0.1) / v3)) * -2.25) * 1e10) * (1e10 * ((((float64((float32(1e10) + -2.25)) / -2.25) - ((((v3 + float64(3.14159)) + (3.14159 / (1.5 * v3))) / v3) * 3.14159)) - float64((0.0 - -2.25))) / ((((0.0 - 0.1) / (-2.25 + v3)) + -2.25) + float64(0.0))))) + v3)) + v3))
        case "\n\t":
            v3 -= -2.25
        }
        v30 := GenS1{F0: 0, F1: v2, F2: "foo bar"}
        v3 = ((float64((3.14159 - 0.1)) * v3) - -2.25)
        _ = v30
    }
    var v31 uint8 = 240
    for fi32 := 0; fi32 < 19; fi32++ {
        v33 := GenS4{F0: -967, F1: 1, F2: 0}
        for fi34 := 0; fi34 < 13; fi34++ {
            var v35 int64 = -1
            if (false || (fi34 < (fi32 >> ((2) & 63)))) {
                v31 = v31
            } else {
                switch ("\n\t" + "\n\t") {
                case "\n\t":
                    v2 -= v2
                }
                if true {
                    switch ("\n\t" + "hello") {
                    case "hello":
                        var v36 uint32 = 0
                        _ = v36
                    case "a":
                        fmt.Println(fi32, v35)
                    default:
                        var v37 uint32 = 671
                        _ = v37
                    }
                    v2 = (fi32 >> ((685) & 63))
                    var v38 int32 = -1
                    _ = v38
                }
                _ = len(v1)
            }
            fmt.Println(fi34, v2)
            _ = fi34
            _ = v35
        }
        _ = fi32
        _ = v33
    }
    var v39 uint16 = 0
    for fi40 := 0; fi40 < 3; fi40++ {
        v31 -= v31
        v41 := GenS2{F0: (true || false)}
        _ = fi40
        _ = v41
    }
    for fi42 := 0; fi42 < 20; fi42++ {
        var v43 uint16 = uint16(v31)
        v44 := GenS5{F0: (v31 + (v31 * (v31 >> ((42) & 63)))), F1: 0, F2: 2}
        v2 += -2147483648
        _ = fi42
        _ = v43
        _ = v44
    }
    switch "" {
    case "":
        v45 := GenS4{F0: int32(v31), F1: -294, F2: (^(v31))}
        _ = v45
    case "hello":
        var v46 float32 = (float32(v3) / -2.25)
        _ = v46
    case "foo bar":
        if false {
            v47 := GenS5{F0: v31, F1: 18446744073709551615, F2: 0}
            fmt.Println(v2, v31)
            v3 = (v3 - v3)
            _ = v47
        } else {
            var v48 int = -885
            for fi49 := 0; fi49 < 9; fi49++ {
                _ = len(v1)
                _ = fi49
            }
            _ = v48
        }
    default:
        var v50 int16 = -32768
        _ = v50
    }
    var v51 int8 = 42
    v52 := GenS1{F0: -2147483648, F1: v2, F2: "\n\t"}
    v3 += v3
    var v53 string = "hello"
    for fi54 := 0; fi54 < 16; fi54++ {
        v4[((("" + v53) + v53) + v53)] = 2147483647
        v39 += (v39 % ((v39) | 1))
        _ = fi54
    }
    v1 = append(v1, 309)
    fmt.Println(v51, v53)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v39
    _ = v4
    _ = v51
    _ = v52
    _ = v53
}
