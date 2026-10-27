package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int; F1 uint32 }
func genFn0() uint64 {
    var v7 int64 = 0
    v8 := []int{-995}
    v8 = append(v8, 128)
    if true {
        v7 = v7
        v7 = v7
    } else {
        if ((42 >= 0) && (1 <= 2147483647)) {
            fmt.Println(v7)
            v9 := GenS1{F0: 1, F1: -58}
            for fi10 := 0; fi10 < 16; fi10++ {
                for fi11 := 0; fi11 < 10; fi11++ {
                    var v12 string = "hello"
                    var v13 int32 = -693
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                if len(v8) > 0 {
                    v8[0] = (fi10 & fi10)
                }
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = fi10
            }
            _ = v9
        } else {
            fmt.Println(v7)
        }
    }
    _ = v7
    _ = v8
    return 29
}
func genFn1(p0 uint32, p1 uint32) uint32 {
    var v14 int32 = 97
    v14 = (v14 % ((42) | 1))
    v14 += v14
    fmt.Println(p1, p0, v14)
    _ = p0
    _ = p1
    _ = v14
    return 115
}

func main() {
    var v1 uint8 = 128
    v2 := []uint32{2147483647}
    var v3 uint32 = 4294967295
    v4 := make(map[uint32]uint32)
    var v5 uint8 = (v1 * (v1 >> ((65535) & 63)))
    var v6 int32 = 0
    var v15 uint32 = (v3 >> ((18446744073709551615) & 63))
    switch -932 {
    case 42:
        v3 = 4294967295
    case 0:
        var v16 int32 = v6
        _ = v16
    }
    switch "" {
    case "foo bar":
        if false {
            v17 := GenS2{F0: 127, F1: 1}
            v5 += v1
            var v18 int64 = 9223372036854775807
            _ = v17
            _ = v18
        }
    default:
        for fi19 := 0; fi19 < 12; fi19++ {
            if len(v2) > 0 {
                v2[0] = v15
            }
            for fi20 := 0; fi20 < 17; fi20++ {
                fmt.Println(v15, v3, v15)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi20
            }
            _ = fi19
        }
    }
    v15 += (v15 % ((genFn1((v3 - genFn1((v3 / ((v15) | 1)), (v15 >> ((890) & 63)))), v3)) | 1))
    var v21 string = ("" + "a")
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v21)
    v22 := GenS2{F0: 2147483647, F1: (v15 % ((v3) | 1))}
    for fi23 := 0; fi23 < 19; fi23++ {
        var v24 int16 = -397
        _ = fi23
        _ = v24
    }
    var v25 float64 = (1.5 * (1.5 - 1.5))
    var v26 int8 = 109
    var v27 uint16 = 1
    var v28 float64 = v25
    v4[(^(v3))] = (^(v15))
    for fi29 := 0; fi29 < 8; fi29++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v15 += v15
        _ = fi29
    }
    v21 = v21
    var v30 int = 42
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
