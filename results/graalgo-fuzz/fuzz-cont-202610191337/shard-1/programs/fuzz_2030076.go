package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint }
func genFn0(p0 uint32, p1 uint8) uint64 {
    p0 -= (p0 * 900)
    var v4 float32 = (((3.14159 * (0.0 + ((0.1 * 0.0) + (1.5 / 3.14159)))) * 1e10) + ((1.5 - 1e10) + (0.1 + float32((1.5 + 1e10)))))
    _ = p0
    _ = p1
    _ = v4
    return 0
}
func genFn1(p0 uint8, p1 string, p2 uint64) uint {
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := []int{}
        v7 := GenS1{F0: -27, F1: (p2 >> ((uint(p2)) & 63)), F2: 18446744073709551615}
        fmt.Println(p1, p2, fi5)
        _ = fi5
        _ = v6
        _ = v7
    }
    p2 -= genFn0(1, 16)
    _ = p0
    _ = p1
    _ = p2
    return 65535
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = 1.5
    var v3 float32 = v2
    var v8 uint32 = 143
    _ = len(v1)
    switch "hello" {
    case "foo bar":
        if true {
            v9 := []int{929, 2147483647, 127}
            v9 = append(v9, -2147483648)
            _ = v9
        } else {
            v3 = (1.5 * v2)
            switch -911 {
            case 128:
                switch 145 {
                case 5:
                    v2 = float32(float64(v2))
                case 2147483647:
                    switch -1 {
                    case 4294967295:
                        fmt.Println(v3, v8)
                    case 42:
                        if true {
                            var v10 uint64 = 463
                            _ = v10
                        }
                    case 128:
                        v8 = (^(v8))
                    }
                }
            case 2147483647:
                fmt.Println(v8)
            default:
                v8 += (v8 + v8)
            }
        }
    default:
        v11 := []int{23, 1, -1, 4294967295}
        _ = v11
    }
    var v12 int16 = 443
    if ((0 != -9223372036854775808) || (true || true)) {
        v8 += (v8 << ((genFn1(127, "hello", genFn0(v8, 122))) & 63))
    } else {
        v2 = v3
        if (127 != 1) {
            v12 += v12
            for fi13 := 0; fi13 < 1; fi13++ {
                v8 = 42
                _ = fi13
            }
        }
        v14 := GenS1{F0: 127, F1: 18446744073709551615, F2: 4294967295}
        _ = v14
    }
    fmt.Println(v12, v12, v8)
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 uint8 = 42
        v16 = v16
        _ = fi15
        _ = v16
    }
    v17 := []int{}
    var v18 uint64 = 2
    for fi19 := 0; fi19 < 14; fi19++ {
        var v20 uint64 = v18
        _ = fi19
        _ = v20
    }
    v17 = append(v17, 1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v8
}
