package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 bool }

func main() {
    v1 := []uint32{}
    var v2 uint16 = 32767
    var v3 float32 = 0.1
    var v4 float32 = v3
    var v5 string = ""
    switch 431 {
    case -708:
        for fi6 := 0; fi6 < 4; fi6++ {
            v4 = v3
            _ = fi6
        }
    case 4294967295:
        for fi7 := 0; fi7 < 6; fi7++ {
            v2 = (v2 - v2)
            _ = fi7
        }
    default:
        var v8 int = 9223372036854775807
        _ = v8
    }
    fmt.Println(v5, v3)
    v9 := GenS1{F0: 2, F1: -263, F2: (((-1 < -918) || false) || true)}
    var v10 uint32 = 982
    switch "foo bar" {
    case "a":
        for fi11 := 0; fi11 < 18; fi11++ {
            _ = len(v1)
            _ = fi11
        }
    default:
        var v12 float32 = 1.5
        _ = v12
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        v5 = (("\n\t" + v5) + "\n\t")
        v3 += ((0.0 / 1.5) - 3.14159)
        if true {
            if len(v1) > 0 {
                v1[0] = (v10 << ((65535) & 63))
            }
            for fi14 := 0; fi14 < 9; fi14++ {
                switch v5 {
                case "foo bar":
                    if (5 > 1) {
                        var v15 uint8 = 116
                        switch (v5 + "a") {
                        case "hello":
                            if len(v1) > 0 {
                                v1[0] = (v10 % (((v10 - 2147483647)) | 1))
                            }
                        case "":
                            v16 := GenS1{F0: 699, F1: (fi14 * (fi14 / ((-850) | 1))), F2: ((^(v2)) <= v2)}
                            _ = v16
                        default:
                            var v17 int16 = 32767
                            _ = v17
                        }
                        _ = v15
                    } else {
                        fmt.Println(fi13)
                    }
                }
                var v18 float32 = float32(1e10)
                v19 := GenS1{F0: 1, F1: 42, F2: (!((!false) || true))}
                _ = fi14
                _ = v18
                _ = v19
            }
        } else {
            var v20 int = -9223372036854775808
            v1 = append(v1, 1)
            _ = v20
        }
        _ = fi13
    }
    var v21 int16 = 0
    v3 -= (v3 + 0.0)
    var v22 float32 = (float32(0.0) / 1.5)
    fmt.Println(v22)
    for fi23 := 0; fi23 < 2; fi23++ {
        var v24 uint = uint(v10)
        _ = fi23
        _ = v24
    }
    switch v5 {
    case "\n\t":
        for fi25 := 0; fi25 < 7; fi25++ {
            var v26 string = (("\n\t" + v5) + "")
            _ = fi25
            _ = v26
        }
    default:
        var v27 uint8 = 0
        _ = v27
    }
    for fi28 := 0; fi28 < 5; fi28++ {
        var v29 float64 = -2.25
        fmt.Println(v21)
        v30 := GenS1{F0: 18446744073709551615, F1: 32767, F2: true}
        _ = fi28
        _ = v29
        _ = v30
    }
    var v31 uint = 468
    v3 = v3
    var v32 uint = (v31 % (((v31 ^ 42)) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v9
}
