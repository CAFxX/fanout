package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int64; F1 int32; F2 bool }

func main() {
    var v1 int64 = 0
    v2 := []string{(("hello" + "") + "a"), "foo bar"}
    var v3 float32 = (1.5 / 1e10)
    var v4 bool = (-9223372036854775808 >= (v1 ^ v1))
    v4 = false
    var v5 uint32 = 65535
    switch -638 {
    case 1:
        fmt.Println(v3, v1, v3)
    case 127:
        var v6 int = 32767
        _ = v6
    case -727:
        if (!v4) {
            for fi7 := 0; fi7 < 20; fi7++ {
                if (-676 != -32768) {
                    var v8 int8 = -75
                    v4 = ((!v4) && v4)
                    fmt.Println(v3, v3)
                    _ = v8
                }
                v1 += (v1 ^ v1)
                v9 := GenS1{F0: 449, F1: (v1 % ((1) | 1))}
                _ = fi7
                _ = v9
            }
        }
    default:
        v3 += (v3 + (float32((((float64(-2.25) + 1e10) * 1e10) / 1.5)) - v3))
    }
    var v10 int8 = -102
    v11 := GenS2{F0: v1, F1: -1, F2: (v1 <= (v1 | (v1 % (((v1 / (((^(v1))) | 1))) | 1))))}
    switch ("a" + "a") {
    case "":
        v4 = v4
    default:
        _ = len(v2)
    }
    var v12 float32 = ((float32(-2.25) - v3) + 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
