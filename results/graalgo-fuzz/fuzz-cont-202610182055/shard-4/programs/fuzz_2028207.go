package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint }
func genFn0(p0 int32) bool {
    fmt.Println(p0)
    v5 := []int{9223372036854775807, 32768, 574}
    p0 = p0
    _ = p0
    _ = v5
    return false
}

func main() {
    var v1 int8 = -60
    var v2 bool = ((42 != 255) || true)
    var v3 int = -295
    var v4 int64 = 498
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := []int{(v3 * v3), (fi6 & 127)}
        if (!v2) {
            v4 -= v4
        }
        var v8 float32 = (1.5 - 1.5)
        _ = fi6
        _ = v7
        _ = v8
    }
    v1 += (v1 >> ((374) & 63))
    var v9 float64 = 3.14159
    if v2 {
        v3 += v3
        v9 -= v9
    } else {
        v10 := GenS1{F0: 32767, F1: v2, F2: 0}
        v11 := []int{(v3 ^ (^(v3)))}
        v3 += v3
        _ = v10
        _ = v11
    }
    var v12 string = ("" + ("hello" + "foo bar"))
    if true {
        if (42 < 42) {
            v13 := GenS1{F0: v3, F1: true, F2: 252}
            _ = v13
        }
    } else {
        fmt.Println(v3, v12, v2)
        v14 := []int{v3, (v3 | (v3 << ((4294967295) & 63))), (v3 + v3), -290}
        _ = v14
    }
    switch 127 {
    case 743:
        var v15 uint64 = 0
        _ = v15
    }
    v9 = 1e10
    var v16 uint = 18446744073709551615
    switch v3 {
    case 42:
        v1 = (v1 >> ((v16) & 63))
    default:
        v17 := GenS2{F0: 538, F1: (v16 >> ((v16) & 63))}
        _ = v17
    }
    v1 += 42
    var v18 float32 = 0.0
    v19 := GenS1{F0: v3, F1: true, F2: 7}
    v20 := []int{}
    if len(v20) > 0 {
        v20[0] = (v3 / (((v3 >> (((v16 >> (((v16 | (v16 ^ v16))) & 63))) & 63))) | 1))
    }
    v21 := GenS1{F0: (v3 & (v3 + -9223372036854775808)), F1: ((v1 / ((v1) | 1)) == (v1 + v1)), F2: 2}
    for fi22 := 0; fi22 < 17; fi22++ {
        var v23 uint32 = 4294967295
        for fi24 := 0; fi24 < 19; fi24++ {
            fmt.Println(v12, v9, v1)
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            _ = fi24
        }
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v9
}
