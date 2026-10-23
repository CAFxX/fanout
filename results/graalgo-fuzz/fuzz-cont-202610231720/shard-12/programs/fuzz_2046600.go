package main

import (
    "fmt"
)

func genFn0() int16 {
    var v5 string = "\n\t"
    var v6 float32 = (((1e10 / (0.1 / -2.25)) * 1.5) / (0.0 - ((0.1 - -2.25) - 0.0)))
    v7 := []int{803, -389}
    _ = len(v7)
    _ = v5
    _ = v6
    _ = v7
    return 1000
}

func main() {
    v1 := []int{128, -2147483648, -2147483648, 0}
    v2 := make(map[int]string)
    var v3 float32 = 0.0
    var v4 uint = 1
    var v8 float64 = 1e10
    var v9 float32 = v3
    v9 = float32(v8)
    switch 128 {
    case -1:
        for fi10 := 0; fi10 < 18; fi10++ {
            switch (-(fi10)) {
            case 2147483647:
                v4 -= 65535
            }
            _ = fi10
        }
    case 2147483647:
        fmt.Println(v3, v9, v3)
    case 32767:
        v8 += 3.14159
    }
    delete(v2, 32768)
    v9 -= v3
    v3 += (v9 * 3.14159)
    fmt.Println(v4, v4)
    switch -860 {
    case -747:
        var v11 uint8 = 22
        _ = v11
    default:
        fmt.Println(v9, v3)
    }
    if (1 < 0) {
        var v12 bool = true
        var v13 float64 = float64(v3)
        _ = v12
        _ = v13
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
