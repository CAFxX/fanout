package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    v3 := make(map[int]uint32)
    var v4 float64 = 0.0
    var v5 float64 = float64(-2.25)
    for fi6 := 0; fi6 < 3; fi6++ {
        fmt.Println(fi6, v5)
        v4 += v5
        _ = fi6
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 bool = ((true || true) && (18446744073709551615 == 42))
        _ = fi7
        _ = v8
    }
    if (0 >= 100) {
        var v9 float32 = 0.1
        _ = v9
    }
    var v10 float32 = -2.25
    if false {
        v10 = 1e10
        for fi11 := 0; fi11 < 2; fi11++ {
            v12 := GenS1{F0: fi11, F1: -32768}
            _ = fi11
            _ = v12
        }
    } else {
        for fi13 := 0; fi13 < 10; fi13++ {
            if false {
                var v14 bool = false
                var v15 int32 = 42
                _ = v14
                _ = v15
            }
            _ = fi13
        }
    }
    v4 -= ((3.14159 - -2.25) - (v5 - 3.14159))
    v16 := []int{0}
    v17 := GenS1{F0: -1, F1: 127}
    v5 -= v4
    var v18 int32 = 1
    if (!("\n\t" < "a")) {
        if len(v16) > 0 {
            v16[0] = -9223372036854775808
        }
    } else {
        var v19 uint64 = 18446744073709551615
        delete(v2, ("" + ("foo bar" + "foo bar")))
        _ = v19
    }
    v5 = (-2.25 / (v4 * ((((((((3.14159 - v4) * ((v4 * v4) + v4)) / (0.1 - v5)) / ((v4 * float64(1e10)) + (((float64(v10) / v5) - -2.25) - 3.14159))) - (((((v5 + (1.5 / (0.1 * v5))) - (v4 - v4)) + v5) + (1e10 / ((v5 - (float64(0.0) / float64(v10))) * (((float64(v10) + ((v4 / (v4 + 0.0)) + ((v4 / v5) - 1e10))) * -2.25) / 1.0)))) + float64((float32(v4) * (v10 * -2.25))))) / 1.0) + (3.14159 * v5)) * v4)))
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
