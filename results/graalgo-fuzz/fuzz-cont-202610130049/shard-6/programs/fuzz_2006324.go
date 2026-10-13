package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int }
func genFn0(p0 int16, p1 int) uint16 {
    v4 := []int{p1, p1}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v4
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 128
    var v3 uint = 2
    var v5 int64 = 1
    v6 := []int{32767}
    v3 = (^(v3))
    v2 = (v2 * v2)
    if (0.1 == 3.14159) {
        var v7 int16 = 437
        if false {
            var v8 uint32 = 2
            v7 += v7
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = v8
        }
        _ = v7
    } else {
        v9 := GenS1{F0: 520}
        v1["a"] = 4294967295
        v10 := GenS2{F0: 709}
        _ = v10
        _ = v9
    }
    v11 := v1[("a" + "foo bar")]
    if (float64((1.5 / 1.0)) < (float64(((float32((-2.25 * ((0.1 - (3.14159 + (0.0 + ((float64((-2.25 - (((1.5 * ((0.0 - 1e10) + (1.5 - (((3.14159 + float32(3.14159)) * 0.1) * 1e10)))) / 0.1) / ((float32(float64((1.5 + 3.14159))) + 0.0) * float32(1.5))))) * ((0.1 + -2.25) * (1e10 * 3.14159))) * 3.14159)))) * float64(0.1)))) / ((float32((3.14159 - 0.1)) * 3.14159) / 3.14159)) + 1e10)) / -2.25)) {
        _ = len(v1)
        v12 := GenS3{F0: (v2 % (((-(v2))) | 1))}
        v5 = v5
        _ = v12
    }
    _ = len(v6)
    v2 = (v2 >> ((v3) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
