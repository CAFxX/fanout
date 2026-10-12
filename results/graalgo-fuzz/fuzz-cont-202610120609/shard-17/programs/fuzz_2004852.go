package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() int16 {
    var v4 uint32 = 235
    v4 = 42
    for fi5 := 0; fi5 < 17; fi5++ {
        fmt.Println(v4, fi5)
        _ = fi5
    }
    v4 += v4
    _ = v4
    return 0
}
func genFn1(p0 int8, p1 uint64) bool {
    var v6 uint16 = 65535
    var v7 float32 = 1.5
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return true
}

func main() {
    v1 := []int{128, -1, 32767, -2147483648}
    v2 := make(map[int]int)
    var v3 uint16 = 93
    var v8 int16 = 32767
    v3 += (v3 * (v3 / (((v3 + (v3 << ((823) & 63)))) | 1)))
    v9 := GenS1{F0: 128}
    fmt.Println(v8, v8)
    for fi10 := 0; fi10 < 3; fi10++ {
        if true {
            for fi11 := 0; fi11 < 20; fi11++ {
                fmt.Println(v3)
                _ = fi11
            }
            v3 = v3
            v12 := GenS1{F0: fi10}
            _ = v12
        }
        for fi13 := 0; fi13 < 14; fi13++ {
            v3 = v3
            v14 := GenS1{F0: fi13}
            _ = fi13
            _ = v14
        }
        _ = fi10
    }
    var v15 float32 = (0.0 + float32(float64((float32(float64(0.0)) * 1e10))))
    v8 = 163
    switch "foo bar" {
    case "hello":
        v16 := GenS1{F0: 127}
        _ = v16
    case "":
        fmt.Println(v8)
    case "a":
        if ((^(v3)) > 0) {
            v2[32767] = 73
        }
    default:
        v3 = (v3 / (((v3 | v3)) | 1))
    }
    v15 -= v15
    var v17 int64 = 9223372036854775807
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
