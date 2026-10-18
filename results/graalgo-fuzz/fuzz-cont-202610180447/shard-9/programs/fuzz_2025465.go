package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint; F2 string }

func main() {
    var v1 float64 = float64(1e10)
    var v2 uint8 = 238
    v3 := []uint32{215}
    v4 := make(map[string]string)
    var v5 float64 = v1
    if len(v3) > 0 {
        v3[0] = 2147483648
    }
    var v6 int = 128
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := GenS1{F0: fi7}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := v4["foo bar"]
            if false {
                var v12 uint8 = (^(v2))
                v13 := v4[v11]
                v1 -= ((v1 - (1.5 - 1.5)) * v5)
                _ = v12
                _ = v13
            } else {
                v2 -= (v2 & v2)
                if false {
                    v11 = v11
                }
                v6 -= 2147483647
                v2 = (v2 ^ 132)
            }
            if false {
                var v14 int64 = -9223372036854775808
                v5 = v5
                _ = v14
            } else {
                v15 := GenS2{F0: 101, F1: 255, F2: v11}
                _ = v15
            }
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    v1 = (float64(-2.25) * 3.14159)
    var v16 float64 = float64(-2.25)
    var v17 bool = false
    fmt.Println(v17, v17)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
