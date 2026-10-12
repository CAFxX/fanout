package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 string }
func genFn0(p0 string, p1 int64, p2 int8) uint16 {
    var v6 bool = true
    for fi7 := 0; fi7 < 15; fi7++ {
        p1 += (p1 | p1)
        _ = fi7
    }
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 709
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{32767, -9223372036854775808}
    v3 := []int{32768}
    var v4 uint64 = 4294967295
    var v5 float32 = 0.1
    v8 := v1["hello"]
    v4 += (^(v4))
    var v9 int32 = 98
    var v10 bool = true
    if v10 {
        v9 = (-(v9))
        v9 += (v9 / (((v9 | v9)) | 1))
    }
    if (180 >= 128) {
        v10 = (-490 > v9)
        switch (v8 ^ (v8 - (v8 >> ((870) & 63)))) {
        case 2147483647:
            var v11 float32 = (v5 * (0.0 * v5))
            _ = v11
        case -911:
            v12 := GenS1{F0: 2147483647}
            _ = v12
        case 458:
            if (1 < 786) {
                v13 := GenS2{F0: int32(v4), F1: "\n\t"}
                _ = v13
            }
        default:
            v5 -= (((-2.25 / v5) - (((0.1 / 1.0) - (1.5 / -2.25)) * ((v5 / 3.14159) / (0.0 + (v5 - 0.0))))) * 0.1)
        }
        if ((((("\n\t" + "hello") + "a") + "foo bar") + "foo bar") == "") {
            if (-32768 >= 0) {
                v9 -= (^(v9))
                var v14 string = "hello"
                _ = v14
            }
            var v15 int8 = -1
            _ = v15
        }
    } else {
        v8 -= (v8 & (v8 % (((v8 * 4294967295)) | 1)))
        v8 = v8
    }
    v10 = (true && (!v10))
    v8 -= v8
    v9 = v9
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
