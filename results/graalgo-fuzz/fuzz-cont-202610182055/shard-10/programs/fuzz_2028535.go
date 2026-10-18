package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 string; F2 uint8 }
func genFn0(p0 bool, p1 uint) uint8 {
    p1 -= p1
    var v7 uint32 = 404
    var v8 bool = ((false && ((!p0) && p0)) || ((float64((1.5 - 3.14159)) + 1e10) < 0.1))
    var v9 float32 = 1e10
    var v10 int32 = -2147483648
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return uint8(p1)
}
func genFn1(p0 int, p1 int16) uint64 {
    p1 -= (p1 + p1)
    p1 -= (-(p1))
    _ = p0
    _ = p1
    return 18446744073709551615
}

func main() {
    var v1 string = "a"
    var v2 int64 = 1
    v3 := make(map[string]uint32)
    v4 := []int{}
    var v5 int = -1
    var v6 int64 = v2
    v5 += (v5 - 42)
    var v11 uint16 = 65535
    if (!false) {
        var v12 uint8 = 128
        switch "" {
        case "":
            v13 := GenS1{F0: v6}
            _ = v13
        default:
            if (false && true) {
                for fi14 := 0; fi14 < 1; fi14++ {
                    if true {
                        v5 -= 2147483647
                        if len(v4) > 0 {
                            v4[0] = -2147483648
                        }
                        var v15 int = v5
                        _ = v15
                    }
                    _ = fi14
                }
            } else {
                var v16 float32 = 1.5
                _ = v16
            }
        }
        for fi17 := 0; fi17 < 5; fi17++ {
            var v18 int64 = v6
            var v19 uint8 = v12
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = v12
    }
    if (364 == 1) {
        v20 := GenS1{F0: (v6 * (v6 | (^(v2))))}
        _ = v20
    }
    var v21 uint32 = 1
    v22 := GenS2{F0: 255, F1: "\n\t", F2: genFn0(false, 18446744073709551615)}
    v23 := GenS1{F0: v2}
    fmt.Println(v6, v11, v5)
    for fi24 := 0; fi24 < 13; fi24++ {
        var v25 int32 = 2147483647
        var v26 int = (fi24 + v5)
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
