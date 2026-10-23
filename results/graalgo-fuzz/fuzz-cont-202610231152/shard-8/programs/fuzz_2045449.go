package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int; F1 uint8 }

func main() {
    var v1 uint32 = 0
    var v2 float32 = 1e10
    v3 := make(map[uint32]string)
    v4 := []int{-9223372036854775808, 871, -1, -2147483648}
    if len(v4) > 0 {
        v4[0] = -2147483648
    }
    var v5 int16 = 42
    v6 := GenS1{F0: v5, F1: 1}
    if true {
        switch ("foo bar" + "a") {
        case "\n\t":
            var v7 uint8 = 1
            _ = v7
        default:
            var v8 int8 = 42
            _ = v8
        }
        fmt.Println(v1)
    }
    switch 9223372036854775807 {
    case 32767:
        v9 := GenS1{F0: v5, F1: 255}
        _ = v9
    case 4294967295:
        if false {
            switch ("foo bar" + "\n\t") {
            case "a":
                fmt.Println(v1, v5, v2)
            case "hello":
                for fi10 := 0; fi10 < 1; fi10++ {
                    _ = len(v4)
                    _ = fi10
                }
            case "":
                for fi11 := 0; fi11 < 13; fi11++ {
                    if false {
                        if (!false) {
                            delete(v3, v1)
                            v5 -= (v5 ^ (v5 & v5))
                            _ = len(v3)
                        } else {
                            var v12 int32 = -514
                            v13 := v3[(^(v1))]
                            _ = v12
                            _ = v13
                        }
                    } else {
                        v1 = 2147483648
                        var v14 string = ("foo bar" + "foo bar")
                        if false {
                            switch "foo bar" {
                            case "a":
                                v5 = v5
                            case "":
                                if len(v4) > 0 {
                                    v4[0] = -911
                                }
                            case "foo bar":
                                fmt.Println(v1, v1)
                            default:
                                if true {
                                    v1 = (v1 >> ((141) & 63))
                                } else {
                                    var v15 int64 = 42
                                    _ = v15
                                }
                            }
                            var v16 bool = (false || (true && (127 != -1)))
                            _ = v16
                        } else {
                            v5 -= v5
                            for fi17 := 0; fi17 < 12; fi17++ {
                                var v18 int16 = (v5 | v5)
                                if true {
                                    v19 := GenS2{F0: (fi17 >> ((255) & 63)), F1: 2}
                                    _ = v19
                                } else {
                                    switch "hello" {
                                    case "a":
                                        v20 := GenS1{F0: (^(v5)), F1: 157}
                                        _ = v20
                                    case "hello":
                                        v3[(v1 ^ 2)] = v14
                                    default:
                                        v21 := GenS2{F0: fi11, F1: 0}
                                        _ = v21
                                    }
                                    v22 := GenS2{F0: (-(fi11)), F1: 255}
                                    v23 := GenS2{F0: 4294967295, F1: 122}
                                    _ = v22
                                    _ = v23
                                }
                                _ = fi17
                                _ = v18
                            }
                            fmt.Println(v1)
                        }
                        _ = v14
                    }
                    _ = fi11
                }
            }
            _ = len(v3)
            v3[v1] = "foo bar"
        }
    case -328:
        v1 -= v1
    default:
        switch -34 {
        case 439:
            delete(v3, (v1 / (((v1 % ((v1) | 1))) | 1)))
        case 32768:
            v2 += float32((float64((v2 / v2)) * (((0.0 * float64(float32(float64(v2)))) - -2.25) + 3.14159)))
        default:
            for fi24 := 0; fi24 < 2; fi24++ {
                if (true || true) {
                    var v25 int64 = 32
                    var v26 int = 32768
                    v3[v1] = ("foo bar" + "")
                    _ = v25
                    _ = v26
                }
                _ = fi24
            }
        }
    }
    v2 = ((float32(((1.5 + 0.0) / 1.0)) / v2) / (((v2 / v2) * (v2 * 1e10)) * 0.0))
    var v27 float64 = (float64((v2 + v2)) - 0.1)
    v27 += v27
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
