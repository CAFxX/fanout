package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint8 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int32, p1 bool) int {
    switch 1 {
    case 32767:
        p1 = (1 < 2)
    default:
        v4 := GenS1{F0: 775, F1: "hello", F2: 127}
        _ = v4
    }
    var v5 int16 = 1000
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v5
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 1
    v3 := make(map[int]int)
    v2 = (v2 << (((^(v2))) & 63))
    v2 = v2
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 int8 = -128
        _ = fi6
        _ = v7
    }
    if true {
        v2 = (v2 ^ 2)
    }
    var v8 uint32 = 1
    v8 -= (^(v8))
    var v9 float32 = 0.0
    v2 = uint(v8)
    fmt.Println(v2, v8, v2)
    v2 += v2
    if (2147483647 >= 903) {
        if (true || true) {
            v3[genFn0(42, false)] = 371
        } else {
            v2 = 255
        }
        v9 = (v9 * v9)
    } else {
        for fi10 := 0; fi10 < 14; fi10++ {
            delete(v1, (v8 / ((v8) | 1)))
            for fi11 := 0; fi11 < 11; fi11++ {
                var v12 uint32 = (v8 << ((2) & 63))
                v13 := GenS2{F0: 0}
                _ = fi11
                _ = v12
                _ = v13
            }
            var v14 int32 = 2147483647
            _ = fi10
            _ = v14
        }
        switch "" {
        case "\n\t":
            v8 += 2147483647
        case "":
            v15 := GenS2{F0: -128}
            _ = v15
        default:
            v8 = (^(v8))
        }
    }
    v2 = (v2 - v2)
    for fi16 := 0; fi16 < 11; fi16++ {
        delete(v3, -409)
        _ = fi16
    }
    v17 := v3[-754]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
