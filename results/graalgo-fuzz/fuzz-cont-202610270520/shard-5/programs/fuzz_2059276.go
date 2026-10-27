package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() uint {
    var v4 int32 = -1
    v4 = v4
    if ((!(0.1 != float64(float32(0.0)))) || true) {
        v5 := GenS1{F0: 0}
        _ = v5
    }
    _ = v4
    return 65535
}

func main() {
    var v1 string = "foo bar"
    var v2 int = -627
    var v3 float32 = (float32(1.5) * (float32(float64(-2.25)) * -2.25))
    v2 = 9223372036854775807
    v2 = (v2 + 2147483647)
    v6 := []int{}
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 int = -2147483648
        fmt.Println(v8, v2, v2)
        v2 = fi7
        _ = fi7
        _ = v8
    }
    v1 = (v1 + v1)
    if len(v6) > 0 {
        v6[0] = -1
    }
    var v9 int32 = 0
    _ = len(v6)
    var v10 float64 = (-2.25 * (float64(1e10) / 1.5))
    v2 += (v2 * (v2 | v2))
    _ = len(v6)
    fmt.Println(v9)
    fmt.Println(v9, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v9
}
