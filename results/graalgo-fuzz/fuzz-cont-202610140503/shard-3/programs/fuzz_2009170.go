package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }

func main() {
    v1 := []string{("\n\t" + "a"), "foo bar"}
    var v2 int64 = 341
    var v3 float64 = ((0.1 + 3.14159) * (((0.0 - ((3.14159 * (1e10 / (1.5 + (0.1 * (0.1 - 0.1))))) - 3.14159)) - 0.0) * 0.1))
    var v4 int16 = -32768
    var v5 float32 = 1e10
    v1 = append(v1, "")
    var v6 int = -873
    fmt.Println(v2)
    v6 -= 128
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint64 = 42
        var v9 bool = false
        _ = fi7
        _ = v8
        _ = v9
    }
    if true {
        if (0 > 100) {
            var v10 uint = 631
            if (2147483648 > 0) {
                switch int(v2) {
                case 32767:
                    for fi11 := 0; fi11 < 6; fi11++ {
                        var v12 int64 = v2
                        v2 = 1
                        _ = len(v1)
                        _ = fi11
                        _ = v12
                    }
                case -1:
                    _ = len(v1)
                case 2147483647:
                    for fi13 := 0; fi13 < 6; fi13++ {
                        switch "" {
                        case "a":
                            v4 += 1
                        case "\n\t":
                            switch 32767 {
                            case 9223372036854775807:
                                if false {
                                    var v14 float32 = float32((v3 / (v3 * 1e10)))
                                    _ = v14
                                } else {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            }
                        case "hello":
                            for fi15 := 0; fi15 < 18; fi15++ {
                                fmt.Println(v4, v6)
                                _ = fi15
                            }
                        default:
                            for fi16 := 0; fi16 < 19; fi16++ {
                                fmt.Println(fi16)
                                _ = len(v1)
                                v1 = append(v1, ("\n\t" + "a"))
                                _ = fi16
                            }
                        }
                        v4 = (v4 >> (((v10 & v10)) & 63))
                        _ = fi13
                    }
                default:
                    var v17 string = "\n\t"
                    _ = v17
                }
                v4 = v4
            }
            _ = v10
        } else {
            var v18 uint32 = 62
            _ = v18
        }
    }
    var v19 int32 = 2147483647
    for fi20 := 0; fi20 < 8; fi20++ {
        for fi21 := 0; fi21 < 16; fi21++ {
            var v22 int16 = 41
            var v23 bool = true
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi20
    }
    if ((1.5 * (v3 - (((v3 / 1.0) * 0.0) / (v3 + 1.5)))) != (v3 * 3.14159)) {
        v3 += ((((v3 + v3) / ((float64(((((v5 + 3.14159) - ((float32(v3) - float32(3.14159)) + v5)) * 0.1) - -2.25)) - ((v3 * v3) + float64((((v5 / float32(v3)) / 3.14159) / v5)))) + (1e10 + v3))) - v3) + 1.5)
        v4 = v4
        fmt.Println(v3, v5, v3)
    } else {
        v2 += (v2 >> ((18446744073709551615) & 63))
        var v24 float32 = 0.0
        _ = v24
    }
    switch "a" {
    case "hello":
        for fi25 := 0; fi25 < 10; fi25++ {
            _ = len(v1)
            _ = fi25
        }
    case "foo bar":
        if false {
            for fi26 := 0; fi26 < 7; fi26++ {
                if ((false && (-70 != -1)) || (439 >= 385)) {
                    var v27 int16 = (v4 ^ 944)
                    for fi28 := 0; fi28 < 8; fi28++ {
                        v27 = v4
                        switch (v6 / ((fi26) | 1)) {
                        case -9223372036854775808:
                            v29 := GenS1{F0: 65535, F1: (v19 << ((2) & 63))}
                            _ = v29
                        case 42:
                            if (!true) {
                                v5 = ((((v5 - (float32(v3) / 1e10)) * float32(v3)) + (float32(v3) / ((float32((float64(v5) - ((v3 * v3) + ((v3 * 1e10) / 3.14159)))) / (-2.25 / 1.0)) * 3.14159))) / 1.0)
                            }
                        case 2147483647:
                            var v30 int = fi26
                            _ = v30
                        }
                        var v31 int = (fi26 | (fi26 / ((fi28) | 1)))
                        _ = fi28
                        _ = v31
                    }
                    _ = v27
                }
                fmt.Println(v4, fi26)
                var v32 uint8 = 2
                _ = fi26
                _ = v32
            }
            if false {
                var v33 uint = 997
                var v34 int16 = int16(v33)
                _ = v33
                _ = v34
            } else {
                var v35 int8 = 127
                _ = v35
            }
        }
    default:
        for fi36 := 0; fi36 < 7; fi36++ {
            if len(v1) > 0 {
                v1[0] = "a"
            }
            _ = fi36
        }
    }
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    v2 += v2
    var v37 int64 = 15
    v1 = append(v1, ("foo bar" + "a"))
    v1 = append(v1, ("" + "a"))
    v5 -= v5
    v1 = append(v1, (("\n\t" + "a") + ("hello" + "hello")))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
