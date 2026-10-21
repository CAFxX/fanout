package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int32; F1 uint64; F2 uint }
func genFn0(p0 string, p1 uint64, p2 uint64) string {
    p1 = 1
    p1 += 763
    v6 := GenS1{F0: 0, F1: 2147483647}
    if (p1 < (p1 % (((p1 >> ((2) & 63))) | 1))) {
        var v7 uint = 4294967295
        if (!false) {
            p0 = p0
        } else {
            var v8 int64 = -1
            for fi9 := 0; fi9 < 7; fi9++ {
                if false {
                    for fi10 := 0; fi10 < 15; fi10++ {
                        p1 -= 2
                        var v11 uint = (v7 | v7)
                        _ = fi10
                        _ = v11
                    }
                } else {
                    fmt.Println(fi9)
                    v8 -= v8
                }
                v12 := GenS1{F0: 127, F1: fi9}
                _ = fi9
                _ = v12
            }
            var v13 uint8 = 254
            _ = v13
            _ = v8
        }
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p0
}
func genFn1(p0 int, p1 int64) uint8 {
    fmt.Println(p1)
    v14 := []int{p0}
    _ = p0
    _ = p1
    _ = v14
    return 2
}

func main() {
    v1 := []string{"", "\n\t", ("" + "")}
    var v2 float64 = 1.5
    v3 := make(map[int]string)
    var v4 uint8 = 255
    var v5 uint32 = 1
    if false {
        v4 += v4
    }
    var v15 int = -2147483648
    fmt.Println(v4)
    v15 -= (v15 | (v15 % ((-9223372036854775808) | 1)))
    var v16 int8 = -74
    v17 := GenS2{F0: 42, F1: 2, F2: 65535}
    v16 = (v16 * (v16 & (^(v16))))
    v2 = v2
    var v18 int64 = int64(v15)
    for fi19 := 0; fi19 < 8; fi19++ {
        v20 := GenS2{F0: -2147483648, F1: 0, F2: 329}
        var v21 uint16 = 32767
        v4 += genFn1((v15 / ((v15) | 1)), v18)
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 uint32 = (v5 + 2)
    var v23 int64 = v18
    fmt.Println(v18)
    delete(v3, v15)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
