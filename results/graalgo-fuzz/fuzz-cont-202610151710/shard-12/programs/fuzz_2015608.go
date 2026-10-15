package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
func genFn0(p0 uint16) uint8 {
    p0 += (^(p0))
    v7 := []int{-2147483648, 42, -710}
    p0 += p0
    _ = p0
    _ = v7
    return 255
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]string)
    var v3 string = ("a" + "hello")
    var v4 int16 = -32768
    var v5 bool = true
    var v6 int = 112
    v8 := []int{v6, (v6 * (v6 >> ((51) & 63))), int(v4)}
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := GenS1{F0: (-(fi9))}
        _ = fi9
        _ = v10
    }
    v6 += (^(v6))
    fmt.Println(v4, v4, v3)
    var v11 int64 = 741
    if v5 {
        var v12 int64 = 615
        if v5 {
            for fi13 := 0; fi13 < 1; fi13++ {
                if (0.1 < ((0.1 / 1.0) + 1.5)) {
                    v14 := GenS2{F0: v3}
                    _ = v14
                }
                var v15 int = (-(v6))
                var v16 float32 = 1.5
                _ = fi13
                _ = v15
                _ = v16
            }
        }
        var v17 int = (v6 & v6)
        _ = v12
        _ = v17
    } else {
        if (0 <= 4294967295) {
            var v18 bool = (v5 && (!true))
            v19 := GenS1{F0: -9223372036854775808}
            _ = v18
            _ = v19
        }
        var v20 uint8 = 2
        _ = v20
    }
    v6 = (v6 >> ((65535) & 63))
    if (v5 || ((0 > int8(v11)) && (((^(v4)) >= (v4 | v4)) && ((1.5 != float64(-2.25)) || ((127 == 0) || (v6 != (v6 >> ((2) & 63)))))))) {
        v4 += 32767
        v5 = (227 > 1)
        var v21 int = v6
        _ = v21
    } else {
        var v22 int64 = -162
        for fi23 := 0; fi23 < 2; fi23++ {
            if v5 {
                var v24 float64 = -2.25
                var v25 uint32 = 4294967295
                _ = v24
                _ = v25
            } else {
                var v26 uint = 65535
                fmt.Println(v3)
                _ = v26
            }
            _ = fi23
        }
        _ = v22
    }
    var v27 int8 = 100
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
