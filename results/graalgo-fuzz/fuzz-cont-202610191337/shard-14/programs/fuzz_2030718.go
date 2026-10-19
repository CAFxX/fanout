package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint16 }
func genFn0(p0 uint16, p1 string, p2 bool) uint {
    for fi7 := 0; fi7 < 7; fi7++ {
        fmt.Println(p1)
        fmt.Println(p1, p0)
        switch fi7 {
        case 4294967295:
            p2 = false
        case -9223372036854775808:
            v8 := []int{fi7, fi7}
            _ = v8
        }
        _ = fi7
    }
    var v9 int16 = 32767
    v10 := GenS1{F0: v9, F1: 42, F2: (p0 % ((p0) | 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 2
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{470}
    var v3 uint8 = 255
    var v4 bool = (false || ((128 == 9223372036854775807) || false))
    var v5 uint32 = 126
    var v6 uint8 = v3
    fmt.Println(v3, v6)
    v5 += (v5 * (v5 % ((2147483647) | 1)))
    v4 = (-1 > -1)
    v2 = append(v2, (v5 & v5))
    v5 -= v5
    v3 -= (^(v6))
    _ = len(v1)
    v4 = (uint16(v5) <= uint16(v5))
    switch -814 {
    case 32768:
        v5 = 2
    case -1:
        if ("hello" > "a") {
            delete(v1, v5)
            var v11 float64 = 3.14159
            _ = v11
        } else {
            v2 = append(v2, (v5 + (v5 | (v5 % ((v5) | 1)))))
        }
    case 128:
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 bool = v4
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if false {
                var v14 uint16 = 20
                v3 = v6
                if false {
                    var v15 int8 = 1
                    _ = len(v1)
                    _ = v15
                } else {
                    fmt.Println(v3, v4, v13)
                    if v4 {
                        if (42 > 1) {
                            var v16 float64 = ((0.1 * (3.14159 + 1e10)) - (-2.25 / 0.1))
                            _ = v16
                        }
                        v17 := GenS1{F0: -808, F1: 214, F2: 2}
                        _ = v17
                    }
                    var v18 uint8 = (v3 ^ v3)
                    _ = v18
                }
                _ = v14
            }
            _ = fi12
            _ = v13
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
