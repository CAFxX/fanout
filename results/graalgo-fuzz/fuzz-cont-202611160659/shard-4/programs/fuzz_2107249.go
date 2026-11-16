package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
func genFn0() uint {
    var v6 uint8 = 255
    var v7 bool = true
    v6 = (v6 ^ v6)
    var v8 string = "\n\t"
    _ = v6
    _ = v7
    _ = v8
    return 255
}
func genFn1(p0 uint8, p1 int, p2 bool) uint {
    p2 = (245 >= 4294967295)
    var v9 uint64 = 1
    fmt.Println(p2, p0, v9)
    var v10 int16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return genFn0()
}

func main() {
    v1 := []uint32{4294967295, 0}
    v2 := make(map[string]int)
    v3 := make(map[int]int)
    var v4 uint64 = 1
    v5 := []int{790, 595}
    v4 = 42
    fmt.Println(v4)
    var v11 uint = 18446744073709551615
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 bool = true
        _ = fi12
        _ = v13
    }
    v3[9223372036854775807] = 1
    var v14 uint16 = 0
    var v15 uint = v11
    v16 := GenS1{F0: -1, F1: (v14 + 2)}
    if false {
        if ((127 == 32768) || (false || true)) {
            fmt.Println(v15, v11, v4)
            var v17 uint32 = 2147483648
            _ = v17
        } else {
            if (true && (!true)) {
                if (1 < -758) {
                    var v18 int64 = 979
                    v19 := GenS1{F0: 0, F1: uint16(v11)}
                    v14 = (^(v14))
                    _ = v18
                    _ = v19
                }
            } else {
                if false {
                    v20 := GenS1{F0: 183, F1: v14}
                    switch "" {
                    case "foo bar":
                        v14 = v14
                    }
                    fmt.Println(v14, v14, v15)
                    _ = v20
                } else {
                    v2[""] = -445
                }
            }
            v14 += v14
        }
        v1 = append(v1, uint32(v4))
    } else {
        v11 -= v11
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
