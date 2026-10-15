package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int8 }
func genFn0() uint32 {
    var v6 float64 = 3.14159
    v7 := GenS1{F0: -128, F1: 1}
    var v8 int8 = 42
    _ = v6
    _ = v7
    _ = v8
    return 89
}

func main() {
    var v1 float32 = 0.1
    var v2 int8 = 1
    v3 := []uint32{377, 0}
    var v4 uint32 = 2147483647
    var v5 int32 = -2147483648
    _ = len(v3)
    v1 = float32(0.0)
    var v9 uint64 = 18446744073709551615
    for fi10 := 0; fi10 < 4; fi10++ {
        v1 = float32(0.1)
        _ = fi10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 string = "a"
        _ = fi11
        _ = v12
    }
    v13 := GenS2{F0: (v4 / ((v4) | 1)), F1: (v2 * (^(v2)))}
    var v14 uint32 = genFn0()
    var v15 uint32 = (v14 & 2147483647)
    var v16 int8 = v2
    for fi17 := 0; fi17 < 4; fi17++ {
        for fi18 := 0; fi18 < 10; fi18++ {
            fmt.Println(fi17)
            var v19 uint = 994
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    v9 += v9
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
