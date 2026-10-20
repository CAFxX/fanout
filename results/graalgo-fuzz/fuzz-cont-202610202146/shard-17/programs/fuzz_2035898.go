package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint }
func genFn0(p0 uint, p1 int8, p2 uint16) int {
    var v6 float32 = ((-2.25 * 0.0) - (1.5 * float32(-2.25)))
    if (!false) {
        if (!((!false) || (p1 == 127))) {
            v7 := GenS1{F0: 65535, F1: 4294967295, F2: (p0 % ((p0) | 1))}
            for fi8 := 0; fi8 < 17; fi8++ {
                fmt.Println(p2, fi8, p1)
                _ = fi8
            }
            if false {
                if ((!false) || true) {
                    for fi9 := 0; fi9 < 14; fi9++ {
                        var v10 int8 = p1
                        v11 := []int{fi9, (^(fi9)), (fi9 % ((fi9) | 1)), int(p2)}
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                }
            }
            _ = v7
        } else {
            fmt.Println(p0, p1, p2)
        }
        fmt.Println(p0, p2)
    }
    v12 := []int{32768, 32767}
    p2 += (p2 % ((p2) | 1))
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v6
    return 2147483647
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = ""
    var v3 uint = 65535
    var v4 uint16 = 550
    v5 := make(map[uint32]int)
    var v13 uint8 = 124
    var v14 float64 = 0.0
    v2 = "foo bar"
    v14 = -2.25
    v13 += (v13 * 2)
    v13 -= (v13 - v13)
    v15 := []int{128, 4294967295, 127, 127}
    fmt.Println(v13, v13)
    v16 := GenS1{F0: 2147483648, F1: 4294967295, F2: v3}
    fmt.Println(v13, v14)
    v3 += 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
