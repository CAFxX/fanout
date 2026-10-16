package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS6 struct { F0 int32; F1 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]bool)
    var v3 uint8 = 0
    var v4 int = 4294967295
    var v5 int = v4
    v6 := []int{-2147483648, (v5 - v5)}
    if (!(!false)) {
        var v7 bool = false
        v3 += v3
        for fi8 := 0; fi8 < 15; fi8++ {
            for fi9 := 0; fi9 < 4; fi9++ {
                v10 := GenS1{F0: fi8, F1: 67}
                _ = fi9
                _ = v10
            }
            if v7 {
                if v7 {
                    v11 := GenS1{F0: (fi8 >> ((65535) & 63)), F1: -122}
                    if v7 {
                        v12 := v2[1]
                        var v13 float64 = 1.5
                        v14 := GenS2{F0: 0, F1: 42}
                        _ = v12
                        _ = v13
                        _ = v14
                    } else {
                        var v15 uint8 = (v3 - (^(v3)))
                        var v16 uint64 = 2
                        v17 := GenS3{F0: 199, F1: ((-74 != 0) && true), F2: 2147483647}
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    _ = v11
                } else {
                    if false {
                        v3 -= v3
                        _ = len(v6)
                    }
                    v18 := GenS2{F0: 378, F1: -828}
                    if (0 > -646) {
                        _ = len(v2)
                        for fi19 := 0; fi19 < 4; fi19++ {
                            fmt.Println(v3)
                            _ = fi19
                        }
                    } else {
                        v2[2147483647] = true
                        v20 := GenS2{F0: 1, F1: 2}
                        if v7 {
                            v7 = ((1.5 + 1.5) < 3.14159)
                            var v21 string = ("foo bar" + ("hello" + ""))
                            var v22 uint8 = 2
                            _ = v21
                            _ = v22
                        }
                        _ = v20
                    }
                    _ = v18
                }
            }
            if (false && v7) {
                if v7 {
                    for fi23 := 0; fi23 < 18; fi23++ {
                        var v24 uint8 = 42
                        if false {
                            v6 = append(v6, 127)
                        }
                        _ = fi23
                        _ = v24
                    }
                    v25 := GenS2{F0: 0, F1: -676}
                    var v26 uint = 603
                    _ = v25
                    _ = v26
                }
                delete(v2, 552)
                v1[42] = 2147483647
            }
            _ = fi8
        }
        _ = v7
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi27 := 0; fi27 < 6; fi27++ {
        v5 -= (-(v5))
        _ = fi27
    }
    v4 = (v5 + 127)
    fmt.Println(v5)
    v6 = append(v6, (v4 << ((168) & 63)))
    var v28 bool = false
    v29 := GenS4{F0: (^(v5))}
    v28 = (((3.14159 / 3.14159) * 3.14159) >= float64((1.5 / (1e10 * 1e10))))
    if len(v6) > 0 {
        v6[0] = (v5 - v4)
    }
    if (0 < 65535) {
        var v30 int32 = 2147483647
        var v31 string = "\n\t"
        var v32 int32 = (v30 ^ int32(v3))
        _ = v30
        _ = v31
        _ = v32
    } else {
        switch ("" + "") {
        case "":
            v33 := GenS5{F0: ((((-2.25 - 3.14159) * 3.14159) - ((1e10 / (1e10 / 0.1)) / ((1e10 / -2.25) + 0.0))) == 3.14159), F1: v3, F2: 128}
            _ = v33
        case "a":
            var v34 int32 = 2147483647
            _ = v34
        case "foo bar":
            if len(v6) > 0 {
                v6[0] = v4
            }
        default:
            if (3.14159 >= 1e10) {
                v28 = false
                v35 := GenS6{F0: 1, F1: 2}
                _ = v35
            }
        }
        v3 -= v3
    }
    switch ("hello" + "a") {
    case "hello":
        for fi36 := 0; fi36 < 6; fi36++ {
            delete(v2, 1)
            fmt.Println(v4, v5)
            _ = fi36
        }
    default:
        if (!(42 > -2147483648)) {
            v4 = 621
            var v37 float64 = float64(float32((((((0.1 - 1.5) + (3.14159 / (3.14159 / (3.14159 - (1.5 / 1.0))))) + -2.25) / 1e10) - 0.0)))
            var v38 float32 = (float32((1.5 / (((-2.25 / v37) + v37) - (0.1 + 3.14159)))) * float32(float64(((3.14159 / (-2.25 + 3.14159)) / -2.25))))
            _ = v37
            _ = v38
        } else {
            var v39 int = (v4 ^ v4)
            for fi40 := 0; fi40 < 3; fi40++ {
                v3 -= v3
                _ = fi40
            }
            v3 += v3
            _ = v39
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
