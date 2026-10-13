package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int8 }

func main() {
    var v1 int8 = -128
    v2 := make(map[uint32]bool)
    v3 := make(map[int]int)
    var v4 uint8 = 1
    v4 += v4
    v5 := GenS1{F0: v4, F1: 18446744073709551615, F2: (v1 | v1)}
    var v6 bool = (true && (!true))
    if v6 {
        v6 = ((!(-142 > 0)) && v6)
        var v7 float32 = (0.1 * -2.25)
        fmt.Println(v1, v6, v7)
        _ = v7
    } else {
        v8 := GenS1{F0: (v4 / (((v4 ^ v4)) | 1)), F1: 1, F2: -1}
        if ((v1 | (v1 / (((v1 + v1)) | 1))) == (v1 % (((v1 << ((286) & 63))) | 1))) {
            if v6 {
                for fi9 := 0; fi9 < 4; fi9++ {
                    fmt.Println(v1)
                    _ = fi9
                }
                if true {
                    for fi10 := 0; fi10 < 7; fi10++ {
                        v11 := GenS1{F0: (v4 % (((v4 & (v4 * v4))) | 1)), F1: 2, F2: (v1 - (v1 - v1))}
                        v6 = false
                        v6 = ((-196 < 1) || v6)
                        _ = fi10
                        _ = v11
                    }
                    delete(v2, 781)
                }
            } else {
                if v6 {
                    v4 -= (v4 << ((42) & 63))
                } else {
                    v6 = (v6 || v6)
                    v6 = v6
                }
                if v6 {
                    var v12 uint = 4294967295
                    _ = v12
                }
            }
            v1 += v1
            v13 := GenS1{F0: (^(v4)), F1: 2, F2: (v1 + -53)}
            _ = v13
        }
        var v14 bool = v6
        _ = v14
        _ = v8
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        fmt.Println(v4, v6, v6)
        if true {
            for fi16 := 0; fi16 < 14; fi16++ {
                fmt.Println(v6, v1)
                v17 := v3[32767]
                v18 := GenS1{F0: 1, F1: 335, F2: v1}
                _ = fi16
                _ = v17
                _ = v18
            }
        } else {
            v2[486] = v6
        }
        _ = fi15
    }
    switch ("a" + (("hello" + ("\n\t" + ("foo bar" + "hello"))) + "hello")) {
    case "a":
        var v19 int32 = 1
        _ = v19
    default:
        fmt.Println(v4)
    }
    if (v6 && ("foo bar" <= "a")) {
        var v20 int16 = 883
        _ = v20
    }
    var v21 uint8 = 255
    if ((v21 << ((65535) & 63)) < v21) {
        v1 = v1
        fmt.Println(v21, v6, v1)
        fmt.Println(v1)
    } else {
        if (9223372036854775807 >= -167) {
            var v22 int16 = -1
            v23 := GenS1{F0: (v21 ^ (^(v21))), F1: 16, F2: -1}
            if v6 {
                delete(v2, 667)
                var v24 uint8 = 255
                var v25 uint64 = 718
                _ = v24
                _ = v25
            }
            _ = v22
            _ = v23
        }
    }
    v21 = (^(v4))
    for fi26 := 0; fi26 < 18; fi26++ {
        v27 := GenS1{F0: (v4 / (((v21 >> ((753) & 63))) | 1)), F1: uint64(v4), F2: v1}
        _ = fi26
        _ = v27
    }
    switch "" {
    case "foo bar":
        var v28 uint16 = 42
        _ = v28
    case "\n\t":
        for fi29 := 0; fi29 < 2; fi29++ {
            var v30 float32 = (-2.25 + 0.1)
            if (float64(((((v30 / -2.25) - 0.1) / (3.14159 + (0.1 - v30))) + v30)) > 1.5) {
                if v6 {
                    var v31 int64 = 749
                    if (v6 && ((fi29 <= fi29) || v6)) {
                        var v32 float32 = float32(float64(((v30 - (0.0 / 1.5)) + (v30 + (((((v30 - 1.5) * (v30 / v30)) * (0.1 / 3.14159)) * -2.25) / v30)))))
                        v30 -= v30
                        _ = v32
                    }
                    for fi33 := 0; fi33 < 6; fi33++ {
                        v34 := GenS1{F0: 120, F1: 4294967295, F2: (v1 - (v1 * v1))}
                        _ = fi33
                        _ = v34
                    }
                    _ = v31
                }
            }
            v6 = true
            _ = fi29
            _ = v30
        }
    case "hello":
        var v35 uint32 = 261
        _ = v35
    }
    for fi36 := 0; fi36 < 12; fi36++ {
        var v37 bool = false
        _ = fi36
        _ = v37
    }
    var v38 float64 = (3.14159 + 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
