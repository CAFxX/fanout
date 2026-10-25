package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 uint64 }
func genFn0(p0 uint64, p1 bool, p2 uint64) int8 {
    p0 = 2
    var v5 int32 = -2147483648
    v6 := []int{}
    var v7 string = "foo bar"
    v5 = (v5 << ((4294967295) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return -1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = float32((1.5 * ((0.1 - 3.14159) * float64(1e10))))
    v3 := make(map[uint32]bool)
    var v4 uint32 = 65535
    var v8 uint16 = 42
    v9 := GenS1{F0: 1}
    v10 := GenS2{F0: 149, F1: 4294967295}
    var v11 uint = 18446744073709551615
    for fi12 := 0; fi12 < 17; fi12++ {
        v13 := GenS2{F0: 144, F1: 2}
        for fi14 := 0; fi14 < 9; fi14++ {
            fmt.Println(v2, v2)
            fmt.Println(v4, v11)
            switch "hello" {
            case "hello":
                var v15 uint16 = 65535
                _ = v15
            case "\n\t":
                v16 := v3[v4]
                _ = v16
            default:
                v4 += v4
            }
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    v2 = v2
    v11 = v11
    var v17 int32 = -419
    v18 := GenS1{F0: 1}
    for fi19 := 0; fi19 < 7; fi19++ {
        if false {
            v11 = v11
        }
        _ = fi19
    }
    var v20 uint32 = uint32(v11)
    v2 += 1e10
    fmt.Println(v8)
    v21 := []int{32768, -973, -9223372036854775808, 1}
    var v22 string = "hello"
    switch 42 {
    case 32767:
        v17 -= v17
    default:
        _ = len(v21)
    }
    v20 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
