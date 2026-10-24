package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int16; F1 int8 }
func genFn0(p0 uint32, p1 uint32, p2 uint16) string {
    p2 += p2
    p2 = (p2 >> ((42) & 63))
    v4 := GenS1{F0: (p2 >> ((42) & 63)), F1: 18446744073709551615}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return "foo bar"
}

func main() {
    var v1 uint32 = 42
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]string)
    v5 := []int{0}
    switch 2147483647 {
    case -2147483648:
        v1 = (v1 * v1)
    }
    v1 = (v1 * (v1 | (v1 ^ v1)))
    v1 = (v1 ^ (v1 - v1))
    var v6 int32 = 1
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v1)
    v7 := GenS2{F0: 1, F1: -1}
    v8 := GenS1{F0: 255, F1: 621}
    v1 = 280
    var v9 uint = 516
    _ = len(v2)
    for fi10 := 0; fi10 < 1; fi10++ {
        v11 := v3[(v1 ^ 65535)]
        v2[v11] = (v1 | (v1 >> (((v9 % ((v9) | 1))) & 63)))
        var v12 uint16 = uint16(v1)
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 uint64 = 2
    var v14 uint8 = uint8(v1)
    v15 := GenS2{F0: 730, F1: 127}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
