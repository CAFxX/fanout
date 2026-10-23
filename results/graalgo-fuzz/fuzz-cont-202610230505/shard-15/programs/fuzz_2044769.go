package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int16; F1 int8; F2 uint32 }
type GenS3 struct { F0 int32; F1 int32 }

func main() {
    v1 := make(map[int]int)
    var v2 int16 = -32768
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]string)
    v5 := []int{9223372036854775807, -2147483648, 2147483647, 42}
    var v6 uint32 = 631
    switch 42 {
    case 9223372036854775807:
        v2 -= (v2 ^ (^(v2)))
    default:
        v7 := GenS1{F0: 42, F1: 127}
        _ = v7
    }
    v6 += (v6 * 42)
    var v8 int8 = -1
    fmt.Println(v2)
    fmt.Println(v6, v2)
    for fi9 := 0; fi9 < 4; fi9++ {
        delete(v1, (fi9 - 32768))
        v10 := GenS2{F0: v2, F1: (v8 / ((94) | 1)), F2: (v6 / ((v6) | 1))}
        var v11 uint16 = 0
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 int8 = v8
    v13 := v3["a"]
    v14 := GenS1{F0: 616, F1: 42}
    switch ("foo bar" + "a") {
    case "":
        var v15 uint8 = 1
        _ = v15
    default:
        fmt.Println(v12, v6, v6)
    }
    v16 := GenS3{F0: 1, F1: 0}
    v2 = (v2 << ((65535) & 63))
    switch 0 {
    case 32767:
        for fi17 := 0; fi17 < 18; fi17++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            for fi18 := 0; fi18 < 7; fi18++ {
                var v19 float32 = 0.0
                switch 9223372036854775807 {
                case 128:
                    v8 = -1
                }
                v20 := GenS3{F0: -2147483648, F1: 42}
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = fi17
        }
    case -9223372036854775808:
        v12 += (v8 | v12)
    case -1:
        v8 = v8
    default:
        v21 := GenS2{F0: 1, F1: int8(v2), F2: (^(v13))}
        _ = v21
    }
    v6 = v6
    v13 -= v13
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
