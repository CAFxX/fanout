package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint64; F2 int64 }
func genFn0(p0 int) int32 {
    var v6 uint8 = 127
    v7 := GenS1{F0: v6}
    _ = p0
    _ = v6
    _ = v7
    return -990
}

func main() {
    var v1 float32 = float32(0.1)
    var v2 int64 = 9223372036854775807
    var v3 uint64 = 18446744073709551615
    var v4 bool = (!(v2 > (v2 % (((^(v2))) | 1))))
    v5 := make(map[string]int)
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS1{F0: 42}
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: v2, F1: 42, F2: v2}
    delete(v5, "hello")
    var v11 bool = false
    v12 := GenS1{F0: 16}
    v1 = v1
    var v13 int8 = 127
    v14 := v5[((("hello" + "a") + "a") + "")]
    switch 112 {
    case 1:
        var v15 int32 = 42
        _ = v15
    case 9223372036854775807:
        if v11 {
            v4 = ((v3 ^ v3) > (^(v3)))
        } else {
            var v16 uint8 = 42
            v1 = v1
            _ = v16
        }
    case 717:
        fmt.Println(v3, v11, v1)
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 uint16 = 255
        fmt.Println(v2, v1)
        var v19 bool = (2 == v3)
        _ = fi17
        _ = v18
        _ = v19
    }
    v20 := GenS2{F0: 1, F1: v3, F2: (^(v2))}
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
