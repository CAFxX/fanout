package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int8, p1 int64, p2 int8) uint8 {
    p0 -= p0
    var v5 float64 = 0.1
    p1 = (p1 - p1)
    p0 += p2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = true
    var v3 int64 = -66
    var v4 int8 = -128
    if v2 {
        v6 := GenS1{F0: 1}
        fmt.Println(v4, v3, v4)
        _ = v6
    } else {
        for fi7 := 0; fi7 < 4; fi7++ {
            v3 += v3
            if v2 {
                v4 += (v4 - -1)
                fmt.Println(v2, v3)
            } else {
                v8 := GenS2{F0: -1}
                var v9 int8 = v4
                if ((-736 != 0) || (42 < 42)) {
                    for fi10 := 0; fi10 < 3; fi10++ {
                        fmt.Println(v4, v4)
                        v4 = int8(fi10)
                        var v11 uint32 = 2147483648
                        _ = fi10
                        _ = v11
                    }
                    var v12 uint = 4294967295
                    v1[0] = (v4 > v4)
                    _ = v12
                }
                _ = v8
                _ = v9
            }
            _ = fi7
        }
        v3 = v3
    }
    v13 := []int{4294967295, 95, 127, 2147483647}
    var v14 int = -2147483648
    for fi15 := 0; fi15 < 11; fi15++ {
        fmt.Println(v2, v2)
        _ = fi15
    }
    fmt.Println(v4, v2, v3)
    var v16 uint = 337
    var v17 int16 = 382
    v17 -= int16(v16)
    v1[2147483647] = (0 != genFn0((v4 & 100), -9223372036854775808, (^(v4))))
    var v18 bool = v2
    v13 = append(v13, v14)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
