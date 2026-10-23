package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int32 }
func genFn0(p0 bool, p1 int64, p2 uint32) uint {
    p2 = p2
    if false {
        p1 = p1
        v5 := GenS1{F0: 294, F1: 42, F2: 2147483647}
        v6 := []int{825, 1, 0}
        _ = v5
        _ = v6
    } else {
        var v7 int32 = 777
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 179
}

func main() {
    v1 := []uint32{42, 42}
    var v2 int64 = 1
    v3 := make(map[uint32]int)
    v4 := []uint32{2147483648}
    for fi8 := 0; fi8 < 10; fi8++ {
        v2 -= (v2 | v2)
        if true {
            v2 = (v2 ^ v2)
        } else {
            v9 := GenS1{F0: 255, F1: 255, F2: 42}
            v2 = (-(v2))
            var v10 bool = (0 == 255)
            _ = v10
            _ = v9
        }
        _ = len(v1)
        _ = fi8
    }
    var v11 string = "a"
    var v12 int32 = -1
    v12 = v12
    v13 := GenS1{F0: 842, F1: 1, F2: v12}
    for fi14 := 0; fi14 < 9; fi14++ {
        switch v11 {
        case "foo bar":
            for fi15 := 0; fi15 < 13; fi15++ {
                var v16 string = ((v11 + ("a" + v11)) + (v11 + ""))
                var v17 int = fi15
                v18 := GenS1{F0: genFn0((1 < 208), (v2 ^ 42), 117), F1: 42, F2: v12}
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
        case "":
            v19 := v3[2147483647]
            _ = v19
        default:
            v12 -= v12
        }
        var v20 int64 = (v2 | (^(v2)))
        _ = fi14
        _ = v20
    }
    var v21 float64 = (float64(float32((float64((3.14159 + 1e10)) / -2.25))) - 3.14159)
    var v22 string = v11
    v22 = v11
    v3[2147483647] = -1
    var v23 uint64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
