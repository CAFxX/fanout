package main

import (
    "fmt"
)

func genFn0(p0 bool, p1 uint16, p2 int32) uint {
    v5 := []int{-775, 128}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 753
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 65
    v3 := make(map[uint32]bool)
    var v4 uint8 = 128
    v2 = uint64(v4)
    delete(v3, uint32(v2))
    if ((!true) || true) {
        for fi6 := 0; fi6 < 6; fi6++ {
            v7 := []int{(fi6 >> ((uint(v2)) & 63)), (fi6 - fi6), (fi6 % ((fi6) | 1)), 128}
            v4 -= 1
            _ = fi6
            _ = v7
        }
    }
    var v8 uint32 = 65535
    var v9 string = "a"
    var v10 uint8 = uint8(v2)
    switch -868 {
    case -1:
        var v11 float64 = 1e10
        _ = v11
    case 32768:
        if (true || (((0.1 + ((3.14159 * -2.25) + (float32(-2.25) / -2.25))) <= 1.5) || ((v4 + (v4 + v10)) == 0))) {
            v9 = "\n\t"
            var v12 float32 = 3.14159
            _ = v12
        }
    case 1:
        var v13 bool = true
        _ = v13
    default:
        v8 = (v8 / (((^(v8))) | 1))
    }
    switch -9223372036854775808 {
    case -921:
        var v14 int64 = -295
        _ = v14
    case 437:
        v10 = 2
    case 128:
        v8 = 42
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        v4 -= v10
        v10 = v4
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 int16 = -855
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
