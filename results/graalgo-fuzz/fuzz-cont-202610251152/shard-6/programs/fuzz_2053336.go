package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint16, p1 int16, p2 int16) int32 {
    v7 := GenS1{F0: 42, F1: p0, F2: (^(p0))}
    var v8 uint32 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -749
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 215
    var v3 float64 = -2.25
    var v4 bool = ((!true) || true)
    v5 := []uint32{0, 4294967295, 1}
    var v6 int16 = 0
    for fi9 := 0; fi9 < 7; fi9++ {
        v6 += v6
        if v4 {
            for fi10 := 0; fi10 < 1; fi10++ {
                v4 = (v6 == v6)
                _ = fi10
            }
        }
        _ = fi9
    }
    if (0 >= 662) {
        v2 = 422
        var v11 string = "foo bar"
        _ = v11
    }
    v3 = (1e10 * (v3 * (float64((0.0 + 0.1)) * ((((float64((-2.25 + 0.0)) * v3) + v3) / 1.0) - v3))))
    if (128 <= 2) {
        var v12 float32 = 1e10
        for fi13 := 0; fi13 < 8; fi13++ {
            if ((((fi13 | fi13) > fi13) && (("hello" + ((("hello" + "") + ("hello" + ("\n\t" + "a"))) + "hello")) != "")) && v4) {
                v4 = v4
                var v14 int8 = -62
                _ = v14
            } else {
                v15 := GenS2{F0: v2}
                _ = v15
            }
            _ = fi13
        }
        fmt.Println(v3, v2, v12)
        _ = v12
    } else {
        fmt.Println(v3, v4)
    }
    var v16 int16 = v6
    v4 = false
    v17 := GenS2{F0: 218}
    v18 := GenS1{F0: -1, F1: 919, F2: 2}
    var v19 float64 = v3
    var v20 bool = (!v4)
    v2 -= v2
    v3 = 3.14159
    v6 += v6
    if len(v5) > 0 {
        v5[0] = 0
    }
    var v21 int64 = 213
    var v22 int8 = 0
    var v23 int32 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
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
