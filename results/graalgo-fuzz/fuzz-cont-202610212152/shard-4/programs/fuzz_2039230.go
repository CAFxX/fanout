package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 string; F2 int32 }

func main() {
    var v1 uint32 = 2147483647
    v2 := []uint32{65535}
    v3 := make(map[uint32]uint32)
    var v4 uint16 = 9
    var v5 int32 = -2147483648
    v6 := []uint32{731, (v1 % ((v1) | 1)), 42}
    delete(v3, v1)
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: 128}
        v4 -= v4
        _ = fi7
        _ = v8
    }
    var v9 int64 = -9223372036854775808
    v1 = v1
    switch "a" {
    case "hello":
        v10 := GenS2{F0: (v4 + v4), F1: "a", F2: -937}
        _ = v10
    case "foo bar":
        v1 = 65535
    case "\n\t":
        if true {
            v11 := GenS2{F0: (v4 | (^(v4))), F1: "", F2: 0}
            v5 += 0
            _ = v11
        }
    }
    v5 = v5
    v1 -= uint32(v4)
    v5 -= -253
    fmt.Println(v5)
    var v12 uint64 = 0
    v13 := GenS1{F0: 127}
    switch "hello" {
    case "\n\t":
        if false {
            fmt.Println(v12, v12)
            if false {
                var v14 float64 = (1.5 - 0.1)
                v4 -= v4
                delete(v3, (^(v1)))
                _ = v14
            } else {
                v4 = v4
            }
            v4 = uint16(v12)
        }
    case "foo bar":
        if false {
            if (("" + "") == "\n\t") {
                switch "a" {
                case "foo bar":
                    if false {
                        var v15 bool = true
                        switch int(v9) {
                        case 42:
                            var v16 int16 = 0
                            _ = v16
                        case 0:
                            v9 = (-(v9))
                        case -2147483648:
                            for fi17 := 0; fi17 < 11; fi17++ {
                                var v18 bool = (v4 >= v4)
                                var v19 int = fi17
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                        }
                        _ = v15
                    } else {
                        v2 = append(v2, v1)
                    }
                case "a":
                    if ((1e10 * 3.14159) <= 0.1) {
                        var v20 int8 = 0
                        _ = v20
                    }
                case "\n\t":
                    var v21 float32 = 0.1
                    _ = v21
                default:
                    delete(v3, (^(v1)))
                }
                fmt.Println(v12, v4)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            var v22 int16 = 0
            _ = v22
        }
    case "a":
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 uint64 = v12
            v25 := GenS2{F0: v4, F1: ("foo bar" + "a"), F2: v5}
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    fmt.Println(v1, v5)
    fmt.Println(v12, v5)
    fmt.Println(v5, v1)
    var v26 float64 = ((1e10 - ((((1.5 - ((-2.25 * ((0.1 - 3.14159) - 0.0)) + -2.25)) + 1e10) / 1.5) - (0.1 - (1.5 / 0.1)))) / ((((1.5 - float64(0.0)) / 1.5) + ((0.0 * (((1.5 + 3.14159) * 0.1) * (3.14159 * ((-2.25 - (3.14159 * 0.0)) + 1e10)))) - (0.1 + 0.1))) + (1.5 / 1.0)))
    var v27 uint16 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
