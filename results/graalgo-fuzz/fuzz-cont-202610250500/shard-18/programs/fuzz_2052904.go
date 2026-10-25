package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 string }
func genFn0(p0 int16, p1 uint32, p2 string) int8 {
    var v5 string = p2
    fmt.Println(p0, p1, p0)
    var v6 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 uint16, p1 bool) string {
    v7 := GenS1{F0: 32768, F1: -1, F2: ""}
    p1 = (p1 && (3.14159 == 0.1))
    var v8 int = 414
    fmt.Println(v8, v8)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return ""
}

func main() {
    var v1 float64 = (float64(0.0) + 1.5)
    v2 := make(map[int]bool)
    v3 := make(map[string]bool)
    var v4 uint = 0
    _ = len(v2)
    v1 = (0.1 + (3.14159 * ((1.5 / 1.0) - 3.14159)))
    delete(v3, "hello")
    var v9 float64 = float64(float32(v1))
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 int32 = -2147483648
        for fi12 := 0; fi12 < 19; fi12++ {
            v1 = (v9 / (v9 + (float64(0.0) - (v9 + 0.0))))
            v13 := GenS1{F0: (^(fi10)), F1: 1, F2: "a"}
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    v4 += v4
    v14 := GenS1{F0: 0, F1: 1000, F2: "a"}
    for fi15 := 0; fi15 < 1; fi15++ {
        fmt.Println(v1, fi15, fi15)
        _ = fi15
    }
    switch 128 {
    case 0:
        fmt.Println(v9, v1)
    case -2147483648:
        v1 -= v9
    case -954:
        v4 = (v4 / ((v4) | 1))
    }
    if (float32(-2.25) > (0.0 + ((0.0 + 0.1) - 0.0))) {
        v4 -= v4
        var v16 uint16 = 366
        var v17 int8 = genFn0(-32768, 624, "a")
        _ = v16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
