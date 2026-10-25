package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint32; F2 int16 }

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 2
    var v3 uint64 = (v2 + 18446744073709551615)
    var v4 int16 = -296
    var v5 int32 = -641
    fmt.Println(v2, v2)
    switch -1 {
    case 32767:
        var v6 int32 = v5
        _ = v6
    default:
        for fi7 := 0; fi7 < 6; fi7++ {
            v8 := GenS1{F0: 65535}
            v9 := GenS1{F0: 2}
            switch 2147483647 {
            case 127:
                for fi10 := 0; fi10 < 12; fi10++ {
                    v11 := GenS2{F0: 115, F1: 1, F2: (v4 | v4)}
                    fmt.Println(fi10, v3, v4)
                    _ = fi10
                    _ = v11
                }
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    fmt.Println(v2)
    v3 = v3
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 int = -778
        v13 += v13
        v1["a"] = "foo bar"
        _ = fi12
        _ = v13
    }
    var v14 int = 0
    v2 -= (v2 / ((v2) | 1))
    var v15 uint8 = 161
    var v16 uint64 = (^(v3))
    v5 = v5
    var v17 int8 = 42
    var v18 uint64 = 18446744073709551615
    var v19 float32 = 1e10
    var v20 int = v14
    var v21 uint8 = 0
    v4 = (v4 | (v4 & v4))
    _ = len(v1)
    v18 = v16
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
