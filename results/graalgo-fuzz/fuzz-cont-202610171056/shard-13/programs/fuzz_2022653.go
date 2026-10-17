package main

import (
    "fmt"
)


func main() {
    var v1 bool = ((("" + "foo bar") != "") && false)
    var v2 int8 = -128
    v3 := make(map[uint32]uint32)
    var v4 int16 = 1
    fmt.Println(v1, v4, v2)
    if ((v2 & v2) == (-(v2))) {
        v5 := []int{-904}
        _ = v5
    } else {
        var v6 uint = 2
        var v7 uint16 = 672
        for fi8 := 0; fi8 < 10; fi8++ {
            _ = len(v3)
            v6 = v6
            _ = fi8
        }
        _ = v6
        _ = v7
    }
    var v9 uint = 402
    v9 -= v9
    var v10 float64 = float64(3.14159)
    for fi11 := 0; fi11 < 2; fi11++ {
        if (!(((3.14159 + ((((1e10 - ((((-2.25 * float32((v10 * float64((1e10 * (float32((0.1 + (v10 / v10))) + ((3.14159 - 3.14159) - (1e10 + float32(v10))))))))) + float32(0.1)) / ((((1.5 - float32(0.1)) / 3.14159) * 1.5) - 1.5)) / ((-2.25 / (1.5 - 1e10)) + 1e10))) / -2.25) / (3.14159 - 0.1)) - 0.1)) - (0.1 + ((0.1 * 3.14159) / (-2.25 / 1e10)))) != 0.0)) {
            var v12 int32 = 42
            _ = v12
        } else {
            _ = len(v3)
        }
        _ = fi11
    }
    fmt.Println(v1, v10)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
