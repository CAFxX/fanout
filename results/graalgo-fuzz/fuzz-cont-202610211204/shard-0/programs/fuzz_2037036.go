package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint16, p1 bool, p2 string) uint32 {
    v5 := GenS1{F0: 18446744073709551615}
    var v6 string = p2
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 65535
}

func main() {
    var v1 uint32 = 0
    var v2 uint = 255
    v3 := make(map[int]bool)
    var v4 uint32 = 0
    v4 = (v4 / (((v4 << (((v2 * v2)) & 63))) | 1))
    fmt.Println(v4, v4, v4)
    var v7 uint64 = 2
    if (!true) {
        v4 += 1
        v1 += (v1 - uint32(v2))
        v8 := []int{-2147483648}
        _ = v8
    } else {
        _ = len(v3)
        if true {
            switch -1 {
            case 32767:
                var v9 int32 = -2147483648
                _ = v9
            default:
                switch 4294967295 {
                case 925:
                    v4 -= 1
                case 42:
                    var v10 uint8 = 2
                    _ = v10
                case -1:
                    var v11 bool = (!false)
                    _ = v11
                default:
                    v7 += v7
                }
            }
        } else {
            if false {
                _ = len(v3)
                var v12 int16 = 32767
                var v13 uint64 = (v7 ^ (v7 ^ (v7 >> ((v2) & 63))))
                _ = v12
                _ = v13
            }
        }
        v2 = v2
    }
    var v14 string = "\n\t"
    for fi15 := 0; fi15 < 10; fi15++ {
        v7 = (v7 * v7)
        var v16 uint16 = 0
        _ = fi15
        _ = v16
    }
    var v17 string = ""
    v3[231] = (1 >= -946)
    var v18 int = 649
    v19 := []int{-9223372036854775808, 32767, v18, -2147483648}
    var v20 float32 = (float32(((float64((1.5 - (1e10 / 1e10))) / float64(-2.25)) / ((-2.25 - 3.14159) * float64(0.1)))) - 0.0)
    v3[42] = (127 >= 42)
    for fi21 := 0; fi21 < 12; fi21++ {
        v14 = v14
        _ = fi21
    }
    var v22 int32 = 1
    if (!true) {
        v17 = v14
        for fi23 := 0; fi23 < 16; fi23++ {
            var v24 uint32 = 0
            _ = fi23
            _ = v24
        }
    }
    switch 32767 {
    case 457:
        var v25 int32 = (v22 / ((v22) | 1))
        _ = v25
    }
    var v26 int32 = (v22 >> (((v2 << ((v2) & 63))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v7
}
