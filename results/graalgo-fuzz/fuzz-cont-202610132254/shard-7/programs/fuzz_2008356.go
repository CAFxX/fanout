package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint }
func genFn0(p0 uint16, p1 string) int32 {
    switch 32767 {
    case 127:
        fmt.Println(p1, p0)
    case -1:
        if false {
            fmt.Println(p0, p0)
            if false {
                var v5 float64 = -2.25
                _ = v5
            } else {
                v6 := GenS1{F0: -971, F1: 32767}
                _ = v6
            }
            var v7 bool = true
            _ = v7
        }
    }
    v8 := GenS2{F0: 9223372036854775807}
    _ = p0
    _ = p1
    _ = v8
    return 1
}

func main() {
    v1 := []string{"a"}
    var v2 float32 = (((((((0.1 - (float32(1.5) / (3.14159 + float32(1.5)))) * 0.1) + (1.5 * ((0.1 / 1.0) + (((0.0 * 3.14159) * 0.1) - (1.5 / 1.0))))) - 3.14159) - 1.5) * float32((3.14159 - ((float64(1.5) * 1.5) + 0.1)))) / (-2.25 * 3.14159))
    v3 := make(map[string]string)
    v4 := make(map[uint32]int)
    v2 = 1e10
    v2 = (((v2 - v2) + (-2.25 * v2)) + -2.25)
    v2 = float32(((((1e10 * 3.14159) * 0.0) / (((float64(float32(3.14159)) - (float64(v2) * (-2.25 + -2.25))) + 1e10) + ((3.14159 / 0.1) / (float64((v2 - (v2 / ((0.1 / 0.1) - v2)))) / 3.14159)))) + 0.0))
    fmt.Println(v2)
    v9 := GenS3{F0: 65535}
    var v10 float64 = 0.0
    if (!true) {
        v2 = v2
        var v11 int64 = 42
        _ = v11
    } else {
        v12 := GenS2{F0: 0}
        _ = v12
    }
    v13 := GenS2{F0: 249}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
