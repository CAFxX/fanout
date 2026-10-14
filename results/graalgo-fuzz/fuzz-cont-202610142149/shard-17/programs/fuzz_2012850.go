package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint8 }
type GenS2 struct { F0 int32; F1 int16 }
func genFn0(p0 uint8) uint {
    if (false && (-2147483648 < 4294967295)) {
        p0 = 127
        var v4 int32 = 2147483647
        fmt.Println(v4, v4)
        _ = v4
    }
    for fi5 := 0; fi5 < 12; fi5++ {
        p0 = p0
        _ = fi5
    }
    v6 := []int{403, -2147483648}
    fmt.Println(p0)
    p0 = p0
    _ = p0
    _ = v6
    return 18446744073709551615
}
func genFn1(p0 int, p1 int8, p2 uint32) uint8 {
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: "hello", F1: 32768, F2: 2}
        _ = fi7
        _ = v8
    }
    v9 := []int{p0, 63}
    var v10 int8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 103
}

func main() {
    var v1 int32 = -2147483648
    var v2 string = "hello"
    v3 := make(map[uint32]bool)
    var v11 int8 = 100
    var v12 int = -1
    var v13 int16 = int16(v12)
    v14 := v3[0]
    fmt.Println(v2, v11)
    for fi15 := 0; fi15 < 11; fi15++ {
        _ = len(v3)
        var v16 int = 4294967295
        _ = fi15
        _ = v16
    }
    if (v11 > (v11 - v11)) {
        var v17 int64 = 9223372036854775807
        _ = v17
    }
    v13 += (v13 % ((v13) | 1))
    v12 = v12
    v18 := GenS2{F0: (v1 / ((v1) | 1)), F1: (v13 + (^(v13)))}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
}
