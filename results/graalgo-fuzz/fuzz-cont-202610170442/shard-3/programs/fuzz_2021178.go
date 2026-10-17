package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint64 }
func genFn0() int {
    v7 := GenS1{F0: 406}
    for fi8 := 0; fi8 < 4; fi8++ {
        fmt.Println(fi8)
        if ((!(!false)) && (112 <= 2)) {
            fmt.Println(fi8)
            var v9 uint8 = 2
            _ = v9
        } else {
            v10 := GenS2{F0: 622, F1: 65535, F2: 1}
            _ = v10
        }
        _ = fi8
    }
    _ = v7
    return 477
}
func genFn1(p0 string, p1 int32) int64 {
    var v11 int32 = (p1 % ((-762) | 1))
    v12 := []int{32768, 1, 32768}
    var v13 int32 = p1
    v14 := GenS2{F0: 143, F1: 758, F2: 1}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 127
    v3 := make(map[uint32]string)
    v4 := []uint32{}
    var v5 float64 = ((float64((1e10 / 0.1)) + (0.0 * float64(1e10))) / 1.5)
    var v6 uint = 1
    if (((3.14159 * v5) + (v5 * 1.5)) < float64((0.1 / float32((0.1 / v5))))) {
        v2 = 100
    }
    var v15 int32 = -355
    switch "hello" {
    case "":
        v16 := GenS1{F0: 2}
        _ = v16
    case "a":
        fmt.Println(v2)
    case "\n\t":
        var v17 uint32 = 65535
        _ = v17
    default:
        v6 = (v6 >> ((v6) & 63))
    }
    v18 := GenS1{F0: 42}
    fmt.Println(v5, v2, v2)
    var v19 uint32 = 719
    v19 = (v19 ^ v19)
    delete(v1, v19)
    if len(v4) > 0 {
        v4[0] = (v19 >> ((v6) & 63))
    }
    _ = len(v4)
    v2 += 5
    v5 += float64((0.1 * -2.25))
    v20 := GenS1{F0: 2}
    var v21 float32 = (1e10 - (-2.25 / 1.5))
    var v22 int32 = int32(v19)
    var v23 uint8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
