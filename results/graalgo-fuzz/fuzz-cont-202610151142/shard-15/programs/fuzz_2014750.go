package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{2147483647}
    var v3 int64 = 9223372036854775807
    var v4 uint32 = 2
    v5 := v1[("" + "\n\t")]
    var v6 float64 = (0.1 + (3.14159 + (1e10 / (0.1 / ((-2.25 - 1.5) / -2.25)))))
    if ((("foo bar" + "\n\t") >= (("hello" + "\n\t") + (("" + "\n\t") + ""))) || true) {
        fmt.Println(v4, v6, v5)
    }
    if (false && (false || false)) {
        var v7 float32 = (3.14159 - float32(v6))
        _ = v7
    } else {
        v8 := GenS1{F0: 32767, F1: 2}
        v9 := GenS1{F0: 32767, F1: 236}
        var v10 uint8 = 42
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 float32 = (1.5 + 0.1)
    switch 32768 {
    case -431:
        v4 = v5
    }
    if true {
        var v12 int64 = v3
        if true {
            switch "hello" {
            case "":
                var v13 uint32 = 42
                _ = v13
            default:
                v5 -= (v4 ^ (v4 | 26))
            }
            v14 := GenS1{F0: 801, F1: 42}
            _ = v14
        }
        _ = v12
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 uint8 = 1
        _ = fi15
        _ = v16
    }
    if (!(!false)) {
        var v17 int16 = 130
        if (2 < 42) {
            fmt.Println(v11, v17, v3)
            var v18 int = 1
            switch v18 {
            case -2147483648:
                if false {
                    v6 = 3.14159
                    for fi19 := 0; fi19 < 16; fi19++ {
                        var v20 string = ((("foo bar" + ("foo bar" + "foo bar")) + "") + "")
                        switch ((v20 + (v20 + (v20 + "a"))) + ("" + v20)) {
                        case "\n\t":
                            for fi21 := 0; fi21 < 7; fi21++ {
                                if (18446744073709551615 >= 495) {
                                    if (!(1e10 <= ((1.5 / 1e10) * v11))) {
                                        var v22 uint16 = 391
                                        _ = v22
                                    } else {
                                        v23 := GenS2{F0: (^(v3))}
                                        var v24 float32 = v11
                                        _ = v23
                                        _ = v24
                                    }
                                    v3 = v3
                                } else {
                                    var v25 float64 = v6
                                    _ = v25
                                }
                                v18 = (v18 & -327)
                                var v26 uint32 = (v5 - (v4 - (v5 << ((4294967295) & 63))))
                                _ = fi21
                                _ = v26
                            }
                        }
                        v6 -= (v6 * (0.1 + 3.14159))
                        _ = fi19
                        _ = v20
                    }
                }
            case 32768:
                var v27 int16 = v17
                _ = v27
            case 9223372036854775807:
                v18 += (v18 >> ((18446744073709551615) & 63))
            default:
                v28 := GenS1{F0: 2, F1: 937}
                _ = v28
            }
            _ = v18
        }
        if false {
            v5 = (v4 % ((471) | 1))
            v4 = (v5 % ((v5) | 1))
        }
        _ = v17
    }
    switch "\n\t" {
    case "foo bar":
        var v29 uint8 = 45
        _ = v29
    case "a":
        v5 += 2
    case "\n\t":
        var v30 uint32 = 2
        _ = v30
    default:
        v3 = 356
    }
    v4 = 2147483648
    v3 += v3
    if (669 < -848) {
        v11 -= float32(float64(v11))
        v5 -= (v5 ^ 65535)
        v6 -= float64(v11)
    } else {
        if true {
            if true {
                v6 -= v6
                fmt.Println(v4)
            }
            for fi31 := 0; fi31 < 18; fi31++ {
                v6 += (3.14159 - v6)
                fmt.Println(v6, v4)
                _ = fi31
            }
        }
        if true {
            v32 := GenS1{F0: 255, F1: 14}
            fmt.Println(v3, v11)
            fmt.Println(v3, v3)
            _ = v32
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
