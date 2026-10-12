package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint32) int64 {
    v7 := GenS1{F0: 42}
    var v8 float32 = ((3.14159 * (3.14159 / (1.5 / (float32(1e10) - 3.14159)))) * ((0.0 + ((3.14159 * 1e10) * ((3.14159 + (-2.25 - (3.14159 * 3.14159))) * 0.0))) - 1.5))
    _ = p0
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 float64 = -2.25
    var v2 int32 = 1
    v3 := make(map[string]int)
    var v4 uint64 = 711
    var v5 int8 = -115
    v6 := []int{585, 2147483647, 127, 4294967295}
    var v9 int8 = v5
    v10 := GenS1{F0: 32767}
    v11 := GenS2{F0: 65535}
    var v12 int8 = (v5 << ((658) & 63))
    if ((v2 - v2) != 42) {
        var v13 uint8 = 0
        if (2 == 0) {
            v2 -= v2
            v1 = 1.5
            v1 = 3.14159
        }
        _ = v13
    }
    fmt.Println(v5, v1)
    fmt.Println(v12)
    fmt.Println(v1)
    switch "" {
    case "a":
        for fi14 := 0; fi14 < 20; fi14++ {
            v15 := v3[("\n\t" + "\n\t")]
            _ = fi14
            _ = v15
        }
    }
    v16 := GenS1{F0: -818}
    v5 = v9
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
