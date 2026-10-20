package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint32 }
func genFn0(p0 bool, p1 string, p2 int64) uint16 {
    if (0 <= 0) {
        v4 := []int{4294967295, 128, -9223372036854775808, 234}
        _ = v4
    }
    fmt.Println(p0)
    var v5 int64 = p2
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := GenS1{F0: "foo bar", F1: 18446744073709551615}
        fmt.Println(fi6, p2, v5)
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{65535, 376, 65535, 12}
    var v3 int16 = 1
    _ = len(v2)
    if true {
        if (true && (float64(float32(-2.25)) <= (3.14159 * 0.1))) {
            v3 = v3
            switch 238 {
            case 401:
                var v8 int8 = -128
                _ = v8
            default:
                if false {
                    v3 -= v3
                } else {
                    v3 += -32768
                    var v9 bool = (-2.25 != (float32(1e10) * float32(float64(0.0))))
                    _ = v9
                }
            }
            for fi10 := 0; fi10 < 19; fi10++ {
                v3 -= (v3 >> ((65535) & 63))
                v11 := GenS1{F0: "hello", F1: 2}
                _ = fi10
                _ = v11
            }
        }
    }
    if false {
        var v12 int16 = v3
        _ = v12
    }
    v13 := GenS1{F0: "\n\t", F1: 4294967295}
    if false {
        v14 := GenS1{F0: "hello", F1: 42}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v14
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        v16 := GenS2{F0: 0, F1: 4294967295}
        _ = fi15
        _ = v16
    }
    switch -285 {
    case -783:
        _ = len(v1)
    case 32767:
        v3 = -1
    default:
        var v17 int = 721
        _ = v17
    }
    fmt.Println(v3)
    var v18 uint16 = 980
    switch 633 {
    case 822:
        v18 = (v18 << ((602) & 63))
    default:
        fmt.Println(v3)
    }
    v3 -= v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v3
}
