package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 int; F2 int8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = -32768
    var v3 uint16 = 2
    var v4 float32 = 0.0
    var v5 uint8 = 2
    switch "a" {
    case "\n\t":
        var v6 float64 = 3.14159
        _ = v6
    default:
        fmt.Println(v2, v5)
    }
    v5 += (v5 / ((56) | 1))
    for fi7 := 0; fi7 < 4; fi7++ {
        switch "" {
        case "foo bar":
            fmt.Println(fi7, v5, v4)
        case "a":
            v8 := GenS1{F0: 2147483648, F1: 32767}
            _ = v8
        }
        if ((v3 + (v3 * v3)) >= v3) {
            v5 += (v5 >> ((42) & 63))
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 uint32 = 0
                var v11 int16 = (v2 * 1)
                v4 += (((0.1 - (v4 - v4)) - -2.25) + v4)
                v12 := GenS2{F0: 0}
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            switch 4294967295 {
            case 2147483647:
                if (!true) {
                    switch "hello" {
                    case "":
                        v1[2] = ("foo bar" + "a")
                    case "foo bar":
                        var v13 uint16 = 1
                        _ = v13
                    default:
                        switch "a" {
                        case "a":
                            v14 := GenS1{F0: 1, F1: (v2 + v2)}
                            _ = v14
                        case "hello":
                            var v15 uint = 42
                            _ = v15
                        case "\n\t":
                            switch (fi7 >> ((95) & 63)) {
                            case -2147483648:
                                v5 = (v5 / ((v5) | 1))
                            case 127:
                                switch (fi7 & (fi7 | (fi7 << ((4294967295) & 63)))) {
                                case 9223372036854775807:
                                    v4 -= float32(((-2.25 + 1e10) / 3.14159))
                                default:
                                    v16 := GenS1{F0: uint32(v3), F1: v2}
                                    _ = v16
                                }
                            default:
                                if false {
                                    switch fi7 {
                                    case 4294967295:
                                        var v17 uint16 = (v3 * (v3 | v3))
                                        _ = v17
                                    }
                                    if (!(!true)) {
                                        var v18 string = ("a" + (((("foo bar" + ("hello" + "")) + "a") + "a") + "hello"))
                                        if true {
                                            var v19 float32 = (float32(0.1) - v4)
                                            _ = v19
                                        }
                                        _ = v18
                                    } else {
                                        var v20 bool = (0 == -9223372036854775808)
                                        for fi21 := 0; fi21 < 10; fi21++ {
                                            var v22 uint64 = 352
                                            delete(v1, 983)
                                            _ = fi21
                                            _ = v22
                                        }
                                        _ = v20
                                    }
                                    for fi23 := 0; fi23 < 3; fi23++ {
                                        v4 = (v4 / (((v4 + 0.0) + (v4 * (((v4 / (v4 + 0.0)) * 1e10) - v4))) * 0.0))
                                        v24 := []int{fi7, -717}
                                        _ = fi23
                                        _ = v24
                                    }
                                }
                            }
                        }
                    }
                    var v25 int8 = -128
                    v26 := GenS2{F0: 1}
                    _ = v25
                    _ = v26
                }
            case 657:
                var v27 int16 = -32768
                _ = v27
            default:
                delete(v1, 2147483647)
            }
        } else {
            if false {
                for fi28 := 0; fi28 < 2; fi28++ {
                    if false {
                        v2 -= (v2 ^ v2)
                        fmt.Println(fi7, v3, v2)
                        var v29 int = (^(fi28))
                        v4 = v4
                        _ = v29
                    } else {
                        if ((((float32(1.5) / float32(0.1)) / (v4 - v4)) * 3.14159) <= (-2.25 * (v4 - (3.14159 / (-2.25 + (1e10 * -2.25)))))) {
                            v5 = (v5 - (v5 / ((v5) | 1)))
                        } else {
                            fmt.Println(v4, fi28, v3)
                        }
                    }
                    v2 -= (^(v2))
                    _ = fi28
                }
                v30 := GenS1{F0: 637, F1: 32767}
                _ = v30
            }
        }
        _ = fi7
    }
    var v31 float64 = float64((v4 - ((1.5 / (v4 + ((v4 - 1e10) * ((v4 - v4) * 3.14159)))) - v4)))
    _ = len(v1)
    v32 := GenS3{F0: 91, F1: 4294967295, F2: 127}
    var v33 int8 = -128
    v3 = (^(v3))
    var v34 uint8 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
