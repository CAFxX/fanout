package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() int32 {
    var v5 int = -672
    var v6 float32 = 1.5
    _ = v5
    _ = v6
    return 2147483647
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[string]bool)
    v3 := make(map[int]uint32)
    var v4 float64 = float64(3.14159)
    v1 = (v4 + v1)
    if false {
        v4 += (v1 + 1.5)
    }
    if true {
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 bool = (42 >= -9223372036854775808)
            var v9 string = "a"
            var v10 uint = 18446744073709551615
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
    } else {
        v11 := []int{32767, 427, 1, -459}
        var v12 uint32 = 0
        _ = v11
        _ = v12
    }
    v13 := []int{-628, 565, 9223372036854775807, 2147483647}
    v3[401] = 4294967295
    if (true || true) {
        v1 = v1
        switch 208 {
        case -910:
            for fi14 := 0; fi14 < 12; fi14++ {
                fmt.Println(fi14, v4, v4)
                if false {
                    v1 = v1
                    var v15 int16 = -1
                    _ = v15
                }
                v16 := GenS1{F0: 1}
                _ = fi14
                _ = v16
            }
        }
    }
    v4 -= (0.1 * ((float64((3.14159 + 3.14159)) + 0.0) + (1.5 / 1.0)))
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
