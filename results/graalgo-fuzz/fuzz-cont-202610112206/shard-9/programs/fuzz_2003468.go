package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint8; F1 int; F2 uint32 }
func genFn0(p0 int32, p1 int64) string {
    var v7 float64 = 3.14159
    v7 += v7
    _ = p0
    _ = p1
    _ = v7
    return ""
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 249
    var v3 uint = 1
    var v4 bool = true
    var v5 int8 = -35
    var v6 int16 = 309
    var v8 uint32 = 2
    switch ("foo bar" + ("\n\t" + ("foo bar" + ("foo bar" + "foo bar")))) {
    case "foo bar":
        var v9 int16 = (v6 >> (((v3 / (((^(v3))) | 1))) & 63))
        _ = v9
    case "a":
        for fi10 := 0; fi10 < 11; fi10++ {
            v4 = (v6 > int16(v8))
            _ = fi10
        }
    default:
        v11 := GenS1{F0: 777, F1: (v6 + v6)}
        _ = v11
    }
    v12 := GenS1{F0: (v8 + 2147483648), F1: v6}
    switch 506 {
    case 127:
        var v13 uint64 = 0
        _ = v13
    }
    fmt.Println(v5)
    var v14 uint = v3
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 uint64 = 42
        var v17 uint64 = v16
        _ = fi15
        _ = v16
        _ = v17
    }
    v6 -= 42
    switch 1 {
    case -884:
        v18 := GenS1{F0: 2, F1: v6}
        _ = v18
    case -47:
        v6 = (-(v6))
    default:
        v6 += (v6 - v6)
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        v20 := []int{}
        v21 := GenS2{F0: v2, F1: fi19, F2: 2}
        v3 += (v3 ^ v3)
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println(v6, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
