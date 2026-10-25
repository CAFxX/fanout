package main

import (
    "fmt"
)

func genFn0(p0 uint64) int32 {
    var v6 int = 127
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int8 = -11
        var v9 uint16 = 904
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int32 = 1
    _ = p0
    _ = v10
    _ = v6
    return (v10 * v10)
}

func main() {
    var v1 uint32 = 2
    var v2 uint8 = 0
    v3 := make(map[string]int)
    var v4 float64 = ((float64(0.0) * ((float64(((float32(((-2.25 - 0.0) / (1e10 - 3.14159))) / 1e10) / 3.14159)) / 0.1) + float64((1.5 * 1e10)))) + float64(float32(((0.1 - (1e10 + 1.5)) * (0.0 + float64(((1.5 - 0.0) + (float32(-2.25) - float32(((((1.5 * 3.14159) + 1e10) - (1.5 - 0.0)) - 0.1))))))))))
    v5 := make(map[string]string)
    if (true || true) {
        v3[("a" + "\n\t")] = 999
    }
    var v11 uint32 = v1
    var v12 uint8 = v2
    var v13 int32 = 42
    fmt.Println(v2, v13)
    var v14 float64 = 1.5
    v4 = float64((((3.14159 / -2.25) + 3.14159) / 1.5))
    delete(v5, "a")
    v15 := []int{-2147483648, -663, 2147483647, 4294967295}
    v2 = v12
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
