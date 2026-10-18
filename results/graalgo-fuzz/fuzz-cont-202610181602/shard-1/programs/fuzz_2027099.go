package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int32) uint64 {
    v4 := GenS1{F0: p0}
    var v5 uint32 = 1
    v6 := GenS1{F0: (p0 / ((p0) | 1))}
    var v7 int16 = 435
    fmt.Println(p0)
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := []int{127, 2147483647, 582, 4294967295}
    v2 := make(map[int]int)
    var v3 uint64 = 4294967295
    v3 = (v3 % (((v3 + 832)) | 1))
    var v8 float64 = (((((float64((3.14159 * float32((0.1 / 1.0)))) - -2.25) / 1e10) - -2.25) * float64(float32((3.14159 - 0.1)))) + float64(3.14159))
    var v9 int16 = -509
    var v10 uint8 = 128
    var v11 float64 = 3.14159
    if ((1 != 575) && ((1 < -552) || (83 != 65535))) {
        v10 = 230
    }
    var v12 uint16 = 0
    if (4294967295 != 0) {
        var v13 int = -2147483648
        v2[v13] = -1
        _ = v13
    } else {
        var v14 float32 = (float32(3.14159) / (1.5 + -2.25))
        var v15 uint = 18446744073709551615
        _ = v14
        _ = v15
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
