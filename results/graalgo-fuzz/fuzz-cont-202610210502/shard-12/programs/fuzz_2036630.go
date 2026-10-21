package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int8 }

func main() {
    var v1 uint64 = 4294967295
    var v2 uint = 65535
    var v3 float64 = 0.0
    var v4 uint32 = 848
    v5 := make(map[string]string)
    var v6 uint32 = (v4 << ((42) & 63))
    v7 := GenS1{F0: 616, F1: 32767}
    fmt.Println(v2)
    if (false && (v1 <= 42)) {
        v3 += 1.5
        _ = len(v5)
        switch ("foo bar" + "hello") {
        case "hello":
            v8 := GenS2{F0: 14}
            _ = v8
        }
    } else {
        for fi9 := 0; fi9 < 3; fi9++ {
            v6 += (v4 | v6)
            _ = fi9
        }
    }
    if false {
        if (false && true) {
            var v10 int8 = -1
            v11 := GenS1{F0: -2147483648, F1: -125}
            _ = v10
            _ = v11
        } else {
            if false {
                var v12 uint16 = 65535
                for fi13 := 0; fi13 < 16; fi13++ {
                    var v14 uint64 = (v1 >> (((v2 | 18446744073709551615)) & 63))
                    switch ("" + "foo bar") {
                    case "\n\t":
                        fmt.Println(v4)
                    default:
                        if (128 < 170) {
                            for fi15 := 0; fi15 < 8; fi15++ {
                                var v16 uint16 = (v12 + (v12 | v12))
                                v2 -= (v2 / ((v2) | 1))
                                v5[("" + "a")] = "hello"
                                _ = fi15
                                _ = v16
                            }
                        }
                    }
                    _ = fi13
                    _ = v14
                }
                _ = v12
            } else {
                var v17 bool = (-2.25 < (float64((0.1 - 1.5)) * (v3 + float64(1.5))))
                if true {
                    var v18 uint16 = 342
                    v1 -= v1
                    var v19 uint16 = v18
                    _ = v18
                    _ = v19
                }
                _ = v17
            }
            v2 -= 5
            var v20 int = -446
            _ = v20
        }
        var v21 bool = true
        var v22 uint8 = 128
        _ = v21
        _ = v22
    }
    var v23 int8 = 0
    if (1e10 < 0.0) {
        var v24 int8 = (v23 - (v23 / (((v23 ^ (v23 >> (((v2 + v2)) & 63)))) | 1)))
        var v25 int32 = 42
        if (0.0 > (v3 - v3)) {
            v26 := GenS1{F0: v25, F1: 32767}
            v5[""] = "foo bar"
            _ = len(v5)
            _ = v26
        }
        _ = v24
        _ = v25
    }
    var v27 int64 = -277
    v28 := []int{81, 32768}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
