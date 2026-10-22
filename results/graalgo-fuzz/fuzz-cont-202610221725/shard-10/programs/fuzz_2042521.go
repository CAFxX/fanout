package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint16; F1 uint; F2 int64 }
type GenS4 struct { F0 uint32; F1 uint16; F2 uint8 }

func main() {
    v1 := []int{403, 32768}
    var v2 uint16 = 32767
    var v3 uint32 = 658
    switch 9223372036854775807 {
    case -2147483648:
        v4 := GenS1{F0: 65535}
        _ = v4
    case 127:
        v2 -= v2
    default:
        v3 += uint32(v2)
    }
    v2 += 706
    v5 := GenS1{F0: (v2 | (^(v2)))}
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            if len(v1) > 0 {
                v1[0] = fi6
            }
            v8 := GenS2{F0: fi7, F1: 127}
            _ = fi7
            _ = v8
        }
        v3 -= (v3 * v3)
        _ = fi6
    }
    v9 := GenS3{F0: (v2 * (v2 * (v2 - 255))), F1: uint(v3), F2: -9223372036854775808}
    fmt.Println(v2, v2)
    var v10 uint32 = 42
    if true {
        if true {
            v3 += (^(v10))
            var v11 uint64 = 2
            _ = v11
        }
        v12 := GenS1{F0: 857}
        if false {
            if (65535 > 42) {
                switch -9223372036854775808 {
                case -2147483648:
                    var v13 uint32 = (v3 - (^(v10)))
                    _ = v13
                case -1:
                    var v14 int8 = -62
                    _ = v14
                }
                v3 = 496
            }
        }
        _ = v12
    } else {
        switch "hello" {
        case "a":
            var v15 float32 = 0.0
            _ = v15
        default:
            v3 = 520
        }
    }
    _ = len(v1)
    if false {
        v10 = (^(v3))
        var v16 int = 4294967295
        _ = v16
    } else {
        v17 := GenS4{F0: (v10 + v3), F1: v2, F2: 162}
        _ = v17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v5
    _ = v9
}
