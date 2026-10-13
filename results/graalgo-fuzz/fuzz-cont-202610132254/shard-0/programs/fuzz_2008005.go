package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int32, p1 uint8, p2 uint16) int64 {
    if true {
        p0 = -2147483648
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := []int{fi7}
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v10 := []int{1, 32768, 128}
    p2 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 9223372036854775807
}

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 1
    var v3 int32 = 42
    v4 := []uint32{2147483647, 2147483648, 571}
    v5 := []uint32{42, 206, 72}
    v6 := make(map[int]bool)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v11 uint = 235
    v2 -= v2
    v12 := GenS1{F0: 42, F1: 413}
    for fi13 := 0; fi13 < 12; fi13++ {
        switch fi13 {
        case 4294967295:
            v14 := GenS1{F0: 909, F1: 2}
            _ = v14
        default:
            for fi15 := 0; fi15 < 12; fi15++ {
                var v16 float32 = 1e10
                _ = len(v5)
                _ = fi15
                _ = v16
            }
        }
        v2 += v2
        var v17 bool = (1 > v3)
        _ = fi13
        _ = v17
    }
    if (2147483647 >= 65535) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        delete(v6, -9223372036854775808)
    } else {
        v2 -= v2
        if len(v5) > 0 {
            v5[0] = 42
        }
        if ((("hello" + "hello") < "\n\t") || ((float64(1e10) / 1.0) == float64(1.5))) {
            var v18 int = 4294967295
            var v19 int32 = 2147483647
            _ = v18
            _ = v19
        }
    }
    _ = len(v5)
    v11 += 288
    v3 = (v3 / ((v3) | 1))
    fmt.Println(v2, v3)
    var v20 int16 = -481
    v11 = (^(v11))
    var v21 float32 = (0.1 * 3.14159)
    switch "hello" {
    case "":
        var v22 float32 = (float32(-2.25) * v21)
        _ = v22
    }
    var v23 int = 0
    fmt.Println(v23, v23)
    v21 = ((float32(-2.25) * v21) + 1e10)
    fmt.Println(v11, v21, v21)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
