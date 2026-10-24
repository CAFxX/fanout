package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 bool }
func genFn0(p0 uint64, p1 uint, p2 uint64) uint16 {
    p2 += p2
    v5 := []int{}
    for fi6 := 0; fi6 < 5; fi6++ {
        if true {
            v5 = append(v5, (fi6 << ((65535) & 63)))
            v7 := GenS1{F0: 410, F1: p0, F2: (false || true)}
            _ = v7
        }
        p0 = p2
        var v8 int16 = 32767
        _ = fi6
        _ = v8
    }
    p2 -= p0
    v9 := GenS1{F0: 0, F1: (p2 % ((4294967295) | 1)), F2: (412 <= 2)}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v9
    return 368
}

func main() {
    var v1 uint = 0
    v2 := make(map[uint32]int)
    var v3 uint16 = 255
    var v4 uint64 = 329
    var v10 bool = false
    var v11 int8 = -112
    v10 = ((0.0 >= ((-2.25 * 1.5) - 0.0)) || ((1.5 + ((0.0 / -2.25) * 1.5)) <= 0.0))
    var v12 int = -2147483648
    v12 -= 9223372036854775807
    fmt.Println(v12, v11, v1)
    var v13 uint8 = 55
    switch -2147483648 {
    case -1:
        var v14 float32 = 0.0
        _ = v14
    case 606:
        var v15 int16 = 0
        _ = v15
    case -113:
        v16 := GenS1{F0: -243, F1: 490, F2: ((v4 * v4) == (v4 % ((v4) | 1)))}
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
