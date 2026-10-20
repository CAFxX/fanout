package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int32; F1 int16; F2 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 uint64; F2 int64 }

func main() {
    var v1 int32 = -2147483648
    var v2 uint = 18446744073709551615
    var v3 float64 = ((-2.25 * (float64((((float32(3.14159) / 1.0) / (1.5 / 3.14159)) - 1e10)) * 0.0)) * (0.0 - -2.25))
    var v4 int = -2147483648
    var v5 uint8 = 1
    v6 := []int{(v4 << ((v2) & 63)), (v4 << ((221) & 63)), (v4 + (-(v4)))}
    v4 = v4
    v7 := GenS1{F0: (v1 & (^(v1))), F1: 100}
    v5 += v5
    switch (-(v4)) {
    case 385:
        v6 = append(v6, (-(v4)))
    case 32767:
        switch "" {
        case "\n\t":
            if true {
                v1 += (-(v1))
                v4 -= 2147483647
                for fi8 := 0; fi8 < 16; fi8++ {
                    var v9 float64 = v3
                    _ = fi8
                    _ = v9
                }
            } else {
                v6 = append(v6, (^(v4)))
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                if false {
                    var v10 int8 = -87
                    v6 = append(v6, (v4 | (-(v4))))
                    v11 := GenS1{F0: (v1 % (((v1 | (v1 >> ((65535) & 63)))) | 1)), F1: -87}
                    _ = v10
                    _ = v11
                } else {
                    switch int(v5) {
                    case -9223372036854775808:
                        v12 := GenS2{F0: 2147483647, F1: 0, F2: (false && ((v5 & (^(v5))) >= 42))}
                        _ = v12
                    }
                    switch "hello" {
                    case "\n\t":
                        if len(v6) > 0 {
                            v6[0] = 42
                        }
                    case "a":
                        v13 := GenS3{F0: 0}
                        _ = v13
                    case "hello":
                        if (9223372036854775807 >= -910) {
                            switch "hello" {
                            case "a":
                                v1 += 0
                            case "foo bar":
                                fmt.Println(v2)
                            case "":
                                var v14 bool = true
                                _ = v14
                            }
                        } else {
                            var v15 uint8 = (v5 + 127)
                            switch 535 {
                            case 498:
                                switch (v4 % ((146) | 1)) {
                                case 2147483647:
                                    if len(v6) > 0 {
                                        v6[0] = (v4 & v4)
                                    }
                                case 0:
                                    if ((^(v2)) == v2) {
                                        v3 += 3.14159
                                        v16 := GenS2{F0: v1, F1: 867, F2: (1 > 927)}
                                        for fi17 := 0; fi17 < 11; fi17++ {
                                            _ = len(v6)
                                            _ = fi17
                                        }
                                        _ = v16
                                    }
                                case 895:
                                    v6 = append(v6, (-(v4)))
                                default:
                                    v15 = 174
                                }
                            case 0:
                                var v18 uint64 = 42
                                _ = v18
                            }
                            if (v3 == (0.0 * (v3 - v3))) {
                                for fi19 := 0; fi19 < 19; fi19++ {
                                    if false {
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        if false {
                                            for fi20 := 0; fi20 < 1; fi20++ {
                                                var v21 bool = false
                                                fmt.Println(v1, v1)
                                                var v22 bool = v21
                                                _ = fi20
                                                _ = v21
                                                _ = v22
                                            }
                                            var v23 uint8 = v5
                                            v2 = (v2 | 422)
                                            _ = v23
                                        } else {
                                            var v24 bool = (false || (false || (("foo bar" + "a") == "")))
                                            _ = v24
                                        }
                                    }
                                    _ = fi19
                                }
                            }
                            _ = v15
                        }
                    default:
                        for fi25 := 0; fi25 < 14; fi25++ {
                            v2 = (v2 & (v2 ^ (v2 - 4294967295)))
                            _ = fi25
                        }
                    }
                    for fi26 := 0; fi26 < 1; fi26++ {
                        fmt.Println(v3, v1, v1)
                        v27 := GenS2{F0: v1, F1: 0, F2: (false && (false || (!(0 > 1))))}
                        _ = fi26
                        _ = v27
                    }
                }
            }
        default:
            fmt.Println(v5)
        }
    case 42:
        if true {
            var v28 uint = v2
            _ = v28
        }
    }
    fmt.Println(v4)
    if (1 == -9223372036854775808) {
        var v29 uint = v2
        v30 := GenS1{F0: v1, F1: -79}
        _ = v29
        _ = v30
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v31 := GenS3{F0: uint32(v5)}
    switch "a" {
    case "hello":
        v32 := GenS4{F0: uint32(v2), F1: 42, F2: 976}
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
