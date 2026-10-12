package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int32; F2 uint }
type GenS3 struct { F0 int16 }

func main() {
    var v1 int8 = 1
    v2 := make(map[uint32]uint32)
    var v3 int16 = -30
    v4 := make(map[string]uint32)
    var v5 string = ""
    var v6 int64 = 1
    for fi7 := 0; fi7 < 7; fi7++ {
        v3 = v3
        _ = fi7
    }
    if (42 < 164) {
        v8 := GenS1{F0: 1}
        _ = v8
    } else {
        for fi9 := 0; fi9 < 10; fi9++ {
            fmt.Println(v3, v6)
            if (128 != 128) {
                v10 := GenS1{F0: 42}
                var v11 int = (-(fi9))
                if ((!(((42 != 640) || true) && (0.1 <= (1.5 - 3.14159)))) || true) {
                    v12 := GenS1{F0: 0}
                    switch "foo bar" {
                    case "foo bar":
                        var v13 bool = false
                        _ = v13
                    default:
                        fmt.Println(v11, fi9, v3)
                    }
                    var v14 int = -404
                    _ = v12
                    _ = v14
                }
                _ = v10
                _ = v11
            } else {
                v1 -= (v1 + (v1 * v1))
                v15 := GenS1{F0: 2147483647}
                _ = v15
            }
            fmt.Println(v3, v5, v5)
            _ = fi9
        }
    }
    var v16 uint64 = 2
    var v17 uint8 = 127
    var v18 uint8 = v17
    v6 = 880
    v17 = (^(v18))
    v18 += uint8(v16)
    v19 := GenS1{F0: 0}
    v17 -= 208
    _ = len(v2)
    v20 := GenS2{F0: 32767, F1: -578, F2: 1}
    var v21 string = ""
    if (false || false) {
        v17 = v18
        switch 128 {
        case -856:
            for fi22 := 0; fi22 < 7; fi22++ {
                var v23 uint16 = 281
                for fi24 := 0; fi24 < 14; fi24++ {
                    v25 := []int{4294967295, 1}
                    fmt.Println(v5, v3)
                    var v26 uint32 = 2147483648
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                v6 = v6
                _ = fi22
                _ = v23
            }
        case -665:
            v27 := GenS2{F0: 32767, F1: 2147483647, F2: 18446744073709551615}
            _ = v27
        }
        v28 := GenS3{F0: (v3 % (((v3 ^ (v3 * (v3 + v3)))) | 1))}
        _ = v28
    } else {
        v29 := GenS3{F0: (-(v3))}
        _ = v29
    }
    v30 := []int{-9223372036854775808, -9223372036854775808, 1}
    var v31 uint16 = 32767
    for fi32 := 0; fi32 < 9; fi32++ {
        v17 += (v17 / ((255) | 1))
        _ = fi32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
