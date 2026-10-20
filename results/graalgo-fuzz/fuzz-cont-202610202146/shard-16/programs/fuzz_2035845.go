package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int32; F1 int8; F2 uint16 }

func main() {
    var v1 float32 = float32(3.14159)
    var v2 float32 = ((0.0 - 0.1) + v1)
    var v3 string = "a"
    var v4 int16 = 184
    v5 := make(map[string]int)
    v6 := []string{v3, "a", "hello", v3}
    v7 := GenS1{F0: 2, F1: 2147483648}
    var v8 uint = 65535
    fmt.Println(v3, v1)
    fmt.Println(v8, v4, v1)
    var v9 uint64 = 121
    v10 := GenS1{F0: 1, F1: 33}
    for fi11 := 0; fi11 < 4; fi11++ {
        v6 = append(v6, v3)
        v3 = v3
        if (!false) {
            var v12 int16 = (v4 * -593)
            var v13 float32 = float32(3.14159)
            v14 := GenS2{F0: -2147483648, F1: 0, F2: 32767}
            _ = v12
            _ = v13
            _ = v14
        } else {
            if (0.1 == 1.5) {
                switch v3 {
                case "":
                    if false {
                        var v15 float32 = (v2 / v2)
                        v2 = (float32((((0.0 * 1.5) - (1.5 + 3.14159)) + 0.0)) + ((0.0 - v2) * v1))
                        v2 = v15
                        _ = v15
                    }
                }
                fmt.Println(v8, v9)
            }
        }
        _ = fi11
    }
    fmt.Println(v1, v8)
    v1 += v1
    var v16 uint16 = 0
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
