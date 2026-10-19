package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 bool) int {
    p0 = true
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := []int{fi7, (fi7 / (((-(fi7))) | 1)), fi7}
        _ = fi7
        _ = v8
    }
    p0 = true
    v9 := []int{32768, 2147483647, -9223372036854775808}
    if (!(-1 < 964)) {
        p0 = (float64(0.1) <= -2.25)
        var v10 float64 = 0.0
        fmt.Println(v10)
        _ = v10
    } else {
        v9 = append(v9, 4294967295)
    }
    _ = p0
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 int32 = 0
    var v2 uint8 = 128
    var v3 int8 = -128
    var v4 uint = 1
    v5 := []string{"\n\t", "a", "a", ""}
    v6 := make(map[uint32]uint32)
    if (458 == int16(v2)) {
        v4 += v4
    }
    v11 := GenS1{F0: -600}
    fmt.Println(v4, v3)
    v2 -= v2
    v3 -= 19
    var v12 uint = v4
    var v13 int = 128
    if ((-1 == -294) || true) {
        var v14 int32 = v1
        switch "hello" {
        case "a":
            var v15 uint64 = 652
            _ = v15
        }
        if true {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
