package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 int16; F2 int64 }
type GenS4 struct { F0 string; F1 int8 }
func genFn0() bool {
    var v7 float64 = ((1.5 + (-2.25 / ((-2.25 * (float64(1.5) * 0.1)) * ((0.1 / -2.25) - (0.0 + (0.0 / -2.25)))))) * 3.14159)
    switch 42 {
    case 2147483647:
        if false {
            v7 = (-2.25 / -2.25)
        } else {
            v8 := []int{}
            _ = v8
        }
    case 32767:
        v9 := []int{128, -721, 127}
        _ = v9
    }
    v7 = (v7 * (1e10 + float64(((0.0 / ((1e10 + 1e10) / 3.14159)) + 1e10))))
    var v10 int16 = 0
    _ = v10
    _ = v7
    return false
}

func main() {
    var v1 int8 = 52
    var v2 uint64 = 42
    var v3 uint = 2
    v4 := make(map[int]int)
    var v5 uint8 = 52
    v6 := []int{127}
    if (624 > 2147483647) {
        var v11 uint = (^(v3))
        _ = v11
    } else {
        v12 := GenS1{F0: 346, F1: 2147483647}
        var v13 string = ("foo bar" + ("hello" + "\n\t"))
        _ = v12
        _ = v13
    }
    v14 := GenS2{F0: 18446744073709551615}
    v4[-2147483648] = 32768
    var v15 int = -9223372036854775808
    v16 := GenS2{F0: (v2 | v2)}
    if (false || false) {
        if true {
            for fi17 := 0; fi17 < 9; fi17++ {
                if len(v6) > 0 {
                    v6[0] = v15
                }
                var v18 float64 = 3.14159
                if len(v6) > 0 {
                    v6[0] = 4294967295
                }
                _ = fi17
                _ = v18
            }
            for fi19 := 0; fi19 < 9; fi19++ {
                switch "foo bar" {
                case "\n\t":
                    for fi20 := 0; fi20 < 3; fi20++ {
                        v5 = v5
                        delete(v4, -1)
                        v1 = (v1 / ((1) | 1))
                        switch (fi20 | (v15 << ((255) & 63))) {
                        case -1:
                            switch 2147483647 {
                            case 0:
                                var v21 uint64 = 0
                                _ = v21
                            default:
                                if false {
                                    var v22 int64 = -1
                                    fmt.Println(v1)
                                    var v23 int8 = (v1 >> (((v3 + uint(v5))) & 63))
                                    _ = v22
                                    _ = v23
                                }
                                var v24 uint64 = v2
                                var v25 float64 = float64(3.14159)
                                delete(v4, (fi19 / ((v15) | 1)))
                                _ = v24
                                _ = v25
                            }
                        case -260:
                            var v26 uint = 0
                            _ = v26
                        case -709:
                            v27 := GenS3{F0: 387, F1: 1, F2: -472}
                            _ = v27
                        default:
                            var v28 float32 = 0.1
                            _ = v28
                        }
                        _ = fi20
                    }
                case "a":
                    var v29 uint32 = 1
                    _ = v29
                case "hello":
                    var v30 int32 = -1
                    _ = v30
                default:
                    switch ("a" + "a") {
                    case "a":
                        v3 -= 42
                    case "hello":
                        v4[672] = fi19
                    case "":
                        switch fi19 {
                        case -282:
                            if ((!("foo bar" >= "hello")) && (false && false)) {
                                v6 = append(v6, (^(fi19)))
                                v31 := GenS1{F0: uint16(v2), F1: 4294967295}
                                var v32 uint = v3
                                _ = v31
                                _ = v32
                            }
                        case -324:
                            _ = len(v6)
                        default:
                            fmt.Println(v5)
                        }
                    }
                }
                for fi33 := 0; fi33 < 6; fi33++ {
                    if ((!false) || (false || true)) {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    } else {
                        if false {
                            var v34 int = fi19
                            _ = v34
                        } else {
                            v1 -= -50
                            var v35 int = (fi19 >> (((^(v3))) & 63))
                            _ = v35
                        }
                        fmt.Println(v3, v3)
                        if len(v6) > 0 {
                            v6[0] = 4294967295
                        }
                    }
                    var v36 uint = 0
                    _ = fi33
                    _ = v36
                }
                var v37 uint32 = 2147483647
                _ = fi19
                _ = v37
            }
        }
        v2 = (v2 >> (((v3 / ((v3) | 1))) & 63))
    }
    fmt.Println(v1)
    v3 -= (v3 - (v3 / ((v3) | 1)))
    var v38 uint32 = 42
    v39 := GenS3{F0: int16(v1), F1: -32768, F2: -990}
    switch "a" {
    case "a":
        v4[(^(v15))] = (^(v15))
    default:
        var v40 int64 = int64(v15)
        _ = v40
    }
    v2 -= (v2 & uint64(v38))
    v41 := GenS4{F0: "a", F1: (v1 & (v1 | (v1 << ((1) & 63))))}
    v15 += 1
    var v42 uint16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
