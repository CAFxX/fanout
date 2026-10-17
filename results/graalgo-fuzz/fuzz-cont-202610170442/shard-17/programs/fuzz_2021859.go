package main

import (
    "fmt"
)

func genFn0(p0 uint8) int {
    p0 -= p0
    var v7 uint = 324
    _ = p0
    _ = v7
    return -70
}
func genFn1(p0 uint8, p1 uint32) uint32 {
    p0 += (p0 & 128)
    if (-387 >= 233) {
        p1 += p1
        p0 += (^(p0))
    }
    _ = p0
    _ = p1
    return 740
}

func main() {
    var v1 int8 = -18
    var v2 float32 = 1.5
    var v3 bool = true
    var v4 int64 = -1
    v5 := make(map[string]bool)
    var v6 uint64 = 4294967295
    var v8 int64 = (v4 - v4)
    v8 += 9223372036854775807
    if v3 {
        fmt.Println(v3, v6, v2)
        v8 = (v8 ^ (v8 ^ v8))
        var v9 uint32 = 1
        _ = v9
    }
    v10 := []int{}
    if ((v2 <= 3.14159) || v3) {
        switch 128 {
        case -2147483648:
            var v11 bool = v3
            _ = v11
        }
    } else {
        delete(v5, "\n\t")
        var v12 float32 = (((v2 * v2) / (v2 / 1.0)) + ((float32((1.5 / ((-2.25 * (0.1 - 1.5)) / (-2.25 + 3.14159)))) / ((-2.25 * v2) * float32(0.0))) + (1.5 / 0.1)))
        _ = v12
    }
    _ = len(v10)
    var v13 uint8 = 2
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    var v14 int16 = -32768
    if v3 {
        v10 = append(v10, 32767)
    } else {
        var v15 bool = true
        fmt.Println(v1, v2)
        switch 128 {
        case 1:
            if (1 <= 2147483648) {
                v16 := v5["\n\t"]
                _ = v16
            }
        }
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
