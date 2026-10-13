package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint }
type GenS2 struct { F0 int8; F1 uint16; F2 int }
func genFn0(p0 uint8, p1 int) string {
    p0 += p0
    v6 := []int{(p1 / (((p1 | 9223372036854775807)) | 1))}
    p0 = (^(p0))
    v7 := GenS1{F0: 42, F1: 0, F2: 255}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return "a"
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]int)
    var v3 float32 = float32((float64(0.0) + (0.0 + ((-2.25 + (3.14159 / -2.25)) - 3.14159))))
    v4 := make(map[uint32]int)
    var v5 int16 = 42
    var v8 uint64 = 1
    v9 := GenS1{F0: v8, F1: 10, F2: 0}
    switch 42 {
    case 42:
        v10 := GenS2{F0: 61, F1: 1, F2: 127}
        _ = v10
    case 32767:
        var v11 int8 = 100
        _ = v11
    case 936:
        v5 = (v5 ^ v5)
    }
    var v12 uint64 = v8
    v13 := GenS1{F0: (v12 + (^(v12))), F1: 0, F2: 520}
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 int16 = (^(v5))
        v8 = (v8 ^ 1)
        v16 := v1[(("hello" + "foo bar") + ("" + "hello"))]
        _ = fi14
        _ = v15
        _ = v16
    }
    v12 = v12
    v5 = 1000
    for fi17 := 0; fi17 < 19; fi17++ {
        v12 -= v8
        var v18 int16 = (v5 / ((v5) | 1))
        _ = fi17
        _ = v18
    }
    fmt.Println(v8, v3)
    fmt.Println(v12, v8, v8)
    v19 := []int{}
    if (4294967295 < 65535) {
        var v20 int16 = 744
        var v21 int8 = -128
        _ = v20
        _ = v21
    }
    v3 += (((float32(((1e10 + (((3.14159 + (float64(((v3 - float32(1.5)) / 1e10)) + (float64((v3 * -2.25)) - (3.14159 - (-2.25 / 1.0))))) + 3.14159) - 0.0)) / float64(1.5))) * v3) + 1.5) + float32(((1e10 + 1e10) * ((float64((3.14159 / v3)) + -2.25) / float64((((1.5 - 1.5) - 1.5) - -2.25))))))
    if false {
        v8 = v8
        _ = len(v19)
        var v22 int8 = -91
        _ = v22
    } else {
        v8 = v12
    }
    var v23 int64 = 955
    v3 -= 0.0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
