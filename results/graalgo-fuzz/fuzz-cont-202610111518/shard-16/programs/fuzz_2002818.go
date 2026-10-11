package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
func genFn0(p0 int8, p1 int, p2 uint64) uint {
    fmt.Println(p1, p0)
    var v6 int32 = -1
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 uint = 255
        v8 = (v8 / ((673) | 1))
        _ = fi7
        _ = v8
    }
    p0 = (p0 >> ((4294967295) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{1, 42, 89, 177}
    var v3 string = "hello"
    v4 := []uint32{157}
    var v5 float32 = 1e10
    var v9 string = v3
    v10 := GenS1{F0: 1, F1: 127}
    v1["\n\t"] = true
    if len(v2) > 0 {
        v2[0] = 344
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        v3 = v9
        fmt.Println(v5, v3, fi11)
        _ = fi11
    }
    switch 42 {
    case 9223372036854775807:
        fmt.Println(v9, v9)
    case 127:
        if false {
            v5 += 0.0
            var v12 uint = 1
            _ = v12
        }
    case -1:
        for fi13 := 0; fi13 < 1; fi13++ {
            for fi14 := 0; fi14 < 14; fi14++ {
                v5 += v5
                _ = fi14
            }
            _ = fi13
        }
    }
    var v15 bool = true
    v5 -= v5
    var v16 bool = (0.1 <= 1e10)
    fmt.Println(v3, v9, v9)
    for fi17 := 0; fi17 < 9; fi17++ {
        fmt.Println(v9, fi17)
        var v18 uint64 = 18446744073709551615
        _ = fi17
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
