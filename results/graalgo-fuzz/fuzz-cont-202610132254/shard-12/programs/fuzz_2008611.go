package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int32 }
func genFn0(p0 uint32) int {
    v5 := []int{38, 792, 232}
    v6 := GenS1{F0: -300, F1: (0 >= 159), F2: 28}
    v7 := GenS1{F0: 0, F1: (true || false), F2: 42}
    p0 = p0
    for fi8 := 0; fi8 < 2; fi8++ {
        fmt.Println(fi8, fi8)
        _ = len(v5)
        _ = fi8
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 127
}

func main() {
    v1 := []uint32{65535, 1, 65535}
    var v2 bool = (100 <= 1)
    var v3 int = 128
    var v4 bool = (v3 != v3)
    switch 32768 {
    case -2147483648:
        v4 = true
    case 862:
        fmt.Println(v2, v4, v4)
    case -9223372036854775808:
        if false {
            v1 = append(v1, 1)
            switch "\n\t" {
            case "\n\t":
                v4 = v4
            default:
                var v9 uint32 = 678
                _ = v9
            }
        }
    }
    var v10 int = (v3 | genFn0(2))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v11 int32 = 2147483647
    var v12 uint16 = 0
    switch "hello" {
    case "\n\t":
        v2 = v4
    case "":
        v3 -= (v10 << ((2) & 63))
    case "a":
        for fi13 := 0; fi13 < 8; fi13++ {
            for fi14 := 0; fi14 < 8; fi14++ {
                v3 -= (v3 - 32768)
                v1 = append(v1, 2147483648)
                _ = fi14
            }
            _ = fi13
        }
    }
    v3 -= (v3 << ((988) & 63))
    if (!(169 < 42)) {
        if v4 {
            v15 := GenS1{F0: -32768, F1: (v4 && (1e10 == ((1e10 / (((0.0 - 1e10) - -2.25) + 0.1)) - 1.5))), F2: -940}
            _ = v15
        }
        if v2 {
            if false {
                v12 += (v12 - v12)
                var v16 float64 = (0.1 * 0.0)
                _ = v16
            }
            var v17 int = 9223372036854775807
            v11 = 2147483647
            _ = v17
        }
        fmt.Println(v4, v2)
    }
    var v18 int8 = -1
    var v19 uint64 = 28
    var v20 uint64 = (v19 % ((v19) | 1))
    var v21 int64 = 1
    if (v11 > -580) {
        v22 := GenS1{F0: -32768, F1: true, F2: v11}
        _ = v22
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
