package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint }
type GenS2 struct { F0 string; F1 uint32; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 string = (("\n\t" + (("a" + (((("foo bar" + (((("a" + "") + "\n\t") + "a") + ("a" + ("\n\t" + "hello")))) + "foo bar") + "\n\t") + "foo bar")) + ("foo bar" + ("hello" + "\n\t")))) + "")
    var v2 int8 = -128
    var v3 uint32 = 292
    var v4 uint64 = 2
    v5 := GenS1{F0: v4, F1: 42, F2: 42}
    var v6 int32 = 2147483647
    v7 := []int{}
    v6 += 729
    var v8 int32 = (v6 | 1)
    v9 := GenS1{F0: (v4 << ((1) & 63)), F1: -2147483648, F2: 947}
    v3 += (v3 >> ((913) & 63))
    for fi10 := 0; fi10 < 3; fi10++ {
        v4 -= v4
        _ = fi10
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            var v13 uint16 = 986
            v14 := GenS2{F0: (v1 + (v1 + "foo bar")), F1: v3, F2: 0}
            if true {
                switch fi12 {
                case 1:
                    v15 := GenS1{F0: 0, F1: fi12, F2: 678}
                    _ = v15
                case 21:
                    switch v1 {
                    case "\n\t":
                        _ = len(v7)
                    default:
                        var v16 uint16 = 65535
                        _ = v16
                    }
                    switch "" {
                    case "hello":
                        if len(v7) > 0 {
                            v7[0] = fi12
                        }
                        switch fi11 {
                        case 0:
                            var v17 float64 = 0.1
                            _ = v17
                        case 982:
                            var v18 int64 = 570
                            var v19 uint8 = 107
                            fmt.Println(v4)
                            _ = v18
                            _ = v19
                        default:
                            fmt.Println(v3)
                        }
                        if (((v2 % ((v2) | 1)) < (v2 ^ v2)) || true) {
                            var v20 float64 = ((0.1 / 1.5) + 3.14159)
                            v3 += v3
                            _ = v20
                        }
                    case "foo bar":
                        var v21 uint16 = 620
                        _ = v21
                    case "\n\t":
                        v3 = 2147483647
                    }
                    v22 := GenS2{F0: ((v1 + (v1 + v1)) + ("" + ((v1 + v1) + v1))), F1: v3, F2: 127}
                    _ = v22
                case -2147483648:
                    v7 = append(v7, (fi12 % ((127) | 1)))
                }
            } else {
                var v23 uint32 = (v3 | (v3 ^ (v3 ^ 42)))
                _ = v23
            }
            var v24 uint64 = 0
            if (!false) {
                var v25 bool = false
                if (-682 == 32767) {
                    v6 = v6
                    fmt.Println(v8, v1)
                } else {
                    v13 = 255
                }
                _ = v25
            } else {
                v4 -= (v4 | 18446744073709551615)
                switch v1 {
                case "hello":
                    v26 := GenS3{F0: 255, F1: v4}
                    _ = v26
                case "\n\t":
                    v27 := GenS2{F0: ("a" + ""), F1: (v3 << ((18446744073709551615) & 63)), F2: 255}
                    _ = v27
                case "":
                    switch v1 {
                    case "foo bar":
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    case "\n\t":
                        if false {
                            var v28 float64 = (0.0 + 1.5)
                            v29 := GenS4{F0: 128, F1: v3}
                            _ = v28
                            _ = v29
                        } else {
                            switch 2147483647 {
                            case -9223372036854775808:
                                v1 = (("" + v1) + "\n\t")
                            case -1:
                                v3 -= v3
                            default:
                                v8 += v6
                            }
                            v24 = (v4 | v4)
                            var v30 uint = 255
                            _ = v30
                        }
                    case "":
                        fmt.Println(v8)
                    default:
                        v31 := GenS4{F0: 128, F1: (^(v3))}
                        fmt.Println(v8)
                        switch v1 {
                        case "foo bar":
                            fmt.Println(v3, v1, fi11)
                        case "":
                            v32 := GenS2{F0: "hello", F1: (^(v3)), F2: 128}
                            _ = v32
                        }
                        var v33 uint64 = uint64(v3)
                        _ = v31
                        _ = v33
                    }
                }
            }
            _ = fi12
            _ = v13
            _ = v14
            _ = v24
        }
        v6 = (v6 << ((122) & 63))
        _ = fi11
    }
    v7 = append(v7, -9223372036854775808)
    var v34 uint8 = 143
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
