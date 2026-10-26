package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int64; F2 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 bool; F2 int64 }
func genFn0(p0 int, p1 string, p2 uint64) int32 {
    v4 := GenS1{F0: 378}
    fmt.Println(p1, p2, p0)
    var v5 uint16 = 231
    var v6 uint16 = uint16(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := []uint32{2147483648, 0, 2, 42}
    var v2 int64 = 9223372036854775807
    var v3 float64 = (3.14159 + -2.25)
    fmt.Println(v3)
    switch "" {
    case "":
        var v7 float32 = float32(float64(0.0))
        _ = v7
    }
    v3 -= (v3 - v3)
    v8 := GenS1{F0: 4294967295}
    v3 += v3
    v2 = -225
    v9 := GenS2{F0: v2, F1: v2, F2: 842}
    var v10 bool = false
    fmt.Println(v3)
    v11 := GenS2{F0: 57, F1: -9223372036854775808, F2: 0}
    for fi12 := 0; fi12 < 3; fi12++ {
        v3 -= (v3 * v3)
        if len(v1) > 0 {
            v1[0] = 910
        }
        if (!v10) {
            var v13 int32 = -1
            v14 := GenS3{F0: ("foo bar" + "\n\t")}
            _ = v13
            _ = v14
        } else {
            switch 128 {
            case 42:
                v15 := GenS4{F0: ("" + "\n\t"), F1: ((v2 <= v2) && v10), F2: (v2 << ((65535) & 63))}
                _ = v15
            }
        }
        _ = fi12
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        v3 -= float64(0.1)
        fmt.Println(fi16)
        _ = fi16
    }
    v17 := GenS1{F0: 823}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
