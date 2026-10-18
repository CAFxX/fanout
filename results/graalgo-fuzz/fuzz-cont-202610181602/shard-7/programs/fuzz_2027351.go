package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint }

func main() {
    var v1 float32 = ((0.0 * 1.5) / 0.1)
    v2 := make(map[int]uint32)
    var v3 uint16 = 84
    var v4 int8 = -7
    for fi5 := 0; fi5 < 11; fi5++ {
        if (!true) {
            fmt.Println(v3, v3, fi5)
            v1 = 0.1
            for fi6 := 0; fi6 < 14; fi6++ {
                switch ("foo bar" + ("\n\t" + ("foo bar" + "hello"))) {
                case "\n\t":
                    if false {
                        fmt.Println(v1, v1)
                        if (6 != 90) {
                            v3 = (v3 % (((v3 - 0)) | 1))
                            var v7 int16 = -32768
                            v4 += v4
                            var v8 int32 = -14
                            _ = v7
                            _ = v8
                        }
                    }
                }
                v1 = v1
                _ = fi6
            }
        }
        v1 += 1e10
        _ = fi5
    }
    fmt.Println(v3, v4)
    fmt.Println(v3, v1, v3)
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := GenS1{F0: 4294967295, F1: 9223372036854775807, F2: 2147483648}
        _ = fi9
        _ = v10
    }
    fmt.Println(v3, v4, v4)
    var v11 uint8 = 1
    v4 -= 0
    v12 := []int{128, -1, 0, 42}
    var v13 int = 128
    switch (v13 % (((v13 & (^(v13)))) | 1)) {
    case 1:
        v14 := GenS2{F0: 1, F1: 18446744073709551615, F2: 305}
        _ = v14
    case -9223372036854775808:
        v11 += v11
    default:
        if false {
            var v15 uint16 = (v3 & (^(v3)))
            _ = v15
        } else {
            for fi16 := 0; fi16 < 3; fi16++ {
                v17 := GenS2{F0: 754, F1: 968, F2: uint(v11)}
                var v18 float32 = 1.5
                if true {
                    v4 += v4
                    v19 := GenS2{F0: 183, F1: 2, F2: 2}
                    fmt.Println(v1, v18, v1)
                    _ = v19
                } else {
                    v18 += (v18 / (v18 * v1))
                }
                _ = fi16
                _ = v17
                _ = v18
            }
        }
    }
    if (-1 > 18) {
        v20 := v2[32767]
        _ = v20
    }
    var v21 int32 = -2147483648
    for fi22 := 0; fi22 < 15; fi22++ {
        v21 = (v21 * -843)
        _ = fi22
    }
    switch "\n\t" {
    case "foo bar":
        v23 := GenS1{F0: 255, F1: 9223372036854775807, F2: 0}
        _ = v23
    default:
        var v24 uint = 517
        _ = v24
    }
    v25 := v2[v13]
    fmt.Println(v4, v1, v25)
    var v26 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
