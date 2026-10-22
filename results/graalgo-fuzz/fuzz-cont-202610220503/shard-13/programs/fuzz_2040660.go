package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint32) string {
    for fi5 := 0; fi5 < 6; fi5++ {
        p0 -= 572
        _ = fi5
    }
    fmt.Println(p0)
    _ = p0
    return "hello"
}
func genFn1(p0 int8, p1 int, p2 uint64) bool {
    var v6 int32 = -2147483648
    v7 := []int{-1, 405, (p1 >> ((0) & 63))}
    if len(v7) > 0 {
        v7[0] = -915
    }
    p2 = 2
    var v8 bool = (!(1e10 != 0.1))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return v8
}

func main() {
    var v1 int32 = -899
    var v2 uint8 = 128
    v3 := []uint32{42, 2147483647}
    v4 := make(map[int]uint32)
    v9 := GenS1{F0: 2}
    for fi10 := 0; fi10 < 8; fi10++ {
        _ = len(v4)
        _ = fi10
    }
    v1 -= v1
    fmt.Println(v1, v1)
    switch 128 {
    case 1:
        v2 = (v2 - (v2 % ((255) | 1)))
    default:
        for fi11 := 0; fi11 < 15; fi11++ {
            fmt.Println(fi11, v1, v1)
            _ = fi11
        }
    }
    delete(v4, -476)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v12 := GenS1{F0: 1}
    switch 885 {
    case -213:
        _ = len(v3)
    case -959:
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 string = ""
            _ = fi13
            _ = v14
        }
    case 9223372036854775807:
        if (1.5 <= float32(1.5)) {
            var v15 float64 = float64(-2.25)
            var v16 float32 = -2.25
            fmt.Println(v16, v1)
            _ = v15
            _ = v16
        }
    default:
        v2 = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
