package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint32, p1 uint64) string {
    fmt.Println(p1, p0)
    v5 := []int{}
    var v6 uint64 = 2
    if true {
        for fi7 := 0; fi7 < 14; fi7++ {
            v5 = append(v5, int(p1))
            fmt.Println(p0, p0)
            _ = fi7
        }
        var v8 float64 = float64(3.14159)
        _ = v8
    } else {
        p1 = v6
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 uint16 = 836
            _ = fi9
            _ = v10
        }
        if len(v5) > 0 {
            v5[0] = 128
        }
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return ("a" + "")
}

func main() {
    v1 := []int{32768, 2147483647, 32767}
    var v2 uint8 = 2
    var v3 int32 = 2147483647
    var v4 int32 = (^(v3))
    v2 -= v2
    var v11 int = 775
    if len(v1) > 0 {
        v1[0] = -2147483648
    }
    v12 := GenS1{F0: 1}
    fmt.Println(v4, v3)
    v11 += (v11 % ((2147483647) | 1))
    var v13 float64 = 3.14159
    if (float32((((float64((0.0 / 3.14159)) + 1.5) + v13) / float64(1.5))) >= -2.25) {
        var v14 int = (v11 * 209)
        v3 -= int32(v2)
        _ = v14
    }
    switch 32767 {
    case -2147483648:
        var v15 int64 = -297
        _ = v15
    case -1:
        v4 += v3
    case 356:
        switch (^(v11)) {
        case 0:
            var v16 int32 = (v3 * (v4 * v4))
            _ = v16
        case 127:
            v11 += 1
        default:
            var v17 uint = 1
            _ = v17
        }
    }
    fmt.Println(v2)
    var v18 uint = uint(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
