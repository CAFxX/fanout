package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 int; F1 uint; F2 uint64 }
func genFn0(p0 uint32) int32 {
    var v5 uint8 = 2
    fmt.Println(v5, p0)
    v5 -= (v5 - (v5 % ((v5) | 1)))
    _ = p0
    _ = v5
    return -2147483648
}

func main() {
    v1 := []string{}
    v2 := make(map[int]int)
    var v3 uint64 = 1
    var v4 int16 = -1
    v2[2147483647] = 411
    v6 := GenS1{F0: (!true), F1: 42, F2: 273}
    v3 = 18446744073709551615
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            v2[(fi7 ^ 32768)] = (fi8 - fi7)
            if len(v1) > 0 {
                v1[0] = "a"
            }
            _ = fi8
        }
        for fi9 := 0; fi9 < 19; fi9++ {
            if ((93 > 127) && (!true)) {
                v3 = v3
                var v10 uint = 740
                _ = len(v1)
                _ = v10
            }
            _ = fi9
        }
        _ = fi7
    }
    v4 = v4
    var v11 uint = 474
    for fi12 := 0; fi12 < 20; fi12++ {
        v13 := GenS2{F0: 127, F1: 627}
        switch fi12 {
        case 2147483647:
            if false {
                switch "" {
                case "":
                    var v14 uint32 = 42
                    _ = v14
                }
                if (42 != 9223372036854775807) {
                    v15 := GenS1{F0: true, F1: 110, F2: (v11 | v11)}
                    var v16 int16 = 32767
                    _ = v15
                    _ = v16
                }
            } else {
                switch "foo bar" {
                case "":
                    for fi17 := 0; fi17 < 17; fi17++ {
                        fmt.Println(v3)
                        fmt.Println(fi12)
                        v18 := GenS3{F0: 32767, F1: (v11 + 18446744073709551615), F2: v3}
                        _ = fi17
                        _ = v18
                    }
                case "a":
                    switch int(v3) {
                    case 127:
                        _ = len(v2)
                    }
                default:
                    if (!false) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        for fi19 := 0; fi19 < 16; fi19++ {
                            v4 = 42
                            if ((true || true) || true) {
                                switch "" {
                                case "foo bar":
                                    switch ("\n\t" + "hello") {
                                    case "":
                                        v3 -= (v3 >> (((v11 / (((^(v11))) | 1))) & 63))
                                    default:
                                        var v20 bool = false
                                        _ = v20
                                    }
                                }
                                var v21 uint = uint(v3)
                                _ = v21
                            } else {
                                var v22 uint64 = 1
                                v23 := v2[fi19]
                                _ = v22
                                _ = v23
                            }
                            v24 := GenS2{F0: 255, F1: -9223372036854775808}
                            _ = fi19
                            _ = v24
                        }
                        if ((180 >= 2147483648) && (!(true && (1.5 < 1.5)))) {
                            var v25 float32 = (0.1 * 1e10)
                            _ = v25
                        }
                        v2[127] = fi12
                    }
                }
            }
        case -1:
            var v26 uint16 = 65535
            _ = v26
        case 0:
            v27 := GenS3{F0: fi12, F1: (^(v11)), F2: 18446744073709551615}
            _ = v27
        default:
            var v28 int16 = (v4 * v4)
            _ = v28
        }
        v3 -= (v3 + (v3 / ((v3) | 1)))
        _ = fi12
        _ = v13
    }
    switch -319 {
    case 4294967295:
        if true {
            var v29 int8 = -128
            _ = v29
        }
    case -1:
        switch "hello" {
        case "\n\t":
            fmt.Println(v3, v3, v11)
        case "hello":
            switch "\n\t" {
            case "\n\t":
                v4 -= (v4 ^ (^(v4)))
            case "a":
                var v30 int16 = (v4 % ((v4) | 1))
                _ = v30
            }
        case "a":
            if false {
                switch 1 {
                case 1:
                    if true {
                        switch "" {
                        case "hello":
                            if false {
                                var v31 int8 = 46
                                _ = v31
                            }
                        case "a":
                            _ = len(v1)
                        case "\n\t":
                            for fi32 := 0; fi32 < 6; fi32++ {
                                var v33 uint32 = 1
                                var v34 int32 = -1
                                _ = fi32
                                _ = v33
                                _ = v34
                            }
                        default:
                            v35 := v2[2147483647]
                            _ = v35
                        }
                        if len(v1) > 0 {
                            v1[0] = ("" + "foo bar")
                        }
                        if true {
                            v2[128] = int(v11)
                            var v36 float64 = (-2.25 + 0.1)
                            fmt.Println(v3, v36)
                            _ = v36
                        }
                    }
                default:
                    fmt.Println(v11)
                }
            }
        default:
            _ = len(v2)
        }
    case 127:
        v4 -= (v4 / (((v4 & (v4 ^ v4))) | 1))
    }
    v3 = v3
    v3 += v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
