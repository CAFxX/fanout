package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 uint8 = 0
    var v2 uint16 = 2
    var v3 int32 = -521
    v4 := []string{}
    var v5 int8 = -7
    v6 := make(map[uint32]int)
    v7 := GenS1{F0: v2}
    v8 := GenS1{F0: (v2 / ((v2) | 1))}
    v5 -= (v5 % ((v5) | 1))
    switch 4294967295 {
    case 582:
        var v9 uint8 = (^(v1))
        _ = v9
    default:
        v10 := GenS1{F0: v2}
        _ = v10
    }
    var v11 uint32 = 832
    if (v1 != v1) {
        var v12 uint = 1
        delete(v6, v11)
        _ = v12
    }
    fmt.Println(v11, v3)
    switch 32768 {
    case -9223372036854775808:
        v13 := GenS2{F0: v2}
        _ = v13
    case 128:
        if true {
            switch 128 {
            case 843:
                v14 := GenS3{F0: -607}
                _ = v14
            default:
                var v15 float32 = (0.1 + (1.5 * 0.0))
                _ = v15
            }
            v6[(v11 & (^(v11)))] = 95
            var v16 int32 = -631
            _ = v16
        } else {
            fmt.Println(v5)
            v1 += 1
        }
    case 1:
        if ((!true) || false) {
            var v17 uint64 = 2
            if (1e10 > (0.1 - -2.25)) {
                delete(v6, (v11 - (^(v11))))
                if (!false) {
                    var v18 uint16 = 1
                    _ = v18
                }
            } else {
                if (!(((v2 % (((^(v2))) | 1)) != 42) && false)) {
                    var v19 int32 = (v3 ^ (v3 >> ((uint(v17)) & 63)))
                    v1 = v1
                    _ = v19
                }
            }
            _ = v17
        }
    }
    v11 -= v11
    var v20 uint8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
