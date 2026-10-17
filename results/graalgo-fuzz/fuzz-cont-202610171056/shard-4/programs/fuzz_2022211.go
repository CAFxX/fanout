package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint, p1 string, p2 uint64) int {
    fmt.Println(p0)
    switch 32767 {
    case 128:
        var v6 bool = (true && ((p0 % ((42) | 1)) < (p0 >> (((p0 << ((p0) & 63))) & 63))))
        _ = v6
    }
    if true {
        v7 := []int{32767, 127}
        _ = v7
    } else {
        v8 := []int{32768, 4294967295, 0, 2147483647}
        p2 -= p2
        _ = v8
    }
    var v9 uint32 = 1
    var v10 uint16 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 32768
}
func genFn1(p0 string, p1 string) uint32 {
    switch ("hello" + p0) {
    case "hello":
        p1 = p1
    case "\n\t":
        p0 = p0
    }
    p0 = p0
    _ = p0
    _ = p1
    return 2147483648
}

func main() {
    var v1 int8 = 1
    v2 := []int{-9223372036854775808, 292, -9223372036854775808, -9223372036854775808}
    v3 := make(map[string]int)
    v4 := []string{("" + "\n\t")}
    var v5 uint64 = 202
    if false {
        var v11 int32 = -980
        if (0 > -979) {
            fmt.Println(v5)
            if (!false) {
                fmt.Println(v1, v11)
                for fi12 := 0; fi12 < 15; fi12++ {
                    v11 -= v11
                    v13 := GenS1{F0: v11}
                    var v14 float32 = 0.0
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                for fi15 := 0; fi15 < 19; fi15++ {
                    var v16 float64 = (((1.5 * (1.5 - (1e10 + (((((float64((float32((1e10 * 0.1)) * 3.14159)) + 1e10) * -2.25) * float64(1e10)) * 0.1) * 0.0)))) - 1e10) * 1.5)
                    fmt.Println(v11, fi15, v1)
                    v16 = 0.0
                    _ = fi15
                    _ = v16
                }
            }
        } else {
            for fi17 := 0; fi17 < 10; fi17++ {
                v18 := GenS1{F0: (v11 | (v11 * -1))}
                _ = fi17
                _ = v18
            }
            fmt.Println(v1)
        }
        var v19 float64 = (((-2.25 - 3.14159) * 1.5) / 1.5)
        _ = v11
        _ = v19
    } else {
        v20 := GenS1{F0: 0}
        for fi21 := 0; fi21 < 4; fi21++ {
            fmt.Println(v1, v5)
            v5 += (v5 << ((253) & 63))
            if ((^(v1)) > v1) {
                v22 := GenS1{F0: 1}
                v1 = v1
                _ = v22
            }
            _ = fi21
        }
        _ = v20
    }
    fmt.Println(v5, v1)
    for fi23 := 0; fi23 < 3; fi23++ {
        v5 = 438
        v4 = append(v4, (("hello" + "\n\t") + ""))
        _ = fi23
    }
    fmt.Println(v5, v1)
    for fi24 := 0; fi24 < 13; fi24++ {
        v25 := GenS2{F0: 255}
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        v1 = (v1 | (^(v1)))
        _ = fi26
    }
    v1 += (^(v1))
    fmt.Println(v5)
    v5 = 2
    _ = len(v3)
    if len(v4) > 0 {
        v4[0] = "a"
    }
    _ = len(v4)
    for fi27 := 0; fi27 < 13; fi27++ {
        for fi28 := 0; fi28 < 15; fi28++ {
            v29 := GenS1{F0: 2147483647}
            _ = fi28
            _ = v29
        }
        fmt.Println(v1)
        fmt.Println(fi27, v5)
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
