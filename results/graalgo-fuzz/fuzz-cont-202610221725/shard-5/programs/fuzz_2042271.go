package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint16 }
func genFn0(p0 int8, p1 uint, p2 int8) int64 {
    var v6 float64 = 1.5
    if (false || true) {
        if true {
            var v7 int32 = 1
            _ = v7
        }
        fmt.Println(p2, p2)
    } else {
        v8 := GenS1{F0: 2, F1: 32767, F2: 860}
        _ = v8
    }
    p0 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}
func genFn1() uint64 {
    var v9 int = 1
    switch "\n\t" {
    case "hello":
        fmt.Println(v9)
    default:
        v10 := []int{v9, 127, 42, v9}
        _ = v10
    }
    v11 := []int{}
    _ = v11
    _ = v9
    return 0
}

func main() {
    var v1 uint64 = 2
    var v2 uint16 = 0
    var v3 uint32 = 946
    var v4 uint = 651
    var v5 uint64 = v1
    var v12 int32 = 0
    var v13 float64 = 1e10
    v13 = v13
    v4 += (^(v4))
    switch 4294967295 {
    case 42:
        var v14 uint32 = (v3 >> ((815) & 63))
        _ = v14
    case 593:
        v13 -= (0.0 + 3.14159)
    }
    var v15 int32 = v12
    v15 -= v15
    for fi16 := 0; fi16 < 8; fi16++ {
        var v17 uint8 = 255
        var v18 uint8 = 10
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: (v5 + (v1 * (^(v1)))), F1: v2, F2: (v2 & v2)}
    v1 += (v1 * genFn1())
    v1 = (v1 * (v5 % (((v5 | (^(v1)))) | 1)))
    fmt.Println(v1, v15, v5)
    v20 := GenS1{F0: v5, F1: (v2 | v2), F2: 32767}
    if ((v2 >> ((v4) & 63)) == 42) {
        v4 -= (v4 >> ((v4) & 63))
    } else {
        var v21 int16 = 42
        var v22 int32 = v15
        _ = v21
        _ = v22
    }
    v4 += v4
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
