package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 uint, p1 int8, p2 uint16) uint32 {
    p1 = p1
    v4 := []int{-9223372036854775808, 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 790
}

func main() {
    var v1 float64 = ((float64((float32(3.14159) / 3.14159)) / 1.0) - 3.14159)
    v2 := []int{}
    v3 := []int{1}
    if (42 >= 1) {
        var v5 uint16 = 797
        v6 := GenS1{F0: 128}
        _ = v5
        _ = v6
    }
    v1 += (v1 / (float64((((((3.14159 + ((-2.25 - 0.1) / float32((v1 - float64(((float32(0.0) - 1e10) + float32((v1 * v1)))))))) / (1.5 / 3.14159)) / float32(0.1)) * float32(float64(1e10))) - 3.14159)) / float64((-2.25 + 1e10))))
    var v7 uint = 4294967295
    var v8 int64 = -225
    v8 = v8
    v1 += ((float64(1e10) / (v1 - (v1 + -2.25))) + v1)
    v1 -= ((((((v1 - (((v1 + 0.1) + ((float64(0.0) * (0.1 + -2.25)) - v1)) + 3.14159)) - 1e10) + (v1 * ((v1 - (v1 / (v1 + v1))) - v1))) + 3.14159) / float64(0.1)) + v1)
    v7 = v7
    v9 := GenS1{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
