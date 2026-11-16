package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int32 }
func genFn0(p0 int8, p1 uint) uint8 {
    p1 = p1
    switch ("foo bar" + "\n\t") {
    case "hello":
        fmt.Println(p0, p0)
    }
    switch 32767 {
    case 2147483647:
        var v6 int64 = -9223372036854775808
        _ = v6
    case 9223372036854775807:
        var v7 int16 = -1
        _ = v7
    default:
        p1 = (^(p1))
    }
    v8 := GenS1{F0: 1, F1: (p0 + 127), F2: -1}
    _ = p0
    _ = p1
    _ = v8
    return 62
}

func main() {
    var v1 int32 = -771
    var v2 bool = true
    v3 := make(map[int]int)
    var v4 bool = true
    var v5 float32 = (0.0 - ((0.0 * 1e10) / (-2.25 / 0.1)))
    v9 := GenS1{F0: 32767, F1: 100, F2: -1}
    v10 := GenS1{F0: -482, F1: 100, F2: 2147483647}
    for fi11 := 0; fi11 < 10; fi11++ {
        fmt.Println(v2)
        fmt.Println(v5, v4, fi11)
        _ = fi11
    }
    v12 := v3[32767]
    if (130 < 0) {
        fmt.Println(v2, v2)
        var v13 uint8 = 128
        _ = v13
    } else {
        v3[(v12 + v12)] = (v12 << ((582) & 63))
    }
    var v14 int16 = 42
    for fi15 := 0; fi15 < 18; fi15++ {
        v2 = (4294967295 > 479)
        _ = fi15
    }
    v12 += v12
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
