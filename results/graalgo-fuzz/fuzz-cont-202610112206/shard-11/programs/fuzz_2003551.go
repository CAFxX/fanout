package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int = 648
    var v2 float32 = 1e10
    var v3 uint = 757
    v4 := []int{(-(v1)), (v1 >> (((v3 ^ 907)) & 63))}
    v2 -= 1.5
    v2 += v2
    var v5 uint16 = 42
    if (!((!false) || (float32(0.1) != v2))) {
        for fi6 := 0; fi6 < 1; fi6++ {
            var v7 float64 = 0.1
            v8 := GenS1{F0: (v1 % (((fi6 << (((v3 ^ (v3 * (v3 % (((v3 << (((v3 - (^(v3)))) & 63))) | 1))))) & 63))) | 1)), F1: -128}
            _ = fi6
            _ = v7
            _ = v8
        }
        var v9 int32 = 2147483647
        _ = v9
    } else {
        for fi10 := 0; fi10 < 4; fi10++ {
            fmt.Println(v3, v1, v5)
            _ = fi10
        }
        v11 := GenS1{F0: v1, F1: -128}
        if (!(!true)) {
            var v12 uint16 = (v5 * (v5 ^ (v5 << (((v3 ^ (v3 | 2))) & 63))))
            switch "hello" {
            case "foo bar":
                var v13 int32 = -383
                _ = v13
            case "a":
                var v14 int64 = int64(v3)
                _ = v14
            default:
                if false {
                    var v15 uint = 2
                    if len(v4) > 0 {
                        v4[0] = v1
                    }
                    _ = len(v4)
                    _ = v15
                }
            }
            _ = v12
        }
        _ = v11
    }
    v3 = (v3 + (v3 & v3))
    v2 -= v2
    if (!false) {
        v16 := GenS2{F0: 1}
        _ = v16
    }
    if false {
        for fi17 := 0; fi17 < 20; fi17++ {
            fmt.Println(v3)
            _ = fi17
        }
        var v18 int64 = 0
        var v19 uint64 = 1
        _ = v18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
