package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint }

func main() {
    var v1 uint = 255
    v2 := []int{}
    var v3 uint32 = 0
    var v4 int64 = 90
    var v5 bool = false
    v2 = append(v2, 42)
    for fi6 := 0; fi6 < 10; fi6++ {
        switch fi6 {
        case 32768:
            v5 = v5
        case -2147483648:
            v4 = 9223372036854775807
        }
        _ = fi6
    }
    v1 = v1
    for fi7 := 0; fi7 < 14; fi7++ {
        v4 = (v4 ^ v4)
        _ = fi7
    }
    if (!((v3 * 0) == (^(v3)))) {
        var v8 int64 = (v4 * (^(v4)))
        if (!(230 <= -32768)) {
            switch 32767 {
            case 353:
                if v5 {
                    if (1e10 < ((-2.25 - -2.25) / -2.25)) {
                        for fi9 := 0; fi9 < 15; fi9++ {
                            var v10 int64 = (v4 / (((^(v8))) | 1))
                            v5 = v5
                            for fi11 := 0; fi11 < 10; fi11++ {
                                fmt.Println(fi9)
                                _ = fi11
                            }
                            _ = fi9
                            _ = v10
                        }
                        v2 = append(v2, 9223372036854775807)
                        v8 -= v4
                    }
                    var v12 float32 = ((-2.25 / 1e10) * (-2.25 / 1e10))
                    _ = v12
                }
            case -422:
                for fi13 := 0; fi13 < 8; fi13++ {
                    v14 := GenS1{F0: 0, F1: -149, F2: (^(v1))}
                    _ = fi13
                    _ = v14
                }
            case 4294967295:
                var v15 float32 = float32(1.5)
                _ = v15
            default:
                fmt.Println(v8, v1)
            }
            v5 = (2147483648 >= (^(v3)))
        }
        _ = v8
    }
    if (42 > 60) {
        if len(v2) > 0 {
            v2[0] = -9223372036854775808
        }
        v16 := GenS1{F0: 39, F1: 32767, F2: (^(v1))}
        v3 = v3
        _ = v16
    }
    v1 -= v1
    v4 = (v4 - 42)
    switch "" {
    case "\n\t":
        if (!(98 < 57)) {
            if (-2147483648 > -937) {
                if v5 {
                    fmt.Println(v4)
                    var v17 int8 = 0
                    var v18 int8 = (v17 - 127)
                    _ = v17
                    _ = v18
                }
            }
            var v19 uint = v1
            _ = v19
        } else {
            v4 += (v4 % (((v4 / (((v4 + (v4 | -9223372036854775808))) | 1))) | 1))
        }
    case "a":
        var v20 uint = (v1 ^ (v1 - 80))
        _ = v20
    case "hello":
        var v21 int8 = 1
        _ = v21
    }
    var v22 uint16 = 42
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        var v24 int64 = v4
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        switch 9223372036854775807 {
        case 32768:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 32767:
            v26 := GenS1{F0: 127, F1: (fi25 - fi25), F2: v1}
            _ = v26
        case -2147483648:
            v27 := GenS1{F0: 42, F1: fi25, F2: (v1 >> (((v1 >> ((181) & 63))) & 63))}
            _ = v27
        }
        switch "foo bar" {
        case "":
            for fi28 := 0; fi28 < 9; fi28++ {
                v22 = (^(v22))
                _ = fi28
            }
        default:
            var v29 int64 = v4
            _ = v29
        }
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
