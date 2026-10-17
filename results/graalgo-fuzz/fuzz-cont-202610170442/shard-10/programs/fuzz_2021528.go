package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }

func main() {
    var v1 int32 = -1
    var v2 uint64 = 705
    var v3 uint16 = 0
    v3 -= (v3 ^ v3)
    var v4 float32 = 3.14159
    if ("a" != "a") {
        var v5 float32 = (v4 + v4)
        _ = v5
    }
    var v6 float32 = v4
    switch -95 {
    case 32767:
        v3 = 2
    default:
        if false {
            for fi7 := 0; fi7 < 12; fi7++ {
                var v8 int32 = (v1 & v1)
                var v9 int64 = 42
                var v10 float64 = (float64((-2.25 - float32(0.0))) - float64(((1e10 + ((3.14159 * 0.1) + v4)) - ((3.14159 - v4) / 1.0))))
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
            for fi11 := 0; fi11 < 4; fi11++ {
                v2 = v2
                var v12 int32 = 2147483647
                v4 -= (v6 * float32(-2.25))
                _ = fi11
                _ = v12
            }
            var v13 bool = true
            _ = v13
        } else {
            v6 += float32(0.0)
            v2 += 2
            v14 := []int{}
            _ = v14
        }
    }
    if false {
        var v15 float64 = (float64((((v4 / (0.0 / (float32(3.14159) + (v6 + v6)))) + float32(float64(float32(((0.1 + 0.1) - 1e10))))) * 1e10)) / -2.25)
        v16 := GenS1{F0: 0, F1: 241}
        var v17 int16 = 32767
        _ = v15
        _ = v16
        _ = v17
    }
    if (!(false && (false && true))) {
        if ((false && (!false)) && (1 == -680)) {
            v18 := GenS1{F0: (v1 - (v1 - (v1 ^ (v1 / ((-2147483648) | 1))))), F1: 181}
            if false {
                v3 = (v3 / (((v3 & v3)) | 1))
            } else {
                v2 -= (v2 << ((4294967295) & 63))
            }
            _ = v18
        }
    } else {
        fmt.Println(v1)
        for fi19 := 0; fi19 < 19; fi19++ {
            v20 := GenS1{F0: 42, F1: 229}
            _ = fi19
            _ = v20
        }
    }
    v21 := GenS1{F0: v1, F1: 0}
    v22 := []int{1, 9223372036854775807}
    fmt.Println(v6, v3, v2)
    v23 := GenS1{F0: 1, F1: 1}
    if len(v22) > 0 {
        v22[0] = 32768
    }
    if true {
        var v24 uint = 65535
        _ = v24
    } else {
        switch -1 {
        case 32767:
            var v25 int32 = v1
            _ = v25
        }
        v26 := GenS1{F0: -2147483648, F1: 1}
        v3 += v3
        _ = v26
    }
    switch "\n\t" {
    case "foo bar":
        if false {
            v6 += (1e10 - 3.14159)
            var v27 uint16 = (^(v3))
            _ = v27
        } else {
            for fi28 := 0; fi28 < 2; fi28++ {
                fmt.Println(v6, v4)
                switch "" {
                case "a":
                    var v29 float64 = float64((v4 + (v4 + v6)))
                    _ = v29
                case "":
                    v30 := GenS1{F0: (^(v1)), F1: 2}
                    _ = v30
                default:
                    v31 := GenS1{F0: (v1 >> ((18446744073709551615) & 63)), F1: 0}
                    _ = v31
                }
                _ = fi28
            }
            var v32 bool = (v1 > 2147483647)
            fmt.Println(v2)
            _ = v32
        }
    default:
        switch "foo bar" {
        case "a":
            if false {
                v1 = (v1 / ((42) | 1))
                var v33 uint = 2
                var v34 uint = (v33 + 65535)
                _ = v33
                _ = v34
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v6
}
